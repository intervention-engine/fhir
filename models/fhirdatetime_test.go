package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/pebbe/util"
	check "gopkg.in/check.v1"
)

type FDSuite struct {
	JSONBlob []byte
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

func (s *FDSuite) SetUpSuite(c *check.C) {}

var _ = check.Suite(&FDSuite{})

type Simple struct {
	Foo []FHIRDateTime `bson:"foo,omitempty" json:"foo,omitempty"`
}

func (s *FDSuite) TestFHIRDateTime(c *check.C) {
	simple := &Simple{}

	data := []byte("{ \"foo\": [\"1991-02-01T10:00:00-05:00\", \"1992-02-01\", \"1993-02-01T10:00:00-05:00\"]}")
	err := json.Unmarshal(data, &simple)
	util.CheckErr(err)

	c.Assert(simple.Foo, check.HasLen, 3)
	loc, err := time.LoadLocation("America/New_York")
	c.Assert(simple.Foo[0].Time.Equal(time.Date(1991, time.February, 1, 10, 0, 0, 0, loc)), check.Equals, true)
	c.Assert(simple.Foo[0].Precision, check.Equals, Precision(Timestamp))
	c.Assert(simple.Foo[1].Time.Equal(time.Date(1992, time.February, 1, 0, 0, 0, 0, time.Local)), check.Equals, true)
	c.Assert(simple.Foo[1].Precision, check.Equals, Precision(Date))
	c.Assert(simple.Foo[2].Time.Equal(time.Date(1993, time.February, 1, 10, 0, 0, 0, loc)), check.Equals, true)
	c.Assert(simple.Foo[2].Precision, check.Equals, Precision(Timestamp))
}
