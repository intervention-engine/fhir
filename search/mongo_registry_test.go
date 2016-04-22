package search

import (
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

type MongoRegistrySuite struct{}

var _ = Suite(&MongoRegistrySuite{})

func (s *MongoRegistrySuite) TestRegisterAndLookupBSONBuilder(c *C) {
	build := func(param SearchParam, search *MongoSearcher) (bson.M, error) {
		return bson.M{"foo": param.(*StringParam).String}, nil
	}

	GlobalMongoRegistry().RegisterBSONBuilder("test", build)
	obtained, err := GlobalMongoRegistry().LookupBSONBuilder("test")
	util.CheckErr(err)
	bmap, err := obtained(&StringParam{String: "bar"}, NewMongoSearcher(nil))
	util.CheckErr(err)
	c.Assert(bmap, HasLen, 1)
	c.Assert(bmap["foo"], Equals, "bar")
}

func (s *MongoRegistrySuite) TestLookupNonExistingBSONBuilder(c *C) {
	obtained, err := GlobalMongoRegistry().LookupBSONBuilder("nope")
	c.Assert(err, Not(IsNil))
	c.Assert(obtained, IsNil)
}
