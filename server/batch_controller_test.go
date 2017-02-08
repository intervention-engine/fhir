package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/intervention-engine/fhir/models"
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BatchControllerSuite struct {
	initialSession *mgo.Session
	MasterSession  *MasterSession
	Engine         *gin.Engine
	Server         *httptest.Server
	Interceptors   map[string]InterceptorList
}

var _ = Suite(&BatchControllerSuite{})

func (s *BatchControllerSuite) SetUpSuite(c *C) {
	// Set gin to release mode because the first printout of all routes makes it hard to see what is failing
	gin.SetMode(gin.ReleaseMode)

	// Set up the database
	var err error
	s.initialSession, err = mgo.Dial("localhost")
	util.CheckErr(err)
	s.MasterSession = NewMasterSession(s.initialSession, "fhir-test")

	// Build routes for testing
	s.Engine = gin.New()
	RegisterRoutes(s.Engine, make(map[string][]gin.HandlerFunc), NewMongoDataAccessLayer(s.MasterSession, s.Interceptors), Config{})

	// Create httptest server
	s.Server = httptest.NewServer(s.Engine)
}

func (s *BatchControllerSuite) TearDownTest(c *C) {
	worker := s.MasterSession.GetWorkerSession()
	worker.DB().DropDatabase()
	worker.Close()
}

func (s *BatchControllerSuite) TearDownSuite(c *C) {
	s.initialSession.Close()
	s.Server.Close()
}

func (s *BatchControllerSuite) TestDeleteEntriesBundle(c *C) {

	worker := s.MasterSession.GetWorkerSession()
	defer worker.Close()

	// Put some records in the database to delete
	condition := &models.Condition{
		Subject: &models.Reference{Reference: "https://example.com/base/Patient/4954037112938410473"},
		Code: &models.CodeableConcept{
			Coding: []models.Coding{
				{System: "Foo", Code: "Bar"},
			},
		},
		VerificationStatus: "confirmed",
	}
	condition.Id = "56afe6b85cdc7ec329dfe6a1"
	condition2 := &models.Condition{
		Subject: &models.Reference{Reference: "https://example.com/base/Patient/4954037112938410473"},
		Code: &models.CodeableConcept{
			Coding: []models.Coding{
				{System: "Foo", Code: "Baz"},
			},
		},
		VerificationStatus: "confirmed",
	}
	condition2.Id = "56afe6b85cdc7ec329dfe6a2"
	encounter := &models.Encounter{
		Status: "finished",
	}
	encounter.Id = "56afe6b85cdc7ec329dfe6a3"
	encounter2 := &models.Encounter{
		Status: "finished",
	}
	encounter2.Id = "56afe6b85cdc7ec329dfe6a4"

	// Insert the conditions and encounters into the db
	condCollection := worker.DB().C("conditions")
	err := condCollection.Insert(condition, condition2)
	util.CheckErr(err)
	encCollection := worker.DB().C("encounters")
	err = encCollection.Insert(encounter, encounter2)
	util.CheckErr(err)

	// Before we test delete, confirm they're really there
	count, err := condCollection.FindId("56afe6b85cdc7ec329dfe6a1").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = condCollection.FindId("56afe6b85cdc7ec329dfe6a2").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6a3").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6a4").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)

	// Now load the bundle with the delete entries and post it
	// Note that it only deletes three of the above resources and it
	// attempts a delete on a non-existent resource (which should succeed)
	data, err := os.Open("../fixtures/delete_entries_bundle.json")
	util.CheckErr(err)
	defer data.Close()

	res, err := http.Post(s.Server.URL+"/", "application/json", data)
	util.CheckErr(err)

	// Successful bundle processing should return a 200
	c.Assert(res.StatusCode, Equals, 200)

	decoder := json.NewDecoder(res.Body)
	responseBundle := &models.Bundle{}
	err = decoder.Decode(responseBundle)
	util.CheckErr(err)

	c.Assert(responseBundle.Type, Equals, "transaction-response")
	c.Assert(*responseBundle.Total, Equals, uint32(4))
	c.Assert(responseBundle.Entry, HasLen, 4)

	for _, entry := range responseBundle.Entry {
		// Everything but the Response should be nil
		c.Assert(entry.Resource, IsNil)
		c.Assert(entry.FullUrl, Equals, "")
		c.Assert(entry.Request, IsNil)
		c.Assert(entry.Search, IsNil)
		c.Assert(len(entry.Link), Equals, 0)

		// response should have 204 status
		c.Assert(entry.Response, NotNil)
		c.Assert(entry.Response.Status, Equals, "204")

		// Everything else in the response should be nil / zero value
		c.Assert(entry.Response.LastModified, IsNil)
		c.Assert(entry.Response.Location, Equals, "")
		c.Assert(entry.Response.Etag, Equals, "") // Since we don't support versioning
	}

	// Now check that the first condition and both encounters were deleted (leaving the 2nd condition)
	count, err = condCollection.FindId("56afe6b85cdc7ec329dfe6a1").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 0)
	count, err = condCollection.FindId("56afe6b85cdc7ec329dfe6a2").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6a3").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 0)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6a4").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 0)
}

