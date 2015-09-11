package search

import (
	"time"

	. "gopkg.in/check.v1"
)

type SearchPTSuite struct {
	MDT *time.Location
}

var _ = Suite(&SearchPTSuite{})

func (s *SearchPTSuite) SetUpSuite(c *C) {
	s.MDT = time.FixedZone("MDT", -7*60*60)
}

/******************************************************************************
 * COMPOSITE
 ******************************************************************************/

var compositeParamInfo = SearchParamInfo{
	Name:       "foo",
	Type:       "composite",
	Composites: []string{"bar", "baz"},
}

func (s *SearchPTSuite) TestCompositeParam(c *C) {
	t := ParseCompositeParam("abc$123", compositeParamInfo)

	c.Assert(t.Name, Equals, "foo")
	c.Assert(t.Type, Equals, "composite")
	c.Assert(t.Paths, HasLen, 0)
	c.Assert(t.Composites, HasLen, 2)
	c.Assert(t.Composites, DeepEquals, []string{"bar", "baz"})
	c.Assert(t.CompositeValues, HasLen, 2)
	c.Assert(t.CompositeValues, DeepEquals, []string{"abc", "123"})
}

func (s *SearchPTSuite) TestCompositeParamWithTokenAndQuantity(c *C) {
	t := ParseCompositeParam("http://hl7.org/fhir/v2/0001|M$5.4|http://unitsofmeasure.org|mg", compositeParamInfo)

	c.Assert(t.Name, Equals, "foo")
	c.Assert(t.Type, Equals, "composite")
	c.Assert(t.Paths, HasLen, 0)
	c.Assert(t.Composites, HasLen, 2)
	c.Assert(t.Composites, DeepEquals, []string{"bar", "baz"})
	c.Assert(t.CompositeValues, HasLen, 2)
	c.Assert(t.CompositeValues, DeepEquals, []string{"http://hl7.org/fhir/v2/0001|M", "5.4|http://unitsofmeasure.org|mg"})
}

/******************************************************************************
 * DATE (Type)
 ******************************************************************************/

func (s *SearchPTSuite) TestDatesToMilliseconds(c *C) {

	d := ParseDate("2013-01-02T12:13:14.999-07:00")
	c.Assert(d.Precision, Equals, Millisecond)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, s.MDT).UnixNano())
	c.Assert(d.String(), Equals, "2013-01-02T12:13:14.999-07:00")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, s.MDT).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 15, 0, s.MDT).UnixNano())

	d = ParseDate("2013-01-02T12:13:14.999Z")
	c.Assert(d.Precision, Equals, Millisecond)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, time.UTC).UnixNano())
	c.Assert(d.String(), Equals, "2013-01-02T12:13:14.999+00:00")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, time.UTC).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 15, 0, time.UTC).UnixNano())

	d = ParseDate("2013-01-02T12:13:14.999")
	c.Assert(d.Precision, Equals, Millisecond)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, time.Local).UnixNano())
	c.Assert(d.String()[:23], Equals, "2013-01-02T12:13:14.999") // don't check the tz since it varies
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 15, 0, time.Local).UnixNano())

	// Test different levels of precision
	d = ParseDate("2013-01-02T12:13:14.9")
	c.Assert(d.Precision, Equals, Millisecond)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 900000000, time.Local).UnixNano())
	c.Assert(d.String()[:23], Equals, "2013-01-02T12:13:14.900") // don't check the tz since it varies
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 900000000, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 901000000, time.Local).UnixNano())

	d = ParseDate("2013-01-02T12:13:14.09")
	c.Assert(d.Precision, Equals, Millisecond)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 90000000, time.Local).UnixNano())
	c.Assert(d.String()[:23], Equals, "2013-01-02T12:13:14.090") // don't check the tz since it varies
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 90000000, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 91000000, time.Local).UnixNano())

	d = ParseDate("2013-01-02T12:13:14.009")
	c.Assert(d.Precision, Equals, Millisecond)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 9000000, time.Local).UnixNano())
	c.Assert(d.String()[:23], Equals, "2013-01-02T12:13:14.009") // don't check the tz since it varies
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 9000000, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 10000000, time.Local).UnixNano())

	d = ParseDate("2013-01-02T12:13:14.987654321")
	c.Assert(d.Precision, Equals, Millisecond)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 987000000, time.Local).UnixNano())
	c.Assert(d.String()[:23], Equals, "2013-01-02T12:13:14.987") // don't check the tz since it varies
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 987000000, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 988000000, time.Local).UnixNano())
}

func (s *SearchPTSuite) TestDatesToSeconds(c *C) {

	d := ParseDate("2013-01-02T12:13:14-07:00")
	c.Assert(d.Precision, Equals, Second)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, s.MDT).UnixNano())
	c.Assert(d.String(), Equals, "2013-01-02T12:13:14-07:00")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, s.MDT).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 15, 0, s.MDT).UnixNano())

	d = ParseDate("2013-01-02T12:13:14Z")
	c.Assert(d.Precision, Equals, Second)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.UTC).UnixNano())
	c.Assert(d.String(), Equals, "2013-01-02T12:13:14+00:00")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.UTC).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 15, 0, time.UTC).UnixNano())

	d = ParseDate("2013-01-02T12:13:14")
	c.Assert(d.Precision, Equals, Second)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.Local).UnixNano())
	c.Assert(d.String()[:19], Equals, "2013-01-02T12:13:14") // don't check the tz since it varies
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 15, 0, time.Local).UnixNano())
}

