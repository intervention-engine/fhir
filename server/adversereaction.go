package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func AdverseReactionIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.AdverseReaction
	c := Database.C("adversereactions")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var adversereactionEntryList []models.AdverseReactionBundleEntry
	for _, adversereaction := range result {
		var entry models.AdverseReactionBundleEntry
		entry.Title = "AdverseReaction " + adversereaction.Id
		entry.Id = adversereaction.Id
		entry.Content = adversereaction
		adversereactionEntryList = append(adversereactionEntryList, entry)
	}

	var bundle models.AdverseReactionBundle
	bundle.Type = "Bundle"
	bundle.Title = "AdverseReaction Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entry = adversereactionEntryList

	log.Println("Setting adversereaction search context")
	context.Set(r, "AdverseReaction", result)
	context.Set(r, "Resource", "AdverseReaction")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func LoadAdverseReaction(r *http.Request) (*models.AdverseReaction, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("adversereactions")
	result := models.AdverseReaction{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting adversereaction read context")
	context.Set(r, "AdverseReaction", result)
	context.Set(r, "Resource", "AdverseReaction")
	return &result, nil
}

func AdverseReactionShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadAdverseReaction(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "AdverseReaction"))
}

func AdverseReactionCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	adversereaction := &models.AdverseReaction{}
	err := decoder.Decode(adversereaction)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("adversereactions")
	i := bson.NewObjectId()
	adversereaction.Id = i.Hex()
	err = c.Insert(adversereaction)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting adversereaction create context")
	context.Set(r, "AdverseReaction", adversereaction)
	context.Set(r, "Resource", "AdverseReaction")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/AdverseReaction/"+i.Hex())
}

func AdverseReactionUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	adversereaction := &models.AdverseReaction{}
	err := decoder.Decode(adversereaction)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("adversereactions")
	adversereaction.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, adversereaction)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting adversereaction update context")
	context.Set(r, "AdverseReaction", adversereaction)
	context.Set(r, "Resource", "AdverseReaction")
	context.Set(r, "Action", "update")
}

func AdverseReactionDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("adversereactions")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting adversereaction delete context")
	context.Set(r, "AdverseReaction", id.Hex())
	context.Set(r, "Resource", "AdverseReaction")
	context.Set(r, "Action", "delete")
}
