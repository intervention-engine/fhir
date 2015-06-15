Intervention Engine FHIR Server
===============================

[![Build Status](https://travis-ci.org/intervention-engine/fhir.svg?branch=master)](https://travis-ci.org/intervention-engine/fhir)

This project provides [HL7 FHIR](http://hl7.org/implement/standards/fhir/) models and a generic
FHIR server implemented in Go. The FHIR server uses MongoDB as storage and contains a small
extension of the Reference model in order to more readily support queries in MongoDB.

Environment
-----------

This project currently uses Go 1.3.3 and is built using the Go toolchain.

To install Go, follow the instructions found at the [Go Website](http://golang.org/doc/install).

Following standard Go practices, you should clone this project to:

    $GOPATH/src/github.com/intervention-engine/fhir

To get all of the dependencies for this project, run:

    go get

in this directory.

To run all of the tests for this project, run:

    go test ./...

in this directory.

This project also requires MongoDB 2.6.* or higher. To install MongoDB, refer to the
[MongoDB installation guide](http://docs.mongodb.org/manual/installation/).

To start the server, simply run server.go:

    go run server.go

Custom Middleware
-----------------

Because this project is a generic FHIR server, it only supports simple CRUD methods for FHIR
resources as-is. In order to provide extensibility, the FHIRServer type has a method called
AddMiddleware that can be called as follows:

    s := server.NewServer("localhost")
    s.AddMiddleware(negroni.HandlerFunc(MyHandler))

where MyHandler is the middleware function that you want to add.

License
-------

Copyright 2015 The MITRE Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
