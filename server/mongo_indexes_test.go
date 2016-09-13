package server

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/dbtest"
)

type MongoIndexesTestSuite struct {
	suite.Suite
	DBServer     *dbtest.DBServer
	EST          *time.Location
	Local        *time.Location
	Session      *mgo.Session
	Database     *mgo.Database
	Engine       *gin.Engine
	Server       *httptest.Server
	Interceptors map[string]InterceptorList
	FixtureID    string
}

func (s *MongoIndexesTestSuite) SetupSuite() {
	s.EST = time.FixedZone("EST", -5*60*60)
	s.Local, _ = time.LoadLocation("Local")

	// Server configuration
	config := DefaultConfig
	config.DatabaseName = "fhir-test"
	config.IndexConfigPath = "../fixtures/test_indexes.conf"

	// Create a temporary directory for the test database
	var err error
	err = os.Mkdir("./testdb", 0775)

	if err != nil {
		panic(err)
	}

	// setup the mongo database
	s.DBServer = &dbtest.DBServer{}
	s.DBServer.SetPath("./testdb")
	s.Session = s.DBServer.Session()
	s.Database = s.Session.DB(config.DatabaseName)

	// Set gin to release mode (less verbose output)
	gin.SetMode(gin.ReleaseMode)

	// Build routes for testing
	s.Engine = gin.New()
	RegisterRoutes(s.Engine, make(map[string][]gin.HandlerFunc), NewMongoDataAccessLayer(s.Database, s.Interceptors), config)

	// Configure test indexes
	ConfigureIndexes(s.Session.Copy(), config)

	// Create httptest server
	s.Server = httptest.NewServer(s.Engine)
}

func (s *MongoIndexesTestSuite) TearDownSuite() {
	s.Session.Close()
	s.DBServer.Wipe()
	s.DBServer.Stop()

	// remove the temporary database directory
	var err error
	err = RemoveContents("./testdb")

	if err != nil {
		panic(err)
	}

	err = os.Remove("./testdb")

	if err != nil {
		panic(err)
	}
}

func TestMongoIndexes(t *testing.T) {
	// bind test suite to go test
	suite.Run(t, new(MongoIndexesTestSuite))
}

func (s *MongoIndexesTestSuite) TestParseIndexes() {

	// Test that we're parsing the index names in the config file correctly
	testIdxConf, err := ioutil.ReadFile("../fixtures/test_indexes.conf")

	if err != nil {
		panic(err)
	}

	allIndexes, err := parseIndexes(string(testIdxConf), "fhir-test")

	s.Nil(err, "The test index configuration file should be parsed without error")

	collectionIndexes, ok := allIndexes["testcollection"]
	s.Equal(true, ok, "'testcollection' should be a key in the parsed indexes map")

	expectedIndexes := []mgo.Index{
		mgo.Index{Key: []string{"foo"}},
		mgo.Index{Key: []string{"-foo"}},
		mgo.Index{Key: []string{"bar.foo"}},
		mgo.Index{Key: []string{"-bar.foo"}},
	}

	s.compareIndexes(expectedIndexes, collectionIndexes)
}

func (s *MongoIndexesTestSuite) TestIndexesCreated() {
	var err error

	// get the "testcollection" collection. This should have been auto-magically
	// created the first time the "server" was run
	c := s.Database.C("testcollection")

	// get the indexes for this collection
	indexes, err := c.Indexes()

	if err != nil {
		panic(err)
	}

	// The indexes *should* be returned in the order they were created.
	// That's how they're returned by queries in the mongo shell at least.
	expectedIndexes := []mgo.Index{
		mgo.Index{Key: []string{"foo"}},
		mgo.Index{Key: []string{"-foo"}},
		mgo.Index{Key: []string{"bar.foo"}},
		mgo.Index{Key: []string{"-bar.foo"}},
	}

	s.compareIndexes(expectedIndexes, indexes)
}

func (s *MongoIndexesTestSuite) compareIndexes(expected, actual []mgo.Index) {

	for _, idx := range actual {
		s.True(len(idx.Key) > 0, "Index should have at least 1 key")

		if idx.Key[0] == "_id" {
			// Skip testing the indexes created by the system
			continue
		}

		s.True(indexInSlice(expected, idx), fmt.Sprintf("Index fhir-test.testcollection: %s was not parsed correctly", idx.Key[0]))
	}
}

func indexInSlice(indexesSlice []mgo.Index, want mgo.Index) bool {
	// Compares two indexes by Key only
	if len(want.Key) > 0 {
		for _, idx := range indexesSlice {
			if want.Key[0] == idx.Key[0] {
				return true
			}
		}
	}
	return false
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