func (s *SearchPTSuite) TestDatesToMinutes(c *C) {

	d := ParseDate("2013-01-02T12:13-07:00")
	c.Assert(d.Precision, Equals, Minute)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 0, 0, s.MDT).UnixNano())
	c.Assert(d.String(), Equals, "2013-01-02T12:13-07:00")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 0, 0, s.MDT).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 14, 0, 0, s.MDT).UnixNano())

	d = ParseDate("2013-01-02T12:13Z")
	c.Assert(d.Precision, Equals, Minute)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 0, 0, time.UTC).UnixNano())
	c.Assert(d.String(), Equals, "2013-01-02T12:13+00:00")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 0, 0, time.UTC).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 14, 0, 0, time.UTC).UnixNano())

	d = ParseDate("2013-01-02T12:13")
	c.Assert(d.Precision, Equals, Minute)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 0, 0, time.Local).UnixNano())
	c.Assert(d.String()[:16], Equals, "2013-01-02T12:13") // don't check the tz since it varies
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 14, 0, 0, time.Local).UnixNano())
}

// NOTE: FHIR spec says that if hours are specified, minutes MUST be specified, so hours-only is invalid

func (s *SearchPTSuite) TestDatesToDays(c *C) {

	// Timezone should be ignored when no time components are included
	d := ParseDate("2013-01-02T-07:00")
	c.Assert(d.Precision, Equals, Day)
	c.Assert(d.Value.Unix(), Equals, time.Date(2013, time.January, 2, 0, 0, 0, 0, time.Local).Unix())
	c.Assert(d.String(), Equals, "2013-01-02")
	c.Assert(d.RangeLowIncl().Unix(), Equals, time.Date(2013, time.January, 2, 0, 0, 0, 0, time.Local).Unix())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 3, 0, 0, 0, 0, time.Local).UnixNano())

	d = ParseDate("2013-01-02Z")
	c.Assert(d.Precision, Equals, Day)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.String(), Equals, "2013-01-02")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 3, 0, 0, 0, 0, time.Local).UnixNano())

	d = ParseDate("2013-01-02")
	c.Assert(d.Precision, Equals, Day)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.String(), Equals, "2013-01-02")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 2, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.January, 3, 0, 0, 0, 0, time.Local).UnixNano())
}

func (s *SearchPTSuite) TestDatesToMonths(c *C) {

	// Timezone should be ignored when no time components are included
	d := ParseDate("2013-01T-07:00")
	c.Assert(d.Precision, Equals, Month)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.String(), Equals, "2013-01")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.February, 1, 0, 0, 0, 0, time.Local).UnixNano())

	d = ParseDate("2013-01Z")
	c.Assert(d.Precision, Equals, Month)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.String(), Equals, "2013-01")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.February, 1, 0, 0, 0, 0, time.Local).UnixNano())

	d = ParseDate("2013-01")
	c.Assert(d.Precision, Equals, Month)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.String(), Equals, "2013-01")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2013, time.February, 1, 0, 0, 0, 0, time.Local).UnixNano())
}

func (s *SearchPTSuite) TestDatesToYears(c *C) {

	// Timezone should be ignored when no time components are included
	d := ParseDate("2013T-07:00")
	c.Assert(d.Precision, Equals, Year)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.String(), Equals, "2013")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2014, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())

	d = ParseDate("2013Z")
	c.Assert(d.Precision, Equals, Year)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.String(), Equals, "2013")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2014, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())

	d = ParseDate("2013")
	c.Assert(d.Precision, Equals, Year)
	c.Assert(d.Value.UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.String(), Equals, "2013")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2014, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
}

func (s *SearchPTSuite) TestLeapAndNonLeapYears(c *C) {

	// Non-Leap Year
	d := ParseDate("1995-02-28")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(1995, time.February, 28, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(1995, time.March, 1, 0, 0, 0, 0, time.Local).UnixNano())

	// Leap Year
	d = ParseDate("1996-02-28")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(1996, time.February, 28, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(1996, time.February, 29, 0, 0, 0, 0, time.Local).UnixNano())

	// Centurial Non-Leap Year (divisible by 4, but centuries are not leap years unless they are divisible by 400)
	d = ParseDate("1900-02-28")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(1900, time.February, 28, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(1900, time.March, 1, 0, 0, 0, 0, time.Local).UnixNano())

	// Centurial Leap Year (divisible by 4, and a century, but also divisible by 400-- so it IS a leap year)
	d = ParseDate("2000-02-28")
	c.Assert(d.RangeLowIncl().UnixNano(), Equals, time.Date(2000, time.February, 28, 0, 0, 0, 0, time.Local).UnixNano())
	c.Assert(d.RangeHighExcl().UnixNano(), Equals, time.Date(2000, time.February, 29, 0, 0, 0, 0, time.Local).UnixNano())
}

/******************************************************************************
 * DATE (Param)
 ******************************************************************************/

var dateParamInfo = SearchParamInfo{
	Name:  "foo",
	Type:  "date",
	Paths: []SearchParamPath{SearchParamPath{Path: "bar", Type: "date"}},
}

func (s *SearchPTSuite) TestDateParamsToMilliseconds(c *C) {

	d := ParseDateParam("2013-01-02T12:13:14.999-07:00", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Millisecond)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, s.MDT).UnixNano())

	d = ParseDateParam("2013-01-02T12:13:14.999Z", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Millisecond)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, time.UTC).UnixNano())

	d = ParseDateParam("2013-01-02T12:13:14.999", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Millisecond)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, time.Local).UnixNano())
}

