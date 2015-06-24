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
	resourceCount, patientCount, encounterCount, conditionCount, immunizationCount, observationCount, procedureCount, diagnosticReportCount, medicationCount, medicationStatementCount := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		output := "Created"
		decoder := json.NewDecoder(r.Body)
		switch {
		case strings.Contains(r.RequestURI, "Patient"):
			if isValid(decoder, &models.Patient{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Patient/%d", resourceCount))
				patientCount++
			}
		case strings.Contains(r.RequestURI, "Encounter"):
			if isValid(decoder, &models.Encounter{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Encounter/%d", resourceCount))
				encounterCount++
			}
		case strings.Contains(r.RequestURI, "Condition"):
			if isValid(decoder, &models.Condition{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Condition/%d", resourceCount))
				conditionCount++
			}
		case strings.Contains(r.RequestURI, "Immunization"):
			if isValid(decoder, &models.Immunization{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Immunization/%d", resourceCount))
				immunizationCount++
			}
		case strings.Contains(r.RequestURI, "Observation"):
			if isValid(decoder, &models.Observation{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Observation/%d", resourceCount))
				observationCount++
			}
		case strings.Contains(r.RequestURI, "Procedure"):
			if isValid(decoder, &models.Procedure{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Procedure/%d", resourceCount))
				procedureCount++
			}
		case strings.Contains(r.RequestURI, "DiagnosticReport"):
			if isValid(decoder, &models.DiagnosticReport{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/DiagnosticReport/%d", resourceCount))
				diagnosticReportCount++
			}
		case strings.Contains(r.RequestURI, "MedicationStatement"):
			if isValid(decoder, &models.MedicationStatement{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/MedicationStatement/%d", resourceCount))
				medicationStatementCount++
			}
		case strings.Contains(r.RequestURI, "Medication"):
			if isValid(decoder, &models.Medication{}) {
				w.Header().Add("Location", fmt.Sprintf("http://localhost/Medication/%d", resourceCount))
				medicationCount++
			}
		}
		fmt.Fprintln(w, output)
		resourceCount++
	}))
	defer ts.Close()

	// Read in the data in FHIR format
	data, err := ioutil.ReadFile("../fixtures/john_peters.json")
	util.CheckErr(err)

	// Do a bunch of junk to properly unmarshal all the data
	// This is needed because:
	// 1. It's a homogenous array, so go doesn't unmarshal it very well
	// 2. The id attribute is configured to not be serialized/deserialized
	type IdAndType struct {
		Id   string `json:"id"`
		Type string `json:"resourceType"`
	}
	idsAndTypes := make([]IdAndType, 20)
	err = json.Unmarshal(data, &idsAndTypes)
	util.CheckErr(err)

	rawMessages := make([]json.RawMessage, 20)
	err = json.Unmarshal(data, &rawMessages)
	util.CheckErr(err)

	fhirmodels := make([]interface{}, 20)
	for i := range fhirmodels {
		var y interface{}
		switch idsAndTypes[i].Type {
		case "Patient":
			y = &models.Patient{Id: idsAndTypes[i].Id}
		case "Encounter":
			y = &models.Encounter{Id: idsAndTypes[i].Id}
		case "Condition":
			y = &models.Condition{Id: idsAndTypes[i].Id}
		case "Observation":
			y = &models.Observation{Id: idsAndTypes[i].Id}
		case "DiagnosticReport":
			y = &models.DiagnosticReport{Id: idsAndTypes[i].Id}
		case "Procedure":
			y = &models.Procedure{Id: idsAndTypes[i].Id}
		case "Medication":
			y = &models.Medication{Id: idsAndTypes[i].Id}
		case "MedicationStatement":
			y = &models.MedicationStatement{Id: idsAndTypes[i].Id}
		case "Immunization":
			y = &models.Immunization{Id: idsAndTypes[i].Id}
		default:
			c.Errorf("Unexpected Type: %s", idsAndTypes[i].Type)
		}
		json.Unmarshal(rawMessages[i], y)
		fhirmodels[i] = y
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
	c.Assert(resourceCount, Equals, 20)
	c.Assert(medicationStatementCount, Equals, 1)
	c.Assert(medicationCount, Equals, 1)

	c.Assert(len(refMap), Equals, 20)
	c.Assert(refMap[idsAndTypes[0].Id], Equals, "http://localhost/Patient/0")
	c.Assert(fhirmodels[0].(*models.Patient).Id, Equals, "0")
}

func (s *UploadSuite) TestExternalReferences(c *C) {
	// Setup the mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Location", "http://localhost/Condition/abc")
		fmt.Fprintln(w, "Created")
	}))
	defer ts.Close()

	condition := &models.Condition{Id: "123"}
	condition.Patient = &models.Reference{Reference: "http://localhost/Patient/0"}

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
				w.Header().Add("Location", "http://localhost/Patient/1")
			}
		case strings.Contains(r.RequestURI, "Condition"):
			if isValid(decoder, &models.Condition{}) {
				w.Header().Add("Location", "http://localhost/Condition/1")
			}
		}
		fmt.Fprintln(w, output)
		resourceCount++
	}))
	defer ts.Close()

	patient := &models.Patient{Id: "a1"}
	condition := &models.Condition{Id: "b2"}
	condition.Patient = &models.Reference{Reference: "cid:a1"}

	// Upload the resources in the wrong order
	refMap, err := UploadResources([]interface{}{condition, patient}, ts.URL)
	util.CheckErr(err)

	// Assert that it processed all resources
	c.Assert(len(refMap), Equals, 2)
	c.Assert(refMap["a1"], Equals, "http://localhost/Patient/1")
	c.Assert(refMap["b2"], Equals, "http://localhost/Condition/1")
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
			match, _ := regexp.MatchString("\\Ahttp://localhost/[^/]+/[0-9a-f]+\\z", ref.Reference)
			if !match {
				fmt.Printf("Invalid reference: %s", ref.Reference)
				return false
			}
		}
		return len(refs) > 0
	}
	return true
}
