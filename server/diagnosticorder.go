package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func DiagnosticOrderIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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
		for key, value := range r.Form {
			splitKey := strings.Split(key, ":")
			if (len(splitKey) > 1) && (splitKey[0] == "subject") {
				err := c.Find(bson.M{"subject.referenceid": value[0]}).All(&result)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}

	var diagnosticorderEntryList []models.DiagnosticOrderBundleEntry
	for _, diagnosticorder := range result {
		var entry models.DiagnosticOrderBundleEntry
		entry.Title = "DiagnosticOrder " + diagnosticorder.Id
		entry.Id = diagnosticorder.Id
		entry.Content = diagnosticorder
		diagnosticorderEntryList = append(diagnosticorderEntryList, entry)
	}

	var bundle models.DiagnosticOrderBundle
	bundle.Type = "Bundle"
	bundle.Title = "DiagnosticOrder Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = diagnosticorderEntryList

	log.Println("Setting diagnosticorder search context")
	context.Set(r, "DiagnosticOrder", result)
	context.Set(r, "Resource", "DiagnosticOrder")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
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
