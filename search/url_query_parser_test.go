package search

import (
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
)

type URLQueryParserSuite struct{}

var _ = Suite(&URLQueryParserSuite{})

func (s *URLQueryParserSuite) TestParseQuery(c *C) {
	p, err := ParseQuery("abc=def&zyx=wvu&lmn=opq&ghi=jkl")
	util.CheckErr(err)
	all := p.All()
	c.Assert(all, HasLen, 4)
	c.Assert(all[0], DeepEquals, URLQueryParameter{Key: "abc", Value: "def"})
	c.Assert(all[1], DeepEquals, URLQueryParameter{Key: "zyx", Value: "wvu"})
	c.Assert(all[2], DeepEquals, URLQueryParameter{Key: "lmn", Value: "opq"})
	c.Assert(all[3], DeepEquals, URLQueryParameter{Key: "ghi", Value: "jkl"})
}

func (s *URLQueryParserSuite) TestParseQueryDuplicateKeys(c *C) {
	p, err := ParseQuery("abc=def&zyx=wvu&abc=123&lmn=opq&ghi=jkl&abc=xyz")
	util.CheckErr(err)
	all := p.All()
	c.Assert(all, HasLen, 6)
	c.Assert(all[0], DeepEquals, URLQueryParameter{Key: "abc", Value: "def"})
	c.Assert(all[1], DeepEquals, URLQueryParameter{Key: "zyx", Value: "wvu"})
	c.Assert(all[2], DeepEquals, URLQueryParameter{Key: "abc", Value: "123"})
	c.Assert(all[3], DeepEquals, URLQueryParameter{Key: "lmn", Value: "opq"})
	c.Assert(all[4], DeepEquals, URLQueryParameter{Key: "ghi", Value: "jkl"})
	c.Assert(all[5], DeepEquals, URLQueryParameter{Key: "abc", Value: "xyz"})
}

func (s *URLQueryParserSuite) TestParseQueryWithEscapedChars(c *C) {
	p, err := ParseQuery("foo%3Abar=foo%26baz")
	util.CheckErr(err)
	all := p.All()
	c.Assert(all, HasLen, 1)
	c.Assert(all[0], DeepEquals, URLQueryParameter{Key: "foo:bar", Value: "foo&baz"})
	c.Assert(p.Get("foo:bar"), Equals, "foo&baz")
	c.Assert(p.GetMulti("foo:bar"), DeepEquals, []string{"foo&baz"})
	values := p.Values()
	c.Assert(values, HasLen, 1)
	c.Assert(values.Get("foo:bar"), Equals, "foo&baz")
}

func (s *URLQueryParserSuite) TestAddSingleParam(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	all := p.All()
	c.Assert(all, HasLen, 1)
	c.Assert(all[0], DeepEquals, URLQueryParameter{Key: "foo", Value: "bar"})
}

func (s *URLQueryParserSuite) TestAddMultipleParams(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Add("foo2", "bar2")
	p.Add("foo3", "bar3")
	all := p.All()
	c.Assert(all, HasLen, 3)
	c.Assert(all[0], DeepEquals, URLQueryParameter{Key: "foo", Value: "bar"})
	c.Assert(all[1], DeepEquals, URLQueryParameter{Key: "foo2", Value: "bar2"})
	c.Assert(all[2], DeepEquals, URLQueryParameter{Key: "foo3", Value: "bar3"})
}

func (s *URLQueryParserSuite) TestAddDuplicateKeyParams(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Add("foo", "baz")
	p.Add("foo", "barz")
	all := p.All()
	c.Assert(all, HasLen, 3)
	c.Assert(all[0], DeepEquals, URLQueryParameter{Key: "foo", Value: "bar"})
	c.Assert(all[1], DeepEquals, URLQueryParameter{Key: "foo", Value: "baz"})
	c.Assert(all[2], DeepEquals, URLQueryParameter{Key: "foo", Value: "barz"})
}

func (s *URLQueryParserSuite) TestSetNonExistingParam(c *C) {
	p := URLQueryParameters{}
	p.Set("foo", "bar")
	all := p.All()
	c.Assert(all, HasLen, 1)
	c.Assert(all[0], DeepEquals, URLQueryParameter{Key: "foo", Value: "bar"})
}

func (s *URLQueryParserSuite) TestSetExistingParam(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Set("foo", "baz")
	all := p.All()
	c.Assert(all, HasLen, 1)
	c.Assert(all[0], DeepEquals, URLQueryParameter{Key: "foo", Value: "baz"})
}

func (s *URLQueryParserSuite) TestSetExistingMultipleParams(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Add("foo2", "bar2")
	p.Add("foo", "baz")
	p.Add("foo3", "bar3")
	p.Set("foo", "barz")
	all := p.All()
	c.Assert(all, HasLen, 3)
	c.Assert(all[0], DeepEquals, URLQueryParameter{Key: "foo", Value: "barz"})
	c.Assert(all[1], DeepEquals, URLQueryParameter{Key: "foo2", Value: "bar2"})
	c.Assert(all[2], DeepEquals, URLQueryParameter{Key: "foo3", Value: "bar3"})
}

