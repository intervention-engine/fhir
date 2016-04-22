package search

import (
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
)

type RegistrySuite struct{}

var _ = Suite(&RegistrySuite{})

func (s *RegistrySuite) TestRegisterAndLookupParameterInfo(c *C) {
	info := SearchParamInfo{
		Resource: "Blah",
		Name:     "foo",
		Type:     "test",
	}

	GlobalRegistry().RegisterParameterInfo(info)
	obtained, err := GlobalRegistry().LookupParameterInfo("Blah", "foo")
	util.CheckErr(err)
	c.Assert(obtained, DeepEquals, info)
}

func (s *RegistrySuite) TestLookupNonExistingParameterInfo(c *C) {
	obtained, err := GlobalRegistry().LookupParameterInfo("Foo", "Bar")
	c.Assert(err, Not(IsNil))
	c.Assert(obtained, DeepEquals, SearchParamInfo{}) // Zero Object
}

func (s *RegistrySuite) TestRegisterAndLookupParameterParser(c *C) {
	info := SearchParamInfo{
		Resource: "Blah",
		Name:     "foo",
		Type:     "test",
	}
	parser := func(info SearchParamInfo, data SearchParamData) (SearchParam, error) {
		return ParseStringParam("bar", info), nil
	}

	GlobalRegistry().RegisterParameterParser("test", parser)
	obtained, err := GlobalRegistry().LookupParameterParser("test")
	util.CheckErr(err)
	param, err := obtained(info, SearchParamData{Value: "bar"})
	util.CheckErr(err)
	c.Assert(param.getInfo(), DeepEquals, info)
	p, v := param.getQueryParamAndValue()
	c.Assert(p, Equals, "foo")
	c.Assert(v, Equals, "bar")
}

func (s *RegistrySuite) TestLookupNonExistingParameterParser(c *C) {
	obtained, err := GlobalRegistry().LookupParameterParser("nope")
	c.Assert(err, Not(IsNil))
	c.Assert(obtained, IsNil)
}
