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

func AuditEventIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.AuditEvent
	c := Database.C("auditevents")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var auditeventEntryList []models.BundleEntryComponent
	for _, auditevent := range result {
		var entry models.BundleEntryComponent
		entry.Resource = auditevent
		auditeventEntryList = append(auditeventEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = auditeventEntryList

	log.Println("Setting auditevent search context")
	context.Set(r, "AuditEvent", result)
	context.Set(r, "Resource", "AuditEvent")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadAuditEvent(r *http.Request) (*models.AuditEvent, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("auditevents")
	result := models.AuditEvent{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting auditevent read context")
	context.Set(r, "AuditEvent", result)
	context.Set(r, "Resource", "AuditEvent")
	return &result, nil
}

func AuditEventShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadAuditEvent(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "AuditEvent"))
}

func AuditEventCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	auditevent := &models.AuditEvent{}
	err := decoder.Decode(auditevent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("auditevents")
	i := bson.NewObjectId()
	auditevent.Id = i.Hex()
	err = c.Insert(auditevent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting auditevent create context")
	context.Set(r, "AuditEvent", auditevent)
	context.Set(r, "Resource", "AuditEvent")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/AuditEvent/"+i.Hex())
}

func AuditEventUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	auditevent := &models.AuditEvent{}
	err := decoder.Decode(auditevent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("auditevents")
	auditevent.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, auditevent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting auditevent update context")
	context.Set(r, "AuditEvent", auditevent)
	context.Set(r, "Resource", "AuditEvent")
	context.Set(r, "Action", "update")
}

func AuditEventDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("auditevents")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting auditevent delete context")
	context.Set(r, "AuditEvent", id.Hex())
	context.Set(r, "Resource", "AuditEvent")
	context.Set(r, "Action", "delete")
}
