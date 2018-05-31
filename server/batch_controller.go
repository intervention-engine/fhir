package server

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/fhir/search"
)

// BatchController handles FHIR batch operations via input bundles
type BatchController struct {
	DAL    DataAccessLayer
	Config Config
}

// NewBatchController creates a new BatchController based on the passed in DAL
func NewBatchController(dal DataAccessLayer, config Config) *BatchController {
	return &BatchController{
		DAL:    dal,
		Config: config,
	}
}

// Post processes and incoming batch request
func (b *BatchController) Post(c *gin.Context) {
	bundle := &models.Bundle{}
	err := FHIRBind(c, bundle)
	if err != nil {
		outcome := &models.OperationOutcome{
			Issue: []models.OperationOutcomeIssueComponent{
				models.OperationOutcomeIssueComponent{
					Severity: "fatal", // fatal means "The issue caused the action to fail, and no further checking could be performed."
					Code:     "structure",
					Diagnostics: err.Error(),
				},
			},
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, outcome)
		return
	}

	// TODO: If type is batch, ensure there are no interdependent resources

	// Loop through the entries, ensuring they have a request and that we support the method,
	// while also creating a new entries array that can be sorted by method.
	entries := make([]*models.BundleEntryComponent, len(bundle.Entry))
	for i := range bundle.Entry {
		if bundle.Entry[i].Request == nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("Entries in a batch operation require a request"))
			return
		}

		switch bundle.Entry[i].Request.Method {
		default:
			c.AbortWithError(http.StatusNotImplemented,
				errors.New("Operation currently unsupported in batch requests: "+bundle.Entry[i].Request.Method))
			return
		case "DELETE":
			if bundle.Entry[i].Request.Url == "" {
				c.AbortWithError(http.StatusBadRequest, errors.New("Batch DELETE must have a URL"))
				return
			}
		case "POST":
			if bundle.Entry[i].Resource == nil {
				c.AbortWithError(http.StatusBadRequest, errors.New("Batch POST must have a resource body"))
				return
			}
		case "PUT":
			if bundle.Entry[i].Resource == nil {
				c.AbortWithError(http.StatusBadRequest, errors.New("Batch PUT must have a resource body"))
				return
			}
			if !strings.Contains(bundle.Entry[i].Request.Url, "/") && !strings.Contains(bundle.Entry[i].Request.Url, "?") {
				c.AbortWithError(http.StatusBadRequest, errors.New("Batch PUT url must have an id or a condition"))
				return
		}
		}
		entries[i] = &bundle.Entry[i]
	}

	sort.Sort(byRequestMethod(entries))

	// Now loop through the entries, assigning new IDs to those that are POST or Conditional PUT and fixing any
	// references to reference the new ID.
	refMap := make(map[string]models.Reference)
	newIDs := make([]string, len(entries))
	createStatus := make([]string, len(entries))
	for i, entry := range entries {
		if entry.Request.Method == "POST" {

			id := ""
			
			if len(entry.Request.IfNoneExist) > 0 {
				// Conditional Create
				query := search.Query{Resource: entry.Request.Url, Query: entry.Request.IfNoneExist}
				existingIds, err := b.DAL.FindIDs(query)
				if err != nil {
					c.AbortWithError(http.StatusInternalServerError, err)
					return
				}

				if len(existingIds) == 0 {
					createStatus[i] = "201"
				} else if len(existingIds) == 1 {
					createStatus[i] = "200"
					id = existingIds[0]
				} else if len(existingIds) > 1 {
					createStatus[i] = "412" // HTTP 412 - Precondition Failed
				}
			} else {
				// Unconditional create
				createStatus[i] = "201"
			}

			if createStatus[i] == "201" {
				// Create a new ID
				id = bson.NewObjectId().Hex()
			newIDs[i] = id
			}

			if len(id) > 0 {
				// Add id to the reference map
			refMap[entry.FullUrl] = models.Reference{
				Reference:    entry.Request.Url + "/" + id,
				Type:         entry.Request.Url,
				ReferencedID: id,
				External:     new(bool),
			}
			// Rewrite the FullUrl using the new ID
			entry.FullUrl = responseURL(c.Request, b.Config, entry.Request.Url, id).String()
			}

		} else if entry.Request.Method == "PUT" && isConditional(entry) {
			// We need to process conditionals referencing temp IDs in a second pass, so skip them here
			if hasTempID(entry.Request.Url) {
				continue
			}

			if err := b.resolveConditionalPut(c.Request, i, entry, newIDs, refMap); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
	}

	// Second pass to take care of conditionals referencing temporary IDs.  Known limitation: if a conditional
	// references a temp ID also defined by a conditional, we error out if it hasn't been resolved yet -- too many
	// rabbit holes.
	for i, entry := range entries {
		if entry.Request.Method == "PUT" && isConditional(entry) {
			// Use a regex to swap out the temp IDs with the new IDs
			for oldID, ref := range refMap {
				re := regexp.MustCompile("([=,])(" + oldID + "|" + url.QueryEscape(oldID) + ")(&|,|$)")
				entry.Request.Url = re.ReplaceAllString(entry.Request.Url, "${1}"+ref.Reference+"${3}")
			}

			if hasTempID(entry.Request.Url) {
				c.AbortWithError(http.StatusNotImplemented,
					errors.New("Cannot resolve conditionals referencing other conditionals"))
				return
			}

			if err := b.resolveConditionalPut(c.Request, i, entry, newIDs, refMap); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
	}

	// Update all the references to the entries (to reflect newly assigned IDs)
	updateAllReferences(entries, refMap)

	// Then make the changes in the database and update the entry response
	for i, entry := range entries {
		switch entry.Request.Method {
		case "DELETE":
			if !isConditional(entry) {
				// It's a normal DELETE
				parts := strings.SplitN(entry.Request.Url, "/", 2)
				if len(parts) != 2 {
					c.AbortWithError(http.StatusInternalServerError,
						fmt.Errorf("Couldn't identify resource and id to delete from %s", entry.Request.Url))
					return
				}
				if err := b.DAL.Delete(parts[1], parts[0]); err != nil && err != ErrNotFound {
					c.AbortWithError(http.StatusInternalServerError, err)
					return
				}
			} else {
				// It's a conditional (query-based) delete
				parts := strings.SplitN(entry.Request.Url, "?", 2)
				query := search.Query{Resource: parts[0], Query: parts[1]}
				if _, err := b.DAL.ConditionalDelete(query); err != nil {
					c.AbortWithError(http.StatusInternalServerError, err)
					return
				}
			}

			entry.Request = nil
			entry.Response = &models.BundleEntryResponseComponent{
				Status: "204",
			}
		case "POST":

			entry.Response = &models.BundleEntryResponseComponent{
				Status:   createStatus[i],
				Location: entry.FullUrl,
			}

			if createStatus[i] == "201" {
				// creating
			if err := b.DAL.PostWithID(newIDs[i], entry.Resource); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
				if meta, ok := models.GetResourceMeta(entry.Resource); ok {
					entry.Response.LastModified = meta.LastUpdated
				}
			} else if createStatus[i] == "200" {
				// have one existing resource
				components := strings.Split(entry.FullUrl, "/")
				existingId := components[len(components)-1]

				existingResource, err := b.DAL.Get(existingId, entry.Request.Url)
				if err != nil {
					c.AbortWithError(http.StatusInternalServerError, err)
					return
			}
				entry.Resource = existingResource
				if meta, ok := models.GetResourceMeta(existingResource); ok {
					if meta != nil && meta.LastUpdated != nil {
				entry.Response.LastModified = meta.LastUpdated
			}
				}
			} else if createStatus[i] == "412" {
				entry.Response.Outcome = &models.OperationOutcome{
					Issue: []models.OperationOutcomeIssueComponent{
						models.OperationOutcomeIssueComponent{
							Severity: "warning",
							Code:     "duplicate",
							Diagnostics: "search criteria were not selective enough",
						},
					},
				}
			}
			entry.Request = nil

		case "PUT":
			// Because we pre-process conditional PUTs, we know this is always a normal PUT operation
			entry.FullUrl = responseURL(c.Request, b.Config, entry.Request.Url).String()
			parts := strings.SplitN(entry.Request.Url, "/", 2)
			if len(parts) != 2 {
				c.AbortWithError(http.StatusInternalServerError,
					fmt.Errorf("Couldn't identify resource and id to put from %s", entry.Request.Url))
				return
			}
			createdNew, err := b.DAL.Put(parts[1], entry.Resource)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			entry.Request = nil
			entry.Response = new(models.BundleEntryResponseComponent)
			entry.Response.Location = entry.FullUrl
			if createdNew {
				entry.Response.Status = "201"
			} else {
				entry.Response.Status = "200"
			}
			if meta, ok := models.GetResourceMeta(entry.Resource); ok {
				entry.Response.LastModified = meta.LastUpdated
			}
		}
	}

	total := uint32(len(entries))
	bundle.Total = &total
	bundle.Type = fmt.Sprintf("%s-response", bundle.Type)

	c.Set("Bundle", bundle)
	c.Set("Resource", "Bundle")
	c.Set("Action", "batch")

	// Send the response

	c.Header("Access-Control-Allow-Origin", "*")
	if c.GetBool("SendXML") {
		converterInt := c.MustGet("FhirFormatConverter")
		converter := converterInt.(*FhirFormatConverter)
		converter.SendXML(bundle, c)
	} else {
		c.JSON(http.StatusOK, bundle)
	}
}

func (b *BatchController) resolveConditionalPut(request *http.Request, entryIndex int, entry *models.BundleEntryComponent, newIDs []string, refMap map[string]models.Reference) error {
	// Do a preflight to either get the existing ID, get a new ID, or detect multiple matches (not allowed)
	parts := strings.SplitN(entry.Request.Url, "?", 2)
	query := search.Query{Resource: parts[0], Query: parts[1]}

	var id string
	if IDs, err := b.DAL.FindIDs(query); err == nil {
		switch len(IDs) {
		case 0:
			id = bson.NewObjectId().Hex()
		case 1:
			id = IDs[0]
		default:
			return ErrMultipleMatches
		}
	} else {
		return err
	}

	// Rewrite the PUT as a normal (non-conditional) PUT
	entry.Request.Url = query.Resource + "/" + id

	// Add the new ID to the reference map
	newIDs[entryIndex] = id
	refMap[entry.FullUrl] = models.Reference{
		Reference:    entry.Request.Url,
		Type:         query.Resource,
		ReferencedID: id,
		External:     new(bool),
	}

	// Rewrite the FullUrl using the new ID
	entry.FullUrl = responseURL(request, b.Config, entry.Request.Url, id).String()

	return nil
}

func updateAllReferences(entries []*models.BundleEntryComponent, refMap map[string]models.Reference) {
	// First, get all the references by reflecting through the fields of each model
	var refs []*models.Reference
	for _, entry := range entries {
		model := entry.Resource
		if model != nil {
			entryRefs := findRefsInValue(reflect.ValueOf(model))
			refs = append(refs, entryRefs...)
		}
	}
	// Then iterate through and update as necessary
	for _, ref := range refs {
		newRef, found := refMap[ref.Reference]
		if found {
			*ref = newRef
		}
	}
}

func findRefsInValue(val reflect.Value) []*models.Reference {
	var refs []*models.Reference

	// Dereference pointers in order to simplify things
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Make sure it's a valid thing, else return right away
	if !val.IsValid() {
		return refs
	}

	// Handle it if it's a ref, otherwise iterate its members for refs
	if val.Type() == reflect.TypeOf(models.Reference{}) {
		refs = append(refs, val.Addr().Interface().(*models.Reference))
	} else if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			subRefs := findRefsInValue(val.Field(i))
			refs = append(refs, subRefs...)
		}
	} else if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			subRefs := findRefsInValue(val.Index(i))
			refs = append(refs, subRefs...)
		}
	}

	return refs
}

