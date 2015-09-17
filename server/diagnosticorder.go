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

func DiagnosticOrderIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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

	var result []models.DiagnosticOrder
	c := Database.C("diagnosticorders")

	r.ParseForm()
	if len(r.Form) == 0 {
		iter := c.Find(nil).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	} else {
		searcher := search.NewMongoSearcher(Database)
		query := search.Query{Resource: "DiagnosticOrder", Query: r.URL.RawQuery}
		err := searcher.CreateQuery(query).All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}

	var diagnosticorderEntryList []models.BundleEntryComponent
	for i := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &result[i]
		diagnosticorderEntryList = append(diagnosticorderEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = diagnosticorderEntryList

	log.Println("Setting diagnosticorder search context")
	context.Set(r, "DiagnosticOrder", result)
	context.Set(r, "Resource", "DiagnosticOrder")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadDiagnosticOrder(r *http.Request) (*models.DiagnosticOrder, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("diagnosticorders")
	result := models.DiagnosticOrder{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting diagnosticorder read context")
	context.Set(r, "DiagnosticOrder", result)
	context.Set(r, "Resource", "DiagnosticOrder")
	return &result, nil
}

func DiagnosticOrderShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadDiagnosticOrder(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "DiagnosticOrder"))
}

func DiagnosticOrderCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	diagnosticorder := &models.DiagnosticOrder{}
	err := decoder.Decode(diagnosticorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("diagnosticorders")
	i := bson.NewObjectId()
	diagnosticorder.Id = i.Hex()
	err = c.Insert(diagnosticorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting diagnosticorder create context")
	context.Set(r, "DiagnosticOrder", diagnosticorder)
	context.Set(r, "Resource", "DiagnosticOrder")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Add("Location", "http://"+host+":3001/DiagnosticOrder/"+i.Hex())
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(diagnosticorder)
}

func DiagnosticOrderUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	diagnosticorder := &models.DiagnosticOrder{}
	err := decoder.Decode(diagnosticorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("diagnosticorders")
	diagnosticorder.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, diagnosticorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting diagnosticorder update context")
	context.Set(r, "DiagnosticOrder", diagnosticorder)
	context.Set(r, "Resource", "DiagnosticOrder")
	context.Set(r, "Action", "update")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(diagnosticorder)
}

func DiagnosticOrderDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("diagnosticorders")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting diagnosticorder delete context")
	context.Set(r, "DiagnosticOrder", id.Hex())
	context.Set(r, "Resource", "DiagnosticOrder")
	context.Set(r, "Action", "delete")
}
