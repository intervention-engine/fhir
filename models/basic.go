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

type Basic struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code           *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Subject        *Reference       `bson:"subject,omitempty" json:"subject,omitempty"`
	Author         *Reference       `bson:"author,omitempty" json:"author,omitempty"`
	Created        *FHIRDateTime    `bson:"created,omitempty" json:"created,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Basic) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Basic"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Basic), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Basic) GetBSON() (interface{}, error) {
	x.ResourceType = "Basic"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "basic" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type basic Basic

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Basic) UnmarshalJSON(data []byte) (err error) {
	x2 := basic{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Basic(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Basic) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Basic"
	} else if x.ResourceType != "Basic" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Basic, instead received %s", x.ResourceType))
	}
	return nil
}

type BasicPlus struct {
	Basic             `bson:",inline"`
	BasicPlusIncludes `bson:",inline"`
}

type BasicPlusIncludes struct {
	IncludedPatientResources             *[]Patient       `bson:"_includedPatientResources,omitempty"`
	IncludedAuthorPractitionerResources  *[]Practitioner  `bson:"_includedAuthorPractitionerResources,omitempty"`
	IncludedAuthorPatientResources       *[]Patient       `bson:"_includedAuthorPatientResources,omitempty"`
	IncludedAuthorRelatedPersonResources *[]RelatedPerson `bson:"_includedAuthorRelatedPersonResources,omitempty"`
}

func (b *BasicPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if b.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*b.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*b.IncludedPatientResources))
	} else if len(*b.IncludedPatientResources) == 1 {
		patient = &(*b.IncludedPatientResources)[0]
	}
	return
}

func (b *BasicPlusIncludes) GetIncludedAuthorPractitionerResource() (practitioner *Practitioner, err error) {
	if b.IncludedAuthorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*b.IncludedAuthorPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*b.IncludedAuthorPractitionerResources))
	} else if len(*b.IncludedAuthorPractitionerResources) == 1 {
		practitioner = &(*b.IncludedAuthorPractitionerResources)[0]
	}
	return
}

func (b *BasicPlusIncludes) GetIncludedAuthorPatientResource() (patient *Patient, err error) {
	if b.IncludedAuthorPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*b.IncludedAuthorPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*b.IncludedAuthorPatientResources))
	} else if len(*b.IncludedAuthorPatientResources) == 1 {
		patient = &(*b.IncludedAuthorPatientResources)[0]
	}
	return
}

func (b *BasicPlusIncludes) GetIncludedAuthorRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if b.IncludedAuthorRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*b.IncludedAuthorRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*b.IncludedAuthorRelatedPersonResources))
	} else if len(*b.IncludedAuthorRelatedPersonResources) == 1 {
		relatedPerson = &(*b.IncludedAuthorRelatedPersonResources)[0]
	}
	return
}

func (b *BasicPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.IncludedPatientResources != nil {
		for _, r := range *b.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if b.IncludedAuthorPractitionerResources != nil {
		for _, r := range *b.IncludedAuthorPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if b.IncludedAuthorPatientResources != nil {
		for _, r := range *b.IncludedAuthorPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if b.IncludedAuthorRelatedPersonResources != nil {
		for _, r := range *b.IncludedAuthorRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
