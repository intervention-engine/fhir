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

func ImmunizationIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.Immunization
	c := Database.C("immunizations")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var bundle models.ImmunizationBundle
	bundle.Type = "Bundle"
	bundle.Title = "Immunization Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entries = result

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func ImmunizationShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("immunizations")

	result := models.Immunization{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(result)
}

func ImmunizationCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	immunization := &models.Immunization{}
	err := decoder.Decode(immunization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("immunizations")
	i := bson.NewObjectId()
	immunization.Id = i.Hex()
	err = c.Insert(immunization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+"/immunization/"+i.Hex())
}

func ImmunizationUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	immunization := &models.Immunization{}
	err := decoder.Decode(immunization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("immunizations")
	immunization.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, immunization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func ImmunizationDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("immunizations")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
