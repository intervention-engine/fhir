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

func MediaIndexHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var result []models.Media
	c := Database.C("medias")

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

	var mediaEntryList []models.BundleEntryComponent
	for _, media := range result {
		var entry models.BundleEntryComponent
		entry.Resource = &media
		mediaEntryList = append(mediaEntryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	var total = uint32(len(result))
	bundle.Total = &total
	bundle.Entry = mediaEntryList

	log.Println("Setting media search context")
	context.Set(r, "Media", result)
	context.Set(r, "Resource", "Media")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(&bundle)
}

func LoadMedia(r *http.Request) (*models.Media, error) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		return nil, errors.New("Invalid id")
	}

	c := Database.C("medias")
	result := models.Media{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		return nil, err
	}

	log.Println("Setting media read context")
	context.Set(r, "Media", result)
	context.Set(r, "Resource", "Media")
	return &result, nil
}

func MediaShowHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "Action", "read")
	_, err := LoadMedia(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(context.Get(r, "Media"))
}

func MediaCreateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	decoder := json.NewDecoder(r.Body)
	media := &models.Media{}
	err := decoder.Decode(media)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medias")
	i := bson.NewObjectId()
	media.Id = i.Hex()
	err = c.Insert(media)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting media create context")
	context.Set(r, "Media", media)
	context.Set(r, "Resource", "Media")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":3001/Media/"+i.Hex())
}

func MediaUpdateHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	media := &models.Media{}
	err := decoder.Decode(media)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medias")
	media.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, media)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting media update context")
	context.Set(r, "Media", media)
	context.Set(r, "Resource", "Media")
	context.Set(r, "Action", "update")
}

func MediaDeleteHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("medias")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting media delete context")
	context.Set(r, "Media", id.Hex())
	context.Set(r, "Resource", "Media")
	context.Set(r, "Action", "delete")
}
