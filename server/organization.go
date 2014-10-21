package server

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"gitlab.mitre.org/fhir/models"
	"github.com/gorilla/mux"
	"os"
)

func OrganizationIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.Organization
	c := Database.C("organizations")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(result)
}

func OrganizationShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	}	else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("organizations")

	result := models.Organization{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(result)
}

func OrganizationCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	organization := &models.Organization{}
	err := decoder.Decode(organization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("organizations")
	i := bson.NewObjectId()
	organization.Id = i.Hex()
	err = c.Insert(organization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://" + host + "/organization/" + i.Hex())
}

func OrganizationUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	}	else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	organization := &models.Organization{}
	err := decoder.Decode(organization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("organizations")
	organization.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, organization)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func OrganizationDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	}	else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("organizations")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
