package models

import (
	"encoding/json"
	"os"
	"time"

	simple "github.com/bitly/go-simplejson"
	"github.com/pebbe/util"
	check "gopkg.in/check.v1"
)

type BundleSuite struct {
}

var _ = check.Suite(&BundleSuite{})

func (s *BundleSuite) TestMarshalJSON(c *check.C) {
	// Build a bundle
	total := uint32(2)
	bundle := Bundle{Type: "searchset", Total: &total}
	bundle.Link = make([]BundleLinkComponent, 1)
	bundle.Link[0].Relation = "self"
	bundle.Link[0].Url = "https://example.com/base/Condition?patient=4954037118555241963"
	bundle.Entry = make([]BundleEntryComponent, 2)
	score := float64(1)
	search := BundleEntrySearchComponent{Mode: "match", Score: &score}
	bundle.Entry[0].Search = &search
	condition := Condition{}
	ref := Reference{Reference: "https://example.com/base/Patient/4954037118555241963"}
	condition.Patient = &ref
	code := CodeableConcept{Text: "Heart failure"}
	code.Coding = make([]Coding, 3)
	code.Coding[0] = Coding{System: "http://snomed.info/sct", Code: "10091002"}
	code.Coding[1] = Coding{System: "http://hl7.org/fhir/sid/icd-9", Code: "428.0"}
	code.Coding[2] = Coding{System: "http://hl7.org/fhir/sid/icd-10", Code: "I50.1"}
	condition.Code = &code
	dateTime := FHIRDateTime{Time: time.Date(2012, time.March, 1, 12, 0, 0, 0, time.UTC), Precision: Precision(Timestamp)}
	condition.OnsetDateTime = &dateTime
	bundle.Entry[0].Resource = condition

	score2 := float64(2)
	search2 := BundleEntrySearchComponent{Mode: "match", Score: &score2}
	bundle.Entry[1].Search = &search2
	condition2 := Condition{}
	ref2 := Reference{Reference: "https://example.com/base/Patient/4954037118555241963"}
	condition2.Patient = &ref2
	code2 := CodeableConcept{Text: "Moderate left ventricular systolic dysfunction"}
	code2.Coding = make([]Coding, 1)
	code2.Coding[0] = Coding{System: "http://snomed.info/sct", Code: "981000124106"}
	condition2.Code = &code2
	dateTime2 := FHIRDateTime{Time: time.Date(2012, time.March, 1, 12, 5, 0, 0, time.UTC), Precision: Precision(Timestamp)}
	condition2.OnsetDateTime = &dateTime2
	bundle.Entry[1].Resource = condition2

	// Marshal it to JSON
	b, err := json.Marshal(&bundle)
	util.CheckErr(err)

	// Check the json
	j, err := simple.NewJson(b)
	util.CheckErr(err)
	c.Assert(j.Get("resourceType").MustString(), check.Equals, "Bundle")
	c.Assert(j.Get("type").MustString(), check.Equals, "searchset")
	c.Assert(j.Get("total").MustInt(), check.Equals, 2)
	c.Assert(j.Get("link").MustArray(), check.HasLen, 1)
	c.Assert(j.Get("link").GetIndex(0).Get("relation").MustString(), check.Equals, "self")
	c.Assert(j.Get("link").GetIndex(0).Get("url").MustString(), check.Equals, "https://example.com/base/Condition?patient=4954037118555241963")
	c.Assert(j.Get("entry").MustArray(), check.HasLen, 2)
	c.Assert(j.Get("entry").GetIndex(0).GetPath("search", "mode").MustString(), check.Equals, "match")
	c.Assert(j.Get("entry").GetIndex(0).GetPath("search", "score").MustInt(), check.Equals, 1)
	r := j.Get("entry").GetIndex(0).GetPath("resource")
	c.Assert(r.GetPath("patient", "reference").MustString(), check.Equals, "https://example.com/base/Patient/4954037118555241963")
	c.Assert(r.GetPath("code", "coding").MustArray(), check.HasLen, 3)
	c.Assert(r.GetPath("code", "coding").GetIndex(0).Get("system").MustString(), check.Equals, "http://snomed.info/sct")
	c.Assert(r.GetPath("code", "coding").GetIndex(0).Get("code").MustString(), check.Equals, "10091002")
	c.Assert(r.GetPath("code", "coding").GetIndex(1).Get("system").MustString(), check.Equals, "http://hl7.org/fhir/sid/icd-9")
	c.Assert(r.GetPath("code", "coding").GetIndex(1).Get("code").MustString(), check.Equals, "428.0")
	c.Assert(r.GetPath("code", "coding").GetIndex(2).Get("system").MustString(), check.Equals, "http://hl7.org/fhir/sid/icd-10")
	c.Assert(r.GetPath("code", "coding").GetIndex(2).Get("code").MustString(), check.Equals, "I50.1")
	c.Assert(r.GetPath("code", "text").MustString(), check.Equals, "Heart failure")
	c.Assert(r.Get("onsetDateTime").MustString(), check.Equals, "2012-03-01T12:00:00Z")
	c.Assert(j.Get("entry").GetIndex(1).GetPath("search", "mode").MustString(), check.Equals, "match")
	c.Assert(j.Get("entry").GetIndex(1).GetPath("search", "score").MustInt(), check.Equals, 2)
	r = j.Get("entry").GetIndex(1).GetPath("resource")
	c.Assert(r.GetPath("patient", "reference").MustString(), check.Equals, "https://example.com/base/Patient/4954037118555241963")
	c.Assert(r.GetPath("code", "coding").MustArray(), check.HasLen, 1)
	c.Assert(r.GetPath("code", "coding").GetIndex(0).Get("system").MustString(), check.Equals, "http://snomed.info/sct")
	c.Assert(r.GetPath("code", "coding").GetIndex(0).Get("code").MustString(), check.Equals, "981000124106")
	c.Assert(r.GetPath("code", "text").MustString(), check.Equals, "Moderate left ventricular systolic dysfunction")
	c.Assert(r.Get("onsetDateTime").MustString(), check.Equals, "2012-03-01T12:05:00Z")
}

