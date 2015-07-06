package upload

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/intervention-engine/fhir/models"
)

/*
 * NOTE: This is a destructive operation.  Resources will be updated with new server-assigned ID and
 * its references will point to server locations of other resources.
 */
func UploadResources(resources []interface{}, baseURL string) (map[string]string, error) {
	refMap := make(map[string]string)
	sortResourcesByDependency(resources)
	for _, t := range resources {
		err := updateReferences(t, refMap)
		if err != nil {
			return refMap, err
		}

		// Remember the old ID so we can create a map from old ID to new location
		oldId := getId(t)

		// Upload the resource and get the new location
		_, err = UploadResource(t, baseURL)
		if err != nil {
			return refMap, err
		}

		// Add entry to map from the old ID to the new relative path
		if oldId != "" {
			refMap[oldId] = reflect.TypeOf(t).Elem().Name() + "/" + getId(t)
		}
	}

	return refMap, nil
}

/*
 * NOTE: This is a destructive operation.  The resources will be updated with new server-assigned ID.
 */
func UploadResource(resource interface{}, baseURL string) (string, error) {
	// FYI: We can post resource w/ bogus id because it doesn't get serialized
	json, _ := json.Marshal(resource)
	body := bytes.NewReader(json)
	url := baseURL + "/" + reflect.TypeOf(resource).Elem().Name()
	response, err := http.Post(url, "application/json+fhir", body)
	if err != nil {
		return "", err
	}
	loc := response.Header.Get("Location")

	regs := []string{".*/([^/]+)/_history/.*", ".*/([^/]+)"}
	for _, reg := range regs {
		creg := regexp.MustCompile(reg)
		if matches := creg.FindStringSubmatch(loc); matches != nil {
			setId(resource, matches[1])
			break
		}
	}

	return loc, nil
}

func updateReferences(resource interface{}, refMap map[string]string) error {
	refs := getAllReferences(resource)
	for _, ref := range refs {
		if err := updateReference(ref, refMap); err != nil {
			return err
		}
	}
	return nil
}

func updateReference(ref *models.Reference, refMap map[string]string) error {
	if ref != nil && strings.HasPrefix(ref.Reference, "cid:") {
		newRef, ok := refMap[strings.TrimPrefix(ref.Reference, "cid:")]
		if ok {
			ref.Reference = newRef
		} else {
			return errors.New(fmt.Sprint("Failed to find updated reference for ", ref))
		}
	}

	return nil
}

func getAllReferences(model interface{}) []*models.Reference {
	refs := make([]*models.Reference, 0)
	s := reflect.ValueOf(model).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.Type() == reflect.TypeOf(&models.Reference{}) && !f.IsNil() {
			refs = append(refs, f.Interface().(*models.Reference))
		} else if f.Type() == reflect.TypeOf([]models.Reference{}) {
			for j := 0; j < f.Len(); j++ {
				refs = append(refs, f.Index(j).Addr().Interface().(*models.Reference))
			}
		}
	}
	return refs
}

func getId(model interface{}) string {
	return reflect.ValueOf(model).Elem().FieldByName("Id").String()
}

func setId(model interface{}, id string) {
	v := reflect.ValueOf(model).Elem().FieldByName("Id")
	if v.CanSet() {
		v.SetString(id)
	}
}

type ByDependency []interface{}

func (d ByDependency) Len() int {
	return len(d)
}
func (d ByDependency) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d ByDependency) Less(i, j int) bool {
	for _, ref := range getAllReferences(d[i]) {
		if strings.TrimPrefix(ref.Reference, "cid:") == getId(d[j]) {
			return false
		}
	}
	return true
}

func sortResourcesByDependency(resources []interface{}) {
	sort.Sort(ByDependency(resources))
}
