package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/intervention-engine/fhir/models"
	"github.com/labstack/echo"
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BatchControllerSuite struct {
	Session *mgo.Session
	Echo    *echo.Echo
	Server  *httptest.Server
}

var _ = Suite(&BatchControllerSuite{})

func (s *BatchControllerSuite) SetUpSuite(c *C) {

	// Set up the database
	var err error
	s.Session, err = mgo.Dial("localhost")
	util.CheckErr(err)
	Database = s.Session.DB("fhir-test")

	// Build routes for testing
	s.Echo = echo.New()
	RegisterRoutes(s.Echo, make(map[string][]echo.Middleware))

	// Create httptest server
	s.Server = httptest.NewServer(s.Echo.Router())
}

func (s *BatchControllerSuite) TearDownSuite(c *C) {
	Database.DropDatabase()
	s.Session.Close()
	s.Server.Close()
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
		coll := Database.C(models.PluralizeLowerResourceName(rName))
		num, err := coll.Find(bson.M{"_id": s.getResourceID(resEntry)}).Count()
		util.CheckErr(err)
		c.Assert(num, Equals, 1)
	}

	// Check patient references
	patientId := responseBundle.Entry[0].Resource.(*models.Patient).Id
	c.Assert(bson.IsObjectIdHex(patientId), Equals, true)
	s.checkReference(c, responseBundle.Entry[1].Resource.(*models.Encounter).Patient, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[2].Resource.(*models.Encounter).Patient, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[3].Resource.(*models.Encounter).Patient, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[4].Resource.(*models.Encounter).Patient, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[5].Resource.(*models.Condition).Patient, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[6].Resource.(*models.Condition).Patient, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[7].Resource.(*models.Condition).Patient, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[8].Resource.(*models.Condition).Patient, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[9].Resource.(*models.Condition).Patient, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[10].Resource.(*models.Observation).Subject, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[11].Resource.(*models.Procedure).Subject, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[12].Resource.(*models.DiagnosticReport).Subject, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[13].Resource.(*models.Observation).Subject, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[14].Resource.(*models.Observation).Subject, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[15].Resource.(*models.Observation).Subject, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[16].Resource.(*models.Procedure).Subject, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[17].Resource.(*models.MedicationStatement).Patient, patientId, "Patient")
	s.checkReference(c, responseBundle.Entry[18].Resource.(*models.Immunization).Patient, patientId, "Patient")

	// Check encounter references
	encounterId := responseBundle.Entry[1].Resource.(*models.Encounter).Id
	c.Assert(bson.IsObjectIdHex(encounterId), Equals, true)
	s.checkReference(c, responseBundle.Entry[10].Resource.(*models.Observation).Encounter, encounterId, "Encounter")
	s.checkReference(c, responseBundle.Entry[11].Resource.(*models.Procedure).Encounter, encounterId, "Encounter")

	// Check dx report references
	dxReportId := responseBundle.Entry[12].Resource.(*models.DiagnosticReport).Id
	c.Assert(bson.IsObjectIdHex(dxReportId), Equals, true)
	s.checkReference(c, &responseBundle.Entry[11].Resource.(*models.Procedure).Report[0], dxReportId, "DiagnosticReport")

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

func (s *BatchControllerSuite) checkReference(c *C, ref *models.Reference, id string, typ string) {
	c.Assert(ref.ReferencedID, Equals, id)
	c.Assert(ref.Type, Equals, typ)
	c.Assert(ref.Reference, Equals, typ+"/"+id)
	c.Assert(*ref.External, Equals, false)
}

func (s *BatchControllerSuite) getResourceID(e models.BundleEntryComponent) string {
	return reflect.ValueOf(e.Resource).Elem().FieldByName("Id").String()
}
