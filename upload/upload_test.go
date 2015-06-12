package upload

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/hdsfhir"
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
)

type UploadSuite struct {
	JSONBlob []byte
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func (s *UploadSuite) SetUpSuite(c *C) {
	data, err := ioutil.ReadFile("../fixtures/john_peters_hds.json")
	util.CheckErr(err)
	s.JSONBlob = data
}

var _ = Suite(&UploadSuite{})

func (s *UploadSuite) TestPostToFHIRServer(c *C) {
	patient := &hdsfhir.Patient{}
	err := json.Unmarshal(s.JSONBlob, patient)
	util.CheckErr(err)
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
	refMap, err := UploadResources(patient.FHIRModels(), ts.URL)
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
	c.Assert(refMap[patient.GetTempID()], Equals, "http://localhost/Patient/0")
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

func getAllReferences(model interface{}) []*models.Reference {
	refs := make([]*models.Reference, 0)
	s := reflect.ValueOf(model).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.Type() == reflect.TypeOf(&models.Reference{}) && !f.IsNil() {
			refs = append(refs, f.Interface().(*models.Reference))
		} else if f.Type() == reflect.TypeOf([]models.Reference{}) {
			for j := 0; j < f.Len(); j++ {
				refs = append(refs, f.Index(j).Addr().Interface().(*models.Reference))
			}
		}
	}
	return refs
}