func (s *BatchControllerSuite) TestConditionalDeleteEntriesBundle(c *C) {

	worker := s.MasterSession.GetWorkerSession()
	defer worker.Close()

	// Put some records in the database to delete
	encounter := &models.Encounter{
		Status: "finished",
	}
	encounter.Id = "56afe6b85cdc7ec329dfe6b1"
	encounter2 := &models.Encounter{
		Status: "planned",
	}
	encounter2.Id = "56afe6b85cdc7ec329dfe6b2"
	encounter3 := &models.Encounter{
		Status: "finished",
	}
	encounter3.Id = "56afe6b85cdc7ec329dfe6b3"
	encounter4 := &models.Encounter{
		Status: "planned",
	}
	encounter4.Id = "56afe6b85cdc7ec329dfe6b4"

	// Insert the encounters into the db
	encCollection := worker.DB().C("encounters")
	err := encCollection.Insert(encounter, encounter2, encounter3, encounter4)
	util.CheckErr(err)

	// Before we test delete, confirm they're really there
	count, err := encCollection.FindId("56afe6b85cdc7ec329dfe6b1").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6b2").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6b3").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6b4").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)

	// Now create a simple bundle with conditional delete of planned encounters
	batch := &models.Bundle{
		Type: "transaction",
		Entry: []models.BundleEntryComponent{
			{
				Request: &models.BundleEntryRequestComponent{
					Method: "DELETE",
					Url:    "Encounter?status=planned",
				},
			},
		},
	}

	data, err := json.Marshal(batch)
	util.CheckErr(err)

	res, err := http.Post(s.Server.URL+"/", "application/json", bytes.NewBuffer(data))
	util.CheckErr(err)

	// Successful bundle processing should return a 200
	c.Assert(res.StatusCode, Equals, 200)

	decoder := json.NewDecoder(res.Body)
	responseBundle := &models.Bundle{}
	err = decoder.Decode(responseBundle)
	util.CheckErr(err)

	c.Assert(responseBundle.Type, Equals, "transaction-response")
	c.Assert(*responseBundle.Total, Equals, uint32(1))
	c.Assert(responseBundle.Entry, HasLen, 1)

	entry := responseBundle.Entry[0]
	// Everything but the Response should be nil
	c.Assert(entry.Resource, IsNil)
	c.Assert(entry.FullUrl, Equals, "")
	c.Assert(entry.Request, IsNil)
	c.Assert(entry.Search, IsNil)
	c.Assert(entry.Link, HasLen, 0)

	// response should have 204 status
	c.Assert(entry.Response, NotNil)
	c.Assert(entry.Response.Status, Equals, "204")

	// Everything else in the response should be nil / zero value
	c.Assert(entry.Response.LastModified, IsNil)
	c.Assert(entry.Response.Location, Equals, "")
	c.Assert(entry.Response.Etag, Equals, "") // Since we don't support versioning

	// Now check that the right encounters were deleted
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6b1").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6b2").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 0)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6b3").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6b4").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 0)
}

