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

type Patient struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active               *bool                           `bson:"active,omitempty" json:"active,omitempty"`
	Name                 []HumanName                     `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []ContactPoint                  `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               string                          `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *FHIRDateTime                   `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                           `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *FHIRDateTime                   `bson:"deceasedDateTime,omitempty" json:"deceasedDateTime,omitempty"`
	Address              []Address                       `bson:"address,omitempty" json:"address,omitempty"`
	MaritalStatus        *CodeableConcept                `bson:"maritalStatus,omitempty" json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *bool                           `bson:"multipleBirthBoolean,omitempty" json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int32                          `bson:"multipleBirthInteger,omitempty" json:"multipleBirthInteger,omitempty"`
	Photo                []Attachment                    `bson:"photo,omitempty" json:"photo,omitempty"`
	Contact              []PatientContactComponent       `bson:"contact,omitempty" json:"contact,omitempty"`
	Animal               *PatientAnimalComponent         `bson:"animal,omitempty" json:"animal,omitempty"`
	Communication        []PatientCommunicationComponent `bson:"communication,omitempty" json:"communication,omitempty"`
	CareProvider         []Reference                     `bson:"careProvider,omitempty" json:"careProvider,omitempty"`
	ManagingOrganization *Reference                      `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Link                 []PatientLinkComponent          `bson:"link,omitempty" json:"link,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Patient) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Patient
	}{
		ResourceType: "Patient",
		Patient:      *resource,
	}
	return json.Marshal(x)
}

// The "patient" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type patient Patient

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Patient) UnmarshalJSON(data []byte) (err error) {
	x2 := patient{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Patient(x2)
	}
	return
}

type PatientContactComponent struct {
	Relationship []CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name         *HumanName        `bson:"name,omitempty" json:"name,omitempty"`
	Telecom      []ContactPoint    `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address      *Address          `bson:"address,omitempty" json:"address,omitempty"`
	Gender       string            `bson:"gender,omitempty" json:"gender,omitempty"`
	Organization *Reference        `bson:"organization,omitempty" json:"organization,omitempty"`
	Period       *Period           `bson:"period,omitempty" json:"period,omitempty"`
}

type PatientAnimalComponent struct {
	Species      *CodeableConcept `bson:"species,omitempty" json:"species,omitempty"`
	Breed        *CodeableConcept `bson:"breed,omitempty" json:"breed,omitempty"`
	GenderStatus *CodeableConcept `bson:"genderStatus,omitempty" json:"genderStatus,omitempty"`
}

type PatientCommunicationComponent struct {
	Language  *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	Preferred *bool            `bson:"preferred,omitempty" json:"preferred,omitempty"`
}

type PatientLinkComponent struct {
	Other *Reference `bson:"other,omitempty" json:"other,omitempty"`
	Type  string     `bson:"type,omitempty" json:"type,omitempty"`
}

type PatientPlus struct {
	Patient             `bson:",inline"`
	PatientPlusIncludes `bson:",inline"`
}

type PatientPlusIncludes struct {
	IncludedLinkResources                     *[]Patient      `bson:"_includedLinkResources,omitempty"`
	IncludedCareproviderPractitionerResources *[]Practitioner `bson:"_includedCareproviderPractitionerResources,omitempty"`
	IncludedCareproviderOrganizationResources *[]Organization `bson:"_includedCareproviderOrganizationResources,omitempty"`
	IncludedOrganizationResources             *[]Organization `bson:"_includedOrganizationResources,omitempty"`
}

func (p *PatientPlusIncludes) GetIncludedLinkResource() (patient *Patient, err error) {
	if p.IncludedLinkResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedLinkResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedLinkResources))
	} else if len(*p.IncludedLinkResources) == 1 {
		patient = &(*p.IncludedLinkResources)[0]
	}
	return
}

func (p *PatientPlusIncludes) GetIncludedCareproviderPractitionerResources() (practitioners []Practitioner, err error) {
	if p.IncludedCareproviderPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *p.IncludedCareproviderPractitionerResources
	}
	return
}

func (p *PatientPlusIncludes) GetIncludedCareproviderOrganizationResources() (organizations []Organization, err error) {
	if p.IncludedCareproviderOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *p.IncludedCareproviderOrganizationResources
	}
	return
}

func (p *PatientPlusIncludes) GetIncludedOrganizationResource() (organization *Organization, err error) {
	if p.IncludedOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResources))
	} else if len(*p.IncludedOrganizationResources) == 1 {
		organization = &(*p.IncludedOrganizationResources)[0]
	}
	return
}

func (p *PatientPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedLinkResources != nil {
		for _, r := range *p.IncludedLinkResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedCareproviderPractitionerResources != nil {
		for _, r := range *p.IncludedCareproviderPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedCareproviderOrganizationResources != nil {
		for _, r := range *p.IncludedCareproviderOrganizationResources {
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