func (s *SearchPTSuite) TestDateParamsToSeconds(c *C) {

	d := ParseDateParam("2013-01-02T12:13:14-07:00", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Second)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, s.MDT).UnixNano())

	d = ParseDateParam("2013-01-02T12:13:14Z", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Second)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.UTC).UnixNano())

	d = ParseDateParam("2013-01-02T12:13:14", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Second)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.Local).UnixNano())
}

func (s *SearchPTSuite) TestDateParamsToMinutes(c *C) {

	d := ParseDateParam("2013-01-02T12:13-07:00", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Minute)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 0, 0, s.MDT).UnixNano())

	d = ParseDateParam("2013-01-02T12:13Z", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Minute)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 0, 0, time.UTC).UnixNano())

	d = ParseDateParam("2013-01-02T12:13", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Minute)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 0, 0, time.Local).UnixNano())
}

// NOTE: FHIR spec says that if hours are specified, minutes MUST be specified, so hours-only is invalid

func (s *SearchPTSuite) TestDateParamsToDays(c *C) {

	d := ParseDateParam("2013-01-02", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Day)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 0, 0, 0, 0, time.Local).UnixNano())
}

func (s *SearchPTSuite) TestDateParamsToMonths(c *C) {

	d := ParseDateParam("2013-01", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Month)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
}

func (s *SearchPTSuite) TestDateParamsToYears(c *C) {

	d := ParseDateParam("2013", dateParamInfo)
	c.Assert(d.Name, Equals, "foo")
	c.Assert(d.Type, Equals, "date")
	c.Assert(d.Paths, HasLen, 1)
	c.Assert(d.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Year)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 1, 0, 0, 0, 0, time.Local).UnixNano())
}

func (s *SearchPTSuite) TestDateParamPrefixes(c *C) {

	// Test prefixes
	d := ParseDateParam("2013-01-02T12:13:14Z", dateParamInfo)
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Second)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.UTC).UnixNano())

	d = ParseDateParam("eq2013-01-02T12:13:14Z", dateParamInfo)
	c.Assert(d.Prefix, Equals, EQ)
	c.Assert(d.Date.Precision, Equals, Second)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.UTC).UnixNano())

	d = ParseDateParam("ne2013-01-02T12:13:14Z", dateParamInfo)
	c.Assert(d.Prefix, Equals, NE)
	c.Assert(d.Date.Precision, Equals, Second)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.UTC).UnixNano())

	d = ParseDateParam("gt2013-01-02T12:13:14Z", dateParamInfo)
	c.Assert(d.Prefix, Equals, GT)
	c.Assert(d.Date.Precision, Equals, Second)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.UTC).UnixNano())

	d = ParseDateParam("lt2013-01-02T12:13:14Z", dateParamInfo)
	c.Assert(d.Prefix, Equals, LT)
	c.Assert(d.Date.Precision, Equals, Second)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.UTC).UnixNano())

	d = ParseDateParam("ge2013-01-02T12:13:14Z", dateParamInfo)
	c.Assert(d.Prefix, Equals, GE)
	c.Assert(d.Date.Precision, Equals, Second)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.UTC).UnixNano())

	d = ParseDateParam("le2013-01-02T12:13:14Z", dateParamInfo)
	c.Assert(d.Prefix, Equals, LE)
	c.Assert(d.Date.Precision, Equals, Second)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.UTC).UnixNano())

	d = ParseDateParam("ap2013-01-02T12:13:14Z", dateParamInfo)
	c.Assert(d.Prefix, Equals, AP)
	c.Assert(d.Date.Precision, Equals, Second)
	c.Assert(d.Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 0, time.UTC).UnixNano())
}

/******************************************************************************
 * NUMBER (Type)
 ******************************************************************************/

func (s *SearchPTSuite) TestNumbersThatAreInts(c *C) {
	n := ParseNumber("100")

	c.Assert(n.Precision, Equals, 0)
	c.Assert(n.Value.RatString(), Equals, "100")
	c.Assert(n.String(), Equals, "100")
	c.Assert(n.RangeLowIncl().RatString(), Equals, "199/2")
	c.Assert(n.RangeHighExcl().RatString(), Equals, "201/2")
}

func (s *SearchPTSuite) TestNumbersThatAreNegativeInts(c *C) {
	n := ParseNumber("-100")

	c.Assert(n.Precision, Equals, 0)
	c.Assert(n.Value.RatString(), Equals, "-100")
	c.Assert(n.String(), Equals, "-100")
	c.Assert(n.RangeLowIncl().RatString(), Equals, "-201/2")
	c.Assert(n.RangeHighExcl().RatString(), Equals, "-199/2")
}

func (s *SearchPTSuite) TestNumbersThatAreDecimals(c *C) {
	n := ParseNumber("0.12345678900000000000")

	c.Assert(n.Precision, Equals, 20)
	c.Assert(n.Value.FloatString(22), Equals, "0.1234567890000000000000")
	c.Assert(n.String(), Equals, "0.12345678900000000000")
	c.Assert(n.RangeLowIncl().FloatString(22), Equals, "0.1234567889999999999950")
	c.Assert(n.RangeHighExcl().FloatString(22), Equals, "0.1234567890000000000050")
}

