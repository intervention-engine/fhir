package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func ImmunizationIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Immunization
	c := Database.C("immunizations")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var immunizationEntryList []models.ImmunizationBundleEntry
	for _, immunization := range result {
		var entry models.ImmunizationBundleEntry
		entry.Title = "Immunization " + immunization.Id
		entry.Id = immunization.Id
		entry.Content = immunization
		immunizationEntryList = append(immunizationEntryList, entry)
	}

	var bundle models.ImmunizationBundle
	bundle.Type = "Bundle"
	bundle.Title = "Immunization Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = immunizationEntryList

	log.Println("Setting immunization search context")
	context.Set(r, "Immunization", result)
	context.Set(r, "Resource", "Immunization")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadImmunization(r *http.Request) (*models.Immunization, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("immunizations")
	result := models.Immunization{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting immunization read context")
	context.Set(r, "Immunization", result)
	context.Set(r, "Resource", "Immunization")
	return &result, nil
}

func ImmunizationShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadImmunization(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Immunization"))
}

func ImmunizationCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	immunization := &models.Immunization{}
	err := decoder.Decode(immunization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("immunizations")
	i := bson.NewObjectId()
	immunization.Id = i.Hex()
	err = c.Insert(immunization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting immunization create context")
	context.Set(r, "Immunization", immunization)
	context.Set(r, "Resource", "Immunization")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Immunization/"+i.Hex())
}

func ImmunizationUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	immunization := &models.Immunization{}
	err := decoder.Decode(immunization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("immunizations")
	immunization.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, immunization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting immunization update context")
	context.Set(r, "Immunization", immunization)
	context.Set(r, "Resource", "Immunization")
	context.Set(r, "Action", "update")
}

func ImmunizationDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("immunizations")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting immunization delete context")
	context.Set(r, "Immunization", id.Hex())
	context.Set(r, "Resource", "Immunization")
	context.Set(r, "Action", "delete")
}