func (s *BatchControllerSuite) TestPostPatientBundle(c *C) {

	worker := s.MasterSession.GetWorkerSession()
	defer worker.Close()

	data, err := os.Open("../fixtures/john_peters_bundle.json")
	util.CheckErr(err)
	defer data.Close()

	decoder := json.NewDecoder(data)
	requestBundle := &models.Bundle{}
	err = decoder.Decode(requestBundle)
	util.CheckErr(err)

	data.Seek(0, 0) // Reset the file pointer
	res, err := http.Post(s.Server.URL+"/", "application/json", data)
	util.CheckErr(err)

	c.Assert(res.StatusCode, Equals, 200)

	decoder = json.NewDecoder(res.Body)
	responseBundle := &models.Bundle{}
	err = decoder.Decode(responseBundle)
	util.CheckErr(err)

	c.Assert(responseBundle.Type, Equals, "transaction-response")
	c.Assert(*responseBundle.Total, Equals, uint32(19))
	c.Assert(responseBundle.Entry, HasLen, 19)

	for i := range responseBundle.Entry {
		resEntry, reqEntry := responseBundle.Entry[i], requestBundle.Entry[i]

		// response resource type should match request resource type
		c.Assert(reflect.TypeOf(resEntry.Resource), Equals, reflect.TypeOf(reqEntry.Resource))

		// full URLs and IDs should be difference in the response
		c.Assert(resEntry.FullUrl, Not(Equals), reqEntry.FullUrl)
		c.Assert(s.getResourceID(resEntry), Not(Equals), s.getResourceID(reqEntry))

		// full URL in response should contain the new ID
		c.Assert(strings.HasSuffix(resEntry.FullUrl, s.getResourceID(resEntry)), Equals, true)

		// resource should have lastUpdatedTime
		m := reflect.ValueOf(resEntry.Resource).Elem().FieldByName("Meta").Interface().(*models.Meta)
		c.Assert(m, NotNil)
		c.Assert(m.LastUpdated, NotNil)
		c.Assert(m.LastUpdated.Precision, Equals, models.Precision(models.Timestamp))
		c.Assert(time.Since(m.LastUpdated.Time).Minutes() < float64(1), Equals, true)

		// response should not contain the request
		c.Assert(resEntry.Request, IsNil)

		// response should have 201 status and location
		c.Assert(resEntry.Response.Status, Equals, "201")
		c.Assert(resEntry.Response.Location, Equals, resEntry.FullUrl)

		// make sure it was stored to the DB
		rName := reflect.TypeOf(resEntry.Resource).Elem().Name()
		coll := worker.DB().C(models.PluralizeLowerResourceName(rName))
		num, err := coll.Find(bson.M{"_id": s.getResourceID(resEntry)}).Count()
		util.CheckErr(err)
		c.Assert(num, Equals, 1)
	}

	// Check patient references
	patientID := responseBundle.Entry[0].Resource.(*models.Patient).Id
	c.Assert(bson.IsObjectIdHex(patientID), Equals, true)
	s.checkReference(c, responseBundle.Entry[1].Resource.(*models.Encounter).Patient, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[2].Resource.(*models.Encounter).Patient, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[3].Resource.(*models.Encounter).Patient, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[4].Resource.(*models.Encounter).Patient, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[5].Resource.(*models.Condition).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[6].Resource.(*models.Condition).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[7].Resource.(*models.Condition).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[8].Resource.(*models.Condition).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[9].Resource.(*models.Condition).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[10].Resource.(*models.Observation).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[11].Resource.(*models.Procedure).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[12].Resource.(*models.DiagnosticReport).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[13].Resource.(*models.Observation).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[14].Resource.(*models.Observation).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[15].Resource.(*models.Observation).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[16].Resource.(*models.Procedure).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[17].Resource.(*models.MedicationStatement).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[18].Resource.(*models.Immunization).Patient, patientID, "Patient")

	// Check encounter references
	encounterID := responseBundle.Entry[1].Resource.(*models.Encounter).Id
	c.Assert(bson.IsObjectIdHex(encounterID), Equals, true)
	s.checkReference(c, responseBundle.Entry[10].Resource.(*models.Observation).Encounter, encounterID, "Encounter")
	s.checkReference(c, responseBundle.Entry[11].Resource.(*models.Procedure).Encounter, encounterID, "Encounter")

	// Check dx report references
	dxReportID := responseBundle.Entry[12].Resource.(*models.DiagnosticReport).Id
	c.Assert(bson.IsObjectIdHex(dxReportID), Equals, true)
	s.checkReference(c, &responseBundle.Entry[11].Resource.(*models.Procedure).Report[0], dxReportID, "DiagnosticReport")

	// Check observation references
	obs0Id := responseBundle.Entry[13].Resource.(*models.Observation).Id
	c.Assert(bson.IsObjectIdHex(obs0Id), Equals, true)
	s.checkReference(c, &responseBundle.Entry[12].Resource.(*models.DiagnosticReport).Result[0], obs0Id, "Observation")
	obs1Id := responseBundle.Entry[14].Resource.(*models.Observation).Id
	c.Assert(bson.IsObjectIdHex(obs1Id), Equals, true)
	s.checkReference(c, &responseBundle.Entry[12].Resource.(*models.DiagnosticReport).Result[1], obs1Id, "Observation")
	obs2Id := responseBundle.Entry[15].Resource.(*models.Observation).Id
	c.Assert(bson.IsObjectIdHex(obs2Id), Equals, true)
	s.checkReference(c, &responseBundle.Entry[12].Resource.(*models.DiagnosticReport).Result[2], obs2Id, "Observation")
}

