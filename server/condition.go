package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func ConditionIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Condition
	c := Database.C("conditions")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	r.ParseForm()
	if (len(r.Form) == 0) {
		iter := c.Find(nil).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	} else {
		for key, value := range r.Form {
			splitKey := strings.Split(key, ":")
			if (len(splitKey) > 1) && (splitKey[0] == "subject") {
				subjectType := splitKey[1]
				referenceString := "http://"+host+":3001/"+subjectType+"/"+value[0]
				err := c.Find(bson.M{"subject.reference": referenceString}).All(&result)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}

	var conditionEntryList []models.ConditionBundleEntry
	for _, condition := range result {
		var entry models.ConditionBundleEntry
		entry.Title = "Condition " + condition.Id
		entry.Id = condition.Id
		entry.Content = condition
		conditionEntryList = append(conditionEntryList, entry)
	}

	var bundle models.ConditionBundle
	bundle.Type = "Bundle"
	bundle.Title = "Condition Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = conditionEntryList

	log.Println("Setting condition search context")
	context.Set(r, "Condition", result)
	context.Set(r, "Resource", "Condition")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadCondition(r *http.Request) (*models.Condition, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("conditions")
	result := models.Condition{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting condition read context")
	context.Set(r, "Condition", result)
	context.Set(r, "Resource", "Condition")
	return &result, nil
}

func ConditionShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadCondition(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Condition"))
}

func ConditionCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	condition := &models.Condition{}
	err := decoder.Decode(condition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("conditions")
	i := bson.NewObjectId()
	condition.Id = i.Hex()
	err = c.Insert(condition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting condition create context")
	context.Set(r, "Condition", condition)
	context.Set(r, "Resource", "Condition")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Condition/"+i.Hex())
}

func ConditionUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	condition := &models.Condition{}
	err := decoder.Decode(condition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("conditions")
	condition.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, condition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting condition update context")
	context.Set(r, "Condition", condition)
	context.Set(r, "Resource", "Condition")
	context.Set(r, "Action", "update")
}

func ConditionDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("conditions")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting condition delete context")
	context.Set(r, "Condition", id.Hex())
	context.Set(r, "Resource", "Condition")
	context.Set(r, "Action", "delete")
}
