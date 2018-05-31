package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/fhir/search"
)

// ResourceController provides the necessary CRUD handlers for a given resource.
type ResourceController struct {
	Name   string
	DAL    DataAccessLayer
	Config Config
}

// NewResourceController creates a new resource controller for the passed in resource name and the passed in
// DataAccessLayer.
func NewResourceController(name string, dal DataAccessLayer, config Config) *ResourceController {
	return &ResourceController{
		Name:   name,
		DAL:    dal,
		Config: config,
	}
}

// IndexHandler handles requests to list resource instances or search for them.
func (rc *ResourceController) IndexHandler(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case *search.Error:
				c.Render(x.HTTPStatus, CustomFhirRenderer{x.OperationOutcome, c})
				return
			default:
				outcome := models.NewOperationOutcome("fatal", "exception", "")
				c.Render(http.StatusInternalServerError, CustomFhirRenderer{outcome, c})
				return
			}
		}
	}()

	searchQuery := search.Query{Resource: rc.Name, Query: c.Request.URL.RawQuery}
	baseURL := responseURL(c.Request, rc.Config, rc.Name)
	bundle, err := rc.DAL.Search(*baseURL, searchQuery)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set("bundle", bundle)
	c.Set("Resource", rc.Name)
	c.Set("Action", "search")

	c.Render(http.StatusOK, CustomFhirRenderer{bundle, c})
}

// LoadResource uses the resource id in the request to get a resource from the DataAccessLayer and store it in the
// context.
func (rc *ResourceController) LoadResource(c *gin.Context) (interface{}, error) {
	result, err := rc.DAL.Get(c.Param("id"), rc.Name)
	if err != nil {
		return nil, err
	}

	c.Set(rc.Name, result)
	c.Set("Resource", rc.Name)
	return result, nil
}

// ShowHandler handles requests to get a particular resource by ID.
func (rc *ResourceController) ShowHandler(c *gin.Context) {
	c.Set("Action", "read")
	resource, err := rc.LoadResource(c)
	if err != nil && err != ErrNotFound {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err == ErrNotFound {
		c.Status(http.StatusNotFound)
		return
	}
	c.Render(http.StatusOK, CustomFhirRenderer{resource, c})
}

// EverythingHandler handles requests for everything related to a Patient or Encounter resource.
func (rc *ResourceController) EverythingHandler(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case *search.Error:
				c.Render(x.HTTPStatus, CustomFhirRenderer{x.OperationOutcome, c})
				return
			default:
				outcome := models.NewOperationOutcome("fatal", "exception", "")
				c.Render(http.StatusInternalServerError, CustomFhirRenderer{outcome, c})
				return
			}
		}
	}()

	// For now we interpret $everything as the union of _include and _revinclude
	query := fmt.Sprintf("_id=%s&_include=*&_revinclude=*", c.Param("id"))

	searchQuery := search.Query{Resource: rc.Name, Query: query}
	baseURL := responseURL(c.Request, rc.Config, rc.Name)
	bundle, err := rc.DAL.Search(*baseURL, searchQuery)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set("bundle", bundle)
	c.Set("Resource", rc.Name)
	c.Set("Action", "search")

	c.Render(http.StatusOK, CustomFhirRenderer{bundle, c})
}

