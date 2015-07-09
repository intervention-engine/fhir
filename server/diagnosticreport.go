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

func DiagnosticReportIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.DiagnosticReport
	c := Database.C("diagnosticreports")

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
			if splitKey[0] == "subject" {
				err := c.Find(bson.M{"subject.referenceid": value[0]}).All(&result)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}

	var diagnosticreportEntryList []models.DiagnosticReportBundleEntry
	for _, diagnosticreport := range result {
		var entry models.DiagnosticReportBundleEntry
		entry.Id = diagnosticreport.Id
		entry.Resource = diagnosticreport
		diagnosticreportEntryList = append(diagnosticreportEntryList, entry)
	}

	var bundle models.DiagnosticReportBundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Total = len(result)
	bundle.Entry = diagnosticreportEntryList

	log.Println("Setting diagnosticreport search context")
	context.Set(r, "DiagnosticReport", result)
	context.Set(r, "Resource", "DiagnosticReport")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadDiagnosticReport(r *http.Request) (*models.DiagnosticReport, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("diagnosticreports")
	result := models.DiagnosticReport{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting diagnosticreport read context")
	context.Set(r, "DiagnosticReport", result)
	context.Set(r, "Resource", "DiagnosticReport")
	return &result, nil
}

func DiagnosticReportShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadDiagnosticReport(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "DiagnosticReport"))
}

func DiagnosticReportCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	diagnosticreport := &models.DiagnosticReport{}
	err := decoder.Decode(diagnosticreport)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("diagnosticreports")
	i := bson.NewObjectId()
	diagnosticreport.Id = i.Hex()
	err = c.Insert(diagnosticreport)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting diagnosticreport create context")
	context.Set(r, "DiagnosticReport", diagnosticreport)
	context.Set(r, "Resource", "DiagnosticReport")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/DiagnosticReport/"+i.Hex())
}

func DiagnosticReportUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	diagnosticreport := &models.DiagnosticReport{}
	err := decoder.Decode(diagnosticreport)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("diagnosticreports")
	diagnosticreport.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, diagnosticreport)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting diagnosticreport update context")
	context.Set(r, "DiagnosticReport", diagnosticreport)
	context.Set(r, "Resource", "DiagnosticReport")
	context.Set(r, "Action", "update")
}

func DiagnosticReportDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("diagnosticreports")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting diagnosticreport delete context")
	context.Set(r, "DiagnosticReport", id.Hex())
	context.Set(r, "Resource", "DiagnosticReport")
	context.Set(r, "Action", "delete")
}
