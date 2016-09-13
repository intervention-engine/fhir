package server

import (
	"fmt"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/dbtest"
)

var expectedIndexes = []mgo.Index{
	mgo.Index{Key: []string{"foo"}},
	mgo.Index{Key: []string{"-foo"}},
	mgo.Index{Key: []string{"bar.foo"}},
	mgo.Index{Key: []string{"-bar.foo"}},
	mgo.Index{Key: []string{"foo", "bar"}},
	mgo.Index{Key: []string{"bar", "foo", "-baz"}},
}

type MongoIndexesTestSuite struct {
	suite.Suite
	DBServer     *dbtest.DBServer
	EST          *time.Location
	Local        *time.Location
	Session      *mgo.Session
	Database     *mgo.Database
	Engine       *gin.Engine
	Server       *httptest.Server
	Config       Config
	Interceptors map[string]InterceptorList
	FixtureID    string
}

func (s *MongoIndexesTestSuite) SetupSuite() {
	s.EST = time.FixedZone("EST", -5*60*60)
	s.Local, _ = time.LoadLocation("Local")

	// Server configuration
	s.Config = DefaultConfig
	s.Config.DatabaseName = "fhir-test"
	s.Config.IndexConfigPath = "../fixtures/test_indexes.conf"

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
	s.Database = s.Session.DB(s.Config.DatabaseName)

	// Set gin to release mode (less verbose output)
	gin.SetMode(gin.ReleaseMode)

	// Build routes for testing
	s.Engine = gin.New()
	RegisterRoutes(s.Engine, make(map[string][]gin.HandlerFunc), NewMongoDataAccessLayer(s.Database, s.Interceptors), s.Config)

	// Create httptest server
	s.Server = httptest.NewServer(s.Engine)
}