// CreateHandler handles requests to create a new resource instance, assigning it a new ID.
func (rc *ResourceController) CreateHandler(c *gin.Context) {
	resource := models.NewStructForResourceName(rc.Name)
	err := FHIRBind(c, resource)
	if err != nil {
		oo := models.NewOperationOutcome("fatal", "structure", err.Error())
		c.Render(http.StatusBadRequest, CustomFhirRenderer{oo, c})
		return
	}

	// check for conditional create
	ifNoneExist := c.GetHeader("If-None-Exist")
	var httpStatus int
	var id string
	if len(ifNoneExist) > 0 {
		query := search.Query{Resource: rc.Name, Query: ifNoneExist}
		httpStatus, id, resource, err = rc.DAL.ConditionalPost(query, resource)
	} else {
		httpStatus = http.StatusCreated
		id, err = rc.DAL.Post(resource)
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set(rc.Name, resource)
	c.Set("Resource", rc.Name)
	c.Set("Action", "create")

	if len(id) > 0 {
		c.Header("Location", responseURL(c.Request, rc.Config, rc.Name, id).String())
	}
	c.Render(httpStatus, CustomFhirRenderer{resource, c})
}

// UpdateHandler handles requests to update a resource having a given ID.  If the resource with that ID does not
// exist, a new resource is created with that ID.
func (rc *ResourceController) UpdateHandler(c *gin.Context) {
	resource := models.NewStructForResourceName(rc.Name)
	err := FHIRBind(c, resource)
	if err != nil {
		oo := models.NewOperationOutcome("fatal", "structure", err.Error())
		c.Render(http.StatusBadRequest, CustomFhirRenderer{oo, c})
		return
	}

	createdNew, err := rc.DAL.Put(c.Param("id"), resource)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set(rc.Name, resource)
	c.Set("Resource", rc.Name)

	c.Header("Location", responseURL(c.Request, rc.Config, rc.Name, c.Param("id")).String())
	if createdNew {
		c.Set("Action", "create")
		c.Render(http.StatusCreated, CustomFhirRenderer{resource, c})
	} else {
		c.Set("Action", "update")
		c.Render(http.StatusOK, CustomFhirRenderer{resource, c})
	}
}

// ConditionalUpdateHandler handles requests for conditional updates.  These requests contain search criteria for the
// resource to update.  If the criteria results in no found resources, a new resource is created.  If the criteria
// results in one found resource, that resource will be updated.  Criteria resulting in more than one found resource
// is considered an error.
func (rc *ResourceController) ConditionalUpdateHandler(c *gin.Context) {
	resource := models.NewStructForResourceName(rc.Name)
	err := FHIRBind(c, resource)
	if err != nil {
		oo := models.NewOperationOutcome("fatal", "structure", err.Error())
		c.Render(http.StatusBadRequest, CustomFhirRenderer{oo, c})
		return
	}

	query := search.Query{Resource: rc.Name, Query: c.Request.URL.RawQuery}
	id, createdNew, err := rc.DAL.ConditionalPut(query, resource)
	if err == ErrMultipleMatches {
		c.AbortWithStatus(http.StatusPreconditionFailed)
		return
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set("Resource", rc.Name)

	c.Header("Location", responseURL(c.Request, rc.Config, rc.Name, id).String())
	if createdNew {
		c.Set("Action", "create")
		c.Render(http.StatusCreated, CustomFhirRenderer{resource, c})
	} else {
		c.Set("Action", "update")
		c.Render(http.StatusOK, CustomFhirRenderer{resource, c})
	}
}

// DeleteHandler handles requests to delete a resource instance identified by its ID.
func (rc *ResourceController) DeleteHandler(c *gin.Context) {
	id := c.Param("id")

	if err := rc.DAL.Delete(id, rc.Name); err != nil && err != ErrNotFound {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set(rc.Name, id)
	c.Set("Resource", rc.Name)
	c.Set("Action", "delete")

	c.Status(http.StatusNoContent)
}

// ConditionalDeleteHandler handles requests to delete resources identified by search criteria.  All resources
// matching the search criteria will be deleted.
func (rc *ResourceController) ConditionalDeleteHandler(c *gin.Context) {
	query := search.Query{Resource: rc.Name, Query: c.Request.URL.RawQuery}
	_, err := rc.DAL.ConditionalDelete(query)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set("Resource", rc.Name)
	c.Set("Action", "delete")

	c.Status(http.StatusNoContent)
}

func responseURL(r *http.Request, config Config, paths ...string) *url.URL {

	if config.ServerURL != "" {
		theURL := fmt.Sprintf("%s/%s", strings.TrimSuffix(config.ServerURL, "/"), strings.Join(paths, "/"))
		responseURL, err := url.Parse(theURL)

		if err == nil {
			return responseURL
		}
	}

	responseURL := url.URL{}

	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		responseURL.Scheme = "https"
	} else {
		responseURL.Scheme = "http"
	}
	responseURL.Host = r.Host
	responseURL.Path = fmt.Sprintf("/%s", strings.Join(paths, "/"))

	return &responseURL
}

// CustomFhirRenderer replaces gin's default JSON renderer and ensures
// that the special characters "<", ">", and "&" are not escaped after the
// the JSON is marshaled. Escaping these special HTML characters is the default
// behavior of Go's json.Marshal().
// It also outputs XML if that is required
type CustomFhirRenderer struct {
	obj interface{}
	c   *gin.Context
}

var fhirJSONContentType = []string{"application/fhir+json; charset=utf-8"}
var fhirXMLContentType = []string{"application/fhir+xml; charset=utf-8"}

func (u CustomFhirRenderer) Render(w http.ResponseWriter) (err error) {

	if u.obj == nil {
		w.Write([]byte(""))
		return
	}

	data, err := json.Marshal(&u.obj)
	if err != nil {
		return
	}

	if u.c.GetBool("SendXML") {
		converterInt := u.c.MustGet("FhirFormatConverter")
		converter := converterInt.(*FhirFormatConverter)
		var xml string
		xml, err = converter.JsonToXml(string(data))
		if err != nil {
			fmt.Printf("ERROR: JsonToXml failed for data: %+v %s\n", u.obj, string(data))
			return
		}
		writeContentType(w, fhirXMLContentType)
		_, err = w.Write([]byte(xml))
	} else {
		// Replace the escaped characters in the data
		data = bytes.Replace(data, []byte("\\u003c"), []byte("<"), -1)
		data = bytes.Replace(data, []byte("\\u003e"), []byte(">"), -1)
		data = bytes.Replace(data, []byte("\\u0026"), []byte("&"), -1)

		writeContentType(w, fhirJSONContentType)
		_, err = w.Write(data)
	}
	return
}

func (u CustomFhirRenderer) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, fhirJSONContentType)
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
