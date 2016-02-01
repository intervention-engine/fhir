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

type FamilyMemberHistory struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier                            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient         *Reference                              `bson:"patient,omitempty" json:"patient,omitempty"`
	Date            *FHIRDateTime                           `bson:"date,omitempty" json:"date,omitempty"`
	Status          string                                  `bson:"status,omitempty" json:"status,omitempty"`
	Name            string                                  `bson:"name,omitempty" json:"name,omitempty"`
	Relationship    *CodeableConcept                        `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Gender          string                                  `bson:"gender,omitempty" json:"gender,omitempty"`
	BornPeriod      *Period                                 `bson:"bornPeriod,omitempty" json:"bornPeriod,omitempty"`
	BornDate        *FHIRDateTime                           `bson:"bornDate,omitempty" json:"bornDate,omitempty"`
	BornString      string                                  `bson:"bornString,omitempty" json:"bornString,omitempty"`
	AgeAge          *Quantity                               `bson:"ageAge,omitempty" json:"ageAge,omitempty"`
	AgeRange        *Range                                  `bson:"ageRange,omitempty" json:"ageRange,omitempty"`
	AgeString       string                                  `bson:"ageString,omitempty" json:"ageString,omitempty"`
	DeceasedBoolean *bool                                   `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedAge     *Quantity                               `bson:"deceasedAge,omitempty" json:"deceasedAge,omitempty"`
	DeceasedRange   *Range                                  `bson:"deceasedRange,omitempty" json:"deceasedRange,omitempty"`
	DeceasedDate    *FHIRDateTime                           `bson:"deceasedDate,omitempty" json:"deceasedDate,omitempty"`
	DeceasedString  string                                  `bson:"deceasedString,omitempty" json:"deceasedString,omitempty"`
	Note            *Annotation                             `bson:"note,omitempty" json:"note,omitempty"`
	Condition       []FamilyMemberHistoryConditionComponent `bson:"condition,omitempty" json:"condition,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "FamilyMemberHistory"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to FamilyMemberHistory), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *FamilyMemberHistory) GetBSON() (interface{}, error) {
	x.ResourceType = "FamilyMemberHistory"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "familyMemberHistory" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type familyMemberHistory FamilyMemberHistory

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *FamilyMemberHistory) UnmarshalJSON(data []byte) (err error) {
	x2 := familyMemberHistory{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = FamilyMemberHistory(x2)
		return x.checkResourceType()
	}
	return
}

func (x *FamilyMemberHistory) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "FamilyMemberHistory"
	} else if x.ResourceType != "FamilyMemberHistory" {
		return errors.New(fmt.Sprintf("Expected resourceType to be FamilyMemberHistory, instead received %s", x.ResourceType))
	}
	return nil
}

type FamilyMemberHistoryConditionComponent struct {
	Code        *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Outcome     *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OnsetAge    *Quantity        `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetRange  *Range           `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetPeriod *Period          `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetString string           `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	Note        *Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}

type FamilyMemberHistoryPlus struct {
	FamilyMemberHistory             `bson:",inline"`
	FamilyMemberHistoryPlusIncludes `bson:",inline"`
}

type FamilyMemberHistoryPlusIncludes struct {
	IncludedPatientResources *[]Patient `bson:"_includedPatientResources,omitempty"`
}

func (f *FamilyMemberHistoryPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if f.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedPatientResources))
	} else if len(*f.IncludedPatientResources) == 1 {
		patient = &(*f.IncludedPatientResources)[0]
	}
	return
}

func (f *FamilyMemberHistoryPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedPatientResources != nil {
		for _, r := range *f.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
