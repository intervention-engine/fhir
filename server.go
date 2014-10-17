package main

import (
	"net/http"
	"gitlab.mitre.org/fhir/server"
	"gopkg.in/mgo.v2"
	"log"
	"fmt"
)

func main() {
	var err error

	// Setup the database
	if server.MongoSession, err = mgo.Dial("localhost"); err != nil {
		panic(err)
	}
	log.Println("Connected to mongodb")
	defer server.MongoSession.Close()

	server.Database = server.MongoSession.DB("fhir")

  server.Router.HandleFunc("/", HomeHandler)

	patient := server.Router.PathPrefix("/patient/{id}").Subrouter()
	patient.Methods("GET").HandlerFunc(server.PatientShowHandler)
	patient.Methods("PUT").HandlerFunc(server.PatientCreateHandler)
	patient.Methods("POST").HandlerFunc(server.PatientUpdateHandler)
	patient.Methods("DELETE").HandlerFunc(server.PatientDeleteHandler)

	http.ListenAndServe(":8080", server.Router)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(rw, "FHIR Server Yay! \\o/")
}
