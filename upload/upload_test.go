package upload

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/intervention-engine/fhir/models"
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
)

type UploadSuite struct {
	JSONBlob []byte
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func (s *UploadSuite) SetUpSuite(c *C) {}

var _ = Suite(&UploadSuite{})

func (s *UploadSuite) TestPostToFHIRServer(c *C) {
	// Setup the mock server
	resourceCount, patientCount, encounterCount, conditionCount, immunizationCount, observationCount, procedureCount, diagnosticReportCount, medicationStatementCount := 0, 0, 0, 0, 0, 0, 0, 0, 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		output := "Created"
		decoder := json.NewDecoder(r.Body)
		switch {
		case strings.Contains(r.RequestURI, "Patient"):
			if isValid(decoder, &models.Patient{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Patient/%d/_history/1", resourceCount))
				patientCount++
			}
		case strings.Contains(r.RequestURI, "Encounter"):
			if isValid(decoder, &models.Encounter{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Encounter/%d/_history/1", resourceCount))
				encounterCount++
			}
		case strings.Contains(r.RequestURI, "Condition"):
			if isValid(decoder, &models.Condition{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Condition/%d/_history/1", resourceCount))
				conditionCount++
			}
		case strings.Contains(r.RequestURI, "Immunization"):
			if isValid(decoder, &models.Immunization{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Immunization/%d/_history/1", resourceCount))
				immunizationCount++
			}
		case strings.Contains(r.RequestURI, "Observation"):
			if isValid(decoder, &models.Observation{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Observation/%d/_history/1", resourceCount))
				observationCount++
			}
		case strings.Contains(r.RequestURI, "Procedure"):
			if isValid(decoder, &models.Procedure{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Procedure/%d/_history/1", resourceCount))
				procedureCount++
			}
		case strings.Contains(r.RequestURI, "DiagnosticReport"):
			if isValid(decoder, &models.DiagnosticReport{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/DiagnosticReport/%d/_history/1", resourceCount))
				diagnosticReportCount++
			}
		case strings.Contains(r.RequestURI, "MedicationStatement"):
			if isValid(decoder, &models.MedicationStatement{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/MedicationStatement/%d/_history/1", resourceCount))
				medicationStatementCount++
			}
		}
		fmt.Fprintln(w, output)
		resourceCount++
	}))
	defer ts.Close()

	// Read in the data in FHIR format
	data, err := ioutil.ReadFile("../fixtures/john_peters.json")
	util.CheckErr(err)

	maps := make([]interface{}, 19)
	err = json.Unmarshal(data, &maps)
	util.CheckErr(err)

	fhirmodels := make([]interface{}, 0, len(maps))
	for _, resourceMap := range maps {
		r := models.MapToResource(resourceMap, true)
		fhirmodels = append(fhirmodels, r)
	}

	// Upload the resources and check the counts
	refMap, err := UploadResources(fhirmodels, ts.URL)

	c.Assert(patientCount, Equals, 1)
	c.Assert(encounterCount, Equals, 4)
	c.Assert(conditionCount, Equals, 5)
	c.Assert(immunizationCount, Equals, 1)
	c.Assert(observationCount, Equals, 4)
	c.Assert(procedureCount, Equals, 2)
	c.Assert(diagnosticReportCount, Equals, 1)
	c.Assert(medicationStatementCount, Equals, 1)

	c.Assert(resourceCount, Equals, 19)
	c.Assert(len(refMap), Equals, 19)

	c.Assert(fhirmodels[0].(*models.Patient).Id, Equals, "0")
}

func (s *UploadSuite) TestExternalReferences(c *C) {
	// Setup the mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Location", "http://localhost/Condition/abc")
		fmt.Fprintln(w, "Created")
	}))
	defer ts.Close()

	condition := &models.Condition{}
	condition.Id = "123"
	condition.Patient = &models.Reference{Reference: "Patient/0"}

	// Upload the resource
	newId, err := UploadResource(condition, ts.URL)
	util.CheckErr(err)
	c.Assert(newId, Equals, "http://localhost/Condition/abc")

	// Upload the resources and check the counts
	refMap, err := UploadResources([]interface{}{condition}, ts.URL)
	util.CheckErr(err)
	c.Assert(len(refMap), Equals, 1)
}

func (s *UploadSuite) TestUnorderedDependencies(c *C) {
	// Setup the mock server
	resourceCount := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		output := "Created"
		decoder := json.NewDecoder(r.Body)
		switch {
		case strings.Contains(r.RequestURI, "Patient"):
			if isValid(decoder, &models.Patient{}) {
				w.Header().Add("Location", "http://localhost/Patient/1/_history/1")
			}
		case strings.Contains(r.RequestURI, "Condition"):
			if isValid(decoder, &models.Condition{}) {
				w.Header().Add("Location", "http://localhost/Condition/1") // Purposefully no _history to test this use case
			}
		}
		fmt.Fprintln(w, output)
		resourceCount++
	}))
	defer ts.Close()

	patient := &models.Patient{}
	patient.Id = "a1"
	condition := &models.Condition{}
	condition.Id = "b2"
	condition.Patient = &models.Reference{Reference: "cid:a1"}

	// Upload the resources in the wrong order
	refMap, err := UploadResources([]interface{}{condition, patient}, ts.URL)
	util.CheckErr(err)

	// Assert that it processed all resources and correctly mapped refs
	c.Assert(len(refMap), Equals, 2)
	c.Assert(refMap["a1"], Equals, "Patient/1")
	c.Assert(refMap["b2"], Equals, "Condition/1")
}

func isValid(decoder *json.Decoder, model interface{}) bool {
	err := decoder.Decode(model)
	if err != nil {
		return false
	}

	_, isPatient := model.(*models.Patient)
	_, isMedication := model.(*models.Medication)
	if !isPatient && !isMedication {
		refs := getAllReferences(model)
		for _, ref := range refs {
			match, _ := regexp.MatchString("\\A[^/]+/[0-9a-f]+\\z", ref.Reference)
			if !match {
				fmt.Printf("Invalid reference: %s", ref.Reference)
				return false
			} else if strings.Contains(ref.Reference, "/_history/") {
				fmt.Printf("Invalid reference (contains _history component): %s", ref.Reference)
				return false
			}
		}
		return len(refs) > 0
	}
	return true
}
