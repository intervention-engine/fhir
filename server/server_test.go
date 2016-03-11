package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/fhir/search"
	"github.com/labstack/echo"
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ServerSuite struct {
	Database  *mgo.Database
	Session   *mgo.Session
	Echo      *echo.Echo
	Server    *httptest.Server
	FixtureId string
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&ServerSuite{})

func (s *ServerSuite) SetUpSuite(c *C) {

	// Set up the database
	var err error
	s.Session, err = mgo.Dial("localhost")
	util.CheckErr(err)
	s.Database = s.Session.DB("fhir-test")

	// Build routes for testing
	s.Echo = echo.New()
	RegisterRoutes(s.Echo, make(map[string][]echo.Middleware), NewMongoDataAccessLayer(s.Database), Config{})

	// Create httptest server
	s.Server = httptest.NewServer(s.Echo.Router())
}

func (s *ServerSuite) SetUpTest(c *C) {
	// Add patient fixture
	p := s.insertPatientFromFixture("../fixtures/patient-example-a.json")
	s.FixtureId = p.Id
}

func (s *ServerSuite) TearDownTest(c *C) {
	s.Database.C("patients").DropCollection()
}

func (s *ServerSuite) TearDownSuite(c *C) {
	s.Database.DropDatabase()
	s.Session.Close()
	s.Server.Close()
}

func (s *ServerSuite) TestGetPatients(c *C) {
	// Add 4 more patients
	for i := 0; i < 4; i++ {
		s.insertPatientFromFixture("../fixtures/patient-example-a.json")
	}
	assertBundleCount(c, s.Server.URL+"/Patient", 5, 5)
}

func (s *ServerSuite) TestGetPatientsWithOptions(c *C) {
	// Add 4 more patients
	for i := 0; i < 4; i++ {
		s.insertPatientFromFixture("../fixtures/patient-example-a.json")
	}
	assertBundleCount(c, s.Server.URL+"/Patient?_count=2", 2, 5)
	assertBundleCount(c, s.Server.URL+"/Patient?_offset=2", 3, 5)
	assertBundleCount(c, s.Server.URL+"/Patient?_count=2&_offset=1", 2, 5)
	assertBundleCount(c, s.Server.URL+"/Patient?_count=2&_offset=4", 1, 5)
	assertBundleCount(c, s.Server.URL+"/Patient?_offset=100", 0, 5)
}

func (s *ServerSuite) TestGetPatientsDefaultLimitIs100(c *C) {
	// Add 100 more patients
	for i := 0; i < 100; i++ {
		s.insertPatientFromFixture("../fixtures/patient-example-a.json")
	}
	assertBundleCount(c, s.Server.URL+"/Patient", 100, 101)
}

