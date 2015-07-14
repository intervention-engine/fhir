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

func SubstanceIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Substance
	c := Database.C("substances")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var substanceEntryList []models.BundleEntryComponent
	for _, substance := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &substance
		substanceEntryList = append(substanceEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = substanceEntryList

	log.Println("Setting substance search context")
	context.Set(r, "Substance", result)
	context.Set(r, "Resource", "Substance")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadSubstance(r *http.Request) (*models.Substance, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("substances")
	result := models.Substance{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting substance read context")
	context.Set(r, "Substance", result)
	context.Set(r, "Resource", "Substance")
	return &result, nil
}

func SubstanceShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadSubstance(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Substance"))
}

func SubstanceCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	substance := &models.Substance{}
	err := decoder.Decode(substance)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("substances")
	i := bson.NewObjectId()
	substance.Id = i.Hex()
	err = c.Insert(substance)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting substance create context")
	context.Set(r, "Substance", substance)
	context.Set(r, "Resource", "Substance")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Substance/"+i.Hex())
}

func SubstanceUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	substance := &models.Substance{}
	err := decoder.Decode(substance)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("substances")
	substance.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, substance)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting substance update context")
	context.Set(r, "Substance", substance)
	context.Set(r, "Resource", "Substance")
	context.Set(r, "Action", "update")
}

func SubstanceDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("substances")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting substance delete context")
	context.Set(r, "Substance", id.Hex())
	context.Set(r, "Resource", "Substance")
	context.Set(r, "Action", "delete")
}
