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

func NutritionOrderIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.NutritionOrder
	c := Database.C("nutritionorders")

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

	var nutritionorderEntryList []models.NutritionOrderBundleEntry
	for _, nutritionorder := range result {
		var entry models.NutritionOrderBundleEntry
		entry.Id = nutritionorder.Id
		entry.Resource = nutritionorder
		nutritionorderEntryList = append(nutritionorderEntryList, entry)
	}

	var bundle models.NutritionOrderBundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Total = len(result)
	bundle.Entry = nutritionorderEntryList

	log.Println("Setting nutritionorder search context")
	context.Set(r, "NutritionOrder", result)
	context.Set(r, "Resource", "NutritionOrder")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadNutritionOrder(r *http.Request) (*models.NutritionOrder, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("nutritionorders")
	result := models.NutritionOrder{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting nutritionorder read context")
	context.Set(r, "NutritionOrder", result)
	context.Set(r, "Resource", "NutritionOrder")
	return &result, nil
}

func NutritionOrderShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadNutritionOrder(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "NutritionOrder"))
}

func NutritionOrderCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	nutritionorder := &models.NutritionOrder{}
	err := decoder.Decode(nutritionorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("nutritionorders")
	i := bson.NewObjectId()
	nutritionorder.Id = i.Hex()
	err = c.Insert(nutritionorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting nutritionorder create context")
	context.Set(r, "NutritionOrder", nutritionorder)
	context.Set(r, "Resource", "NutritionOrder")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/NutritionOrder/"+i.Hex())
}

func NutritionOrderUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	nutritionorder := &models.NutritionOrder{}
	err := decoder.Decode(nutritionorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("nutritionorders")
	nutritionorder.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, nutritionorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting nutritionorder update context")
	context.Set(r, "NutritionOrder", nutritionorder)
	context.Set(r, "Resource", "NutritionOrder")
	context.Set(r, "Action", "update")
}

func NutritionOrderDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("nutritionorders")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting nutritionorder delete context")
	context.Set(r, "NutritionOrder", id.Hex())
	context.Set(r, "Resource", "NutritionOrder")
	context.Set(r, "Action", "delete")
}
