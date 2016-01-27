// Copyright (c) 2011-2015, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import "encoding/json"

type DataElement struct {
	DomainResource `bson:",inline"`
	Url            string                        `bson:"url,omitempty" json:"url,omitempty"`
	Identifier     []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version        string                        `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                        `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                        `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                         `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                        `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []DataElementContactComponent `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                 `bson:"date,omitempty" json:"date,omitempty"`
	UseContext     []CodeableConcept             `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Copyright      string                        `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Stringency     string                        `bson:"stringency,omitempty" json:"stringency,omitempty"`
	Mapping        []DataElementMappingComponent `bson:"mapping,omitempty" json:"mapping,omitempty"`
	Element        []ElementDefinition           `bson:"element,omitempty" json:"element,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DataElement) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		DataElement
	}{
		ResourceType: "DataElement",
		DataElement:  *resource,
	}
	return json.Marshal(x)
}

// The "dataElement" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type dataElement DataElement

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DataElement) UnmarshalJSON(data []byte) (err error) {
	x2 := dataElement{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DataElement(x2)
	}
	return
}

type DataElementContactComponent struct {
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type DataElementMappingComponent struct {
	Identity string `bson:"identity,omitempty" json:"identity,omitempty"`
	Uri      string `bson:"uri,omitempty" json:"uri,omitempty"`
	Name     string `bson:"name,omitempty" json:"name,omitempty"`
	Comments string `bson:"comments,omitempty" json:"comments,omitempty"`
}

type DataElementPlus struct {
	DataElement             `bson:",inline"`
	DataElementPlusIncludes `bson:",inline"`
}

type DataElementPlusIncludes struct {
}

func (d *DataElementPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}
