package server

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"gitlab.mitre.org/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func PatientIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.Patient
	c := Database.C("patients")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var bundle models.PatientBundle
	bundle.Type = "Bundle"
	bundle.Title = "Patient Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entries = result

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func PatientShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("patients")

	result := models.Patient{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(result)
}

func PatientCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	patient := &models.Patient{}
	err := decoder.Decode(patient)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("patients")
	i := bson.NewObjectId()
	patient.Id = i.Hex()
	err = c.Insert(patient)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	f := Database.C("facts")
	fact := patient.ToFact()
	err = f.Insert(fact)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":8080/Patient/"+i.Hex())
}

func PatientUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	patient := &models.Patient{}
	err := decoder.Decode(patient)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("patients")
	patient.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, patient)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	f := Database.C("facts")
	fact := patient.ToFact()
	tempFact := models.Fact{}
	err = f.Find(bson.M{"targetid": id.Hex()}).One(&tempFact)

	fact.Id = tempFact.Id

	err = f.Update(bson.M{"targetid": id.Hex()}, fact)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func PatientDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("patients")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	f := Database.C("facts")
	err = f.Remove(bson.M{"targetid": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
