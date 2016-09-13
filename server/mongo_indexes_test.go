package server

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2"
)

type MongoIndexesTestSuite struct {
	suite.Suite
	Database     *mgo.Database
	Session      *mgo.Session
	Engine       *gin.Engine
	Server       *httptest.Server
	Interceptors map[string]InterceptorList
	FixtureID    string
}

func (s *MongoIndexesTestSuite) SetupSuite() {
	// Server configuration
	config := DefaultConfig
	config.DatabaseName = "fhir-test"
	config.IndexConfigPath = "../fixtures/test_indexes.conf"

	// setup the mongo database
	var err error
	s.Session, err = mgo.Dial(config.ServerURL)

	if err != nil {
		panic(err)
	}

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
	s.Database.DropDatabase()
	s.Session.Close()
	s.Server.Close()
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
