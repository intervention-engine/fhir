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

func CoverageIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Coverage
	c := Database.C("coverages")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var coverageEntryList []models.CoverageBundleEntry
	for _, coverage := range result {
		var entry models.CoverageBundleEntry
		entry.Id = coverage.Id
		entry.Resource = coverage
		coverageEntryList = append(coverageEntryList, entry)
	}

	var bundle models.CoverageBundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Total = len(result)
	bundle.Entry = coverageEntryList

	log.Println("Setting coverage search context")
	context.Set(r, "Coverage", result)
	context.Set(r, "Resource", "Coverage")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadCoverage(r *http.Request) (*models.Coverage, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("coverages")
	result := models.Coverage{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting coverage read context")
	context.Set(r, "Coverage", result)
	context.Set(r, "Resource", "Coverage")
	return &result, nil
}

func CoverageShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadCoverage(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Coverage"))
}

func CoverageCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	coverage := &models.Coverage{}
	err := decoder.Decode(coverage)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("coverages")
	i := bson.NewObjectId()
	coverage.Id = i.Hex()
	err = c.Insert(coverage)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting coverage create context")
	context.Set(r, "Coverage", coverage)
	context.Set(r, "Resource", "Coverage")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Coverage/"+i.Hex())
}

func CoverageUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	coverage := &models.Coverage{}
	err := decoder.Decode(coverage)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("coverages")
	coverage.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, coverage)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting coverage update context")
	context.Set(r, "Coverage", coverage)
	context.Set(r, "Resource", "Coverage")
	context.Set(r, "Action", "update")
}

func CoverageDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("coverages")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting coverage delete context")
	context.Set(r, "Coverage", id.Hex())
	context.Set(r, "Resource", "Coverage")
	context.Set(r, "Action", "delete")
}
