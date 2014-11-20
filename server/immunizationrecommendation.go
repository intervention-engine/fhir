package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gitlab.mitre.org/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func ImmunizationRecommendationIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.ImmunizationRecommendation
	c := Database.C("immunizationrecommendations")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var bundle models.ImmunizationRecommendationBundle
	bundle.Type = "Bundle"
	bundle.Title = "ImmunizationRecommendation Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entries = result

	log.Println("Setting immunizationrecommendation search context")
	context.Set(r, "ImmunizationRecommendation", result)
	context.Set(r, "Resource", "ImmunizationRecommendation")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func ImmunizationRecommendationShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("immunizationrecommendations")

	result := models.ImmunizationRecommendation{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting immunizationrecommendation read context")
	context.Set(r, "ImmunizationRecommendation", result)
	context.Set(r, "Resource", "ImmunizationRecommendation")
	context.Set(r, "Action", "read")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(result)
}

func ImmunizationRecommendationCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	immunizationrecommendation := &models.ImmunizationRecommendation{}
	err := decoder.Decode(immunizationrecommendation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("immunizationrecommendations")
	i := bson.NewObjectId()
	immunizationrecommendation.Id = i.Hex()
	err = c.Insert(immunizationrecommendation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting immunizationrecommendation create context")
	context.Set(r, "ImmunizationRecommendation", immunizationrecommendation)
	context.Set(r, "Resource", "ImmunizationRecommendation")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":8080/ImmunizationRecommendation/"+i.Hex())
}

func ImmunizationRecommendationUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	immunizationrecommendation := &models.ImmunizationRecommendation{}
	err := decoder.Decode(immunizationrecommendation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("immunizationrecommendations")
	immunizationrecommendation.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, immunizationrecommendation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting immunizationrecommendation update context")
	context.Set(r, "ImmunizationRecommendation", immunizationrecommendation)
	context.Set(r, "Resource", "ImmunizationRecommendation")
	context.Set(r, "Action", "update")
}

func ImmunizationRecommendationDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("immunizationrecommendations")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting immunizationrecommendation delete context")
	context.Set(r, "ImmunizationRecommendation", id.Hex())
	context.Set(r, "Resource", "ImmunizationRecommendation")
	context.Set(r, "Action", "delete")
}