func (s *BatchControllerSuite) TestPutEntriesBundle(c *C) {

	worker := s.MasterSession.GetWorkerSession()
	defer worker.Close()

	// Put some records in the database to update
	patient := &models.Patient{
		Identifier: []models.Identifier{
			{System: "http://test.org/simple", Value: "doejohn"},
		},
		Name: []models.HumanName{
			{Given: []string{"John"}, Family: "Doe"},
		},
	}
	patient.Id = "56afe6b85cdc7ec329dfe6a0"
	condition := &models.Condition{
		Subject: &models.Reference{
			Type:         "Patient",
			Reference:    s.Server.URL + "/Patient/56afe6b85cdc7ec329dfe6a0",
			ReferencedID: "56afe6b85cdc7ec329dfe6a0",
			External:     new(bool),
		},
		Code: &models.CodeableConcept{
			Coding: []models.Coding{
				{System: "Foo", Code: "Bar"},
			},
		},
		VerificationStatus: "confirmed",
	}
	condition.Id = "56afe6b85cdc7ec329dfe6a1"
	condition2 := &models.Condition{
		Subject: &models.Reference{
			Type:         "Patient",
			Reference:    s.Server.URL + "/Patient/56afe6b85cdc7ec329dfe6a0",
			ReferencedID: "56afe6b85cdc7ec329dfe6a0",
			External:     new(bool),
		},
		Code: &models.CodeableConcept{
			Coding: []models.Coding{
				{System: "Foo", Code: "Baz"},
			},
		},
		VerificationStatus: "confirmed",
	}
	condition2.Id = "56afe6b85cdc7ec329dfe6a2"

	// Insert the conditions into the db
	patCollection := worker.DB().C("patients")
	err := patCollection.Insert(patient)
	util.CheckErr(err)
	condCollection := worker.DB().C("conditions")
	err = condCollection.Insert(condition, condition2)
	util.CheckErr(err)

	// Now load the bundle with the put entries and post it.
	data, err := os.Open("../fixtures/put_entries_bundle.json")
	util.CheckErr(err)
	defer data.Close()

	res, err := http.Post(s.Server.URL+"/", "application/json", data)
	util.CheckErr(err)

	// Successful bundle processing should return a 200
	c.Assert(res.StatusCode, Equals, 200)

	decoder := json.NewDecoder(res.Body)
	responseBundle := &models.Bundle{}
	err = decoder.Decode(responseBundle)
	util.CheckErr(err)

	c.Assert(responseBundle.Type, Equals, "transaction-response")
	c.Assert(*responseBundle.Total, Equals, uint32(4))
	c.Assert(responseBundle.Entry, HasLen, 4)

	patEntry := responseBundle.Entry[0]

	// response resource type should match request resource type
	c.Assert(patEntry.Resource, FitsTypeOf, &models.Patient{})

	// full URLs and IDs should contain correct ID in response
	c.Assert(patEntry.FullUrl, Equals, s.Server.URL+"/Patient/56afe6b85cdc7ec329dfe6a0")
	c.Assert(s.getResourceID(patEntry), Equals, "56afe6b85cdc7ec329dfe6a0")

	// resource should have lastUpdatedTime
	m := reflect.ValueOf(patEntry.Resource).Elem().FieldByName("Meta").Interface().(*models.Meta)
	c.Assert(m, NotNil)
	c.Assert(m.LastUpdated, NotNil)
	c.Assert(m.LastUpdated.Precision, Equals, models.Precision(models.Timestamp))
	c.Assert(time.Since(m.LastUpdated.Time).Minutes() < float64(1), Equals, true)

	// response should not contain the request
	c.Assert(patEntry.Request, IsNil)

	// response should have 200 status and location
	c.Assert(patEntry.Response.Status, Equals, "200")
	c.Assert(patEntry.Response.Location, Equals, patEntry.FullUrl)

	// Now check other entries
	expectedIDs := []string{"56afe6b85cdc7ec329dfe6a2", "56afe6b85cdc7ec329dfe6a3", "56afe6b85cdc7ec329dfe6a1"}
	for i := 1; i < len(responseBundle.Entry); i++ {
		resEntry := responseBundle.Entry[i]

		// response resource type should be a condition
		c.Assert(resEntry.Resource, FitsTypeOf, &models.Condition{})

		// Reference to patient should be to upserted patient
		s.checkReference(c, resEntry.Resource.(*models.Condition).Subject, "56afe6b85cdc7ec329dfe6a0", "Patient")

		// check full URL and ID match expected values
		c.Assert(resEntry.FullUrl, Equals, s.Server.URL+"/Condition/"+expectedIDs[i-1])
		c.Assert(s.getResourceID(resEntry), Equals, expectedIDs[i-1])

		// resource should have lastUpdatedTime
		m := reflect.ValueOf(resEntry.Resource).Elem().FieldByName("Meta").Interface().(*models.Meta)
		c.Assert(m, NotNil)
		c.Assert(m.LastUpdated, NotNil)
		c.Assert(m.LastUpdated.Precision, Equals, models.Precision(models.Timestamp))
		c.Assert(time.Since(m.LastUpdated.Time).Minutes() < float64(1), Equals, true)

		// response should not contain the request
		c.Assert(resEntry.Request, IsNil)

		// response should have 200 or 201 status and location
		switch i {
		case 1, 3:
			c.Assert(resEntry.Response.Status, Equals, "200")
		case 2:
			c.Assert(resEntry.Response.Status, Equals, "201")
		}
		c.Assert(resEntry.Response.Location, Equals, resEntry.FullUrl)
	}

	// Now do a quick content check
	count, err := condCollection.Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 3)

	pat1 := models.Patient{}
	err = patCollection.FindId("56afe6b85cdc7ec329dfe6a0").One(&pat1)
	util.CheckErr(err)
	c.Assert(pat1.Gender, Equals, "male")

	cond1 := models.Condition{}
	err = condCollection.FindId("56afe6b85cdc7ec329dfe6a1").One(&cond1)
	util.CheckErr(err)
	c.Assert(cond1.Code.Coding, HasLen, 1)
	c.Assert(cond1.Code.Coding[0].Code, Equals, "Bar2")

	cond2 := models.Condition{}
	err = condCollection.FindId("56afe6b85cdc7ec329dfe6a2").One(&cond2)
	util.CheckErr(err)
	c.Assert(cond2.Code.Coding, HasLen, 1)
	c.Assert(cond2.Code.Coding[0].Code, Equals, "Baz2")

	cond3 := models.Condition{}
	err = condCollection.FindId("56afe6b85cdc7ec329dfe6a3").One(&cond3)
	util.CheckErr(err)
	c.Assert(cond3.Code.Coding, HasLen, 1)
	c.Assert(cond3.Code.Coding[0].Code, Equals, "Bat")
}

