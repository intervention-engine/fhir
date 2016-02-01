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

type Goal struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject              *Reference             `bson:"subject,omitempty" json:"subject,omitempty"`
	StartDate            *FHIRDateTime          `bson:"startDate,omitempty" json:"startDate,omitempty"`
	StartCodeableConcept *CodeableConcept       `bson:"startCodeableConcept,omitempty" json:"startCodeableConcept,omitempty"`
	TargetDate           *FHIRDateTime          `bson:"targetDate,omitempty" json:"targetDate,omitempty"`
	TargetDuration       *Quantity              `bson:"targetDuration,omitempty" json:"targetDuration,omitempty"`
	Category             []CodeableConcept      `bson:"category,omitempty" json:"category,omitempty"`
	Description          string                 `bson:"description,omitempty" json:"description,omitempty"`
	Status               string                 `bson:"status,omitempty" json:"status,omitempty"`
	StatusDate           *FHIRDateTime          `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	StatusReason         *CodeableConcept       `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Author               *Reference             `bson:"author,omitempty" json:"author,omitempty"`
	Priority             *CodeableConcept       `bson:"priority,omitempty" json:"priority,omitempty"`
	Addresses            []Reference            `bson:"addresses,omitempty" json:"addresses,omitempty"`
	Note                 []Annotation           `bson:"note,omitempty" json:"note,omitempty"`
	Outcome              []GoalOutcomeComponent `bson:"outcome,omitempty" json:"outcome,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Goal) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Goal"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Goal), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Goal) GetBSON() (interface{}, error) {
	x.ResourceType = "Goal"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "goal" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type goal Goal

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Goal) UnmarshalJSON(data []byte) (err error) {
	x2 := goal{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Goal(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Goal) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Goal"
	} else if x.ResourceType != "Goal" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Goal, instead received %s", x.ResourceType))
	}
	return nil
}

type GoalOutcomeComponent struct {
	ResultCodeableConcept *CodeableConcept `bson:"resultCodeableConcept,omitempty" json:"resultCodeableConcept,omitempty"`
	ResultReference       *Reference       `bson:"resultReference,omitempty" json:"resultReference,omitempty"`
}

type GoalPlus struct {
	Goal             `bson:",inline"`
	GoalPlusIncludes `bson:",inline"`
}

type GoalPlusIncludes struct {
	IncludedPatientResources             *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedSubjectGroupResources        *[]Group        `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectOrganizationResources *[]Organization `bson:"_includedSubjectOrganizationResources,omitempty"`
	IncludedSubjectPatientResources      *[]Patient      `bson:"_includedSubjectPatientResources,omitempty"`
}

func (g *GoalPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if g.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*g.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*g.IncludedPatientResources))
	} else if len(*g.IncludedPatientResources) == 1 {
		patient = &(*g.IncludedPatientResources)[0]
	}
	return
}

func (g *GoalPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if g.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*g.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*g.IncludedSubjectGroupResources))
	} else if len(*g.IncludedSubjectGroupResources) == 1 {
		group = &(*g.IncludedSubjectGroupResources)[0]
	}
	return
}

func (g *GoalPlusIncludes) GetIncludedSubjectOrganizationResource() (organization *Organization, err error) {
	if g.IncludedSubjectOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*g.IncludedSubjectOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*g.IncludedSubjectOrganizationResources))
	} else if len(*g.IncludedSubjectOrganizationResources) == 1 {
		organization = &(*g.IncludedSubjectOrganizationResources)[0]
	}
	return
}

func (g *GoalPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if g.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*g.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*g.IncludedSubjectPatientResources))
	} else if len(*g.IncludedSubjectPatientResources) == 1 {
		patient = &(*g.IncludedSubjectPatientResources)[0]
	}
	return
}

func (g *GoalPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.IncludedPatientResources != nil {
		for _, r := range *g.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if g.IncludedSubjectGroupResources != nil {
		for _, r := range *g.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if g.IncludedSubjectOrganizationResources != nil {
		for _, r := range *g.IncludedSubjectOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if g.IncludedSubjectPatientResources != nil {
		for _, r := range *g.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
