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

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Parameters struct {
	Resource  `bson:",inline"`
	Parameter []ParametersParameterComponent `bson:"parameter,omitempty" json:"parameter,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Parameters) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Parameters"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Parameters), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Parameters) GetBSON() (interface{}, error) {
	x.ResourceType = "Parameters"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "parameters" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type parameters Parameters

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Parameters) UnmarshalJSON(data []byte) (err error) {
	x2 := parameters{}
	if err = json.Unmarshal(data, &x2); err == nil {
		*x = Parameters(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Parameters) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Parameters"
	} else if x.ResourceType != "Parameters" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Parameters, instead received %s", x.ResourceType))
	}
	return nil
}

type ParametersParameterComponent struct {
	Name                 string                         `bson:"name,omitempty" json:"name,omitempty"`
	ValueString          string                         `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueInteger         *int32                         `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDateTime        *FHIRDateTime                  `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueBoolean         *bool                          `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCodeableConcept *CodeableConcept               `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueRange           *Range                         `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	Resource             interface{}                    `bson:"resource,omitempty" json:"resource,omitempty"`
	Part                 []ParametersParameterComponent `bson:"part,omitempty" json:"part,omitempty"`
}

// The "parametersParameterComponent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type parametersParameterComponent ParametersParameterComponent

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ParametersParameterComponent) UnmarshalJSON(data []byte) (err error) {
	x2 := parametersParameterComponent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Resource != nil {
			x2.Resource = MapToResource(x2.Resource, true)
		}
		*x = ParametersParameterComponent(x2)
	}
	return
}
