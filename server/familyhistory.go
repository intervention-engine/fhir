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

func FamilyHistoryIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.FamilyHistory
	c := Database.C("familyhistorys")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var bundle models.FamilyHistoryBundle
	bundle.Type = "Bundle"
	bundle.Title = "FamilyHistory Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entries = result

	log.Println("Setting familyhistory search context")
	context.Set(r, "FamilyHistory", result)
	context.Set(r, "Resource", "FamilyHistory")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadFamilyHistory(r *http.Request) (*models.FamilyHistory, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("familyhistorys")
	result := models.FamilyHistory{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting familyhistory read context")
	context.Set(r, "FamilyHistory", result)
	context.Set(r, "Resource", "FamilyHistory")
	return &result, nil
}

func FamilyHistoryShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadFamilyHistory(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "FamilyHistory"))
}

func FamilyHistoryCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	familyhistory := &models.FamilyHistory{}
	err := decoder.Decode(familyhistory)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("familyhistorys")
	i := bson.NewObjectId()
	familyhistory.Id = i.Hex()
	err = c.Insert(familyhistory)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting familyhistory create context")
	context.Set(r, "FamilyHistory", familyhistory)
	context.Set(r, "Resource", "FamilyHistory")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/FamilyHistory/"+i.Hex())
}

func FamilyHistoryUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	familyhistory := &models.FamilyHistory{}
	err := decoder.Decode(familyhistory)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("familyhistorys")
	familyhistory.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, familyhistory)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting familyhistory update context")
	context.Set(r, "FamilyHistory", familyhistory)
	context.Set(r, "Resource", "FamilyHistory")
	context.Set(r, "Action", "update")
}

func FamilyHistoryDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("familyhistorys")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting familyhistory delete context")
	context.Set(r, "FamilyHistory", id.Hex())
	context.Set(r, "Resource", "FamilyHistory")
	context.Set(r, "Action", "delete")
}
