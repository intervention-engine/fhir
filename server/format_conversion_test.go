package server

import (
	"fmt"
	"reflect"
	"io/ioutil"
	"encoding/json"
	"encoding/xml"
	. "gopkg.in/check.v1"
	"github.com/eug48/fhir/models"
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

func (s *FormatConversionSuite) TestObservation(c *C) {
	jsonBytes, err := ioutil.ReadFile("../fixtures/bundle-crucible-1-observation.json")
	c.Assert(err, IsNil)
	obj := &models.Observation{}
	err = json.Unmarshal(jsonBytes, obj)
	c.Assert(*obj.ValueQuantity.Value, Equals, 170.0)
	c.Assert(err, IsNil)
}
func (s *FormatConversionSuite) TestBundle2_json(c *C) {
	jsonBytes, err := ioutil.ReadFile("../fixtures/bundle-crucible-1.json")
	c.Assert(err, IsNil)
	// jsonString := string(jsonBytes)

	obj := &models.Bundle{}
	// fmt.Printf("TestBundle2_json: %s\n", jsonString)
	err = json.Unmarshal(jsonBytes, obj)
	// fmt.Printf("%#v", obj)
	observation1 := obj.Entry[1].Resource.(*models.Observation)

	c.Assert(observation1.Status, Equals, "final")
	c.Assert(*observation1.ValueQuantity.Value, Equals, 170.0)
	c.Assert(err, IsNil)
}
func (s *FormatConversionSuite) TestBundle2_convert_xml(c *C) {
	xmlBytes, err := ioutil.ReadFile("../fixtures/bundle-crucible-1.xml")
	c.Assert(err, IsNil)
	// jsonStr := string(jsonBytes)

	converter := NewFhirFormatConverter()
	jsonString, err := converter.XmlToJson(string(xmlBytes))
	c.Assert(err, IsNil)

	obj := &models.Bundle{}
	// fmt.Printf("TestBundle2_convert_xml: %s\n", jsonString)
	err = json.Unmarshal([]byte(jsonString), obj)
	c.Assert(err, IsNil)

	// fmt.Printf("%#v", obj)
	observation1 := obj.Entry[1].Resource.(*models.Observation)

	c.Assert(observation1.Status, Equals, "final")
	c.Assert(*observation1.ValueQuantity.Value, Equals, 170.0)
}

func areEqualJSON(s1, s2 string) (bool, error) {
	// thanks to turtlemonvh https://gist.github.com/turtlemonvh/e4f7404e28387fadb8ad275a99596f67

	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling s1: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling s22: %s", err.Error())
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
