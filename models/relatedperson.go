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

type RelatedPerson struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient        *Reference       `bson:"patient,omitempty" json:"patient,omitempty"`
	Relationship   *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name           *HumanName       `bson:"name,omitempty" json:"name,omitempty"`
	Telecom        []ContactPoint   `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender         string           `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate      *FHIRDateTime    `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address        []Address        `bson:"address,omitempty" json:"address,omitempty"`
	Photo          []Attachment     `bson:"photo,omitempty" json:"photo,omitempty"`
	Period         *Period          `bson:"period,omitempty" json:"period,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *RelatedPerson) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "RelatedPerson"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to RelatedPerson), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *RelatedPerson) GetBSON() (interface{}, error) {
	x.ResourceType = "RelatedPerson"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "relatedPerson" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type relatedPerson RelatedPerson

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *RelatedPerson) UnmarshalJSON(data []byte) (err error) {
	x2 := relatedPerson{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = RelatedPerson(x2)
		return x.checkResourceType()
	}
	return
}

func (x *RelatedPerson) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "RelatedPerson"
	} else if x.ResourceType != "RelatedPerson" {
		return errors.New(fmt.Sprintf("Expected resourceType to be RelatedPerson, instead received %s", x.ResourceType))
	}
	return nil
}

type RelatedPersonPlus struct {
	RelatedPerson             `bson:",inline"`
	RelatedPersonPlusIncludes `bson:",inline"`
}

type RelatedPersonPlusIncludes struct {
	IncludedPatientResources *[]Patient `bson:"_includedPatientResources,omitempty"`
}

func (r *RelatedPersonPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if r.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResources))
	} else if len(*r.IncludedPatientResources) == 1 {
		patient = &(*r.IncludedPatientResources)[0]
	}
	return
}

func (r *RelatedPersonPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPatientResources != nil {
		for _, r := range *r.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