func (s *ServerSuite) TestGetPatientsPaging(c *C) {
	// Add 39 more patients
	for i := 0; i < 39; i++ {
		s.insertPatientFromFixture("../fixtures/patient-example-a.json")
	}

	// Default counts and less results than count
	bundle := performSearch(c, s.Server.URL+"/Patient")
	c.Assert(bundle.Link, HasLen, 3)
	assertPagingLink(c, bundle.Link[0], "self", 100, 0)
	assertPagingLink(c, bundle.Link[1], "first", 100, 0)
	assertPagingLink(c, bundle.Link[2], "last", 100, 0)

	// More results than count, first page
	bundle = performSearch(c, s.Server.URL+"/Patient?_count=10")
	c.Assert(bundle.Link, HasLen, 4)
	assertPagingLink(c, bundle.Link[0], "self", 10, 0)
	assertPagingLink(c, bundle.Link[1], "first", 10, 0)
	assertPagingLink(c, bundle.Link[2], "next", 10, 10)
	assertPagingLink(c, bundle.Link[3], "last", 10, 30)

	// More results than count, middle page
	bundle = performSearch(c, s.Server.URL+"/Patient?_count=10&_offset=20")
	c.Assert(bundle.Link, HasLen, 5)
	assertPagingLink(c, bundle.Link[0], "self", 10, 20)
	assertPagingLink(c, bundle.Link[1], "first", 10, 0)
	assertPagingLink(c, bundle.Link[2], "previous", 10, 10)
	assertPagingLink(c, bundle.Link[3], "next", 10, 30)
	assertPagingLink(c, bundle.Link[4], "last", 10, 30)

	// More results than count, last page
	bundle = performSearch(c, s.Server.URL+"/Patient?_count=10&_offset=30")
	c.Assert(bundle.Link, HasLen, 4)
	assertPagingLink(c, bundle.Link[0], "self", 10, 30)
	assertPagingLink(c, bundle.Link[1], "first", 10, 0)
	assertPagingLink(c, bundle.Link[2], "previous", 10, 20)
	assertPagingLink(c, bundle.Link[3], "last", 10, 30)

	// More results than count, uneven last page
	bundle = performSearch(c, s.Server.URL+"/Patient?_count=10&_offset=25")
	c.Assert(bundle.Link, HasLen, 5)
	assertPagingLink(c, bundle.Link[0], "self", 10, 25)
	assertPagingLink(c, bundle.Link[1], "first", 10, 0)
	assertPagingLink(c, bundle.Link[2], "previous", 10, 15)
	assertPagingLink(c, bundle.Link[3], "next", 10, 35)
	assertPagingLink(c, bundle.Link[4], "last", 10, 35)

	// More results than count, uneven previous page and last page
	bundle = performSearch(c, s.Server.URL+"/Patient?_count=10&_offset=5")
	c.Assert(bundle.Link, HasLen, 5)
	assertPagingLink(c, bundle.Link[0], "self", 10, 5)
	assertPagingLink(c, bundle.Link[1], "first", 10, 0)
	assertPagingLink(c, bundle.Link[2], "previous", 5, 0)
	assertPagingLink(c, bundle.Link[3], "next", 10, 15)
	assertPagingLink(c, bundle.Link[4], "last", 10, 35)

	// Search with other search criteria and results
	bundle = performSearch(c, s.Server.URL+"/Patient?_count=10&gender=male")
	c.Assert(bundle.Link, HasLen, 4)
	assertPagingLink(c, bundle.Link[0], "self", 10, 0)
	assertPagingLink(c, bundle.Link[1], "first", 10, 0)
	assertPagingLink(c, bundle.Link[2], "next", 10, 10)
	assertPagingLink(c, bundle.Link[3], "last", 10, 30)

	// Search with no results
	bundle = performSearch(c, s.Server.URL+"/Patient?_count=10&gender=FOO")
	c.Assert(bundle.Link, HasLen, 3)
	assertPagingLink(c, bundle.Link[0], "self", 10, 0)
	assertPagingLink(c, bundle.Link[1], "first", 10, 0)
	assertPagingLink(c, bundle.Link[2], "last", 10, 0)

	// Search with out of bounds offset
	bundle = performSearch(c, s.Server.URL+"/Patient?_count=10&_offset=1000")
	c.Assert(bundle.Link, HasLen, 4)
	assertPagingLink(c, bundle.Link[0], "self", 10, 1000)
	assertPagingLink(c, bundle.Link[1], "first", 10, 0)
	assertPagingLink(c, bundle.Link[2], "previous", 10, 990)
	assertPagingLink(c, bundle.Link[3], "last", 10, 30)

	// Search with negative offset
	bundle = performSearch(c, s.Server.URL+"/Patient?_offset=-10")
	c.Assert(bundle.Link, HasLen, 3)
	assertPagingLink(c, bundle.Link[0], "self", 100, 0)
	assertPagingLink(c, bundle.Link[1], "first", 100, 0)
	assertPagingLink(c, bundle.Link[2], "last", 100, 0)

	// Search with negative count
	bundle = performSearch(c, s.Server.URL+"/Patient?_count=-10")
	c.Assert(bundle.Link, HasLen, 3)
	assertPagingLink(c, bundle.Link[0], "self", 100, 0)
	assertPagingLink(c, bundle.Link[1], "first", 100, 0)
	assertPagingLink(c, bundle.Link[2], "last", 100, 0)
}

func (s *ServerSuite) TestGetPatientSearchPagingPreservesSearchParams(c *C) {
	// Add 39 more patients
	for i := 0; i < 39; i++ {
		s.insertPatientFromFixture("../fixtures/patient-example-a.json")
	}

	// Default counts and less results than count
	bundle := performSearch(c, s.Server.URL+"/Patient?gender=male&name=Donald&name=Duck")
	v := url.Values{}
	v.Set("gender", "male")
	v.Add("name", "Donald")
	v.Add("name", "Duck")
	c.Assert(bundle.Link, HasLen, 3)
	assertPagingLinkWithParams(c, bundle.Link[0], "self", v, 100, 0)
	assertPagingLinkWithParams(c, bundle.Link[1], "first", v, 100, 0)
	assertPagingLinkWithParams(c, bundle.Link[2], "last", v, 100, 0)

	// More results than count, first page
	bundle = performSearch(c, s.Server.URL+"/Patient?gender=male&name=Donald&name=Duck&_count=10")
	c.Assert(bundle.Link, HasLen, 4)
	assertPagingLinkWithParams(c, bundle.Link[0], "self", v, 10, 0)
	assertPagingLinkWithParams(c, bundle.Link[1], "first", v, 10, 0)
	assertPagingLinkWithParams(c, bundle.Link[2], "next", v, 10, 10)
	assertPagingLinkWithParams(c, bundle.Link[3], "last", v, 10, 30)

	// More results than count, middle page
	bundle = performSearch(c, s.Server.URL+"/Patient?gender=male&name=Donald&name=Duck&_count=10&_offset=20")
	c.Assert(bundle.Link, HasLen, 5)
	assertPagingLinkWithParams(c, bundle.Link[0], "self", v, 10, 20)
	assertPagingLinkWithParams(c, bundle.Link[1], "first", v, 10, 0)
	assertPagingLinkWithParams(c, bundle.Link[2], "previous", v, 10, 10)
	assertPagingLinkWithParams(c, bundle.Link[3], "next", v, 10, 30)
	assertPagingLinkWithParams(c, bundle.Link[4], "last", v, 10, 30)
}

