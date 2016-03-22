package server

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
)

type SmartAuthSuite struct {
}

var _ = Suite(&SmartAuthSuite{})

func (s *SmartAuthSuite) SetUpTest(c *C) {
}

func (s *SmartAuthSuite) TestGetPatientWithoutScopes(c *C) {
	rr := s.SetUpRequest("GET", "")
	c.Assert(rr.Code, Equals, http.StatusForbidden)
}

func (s *SmartAuthSuite) TestGetPatientWithWriteScopes(c *C) {
	rr := s.SetUpRequest("GET", "user/Patient.write")
	c.Assert(rr.Code, Equals, http.StatusForbidden)
}

func (s *SmartAuthSuite) TestGetPatientWithScopes(c *C) {
	rr := s.SetUpRequest("GET", "user/Patient.read")
	c.Assert(rr.Code, Equals, http.StatusOK)
	c.Assert(rr.Body.String(), Equals, "Hello")
}

func (s *SmartAuthSuite) TestGetPatientWithMultipleScopes(c *C) {
	rr := s.SetUpRequest("GET", "user/Patient.read user/Observation.* user/Condition.write")
	c.Assert(rr.Code, Equals, http.StatusOK)
	c.Assert(rr.Body.String(), Equals, "Hello")
}

func (s *SmartAuthSuite) TestGetPatientWithWildcard(c *C) {
	rr := s.SetUpRequest("GET", "user/Patient.*")
	c.Assert(rr.Code, Equals, http.StatusOK)
	c.Assert(rr.Body.String(), Equals, "Hello")
}

func (s *SmartAuthSuite) TestGetPatientWithAllWildcard(c *C) {
	rr := s.SetUpRequest("GET", "user/*.*")
	c.Assert(rr.Code, Equals, http.StatusOK)
	c.Assert(rr.Body.String(), Equals, "Hello")
}

func (s *SmartAuthSuite) TestPostPatientWithScopes(c *C) {
	rr := s.SetUpRequest("POST", "user/Patient.write")
	c.Assert(rr.Code, Equals, http.StatusOK)
	c.Assert(rr.Body.String(), Equals, "Hello")
}

func (s *SmartAuthSuite) SetUpRequest(method, scopes string) *httptest.ResponseRecorder {
	r, err := http.NewRequest(method, "/", nil)
	util.CheckErr(err)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-DELEGATED", "true")
	if scopes != "" {
		r.Header.Add("X-SCOPE", scopes)
	}
	e := gin.New()
	rw := httptest.NewRecorder()
	noop := func(c *gin.Context) { c.String(http.StatusOK, "Hello") }
	authHandler := SmartAuthHandler("Patient")
	e.GET("/", authHandler, noop)
	e.POST("/", authHandler, noop)
	e.ServeHTTP(rw, r)
	return rw
}
