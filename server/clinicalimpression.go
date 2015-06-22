package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func ClinicalImpressionIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.ClinicalImpression
	c := Database.C("clinicalimpressions")

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
			if (len(splitKey) > 1) && (splitKey[0] == "patient") {
				err := c.Find(bson.M{"patient.referenceid": value[0]}).All(&result)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}

	var clinicalimpressionEntryList []models.ClinicalImpressionBundleEntry
	for _, clinicalimpression := range result {
		var entry models.ClinicalImpressionBundleEntry
		entry.Title = "ClinicalImpression " + clinicalimpression.Id
		entry.Id = clinicalimpression.Id
		entry.Content = clinicalimpression
		clinicalimpressionEntryList = append(clinicalimpressionEntryList, entry)
	}

	var bundle models.ClinicalImpressionBundle
	bundle.Type = "Bundle"
	bundle.Title = "ClinicalImpression Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = clinicalimpressionEntryList

	log.Println("Setting clinicalimpression search context")
	context.Set(r, "ClinicalImpression", result)
	context.Set(r, "Resource", "ClinicalImpression")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadClinicalImpression(r *http.Request) (*models.ClinicalImpression, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("clinicalimpressions")
	result := models.ClinicalImpression{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting clinicalimpression read context")
	context.Set(r, "ClinicalImpression", result)
	context.Set(r, "Resource", "ClinicalImpression")
	return &result, nil
}

func ClinicalImpressionShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadClinicalImpression(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "ClinicalImpression"))
}

func ClinicalImpressionCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	clinicalimpression := &models.ClinicalImpression{}
	err := decoder.Decode(clinicalimpression)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("clinicalimpressions")
	i := bson.NewObjectId()
	clinicalimpression.Id = i.Hex()
	err = c.Insert(clinicalimpression)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting clinicalimpression create context")
	context.Set(r, "ClinicalImpression", clinicalimpression)
	context.Set(r, "Resource", "ClinicalImpression")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/ClinicalImpression/"+i.Hex())
}

func ClinicalImpressionUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	clinicalimpression := &models.ClinicalImpression{}
	err := decoder.Decode(clinicalimpression)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("clinicalimpressions")
	clinicalimpression.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, clinicalimpression)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting clinicalimpression update context")
	context.Set(r, "ClinicalImpression", clinicalimpression)
	context.Set(r, "Resource", "ClinicalImpression")
	context.Set(r, "Action", "update")
}

func ClinicalImpressionDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("clinicalimpressions")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting clinicalimpression delete context")
	context.Set(r, "ClinicalImpression", id.Hex())
	context.Set(r, "Resource", "ClinicalImpression")
	context.Set(r, "Action", "delete")
}