func (s *ServerSuite) TestGetPatient(c *C) {
	res, err := http.Get(s.Server.URL + "/Patient/" + s.FixtureId)
	util.CheckErr(err)

	decoder := json.NewDecoder(res.Body)
	patient := &models.Patient{}
	err = decoder.Decode(patient)
	util.CheckErr(err)
	c.Assert(patient.Name[0].Family[0], Equals, "Donald")
}

func (s *ServerSuite) TestShowPatient(c *C) {
	res, err := http.Get(s.Server.URL + "/Patient")
	util.CheckErr(err)

	decoder := json.NewDecoder(res.Body)
	patientBundle := &models.Bundle{}
	err = decoder.Decode(patientBundle)
	util.CheckErr(err)

	var result []models.Patient
	collection := s.Database.C("patients")
	iter := collection.Find(nil).Iter()
	err = iter.All(&result)
	util.CheckErr(err)

	c.Assert(int(*patientBundle.Total), Equals, len(result))
}

func (s *ServerSuite) TestCreatePatient(c *C) {
	data, err := os.Open("../fixtures/patient-example-b.json")
	util.CheckErr(err)
	defer data.Close()

	res, err := http.Post(s.Server.URL+"/Patient", "application/json", data)
	util.CheckErr(err)

	splitLocation := strings.Split(res.Header["Location"][0], "/")
	createdPatientId := splitLocation[len(splitLocation)-1]

	s.checkCreatedPatient(createdPatientId, c)
}

func (s *ServerSuite) TestCreatePatientByPut(c *C) {
	data, err := os.Open("../fixtures/patient-example-b.json")
	util.CheckErr(err)
	defer data.Close()

	createdPatientId := bson.NewObjectId().Hex()

	req, err := http.NewRequest("PUT", s.Server.URL+"/Patient/"+createdPatientId, data)
	util.CheckErr(err)

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	_, err = client.Do(req)
	util.CheckErr(err)

	s.checkCreatedPatient(createdPatientId, c)
}

func (s *ServerSuite) checkCreatedPatient(createdPatientId string, c *C) {
	patientCollection := s.Database.C("patients")
	patient := models.Patient{}
	err := patientCollection.Find(bson.M{"_id": createdPatientId}).One(&patient)
	util.CheckErr(err)
	c.Assert(patient.Name[0].Family[0], Equals, "Daffy")
	c.Assert(patient.Meta, NotNil)
	c.Assert(patient.Meta.LastUpdated, NotNil)
	c.Assert(patient.Meta.LastUpdated.Precision, Equals, models.Precision(models.Timestamp))
	since := time.Since(patient.Meta.LastUpdated.Time)
	c.Assert(since.Hours() < float64(1), Equals, true)
	c.Assert(since.Minutes() < float64(1), Equals, true)
}

func (s *ServerSuite) TestGetConditionsWithIncludes(c *C) {
	// Add 1 more patient
	patient := s.insertPatientFromFixture("../fixtures/patient-example-a.json")

	// Add condition
	data, err := os.Open("../fixtures/condition.json")
	util.CheckErr(err)
	defer data.Close()
	decoder := json.NewDecoder(data)
	condition := &models.Condition{}
	err = decoder.Decode(condition)
	util.CheckErr(err)
	// Set condition patient
	condition.Patient = &models.Reference{
		Reference:    "Patient/" + patient.Id,
		Type:         "Patient",
		ReferencedID: patient.Id,
		External:     new(bool),
	}
	condition.Id = bson.NewObjectId().Hex()
	err = s.Database.C("conditions").Insert(condition)
	util.CheckErr(err)

	assertBundleCount(c, s.Server.URL+"/Condition", 1, 1)
	b := assertBundleCount(c, s.Server.URL+"/Condition?_include=Condition:patient", 2, 1)
	c.Assert(b.Entry[0].Resource, FitsTypeOf, &models.Condition{})
	c.Assert(b.Entry[0].Search.Mode, Equals, "match")
	c.Assert(b.Entry[1].Resource, FitsTypeOf, &models.Patient{})
	c.Assert(b.Entry[1].Search.Mode, Equals, "include")
}