func isConditional(entry *models.BundleEntryComponent) bool {
	if entry.Request == nil {
		return false
	} else if entry.Request.Method != "PUT" && entry.Request.Method != "DELETE" {
		return false
	}
	return !strings.Contains(entry.Request.Url, "/") || strings.Contains(entry.Request.Url, "?")
}

func hasTempID(str string) bool {

	// do not match URLs like Patient?identifier=urn:oid:0.1.2.3.4.5.6.7|urn:uuid:6002c2ab-9571-4db7-9a79-87163475b071
	tempIdRegexp := regexp.MustCompile("([=,])(urn:uuid:|urn%3Auuid%3A)[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}(&|,|$)")
	matches := tempIdRegexp.MatchString(str)

	// hasPrefix := strings.HasPrefix(str, "urn:uuid:") || strings.HasPrefix(str, "urn%3Auuid%3A")
	// contains := strings.Contains(str, "urn:uuid:") || strings.Contains(str, "urn%3Auuid%3A")
	// if matches != contains {
		// fmt.Printf("re != contains (re = %t): %s\n", matches, str)
	// }

	return matches
}

// Support sorting by request method, as defined in the spec
type byRequestMethod []*models.BundleEntryComponent

func (e byRequestMethod) Len() int {
	return len(e)
}
func (e byRequestMethod) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e byRequestMethod) Less(i, j int) bool {
	methodMap := map[string]int{"DELETE": 0, "POST": 1, "PUT": 2, "GET": 3}
	return methodMap[e[i].Request.Method] < methodMap[e[j].Request.Method]
}
