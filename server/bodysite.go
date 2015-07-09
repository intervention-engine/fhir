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

func BodySiteIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.BodySite
	c := Database.C("bodysites")

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

	var bodysiteEntryList []models.BodySiteBundleEntry
	for _, bodysite := range result {
		var entry models.BodySiteBundleEntry
		entry.Id = bodysite.Id
		entry.Resource = bodysite
		bodysiteEntryList = append(bodysiteEntryList, entry)
	}

	var bundle models.BodySiteBundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Total = len(result)
	bundle.Entry = bodysiteEntryList

	log.Println("Setting bodysite search context")
	context.Set(r, "BodySite", result)
	context.Set(r, "Resource", "BodySite")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadBodySite(r *http.Request) (*models.BodySite, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("bodysites")
	result := models.BodySite{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting bodysite read context")
	context.Set(r, "BodySite", result)
	context.Set(r, "Resource", "BodySite")
	return &result, nil
}

func BodySiteShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadBodySite(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "BodySite"))
}

func BodySiteCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	bodysite := &models.BodySite{}
	err := decoder.Decode(bodysite)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("bodysites")
	i := bson.NewObjectId()
	bodysite.Id = i.Hex()
	err = c.Insert(bodysite)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting bodysite create context")
	context.Set(r, "BodySite", bodysite)
	context.Set(r, "Resource", "BodySite")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/BodySite/"+i.Hex())
}

func BodySiteUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	bodysite := &models.BodySite{}
	err := decoder.Decode(bodysite)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("bodysites")
	bodysite.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, bodysite)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting bodysite update context")
	context.Set(r, "BodySite", bodysite)
	context.Set(r, "Resource", "BodySite")
	context.Set(r, "Action", "update")
}

func BodySiteDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("bodysites")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting bodysite delete context")
	context.Set(r, "BodySite", id.Hex())
	context.Set(r, "Resource", "BodySite")
	context.Set(r, "Action", "delete")
}
