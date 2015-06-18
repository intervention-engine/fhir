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

func ContractIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Contract
	c := Database.C("contracts")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var contractEntryList []models.ContractBundleEntry
	for _, contract := range result {
		var entry models.ContractBundleEntry
		entry.Title = "Contract " + contract.Id
		entry.Id = contract.Id
		entry.Content = contract
		contractEntryList = append(contractEntryList, entry)
	}

	var bundle models.ContractBundle
	bundle.Type = "Bundle"
	bundle.Title = "Contract Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = contractEntryList

	log.Println("Setting contract search context")
	context.Set(r, "Contract", result)
	context.Set(r, "Resource", "Contract")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadContract(r *http.Request) (*models.Contract, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("contracts")
	result := models.Contract{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting contract read context")
	context.Set(r, "Contract", result)
	context.Set(r, "Resource", "Contract")
	return &result, nil
}

func ContractShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadContract(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Contract"))
}

func ContractCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	contract := &models.Contract{}
	err := decoder.Decode(contract)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("contracts")
	i := bson.NewObjectId()
	contract.Id = i.Hex()
	err = c.Insert(contract)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting contract create context")
	context.Set(r, "Contract", contract)
	context.Set(r, "Resource", "Contract")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Contract/"+i.Hex())
}

func ContractUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	contract := &models.Contract{}
	err := decoder.Decode(contract)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("contracts")
	contract.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, contract)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting contract update context")
	context.Set(r, "Contract", contract)
	context.Set(r, "Resource", "Contract")
	context.Set(r, "Action", "update")
}

func ContractDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("contracts")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting contract delete context")
	context.Set(r, "Contract", id.Hex())
	context.Set(r, "Resource", "Contract")
	context.Set(r, "Action", "delete")
}
