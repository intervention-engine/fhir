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

func QuestionnaireResponseIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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

	var result []models.QuestionnaireResponse
	c := Database.C("questionnaireresponses")

	r.ParseForm()
	if len(r.Form) == 0 {
		iter := c.Find(nil).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	} else {
		searcher := search.NewMongoSearcher(Database)
		query := search.Query{Resource: "QuestionnaireResponse", Query: r.URL.RawQuery}
		err := searcher.CreateQuery(query).All(&result)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}

	var questionnaireresponseEntryList []models.BundleEntryComponent
	for i := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &result[i]
		questionnaireresponseEntryList = append(questionnaireresponseEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = questionnaireresponseEntryList

	log.Println("Setting questionnaireresponse search context")
	context.Set(r, "QuestionnaireResponse", result)
	context.Set(r, "Resource", "QuestionnaireResponse")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadQuestionnaireResponse(r *http.Request) (*models.QuestionnaireResponse, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("questionnaireresponses")
	result := models.QuestionnaireResponse{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting questionnaireresponse read context")
	context.Set(r, "QuestionnaireResponse", result)
	context.Set(r, "Resource", "QuestionnaireResponse")
	return &result, nil
}

func QuestionnaireResponseShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadQuestionnaireResponse(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "QuestionnaireResponse"))
}

func QuestionnaireResponseCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	questionnaireresponse := &models.QuestionnaireResponse{}
	err := decoder.Decode(questionnaireresponse)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("questionnaireresponses")
	i := bson.NewObjectId()
	questionnaireresponse.Id = i.Hex()
	err = c.Insert(questionnaireresponse)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting questionnaireresponse create context")
	context.Set(r, "QuestionnaireResponse", questionnaireresponse)
	context.Set(r, "Resource", "QuestionnaireResponse")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Add("Location", "http://"+host+":3001/QuestionnaireResponse/"+i.Hex())
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(questionnaireresponse)
}

func QuestionnaireResponseUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	questionnaireresponse := &models.QuestionnaireResponse{}
	err := decoder.Decode(questionnaireresponse)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("questionnaireresponses")
	questionnaireresponse.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, questionnaireresponse)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting questionnaireresponse update context")
	context.Set(r, "QuestionnaireResponse", questionnaireresponse)
	context.Set(r, "Resource", "QuestionnaireResponse")
	context.Set(r, "Action", "update")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(questionnaireresponse)
}

func QuestionnaireResponseDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("questionnaireresponses")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting questionnaireresponse delete context")
	context.Set(r, "QuestionnaireResponse", id.Hex())
	context.Set(r, "Resource", "QuestionnaireResponse")
	context.Set(r, "Action", "delete")
}
