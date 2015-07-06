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

func NamingSystemIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.NamingSystem
	c := Database.C("namingsystems")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var namingsystemEntryList []models.NamingSystemBundleEntry
	for _, namingsystem := range result {
		var entry models.NamingSystemBundleEntry
		entry.Title = "NamingSystem " + namingsystem.Id
		entry.Id = namingsystem.Id
		entry.Content = namingsystem
		namingsystemEntryList = append(namingsystemEntryList, entry)
	}

	var bundle models.NamingSystemBundle
	bundle.Type = "Bundle"
	bundle.Title = "NamingSystem Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = namingsystemEntryList

	log.Println("Setting namingsystem search context")
	context.Set(r, "NamingSystem", result)
	context.Set(r, "Resource", "NamingSystem")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadNamingSystem(r *http.Request) (*models.NamingSystem, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("namingsystems")
	result := models.NamingSystem{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting namingsystem read context")
	context.Set(r, "NamingSystem", result)
	context.Set(r, "Resource", "NamingSystem")
	return &result, nil
}

func NamingSystemShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadNamingSystem(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "NamingSystem"))
}

func NamingSystemCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	namingsystem := &models.NamingSystem{}
	err := decoder.Decode(namingsystem)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("namingsystems")
	i := bson.NewObjectId()
	namingsystem.Id = i.Hex()
	err = c.Insert(namingsystem)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting namingsystem create context")
	context.Set(r, "NamingSystem", namingsystem)
	context.Set(r, "Resource", "NamingSystem")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/NamingSystem/"+i.Hex())
}

func NamingSystemUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	namingsystem := &models.NamingSystem{}
	err := decoder.Decode(namingsystem)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("namingsystems")
	namingsystem.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, namingsystem)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting namingsystem update context")
	context.Set(r, "NamingSystem", namingsystem)
	context.Set(r, "Resource", "NamingSystem")
	context.Set(r, "Action", "update")
}

func NamingSystemDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("namingsystems")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting namingsystem delete context")
	context.Set(r, "NamingSystem", id.Hex())
	context.Set(r, "Resource", "NamingSystem")
	context.Set(r, "Action", "delete")
}
