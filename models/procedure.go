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

type Procedure struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject               *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Status                string                          `bson:"status,omitempty" json:"status,omitempty"`
	Category              *CodeableConcept                `bson:"category,omitempty" json:"category,omitempty"`
	Code                  *CodeableConcept                `bson:"code,omitempty" json:"code,omitempty"`
	NotPerformed          *bool                           `bson:"notPerformed,omitempty" json:"notPerformed,omitempty"`
	ReasonNotPerformed    []CodeableConcept               `bson:"reasonNotPerformed,omitempty" json:"reasonNotPerformed,omitempty"`
	BodySite              []CodeableConcept               `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	ReasonCodeableConcept *CodeableConcept                `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                      `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Performer             []ProcedurePerformerComponent   `bson:"performer,omitempty" json:"performer,omitempty"`
	PerformedDateTime     *FHIRDateTime                   `bson:"performedDateTime,omitempty" json:"performedDateTime,omitempty"`
	PerformedPeriod       *Period                         `bson:"performedPeriod,omitempty" json:"performedPeriod,omitempty"`
	Encounter             *Reference                      `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Location              *Reference                      `bson:"location,omitempty" json:"location,omitempty"`
	Outcome               *CodeableConcept                `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Report                []Reference                     `bson:"report,omitempty" json:"report,omitempty"`
	Complication          []CodeableConcept               `bson:"complication,omitempty" json:"complication,omitempty"`
	FollowUp              []CodeableConcept               `bson:"followUp,omitempty" json:"followUp,omitempty"`
	Request               *Reference                      `bson:"request,omitempty" json:"request,omitempty"`
	Notes                 []Annotation                    `bson:"notes,omitempty" json:"notes,omitempty"`
	FocalDevice           []ProcedureFocalDeviceComponent `bson:"focalDevice,omitempty" json:"focalDevice,omitempty"`
	Used                  []Reference                     `bson:"used,omitempty" json:"used,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Procedure) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Procedure"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Procedure), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Procedure) GetBSON() (interface{}, error) {
	x.ResourceType = "Procedure"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "procedure" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type procedure Procedure

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Procedure) UnmarshalJSON(data []byte) (err error) {
	x2 := procedure{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Procedure(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Procedure) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Procedure"
	} else if x.ResourceType != "Procedure" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Procedure, instead received %s", x.ResourceType))
	}
	return nil
}

type ProcedurePerformerComponent struct {
	Actor *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
	Role  *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ProcedureFocalDeviceComponent struct {
	Action      *CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Manipulated *Reference       `bson:"manipulated,omitempty" json:"manipulated,omitempty"`
}

type ProcedurePlus struct {
	Procedure             `bson:",inline"`
	ProcedurePlusIncludes `bson:",inline"`
}

type ProcedurePlusIncludes struct {
	IncludedPerformerPractitionerResources  *[]Practitioner  `bson:"_includedPerformerPractitionerResources,omitempty"`
	IncludedPerformerOrganizationResources  *[]Organization  `bson:"_includedPerformerOrganizationResources,omitempty"`
	IncludedPerformerPatientResources       *[]Patient       `bson:"_includedPerformerPatientResources,omitempty"`
	IncludedPerformerRelatedPersonResources *[]RelatedPerson `bson:"_includedPerformerRelatedPersonResources,omitempty"`
	IncludedSubjectGroupResources           *[]Group         `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectPatientResources         *[]Patient       `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedPatientResources                *[]Patient       `bson:"_includedPatientResources,omitempty"`
	IncludedLocationResources               *[]Location      `bson:"_includedLocationResources,omitempty"`
	IncludedEncounterResources              *[]Encounter     `bson:"_includedEncounterResources,omitempty"`
}

func (p *ProcedurePlusIncludes) GetIncludedPerformerPractitionerResource() (practitioner *Practitioner, err error) {
	if p.IncludedPerformerPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedPerformerPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedPerformerPractitionerResources))
	} else if len(*p.IncludedPerformerPractitionerResources) == 1 {
		practitioner = &(*p.IncludedPerformerPractitionerResources)[0]
	}
	return
}

func (p *ProcedurePlusIncludes) GetIncludedPerformerOrganizationResource() (organization *Organization, err error) {
	if p.IncludedPerformerOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedPerformerOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedPerformerOrganizationResources))
	} else if len(*p.IncludedPerformerOrganizationResources) == 1 {
		organization = &(*p.IncludedPerformerOrganizationResources)[0]
	}
	return
}

func (p *ProcedurePlusIncludes) GetIncludedPerformerPatientResource() (patient *Patient, err error) {
	if p.IncludedPerformerPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPerformerPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPerformerPatientResources))
	} else if len(*p.IncludedPerformerPatientResources) == 1 {
		patient = &(*p.IncludedPerformerPatientResources)[0]
	}
	return
}

func (p *ProcedurePlusIncludes) GetIncludedPerformerRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedPerformerRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedPerformerRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedPerformerRelatedPersonResources))
	} else if len(*p.IncludedPerformerRelatedPersonResources) == 1 {
		relatedPerson = &(*p.IncludedPerformerRelatedPersonResources)[0]
	}
	return
}

func (p *ProcedurePlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if p.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*p.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*p.IncludedSubjectGroupResources))
	} else if len(*p.IncludedSubjectGroupResources) == 1 {
		group = &(*p.IncludedSubjectGroupResources)[0]
	}
	return
}

func (p *ProcedurePlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if p.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedSubjectPatientResources))
	} else if len(*p.IncludedSubjectPatientResources) == 1 {
		patient = &(*p.IncludedSubjectPatientResources)[0]
	}
	return
}

func (p *ProcedurePlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if p.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResources))
	} else if len(*p.IncludedPatientResources) == 1 {
		patient = &(*p.IncludedPatientResources)[0]
	}
	return
}

func (p *ProcedurePlusIncludes) GetIncludedLocationResource() (location *Location, err error) {
	if p.IncludedLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*p.IncludedLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*p.IncludedLocationResources))
	} else if len(*p.IncludedLocationResources) == 1 {
		location = &(*p.IncludedLocationResources)[0]
	}
	return
}

func (p *ProcedurePlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if p.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*p.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*p.IncludedEncounterResources))
	} else if len(*p.IncludedEncounterResources) == 1 {
		encounter = &(*p.IncludedEncounterResources)[0]
	}
	return
}

func (p *ProcedurePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPerformerPractitionerResources != nil {
		for _, r := range *p.IncludedPerformerPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedPerformerOrganizationResources != nil {
		for _, r := range *p.IncludedPerformerOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedPerformerPatientResources != nil {
		for _, r := range *p.IncludedPerformerPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedPerformerRelatedPersonResources != nil {
		for _, r := range *p.IncludedPerformerRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedSubjectGroupResources != nil {
		for _, r := range *p.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedSubjectPatientResources != nil {
		for _, r := range *p.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedPatientResources != nil {
		for _, r := range *p.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedLocationResources != nil {
		for _, r := range *p.IncludedLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedEncounterResources != nil {
		for _, r := range *p.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
