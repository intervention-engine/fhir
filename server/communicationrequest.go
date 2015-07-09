package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func CommunicationRequestIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.CommunicationRequest
	c := Database.C("communicationrequests")

	r.ParseForm()
	if len(r.Form) == 0 {
		iter := c.Find(nil).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	} else {
		for key, value := range r.Form {
			splitKey := strings.Split(key, ":")
			if splitKey[0] == "subject" {
				err := c.Find(bson.M{"subject.referenceid": value[0]}).All(&result)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}

	var communicationrequestEntryList []models.CommunicationRequestBundleEntry
	for _, communicationrequest := range result {
		var entry models.CommunicationRequestBundleEntry
		entry.Id = communicationrequest.Id
		entry.Resource = communicationrequest
		communicationrequestEntryList = append(communicationrequestEntryList, entry)
	}

	var bundle models.CommunicationRequestBundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Total = len(result)
	bundle.Entry = communicationrequestEntryList

	log.Println("Setting communicationrequest search context")
	context.Set(r, "CommunicationRequest", result)
	context.Set(r, "Resource", "CommunicationRequest")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadCommunicationRequest(r *http.Request) (*models.CommunicationRequest, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("communicationrequests")
	result := models.CommunicationRequest{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting communicationrequest read context")
	context.Set(r, "CommunicationRequest", result)
	context.Set(r, "Resource", "CommunicationRequest")
	return &result, nil
}

func CommunicationRequestShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadCommunicationRequest(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "CommunicationRequest"))
}

func CommunicationRequestCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	communicationrequest := &models.CommunicationRequest{}
	err := decoder.Decode(communicationrequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("communicationrequests")
	i := bson.NewObjectId()
	communicationrequest.Id = i.Hex()
	err = c.Insert(communicationrequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting communicationrequest create context")
	context.Set(r, "CommunicationRequest", communicationrequest)
	context.Set(r, "Resource", "CommunicationRequest")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/CommunicationRequest/"+i.Hex())
}

func CommunicationRequestUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	communicationrequest := &models.CommunicationRequest{}
	err := decoder.Decode(communicationrequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("communicationrequests")
	communicationrequest.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, communicationrequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting communicationrequest update context")
	context.Set(r, "CommunicationRequest", communicationrequest)
	context.Set(r, "Resource", "CommunicationRequest")
	context.Set(r, "Action", "update")
}

func CommunicationRequestDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("communicationrequests")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting communicationrequest delete context")
	context.Set(r, "CommunicationRequest", id.Hex())
	context.Set(r, "Resource", "CommunicationRequest")
	context.Set(r, "Action", "delete")
}
