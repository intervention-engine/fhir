package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func ObservationIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Observation
	c := Database.C("observations")

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
			if splitKey[0] == "subject" {
				err := c.Find(bson.M{"subject.referenceid": value[0]}).All(&result)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}

	var observationEntryList []models.BundleEntryComponent
	for _, observation := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &observation
		observationEntryList = append(observationEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = observationEntryList

	log.Println("Setting observation search context")
	context.Set(r, "Observation", result)
	context.Set(r, "Resource", "Observation")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadObservation(r *http.Request) (*models.Observation, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("observations")
	result := models.Observation{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting observation read context")
	context.Set(r, "Observation", result)
	context.Set(r, "Resource", "Observation")
	return &result, nil
}

func ObservationShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadObservation(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Observation"))
}

func ObservationCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	observation := &models.Observation{}
	err := decoder.Decode(observation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("observations")
	i := bson.NewObjectId()
	observation.Id = i.Hex()
	err = c.Insert(observation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting observation create context")
	context.Set(r, "Observation", observation)
	context.Set(r, "Resource", "Observation")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Observation/"+i.Hex())
}

func ObservationUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	observation := &models.Observation{}
	err := decoder.Decode(observation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("observations")
	observation.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, observation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting observation update context")
	context.Set(r, "Observation", observation)
	context.Set(r, "Resource", "Observation")
	context.Set(r, "Action", "update")
}

func ObservationDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("observations")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting observation delete context")
	context.Set(r, "Observation", id.Hex())
	context.Set(r, "Resource", "Observation")
	context.Set(r, "Action", "delete")
}