func (s *ServerSuite) TestWrongResource(c *C) {
	data, err := os.Open("../fixtures/patient-wrong-type.json")
	util.CheckErr(err)
	defer data.Close()

	res, err := http.Post(s.Server.URL+"/Patient", "application/json", data)
	util.CheckErr(err)

	c.Assert(res.StatusCode, Equals, http.StatusBadRequest)
}

func (s *ServerSuite) TestUpdatePatient(c *C) {
	data, err := os.Open("../fixtures/patient-example-c.json")
	util.CheckErr(err)
	defer data.Close()

	client := &http.Client{}
	req, err := http.NewRequest("PUT", s.Server.URL+"/Patient/"+s.FixtureId, data)
	req.Header.Add("Content-Type", "application/json")
	util.CheckErr(err)
	_, err = client.Do(req)

	patientCollection := s.Database.C("patients")
	patient := models.Patient{}
	err = patientCollection.Find(bson.M{"_id": s.FixtureId}).One(&patient)
	util.CheckErr(err)
	c.Assert(patient.Name[0].Family[0], Equals, "Darkwing")
	c.Assert(patient.Meta, NotNil)
	c.Assert(patient.Meta.LastUpdated, NotNil)
	c.Assert(patient.Meta.LastUpdated.Precision, Equals, models.Precision(models.Timestamp))
	since := time.Since(patient.Meta.LastUpdated.Time)
	c.Assert(since.Hours() < float64(1), Equals, true)
	c.Assert(since.Minutes() < float64(1), Equals, true)
}

func (s *ServerSuite) TestDeletePatient(c *C) {

	data, err := os.Open("../fixtures/patient-example-d.json")
	util.CheckErr(err)
	defer data.Close()

	res, err := http.Post(s.Server.URL+"/Patient", "application/json", data)
	util.CheckErr(err)

	splitLocation := strings.Split(res.Header["Location"][0], "/")
	createdPatientId := splitLocation[len(splitLocation)-1]

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", s.Server.URL+"/Patient/"+createdPatientId, nil)
	util.CheckErr(err)
	_, err = client.Do(req)

	patientCollection := s.Database.C("patients")

	count, err := patientCollection.Find(bson.M{"_id": createdPatientId}).Count()
	c.Assert(count, Equals, 0)
}

func performSearch(c *C, url string) *models.Bundle {
	res, err := http.Get(url)
	util.CheckErr(err)
	decoder := json.NewDecoder(res.Body)
	bundle := &models.Bundle{}
	err = decoder.Decode(bundle)
	util.CheckErr(err)
	return bundle
}

func assertBundleCount(c *C, url string, expectedResults int, expectedTotal int) *models.Bundle {
	bundle := performSearch(c, url)
	c.Assert(len(bundle.Entry), Equals, expectedResults)
	c.Assert(*bundle.Total, Equals, uint32(expectedTotal))
	return bundle
}

func assertPagingLink(c *C, link models.BundleLinkComponent, relation string, count int, offset int) {
	c.Assert(link.Relation, Equals, relation)

	urlStr := link.Url
	urlUrl, err := url.Parse(urlStr)
	util.CheckErr(err)
	v := urlUrl.Query()

	c.Assert(v.Get(search.CountParam), Equals, fmt.Sprint(count))
	c.Assert(v.Get(search.OffsetParam), Equals, fmt.Sprint(offset))
}

func assertPagingLinkWithParams(c *C, link models.BundleLinkComponent, relation string, values url.Values, count int, offset int) {
	c.Assert(link.Relation, Equals, relation)

	urlStr := link.Url
	urlUrl, err := url.Parse(urlStr)
	util.CheckErr(err)
	v := urlUrl.Query()

	for key, val := range values {
		c.Assert(v[key], DeepEquals, val)
	}
	c.Assert(v.Get(search.CountParam), Equals, fmt.Sprint(count))
	c.Assert(v.Get(search.OffsetParam), Equals, fmt.Sprint(offset))
}

func (s *ServerSuite) insertPatientFromFixture(filePath string) *models.Patient {
	patientCollection := s.Database.C("patients")
	patient := loadPatientFromFixture(filePath)
	patient.Id = bson.NewObjectId().Hex()
	err := patientCollection.Insert(patient)
	util.CheckErr(err)
	return patient
}

func loadPatientFromFixture(fileName string) *models.Patient {
	data, err := os.Open(fileName)
	util.CheckErr(err)
	defer data.Close()

	decoder := json.NewDecoder(data)
	patient := &models.Patient{}
	err = decoder.Decode(patient)
	util.CheckErr(err)
	return patient
}
