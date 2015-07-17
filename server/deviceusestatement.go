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

func DeviceUseStatementIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.DeviceUseStatement
	c := Database.C("deviceusestatements")

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

	var deviceusestatementEntryList []models.BundleEntryComponent
	for i := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &result[i]
		deviceusestatementEntryList = append(deviceusestatementEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = deviceusestatementEntryList

	log.Println("Setting deviceusestatement search context")
	context.Set(r, "DeviceUseStatement", result)
	context.Set(r, "Resource", "DeviceUseStatement")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadDeviceUseStatement(r *http.Request) (*models.DeviceUseStatement, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("deviceusestatements")
	result := models.DeviceUseStatement{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting deviceusestatement read context")
	context.Set(r, "DeviceUseStatement", result)
	context.Set(r, "Resource", "DeviceUseStatement")
	return &result, nil
}

func DeviceUseStatementShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadDeviceUseStatement(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "DeviceUseStatement"))
}

func DeviceUseStatementCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	deviceusestatement := &models.DeviceUseStatement{}
	err := decoder.Decode(deviceusestatement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("deviceusestatements")
	i := bson.NewObjectId()
	deviceusestatement.Id = i.Hex()
	err = c.Insert(deviceusestatement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting deviceusestatement create context")
	context.Set(r, "DeviceUseStatement", deviceusestatement)
	context.Set(r, "Resource", "DeviceUseStatement")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/DeviceUseStatement/"+i.Hex())
}

func DeviceUseStatementUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	deviceusestatement := &models.DeviceUseStatement{}
	err := decoder.Decode(deviceusestatement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("deviceusestatements")
	deviceusestatement.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, deviceusestatement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting deviceusestatement update context")
	context.Set(r, "DeviceUseStatement", deviceusestatement)
	context.Set(r, "Resource", "DeviceUseStatement")
	context.Set(r, "Action", "update")
}

func DeviceUseStatementDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("deviceusestatements")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting deviceusestatement delete context")
	context.Set(r, "DeviceUseStatement", id.Hex())
	context.Set(r, "Resource", "DeviceUseStatement")
	context.Set(r, "Action", "delete")
}
