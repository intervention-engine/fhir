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

type EnrollmentRequest struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ruleset         *Coding       `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset *Coding       `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created         *FHIRDateTime `bson:"created,omitempty" json:"created,omitempty"`
	Target          *Reference    `bson:"target,omitempty" json:"target,omitempty"`
	Provider        *Reference    `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization    *Reference    `bson:"organization,omitempty" json:"organization,omitempty"`
	Subject         *Reference    `bson:"subject,omitempty" json:"subject,omitempty"`
	Coverage        *Reference    `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Relationship    *Coding       `bson:"relationship,omitempty" json:"relationship,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *EnrollmentRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "EnrollmentRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to EnrollmentRequest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *EnrollmentRequest) GetBSON() (interface{}, error) {
	x.ResourceType = "EnrollmentRequest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "enrollmentRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type enrollmentRequest EnrollmentRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *EnrollmentRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := enrollmentRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = EnrollmentRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *EnrollmentRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "EnrollmentRequest"
	} else if x.ResourceType != "EnrollmentRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be EnrollmentRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type EnrollmentRequestPlus struct {
	EnrollmentRequest             `bson:",inline"`
	EnrollmentRequestPlusIncludes `bson:",inline"`
}

type EnrollmentRequestPlusIncludes struct {
	IncludedSubjectResources *[]Patient `bson:"_includedSubjectResources,omitempty"`
	IncludedPatientResources *[]Patient `bson:"_includedPatientResources,omitempty"`
}

func (e *EnrollmentRequestPlusIncludes) GetIncludedSubjectResource() (patient *Patient, err error) {
	if e.IncludedSubjectResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedSubjectResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedSubjectResources))
	} else if len(*e.IncludedSubjectResources) == 1 {
		patient = &(*e.IncludedSubjectResources)[0]
	}
	return
}

func (e *EnrollmentRequestPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if e.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResources))
	} else if len(*e.IncludedPatientResources) == 1 {
		patient = &(*e.IncludedPatientResources)[0]
	}
	return
}

func (e *EnrollmentRequestPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedSubjectResources != nil {
		for _, r := range *e.IncludedSubjectResources {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPatientResources != nil {
		for _, r := range *e.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
