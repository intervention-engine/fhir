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

func EnrollmentRequestIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.EnrollmentRequest
	c := Database.C("enrollmentrequests")

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

	var enrollmentrequestEntryList []models.BundleEntryComponent
	for _, enrollmentrequest := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &enrollmentrequest
		enrollmentrequestEntryList = append(enrollmentrequestEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = enrollmentrequestEntryList

	log.Println("Setting enrollmentrequest search context")
	context.Set(r, "EnrollmentRequest", result)
	context.Set(r, "Resource", "EnrollmentRequest")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadEnrollmentRequest(r *http.Request) (*models.EnrollmentRequest, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("enrollmentrequests")
	result := models.EnrollmentRequest{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting enrollmentrequest read context")
	context.Set(r, "EnrollmentRequest", result)
	context.Set(r, "Resource", "EnrollmentRequest")
	return &result, nil
}

func EnrollmentRequestShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadEnrollmentRequest(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "EnrollmentRequest"))
}

func EnrollmentRequestCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	enrollmentrequest := &models.EnrollmentRequest{}
	err := decoder.Decode(enrollmentrequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("enrollmentrequests")
	i := bson.NewObjectId()
	enrollmentrequest.Id = i.Hex()
	err = c.Insert(enrollmentrequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting enrollmentrequest create context")
	context.Set(r, "EnrollmentRequest", enrollmentrequest)
	context.Set(r, "Resource", "EnrollmentRequest")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/EnrollmentRequest/"+i.Hex())
}

func EnrollmentRequestUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	enrollmentrequest := &models.EnrollmentRequest{}
	err := decoder.Decode(enrollmentrequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("enrollmentrequests")
	enrollmentrequest.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, enrollmentrequest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting enrollmentrequest update context")
	context.Set(r, "EnrollmentRequest", enrollmentrequest)
	context.Set(r, "Resource", "EnrollmentRequest")
	context.Set(r, "Action", "update")
}

func EnrollmentRequestDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("enrollmentrequests")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting enrollmentrequest delete context")
	context.Set(r, "EnrollmentRequest", id.Hex())
	context.Set(r, "Resource", "EnrollmentRequest")
	context.Set(r, "Action", "delete")
}