func (s *SearchPTSuite) TestNumbersThatAreNegativeDecimals(c *C) {
	n := ParseNumber("-0.12345678900000000000")

	c.Assert(n.Precision, Equals, 20)
	c.Assert(n.Value.FloatString(22), Equals, "-0.1234567890000000000000")
	c.Assert(n.String(), Equals, "-0.12345678900000000000")
	c.Assert(n.RangeLowIncl().FloatString(22), Equals, "-0.1234567890000000000050")
	c.Assert(n.RangeHighExcl().FloatString(22), Equals, "-0.1234567889999999999950")
}

/******************************************************************************
 * NUMBER (Param)
 ******************************************************************************/

var numberParamInfo = SearchParamInfo{
	Name:  "foo",
	Type:  "number",
	Paths: []SearchParamPath{SearchParamPath{Path: "bar", Type: "number"}},
}

func (s *SearchPTSuite) TestNumberParamsThatAreInts(c *C) {
	n := ParseNumberParam("100", numberParamInfo)

	c.Assert(n.Name, Equals, "foo")
	c.Assert(n.Type, Equals, "number")
	c.Assert(n.Paths, HasLen, 1)
	c.Assert(n.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "number"})
	c.Assert(n.Prefix, Equals, EQ)
	c.Assert(n.Number.String(), Equals, "100")
	f, _ := n.Number.Value.Float64()
	c.Assert(f, Equals, float64(100))
	f, _ = n.Number.RangeLowIncl().Float64()
	c.Assert(f, Equals, float64(99.5))
	f, _ = n.Number.RangeHighExcl().Float64()
	c.Assert(f, Equals, float64(100.5))
}

func (s *SearchPTSuite) TestNumberParamsThatAreNegativeInts(c *C) {
	n := ParseNumberParam("-100", numberParamInfo)

	c.Assert(n.Name, Equals, "foo")
	c.Assert(n.Type, Equals, "number")
	c.Assert(n.Paths, HasLen, 1)
	c.Assert(n.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "number"})
	c.Assert(n.Prefix, Equals, EQ)
	c.Assert(n.Number.String(), Equals, "-100")
}

func (s *SearchPTSuite) TestNumberParamsThatAreDecimals(c *C) {
	n := ParseNumberParam("100.00", numberParamInfo)

	c.Assert(n.Name, Equals, "foo")
	c.Assert(n.Type, Equals, "number")
	c.Assert(n.Paths, HasLen, 1)
	c.Assert(n.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "number"})
	c.Assert(n.Prefix, Equals, EQ)
	c.Assert(n.Number.String(), Equals, "100.00")
	f, _ := n.Number.Value.Float64()
	c.Assert(f, Equals, float64(100))
	f, _ = n.Number.RangeLowIncl().Float64()
	c.Assert(f, Equals, float64(99.995))
	f, _ = n.Number.RangeHighExcl().Float64()
	c.Assert(f, Equals, float64(100.005))
}

func (s *SearchPTSuite) TestNumberParamsThatAreNegativeDecimals(c *C) {
	n := ParseNumberParam("-100.00", numberParamInfo)

	c.Assert(n.Name, Equals, "foo")
	c.Assert(n.Type, Equals, "number")
	c.Assert(n.Paths, HasLen, 1)
	c.Assert(n.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "number"})
	c.Assert(n.Prefix, Equals, EQ)
	c.Assert(n.Number.String(), Equals, "-100.00")
}

func (s *SearchPTSuite) TestNumberParamPrefixes(c *C) {
	n := ParseNumberParam("100", numberParamInfo)
	c.Assert(n.Prefix, Equals, EQ)
	c.Assert(n.Number.String(), Equals, "100")

	n = ParseNumberParam("eq100", numberParamInfo)
	c.Assert(n.Prefix, Equals, EQ)
	c.Assert(n.Number.String(), Equals, "100")

	n = ParseNumberParam("ne100", numberParamInfo)
	c.Assert(n.Prefix, Equals, NE)
	c.Assert(n.Number.String(), Equals, "100")

	n = ParseNumberParam("gt100", numberParamInfo)
	c.Assert(n.Prefix, Equals, GT)
	c.Assert(n.Number.String(), Equals, "100")

	n = ParseNumberParam("lt100", numberParamInfo)
	c.Assert(n.Prefix, Equals, LT)
	c.Assert(n.Number.String(), Equals, "100")

	n = ParseNumberParam("ge100", numberParamInfo)
	c.Assert(n.Prefix, Equals, GE)
	c.Assert(n.Number.String(), Equals, "100")

	n = ParseNumberParam("le100", numberParamInfo)
	c.Assert(n.Prefix, Equals, LE)
	c.Assert(n.Number.String(), Equals, "100")

	n = ParseNumberParam("ap100", numberParamInfo)
	c.Assert(n.Prefix, Equals, AP)
	c.Assert(n.Number.String(), Equals, "100")
}

/******************************************************************************
 * QUANTITY
 ******************************************************************************/

var quantityParamInfo = SearchParamInfo{
	Name:  "foo",
	Type:  "quantity",
	Paths: []SearchParamPath{SearchParamPath{Path: "bar", Type: "quantity"}},
}

func (s *SearchPTSuite) TestQuantitiesWithSystemsAndUnits(c *C) {
	q := ParseQuantityParam("5.4|http://unitsofmeasure.org|mg", quantityParamInfo)

	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, EQ)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "http://unitsofmeasure.org")
	c.Assert(q.Code, Equals, "mg")
}