func (s *URLQueryParserSuite) TestGet(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Add("foo2", "bar2")
	p.Add("foo3", "bar3")
	c.Assert(p.Get("foo"), Equals, "bar")
	c.Assert(p.Get("foo2"), Equals, "bar2")
	c.Assert(p.Get("foo3"), Equals, "bar3")
}

func (s *URLQueryParserSuite) TestGetOnDuplicateKeyReturnsFirst(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Add("foo2", "bar2")
	p.Add("foo", "baz")
	c.Assert(p.Get("foo"), Equals, "bar")
}

func (s *URLQueryParserSuite) TestGetNonExistingKey(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Add("foo2", "bar2")
	p.Add("foo3", "bar3")
	c.Assert(p.Get("bar"), Equals, "")
}

func (s *URLQueryParserSuite) TestGetMulti(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Add("foo2", "bar2")
	p.Add("foo3", "bar3")
	c.Assert(p.GetMulti("foo"), DeepEquals, []string{"bar"})
	c.Assert(p.GetMulti("foo2"), DeepEquals, []string{"bar2"})
	c.Assert(p.GetMulti("foo3"), DeepEquals, []string{"bar3"})
}

func (s *URLQueryParserSuite) TestGetMultiOnDuplicateKeys(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Add("foo2", "bar2")
	p.Add("foo", "baz")
	p.Add("foo3", "bar3")
	p.Add("foo", "barz")
	c.Assert(p.GetMulti("foo"), DeepEquals, []string{"bar", "baz", "barz"})
	c.Assert(p.GetMulti("foo2"), DeepEquals, []string{"bar2"})
	c.Assert(p.GetMulti("foo3"), DeepEquals, []string{"bar3"})
}

func (s *URLQueryParserSuite) TestGetMultiNonExistingKey(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Add("foo2", "bar2")
	p.Add("foo3", "bar3")
	c.Assert(p.GetMulti("bar"), HasLen, 0)
}

func (s *URLQueryParserSuite) TestAll(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Add("foo3", "bar3")
	p.Add("foo3", "bar3.1")
	p.Add("foo2", "bar2")
	p.Add("foo", "bar.1")
	all := p.All()
	c.Assert(all, HasLen, 5)
	c.Assert(all[0], DeepEquals, URLQueryParameter{Key: "foo", Value: "bar"})
	c.Assert(all[1], DeepEquals, URLQueryParameter{Key: "foo3", Value: "bar3"})
	c.Assert(all[2], DeepEquals, URLQueryParameter{Key: "foo3", Value: "bar3.1"})
	c.Assert(all[3], DeepEquals, URLQueryParameter{Key: "foo2", Value: "bar2"})
	c.Assert(all[4], DeepEquals, URLQueryParameter{Key: "foo", Value: "bar.1"})
}

func (s *URLQueryParserSuite) TestAllEmpty(c *C) {
	p := URLQueryParameters{}
	c.Assert(p.All(), HasLen, 0)
}

func (s *URLQueryParserSuite) TestValues(c *C) {
	p := URLQueryParameters{}
	p.Add("foo", "bar")
	p.Add("foo3", "bar3")
	p.Add("foo3", "bar3.1")
	p.Add("foo2", "bar2")
	p.Add("foo", "bar.1")
	values := p.Values()
	c.Assert(values, HasLen, 3)
	c.Assert(values["foo"], DeepEquals, []string{"bar", "bar.1"})
	c.Assert(values["foo2"], DeepEquals, []string{"bar2"})
	c.Assert(values["foo3"], DeepEquals, []string{"bar3", "bar3.1"})
}

func (s *URLQueryParserSuite) TestValuesEmpty(c *C) {
	p := URLQueryParameters{}
	c.Assert(p.Values(), HasLen, 0)
}

func (s *URLQueryParserSuite) TestEncode(c *C) {
	p := URLQueryParameters{}
	p.Add("abc", "def")
	p.Add("zyx", "wvu")
	p.Add("abc", "123")
	p.Add("lmn", "opq")
	p.Add("ghi", "jkl")
	p.Add("abc", "xyz")
	c.Assert(p.Encode(), Equals, "abc=def&zyx=wvu&abc=123&lmn=opq&ghi=jkl&abc=xyz")
}

func (s *URLQueryParserSuite) TestEncodeEscapesSpecialChars(c *C) {
	p := URLQueryParameters{}
	p.Add("foo:bar", "foo&baz")
	c.Assert(p.Encode(), Equals, "foo%3Abar=foo%26baz")
}
