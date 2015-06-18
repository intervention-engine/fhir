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

func FamilyMemberHistoryIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.FamilyMemberHistory
	c := Database.C("familymemberhistorys")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var familymemberhistoryEntryList []models.FamilyMemberHistoryBundleEntry
	for _, familymemberhistory := range result {
		var entry models.FamilyMemberHistoryBundleEntry
		entry.Title = "FamilyMemberHistory " + familymemberhistory.Id
		entry.Id = familymemberhistory.Id
		entry.Content = familymemberhistory
		familymemberhistoryEntryList = append(familymemberhistoryEntryList, entry)
	}

	var bundle models.FamilyMemberHistoryBundle
	bundle.Type = "Bundle"
	bundle.Title = "FamilyMemberHistory Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = familymemberhistoryEntryList

	log.Println("Setting familymemberhistory search context")
	context.Set(r, "FamilyMemberHistory", result)
	context.Set(r, "Resource", "FamilyMemberHistory")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadFamilyMemberHistory(r *http.Request) (*models.FamilyMemberHistory, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("familymemberhistorys")
	result := models.FamilyMemberHistory{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting familymemberhistory read context")
	context.Set(r, "FamilyMemberHistory", result)
	context.Set(r, "Resource", "FamilyMemberHistory")
	return &result, nil
}

func FamilyMemberHistoryShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadFamilyMemberHistory(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "FamilyMemberHistory"))
}

func FamilyMemberHistoryCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	familymemberhistory := &models.FamilyMemberHistory{}
	err := decoder.Decode(familymemberhistory)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("familymemberhistorys")
	i := bson.NewObjectId()
	familymemberhistory.Id = i.Hex()
	err = c.Insert(familymemberhistory)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting familymemberhistory create context")
	context.Set(r, "FamilyMemberHistory", familymemberhistory)
	context.Set(r, "Resource", "FamilyMemberHistory")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/FamilyMemberHistory/"+i.Hex())
}

func FamilyMemberHistoryUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	familymemberhistory := &models.FamilyMemberHistory{}
	err := decoder.Decode(familymemberhistory)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("familymemberhistorys")
	familymemberhistory.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, familymemberhistory)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting familymemberhistory update context")
	context.Set(r, "FamilyMemberHistory", familymemberhistory)
	context.Set(r, "Resource", "FamilyMemberHistory")
	context.Set(r, "Action", "update")
}

func FamilyMemberHistoryDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("familymemberhistorys")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting familymemberhistory delete context")
	context.Set(r, "FamilyMemberHistory", id.Hex())
	context.Set(r, "Resource", "FamilyMemberHistory")
	context.Set(r, "Action", "delete")
}
