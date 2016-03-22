package server

import (
	"bytes"
	"encoding/json"
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
	Database *mgo.Database
	Session  *mgo.Session
	Engine   *gin.Engine
	Server   *httptest.Server
}

var _ = Suite(&BatchControllerSuite{})

func (s *BatchControllerSuite) SetUpSuite(c *C) {

	// Set up the database
	var err error
	s.Session, err = mgo.Dial("localhost")
	util.CheckErr(err)
	s.Database = s.Session.DB("fhir-test")

	// Build routes for testing
	s.Engine = gin.New()
	RegisterRoutes(s.Engine, make(map[string][]gin.HandlerFunc), NewMongoDataAccessLayer(s.Database), Config{})

	// Create httptest server
	s.Server = httptest.NewServer(s.Engine)
}

func (s *BatchControllerSuite) TearDownSuite(c *C) {
	s.Database.DropDatabase()
	s.Session.Close()
	s.Server.Close()
}

func (s *BatchControllerSuite) TestDeleteEntriesBundle(c *C) {
	// Put some records in the database to delete
	condition := &models.Condition{
		Patient: &models.Reference{Reference: "https://example.com/base/Patient/4954037112938410473"},
		Code: &models.CodeableConcept{
			Coding: []models.Coding{
				{System: "Foo", Code: "Bar"},
			},
		},
		VerificationStatus: "confirmed",
	}
	condition.Id = "56afe6b85cdc7ec329dfe6a1"
	condition2 := &models.Condition{
		Patient: &models.Reference{Reference: "https://example.com/base/Patient/4954037112938410473"},
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
	condCollection := s.Database.C("conditions")
	err := condCollection.Insert(condition, condition2)
	util.CheckErr(err)
	encCollection := s.Database.C("encounters")
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
	encCollection := s.Database.C("encounters")
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

func (s *BatchControllerSuite) TestUploadPatientBundle(c *C) {
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
		since := time.Since(m.LastUpdated.Time)
		c.Assert(since.Hours() < float64(1), Equals, true)
		c.Assert(since.Minutes() < float64(1), Equals, true)

		// response should not contain the request
		c.Assert(resEntry.Request, IsNil)

		// response should have 201 status and location
		c.Assert(resEntry.Response.Status, Equals, "201")
		c.Assert(resEntry.Response.Location, Equals, resEntry.FullUrl)

		// make sure it was stored to the DB
		rName := reflect.TypeOf(resEntry.Resource).Elem().Name()
		coll := s.Database.C(models.PluralizeLowerResourceName(rName))
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
	s.checkReference(c, responseBundle.Entry[5].Resource.(*models.Condition).Patient, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[6].Resource.(*models.Condition).Patient, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[7].Resource.(*models.Condition).Patient, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[8].Resource.(*models.Condition).Patient, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[9].Resource.(*models.Condition).Patient, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[10].Resource.(*models.Observation).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[11].Resource.(*models.Procedure).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[12].Resource.(*models.DiagnosticReport).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[13].Resource.(*models.Observation).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[14].Resource.(*models.Observation).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[15].Resource.(*models.Observation).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[16].Resource.(*models.Procedure).Subject, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[17].Resource.(*models.MedicationStatement).Patient, patientID, "Patient")
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

func (s *BatchControllerSuite) TestAllSupportedMethodsBundle(c *C) {
	// Put some records in the database to delete
	condition := &models.Condition{
		Patient: &models.Reference{Reference: "https://example.com/base/Patient/4954037112938410473"},
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

	// Put those records in the db to delete
	encCollection := s.Database.C("encounters")
	err := encCollection.Insert(encounter)
	util.CheckErr(err)
	condCollection := s.Database.C("conditions")
	err = condCollection.Insert(condition)
	util.CheckErr(err)

	// Before we test delete, confirm they're really there
	count, err := condCollection.FindId("56afe6b85cdc7ec329dfe6a5").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)
	count, err = encCollection.FindId("56afe6b85cdc7ec329dfe6a6").Count()
	util.CheckErr(err)
	c.Assert(count, Equals, 1)

	// Load the bundle with delete / post entries and post it
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
	c.Assert(*responseBundle.Total, Equals, uint32(5))
	c.Assert(responseBundle.Entry, HasLen, 5)

	// First check the deleted resources (first two entries)
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
		since := time.Since(m.LastUpdated.Time)
		c.Assert(since.Hours() < float64(1), Equals, true)
		c.Assert(since.Minutes() < float64(1), Equals, true)

		// response should not contain the request
		c.Assert(resEntry.Request, IsNil)

		// response should have 201 status and location
		c.Assert(resEntry.Response.Status, Equals, "201")
		c.Assert(resEntry.Response.Location, Equals, resEntry.FullUrl)

		// make sure it was stored to the DB
		rName := reflect.TypeOf(resEntry.Resource).Elem().Name()
		coll := s.Database.C(models.PluralizeLowerResourceName(rName))
		num, err := coll.Find(bson.M{"_id": s.getResourceID(resEntry)}).Count()
		util.CheckErr(err)
		c.Assert(num, Equals, 1)
	}

	// Check patient references
	patientID := responseBundle.Entry[2].Resource.(*models.Patient).Id
	c.Assert(bson.IsObjectIdHex(patientID), Equals, true)
	s.checkReference(c, responseBundle.Entry[3].Resource.(*models.Encounter).Patient, patientID, "Patient")
	s.checkReference(c, responseBundle.Entry[4].Resource.(*models.Condition).Patient, patientID, "Patient")
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
