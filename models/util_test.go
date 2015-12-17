package models

import (
	"encoding/json"
	"os"

	"github.com/pebbe/util"
	check "gopkg.in/check.v1"
)

type UtilSuite struct {
}

var _ = check.Suite(&UtilSuite{})

func (s *UtilSuite) TestUnmarshalJSON(c *check.C) {
	i := LoadMapFromFixture("../fixtures/condition.json")
	t := make(map[string]interface{})
	c.Assert(i, check.FitsTypeOf, t)

	r, ok := MapToResource(i, false).(Condition)
	c.Assert(ok, check.Equals, true)
	c.Assert(r.Patient.Reference, check.Equals, "https://example.com/base/Patient/4954037118555241963")
	c.Assert(r.Code.Coding, check.HasLen, 3)
	c.Assert(r.Code.MatchesCode("http://snomed.info/sct", "10091002"), check.Equals, true)
	c.Assert(r.Code.MatchesCode("http://hl7.org/fhir/sid/icd-9", "428.0"), check.Equals, true)
	c.Assert(r.Code.MatchesCode("http://hl7.org/fhir/sid/icd-10", "I50.1"), check.Equals, true)
	c.Assert(r.Code.Text, check.Equals, "Heart failure")
	c.Assert(r.OnsetDateTime.Precision, check.Equals, Precision(Timestamp))
	c.Assert(r.OnsetDateTime.Time.Unix(), check.Equals, int64(1330603200))
}

func LoadMapFromFixture(fileName string) interface{} {
	data, err := os.Open(fileName)
	util.CheckErr(err)
	defer data.Close()

	decoder := json.NewDecoder(data)
	i := make(map[string]interface{})
	err = decoder.Decode(&i)
	util.CheckErr(err)
	return i
}
