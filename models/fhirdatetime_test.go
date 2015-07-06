package models

import (
	"encoding/json"
	"testing"

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
	c.Assert(simple.Foo[0].Time.Unix(), check.Equals, int64(665420400))
	c.Assert(simple.Foo[0].Precision, check.Equals, Precision(Timestamp))
	c.Assert(simple.Foo[1].Time.Unix(), check.Equals, int64(696902400))
	c.Assert(simple.Foo[1].Precision, check.Equals, Precision(Date))
	c.Assert(simple.Foo[2].Time.Unix(), check.Equals, int64(728578800))
	c.Assert(simple.Foo[2].Precision, check.Equals, Precision(Timestamp))
}