func (s *BatchControllerSuite) TestConditionalUpdatesBundle(c *C) {

	worker := s.MasterSession.GetWorkerSession()
	defer worker.Close()

	data, err := os.Open("../fixtures/conditional_update_bundle.json")
	util.CheckErr(err)
	defer data.Close()

	decoder := json.NewDecoder(data)
	requestBundle := &models.Bundle{}
	err = decoder.Decode(requestBundle)
	util.CheckErr(err)

	// Do the initial post
	data.Seek(0, 0) // Reset the file pointer
	res, err := http.Post(s.Server.URL+"/", "application/json", data)
	util.CheckErr(err)

	c.Assert(res.StatusCode, Equals, 200)

	decoder = json.NewDecoder(res.Body)
	responseBundle := &models.Bundle{}
	err = decoder.Decode(responseBundle)
	util.CheckErr(err)

	c.Assert(responseBundle.Type, Equals, "transaction-response")
	c.Assert(*responseBundle.Total, Equals, uint32(20))
	c.Assert(responseBundle.Entry, HasLen, 20)

	// Check all of the response status for created (vs updated)
	for _, entry := range responseBundle.Entry {
		c.Assert(entry.Response.Status, Equals, "201")
	}

	// Do it again!
	data, err = os.Open("../fixtures/conditional_update_bundle.json")
	util.CheckErr(err)
	defer data.Close()

	res, err = http.Post(s.Server.URL+"/", "application/json", data)
	util.CheckErr(err)

	c.Assert(res.StatusCode, Equals, 200)

	decoder = json.NewDecoder(res.Body)
	response2Bundle := &models.Bundle{}
	err = decoder.Decode(response2Bundle)
	util.CheckErr(err)

	c.Assert(response2Bundle.Type, Equals, "transaction-response")
	c.Assert(*response2Bundle.Total, Equals, uint32(20))
	c.Assert(response2Bundle.Entry, HasLen, 20)

	// Now check all of the response status for updated (vs created)
	for _, entry := range response2Bundle.Entry {
		c.Assert(entry.Response.Status, Equals, "200")
	}
}

