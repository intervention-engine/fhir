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

type Person struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name                 []HumanName           `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []ContactPoint        `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               string                `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *FHIRDateTime         `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address              []Address             `bson:"address,omitempty" json:"address,omitempty"`
	Photo                *Attachment           `bson:"photo,omitempty" json:"photo,omitempty"`
	ManagingOrganization *Reference            `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Active               *bool                 `bson:"active,omitempty" json:"active,omitempty"`
	Link                 []PersonLinkComponent `bson:"link,omitempty" json:"link,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Person) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Person"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Person), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Person) GetBSON() (interface{}, error) {
	x.ResourceType = "Person"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "person" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type person Person

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Person) UnmarshalJSON(data []byte) (err error) {
	x2 := person{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Person(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Person) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Person"
	} else if x.ResourceType != "Person" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Person, instead received %s", x.ResourceType))
	}
	return nil
}

type PersonLinkComponent struct {
	Target    *Reference `bson:"target,omitempty" json:"target,omitempty"`
	Assurance string     `bson:"assurance,omitempty" json:"assurance,omitempty"`
}

type PersonPlus struct {
	Person             `bson:",inline"`
	PersonPlusIncludes `bson:",inline"`
}

type PersonPlusIncludes struct {
	IncludedPractitionerResources      *[]Practitioner  `bson:"_includedPractitionerResources,omitempty"`
	IncludedLinkPractitionerResources  *[]Practitioner  `bson:"_includedLinkPractitionerResources,omitempty"`
	IncludedLinkPatientResources       *[]Patient       `bson:"_includedLinkPatientResources,omitempty"`
	IncludedLinkPersonResources        *[]Person        `bson:"_includedLinkPersonResources,omitempty"`
	IncludedLinkRelatedPersonResources *[]RelatedPerson `bson:"_includedLinkRelatedPersonResources,omitempty"`
	IncludedRelatedpersonResources     *[]RelatedPerson `bson:"_includedRelatedpersonResources,omitempty"`
	IncludedPatientResources           *[]Patient       `bson:"_includedPatientResources,omitempty"`
	IncludedOrganizationResources      *[]Organization  `bson:"_includedOrganizationResources,omitempty"`
}

func (p *PersonPlusIncludes) GetIncludedPractitionerResource() (practitioner *Practitioner, err error) {
	if p.IncludedPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedPractitionerResources))
	} else if len(*p.IncludedPractitionerResources) == 1 {
		practitioner = &(*p.IncludedPractitionerResources)[0]
	}
	return
}

func (p *PersonPlusIncludes) GetIncludedLinkPractitionerResource() (practitioner *Practitioner, err error) {
	if p.IncludedLinkPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedLinkPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedLinkPractitionerResources))
	} else if len(*p.IncludedLinkPractitionerResources) == 1 {
		practitioner = &(*p.IncludedLinkPractitionerResources)[0]
	}
	return
}

func (p *PersonPlusIncludes) GetIncludedLinkPatientResource() (patient *Patient, err error) {
	if p.IncludedLinkPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedLinkPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedLinkPatientResources))
	} else if len(*p.IncludedLinkPatientResources) == 1 {
		patient = &(*p.IncludedLinkPatientResources)[0]
	}
	return
}

func (p *PersonPlusIncludes) GetIncludedLinkPersonResource() (person *Person, err error) {
	if p.IncludedLinkPersonResources == nil {
		err = errors.New("Included people not requested")
	} else if len(*p.IncludedLinkPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 person, but found %d", len(*p.IncludedLinkPersonResources))
	} else if len(*p.IncludedLinkPersonResources) == 1 {
		person = &(*p.IncludedLinkPersonResources)[0]
	}
	return
}

func (p *PersonPlusIncludes) GetIncludedLinkRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedLinkRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedLinkRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedLinkRelatedPersonResources))
	} else if len(*p.IncludedLinkRelatedPersonResources) == 1 {
		relatedPerson = &(*p.IncludedLinkRelatedPersonResources)[0]
	}
	return
}

func (p *PersonPlusIncludes) GetIncludedRelatedpersonResource() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedRelatedpersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedRelatedpersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedRelatedpersonResources))
	} else if len(*p.IncludedRelatedpersonResources) == 1 {
		relatedPerson = &(*p.IncludedRelatedpersonResources)[0]
	}
	return
}

func (p *PersonPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if p.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResources))
	} else if len(*p.IncludedPatientResources) == 1 {
		patient = &(*p.IncludedPatientResources)[0]
	}
	return
}

func (p *PersonPlusIncludes) GetIncludedOrganizationResource() (organization *Organization, err error) {
	if p.IncludedOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResources))
	} else if len(*p.IncludedOrganizationResources) == 1 {
		organization = &(*p.IncludedOrganizationResources)[0]
	}
	return
}

func (p *PersonPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResources != nil {
		for _, r := range *p.IncludedPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedLinkPractitionerResources != nil {
		for _, r := range *p.IncludedLinkPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedLinkPatientResources != nil {
		for _, r := range *p.IncludedLinkPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedLinkPersonResources != nil {
		for _, r := range *p.IncludedLinkPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedLinkRelatedPersonResources != nil {
		for _, r := range *p.IncludedLinkRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedRelatedpersonResources != nil {
		for _, r := range *p.IncludedRelatedpersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedPatientResources != nil {
		for _, r := range *p.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedOrganizationResources != nil {
		for _, r := range *p.IncludedOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
