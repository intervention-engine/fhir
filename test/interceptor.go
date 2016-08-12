package main

import (
	"flag"
	"fmt"
	"github.com/intervention-engine/fhir/server"
	"reflect"
)

func main() {
	noint := flag.Bool("noint", false, "Run the test server without interceptors")
	flag.Parse()
	s := server.NewServer("localhost")

	if !*noint {
		setupTestInterceptors(s)
	}

	config := server.Config{UseSmartAuth: false}
	s.Run(config)
}

// With this test server running, verfiy the following (by viewing server log):
// =======================================================================================
// 1.  GET    /Patient           -- verify that no interceptor is called
// 2.  GET    /Condition         -- verify that no interceptor is called
// 3.  POST   /Patient           -- verify that BOTH POST interceptors are called
// 4.  POST   /Condition         -- verify that only the allPostsInterceptor is called
// 5.  PUT    /Patient/:id       -- verify that BOTH PUT interceptors are called
// 6.  PUT    /Condition/:id     -- verify that only the allPutsInterceptor is called
// 7.  DELETE /Patient/:id       -- verify that BOTH DELETE interceptors are called
// 8.  DELETE /Condition/:id     -- verify that only the allDeletesInterceptor is called
// 9-10: repeat steps 3 and 4
// 11: PUT    /Patient?_id=:id   -- verify that BOTH PUT interceptors are called
// 12: PUT    /Condition?_id=:id -- verify that only the allPutsInterceptor is called
// 13: DELETE /Patient?_id=:id   -- verify that BOTH DELETE interceptors are called
// 14: DELETE /Condition?_id=:id -- verify that only the allDeletesInterceptor is called
// =======================================================================================
// Next, run ./test -noint (run the test server without any interceptors) and verify that
// the new interceptor logic does not interfere with normal server operation.
//
// You can get sample Patient and Condition JSON objects to PUT/POST from:
// https://syntheticmass.mitre.org/fhir/baseDstu3/Patient
// https://syntheticmass.mitre.org/fhir/baseDstu3/Condition
//
func setupTestInterceptors(s *server.FHIRServer) {
	s.AddInterceptor("POST", "Patient", postInterceptor)
	s.AddInterceptor("POST", "*", allPostsInterceptor)

	s.AddInterceptor("PUT", "Patient", putInterceptor)
	s.AddInterceptor("PUT", "*", allPutsInterceptor)

	s.AddInterceptor("DELETE", "Patient", deleteInterceptor)
	s.AddInterceptor("DELETE", "*", allDeletesInterceptor)
}

func postInterceptor(resource interface{}) {
	fmt.Printf("POST intercepted for resource: %s\n", getResourceType(resource))
}

func allPostsInterceptor(resource interface{}) {
	fmt.Printf("POST intercepted for ALL resources\n")
}

func putInterceptor(resource interface{}) {
	fmt.Printf("PUT intercepted for resource: %s\n", getResourceType(resource))
}

func allPutsInterceptor(resource interface{}) {
	fmt.Printf("PUT intercepted for ALL resources\n")
}

func deleteInterceptor(resource interface{}) {
	fmt.Printf("DELETE intercepted for resource: %s\n", getResourceType(resource))
}

func allDeletesInterceptor(resource interface{}) {
	fmt.Printf("DELETE intercepted for ALL resources\n")
}

func getResourceType(resource interface{}) string {
	resType := reflect.TypeOf(resource).Elem().Name()
	return resType
}
