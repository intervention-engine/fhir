package main

import (
	"flag"
	"fmt"
	"github.com/intervention-engine/fhir/server"
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
// ================================================================================================
// 1.  GET    /Patient           -- verify that no interceptor is called
// 2.  GET    /Condition         -- verify that no interceptor is called
// 3.  POST   /Patient           -- verify that BOTH Create interceptors are called
// 4.  POST   /Condition         -- verify that only the TestUniversalCreateInterceptor is called
// 5.  PUT    /Patient/:id       -- verify that BOTH Update interceptors are called
// 6.  PUT    /Condition/:id     -- verify that only the TestUniversalUpdateInterceptor is called
// 7.  DELETE /Patient/:id       -- verify that BOTH Delete interceptors are called
// 8.  DELETE /Condition/:id     -- verify that only the TestUniversalDeleteInterceptor is called
// 9-10: repeat steps 3 and 4
// 11: PUT    /Patient?_id=:id   -- verify that BOTH Update interceptors are called
// 12: PUT    /Condition?_id=:id -- verify that only the TestUniversalUpdateInterceptor is called
// 13: DELETE /Patient?_id=:id   -- verify that BOTH Delete interceptors are called
// 14: DELETE /Condition?_id=:id -- verify that only the TestUniversalDeleteInterceptor is called
// ================================================================================================
// Next, run ./test -noint (run the test server without any interceptors) and verify that
// the new interceptor logic does not interfere with normal server operation.
//
// You can get sample Patient and Condition JSON objects to PUT/POST from:
// https://syntheticmass.mitre.org/fhir/baseDstu3/Patient
// https://syntheticmass.mitre.org/fhir/baseDstu3/Condition
//
func setupTestInterceptors(s *server.FHIRServer) {
	s.AddInterceptor("Create", "Patient", &TestPatientCreateInterceptor{})
	s.AddInterceptor("Update", "Patient", &TestPatientUpdateInterceptor{})
	s.AddInterceptor("Delete", "Patient", &TestPatientDeleteInterceptor{})
	s.AddInterceptor("Create", "*", &TestUniversalCreateInterceptor{})
	s.AddInterceptor("Update", "*", &TestUniversalUpdateInterceptor{})
	s.AddInterceptor("Delete", "*", &TestUniversalDeleteInterceptor{})
}

// Interceptors that will be registered to operate on Patient resources only:
// ----------------------------------------------------------------------------

// TestPatientCreateInterceptor operates on a Patient resource after it is created
type TestPatientCreateInterceptor struct{}

func (s *TestPatientCreateInterceptor) Before(resource interface{}) {}

func (s *TestPatientCreateInterceptor) After(resource interface{}) {
	fmt.Println("TestPatientCreateInterceptor: After()")
}

func (s *TestPatientCreateInterceptor) OnError(err error, resource interface{}) {}

// TestPatientUpdateInterceptor operates on a Patient resource both before and
// after it is updated
type TestPatientUpdateInterceptor struct{}

func (s *TestPatientUpdateInterceptor) Before(resource interface{}) {
	fmt.Println("TestPatientUpdateInterceptor: Before()")
}

func (s *TestPatientUpdateInterceptor) After(resource interface{}) {
	fmt.Println("TestPatientUpdateInterceptor: After()")
}

func (s *TestPatientUpdateInterceptor) OnError(err error, resource interface{}) {}

// TestPatientDeleteInterceptor operates on a Patient resource only before it is deleted
type TestPatientDeleteInterceptor struct{}

func (s *TestPatientDeleteInterceptor) Before(resource interface{}) {
	fmt.Println("TestPatientDeleteInterceptor: Before()")
}

func (s *TestPatientDeleteInterceptor) After(resource interface{}) {}

func (s *TestPatientDeleteInterceptor) OnError(err error, resource interface{}) {}

// Interceptors that will be registered to operate on ANY resource:
// ----------------------------------------------------------------------------

// TestUniversalCreateInterceptor operates on any resource after it is created
type TestUniversalCreateInterceptor struct{}

func (s *TestUniversalCreateInterceptor) Before(resource interface{}) {}

func (s *TestUniversalCreateInterceptor) After(resource interface{}) {
	fmt.Println("TestUniversalCreateInterceptor: After()")
}

func (s *TestUniversalCreateInterceptor) OnError(err error, resource interface{}) {}

// TestUniversalUpdateInterceptor operates on any resource both before and after
// it is updated
type TestUniversalUpdateInterceptor struct{}

func (s *TestUniversalUpdateInterceptor) Before(resource interface{}) {
	fmt.Println("TestUniversalUpdateInterceptor: Before()")
}

func (s *TestUniversalUpdateInterceptor) After(resource interface{}) {
	fmt.Println("TestUniversalUpdateInterceptor: After()")
}

func (s *TestUniversalUpdateInterceptor) OnError(err error, resource interface{}) {}

// TestUniversalDeleteInterceptor operates on any resource after it is deleted
type TestUniversalDeleteInterceptor struct{}

func (s *TestUniversalDeleteInterceptor) Before(resource interface{}) {}

func (s *TestUniversalDeleteInterceptor) After(resource interface{}) {
	fmt.Println("TestUniversalDeleteInterceptor: After()")
}

func (s *TestUniversalDeleteInterceptor) OnError(err error, resource interface{}) {}
