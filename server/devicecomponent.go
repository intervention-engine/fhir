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

func DeviceComponentIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.DeviceComponent
	c := Database.C("devicecomponents")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var devicecomponentEntryList []models.BundleEntryComponent
	for _, devicecomponent := range result {
		var entry models.BundleEntryComponent
		entry.Resource = devicecomponent
		devicecomponentEntryList = append(devicecomponentEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = devicecomponentEntryList

	log.Println("Setting devicecomponent search context")
	context.Set(r, "DeviceComponent", result)
	context.Set(r, "Resource", "DeviceComponent")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadDeviceComponent(r *http.Request) (*models.DeviceComponent, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("devicecomponents")
	result := models.DeviceComponent{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting devicecomponent read context")
	context.Set(r, "DeviceComponent", result)
	context.Set(r, "Resource", "DeviceComponent")
	return &result, nil
}

func DeviceComponentShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadDeviceComponent(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "DeviceComponent"))
}

func DeviceComponentCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	devicecomponent := &models.DeviceComponent{}
	err := decoder.Decode(devicecomponent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("devicecomponents")
	i := bson.NewObjectId()
	devicecomponent.Id = i.Hex()
	err = c.Insert(devicecomponent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting devicecomponent create context")
	context.Set(r, "DeviceComponent", devicecomponent)
	context.Set(r, "Resource", "DeviceComponent")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/DeviceComponent/"+i.Hex())
}

func DeviceComponentUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	devicecomponent := &models.DeviceComponent{}
	err := decoder.Decode(devicecomponent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("devicecomponents")
	devicecomponent.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, devicecomponent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting devicecomponent update context")
	context.Set(r, "DeviceComponent", devicecomponent)
	context.Set(r, "Resource", "DeviceComponent")
	context.Set(r, "Action", "update")
}

func DeviceComponentDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("devicecomponents")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting devicecomponent delete context")
	context.Set(r, "DeviceComponent", id.Hex())
	context.Set(r, "Resource", "DeviceComponent")
	context.Set(r, "Action", "delete")
}
