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

func DataElementIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.DataElement
	c := Database.C("dataelements")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var dataelementEntryList []models.DataElementBundleEntry
	for _, dataelement := range result {
		var entry models.DataElementBundleEntry
		entry.Id = dataelement.Id
		entry.Resource = dataelement
		dataelementEntryList = append(dataelementEntryList, entry)
	}

	var bundle models.DataElementBundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Total = len(result)
	bundle.Entry = dataelementEntryList

	log.Println("Setting dataelement search context")
	context.Set(r, "DataElement", result)
	context.Set(r, "Resource", "DataElement")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadDataElement(r *http.Request) (*models.DataElement, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("dataelements")
	result := models.DataElement{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting dataelement read context")
	context.Set(r, "DataElement", result)
	context.Set(r, "Resource", "DataElement")
	return &result, nil
}

func DataElementShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadDataElement(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "DataElement"))
}

func DataElementCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	dataelement := &models.DataElement{}
	err := decoder.Decode(dataelement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("dataelements")
	i := bson.NewObjectId()
	dataelement.Id = i.Hex()
	err = c.Insert(dataelement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting dataelement create context")
	context.Set(r, "DataElement", dataelement)
	context.Set(r, "Resource", "DataElement")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/DataElement/"+i.Hex())
}

func DataElementUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	dataelement := &models.DataElement{}
	err := decoder.Decode(dataelement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("dataelements")
	dataelement.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, dataelement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting dataelement update context")
	context.Set(r, "DataElement", dataelement)
	context.Set(r, "Resource", "DataElement")
	context.Set(r, "Action", "update")
}

func DataElementDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("dataelements")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting dataelement delete context")
	context.Set(r, "DataElement", id.Hex())
	context.Set(r, "Resource", "DataElement")
	context.Set(r, "Action", "delete")
}
