package server

import (
	"fmt"
	"reflect"
	"io/ioutil"
	"encoding/json"
	"encoding/xml"
	. "gopkg.in/check.v1"
)

type FormatConversionSuite struct {
}

var _ = Suite(&FormatConversionSuite{})

func (s *FormatConversionSuite) TestBundle(c *C) {
	converter := NewFhirFormatConverter()

	xml, err := ioutil.ReadFile("../fixtures/bundle-transaction.xml")
	c.Assert(err, IsNil)
	json, err := ioutil.ReadFile("../fixtures/bundle-transaction.json")
	c.Assert(err, IsNil)

	result, err := converter.XmlToJson(string(xml))
	c.Assert(err, IsNil)
	areEqual, err := areEqualJSON(result, string(json))
	c.Assert(err, IsNil)
	c.Assert(areEqual, Equals, true)
	// c.Assert(strings.Replace(result, "\n", "\r\n", -1), Equals, string(json))

	result, err = converter.JsonToXml(string(json))
	c.Assert(err, IsNil)
	areEqual, err = areEqualXML(result, string(xml))
	c.Assert(err, IsNil)
	c.Assert(areEqual, Equals, true)
	// c.Assert(strings.Replace(result, "\n", "\r\n", -1), Equals, string(xml))
}

func areEqualJSON(s1, s2 string) (bool, error) {
	// thanks to turtlemonvh https://gist.github.com/turtlemonvh/e4f7404e28387fadb8ad275a99596f67

	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

func areEqualXML(s1, s2 string) (bool, error) {
	// thanks to turtlemonvh https://gist.github.com/turtlemonvh/e4f7404e28387fadb8ad275a99596f67

	var o1 interface{}
	var o2 interface{}

	var err error
	err = xml.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = xml.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}
