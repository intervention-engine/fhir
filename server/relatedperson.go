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

func RelatedPersonIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.RelatedPerson
	c := Database.C("relatedpersons")

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

	var relatedpersonEntryList []models.RelatedPersonBundleEntry
	for _, relatedperson := range result {
		var entry models.RelatedPersonBundleEntry
		entry.Id = relatedperson.Id
		entry.Resource = relatedperson
		relatedpersonEntryList = append(relatedpersonEntryList, entry)
	}

	var bundle models.RelatedPersonBundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Total = len(result)
	bundle.Entry = relatedpersonEntryList

	log.Println("Setting relatedperson search context")
	context.Set(r, "RelatedPerson", result)
	context.Set(r, "Resource", "RelatedPerson")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadRelatedPerson(r *http.Request) (*models.RelatedPerson, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("relatedpersons")
	result := models.RelatedPerson{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting relatedperson read context")
	context.Set(r, "RelatedPerson", result)
	context.Set(r, "Resource", "RelatedPerson")
	return &result, nil
}

func RelatedPersonShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadRelatedPerson(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "RelatedPerson"))
}

func RelatedPersonCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	relatedperson := &models.RelatedPerson{}
	err := decoder.Decode(relatedperson)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("relatedpersons")
	i := bson.NewObjectId()
	relatedperson.Id = i.Hex()
	err = c.Insert(relatedperson)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting relatedperson create context")
	context.Set(r, "RelatedPerson", relatedperson)
	context.Set(r, "Resource", "RelatedPerson")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/RelatedPerson/"+i.Hex())
}

func RelatedPersonUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	relatedperson := &models.RelatedPerson{}
	err := decoder.Decode(relatedperson)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("relatedpersons")
	relatedperson.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, relatedperson)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting relatedperson update context")
	context.Set(r, "RelatedPerson", relatedperson)
	context.Set(r, "Resource", "RelatedPerson")
	context.Set(r, "Action", "update")
}

func RelatedPersonDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("relatedpersons")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting relatedperson delete context")
	context.Set(r, "RelatedPerson", id.Hex())
	context.Set(r, "Resource", "RelatedPerson")
	context.Set(r, "Action", "delete")
}
