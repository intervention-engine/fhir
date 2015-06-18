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

func VisionPrescriptionIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.VisionPrescription
	c := Database.C("visionprescriptions")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var visionprescriptionEntryList []models.VisionPrescriptionBundleEntry
	for _, visionprescription := range result {
		var entry models.VisionPrescriptionBundleEntry
		entry.Title = "VisionPrescription " + visionprescription.Id
		entry.Id = visionprescription.Id
		entry.Content = visionprescription
		visionprescriptionEntryList = append(visionprescriptionEntryList, entry)
	}

	var bundle models.VisionPrescriptionBundle
	bundle.Type = "Bundle"
	bundle.Title = "VisionPrescription Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = visionprescriptionEntryList

	log.Println("Setting visionprescription search context")
	context.Set(r, "VisionPrescription", result)
	context.Set(r, "Resource", "VisionPrescription")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadVisionPrescription(r *http.Request) (*models.VisionPrescription, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("visionprescriptions")
	result := models.VisionPrescription{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting visionprescription read context")
	context.Set(r, "VisionPrescription", result)
	context.Set(r, "Resource", "VisionPrescription")
	return &result, nil
}

func VisionPrescriptionShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadVisionPrescription(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "VisionPrescription"))
}

func VisionPrescriptionCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	visionprescription := &models.VisionPrescription{}
	err := decoder.Decode(visionprescription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("visionprescriptions")
	i := bson.NewObjectId()
	visionprescription.Id = i.Hex()
	err = c.Insert(visionprescription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting visionprescription create context")
	context.Set(r, "VisionPrescription", visionprescription)
	context.Set(r, "Resource", "VisionPrescription")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/VisionPrescription/"+i.Hex())
}

func VisionPrescriptionUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	visionprescription := &models.VisionPrescription{}
	err := decoder.Decode(visionprescription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("visionprescriptions")
	visionprescription.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, visionprescription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting visionprescription update context")
	context.Set(r, "VisionPrescription", visionprescription)
	context.Set(r, "Resource", "VisionPrescription")
	context.Set(r, "Action", "update")
}

func VisionPrescriptionDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("visionprescriptions")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting visionprescription delete context")
	context.Set(r, "VisionPrescription", id.Hex())
	context.Set(r, "Resource", "VisionPrescription")
	context.Set(r, "Action", "delete")
}
