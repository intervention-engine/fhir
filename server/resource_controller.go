package server

// TODO: This code can and should be cleaned up.  For now, it is more or less a port of the code that used to exist
// for every resource controller.

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"reflect"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/fhir/search"
	"gopkg.in/mgo.v2/bson"
)

type ResourceController struct {
	Name string
}

func (rc *ResourceController) IndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if r := recover(); r != nil {
			rw.Header().Set("Content-Type", "application/json; charset=utf-8")
			switch x := r.(type) {
			case search.Error:
				rw.WriteHeader(x.HTTPStatus)
				json.NewEncoder(rw).Encode(x.OperationOutcome)
				return
			default:
				outcome := &models.OperationOutcome{
					Issue: []models.OperationOutcomeIssueComponent{
						models.OperationOutcomeIssueComponent{
							Severity: "fatal",
							Code:     "exception",
						},
					},
				}
				rw.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(rw).Encode(outcome)
			}
		}
	}()

	result := models.NewSliceForResourceName(rc.Name, 0, 0)

	// Create and execute the Mongo query based on the http query params
	searcher := search.NewMongoSearcher(Database)
	searchQuery := search.Query{Resource: rc.Name, Query: r.URL.RawQuery}
	mgoQuery := searcher.CreateQuery(searchQuery)
	err := mgoQuery.All(result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var entryList []models.BundleEntryComponent
	resultVal := reflect.ValueOf(result).Elem()
	for i := 0; i < resultVal.Len(); i++ {
		var entry models.BundleEntryComponent
		entry.Resource = resultVal.Index(i).Addr().Interface()
		entryList = append(entryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Entry = entryList

	// Need to get the true total (not just how many were returned in this response)
	options := searchQuery.Options()
	var total uint32
	if resultVal.Len() == options.Count || resultVal.Len() == 0 {
		// Need to get total count from the server, since there may be more or the offset was too high
		intTotal, err := searcher.CreateQueryWithoutOptions(searchQuery).Count()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		total = uint32(intTotal)
	} else {
		// We can figure out the total by adding the offset and # results returned
		total = uint32(options.Offset + resultVal.Len())
	}
	bundle.Total = &total

	context.Set(r, rc.Name, reflect.ValueOf(result).Elem().Interface())
	context.Set(r, "Resource", rc.Name)
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func (rc *ResourceController) LoadResource(r *http.Request) (interface{}, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C(models.PluralizeLowerResourceName(rc.Name))
	result := models.NewStructForResourceName(rc.Name)
	err := c.Find(bson.M{"_id": id.Hex()}).One(result)
	if err != nil {
		return nil, err
	}

	context.Set(r, rc.Name, result)
	context.Set(r, "Resource", rc.Name)
	return result, nil
}

func (rc *ResourceController) ShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := rc.LoadResource(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, rc.Name))
}

func (rc *ResourceController) CreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	resource := models.NewStructForResourceName(rc.Name)
	err := decoder.Decode(resource)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C(models.PluralizeLowerResourceName(rc.Name))
	i := bson.NewObjectId()
	reflect.ValueOf(resource).Elem().FieldByName("Id").SetString(i.Hex())
	err = c.Insert(resource)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	context.Set(r, rc.Name, resource)
	context.Set(r, "Resource", rc.Name)
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Add("Location", "http://"+host+":3001/"+rc.Name+"/"+i.Hex())
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(resource)
}

func (rc *ResourceController) UpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	resource := models.NewStructForResourceName(rc.Name)
	err := decoder.Decode(resource)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C(models.PluralizeLowerResourceName(rc.Name))
	reflect.ValueOf(resource).Elem().FieldByName("Id").SetString(id.Hex())
	err = c.Update(bson.M{"_id": id.Hex()}, resource)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	context.Set(r, rc.Name, resource)
	context.Set(r, "Resource", rc.Name)
	context.Set(r, "Action", "update")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(resource)
}

func (rc *ResourceController) DeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C(models.PluralizeLowerResourceName(rc.Name))

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	context.Set(r, rc.Name, id.Hex())
	context.Set(r, "Resource", rc.Name)
	context.Set(r, "Action", "delete")
}
