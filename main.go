package main

import "github.com/intervention-engine/fhir/server"

func main() {
	server.NewServer("mongodb://localhost:27017").Run(server.DefaultConfig)
}
