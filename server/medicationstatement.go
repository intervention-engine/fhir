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

func MedicationStatementIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.MedicationStatement
	c := Database.C("medicationstatements")

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

	var medicationstatementEntryList []models.MedicationStatementBundleEntry
	for _, medicationstatement := range result {
		var entry models.MedicationStatementBundleEntry
		entry.Id = medicationstatement.Id
		entry.Resource = medicationstatement
		medicationstatementEntryList = append(medicationstatementEntryList, entry)
	}

	var bundle models.MedicationStatementBundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Total = len(result)
	bundle.Entry = medicationstatementEntryList

	log.Println("Setting medicationstatement search context")
	context.Set(r, "MedicationStatement", result)
	context.Set(r, "Resource", "MedicationStatement")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadMedicationStatement(r *http.Request) (*models.MedicationStatement, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("medicationstatements")
	result := models.MedicationStatement{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting medicationstatement read context")
	context.Set(r, "MedicationStatement", result)
	context.Set(r, "Resource", "MedicationStatement")
	return &result, nil
}

func MedicationStatementShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadMedicationStatement(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "MedicationStatement"))
}

func MedicationStatementCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	medicationstatement := &models.MedicationStatement{}
	err := decoder.Decode(medicationstatement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medicationstatements")
	i := bson.NewObjectId()
	medicationstatement.Id = i.Hex()
	err = c.Insert(medicationstatement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting medicationstatement create context")
	context.Set(r, "MedicationStatement", medicationstatement)
	context.Set(r, "Resource", "MedicationStatement")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/MedicationStatement/"+i.Hex())
}

func MedicationStatementUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	medicationstatement := &models.MedicationStatement{}
	err := decoder.Decode(medicationstatement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medicationstatements")
	medicationstatement.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, medicationstatement)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting medicationstatement update context")
	context.Set(r, "MedicationStatement", medicationstatement)
	context.Set(r, "Resource", "MedicationStatement")
	context.Set(r, "Action", "update")
}

func MedicationStatementDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("medicationstatements")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting medicationstatement delete context")
	context.Set(r, "MedicationStatement", id.Hex())
	context.Set(r, "Resource", "MedicationStatement")
	context.Set(r, "Action", "delete")
}
