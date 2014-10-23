package server

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var (
	MongoSession *mgo.Session
	Database     *mgo.Database

	Router = mux.NewRouter()
)
