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

func PaymentReconciliationIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.PaymentReconciliation
	c := Database.C("paymentreconciliations")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var paymentreconciliationEntryList []models.PaymentReconciliationBundleEntry
	for _, paymentreconciliation := range result {
		var entry models.PaymentReconciliationBundleEntry
		entry.Id = paymentreconciliation.Id
		entry.Resource = paymentreconciliation
		paymentreconciliationEntryList = append(paymentreconciliationEntryList, entry)
	}

	var bundle models.PaymentReconciliationBundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Total = len(result)
	bundle.Entry = paymentreconciliationEntryList

	log.Println("Setting paymentreconciliation search context")
	context.Set(r, "PaymentReconciliation", result)
	context.Set(r, "Resource", "PaymentReconciliation")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadPaymentReconciliation(r *http.Request) (*models.PaymentReconciliation, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("paymentreconciliations")
	result := models.PaymentReconciliation{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting paymentreconciliation read context")
	context.Set(r, "PaymentReconciliation", result)
	context.Set(r, "Resource", "PaymentReconciliation")
	return &result, nil
}

func PaymentReconciliationShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadPaymentReconciliation(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "PaymentReconciliation"))
}

func PaymentReconciliationCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	paymentreconciliation := &models.PaymentReconciliation{}
	err := decoder.Decode(paymentreconciliation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("paymentreconciliations")
	i := bson.NewObjectId()
	paymentreconciliation.Id = i.Hex()
	err = c.Insert(paymentreconciliation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting paymentreconciliation create context")
	context.Set(r, "PaymentReconciliation", paymentreconciliation)
	context.Set(r, "Resource", "PaymentReconciliation")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/PaymentReconciliation/"+i.Hex())
}

func PaymentReconciliationUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	paymentreconciliation := &models.PaymentReconciliation{}
	err := decoder.Decode(paymentreconciliation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("paymentreconciliations")
	paymentreconciliation.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, paymentreconciliation)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting paymentreconciliation update context")
	context.Set(r, "PaymentReconciliation", paymentreconciliation)
	context.Set(r, "Resource", "PaymentReconciliation")
	context.Set(r, "Action", "update")
}

func PaymentReconciliationDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("paymentreconciliations")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting paymentreconciliation delete context")
	context.Set(r, "PaymentReconciliation", id.Hex())
	context.Set(r, "Resource", "PaymentReconciliation")
	context.Set(r, "Action", "delete")
}
