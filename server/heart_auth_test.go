package server

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
)

type HEARTScopesSuite struct {
}

var _ = Suite(&HEARTScopesSuite{})

func (s *HEARTScopesSuite) SetUpTest(c *C) {
}

func (s *HEARTScopesSuite) TestGetPatientWithoutScopes(c *C) {
	rr := s.SetUpRequest("GET", "")
	c.Assert(rr.Code, Equals, http.StatusForbidden)
}

func (s *HEARTScopesSuite) TestGetPatientWithWriteScopes(c *C) {
	rr := s.SetUpRequest("GET", "user/Patient.write")
	c.Assert(rr.Code, Equals, http.StatusForbidden)
}

func (s *HEARTScopesSuite) TestGetPatientWithScopes(c *C) {
	rr := s.SetUpRequest("GET", "user/Patient.read")
	c.Assert(rr.Code, Equals, http.StatusOK)
	c.Assert(rr.Body.String(), Equals, "Hello")
}

func (s *HEARTScopesSuite) TestGetPatientWithMultipleScopes(c *C) {
	rr := s.SetUpRequest("GET", "user/Patient.read user/Observation.* user/Condition.write")
	c.Assert(rr.Code, Equals, http.StatusOK)
	c.Assert(rr.Body.String(), Equals, "Hello")
}

func (s *HEARTScopesSuite) TestGetPatientWithWildcard(c *C) {
	rr := s.SetUpRequest("GET", "user/Patient.*")
	c.Assert(rr.Code, Equals, http.StatusOK)
	c.Assert(rr.Body.String(), Equals, "Hello")
}

func (s *HEARTScopesSuite) TestGetPatientWithAllWildcard(c *C) {
	rr := s.SetUpRequest("GET", "user/*.*")
	c.Assert(rr.Code, Equals, http.StatusOK)
	c.Assert(rr.Body.String(), Equals, "Hello")
}

func (s *HEARTScopesSuite) TestPostPatientWithScopes(c *C) {
	rr := s.SetUpRequest("POST", "user/Patient.write")
	c.Assert(rr.Code, Equals, http.StatusOK)
	c.Assert(rr.Body.String(), Equals, "Hello")
}

func (s *HEARTScopesSuite) SetUpRequest(method, scopes string) *httptest.ResponseRecorder {
	r, err := http.NewRequest(method, "/", nil)
	util.CheckErr(err)
	r.Header.Add("Content-Type", "application/json")
	mockTokenIntrospection := func(c *gin.Context) {
		if scopes != "" {
			c.Set("scopes", strings.Split(scopes, " "))
		}
	}

	e := gin.New()
	rw := httptest.NewRecorder()
	noop := func(c *gin.Context) { c.String(http.StatusOK, "Hello") }
	authHandler := HEARTScopesHandler("Patient")
	e.GET("/", mockTokenIntrospection, authHandler, noop)
	e.POST("/", mockTokenIntrospection, authHandler, noop)
	e.ServeHTTP(rw, r)
	return rw
}
