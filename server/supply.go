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

func SupplyIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Supply
	c := Database.C("supplys")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var supplyEntryList []models.SupplyBundleEntry
	for _, supply := range result {
		var entry models.SupplyBundleEntry
		entry.Title = "Supply " + supply.Id
		entry.Id = supply.Id
		entry.Content = supply
		supplyEntryList = append(supplyEntryList, entry)
	}

	var bundle models.SupplyBundle
	bundle.Type = "Bundle"
	bundle.Title = "Supply Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = supplyEntryList

	log.Println("Setting supply search context")
	context.Set(r, "Supply", result)
	context.Set(r, "Resource", "Supply")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadSupply(r *http.Request) (*models.Supply, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("supplys")
	result := models.Supply{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting supply read context")
	context.Set(r, "Supply", result)
	context.Set(r, "Resource", "Supply")
	return &result, nil
}

func SupplyShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadSupply(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Supply"))
}

func SupplyCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	supply := &models.Supply{}
	err := decoder.Decode(supply)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("supplys")
	i := bson.NewObjectId()
	supply.Id = i.Hex()
	err = c.Insert(supply)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting supply create context")
	context.Set(r, "Supply", supply)
	context.Set(r, "Resource", "Supply")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Supply/"+i.Hex())
}

func SupplyUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	supply := &models.Supply{}
	err := decoder.Decode(supply)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("supplys")
	supply.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, supply)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting supply update context")
	context.Set(r, "Supply", supply)
	context.Set(r, "Resource", "Supply")
	context.Set(r, "Action", "update")
}

func SupplyDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("supplys")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting supply delete context")
	context.Set(r, "Supply", id.Hex())
	context.Set(r, "Resource", "Supply")
	context.Set(r, "Action", "delete")
}
