package server

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"gitlab.mitre.org/fhir/models"
	"github.com/gorilla/mux"
	"os"
)

func ValueSetIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.ValueSet
	c := Database.C("valuesets")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(result)
}

func ValueSetShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	}	else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("valuesets")

	result := models.ValueSet{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(result)
}

func ValueSetCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	valueset := &models.ValueSet{}
	err := decoder.Decode(valueset)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("valuesets")
	i := bson.NewObjectId()
	valueset.Id = i.Hex()
	err = c.Insert(valueset)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://" + host + "/valueset/" + i.Hex())
}

func ValueSetUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	}	else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	valueset := &models.ValueSet{}
	err := decoder.Decode(valueset)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("valuesets")
	valueset.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, valueset)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func ValueSetDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	}	else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("valuesets")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