func (s *SearchPTSuite) TestQuantitiesWithOnlyUnits(c *C) {
	q := ParseQuantityParam("5.4||mg", quantityParamInfo)

	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, EQ)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "")
	c.Assert(q.Code, Equals, "mg")
}

func (s *SearchPTSuite) TestQuantitiesWithNoUnits(c *C) {
	q := ParseQuantityParam("5.4", quantityParamInfo)

	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, EQ)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "")
	c.Assert(q.Code, Equals, "")
}

func (s *SearchPTSuite) TestNegativeQuantities(c *C) {
	q := ParseQuantityParam("-10|http://unitsofmeasure.org|mg", quantityParamInfo)

	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, EQ)
	c.Assert(q.Number.String(), Equals, "-10")
	c.Assert(q.System, Equals, "http://unitsofmeasure.org")
	c.Assert(q.Code, Equals, "mg")
}

func (s *SearchPTSuite) TestQuantitiesWithEscapedPipesAndSlashes(c *C) {
	q := ParseQuantityParam("5.4|foo\\|bar|foo\\\\\\|baz", quantityParamInfo)

	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, EQ)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "foo|bar")
	c.Assert(q.Code, Equals, "foo\\|baz")
}

func (s *SearchPTSuite) TestQuantityPrefixes(c *C) {
	q := ParseQuantityParam("5.4|http://unitsofmeasure.org|mg", quantityParamInfo)
	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, EQ)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "http://unitsofmeasure.org")
	c.Assert(q.Code, Equals, "mg")

	q = ParseQuantityParam("eq5.4|http://unitsofmeasure.org|mg", quantityParamInfo)
	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, EQ)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "http://unitsofmeasure.org")
	c.Assert(q.Code, Equals, "mg")

	q = ParseQuantityParam("ne5.4|http://unitsofmeasure.org|mg", quantityParamInfo)
	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, NE)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "http://unitsofmeasure.org")
	c.Assert(q.Code, Equals, "mg")

	q = ParseQuantityParam("gt5.4|http://unitsofmeasure.org|mg", quantityParamInfo)
	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, GT)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "http://unitsofmeasure.org")
	c.Assert(q.Code, Equals, "mg")

	q = ParseQuantityParam("lt5.4|http://unitsofmeasure.org|mg", quantityParamInfo)
	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, LT)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "http://unitsofmeasure.org")
	c.Assert(q.Code, Equals, "mg")

	q = ParseQuantityParam("ge5.4|http://unitsofmeasure.org|mg", quantityParamInfo)
	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, GE)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "http://unitsofmeasure.org")
	c.Assert(q.Code, Equals, "mg")

	q = ParseQuantityParam("le5.4|http://unitsofmeasure.org|mg", quantityParamInfo)
	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, LE)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "http://unitsofmeasure.org")
	c.Assert(q.Code, Equals, "mg")

	q = ParseQuantityParam("ap5.4|http://unitsofmeasure.org|mg", quantityParamInfo)
	c.Assert(q.Name, Equals, "foo")
	c.Assert(q.Type, Equals, "quantity")
	c.Assert(q.Paths, HasLen, 1)
	c.Assert(q.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "quantity"})
	c.Assert(q.Prefix, Equals, AP)
	c.Assert(q.Number.String(), Equals, "5.4")
	c.Assert(q.System, Equals, "http://unitsofmeasure.org")
	c.Assert(q.Code, Equals, "mg")
}

/******************************************************************************
 * REFERENCE
 ******************************************************************************/

var referenceParamInfo = SearchParamInfo{
	Name:    "foo",
	Type:    "reference",
	Paths:   []SearchParamPath{SearchParamPath{Path: "bar", Type: "reference"}},
	Targets: []string{"Patient"},
}

func (s *SearchPTSuite) TestReferenceID(c *C) {
	r := ParseReferenceParam("23", referenceParamInfo)

	c.Assert(r.Name, Equals, "foo")
	c.Assert(r.Type, Equals, "reference")
	c.Assert(r.Paths, HasLen, 1)
	c.Assert(r.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "reference"})

	c.Assert(r.Reference, FitsTypeOf, LocalReference{})
	lRef := r.Reference.(LocalReference)
	c.Assert(lRef.ID, Equals, "23")
	c.Assert(lRef.Type, Equals, "Patient")
}

func (s *SearchPTSuite) TestReferenceIDWithModifier(c *C) {
	modInfo := referenceParamInfo
	modInfo.Modifier = "Patient"
	r := ParseReferenceParam("23", modInfo)

	c.Assert(r.Name, Equals, "foo")
	c.Assert(r.Type, Equals, "reference")
	c.Assert(r.Paths, HasLen, 1)
	c.Assert(r.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "reference"})
	c.Assert(r.Modifier, Equals, "Patient")

	c.Assert(r.Reference, FitsTypeOf, LocalReference{})
	lRef := r.Reference.(LocalReference)
	c.Assert(lRef.ID, Equals, "23")
	c.Assert(lRef.Type, Equals, "Patient")
}

func (s *SearchPTSuite) TestReferenceIDWithMismatchedModifier(c *C) {
	modInfo := referenceParamInfo
	modInfo.Modifier = "Condition"
	c.Assert(func() { ParseReferenceParam("23", modInfo) }, Panics, InvalidSearchError("target is Patient, but type is Condition"))
}