func (s *BundleSuite) TestUnmarshalJSON(c *check.C) {
	// Load a JSON Bundle
	bundle := LoadBundleFromFixture("../fixtures/search_response.json")

	// Check that it unmarshalled correctly
	c.Assert(bundle.Type, check.Equals, "searchset")
	c.Assert(*bundle.Total, check.Equals, uint32(2))
	c.Assert(bundle.Link, check.HasLen, 1)
	c.Assert(bundle.Link[0].Relation, check.Equals, "self")
	c.Assert(bundle.Link[0].Url, check.Equals, "https://example.com/base/Condition?patient=4954037118555241963")
	c.Assert(bundle.Entry, check.HasLen, 2)
	c.Assert(bundle.Entry[0].Search.Mode, check.Equals, "match")
	c.Assert(*bundle.Entry[0].Search.Score, check.Equals, float64(1))
	r, ok := bundle.Entry[0].Resource.(*Condition)
	c.Assert(ok, check.Equals, true)
	c.Assert(r.Patient.Reference, check.Equals, "https://example.com/base/Patient/4954037118555241963")
	c.Assert(r.Code.Coding, check.HasLen, 3)
	c.Assert(r.Code.MatchesCode("http://snomed.info/sct", "10091002"), check.Equals, true)
	c.Assert(r.Code.MatchesCode("http://hl7.org/fhir/sid/icd-9", "428.0"), check.Equals, true)
	c.Assert(r.Code.MatchesCode("http://hl7.org/fhir/sid/icd-10", "I50.1"), check.Equals, true)
	c.Assert(r.Code.Text, check.Equals, "Heart failure")
	c.Assert(r.OnsetDateTime.Precision, check.Equals, Precision(Timestamp))
	c.Assert(r.OnsetDateTime.Time.Unix(), check.Equals, int64(1330603200))
	c.Assert(bundle.Entry[1].Search.Mode, check.Equals, "match")
	c.Assert(*bundle.Entry[1].Search.Score, check.Equals, float64(2))
	r, ok = bundle.Entry[1].Resource.(*Condition)
	c.Assert(ok, check.Equals, true)
	c.Assert(r.Patient.Reference, check.Equals, "https://example.com/base/Patient/4954037118555241963")
	c.Assert(r.Code.Coding, check.HasLen, 1)
	c.Assert(r.Code.MatchesCode("http://snomed.info/sct", "981000124106"), check.Equals, true)
	c.Assert(r.Code.Text, check.Equals, "Moderate left ventricular systolic dysfunction")
	c.Assert(r.OnsetDateTime.Precision, check.Equals, Precision(Timestamp))
	c.Assert(r.OnsetDateTime.Time.Unix(), check.Equals, int64(1330603500))
}

func LoadBundleFromFixture(fileName string) *Bundle {
	data, err := os.Open(fileName)
	util.CheckErr(err)
	defer data.Close()

	decoder := json.NewDecoder(data)
	bundle := &Bundle{}
	err = decoder.Decode(bundle)
	util.CheckErr(err)
	return bundle
}
