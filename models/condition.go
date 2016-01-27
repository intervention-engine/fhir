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

type Condition struct {
	DomainResource     `bson:",inline"`
	Identifier         []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient            *Reference                   `bson:"patient,omitempty" json:"patient,omitempty"`
	Encounter          *Reference                   `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Asserter           *Reference                   `bson:"asserter,omitempty" json:"asserter,omitempty"`
	DateRecorded       *FHIRDateTime                `bson:"dateRecorded,omitempty" json:"dateRecorded,omitempty"`
	Code               *CodeableConcept             `bson:"code,omitempty" json:"code,omitempty"`
	Category           *CodeableConcept             `bson:"category,omitempty" json:"category,omitempty"`
	ClinicalStatus     string                       `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus string                       `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Severity           *CodeableConcept             `bson:"severity,omitempty" json:"severity,omitempty"`
	OnsetDateTime      *FHIRDateTime                `bson:"onsetDateTime,omitempty" json:"onsetDateTime,omitempty"`
	OnsetAge           *Quantity                    `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                      `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                       `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetString        string                       `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	AbatementDateTime  *FHIRDateTime                `bson:"abatementDateTime,omitempty" json:"abatementDateTime,omitempty"`
	AbatementAge       *Quantity                    `bson:"abatementAge,omitempty" json:"abatementAge,omitempty"`
	AbatementBoolean   *bool                        `bson:"abatementBoolean,omitempty" json:"abatementBoolean,omitempty"`
	AbatementPeriod    *Period                      `bson:"abatementPeriod,omitempty" json:"abatementPeriod,omitempty"`
	AbatementRange     *Range                       `bson:"abatementRange,omitempty" json:"abatementRange,omitempty"`
	AbatementString    string                       `bson:"abatementString,omitempty" json:"abatementString,omitempty"`
	Stage              *ConditionStageComponent     `bson:"stage,omitempty" json:"stage,omitempty"`
	Evidence           []ConditionEvidenceComponent `bson:"evidence,omitempty" json:"evidence,omitempty"`
	BodySite           []CodeableConcept            `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Notes              string                       `bson:"notes,omitempty" json:"notes,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Condition) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Condition
	}{
		ResourceType: "Condition",
		Condition:    *resource,
	}
	return json.Marshal(x)
}

// The "condition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type condition Condition

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Condition) UnmarshalJSON(data []byte) (err error) {
	x2 := condition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Condition(x2)
	}
	return
}

type ConditionStageComponent struct {
	Summary    *CodeableConcept `bson:"summary,omitempty" json:"summary,omitempty"`
	Assessment []Reference      `bson:"assessment,omitempty" json:"assessment,omitempty"`
}

type ConditionEvidenceComponent struct {
	Code   *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Detail []Reference      `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ConditionPlus struct {
	Condition             `bson:",inline"`
	ConditionPlusIncludes `bson:",inline"`
}

type ConditionPlusIncludes struct {
	IncludedEncounterResources            *[]Encounter    `bson:"_includedEncounterResources,omitempty"`
	IncludedAsserterPractitionerResources *[]Practitioner `bson:"_includedAsserterPractitionerResources,omitempty"`
	IncludedAsserterPatientResources      *[]Patient      `bson:"_includedAsserterPatientResources,omitempty"`
	IncludedPatientResources              *[]Patient      `bson:"_includedPatientResources,omitempty"`
}

func (c *ConditionPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if c.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*c.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*c.IncludedEncounterResources))
	} else if len(*c.IncludedEncounterResources) == 1 {
		encounter = &(*c.IncludedEncounterResources)[0]
	}
	return
}

func (c *ConditionPlusIncludes) GetIncludedAsserterPractitionerResource() (practitioner *Practitioner, err error) {
	if c.IncludedAsserterPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedAsserterPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedAsserterPractitionerResources))
	} else if len(*c.IncludedAsserterPractitionerResources) == 1 {
		practitioner = &(*c.IncludedAsserterPractitionerResources)[0]
	}
	return
}

func (c *ConditionPlusIncludes) GetIncludedAsserterPatientResource() (patient *Patient, err error) {
	if c.IncludedAsserterPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedAsserterPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedAsserterPatientResources))
	} else if len(*c.IncludedAsserterPatientResources) == 1 {
		patient = &(*c.IncludedAsserterPatientResources)[0]
	}
	return
}

func (c *ConditionPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if c.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResources))
	} else if len(*c.IncludedPatientResources) == 1 {
		patient = &(*c.IncludedPatientResources)[0]
	}
	return
}

func (c *ConditionPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedEncounterResources != nil {
		for _, r := range *c.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedAsserterPractitionerResources != nil {
		for _, r := range *c.IncludedAsserterPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedAsserterPatientResources != nil {
		for _, r := range *c.IncludedAsserterPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResources != nil {
		for _, r := range *c.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
