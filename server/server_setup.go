package server

import (
	"github.com/codegangsta/negroni"
	"gopkg.in/mgo.v2"
	"log"
)

type FHIRServer struct {
	Middleware   []negroni.Handler
	DatabaseHost string
}

func (f *FHIRServer) AddMiddleware(middleware negroni.Handler) {
	f.Middleware = append(f.Middleware, middleware)
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
	Router.StrictSlash(true)
	Router.KeepContext = true

	RegisterRoutes()

	n := negroni.Classic()
	for _, m := range f.Middleware {
		n.Use(m)
	}
	n.UseHandler(Router)
	n.Run(":3001")
}
