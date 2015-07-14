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

func MedicationPrescriptionIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.MedicationPrescription
	c := Database.C("medicationprescriptions")

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

	var medicationprescriptionEntryList []models.BundleEntryComponent
	for _, medicationprescription := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &medicationprescription
		medicationprescriptionEntryList = append(medicationprescriptionEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = medicationprescriptionEntryList

	log.Println("Setting medicationprescription search context")
	context.Set(r, "MedicationPrescription", result)
	context.Set(r, "Resource", "MedicationPrescription")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadMedicationPrescription(r *http.Request) (*models.MedicationPrescription, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("medicationprescriptions")
	result := models.MedicationPrescription{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting medicationprescription read context")
	context.Set(r, "MedicationPrescription", result)
	context.Set(r, "Resource", "MedicationPrescription")
	return &result, nil
}

func MedicationPrescriptionShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadMedicationPrescription(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "MedicationPrescription"))
}

func MedicationPrescriptionCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	medicationprescription := &models.MedicationPrescription{}
	err := decoder.Decode(medicationprescription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medicationprescriptions")
	i := bson.NewObjectId()
	medicationprescription.Id = i.Hex()
	err = c.Insert(medicationprescription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting medicationprescription create context")
	context.Set(r, "MedicationPrescription", medicationprescription)
	context.Set(r, "Resource", "MedicationPrescription")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/MedicationPrescription/"+i.Hex())
}

func MedicationPrescriptionUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	medicationprescription := &models.MedicationPrescription{}
	err := decoder.Decode(medicationprescription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medicationprescriptions")
	medicationprescription.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, medicationprescription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting medicationprescription update context")
	context.Set(r, "MedicationPrescription", medicationprescription)
	context.Set(r, "Resource", "MedicationPrescription")
	context.Set(r, "Action", "update")
}

func MedicationPrescriptionDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("medicationprescriptions")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting medicationprescription delete context")
	context.Set(r, "MedicationPrescription", id.Hex())
	context.Set(r, "Resource", "MedicationPrescription")
	context.Set(r, "Action", "delete")
}
