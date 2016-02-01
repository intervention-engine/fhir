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

type Composition struct {
	DomainResource  `bson:",inline"`
	Identifier      *Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Date            *FHIRDateTime                  `bson:"date,omitempty" json:"date,omitempty"`
	Type            *CodeableConcept               `bson:"type,omitempty" json:"type,omitempty"`
	Class           *CodeableConcept               `bson:"class,omitempty" json:"class,omitempty"`
	Title           string                         `bson:"title,omitempty" json:"title,omitempty"`
	Status          string                         `bson:"status,omitempty" json:"status,omitempty"`
	Confidentiality string                         `bson:"confidentiality,omitempty" json:"confidentiality,omitempty"`
	Subject         *Reference                     `bson:"subject,omitempty" json:"subject,omitempty"`
	Author          []Reference                    `bson:"author,omitempty" json:"author,omitempty"`
	Attester        []CompositionAttesterComponent `bson:"attester,omitempty" json:"attester,omitempty"`
	Custodian       *Reference                     `bson:"custodian,omitempty" json:"custodian,omitempty"`
	Event           []CompositionEventComponent    `bson:"event,omitempty" json:"event,omitempty"`
	Encounter       *Reference                     `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Section         []CompositionSectionComponent  `bson:"section,omitempty" json:"section,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Composition) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Composition"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Composition), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Composition) GetBSON() (interface{}, error) {
	x.ResourceType = "Composition"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "composition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type composition Composition

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Composition) UnmarshalJSON(data []byte) (err error) {
	x2 := composition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Composition(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Composition) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Composition"
	} else if x.ResourceType != "Composition" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Composition, instead received %s", x.ResourceType))
	}
	return nil
}

type CompositionAttesterComponent struct {
	Mode  []string      `bson:"mode,omitempty" json:"mode,omitempty"`
	Time  *FHIRDateTime `bson:"time,omitempty" json:"time,omitempty"`
	Party *Reference    `bson:"party,omitempty" json:"party,omitempty"`
}

type CompositionEventComponent struct {
	Code   []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Period *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Detail []Reference       `bson:"detail,omitempty" json:"detail,omitempty"`
}

type CompositionSectionComponent struct {
	Title       string                        `bson:"title,omitempty" json:"title,omitempty"`
	Code        *CodeableConcept              `bson:"code,omitempty" json:"code,omitempty"`
	Text        *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	Mode        string                        `bson:"mode,omitempty" json:"mode,omitempty"`
	OrderedBy   *CodeableConcept              `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Entry       []Reference                   `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason *CodeableConcept              `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
	Section     []CompositionSectionComponent `bson:"section,omitempty" json:"section,omitempty"`
}

type CompositionPlus struct {
	Composition             `bson:",inline"`
	CompositionPlusIncludes `bson:",inline"`
}

type CompositionPlusIncludes struct {
	IncludedAuthorPractitionerResources   *[]Practitioner  `bson:"_includedAuthorPractitionerResources,omitempty"`
	IncludedAuthorDeviceResources         *[]Device        `bson:"_includedAuthorDeviceResources,omitempty"`
	IncludedAuthorPatientResources        *[]Patient       `bson:"_includedAuthorPatientResources,omitempty"`
	IncludedAuthorRelatedPersonResources  *[]RelatedPerson `bson:"_includedAuthorRelatedPersonResources,omitempty"`
	IncludedEncounterResources            *[]Encounter     `bson:"_includedEncounterResources,omitempty"`
	IncludedAttesterPractitionerResources *[]Practitioner  `bson:"_includedAttesterPractitionerResources,omitempty"`
	IncludedAttesterOrganizationResources *[]Organization  `bson:"_includedAttesterOrganizationResources,omitempty"`
	IncludedAttesterPatientResources      *[]Patient       `bson:"_includedAttesterPatientResources,omitempty"`
	IncludedPatientResources              *[]Patient       `bson:"_includedPatientResources,omitempty"`
}

func (c *CompositionPlusIncludes) GetIncludedAuthorPractitionerResources() (practitioners []Practitioner, err error) {
	if c.IncludedAuthorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *c.IncludedAuthorPractitionerResources
	}
	return
}

func (c *CompositionPlusIncludes) GetIncludedAuthorDeviceResources() (devices []Device, err error) {
	if c.IncludedAuthorDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *c.IncludedAuthorDeviceResources
	}
	return
}

func (c *CompositionPlusIncludes) GetIncludedAuthorPatientResources() (patients []Patient, err error) {
	if c.IncludedAuthorPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedAuthorPatientResources
	}
	return
}

func (c *CompositionPlusIncludes) GetIncludedAuthorRelatedPersonResources() (relatedPeople []RelatedPerson, err error) {
	if c.IncludedAuthorRelatedPersonResources == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *c.IncludedAuthorRelatedPersonResources
	}
	return
}

func (c *CompositionPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if c.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*c.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*c.IncludedEncounterResources))
	} else if len(*c.IncludedEncounterResources) == 1 {
		encounter = &(*c.IncludedEncounterResources)[0]
	}
	return
}

func (c *CompositionPlusIncludes) GetIncludedAttesterPractitionerResource() (practitioner *Practitioner, err error) {
	if c.IncludedAttesterPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedAttesterPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedAttesterPractitionerResources))
	} else if len(*c.IncludedAttesterPractitionerResources) == 1 {
		practitioner = &(*c.IncludedAttesterPractitionerResources)[0]
	}
	return
}

func (c *CompositionPlusIncludes) GetIncludedAttesterOrganizationResource() (organization *Organization, err error) {
	if c.IncludedAttesterOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedAttesterOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedAttesterOrganizationResources))
	} else if len(*c.IncludedAttesterOrganizationResources) == 1 {
		organization = &(*c.IncludedAttesterOrganizationResources)[0]
	}
	return
}

func (c *CompositionPlusIncludes) GetIncludedAttesterPatientResource() (patient *Patient, err error) {
	if c.IncludedAttesterPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedAttesterPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedAttesterPatientResources))
	} else if len(*c.IncludedAttesterPatientResources) == 1 {
		patient = &(*c.IncludedAttesterPatientResources)[0]
	}
	return
}

func (c *CompositionPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if c.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResources))
	} else if len(*c.IncludedPatientResources) == 1 {
		patient = &(*c.IncludedPatientResources)[0]
	}
	return
}

func (c *CompositionPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedAuthorPractitionerResources != nil {
		for _, r := range *c.IncludedAuthorPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedAuthorDeviceResources != nil {
		for _, r := range *c.IncludedAuthorDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedAuthorPatientResources != nil {
		for _, r := range *c.IncludedAuthorPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedAuthorRelatedPersonResources != nil {
		for _, r := range *c.IncludedAuthorRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedEncounterResources != nil {
		for _, r := range *c.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedAttesterPractitionerResources != nil {
		for _, r := range *c.IncludedAttesterPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedAttesterOrganizationResources != nil {
		for _, r := range *c.IncludedAttesterOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedAttesterPatientResources != nil {
		for _, r := range *c.IncludedAttesterPatientResources {
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