func (s *MongoIndexesTestSuite) TearDownSuite() {
	s.Session.Close()
	s.DBServer.Wipe()
	s.DBServer.Stop()

	// remove the temporary database directory
	var err error
	err = removeContents("./testdb")

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

func (s *MongoIndexesTestSuite) TestParseIndexesStandardIndexAsc() {

	indexStr := "fhir-test.testcollection.foo_1"
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(err, "Should return without error")

	indexes, ok := indexMap["testcollection"]
	s.True(ok, "'testcollection' should be a key in the index map")
	s.Equal(len(indexes), 1, "indexMap[testcollection] should contain one index")
	s.Equal(len(indexes[0].Key), 1, "The created index should contain one key")
	s.Equal(indexes[0].Key[0], "foo", "The index key should be 'foo'")

	// We only need to check this once, since it's done for all successful indexes
	s.True(indexes[0].Background, "The index should be set to build in the background")
}

func (s *MongoIndexesTestSuite) TestParseIndexesStandardIndexDesc() {

	indexStr := "fhir-test.testcollection.foo_-1"
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(err, "Should return without error")

	indexes, ok := indexMap["testcollection"]
	s.True(ok, "'testcollection' should be a key in the index map")
	s.Equal(len(indexes), 1, "indexMap[testcollection] should contain one index")
	s.Equal(len(indexes[0].Key), 1, "The created index should contain one key")
	s.Equal(indexes[0].Key[0], "-foo", "The index key should be '-foo'")
}

func (s *MongoIndexesTestSuite) TestParseIndexesCompoundIndexAsc() {

	indexStr := "fhir-test.testcollection.(foo_1, bar_1)"
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(err, "Should return without error")

	indexes, ok := indexMap["testcollection"]
	s.True(ok, "'testcollection' should be a key in the index map")
	s.Equal(len(indexes), 1, "indexMap[testcollection] should contain one index")
	s.Equal(len(indexes[0].Key), 2, "The created index should contain 2 keys")
	s.Equal(indexes[0].Key[0], "foo", "The prefix index key should be 'foo'")
	s.Equal(indexes[0].Key[1], "bar", "The second index key should be 'bar'")
}

func (s *MongoIndexesTestSuite) TestParseIndexesCompoundIndexDesc() {

	indexStr := "fhir-test.testcollection.(foo_-1, bar_-1)"
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(err, "Should return without error")

	indexes, ok := indexMap["testcollection"]
	s.True(ok, "'testcollection' should be a key in the index map")
	s.Equal(len(indexes), 1, "indexMap[testcollection] should contain one index")
	s.Equal(len(indexes[0].Key), 2, "The created index should contain 2 keys")
	s.Equal(indexes[0].Key[0], "-foo", "The prefix index key should be '-foo'")
	s.Equal(indexes[0].Key[1], "-bar", "The second index key should be '-bar'")
}

func (s *MongoIndexesTestSuite) TestParseIndexesCompoundIndexMixed() {

	indexStr := "fhir-test.testcollection.(foo_-1, bar_1)"
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(err, "Should return without error")

	indexes, ok := indexMap["testcollection"]
	s.True(ok, "'testcollection' should be a key in the index map")
	s.Equal(len(indexes), 1, "indexMap[testcollection] should contain one index")
	s.Equal(len(indexes[0].Key), 2, "The created index should contain 2 keys")
	s.Equal(indexes[0].Key[0], "-foo", "The prefix index key should be '-foo'")
	s.Equal(indexes[0].Key[1], "bar", "The second index key should be 'bar'")
}

func (s *MongoIndexesTestSuite) TestParseIndexesNoIndexes() {

	indexStr := ""
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	keys := getKeys(indexMap)
	s.Nil(err, "Should return without error")
	s.Equal(len(keys), 0, "Should return a map with no indexes in it")
}

func (s *MongoIndexesTestSuite) TestParseIndexesBadIndexFormat() {

	indexStr := "asdfasdf"
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(indexMap, "IndexMap should be nil")
	s.NotNil(err, "Should return an error")
	s.Equal(err.Error(), "Index 'asdfasdf' is invalid: Not of format <db_name>.<collection_name>.<index(es)>", "Error message should match expected", "Unexpected error returned")
}

func (s *MongoIndexesTestSuite) TestParseIndexesBadDbName() {

	indexStr := "foo.testcollection.(foo_-1, bar_1)"
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(indexMap, "IndexMap should be nil")
	s.NotNil(err, "Should return an error")
	s.Equal(err.Error(), "Index 'foo.testcollection.(foo_-1, bar_1)' is invalid: DB name does not match server configuration", "Unexpected error returned")
}

func (s *MongoIndexesTestSuite) TestParseIndexesNoCollectionName() {

	indexStr := "fhir-test..(foo_-1, bar_1)"
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(indexMap, "IndexMap should be nil")
	s.NotNil(err, "Should return an error")
	s.Equal(err.Error(), "Index 'fhir-test..(foo_-1, bar_1)' is invalid: No collection name given", "Unexpected error returned")
}

func (s *MongoIndexesTestSuite) TestParseIndexesNoKeys() {

	indexStr := "fhir-test.testcollection."
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(indexMap, "IndexMap should be nil")
	s.NotNil(err, "Should return an error")
	s.Equal(err.Error(), "Index 'fhir-test.testcollection.' is invalid: No index key(s) given", "Unexpected error returned")
}

func (s *MongoIndexesTestSuite) TestParseIndexesBadStandardKeyFormat() {

	indexStr := "fhir-test.testcollection.foo"
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(indexMap, "IndexMap should be nil")
	s.NotNil(err, "Should return an error")
	s.Equal(err.Error(), "Index 'fhir-test.testcollection.foo' is invalid: Standard key not of format: <key>_(-)1", "Unexpected error returned")
}

func (s *MongoIndexesTestSuite) TestParseIndexesBadCompoundKeyFormat() {

	indexStr := "fhir-test.testcollection.(foobar"
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(indexMap, "IndexMap should be nil")
	s.NotNil(err, "Should return an error")
	s.Equal(err.Error(), "Index 'fhir-test.testcollection.(foobar' is invalid: Compound key not of format: (<key1>_(-)1, <key2>_(-)1, ...)", "Unexpected error returned")
}

func (s *MongoIndexesTestSuite) TestParseIndexesBadCompoundKeySubKeyFormat() {

	indexStr := "fhir-test.testcollection.(foo, bar_1)"
	indexMap, err := parseIndexes(indexStr, "fhir-test")
	s.Nil(indexMap, "IndexMap should be nil")
	s.NotNil(err, "Should return an error")
	s.Equal(err.Error(), "Index 'fhir-test.testcollection.(foo, bar_1)' is invalid: Compound key sub-key not of format: <key>_(-)1", "Unexpected error returned")
}

func (s *MongoIndexesTestSuite) TestConfigureIndexes() {
	// Configure test indexes
	ConfigureIndexes(s.Session.Copy(), s.Config)

	// get the "testcollection" collection. This should have been auto-magically
	// created by ConfigureIndexes
	c := s.Database.C("testcollection")

	// get the indexes for this collection
	indexes, err := c.Indexes()

	if err != nil {
		panic(err)
	}

	// The indexes *should* be returned in the order they were created.
	// That's how they're returned by queries in the mongo shell at least.
	s.compareIndexes(expectedIndexes, indexes)
}

func (s *MongoIndexesTestSuite) TestConfigureIndexesNoConfigFile() {

	s.Config.IndexConfigPath = "./does_not_exist.conf"
	s.NotPanics(func() { ConfigureIndexes(s.Session.Copy(), s.Config) }, "Should not panic if no config file is found")
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

func removeContents(dir string) error {
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

// getKeys returns a slice of all keys in an IndexMap in sorted order
func getKeys(keyMap IndexMap) []string {
	keys := make([]string, len(keyMap))
	var i = 0
	for k := range keyMap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
