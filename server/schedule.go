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

func ScheduleIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Schedule
	c := Database.C("schedules")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var scheduleEntryList []models.ScheduleBundleEntry
	for _, schedule := range result {
		var entry models.ScheduleBundleEntry
		entry.Title = "Schedule " + schedule.Id
		entry.Id = schedule.Id
		entry.Content = schedule
		scheduleEntryList = append(scheduleEntryList, entry)
	}

	var bundle models.ScheduleBundle
	bundle.Type = "Bundle"
	bundle.Title = "Schedule Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = scheduleEntryList

	log.Println("Setting schedule search context")
	context.Set(r, "Schedule", result)
	context.Set(r, "Resource", "Schedule")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadSchedule(r *http.Request) (*models.Schedule, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("schedules")
	result := models.Schedule{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting schedule read context")
	context.Set(r, "Schedule", result)
	context.Set(r, "Resource", "Schedule")
	return &result, nil
}

func ScheduleShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadSchedule(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Schedule"))
}

func ScheduleCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	schedule := &models.Schedule{}
	err := decoder.Decode(schedule)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("schedules")
	i := bson.NewObjectId()
	schedule.Id = i.Hex()
	err = c.Insert(schedule)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting schedule create context")
	context.Set(r, "Schedule", schedule)
	context.Set(r, "Resource", "Schedule")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Schedule/"+i.Hex())
}

func ScheduleUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	schedule := &models.Schedule{}
	err := decoder.Decode(schedule)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("schedules")
	schedule.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, schedule)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting schedule update context")
	context.Set(r, "Schedule", schedule)
	context.Set(r, "Resource", "Schedule")
	context.Set(r, "Action", "update")
}

func ScheduleDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("schedules")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting schedule delete context")
	context.Set(r, "Schedule", id.Hex())
	context.Set(r, "Resource", "Schedule")
	context.Set(r, "Action", "delete")
}
