package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/fhir/search"
	"gopkg.in/mgo.v2/bson"
)

func ProcedureRequestIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if r := recover(); r != nil {
			rw.Header().Set("Content-Type", "application/json; charset=utf-8")
			switch x := r.(type) {
			case search.SearchError:
				rw.WriteHeader(x.HTTPStatus())
				json.NewEncoder(rw).Encode(x.OperationOutcome())
				return
			default:
				e := search.InternalServerError(fmt.Sprintf("%s", x))
				rw.WriteHeader(e.HTTPStatus())
				json.NewEncoder(rw).Encode(e.OperationOutcome())
			}
		}
	}()

	var result []models.ProcedureRequest
	c := Database.C("procedurerequests")

	r.ParseForm()
	if len(r.Form) == 0 {
		iter := c.Find(nil).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	} else {
		searcher := search.NewMongoSearcher(Database)
		query := search.Query{Resource: "ProcedureRequest", Query: r.URL.RawQuery}
		err := searcher.CreateQuery(query).All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}

	var procedurerequestEntryList []models.BundleEntryComponent
	for i := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &result[i]
		procedurerequestEntryList = append(procedurerequestEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = procedurerequestEntryList

	log.Println("Setting procedurerequest search context")
	context.Set(r, "ProcedureRequest", result)
	context.Set(r, "Resource", "ProcedureRequest")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadProcedureRequest(r *http.Request) (*models.ProcedureRequest, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("procedurerequests")
	result := models.ProcedureRequest{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting procedurerequest read context")
	context.Set(r, "ProcedureRequest", result)
	context.Set(r, "Resource", "ProcedureRequest")
	return &result, nil
}

func ProcedureRequestShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadProcedureRequest(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "ProcedureRequest"))
}

func ProcedureRequestCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	procedurerequest := &models.ProcedureRequest{}
	err := decoder.Decode(procedurerequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("procedurerequests")
	i := bson.NewObjectId()
	procedurerequest.Id = i.Hex()
	err = c.Insert(procedurerequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting procedurerequest create context")
	context.Set(r, "ProcedureRequest", procedurerequest)
	context.Set(r, "Resource", "ProcedureRequest")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Add("Location", "http://"+host+":3001/ProcedureRequest/"+i.Hex())
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(procedurerequest)
}

func ProcedureRequestUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	procedurerequest := &models.ProcedureRequest{}
	err := decoder.Decode(procedurerequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("procedurerequests")
	procedurerequest.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, procedurerequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting procedurerequest update context")
	context.Set(r, "ProcedureRequest", procedurerequest)
	context.Set(r, "Resource", "ProcedureRequest")
	context.Set(r, "Action", "update")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(procedurerequest)
}

func ProcedureRequestDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("procedurerequests")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting procedurerequest delete context")
	context.Set(r, "ProcedureRequest", id.Hex())
	context.Set(r, "Resource", "ProcedureRequest")
	context.Set(r, "Action", "delete")
}