func (s *SearchPTSuite) TestReferenceTypeAndId(c *C) {
	r := ParseReferenceParam("Patient/23", referenceParamInfo)

	c.Assert(r.Name, Equals, "foo")
	c.Assert(r.Type, Equals, "reference")
	c.Assert(r.Paths, HasLen, 1)
	c.Assert(r.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "reference"})

	c.Assert(r.Reference, FitsTypeOf, LocalReference{})
	lRef := r.Reference.(LocalReference)
	c.Assert(lRef.ID, Equals, "23")
	c.Assert(lRef.Type, Equals, "Patient")
}

func (s *SearchPTSuite) TestReferenceTypeAndIDWithMismatchedType(c *C) {
	modInfo := referenceParamInfo
	c.Assert(func() { ParseReferenceParam("Condition/23", modInfo) }, Panics, InvalidSearchError("target is Patient, but type is Condition"))
}

func (s *SearchPTSuite) TestReferenceTypeAndIdWithModifier(c *C) {
	modInfo := referenceParamInfo
	modInfo.Modifier = "Patient"
	r := ParseReferenceParam("Patient/23", modInfo)

	c.Assert(r.Name, Equals, "foo")
	c.Assert(r.Type, Equals, "reference")
	c.Assert(r.Paths, HasLen, 1)
	c.Assert(r.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "reference"})
	c.Assert(r.Modifier, Equals, "Patient")

	c.Assert(r.Reference, FitsTypeOf, LocalReference{})
	lRef := r.Reference.(LocalReference)
	c.Assert(lRef.ID, Equals, "23")
	c.Assert(lRef.Type, Equals, "Patient")
}

func (s *SearchPTSuite) TestReferenceTypeAndIdWithMismatchedModifier(c *C) {
	modInfo := referenceParamInfo
	modInfo.Modifier = "Condition"
	c.Assert(func() { ParseReferenceParam("Patient/23", modInfo) }, Panics, InvalidSearchError("reference modifier is Condition, but type is Patient"))
}

func (s *SearchPTSuite) TestReferenceAbsoluteURL(c *C) {
	r := ParseReferenceParam("http://acme.org/fhir/Patient/23", referenceParamInfo)

	c.Assert(r.Name, Equals, "foo")
	c.Assert(r.Type, Equals, "reference")
	c.Assert(r.Paths, HasLen, 1)
	c.Assert(r.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "reference"})

	c.Assert(r.Reference, FitsTypeOf, ExternalReference{})
	eRef := r.Reference.(ExternalReference)
	c.Assert(eRef.URL, Equals, "http://acme.org/fhir/Patient/23")
	c.Assert(eRef.Type, Equals, "Patient")
}

func (s *SearchPTSuite) TestReferenceAbsoluteURLWithMismatchedType(c *C) {
	modInfo := referenceParamInfo
	c.Assert(func() { ParseReferenceParam("http://acme.org/fhir/Condition/23", modInfo) }, Panics, InvalidSearchError("target is Patient, but type is Condition"))
}

func (s *SearchPTSuite) TestReferenceAbsoluteURLWithModifier(c *C) {
	modInfo := referenceParamInfo
	modInfo.Modifier = "Patient"
	r := ParseReferenceParam("http://acme.org/fhir/Patient/23", modInfo)

	c.Assert(r.Name, Equals, "foo")
	c.Assert(r.Type, Equals, "reference")
	c.Assert(r.Paths, HasLen, 1)
	c.Assert(r.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "reference"})
	c.Assert(r.Modifier, Equals, "Patient")

	c.Assert(r.Reference, FitsTypeOf, ExternalReference{})
	eRef := r.Reference.(ExternalReference)
	c.Assert(eRef.URL, Equals, "http://acme.org/fhir/Patient/23")
	c.Assert(eRef.Type, Equals, "Patient")
}

func (s *SearchPTSuite) TestReferenceAbsoluteURLWithMismatchedModifier(c *C) {
	modInfo := referenceParamInfo
	modInfo.Modifier = "Condition"
	c.Assert(func() { ParseReferenceParam("http://acme.org/fhir/Patient/23", modInfo) }, Panics, InvalidSearchError("reference modifier is Condition, but type is Patient"))
}

func (s *SearchPTSuite) TestReferenceChainedQuery(c *C) {
	modInfo := referenceParamInfo
	modInfo.Postfix = "name"
	r := ParseReferenceParam("Peter", modInfo)

	c.Assert(r.Name, Equals, "foo")
	c.Assert(r.Type, Equals, "reference")
	c.Assert(r.Paths, HasLen, 1)
	c.Assert(r.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "reference"})

	c.Assert(r.Reference, FitsTypeOf, ChainedQueryReference{})
	qRef := r.Reference.(ChainedQueryReference)
	c.Assert(qRef.ChainedQuery, DeepEquals, Query{Resource: "Patient", Query: "name=Peter"})
	c.Assert(qRef.Type, Equals, "Patient")
}

func (s *SearchPTSuite) TestReferenceChainedQueryWithModifier(c *C) {
	modInfo := referenceParamInfo
	modInfo.Modifier = "Patient"
	modInfo.Postfix = "name"
	r := ParseReferenceParam("Peter", modInfo)

	c.Assert(r.Name, Equals, "foo")
	c.Assert(r.Type, Equals, "reference")
	c.Assert(r.Paths, HasLen, 1)
	c.Assert(r.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "reference"})

	c.Assert(r.Reference, FitsTypeOf, ChainedQueryReference{})
	qRef := r.Reference.(ChainedQueryReference)
	c.Assert(qRef.ChainedQuery, DeepEquals, Query{Resource: "Patient", Query: "name=Peter"})
	c.Assert(qRef.Type, Equals, "Patient")
}

