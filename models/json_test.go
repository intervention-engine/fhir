package models

import (
	"encoding/json"
	"io/ioutil"
	"time"

	simple "github.com/bitly/go-simplejson"
	"github.com/pebbe/util"
	check "gopkg.in/check.v1"
)

type JSONSuite struct {
}

var _ = check.Suite(&JSONSuite{})

func (s *JSONSuite) TestUnmarshalJSON(c *check.C) {
	file, err := ioutil.ReadFile("../fixtures/loaded_condition.json")
	util.CheckErr(err)

	r := &Condition{}
	err = json.Unmarshal(file, r)
	util.CheckErr(err)

	c.Assert(r.Id, check.Equals, "8664777288161060797")
	c.Assert(r.Meta, check.NotNil)
	c.Assert(r.Meta.VersionId, check.Equals, "8664777288161060797-1")
	c.Assert(r.Meta.LastUpdated.Precision, check.Equals, Precision(Timestamp))
	c.Assert(r.Meta.LastUpdated.Time.Unix(), check.Equals, int64(1392208200))
	c.Assert(r.Meta.Tag, check.HasLen, 2)
	c.Assert(r.Meta.Tag[0].System, check.Equals, "http://intervention-engine.org/tag")
	c.Assert(r.Meta.Tag[0].Code, check.Equals, "FOO")
	c.Assert(r.Meta.Tag[1].System, check.Equals, "http://intervention-engine.org/tag")
	c.Assert(r.Meta.Tag[1].Code, check.Equals, "BAR")
	c.Assert(r.Text.Status, check.Equals, "additional")
	c.Assert(r.Text.Div, check.Equals, "<div>HTML in JavaScript.  Wow.</div>")
	c.Assert(r.Contained, check.HasLen, 1)
	c.Assert(r.Contained[0], check.FitsTypeOf, &Practitioner{})
	contained := r.Contained[0].(*Practitioner)
	c.Assert(contained.Id, check.Equals, "pract1")
	c.Assert(contained.Name.Family, check.HasLen, 1)
	c.Assert(contained.Name.Family[0], check.Equals, "Doofenshmirtz")
	c.Assert(contained.Name.Given, check.HasLen, 1)
	c.Assert(contained.Name.Given[0], check.Equals, "Heinz")
	c.Assert(contained.Name.Suffix, check.HasLen, 1)
	c.Assert(contained.Name.Suffix[0], check.Equals, "MD")
	c.Assert(r.Patient.Reference, check.Equals, "https://example.com/base/Patient/4954037118555241963")
	c.Assert(r.Asserter.Reference, check.Equals, "#pract1")
	c.Assert(r.Code.Coding, check.HasLen, 3)
	c.Assert(r.Code.MatchesCode("http://snomed.info/sct", "10091002"), check.Equals, true)
	c.Assert(r.Code.MatchesCode("http://hl7.org/fhir/sid/icd-9", "428.0"), check.Equals, true)
	c.Assert(r.Code.MatchesCode("http://hl7.org/fhir/sid/icd-10", "I50.1"), check.Equals, true)
	c.Assert(r.Code.Text, check.Equals, "Heart failure")
	c.Assert(r.OnsetDateTime.Precision, check.Equals, Precision(Timestamp))
	c.Assert(r.OnsetDateTime.Time.Unix(), check.Equals, int64(1330603200))
}

