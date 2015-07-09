package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func SearchParameterIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.SearchParameter
	c := Database.C("searchparameters")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var searchparameterEntryList []models.SearchParameterBundleEntry
	for _, searchparameter := range result {
		var entry models.SearchParameterBundleEntry
		entry.Id = searchparameter.Id
		entry.Resource = searchparameter
		searchparameterEntryList = append(searchparameterEntryList, entry)
	}

	var bundle models.SearchParameterBundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Total = len(result)
	bundle.Entry = searchparameterEntryList

	log.Println("Setting searchparameter search context")
	context.Set(r, "SearchParameter", result)
	context.Set(r, "Resource", "SearchParameter")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadSearchParameter(r *http.Request) (*models.SearchParameter, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("searchparameters")
	result := models.SearchParameter{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting searchparameter read context")
	context.Set(r, "SearchParameter", result)
	context.Set(r, "Resource", "SearchParameter")
	return &result, nil
}

func SearchParameterShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadSearchParameter(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "SearchParameter"))
}

func SearchParameterCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	searchparameter := &models.SearchParameter{}
	err := decoder.Decode(searchparameter)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("searchparameters")
	i := bson.NewObjectId()
	searchparameter.Id = i.Hex()
	err = c.Insert(searchparameter)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting searchparameter create context")
	context.Set(r, "SearchParameter", searchparameter)
	context.Set(r, "Resource", "SearchParameter")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/SearchParameter/"+i.Hex())
}

func SearchParameterUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	searchparameter := &models.SearchParameter{}
	err := decoder.Decode(searchparameter)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("searchparameters")
	searchparameter.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, searchparameter)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting searchparameter update context")
	context.Set(r, "SearchParameter", searchparameter)
	context.Set(r, "Resource", "SearchParameter")
	context.Set(r, "Action", "update")
}

func SearchParameterDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("searchparameters")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting searchparameter delete context")
	context.Set(r, "SearchParameter", id.Hex())
	context.Set(r, "Resource", "SearchParameter")
	context.Set(r, "Action", "delete")
}
