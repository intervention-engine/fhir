package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func ConceptMapIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.ConceptMap
	c := Database.C("conceptmaps")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var conceptmapEntryList []models.BundleEntryComponent
	for _, conceptmap := range result {
		var entry models.BundleEntryComponent
		entry.Resource = conceptmap
		conceptmapEntryList = append(conceptmapEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = conceptmapEntryList

	log.Println("Setting conceptmap search context")
	context.Set(r, "ConceptMap", result)
	context.Set(r, "Resource", "ConceptMap")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadConceptMap(r *http.Request) (*models.ConceptMap, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("conceptmaps")
	result := models.ConceptMap{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting conceptmap read context")
	context.Set(r, "ConceptMap", result)
	context.Set(r, "Resource", "ConceptMap")
	return &result, nil
}

func ConceptMapShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadConceptMap(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "ConceptMap"))
}

func ConceptMapCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	conceptmap := &models.ConceptMap{}
	err := decoder.Decode(conceptmap)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("conceptmaps")
	i := bson.NewObjectId()
	conceptmap.Id = i.Hex()
	err = c.Insert(conceptmap)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting conceptmap create context")
	context.Set(r, "ConceptMap", conceptmap)
	context.Set(r, "Resource", "ConceptMap")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/ConceptMap/"+i.Hex())
}

func ConceptMapUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	conceptmap := &models.ConceptMap{}
	err := decoder.Decode(conceptmap)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("conceptmaps")
	conceptmap.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, conceptmap)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting conceptmap update context")
	context.Set(r, "ConceptMap", conceptmap)
	context.Set(r, "Resource", "ConceptMap")
	context.Set(r, "Action", "update")
}

func ConceptMapDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("conceptmaps")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting conceptmap delete context")
	context.Set(r, "ConceptMap", id.Hex())
	context.Set(r, "Resource", "ConceptMap")
	context.Set(r, "Action", "delete")
}
