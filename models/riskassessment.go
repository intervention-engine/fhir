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

type RiskAssessment struct {
	DomainResource `bson:",inline"`
	Subject        *Reference                          `bson:"subject,omitempty" json:"subject,omitempty"`
	Date           *FHIRDateTime                       `bson:"date,omitempty" json:"date,omitempty"`
	Condition      *Reference                          `bson:"condition,omitempty" json:"condition,omitempty"`
	Encounter      *Reference                          `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Performer      *Reference                          `bson:"performer,omitempty" json:"performer,omitempty"`
	Identifier     *Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Method         *CodeableConcept                    `bson:"method,omitempty" json:"method,omitempty"`
	Basis          []Reference                         `bson:"basis,omitempty" json:"basis,omitempty"`
	Prediction     []RiskAssessmentPredictionComponent `bson:"prediction,omitempty" json:"prediction,omitempty"`
	Mitigation     string                              `bson:"mitigation,omitempty" json:"mitigation,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *RiskAssessment) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "RiskAssessment"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to RiskAssessment), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *RiskAssessment) GetBSON() (interface{}, error) {
	x.ResourceType = "RiskAssessment"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "riskAssessment" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type riskAssessment RiskAssessment

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *RiskAssessment) UnmarshalJSON(data []byte) (err error) {
	x2 := riskAssessment{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = RiskAssessment(x2)
		return x.checkResourceType()
	}
	return
}

func (x *RiskAssessment) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "RiskAssessment"
	} else if x.ResourceType != "RiskAssessment" {
		return errors.New(fmt.Sprintf("Expected resourceType to be RiskAssessment, instead received %s", x.ResourceType))
	}
	return nil
}

type RiskAssessmentPredictionComponent struct {
	Outcome                    *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	ProbabilityDecimal         *float64         `bson:"probabilityDecimal,omitempty" json:"probabilityDecimal,omitempty"`
	ProbabilityRange           *Range           `bson:"probabilityRange,omitempty" json:"probabilityRange,omitempty"`
	ProbabilityCodeableConcept *CodeableConcept `bson:"probabilityCodeableConcept,omitempty" json:"probabilityCodeableConcept,omitempty"`
	RelativeRisk               *float64         `bson:"relativeRisk,omitempty" json:"relativeRisk,omitempty"`
	WhenPeriod                 *Period          `bson:"whenPeriod,omitempty" json:"whenPeriod,omitempty"`
	WhenRange                  *Range           `bson:"whenRange,omitempty" json:"whenRange,omitempty"`
	Rationale                  string           `bson:"rationale,omitempty" json:"rationale,omitempty"`
}

type RiskAssessmentPlus struct {
	RiskAssessment             `bson:",inline"`
	RiskAssessmentPlusIncludes `bson:",inline"`
}

type RiskAssessmentPlusIncludes struct {
	IncludedConditionResources             *[]Condition    `bson:"_includedConditionResources,omitempty"`
	IncludedPerformerPractitionerResources *[]Practitioner `bson:"_includedPerformerPractitionerResources,omitempty"`
	IncludedPerformerDeviceResources       *[]Device       `bson:"_includedPerformerDeviceResources,omitempty"`
	IncludedSubjectGroupResources          *[]Group        `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectPatientResources        *[]Patient      `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedPatientResources               *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedEncounterResources             *[]Encounter    `bson:"_includedEncounterResources,omitempty"`
}

func (r *RiskAssessmentPlusIncludes) GetIncludedConditionResource() (condition *Condition, err error) {
	if r.IncludedConditionResources == nil {
		err = errors.New("Included conditions not requested")
	} else if len(*r.IncludedConditionResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 condition, but found %d", len(*r.IncludedConditionResources))
	} else if len(*r.IncludedConditionResources) == 1 {
		condition = &(*r.IncludedConditionResources)[0]
	}
	return
}

func (r *RiskAssessmentPlusIncludes) GetIncludedPerformerPractitionerResource() (practitioner *Practitioner, err error) {
	if r.IncludedPerformerPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*r.IncludedPerformerPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*r.IncludedPerformerPractitionerResources))
	} else if len(*r.IncludedPerformerPractitionerResources) == 1 {
		practitioner = &(*r.IncludedPerformerPractitionerResources)[0]
	}
	return
}

func (r *RiskAssessmentPlusIncludes) GetIncludedPerformerDeviceResource() (device *Device, err error) {
	if r.IncludedPerformerDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*r.IncludedPerformerDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*r.IncludedPerformerDeviceResources))
	} else if len(*r.IncludedPerformerDeviceResources) == 1 {
		device = &(*r.IncludedPerformerDeviceResources)[0]
	}
	return
}

func (r *RiskAssessmentPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if r.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*r.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*r.IncludedSubjectGroupResources))
	} else if len(*r.IncludedSubjectGroupResources) == 1 {
		group = &(*r.IncludedSubjectGroupResources)[0]
	}
	return
}

func (r *RiskAssessmentPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if r.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedSubjectPatientResources))
	} else if len(*r.IncludedSubjectPatientResources) == 1 {
		patient = &(*r.IncludedSubjectPatientResources)[0]
	}
	return
}

func (r *RiskAssessmentPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if r.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResources))
	} else if len(*r.IncludedPatientResources) == 1 {
		patient = &(*r.IncludedPatientResources)[0]
	}
	return
}

func (r *RiskAssessmentPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if r.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*r.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*r.IncludedEncounterResources))
	} else if len(*r.IncludedEncounterResources) == 1 {
		encounter = &(*r.IncludedEncounterResources)[0]
	}
	return
}

func (r *RiskAssessmentPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedConditionResources != nil {
		for _, r := range *r.IncludedConditionResources {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedPerformerPractitionerResources != nil {
		for _, r := range *r.IncludedPerformerPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedPerformerDeviceResources != nil {
		for _, r := range *r.IncludedPerformerDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedSubjectGroupResources != nil {
		for _, r := range *r.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedSubjectPatientResources != nil {
		for _, r := range *r.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedPatientResources != nil {
		for _, r := range *r.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedEncounterResources != nil {
		for _, r := range *r.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
