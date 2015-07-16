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

func QuestionnaireIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Questionnaire
	c := Database.C("questionnaires")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var questionnaireEntryList []models.BundleEntryComponent
	for _, questionnaire := range result {
		var entry models.BundleEntryComponent
		entry.Resource = questionnaire
		questionnaireEntryList = append(questionnaireEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = questionnaireEntryList

	log.Println("Setting questionnaire search context")
	context.Set(r, "Questionnaire", result)
	context.Set(r, "Resource", "Questionnaire")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadQuestionnaire(r *http.Request) (*models.Questionnaire, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("questionnaires")
	result := models.Questionnaire{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting questionnaire read context")
	context.Set(r, "Questionnaire", result)
	context.Set(r, "Resource", "Questionnaire")
	return &result, nil
}

func QuestionnaireShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadQuestionnaire(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Questionnaire"))
}

func QuestionnaireCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	questionnaire := &models.Questionnaire{}
	err := decoder.Decode(questionnaire)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("questionnaires")
	i := bson.NewObjectId()
	questionnaire.Id = i.Hex()
	err = c.Insert(questionnaire)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting questionnaire create context")
	context.Set(r, "Questionnaire", questionnaire)
	context.Set(r, "Resource", "Questionnaire")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Questionnaire/"+i.Hex())
}

func QuestionnaireUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	questionnaire := &models.Questionnaire{}
	err := decoder.Decode(questionnaire)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("questionnaires")
	questionnaire.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, questionnaire)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting questionnaire update context")
	context.Set(r, "Questionnaire", questionnaire)
	context.Set(r, "Resource", "Questionnaire")
	context.Set(r, "Action", "update")
}

func QuestionnaireDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("questionnaires")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting questionnaire delete context")
	context.Set(r, "Questionnaire", id.Hex())
	context.Set(r, "Resource", "Questionnaire")
	context.Set(r, "Action", "delete")
}
