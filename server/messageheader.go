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

func MessageHeaderIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.MessageHeader
	c := Database.C("messageheaders")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var messageheaderEntryList []models.BundleEntryComponent
	for _, messageheader := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &messageheader
		messageheaderEntryList = append(messageheaderEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = messageheaderEntryList

	log.Println("Setting messageheader search context")
	context.Set(r, "MessageHeader", result)
	context.Set(r, "Resource", "MessageHeader")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadMessageHeader(r *http.Request) (*models.MessageHeader, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("messageheaders")
	result := models.MessageHeader{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting messageheader read context")
	context.Set(r, "MessageHeader", result)
	context.Set(r, "Resource", "MessageHeader")
	return &result, nil
}

func MessageHeaderShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadMessageHeader(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "MessageHeader"))
}

func MessageHeaderCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	messageheader := &models.MessageHeader{}
	err := decoder.Decode(messageheader)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("messageheaders")
	i := bson.NewObjectId()
	messageheader.Id = i.Hex()
	err = c.Insert(messageheader)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting messageheader create context")
	context.Set(r, "MessageHeader", messageheader)
	context.Set(r, "Resource", "MessageHeader")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/MessageHeader/"+i.Hex())
}

func MessageHeaderUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	messageheader := &models.MessageHeader{}
	err := decoder.Decode(messageheader)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("messageheaders")
	messageheader.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, messageheader)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting messageheader update context")
	context.Set(r, "MessageHeader", messageheader)
	context.Set(r, "Resource", "MessageHeader")
	context.Set(r, "Action", "update")
}

func MessageHeaderDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("messageheaders")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting messageheader delete context")
	context.Set(r, "MessageHeader", id.Hex())
	context.Set(r, "Resource", "MessageHeader")
	context.Set(r, "Action", "delete")
}
