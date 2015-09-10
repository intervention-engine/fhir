package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/fhir/search"
	"gopkg.in/mgo.v2/bson"
)

func RiskAssessmentIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case search.UnsupportedError:
				http.Error(rw, x.Error(), http.StatusNotImplemented)
			case search.InvalidSearchError:
				http.Error(rw, x.Error(), http.StatusBadRequest)
			default:
				http.Error(rw, fmt.Sprintf("%s", x), http.StatusInternalServerError)
			}
		}
	}()

	var result []models.RiskAssessment
	c := Database.C("riskassessments")

	r.ParseForm()
	if len(r.Form) == 0 {
		iter := c.Find(nil).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	} else {
		searcher := search.NewMongoSearcher(Database)
		query := search.Query{Resource: "RiskAssessment", Query: r.URL.RawQuery}
		err := searcher.CreateQuery(query).All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}

	var riskassessmentEntryList []models.BundleEntryComponent
	for i := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &result[i]
		riskassessmentEntryList = append(riskassessmentEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = riskassessmentEntryList

	log.Println("Setting riskassessment search context")
	context.Set(r, "RiskAssessment", result)
	context.Set(r, "Resource", "RiskAssessment")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadRiskAssessment(r *http.Request) (*models.RiskAssessment, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("riskassessments")
	result := models.RiskAssessment{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting riskassessment read context")
	context.Set(r, "RiskAssessment", result)
	context.Set(r, "Resource", "RiskAssessment")
	return &result, nil
}

func RiskAssessmentShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadRiskAssessment(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "RiskAssessment"))
}

func RiskAssessmentCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	riskassessment := &models.RiskAssessment{}
	err := decoder.Decode(riskassessment)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("riskassessments")
	i := bson.NewObjectId()
	riskassessment.Id = i.Hex()
	err = c.Insert(riskassessment)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting riskassessment create context")
	context.Set(r, "RiskAssessment", riskassessment)
	context.Set(r, "Resource", "RiskAssessment")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/RiskAssessment/"+i.Hex())
	rw.WriteHeader(http.StatusCreated)
}

func RiskAssessmentUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	riskassessment := &models.RiskAssessment{}
	err := decoder.Decode(riskassessment)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("riskassessments")
	riskassessment.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, riskassessment)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting riskassessment update context")
	context.Set(r, "RiskAssessment", riskassessment)
	context.Set(r, "Resource", "RiskAssessment")
	context.Set(r, "Action", "update")
}

func RiskAssessmentDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("riskassessments")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting riskassessment delete context")
	context.Set(r, "RiskAssessment", id.Hex())
	context.Set(r, "Resource", "RiskAssessment")
	context.Set(r, "Action", "delete")
}
