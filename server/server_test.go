package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
  "bytes"
  "strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/pebbe/util"
	"gitlab.mitre.org/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
  . "gopkg.in/check.v1"
)

type ServerSuite struct {
	Session   *mgo.Session
	Router    *mux.Router
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
	Database = s.Session.DB("fhir-test")

	// Build routes for testing
	s.Router = mux.NewRouter()
	s.Router.StrictSlash(true)
	s.Router.KeepContext = true
	RegisterRoutes(s.Router, make(map[string][]negroni.Handler))

	// Create httptest server
	s.Server = httptest.NewServer(s.Router)

	// Add patient fixture
	patientCollection := Database.C("patients")
	patient := LoadPatientFromFixture("../fixtures/patient-example-a.json")
	i := bson.NewObjectId()
	s.FixtureId = i.Hex()
	patient.Id = s.FixtureId
	err = patientCollection.Insert(patient)
	util.CheckErr(err)
}

func (s *ServerSuite) TearDownSuite(c *C) {
  Database.DropDatabase()
  s.Session.Close()
	s.Server.Close()
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
  patientBundle := &models.PatientBundle{}
  err = decoder.Decode(patientBundle)
  util.CheckErr(err)

  var result []models.Patient
  collection := Database.C("patients")
  iter := collection.Find(nil).Iter()
  err = iter.All(&result)
  util.CheckErr(err)

  c.Assert(patientBundle.TotalResults, Equals, len(result))
  c.Assert(patientBundle.Title, Equals, "Patient Index")
}

func (s *ServerSuite) TestCreatePatient(c *C) {
  createPatient := LoadPatientFromFixture("../fixtures/patient-example-b.json")
  var buf bytes.Buffer
  encoder := json.NewEncoder(&buf)
  encoder.Encode(createPatient)
  res, err := http.Post(s.Server.URL + "/Patient", "application/json", &buf)
  util.CheckErr(err)

  splitLocation := strings.Split(res.Header["Location"][0], "/")
  createdPatientId := splitLocation[len(splitLocation)-1]

  res, err = http.Get(s.Server.URL + "/Patient/" + createdPatientId)
  util.CheckErr(err)

  decoder := json.NewDecoder(res.Body)
  patient := &models.Patient{}
  err = decoder.Decode(patient)
  util.CheckErr(err)
  c.Assert(patient.Name[0].Family[0], Equals, "Daffy")
}

func (s *ServerSuite) TestUpdatePatient(c *C) {
  updatePatient := LoadPatientFromFixture("../fixtures/patient-example-c.json")
  var buf bytes.Buffer
  encoder := json.NewEncoder(&buf)
  encoder.Encode(updatePatient)

  client := &http.Client{}
  req, err := http.NewRequest("PUT", s.Server.URL + "/Patient/" + s.FixtureId, &buf)
  util.CheckErr(err)
  res, err := client.Do(req)

  res, err = http.Get(s.Server.URL + "/Patient/" + s.FixtureId)
  util.CheckErr(err)

  decoder := json.NewDecoder(res.Body)
  patient := &models.Patient{}
  err = decoder.Decode(patient)
  util.CheckErr(err)
  c.Assert(patient.Name[0].Family[0], Equals, "Darkwing")
}

func (s *ServerSuite) TestDeletePatient(c *C) {

  createPatient := LoadPatientFromFixture("../fixtures/patient-example-d.json")
  var buf bytes.Buffer
  encoder := json.NewEncoder(&buf)
  encoder.Encode(createPatient)
  res, err := http.Post(s.Server.URL + "/Patient", "application/json", &buf)
  util.CheckErr(err)

  splitLocation := strings.Split(res.Header["Location"][0], "/")
  createdPatientId := splitLocation[len(splitLocation)-1]

  client := &http.Client{}
  req, err := http.NewRequest("DELETE", s.Server.URL + "/Patient/" + createdPatientId, nil)
  util.CheckErr(err)
  _, err = client.Do(req)

  patientCollection := Database.C("patients")

  count, err := patientCollection.Find(bson.M{"_id": createdPatientId}).Count()
  c.Assert(count, Equals, 0)
}

func LoadPatientFromFixture(fileName string) *models.Patient {
	data, err := os.Open(fileName)
	defer data.Close()
	util.CheckErr(err)
	decoder := json.NewDecoder(data)
	patient := &models.Patient{}
	err = decoder.Decode(patient)
	util.CheckErr(err)
	return patient
}