func (s *BatchControllerSuite) TestAllSupportedMethodsBundle(c *C) {

	worker := s.MasterSession.GetWorkerSession()
	defer worker.Close()

	// Create some records to delete or update
	condition := &models.Condition{
		Subject: &models.Reference{Reference: "https://example.com/base/Patient/4954037112938410473"},
		Code: &models.CodeableConcept{
			Coding: []models.Coding{
				{System: "Foo", Code: "Bar"},
			},
		},
		VerificationStatus: "confirmed",
	}
	condition.Id = "56afe6b85cdc7ec329dfe6a5"
	encounter := &models.Encounter{
		Status: "finished",
	}
	encounter.Id = "56afe6b85cdc7ec329dfe6a6"
	encounter2 := &models.Encounter{
		Status: "planned",
	}
	encounter2.Id = "56afe6b85cdc7ec329dfe6a7"

	// Put those records in the db to delete or update
	encCollection := worker.DB().C("encounters")
	err := encCollection.Insert(encounter, encounter2)
	util.CheckErr(err)
	condCollection := worker.DB().C("conditions")
	err = condCollection.Insert(condition)
	util.CheckErr(err)

	// Before we test delete, confirm they're really there
	count, err := condCollection.FindId("56afe6b85cdc7ec329dfe6a5").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6a6").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6a7").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)

	// Load the bundle with delete / post /put entries and post it
	data, err := os.Open("../fixtures/all_supported_methods_bundle.json")
	util.CheckErr(err)
	defer data.Close()

	decoder := json.NewDecoder(data)
	requestBundle := &models.Bundle{}
	err = decoder.Decode(requestBundle)
	util.CheckErr(err)

	data.Seek(0, 0) // Reset the file pointer
	res, err := http.Post(s.Server.URL+"/", "application/json", data)
	util.CheckErr(err)

	c.Assert(res.StatusCode, Equals, 200)

	decoder = json.NewDecoder(res.Body)
	responseBundle := &models.Bundle{}
	err = decoder.Decode(responseBundle)
	util.CheckErr(err)

	c.Assert(responseBundle.Type, Equals, "transaction-response")
	c.Assert(*responseBundle.Total, Equals, uint32(7))
	c.Assert(responseBundle.Entry, HasLen, 7)

	// First check the DELETEd resources (first two entries)
	for i := 0; i < 2; i++ {
		entry := responseBundle.Entry[i]

		// Everything but the Response should be nil
		c.Assert(entry.Resource, IsNil)
		c.Assert(entry.FullUrl, Equals, "")
		c.Assert(entry.Request, IsNil)
		c.Assert(entry.Search, IsNil)
		c.Assert(len(entry.Link), Equals, 0)

		// response should have 204 status
		c.Assert(entry.Response, NotNil)
		c.Assert(entry.Response.Status, Equals, "204")

		// Everything else in the response should be nil / zero value
		c.Assert(entry.Response.LastModified, IsNil)
		c.Assert(entry.Response.Location, Equals, "")
		c.Assert(entry.Response.Etag, Equals, "") // Since we don't support versioning
	}

	count, err = condCollection.FindId("56afe6b85cdc7ec329dfe6a5").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 0)
	count, err = condCollection.FindId("56afe6b85cdc7ec329dfe6a6").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 0)

	// Then check the POSTed resources
	for i := 2; i < 5; i++ {
		resEntry, reqEntry := responseBundle.Entry[i], requestBundle.Entry[i]

		// response resource type should match request resource type
		c.Assert(reflect.TypeOf(resEntry.Resource), Equals, reflect.TypeOf(reqEntry.Resource))

		// full URLs and IDs should be difference in the response
		c.Assert(resEntry.FullUrl, Not(Equals), reqEntry.FullUrl)
		c.Assert(s.getResourceID(resEntry), Not(Equals), s.getResourceID(reqEntry))

		// full URL in response should contain the new ID
		c.Assert(strings.HasSuffix(resEntry.FullUrl, s.getResourceID(resEntry)), Equals, true)

		// resource should have lastUpdatedTime
		m := reflect.ValueOf(resEntry.Resource).Elem().FieldByName("Meta").Interface().(*models.Meta)
		c.Assert(m, NotNil)
		c.Assert(m.LastUpdated, NotNil)
		c.Assert(m.LastUpdated.Precision, Equals, models.Precision(models.Timestamp))
		c.Assert(time.Since(m.LastUpdated.Time).Minutes() < float64(1), Equals, true)

		// response should not contain the request
		c.Assert(resEntry.Request, IsNil)

		// response should have 201 status and location
		c.Assert(resEntry.Response.Status, Equals, "201")
		c.Assert(resEntry.Response.Location, Equals, resEntry.FullUrl)

		// make sure it was stored to the DB
		rName := reflect.TypeOf(resEntry.Resource).Elem().Name()
		coll := worker.DB().C(models.PluralizeLowerResourceName(rName))
		num, err := coll.Find(bson.M{"_id": s.getResourceID(resEntry)}).Count()
		util.CheckErr(err)
		c.Assert(num, Equals, 1)
	}

	// Then check the PUTted resources
	expectedIDs := []string{"56afe6b85cdc7ec329dfe6a7", "56afe6b85cdc7ec329dfe6a8"}
	for i := 5; i < 7; i++ {
		resEntry := responseBundle.Entry[i]

		// response resource type should be an encounter
		c.Assert(resEntry.Resource, FitsTypeOf, &models.Encounter{})

		// check full URL and ID match expected values
		c.Assert(resEntry.FullUrl, Equals, s.Server.URL+"/Encounter/"+expectedIDs[i-5])
		c.Assert(s.getResourceID(resEntry), Equals, expectedIDs[i-5])

		// resource should have lastUpdatedTime
		m := reflect.ValueOf(resEntry.Resource).Elem().FieldByName("Meta").Interface().(*models.Meta)
		c.Assert(m, NotNil)
		c.Assert(m.LastUpdated, NotNil)
		c.Assert(m.LastUpdated.Precision, Equals, models.Precision(models.Timestamp))
		c.Assert(time.Since(m.LastUpdated.Time).Minutes() < float64(1), Equals, true)

		// response should not contain the request
		c.Assert(resEntry.Request, IsNil)

		// response should have 200 or 201 status and location
		c.Assert(resEntry.Response.Status, Equals, fmt.Sprint(195+i)) // this just happens to work out (200, 201)
		c.Assert(resEntry.Response.Location, Equals, resEntry.FullUrl)
	}

	// Quick content check on the PUTs
	enc1 := models.Encounter{}
	err = encCollection.FindId("56afe6b85cdc7ec329dfe6a7").One(&enc1)
	util.CheckErr(err)
	c.Assert(enc1.Status, Equals, "finished")
	c.Assert(enc1.Period.Start.Time.Equal(time.Date(2011, 12, 1, 13, 0, 0, 0, time.UTC)), Equals, true)
	c.Assert(enc1.Period.End.Time.Equal(time.Date(2011, 12, 1, 14, 0, 0, 0, time.UTC)), Equals, true)

	enc2 := models.Encounter{}
	err = encCollection.FindId("56afe6b85cdc7ec329dfe6a8").One(&enc2)
	util.CheckErr(err)
	c.Assert(enc2.Status, Equals, "planned")
	c.Assert(enc2.Period, IsNil)

	// Check patient references
	patientID := responseBundle.Entry[2].Resource.(*models.Patient).Id
	c.Assert(bson.IsObjectIdHex(patientID), Equals, true)
	s.checkReference(c, responseBundle.Entry[3].Resource.(*models.Encounter).Patient, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[4].Resource.(*models.Condition).Subject, patientID, "Patient")
}

func (s *BatchControllerSuite) checkReference(c *C, ref *models.Reference, id string, typ string) {
	c.Assert(ref.ReferencedID, Equals, id)
	c.Assert(ref.Type, Equals, typ)
	c.Assert(ref.Reference, Equals, typ+"/"+id)
	c.Assert(*ref.External, Equals, false)
}

func (s *BatchControllerSuite) getResourceID(e models.BundleEntryComponent) string {
	return reflect.ValueOf(e.Resource).Elem().FieldByName("Id").String()
}
