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

func EpisodeOfCareIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.EpisodeOfCare
	c := Database.C("episodeofcares")

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
			if splitKey[0] == "patient" {
				err := c.Find(bson.M{"patient.referenceid": value[0]}).All(&result)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}

	var episodeofcareEntryList []models.BundleEntryComponent
	for _, episodeofcare := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &episodeofcare
		episodeofcareEntryList = append(episodeofcareEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = episodeofcareEntryList

	log.Println("Setting episodeofcare search context")
	context.Set(r, "EpisodeOfCare", result)
	context.Set(r, "Resource", "EpisodeOfCare")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadEpisodeOfCare(r *http.Request) (*models.EpisodeOfCare, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("episodeofcares")
	result := models.EpisodeOfCare{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting episodeofcare read context")
	context.Set(r, "EpisodeOfCare", result)
	context.Set(r, "Resource", "EpisodeOfCare")
	return &result, nil
}

func EpisodeOfCareShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadEpisodeOfCare(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "EpisodeOfCare"))
}

func EpisodeOfCareCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	episodeofcare := &models.EpisodeOfCare{}
	err := decoder.Decode(episodeofcare)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("episodeofcares")
	i := bson.NewObjectId()
	episodeofcare.Id = i.Hex()
	err = c.Insert(episodeofcare)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting episodeofcare create context")
	context.Set(r, "EpisodeOfCare", episodeofcare)
	context.Set(r, "Resource", "EpisodeOfCare")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/EpisodeOfCare/"+i.Hex())
}

func EpisodeOfCareUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	episodeofcare := &models.EpisodeOfCare{}
	err := decoder.Decode(episodeofcare)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("episodeofcares")
	episodeofcare.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, episodeofcare)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting episodeofcare update context")
	context.Set(r, "EpisodeOfCare", episodeofcare)
	context.Set(r, "Resource", "EpisodeOfCare")
	context.Set(r, "Action", "update")
}

func EpisodeOfCareDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("episodeofcares")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting episodeofcare delete context")
	context.Set(r, "EpisodeOfCare", id.Hex())
	context.Set(r, "Resource", "EpisodeOfCare")
	context.Set(r, "Action", "delete")
}
