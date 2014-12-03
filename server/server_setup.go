package server

import (
	"log"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

type FHIRServer struct {
	DatabaseHost     string
	Router           *mux.Router
	MiddlewareConfig map[string][]negroni.Handler
}

func (f *FHIRServer) AddMiddleware(key string, middleware negroni.Handler) {
	f.MiddlewareConfig[key] = append(f.MiddlewareConfig[key], middleware)
}

func (f *FHIRServer) Run() {
	var err error

	// Setup the database
	if MongoSession, err = mgo.Dial(f.DatabaseHost); err != nil {
		panic(err)
	}
	log.Println("Connected to mongodb")
	defer MongoSession.Close()

	Database = MongoSession.DB("fhir")
	f.Router = mux.NewRouter()
	f.Router.StrictSlash(true)
	f.Router.KeepContext = true

	RegisterRoutes(f.Router, f.MiddlewareConfig)

	n := negroni.Classic()
	// for _, m := range f.Middleware {
	// 	n.Use(m)
	// }
	n.UseHandler(f.Router)
	n.Run(":3001")
}
