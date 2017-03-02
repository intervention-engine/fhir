package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gopkg.in/mgo.v2/dbtest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type MiddlewareTestSuite struct {
	suite.Suite
	DBServer      *dbtest.DBServer
	MasterSession *MasterSession
}

func TestMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}

func (m *MiddlewareTestSuite) SetupSuite() {
	// Create a temporary directory for the test database
	var err error
	err = os.Mkdir("./testdb", 0775)

	if err != nil {
		panic(err)
	}

	// setup the mongo database
	m.DBServer = &dbtest.DBServer{}
	m.DBServer.SetPath("./testdb")
	m.MasterSession = NewMasterSession(m.DBServer.Session(), "fhir-test")

	// Set gin to release mode (less verbose output)
	gin.SetMode(gin.ReleaseMode)
}

func (m *MiddlewareTestSuite) TearDownSuite() {
	m.MasterSession.session.Close()
	m.DBServer.Wipe()
	m.DBServer.Stop()

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

func (m *MiddlewareTestSuite) TestRejectXML() {
	e := gin.New()
	e.Use(AbortNonJSONRequestsMiddleware)
	RegisterRoutes(e, nil, NewMongoDataAccessLayer(m.MasterSession, nil, DefaultConfig), DefaultConfig)
	server := httptest.NewServer(e)

	req, err := http.NewRequest("GET", server.URL+"/Patient", nil)
	m.NoError(err)
	req.Header.Add("Accept", "application/xml")
	resp, err := http.DefaultClient.Do(req)
	m.Equal(http.StatusNotAcceptable, resp.StatusCode)
}

func (m *MiddlewareTestSuite) TestReadOnlyMode() {
	e := gin.New()
	e.Use(ReadOnlyMiddleware)
	config := DefaultConfig
	config.ReadOnly = true
	RegisterRoutes(e, nil, NewMongoDataAccessLayer(m.MasterSession, nil, config), config)
	server := httptest.NewServer(e)

	req, err := http.NewRequest("POST", server.URL+"/Patient", nil)
	m.NoError(err)
	resp, err := http.DefaultClient.Do(req)
	m.Equal(http.StatusMethodNotAllowed, resp.StatusCode)
}
