package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
)

type SmartAuthSuite struct {
	Echo *echo.Echo
}

var _ = Suite(&SmartAuthSuite{})

func (s *SmartAuthSuite) SetUpTest(c *C) {
	s.Echo = echo.New()
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
	var buf bytes.Buffer
	r, err := http.NewRequest(method, "http://fhir-server/", &buf)
	util.CheckErr(err)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-DELEGATED", "true")
	if scopes != "" {
		r.Header.Add("X-SCOPE", scopes)
	}
	rw := httptest.NewRecorder()
	ctx := echo.NewContext(r, echo.NewResponse(rw, s.Echo), s.Echo)

	noop := func(c *echo.Context) error { return c.String(http.StatusOK, "Hello") }
	handlerGenerator := SmartAuthHandler("Patient")
	handler := handlerGenerator(noop)
	err = handler(ctx)
	util.CheckErr(err)
	return rw
}
