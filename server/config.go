package server

import (
  "gopkg.in/mgo.v2"
  "github.com/gorilla/mux"
)

var (
  MongoSession *mgo.Session
  Database     *mgo.Database

  Router = mux.NewRouter()
)
