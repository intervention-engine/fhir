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
	"github.com/intervention-engine/fhir/search"
	"gopkg.in/mgo.v2/bson"
)

func ExplanationOfBenefitIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.ExplanationOfBenefit
	c := Database.C("explanationofbenefits")

	r.ParseForm()
	if len(r.Form) == 0 {
		iter := c.Find(nil).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	} else {
		searcher := search.NewMongoSearcher(Database)
		query := search.Query{Resource: "ExplanationOfBenefit", Query: r.URL.RawQuery}
		err := searcher.CreateQuery(query).All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}

	var explanationofbenefitEntryList []models.BundleEntryComponent
	for i := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &result[i]
		explanationofbenefitEntryList = append(explanationofbenefitEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = explanationofbenefitEntryList

	log.Println("Setting explanationofbenefit search context")
	context.Set(r, "ExplanationOfBenefit", result)
	context.Set(r, "Resource", "ExplanationOfBenefit")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadExplanationOfBenefit(r *http.Request) (*models.ExplanationOfBenefit, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("explanationofbenefits")
	result := models.ExplanationOfBenefit{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting explanationofbenefit read context")
	context.Set(r, "ExplanationOfBenefit", result)
	context.Set(r, "Resource", "ExplanationOfBenefit")
	return &result, nil
}

func ExplanationOfBenefitShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadExplanationOfBenefit(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "ExplanationOfBenefit"))
}

func ExplanationOfBenefitCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	explanationofbenefit := &models.ExplanationOfBenefit{}
	err := decoder.Decode(explanationofbenefit)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("explanationofbenefits")
	i := bson.NewObjectId()
	explanationofbenefit.Id = i.Hex()
	err = c.Insert(explanationofbenefit)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting explanationofbenefit create context")
	context.Set(r, "ExplanationOfBenefit", explanationofbenefit)
	context.Set(r, "Resource", "ExplanationOfBenefit")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/ExplanationOfBenefit/"+i.Hex())
}

func ExplanationOfBenefitUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	explanationofbenefit := &models.ExplanationOfBenefit{}
	err := decoder.Decode(explanationofbenefit)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("explanationofbenefits")
	explanationofbenefit.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, explanationofbenefit)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting explanationofbenefit update context")
	context.Set(r, "ExplanationOfBenefit", explanationofbenefit)
	context.Set(r, "Resource", "ExplanationOfBenefit")
	context.Set(r, "Action", "update")
}

func ExplanationOfBenefitDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("explanationofbenefits")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting explanationofbenefit delete context")
	context.Set(r, "ExplanationOfBenefit", id.Hex())
	context.Set(r, "Resource", "ExplanationOfBenefit")
	context.Set(r, "Action", "delete")
}
