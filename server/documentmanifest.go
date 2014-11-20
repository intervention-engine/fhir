package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gitlab.mitre.org/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func DocumentManifestIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.DocumentManifest
	c := Database.C("documentmanifests")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var bundle models.DocumentManifestBundle
	bundle.Type = "Bundle"
	bundle.Title = "DocumentManifest Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entries = result

	log.Println("Setting documentmanifest search context")
	context.Set(r, "DocumentManifest", result)
	context.Set(r, "Resource", "DocumentManifest")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func DocumentManifestShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("documentmanifests")

	result := models.DocumentManifest{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting documentmanifest read context")
	context.Set(r, "DocumentManifest", result)
	context.Set(r, "Resource", "DocumentManifest")
	context.Set(r, "Action", "read")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(result)
}

func DocumentManifestCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	documentmanifest := &models.DocumentManifest{}
	err := decoder.Decode(documentmanifest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("documentmanifests")
	i := bson.NewObjectId()
	documentmanifest.Id = i.Hex()
	err = c.Insert(documentmanifest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting documentmanifest create context")
	context.Set(r, "DocumentManifest", documentmanifest)
	context.Set(r, "Resource", "DocumentManifest")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":8080/DocumentManifest/"+i.Hex())
}

func DocumentManifestUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	documentmanifest := &models.DocumentManifest{}
	err := decoder.Decode(documentmanifest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("documentmanifests")
	documentmanifest.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, documentmanifest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting documentmanifest update context")
	context.Set(r, "DocumentManifest", documentmanifest)
	context.Set(r, "Resource", "DocumentManifest")
	context.Set(r, "Action", "update")
}

func DocumentManifestDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("documentmanifests")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting documentmanifest delete context")
	context.Set(r, "DocumentManifest", id.Hex())
	context.Set(r, "Resource", "DocumentManifest")
	context.Set(r, "Action", "delete")
}
