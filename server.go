package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"gitlab.mitre.org/intervention-engine/fhir/ie"
	"gitlab.mitre.org/intervention-engine/fhir/server"
)

func main() {
	s := server.FHIRServer{DatabaseHost: "localhost", Middleware: make([]negroni.Handler, 0)}
	s.AddMiddleware(negroni.HandlerFunc(ie.PatientHandler))
  s.AddMiddleware(negroni.HandlerFunc(ie.FactHandler))

	s.Run()
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "FHIR Server Yay! \\o/")
}
