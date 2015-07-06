package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func TestScriptIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.TestScript
	c := Database.C("testscripts")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var testscriptEntryList []models.TestScriptBundleEntry
	for _, testscript := range result {
		var entry models.TestScriptBundleEntry
		entry.Title = "TestScript " + testscript.Id
		entry.Id = testscript.Id
		entry.Content = testscript
		testscriptEntryList = append(testscriptEntryList, entry)
	}

	var bundle models.TestScriptBundle
	bundle.Type = "Bundle"
	bundle.Title = "TestScript Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = testscriptEntryList

	log.Println("Setting testscript search context")
	context.Set(r, "TestScript", result)
	context.Set(r, "Resource", "TestScript")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadTestScript(r *http.Request) (*models.TestScript, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("testscripts")
	result := models.TestScript{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting testscript read context")
	context.Set(r, "TestScript", result)
	context.Set(r, "Resource", "TestScript")
	return &result, nil
}

func TestScriptShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadTestScript(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "TestScript"))
}

func TestScriptCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	testscript := &models.TestScript{}
	err := decoder.Decode(testscript)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("testscripts")
	i := bson.NewObjectId()
	testscript.Id = i.Hex()
	err = c.Insert(testscript)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting testscript create context")
	context.Set(r, "TestScript", testscript)
	context.Set(r, "Resource", "TestScript")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/TestScript/"+i.Hex())
}

func TestScriptUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	testscript := &models.TestScript{}
	err := decoder.Decode(testscript)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("testscripts")
	testscript.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, testscript)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting testscript update context")
	context.Set(r, "TestScript", testscript)
	context.Set(r, "Resource", "TestScript")
	context.Set(r, "Action", "update")
}

func TestScriptDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("testscripts")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting testscript delete context")
	context.Set(r, "TestScript", id.Hex())
	context.Set(r, "Resource", "TestScript")
	context.Set(r, "Action", "delete")
}
