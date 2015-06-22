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

func SpecimenIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Specimen
	c := Database.C("specimens")

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

	var specimenEntryList []models.SpecimenBundleEntry
	for _, specimen := range result {
		var entry models.SpecimenBundleEntry
		entry.Title = "Specimen " + specimen.Id
		entry.Id = specimen.Id
		entry.Content = specimen
		specimenEntryList = append(specimenEntryList, entry)
	}

	var bundle models.SpecimenBundle
	bundle.Type = "Bundle"
	bundle.Title = "Specimen Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = specimenEntryList

	log.Println("Setting specimen search context")
	context.Set(r, "Specimen", result)
	context.Set(r, "Resource", "Specimen")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadSpecimen(r *http.Request) (*models.Specimen, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("specimens")
	result := models.Specimen{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting specimen read context")
	context.Set(r, "Specimen", result)
	context.Set(r, "Resource", "Specimen")
	return &result, nil
}

func SpecimenShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadSpecimen(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Specimen"))
}

func SpecimenCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	specimen := &models.Specimen{}
	err := decoder.Decode(specimen)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("specimens")
	i := bson.NewObjectId()
	specimen.Id = i.Hex()
	err = c.Insert(specimen)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting specimen create context")
	context.Set(r, "Specimen", specimen)
	context.Set(r, "Resource", "Specimen")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Specimen/"+i.Hex())
}

func SpecimenUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	specimen := &models.Specimen{}
	err := decoder.Decode(specimen)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("specimens")
	specimen.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, specimen)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting specimen update context")
	context.Set(r, "Specimen", specimen)
	context.Set(r, "Resource", "Specimen")
	context.Set(r, "Action", "update")
}

func SpecimenDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("specimens")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting specimen delete context")
	context.Set(r, "Specimen", id.Hex())
	context.Set(r, "Resource", "Specimen")
	context.Set(r, "Action", "delete")
}
