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

func OrganizationIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Organization
	c := Database.C("organizations")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var organizationEntryList []models.OrganizationBundleEntry
	for _, organization := range result {
		var entry models.OrganizationBundleEntry
		entry.Title = "Organization " + organization.Id
		entry.Id = organization.Id
		entry.Content = organization
		organizationEntryList = append(organizationEntryList, entry)
	}

	var bundle models.OrganizationBundle
	bundle.Type = "Bundle"
	bundle.Title = "Organization Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = organizationEntryList

	log.Println("Setting organization search context")
	context.Set(r, "Organization", result)
	context.Set(r, "Resource", "Organization")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadOrganization(r *http.Request) (*models.Organization, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("organizations")
	result := models.Organization{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting organization read context")
	context.Set(r, "Organization", result)
	context.Set(r, "Resource", "Organization")
	return &result, nil
}

func OrganizationShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadOrganization(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Organization"))
}

func OrganizationCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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

	log.Println("Setting organization create context")
	context.Set(r, "Organization", organization)
	context.Set(r, "Resource", "Organization")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Organization/"+i.Hex())
}

func OrganizationUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
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

	log.Println("Setting organization update context")
	context.Set(r, "Organization", organization)
	context.Set(r, "Resource", "Organization")
	context.Set(r, "Action", "update")
}

func OrganizationDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("organizations")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting organization delete context")
	context.Set(r, "Organization", id.Hex())
	context.Set(r, "Resource", "Organization")
	context.Set(r, "Action", "delete")
}
