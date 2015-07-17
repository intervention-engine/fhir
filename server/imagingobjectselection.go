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

func ImagingObjectSelectionIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.ImagingObjectSelection
	c := Database.C("imagingobjectselections")

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
			if splitKey[0] == "patient" {
				err := c.Find(bson.M{"patient.referenceid": value[0]}).All(&result)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}

	var imagingobjectselectionEntryList []models.BundleEntryComponent
	for i := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &result[i]
		imagingobjectselectionEntryList = append(imagingobjectselectionEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = imagingobjectselectionEntryList

	log.Println("Setting imagingobjectselection search context")
	context.Set(r, "ImagingObjectSelection", result)
	context.Set(r, "Resource", "ImagingObjectSelection")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadImagingObjectSelection(r *http.Request) (*models.ImagingObjectSelection, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("imagingobjectselections")
	result := models.ImagingObjectSelection{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting imagingobjectselection read context")
	context.Set(r, "ImagingObjectSelection", result)
	context.Set(r, "Resource", "ImagingObjectSelection")
	return &result, nil
}

func ImagingObjectSelectionShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadImagingObjectSelection(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "ImagingObjectSelection"))
}

func ImagingObjectSelectionCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	imagingobjectselection := &models.ImagingObjectSelection{}
	err := decoder.Decode(imagingobjectselection)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("imagingobjectselections")
	i := bson.NewObjectId()
	imagingobjectselection.Id = i.Hex()
	err = c.Insert(imagingobjectselection)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting imagingobjectselection create context")
	context.Set(r, "ImagingObjectSelection", imagingobjectselection)
	context.Set(r, "Resource", "ImagingObjectSelection")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/ImagingObjectSelection/"+i.Hex())
}

func ImagingObjectSelectionUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	imagingobjectselection := &models.ImagingObjectSelection{}
	err := decoder.Decode(imagingobjectselection)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("imagingobjectselections")
	imagingobjectselection.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, imagingobjectselection)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting imagingobjectselection update context")
	context.Set(r, "ImagingObjectSelection", imagingobjectselection)
	context.Set(r, "Resource", "ImagingObjectSelection")
	context.Set(r, "Action", "update")
}

func ImagingObjectSelectionDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("imagingobjectselections")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting imagingobjectselection delete context")
	context.Set(r, "ImagingObjectSelection", id.Hex())
	context.Set(r, "Resource", "ImagingObjectSelection")
	context.Set(r, "Action", "delete")
}
