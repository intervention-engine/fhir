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

func MedicationIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Medication
	c := Database.C("medications")
	r.ParseForm()
	if len(r.Form) == 0 {
		iter := c.Find(nil).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	} else {
		codeValue := r.Form.Get("code")
		splitCode := strings.Split(codeValue, "|")
		if len(splitCode) > 1 {
			codeSystem := splitCode[0]
			code := splitCode[1]
			err := c.Find(bson.M{"code.coding.system": codeSystem, "code.coding.code": code}).All(&result)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		}
	}

	var medicationEntryList []models.MedicationBundleEntry
	for _, medication := range result {
		var entry models.MedicationBundleEntry
		entry.Id = medication.Id
		entry.Resource = medication
		medicationEntryList = append(medicationEntryList, entry)
	}

	var bundle models.MedicationBundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Total = len(result)
	bundle.Entry = medicationEntryList

	log.Println("Setting medication search context")
	context.Set(r, "Medication", result)
	context.Set(r, "Resource", "Medication")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadMedication(r *http.Request) (*models.Medication, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("medications")
	result := models.Medication{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting medication read context")
	context.Set(r, "Medication", result)
	context.Set(r, "Resource", "Medication")
	return &result, nil
}

func MedicationShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadMedication(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Medication"))
}

func MedicationCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	medication := &models.Medication{}
	err := decoder.Decode(medication)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medications")
	i := bson.NewObjectId()
	medication.Id = i.Hex()
	err = c.Insert(medication)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting medication create context")
	context.Set(r, "Medication", medication)
	context.Set(r, "Resource", "Medication")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Medication/"+i.Hex())
}

func MedicationUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	medication := &models.Medication{}
	err := decoder.Decode(medication)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medications")
	medication.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, medication)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting medication update context")
	context.Set(r, "Medication", medication)
	context.Set(r, "Resource", "Medication")
	context.Set(r, "Action", "update")
}

func MedicationDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("medications")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting medication delete context")
	context.Set(r, "Medication", id.Hex())
	context.Set(r, "Resource", "Medication")
	context.Set(r, "Action", "delete")
}
