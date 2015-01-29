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

func AppointmentResponseIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.AppointmentResponse
	c := Database.C("appointmentresponses")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var appointmentresponseEntryList []models.AppointmentResponseBundleEntry
	for _, appointmentresponse := range result {
		var entry models.AppointmentResponseBundleEntry
		entry.Title = "AppointmentResponse " + appointmentresponse.Id
		entry.Id = appointmentresponse.Id
		entry.Content = appointmentresponse
		appointmentresponseEntryList = append(appointmentresponseEntryList, entry)
	}

	var bundle models.AppointmentResponseBundle
	bundle.Type = "Bundle"
	bundle.Title = "AppointmentResponse Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = appointmentresponseEntryList

	log.Println("Setting appointmentresponse search context")
	context.Set(r, "AppointmentResponse", result)
	context.Set(r, "Resource", "AppointmentResponse")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadAppointmentResponse(r *http.Request) (*models.AppointmentResponse, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("appointmentresponses")
	result := models.AppointmentResponse{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting appointmentresponse read context")
	context.Set(r, "AppointmentResponse", result)
	context.Set(r, "Resource", "AppointmentResponse")
	return &result, nil
}

func AppointmentResponseShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadAppointmentResponse(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "AppointmentResponse"))
}

func AppointmentResponseCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	appointmentresponse := &models.AppointmentResponse{}
	err := decoder.Decode(appointmentresponse)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("appointmentresponses")
	i := bson.NewObjectId()
	appointmentresponse.Id = i.Hex()
	err = c.Insert(appointmentresponse)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting appointmentresponse create context")
	context.Set(r, "AppointmentResponse", appointmentresponse)
	context.Set(r, "Resource", "AppointmentResponse")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/AppointmentResponse/"+i.Hex())
}

func AppointmentResponseUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	appointmentresponse := &models.AppointmentResponse{}
	err := decoder.Decode(appointmentresponse)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("appointmentresponses")
	appointmentresponse.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, appointmentresponse)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting appointmentresponse update context")
	context.Set(r, "AppointmentResponse", appointmentresponse)
	context.Set(r, "Resource", "AppointmentResponse")
	context.Set(r, "Action", "update")
}

func AppointmentResponseDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("appointmentresponses")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting appointmentresponse delete context")
	context.Set(r, "AppointmentResponse", id.Hex())
	context.Set(r, "Resource", "AppointmentResponse")
	context.Set(r, "Action", "delete")
}
