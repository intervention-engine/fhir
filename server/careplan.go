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

func CarePlanIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.CarePlan
	c := Database.C("careplans")

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

	var careplanEntryList []models.CarePlanBundleEntry
	for _, careplan := range result {
		var entry models.CarePlanBundleEntry
		entry.Title = "CarePlan " + careplan.Id
		entry.Id = careplan.Id
		entry.Content = careplan
		careplanEntryList = append(careplanEntryList, entry)
	}

	var bundle models.CarePlanBundle
	bundle.Type = "Bundle"
	bundle.Title = "CarePlan Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = careplanEntryList

	log.Println("Setting careplan search context")
	context.Set(r, "CarePlan", result)
	context.Set(r, "Resource", "CarePlan")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadCarePlan(r *http.Request) (*models.CarePlan, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("careplans")
	result := models.CarePlan{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting careplan read context")
	context.Set(r, "CarePlan", result)
	context.Set(r, "Resource", "CarePlan")
	return &result, nil
}

func CarePlanShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadCarePlan(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "CarePlan"))
}

func CarePlanCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	careplan := &models.CarePlan{}
	err := decoder.Decode(careplan)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("careplans")
	i := bson.NewObjectId()
	careplan.Id = i.Hex()
	err = c.Insert(careplan)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting careplan create context")
	context.Set(r, "CarePlan", careplan)
	context.Set(r, "Resource", "CarePlan")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/CarePlan/"+i.Hex())
}

func CarePlanUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	careplan := &models.CarePlan{}
	err := decoder.Decode(careplan)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("careplans")
	careplan.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, careplan)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting careplan update context")
	context.Set(r, "CarePlan", careplan)
	context.Set(r, "Resource", "CarePlan")
	context.Set(r, "Action", "update")
}

func CarePlanDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("careplans")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting careplan delete context")
	context.Set(r, "CarePlan", id.Hex())
	context.Set(r, "Resource", "CarePlan")
	context.Set(r, "Action", "delete")
}