func (s *SearchPTSuite) TestReferenceChainedQueryWithMismatchedModifier(c *C) {
	modInfo := referenceParamInfo
	modInfo.Modifier = "Condition"
	modInfo.Postfix = "name"
	c.Assert(func() { ParseReferenceParam("Peter", modInfo) }, Panics, InvalidSearchError("target is Patient, but type is Condition"))
}

/******************************************************************************
 * STRING
 ******************************************************************************/

var stringParamInfo = SearchParamInfo{
	Name:  "foo",
	Type:  "string",
	Paths: []SearchParamPath{SearchParamPath{Path: "bar", Type: "string"}},
}

func (s *SearchPTSuite) TestStringParam(c *C) {
	st := ParseStringParam("Hello World", stringParamInfo)

	c.Assert(st.Name, Equals, "foo")
	c.Assert(st.Type, Equals, "string")
	c.Assert(st.Paths, HasLen, 1)
	c.Assert(st.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "string"})
	c.Assert(st.String, Equals, "Hello World")
}

/******************************************************************************
 * TOKEN
 ******************************************************************************/

var tokenParamInfo = SearchParamInfo{
	Name:  "foo",
	Type:  "token",
	Paths: []SearchParamPath{SearchParamPath{Path: "bar", Type: "CodeableConcept"}},
}

func (s *SearchPTSuite) TestTokenParamCode(c *C) {
	t := ParseTokenParam("M", tokenParamInfo)

	c.Assert(t.Name, Equals, "foo")
	c.Assert(t.Type, Equals, "token")
	c.Assert(t.Paths, HasLen, 1)
	c.Assert(t.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "CodeableConcept"})
	c.Assert(t.AnySystem, Equals, true)
	c.Assert(t.Code, Equals, "M")
	c.Assert(t.System, Equals, "")
}

func (s *SearchPTSuite) TestTokenParamSystemAndCode(c *C) {
	t := ParseTokenParam("http://hl7.org/fhir/v2/0001|M", tokenParamInfo)

	c.Assert(t.Name, Equals, "foo")
	c.Assert(t.Type, Equals, "token")
	c.Assert(t.Paths, HasLen, 1)
	c.Assert(t.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "CodeableConcept"})
	c.Assert(t.AnySystem, Equals, false)
	c.Assert(t.Code, Equals, "M")
	c.Assert(t.System, Equals, "http://hl7.org/fhir/v2/0001")
}

func (s *SearchPTSuite) TestTokenParamSystemlessCode(c *C) {
	t := ParseTokenParam("|M", tokenParamInfo)

	c.Assert(t.Name, Equals, "foo")
	c.Assert(t.Type, Equals, "token")
	c.Assert(t.Paths, HasLen, 1)
	c.Assert(t.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "CodeableConcept"})
	c.Assert(t.AnySystem, Equals, false)
	c.Assert(t.Code, Equals, "M")
	c.Assert(t.System, Equals, "")
}

func (s *SearchPTSuite) TestTokenParamsWithEscapedPipesAndSlashes(c *C) {
	t := ParseTokenParam("foo\\|bar", tokenParamInfo)

	c.Assert(t.Name, Equals, "foo")
	c.Assert(t.Type, Equals, "token")
	c.Assert(t.Paths, HasLen, 1)
	c.Assert(t.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "CodeableConcept"})
	c.Assert(t.AnySystem, Equals, true)
	c.Assert(t.Code, Equals, "foo|bar")
	c.Assert(t.System, Equals, "")

	t = ParseTokenParam("foo\\|bar|foo\\\\\\|baz", tokenParamInfo)

	c.Assert(t.Name, Equals, "foo")
	c.Assert(t.Type, Equals, "token")
	c.Assert(t.Paths, HasLen, 1)
	c.Assert(t.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "CodeableConcept"})
	c.Assert(t.AnySystem, Equals, false)
	c.Assert(t.Code, Equals, "foo\\|baz")
	c.Assert(t.System, Equals, "foo|bar")
}

/******************************************************************************
 * URI
 ******************************************************************************/

var uriParamInfo = SearchParamInfo{
	Name:  "foo",
	Type:  "uri",
	Paths: []SearchParamPath{SearchParamPath{Path: "bar", Type: "uri"}},
}

func (s *SearchPTSuite) TestURIParam(c *C) {
	u := ParseURIParam("http://acme.org/fhir/ValueSet/123", uriParamInfo)

	c.Assert(u.Name, Equals, "foo")
	c.Assert(u.Type, Equals, "uri")
	c.Assert(u.Paths, HasLen, 1)
	c.Assert(u.Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "uri"})
	c.Assert(u.URI, Equals, "http://acme.org/fhir/ValueSet/123")
}

/******************************************************************************
 * OR
 ******************************************************************************/