func (s *JSONSuite) TestMarshalJSON(c *check.C) {
	r := &Condition{}
	r.Id = "8664777288161060797"
	r.Meta = &Meta{
		VersionId:   "8664777288161060797-1",
		LastUpdated: &FHIRDateTime{Time: time.Date(2014, time.February, 12, 12, 30, 0, 0, time.UTC), Precision: Precision(Timestamp)},
		Tag: []Coding{
			{System: "http://intervention-engine.org/tag", Code: "FOO"},
			{System: "http://intervention-engine.org/tag", Code: "BAR"},
		},
	}
	r.Text = &Narrative{
		Status: "additional",
		Div:    "<div>HTML in JavaScript.  Wow.</div>",
	}
	r.Contained = make([]interface{}, 1)
	r.Contained[0] = &Practitioner{
		DomainResource: DomainResource{Resource: Resource{Id: "pract1"}},
		Name: &HumanName{
			Family: []string{"Doofenshmirtz"},
			Given:  []string{"Heinz"},
			Suffix: []string{"MD"},
		},
	}
	r.Patient = &Reference{Reference: "https://example.com/base/Patient/4954037118555241963"}
	r.Asserter = &Reference{Reference: "#pract1"}
	r.Code = &CodeableConcept{
		Coding: []Coding{
			{System: "http://snomed.info/sct", Code: "10091002"},
			{System: "http://hl7.org/fhir/sid/icd-9", Code: "428.0"},
			{System: "http://hl7.org/fhir/sid/icd-10", Code: "I50.1"},
		},
		Text: "Heart failure",
	}
	r.OnsetDateTime = &FHIRDateTime{Time: time.Date(2012, time.March, 01, 12, 0, 0, 0, time.UTC), Precision: Precision(Timestamp)}

	b, err := json.Marshal(r)
	util.CheckErr(err)

	j, err := simple.NewJson(b)
	util.CheckErr(err)

	c.Assert(j.Get("resourceType").MustString(), check.Equals, "Condition")
	c.Assert(j.GetPath("meta", "versionId").MustString(), check.Equals, "8664777288161060797-1")
	c.Assert(j.GetPath("meta", "lastUpdated").MustString(), check.Equals, "2014-02-12T12:30:00Z")
	c.Assert(j.GetPath("meta", "tag").GetIndex(0).Get("system").MustString(), check.Equals, "http://intervention-engine.org/tag")
	c.Assert(j.GetPath("meta", "tag").GetIndex(0).Get("code").MustString(), check.Equals, "FOO")
	c.Assert(j.GetPath("meta", "tag").GetIndex(1).Get("system").MustString(), check.Equals, "http://intervention-engine.org/tag")
	c.Assert(j.GetPath("meta", "tag").GetIndex(1).Get("code").MustString(), check.Equals, "BAR")
	c.Assert(j.GetPath("text", "status").MustString(), check.Equals, "additional")
	c.Assert(j.GetPath("text", "div").MustString(), check.Equals, "<div>HTML in JavaScript.  Wow.</div>")
	contained := j.Get("contained").GetIndex(0)
	c.Assert(contained.Get("resourceType").MustString(), check.Equals, "Practitioner")
	c.Assert(contained.GetPath("name", "family").GetIndex(0).MustString(), check.Equals, "Doofenshmirtz")
	c.Assert(contained.GetPath("name", "given").GetIndex(0).MustString(), check.Equals, "Heinz")
	c.Assert(contained.GetPath("name", "suffix").GetIndex(0).MustString(), check.Equals, "MD")
	c.Assert(j.GetPath("patient", "reference").MustString(), check.Equals, "https://example.com/base/Patient/4954037118555241963")
	c.Assert(j.GetPath("asserter", "reference").MustString(), check.Equals, "#pract1")
	c.Assert(j.GetPath("code", "coding").MustArray(), check.HasLen, 3)
	c.Assert(j.GetPath("code", "coding").GetIndex(0).Get("system").MustString(), check.Equals, "http://snomed.info/sct")
	c.Assert(j.GetPath("code", "coding").GetIndex(0).Get("code").MustString(), check.Equals, "10091002")
	c.Assert(j.GetPath("code", "coding").GetIndex(1).Get("system").MustString(), check.Equals, "http://hl7.org/fhir/sid/icd-9")
	c.Assert(j.GetPath("code", "coding").GetIndex(1).Get("code").MustString(), check.Equals, "428.0")
	c.Assert(j.GetPath("code", "coding").GetIndex(2).Get("system").MustString(), check.Equals, "http://hl7.org/fhir/sid/icd-10")
	c.Assert(j.GetPath("code", "coding").GetIndex(2).Get("code").MustString(), check.Equals, "I50.1")
	c.Assert(j.GetPath("code", "text").MustString(), check.Equals, "Heart failure")
	c.Assert(j.Get("onsetDateTime").MustString(), check.Equals, "2012-03-01T12:00:00Z")
}
