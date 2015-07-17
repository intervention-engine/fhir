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

func QuestionnaireAnswersIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.QuestionnaireAnswers
	c := Database.C("questionnaireanswerss")

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

	var questionnaireanswersEntryList []models.BundleEntryComponent
	for i := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &result[i]
		questionnaireanswersEntryList = append(questionnaireanswersEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = questionnaireanswersEntryList

	log.Println("Setting questionnaireanswers search context")
	context.Set(r, "QuestionnaireAnswers", result)
	context.Set(r, "Resource", "QuestionnaireAnswers")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadQuestionnaireAnswers(r *http.Request) (*models.QuestionnaireAnswers, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("questionnaireanswerss")
	result := models.QuestionnaireAnswers{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting questionnaireanswers read context")
	context.Set(r, "QuestionnaireAnswers", result)
	context.Set(r, "Resource", "QuestionnaireAnswers")
	return &result, nil
}

func QuestionnaireAnswersShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadQuestionnaireAnswers(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "QuestionnaireAnswers"))
}

func QuestionnaireAnswersCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	questionnaireanswers := &models.QuestionnaireAnswers{}
	err := decoder.Decode(questionnaireanswers)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("questionnaireanswerss")
	i := bson.NewObjectId()
	questionnaireanswers.Id = i.Hex()
	err = c.Insert(questionnaireanswers)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting questionnaireanswers create context")
	context.Set(r, "QuestionnaireAnswers", questionnaireanswers)
	context.Set(r, "Resource", "QuestionnaireAnswers")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/QuestionnaireAnswers/"+i.Hex())
}

func QuestionnaireAnswersUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	questionnaireanswers := &models.QuestionnaireAnswers{}
	err := decoder.Decode(questionnaireanswers)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("questionnaireanswerss")
	questionnaireanswers.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, questionnaireanswers)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting questionnaireanswers update context")
	context.Set(r, "QuestionnaireAnswers", questionnaireanswers)
	context.Set(r, "Resource", "QuestionnaireAnswers")
	context.Set(r, "Action", "update")
}

func QuestionnaireAnswersDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("questionnaireanswerss")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting questionnaireanswers delete context")
	context.Set(r, "QuestionnaireAnswers", id.Hex())
	context.Set(r, "Resource", "QuestionnaireAnswers")
	context.Set(r, "Action", "delete")
}
