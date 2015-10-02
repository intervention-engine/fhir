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
	"github.com/intervention-engine/fhir/search"
	"gopkg.in/mgo.v2/bson"
)

func MedicationOrderIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if r := recover(); r != nil {
			rw.Header().Set("Content-Type", "application/json; charset=utf-8")
			switch x := r.(type) {
			case search.Error:
				rw.WriteHeader(x.HTTPStatus)
				json.NewEncoder(rw).Encode(x.OperationOutcome)
				return
			default:
				outcome := &models.OperationOutcome{
					Issue: []models.OperationOutcomeIssueComponent{
						models.OperationOutcomeIssueComponent{
							Severity: "fatal",
							Code:     "exception",
						},
					},
				}
				rw.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(rw).Encode(outcome)
			}
		}
	}()

	var result []models.MedicationOrder
	c := Database.C("medicationorders")

	r.ParseForm()
	if len(r.Form) == 0 {
		iter := c.Find(nil).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	} else {
		searcher := search.NewMongoSearcher(Database)
		query := search.Query{Resource: "MedicationOrder", Query: r.URL.RawQuery}
		err := searcher.CreateQuery(query).All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}

	var medicationorderEntryList []models.BundleEntryComponent
	for i := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &result[i]
		medicationorderEntryList = append(medicationorderEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = medicationorderEntryList

	log.Println("Setting medicationorder search context")
	context.Set(r, "MedicationOrder", result)
	context.Set(r, "Resource", "MedicationOrder")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadMedicationOrder(r *http.Request) (*models.MedicationOrder, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("medicationorders")
	result := models.MedicationOrder{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting medicationorder read context")
	context.Set(r, "MedicationOrder", result)
	context.Set(r, "Resource", "MedicationOrder")
	return &result, nil
}

func MedicationOrderShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadMedicationOrder(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "MedicationOrder"))
}

func MedicationOrderCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	medicationorder := &models.MedicationOrder{}
	err := decoder.Decode(medicationorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medicationorders")
	i := bson.NewObjectId()
	medicationorder.Id = i.Hex()
	err = c.Insert(medicationorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting medicationorder create context")
	context.Set(r, "MedicationOrder", medicationorder)
	context.Set(r, "Resource", "MedicationOrder")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Add("Location", "http://"+host+":3001/MedicationOrder/"+i.Hex())
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(medicationorder)
}

func MedicationOrderUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	medicationorder := &models.MedicationOrder{}
	err := decoder.Decode(medicationorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medicationorders")
	medicationorder.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, medicationorder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting medicationorder update context")
	context.Set(r, "MedicationOrder", medicationorder)
	context.Set(r, "Resource", "MedicationOrder")
	context.Set(r, "Action", "update")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(medicationorder)
}

func MedicationOrderDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("medicationorders")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting medicationorder delete context")
	context.Set(r, "MedicationOrder", id.Hex())
	context.Set(r, "Resource", "MedicationOrder")
	context.Set(r, "Action", "delete")
}