// Or dates
func (s *SearchPTSuite) TestOrDateParams(c *C) {
	o := ParseOrParam([]string{"2013-01-02T12:13:14.999-07:00", "2013-01-02T12:13:14.999Z", "2013-01-02T12:13:14.999"}, dateParamInfo)
	c.Assert(o.Name, Equals, "foo")
	c.Assert(o.Type, Equals, "or")
	c.Assert(o.Paths, HasLen, 0)
	c.Assert(o.Composites, HasLen, 0)
	c.Assert(o.Items, HasLen, 3)

	for i := 0; i < 3; i++ {
		c.Assert(o.Items[i], FitsTypeOf, &DateParam{})
		c.Assert(o.Items[i].getInfo().Name, Equals, "foo")
		c.Assert(o.Items[i].getInfo().Type, Equals, "date")
		c.Assert(o.Items[i].getInfo().Paths, HasLen, 1)
		c.Assert(o.Items[i].getInfo().Paths[0], DeepEquals, SearchParamPath{Path: "bar", Type: "date"})
		c.Assert(o.Items[i].(*DateParam).Prefix, Equals, EQ)
		c.Assert(o.Items[i].(*DateParam).Date.Precision, Equals, Millisecond)
		switch i {
		case 0:
			c.Assert(o.Items[i].(*DateParam).Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, s.MDT).UnixNano())
		case 1:
			c.Assert(o.Items[i].(*DateParam).Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, time.UTC).UnixNano())
		case 2:
			c.Assert(o.Items[i].(*DateParam).Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, time.Local).UnixNano())
		}
	}
}

func (s *SearchPTSuite) TestOrQueryIsParsedCorrectly(c *C) {
	q := Query{"Condition", "onset=2013-01-02T12:13:14.999-07:00,2013-01-02T12:13:14.999Z,2013-01-02T12:13:14.999&code=foo|bar"}
	p := q.Params()

	c.Assert(p, HasLen, 2)

	onset, code := 0, 1
	if _, ok := p[0].(*TokenParam); ok {
		onset, code = 1, 0
	}

	c.Assert(p[onset], FitsTypeOf, &OrParam{})
	c.Assert(p[onset].getInfo().Name, Equals, "onset")
	c.Assert(p[onset].getInfo().Type, Equals, "or")
	c.Assert(p[onset].getInfo().Paths, HasLen, 0)
	c.Assert(p[onset].getInfo().Composites, HasLen, 0)
	c.Assert(p[onset].(*OrParam).Items, HasLen, 3)
	for i := 0; i < 3; i++ {
		c.Assert(p[onset].(*OrParam).Items[i], FitsTypeOf, &DateParam{})
		c.Assert(p[onset].(*OrParam).Items[i].getInfo().Name, Equals, "onset")
		c.Assert(p[onset].(*OrParam).Items[i].getInfo().Type, Equals, "date")
		c.Assert(p[onset].(*OrParam).Items[i].getInfo().Paths, HasLen, 2)
		c.Assert(p[onset].(*OrParam).Items[i].(*DateParam).Prefix, Equals, EQ)
		c.Assert(p[onset].(*OrParam).Items[i].(*DateParam).Date.Precision, Equals, Millisecond)
		switch i {
		case 0:
			c.Assert(p[onset].(*OrParam).Items[i].(*DateParam).Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, s.MDT).UnixNano())
		case 1:
			c.Assert(p[onset].(*OrParam).Items[i].(*DateParam).Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, time.UTC).UnixNano())
		case 2:
			c.Assert(p[onset].(*OrParam).Items[i].(*DateParam).Date.Value.UnixNano(), Equals, time.Date(2013, time.January, 2, 12, 13, 14, 999000000, time.Local).UnixNano())
		}
	}

	c.Assert(p[code], FitsTypeOf, &TokenParam{})
	c.Assert(p[code].getInfo().Name, Equals, "code")
	c.Assert(p[code].getInfo().Type, Equals, "token")
	c.Assert(p[code].getInfo().Paths, HasLen, 1)
	c.Assert(p[code].getInfo().Composites, HasLen, 0)
	c.Assert(p[code].(*TokenParam).System, Equals, "foo")
	c.Assert(p[code].(*TokenParam).Code, Equals, "bar")
	c.Assert(p[code].(*TokenParam).AnySystem, Equals, false)
}

/******************************************************************************
 * PREFIX
 ******************************************************************************/

func (s *SearchPTSuite) TestPrefixes(c *C) {
	x, y := ExtractPrefixAndValue("eq10")
	c.Assert(x, Equals, EQ)
	c.Assert(y, Equals, "10")

	x, y = ExtractPrefixAndValue("ne10")
	c.Assert(x, Equals, NE)
	c.Assert(y, Equals, "10")

	x, y = ExtractPrefixAndValue("gt10")
	c.Assert(x, Equals, GT)
	c.Assert(y, Equals, "10")

	x, y = ExtractPrefixAndValue("lt10")
	c.Assert(x, Equals, LT)
	c.Assert(y, Equals, "10")

	x, y = ExtractPrefixAndValue("ge10")
	c.Assert(x, Equals, GE)
	c.Assert(y, Equals, "10")

	x, y = ExtractPrefixAndValue("le10")
	c.Assert(x, Equals, LE)
	c.Assert(y, Equals, "10")

	x, y = ExtractPrefixAndValue("ap10")
	c.Assert(x, Equals, AP)
	c.Assert(y, Equals, "10")
}

func (s *SearchPTSuite) TestPrefixDefault(c *C) {
	x, y := ExtractPrefixAndValue("10")
	c.Assert(x, Equals, EQ)
	c.Assert(y, Equals, "10")
}
