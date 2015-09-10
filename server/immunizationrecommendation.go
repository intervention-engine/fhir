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

func ImmunizationRecommendationIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if r := recover(); r != nil {
			rw.Header().Set("Content-Type", "application/json; charset=utf-8")
			switch x := r.(type) {
			case search.SearchError:
				rw.WriteHeader(x.HTTPStatus())
				json.NewEncoder(rw).Encode(x.OperationOutcome())
				return
			default:
				e := search.InternalServerError(fmt.Sprintf("%s", x))
				rw.WriteHeader(e.HTTPStatus())
				json.NewEncoder(rw).Encode(e.OperationOutcome())
			}
		}
	}()

	var result []models.ImmunizationRecommendation
	c := Database.C("immunizationrecommendations")

	r.ParseForm()
	if len(r.Form) == 0 {
		iter := c.Find(nil).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	} else {
		searcher := search.NewMongoSearcher(Database)
		query := search.Query{Resource: "ImmunizationRecommendation", Query: r.URL.RawQuery}
		err := searcher.CreateQuery(query).All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}

	var immunizationrecommendationEntryList []models.BundleEntryComponent
	for i := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &result[i]
		immunizationrecommendationEntryList = append(immunizationrecommendationEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = immunizationrecommendationEntryList

	log.Println("Setting immunizationrecommendation search context")
	context.Set(r, "ImmunizationRecommendation", result)
	context.Set(r, "Resource", "ImmunizationRecommendation")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadImmunizationRecommendation(r *http.Request) (*models.ImmunizationRecommendation, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("immunizationrecommendations")
	result := models.ImmunizationRecommendation{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting immunizationrecommendation read context")
	context.Set(r, "ImmunizationRecommendation", result)
	context.Set(r, "Resource", "ImmunizationRecommendation")
	return &result, nil
}

func ImmunizationRecommendationShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadImmunizationRecommendation(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "ImmunizationRecommendation"))
}

func ImmunizationRecommendationCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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

	rw.Header().Add("Location", "http://"+host+":3001/ImmunizationRecommendation/"+i.Hex())
	rw.WriteHeader(http.StatusCreated)
}

func ImmunizationRecommendationUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

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

func ImmunizationRecommendationDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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
