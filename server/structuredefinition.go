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

func StructureDefinitionIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.StructureDefinition
	c := Database.C("structuredefinitions")
	r.ParseForm()
	if len(r.Form) == 0 {
		iter := c.Find(nil).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	} else {
		codeValue := r.Form.Get("code")
		splitCode := strings.Split(codeValue, "|")
		if len(splitCode) > 1 {
			codeSystem := splitCode[0]
			code := splitCode[1]
			err := c.Find(bson.M{"code.coding.system": codeSystem, "code.coding.code": code}).All(&result)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		}
	}

	var structuredefinitionEntryList []models.BundleEntryComponent
	for _, structuredefinition := range result {
		var entry models.BundleEntryComponent
		entry.Resource = structuredefinition
		structuredefinitionEntryList = append(structuredefinitionEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = structuredefinitionEntryList

	log.Println("Setting structuredefinition search context")
	context.Set(r, "StructureDefinition", result)
	context.Set(r, "Resource", "StructureDefinition")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadStructureDefinition(r *http.Request) (*models.StructureDefinition, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("structuredefinitions")
	result := models.StructureDefinition{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting structuredefinition read context")
	context.Set(r, "StructureDefinition", result)
	context.Set(r, "Resource", "StructureDefinition")
	return &result, nil
}

func StructureDefinitionShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadStructureDefinition(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "StructureDefinition"))
}

func StructureDefinitionCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	structuredefinition := &models.StructureDefinition{}
	err := decoder.Decode(structuredefinition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("structuredefinitions")
	i := bson.NewObjectId()
	structuredefinition.Id = i.Hex()
	err = c.Insert(structuredefinition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting structuredefinition create context")
	context.Set(r, "StructureDefinition", structuredefinition)
	context.Set(r, "Resource", "StructureDefinition")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/StructureDefinition/"+i.Hex())
}

func StructureDefinitionUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	structuredefinition := &models.StructureDefinition{}
	err := decoder.Decode(structuredefinition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("structuredefinitions")
	structuredefinition.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, structuredefinition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting structuredefinition update context")
	context.Set(r, "StructureDefinition", structuredefinition)
	context.Set(r, "Resource", "StructureDefinition")
	context.Set(r, "Action", "update")
}

func StructureDefinitionDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("structuredefinitions")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting structuredefinition delete context")
	context.Set(r, "StructureDefinition", id.Hex())
	context.Set(r, "Resource", "StructureDefinition")
	context.Set(r, "Action", "delete")
}
