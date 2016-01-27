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

type ProcedureRequest struct {
	DomainResource          `bson:",inline"`
	Identifier              []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject                 *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	Code                    *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	BodySite                []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	ReasonCodeableConcept   *CodeableConcept  `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference         *Reference        `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	ScheduledDateTime       *FHIRDateTime     `bson:"scheduledDateTime,omitempty" json:"scheduledDateTime,omitempty"`
	ScheduledPeriod         *Period           `bson:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	ScheduledTiming         *Timing           `bson:"scheduledTiming,omitempty" json:"scheduledTiming,omitempty"`
	Encounter               *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Performer               *Reference        `bson:"performer,omitempty" json:"performer,omitempty"`
	Status                  string            `bson:"status,omitempty" json:"status,omitempty"`
	Notes                   []Annotation      `bson:"notes,omitempty" json:"notes,omitempty"`
	AsNeededBoolean         *bool             `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept  `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	OrderedOn               *FHIRDateTime     `bson:"orderedOn,omitempty" json:"orderedOn,omitempty"`
	Orderer                 *Reference        `bson:"orderer,omitempty" json:"orderer,omitempty"`
	Priority                string            `bson:"priority,omitempty" json:"priority,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ProcedureRequest) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		ProcedureRequest
	}{
		ResourceType:     "ProcedureRequest",
		ProcedureRequest: *resource,
	}
	return json.Marshal(x)
}

// The "procedureRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type procedureRequest ProcedureRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ProcedureRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := procedureRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ProcedureRequest(x2)
	}
	return
}

type ProcedureRequestPlus struct {
	ProcedureRequest             `bson:",inline"`
	ProcedureRequestPlusIncludes `bson:",inline"`
}

type ProcedureRequestPlusIncludes struct {
	IncludedPerformerPractitionerResources  *[]Practitioner  `bson:"_includedPerformerPractitionerResources,omitempty"`
	IncludedPerformerOrganizationResources  *[]Organization  `bson:"_includedPerformerOrganizationResources,omitempty"`
	IncludedPerformerPatientResources       *[]Patient       `bson:"_includedPerformerPatientResources,omitempty"`
	IncludedPerformerRelatedPersonResources *[]RelatedPerson `bson:"_includedPerformerRelatedPersonResources,omitempty"`
	IncludedSubjectGroupResources           *[]Group         `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectPatientResources         *[]Patient       `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedPatientResources                *[]Patient       `bson:"_includedPatientResources,omitempty"`
	IncludedOrdererPractitionerResources    *[]Practitioner  `bson:"_includedOrdererPractitionerResources,omitempty"`
	IncludedOrdererDeviceResources          *[]Device        `bson:"_includedOrdererDeviceResources,omitempty"`
	IncludedOrdererPatientResources         *[]Patient       `bson:"_includedOrdererPatientResources,omitempty"`
	IncludedOrdererRelatedPersonResources   *[]RelatedPerson `bson:"_includedOrdererRelatedPersonResources,omitempty"`
	IncludedEncounterResources              *[]Encounter     `bson:"_includedEncounterResources,omitempty"`
}

func (p *ProcedureRequestPlusIncludes) GetIncludedPerformerPractitionerResource() (practitioner *Practitioner, err error) {
	if p.IncludedPerformerPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedPerformerPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedPerformerPractitionerResources))
	} else if len(*p.IncludedPerformerPractitionerResources) == 1 {
		practitioner = &(*p.IncludedPerformerPractitionerResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedPerformerOrganizationResource() (organization *Organization, err error) {
	if p.IncludedPerformerOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedPerformerOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedPerformerOrganizationResources))
	} else if len(*p.IncludedPerformerOrganizationResources) == 1 {
		organization = &(*p.IncludedPerformerOrganizationResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedPerformerPatientResource() (patient *Patient, err error) {
	if p.IncludedPerformerPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPerformerPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPerformerPatientResources))
	} else if len(*p.IncludedPerformerPatientResources) == 1 {
		patient = &(*p.IncludedPerformerPatientResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedPerformerRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedPerformerRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedPerformerRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedPerformerRelatedPersonResources))
	} else if len(*p.IncludedPerformerRelatedPersonResources) == 1 {
		relatedPerson = &(*p.IncludedPerformerRelatedPersonResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if p.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*p.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*p.IncludedSubjectGroupResources))
	} else if len(*p.IncludedSubjectGroupResources) == 1 {
		group = &(*p.IncludedSubjectGroupResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if p.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedSubjectPatientResources))
	} else if len(*p.IncludedSubjectPatientResources) == 1 {
		patient = &(*p.IncludedSubjectPatientResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if p.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResources))
	} else if len(*p.IncludedPatientResources) == 1 {
		patient = &(*p.IncludedPatientResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedOrdererPractitionerResource() (practitioner *Practitioner, err error) {
	if p.IncludedOrdererPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedOrdererPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedOrdererPractitionerResources))
	} else if len(*p.IncludedOrdererPractitionerResources) == 1 {
		practitioner = &(*p.IncludedOrdererPractitionerResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedOrdererDeviceResource() (device *Device, err error) {
	if p.IncludedOrdererDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*p.IncludedOrdererDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*p.IncludedOrdererDeviceResources))
	} else if len(*p.IncludedOrdererDeviceResources) == 1 {
		device = &(*p.IncludedOrdererDeviceResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedOrdererPatientResource() (patient *Patient, err error) {
	if p.IncludedOrdererPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedOrdererPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedOrdererPatientResources))
	} else if len(*p.IncludedOrdererPatientResources) == 1 {
		patient = &(*p.IncludedOrdererPatientResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedOrdererRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedOrdererRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedOrdererRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedOrdererRelatedPersonResources))
	} else if len(*p.IncludedOrdererRelatedPersonResources) == 1 {
		relatedPerson = &(*p.IncludedOrdererRelatedPersonResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if p.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*p.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*p.IncludedEncounterResources))
	} else if len(*p.IncludedEncounterResources) == 1 {
		encounter = &(*p.IncludedEncounterResources)[0]
	}
	return
}

func (p *ProcedureRequestPlusIncludes) GetIncludedResources() map[string]interface{} {
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
	if p.IncludedOrdererPractitionerResources != nil {
		for _, r := range *p.IncludedOrdererPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedOrdererDeviceResources != nil {
		for _, r := range *p.IncludedOrdererDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedOrdererPatientResources != nil {
		for _, r := range *p.IncludedOrdererPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedOrdererRelatedPersonResources != nil {
		for _, r := range *p.IncludedOrdererRelatedPersonResources {
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
