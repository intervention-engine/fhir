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

type Flag struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Category       *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Status         string           `bson:"status,omitempty" json:"status,omitempty"`
	Period         *Period          `bson:"period,omitempty" json:"period,omitempty"`
	Subject        *Reference       `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter      *Reference       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Author         *Reference       `bson:"author,omitempty" json:"author,omitempty"`
	Code           *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Flag) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Flag"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Flag), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Flag) GetBSON() (interface{}, error) {
	x.ResourceType = "Flag"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "flag" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type flag Flag

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Flag) UnmarshalJSON(data []byte) (err error) {
	x2 := flag{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Flag(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Flag) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Flag"
	} else if x.ResourceType != "Flag" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Flag, instead received %s", x.ResourceType))
	}
	return nil
}

type FlagPlus struct {
	Flag             `bson:",inline"`
	FlagPlusIncludes `bson:",inline"`
}

type FlagPlusIncludes struct {
	IncludedSubjectPractitionerResources *[]Practitioner `bson:"_includedSubjectPractitionerResources,omitempty"`
	IncludedSubjectGroupResources        *[]Group        `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectOrganizationResources *[]Organization `bson:"_includedSubjectOrganizationResources,omitempty"`
	IncludedSubjectPatientResources      *[]Patient      `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedSubjectLocationResources     *[]Location     `bson:"_includedSubjectLocationResources,omitempty"`
	IncludedPatientResources             *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedAuthorPractitionerResources  *[]Practitioner `bson:"_includedAuthorPractitionerResources,omitempty"`
	IncludedAuthorOrganizationResources  *[]Organization `bson:"_includedAuthorOrganizationResources,omitempty"`
	IncludedAuthorDeviceResources        *[]Device       `bson:"_includedAuthorDeviceResources,omitempty"`
	IncludedAuthorPatientResources       *[]Patient      `bson:"_includedAuthorPatientResources,omitempty"`
	IncludedEncounterResources           *[]Encounter    `bson:"_includedEncounterResources,omitempty"`
}

func (f *FlagPlusIncludes) GetIncludedSubjectPractitionerResource() (practitioner *Practitioner, err error) {
	if f.IncludedSubjectPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*f.IncludedSubjectPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*f.IncludedSubjectPractitionerResources))
	} else if len(*f.IncludedSubjectPractitionerResources) == 1 {
		practitioner = &(*f.IncludedSubjectPractitionerResources)[0]
	}
	return
}

func (f *FlagPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if f.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*f.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*f.IncludedSubjectGroupResources))
	} else if len(*f.IncludedSubjectGroupResources) == 1 {
		group = &(*f.IncludedSubjectGroupResources)[0]
	}
	return
}

func (f *FlagPlusIncludes) GetIncludedSubjectOrganizationResource() (organization *Organization, err error) {
	if f.IncludedSubjectOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*f.IncludedSubjectOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*f.IncludedSubjectOrganizationResources))
	} else if len(*f.IncludedSubjectOrganizationResources) == 1 {
		organization = &(*f.IncludedSubjectOrganizationResources)[0]
	}
	return
}

func (f *FlagPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if f.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedSubjectPatientResources))
	} else if len(*f.IncludedSubjectPatientResources) == 1 {
		patient = &(*f.IncludedSubjectPatientResources)[0]
	}
	return
}

func (f *FlagPlusIncludes) GetIncludedSubjectLocationResource() (location *Location, err error) {
	if f.IncludedSubjectLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*f.IncludedSubjectLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*f.IncludedSubjectLocationResources))
	} else if len(*f.IncludedSubjectLocationResources) == 1 {
		location = &(*f.IncludedSubjectLocationResources)[0]
	}
	return
}

func (f *FlagPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if f.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedPatientResources))
	} else if len(*f.IncludedPatientResources) == 1 {
		patient = &(*f.IncludedPatientResources)[0]
	}
	return
}

func (f *FlagPlusIncludes) GetIncludedAuthorPractitionerResource() (practitioner *Practitioner, err error) {
	if f.IncludedAuthorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*f.IncludedAuthorPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*f.IncludedAuthorPractitionerResources))
	} else if len(*f.IncludedAuthorPractitionerResources) == 1 {
		practitioner = &(*f.IncludedAuthorPractitionerResources)[0]
	}
	return
}

func (f *FlagPlusIncludes) GetIncludedAuthorOrganizationResource() (organization *Organization, err error) {
	if f.IncludedAuthorOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*f.IncludedAuthorOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*f.IncludedAuthorOrganizationResources))
	} else if len(*f.IncludedAuthorOrganizationResources) == 1 {
		organization = &(*f.IncludedAuthorOrganizationResources)[0]
	}
	return
}

func (f *FlagPlusIncludes) GetIncludedAuthorDeviceResource() (device *Device, err error) {
	if f.IncludedAuthorDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*f.IncludedAuthorDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*f.IncludedAuthorDeviceResources))
	} else if len(*f.IncludedAuthorDeviceResources) == 1 {
		device = &(*f.IncludedAuthorDeviceResources)[0]
	}
	return
}

func (f *FlagPlusIncludes) GetIncludedAuthorPatientResource() (patient *Patient, err error) {
	if f.IncludedAuthorPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedAuthorPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedAuthorPatientResources))
	} else if len(*f.IncludedAuthorPatientResources) == 1 {
		patient = &(*f.IncludedAuthorPatientResources)[0]
	}
	return
}

func (f *FlagPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if f.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*f.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*f.IncludedEncounterResources))
	} else if len(*f.IncludedEncounterResources) == 1 {
		encounter = &(*f.IncludedEncounterResources)[0]
	}
	return
}

func (f *FlagPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedSubjectPractitionerResources != nil {
		for _, r := range *f.IncludedSubjectPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedSubjectGroupResources != nil {
		for _, r := range *f.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedSubjectOrganizationResources != nil {
		for _, r := range *f.IncludedSubjectOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedSubjectPatientResources != nil {
		for _, r := range *f.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedSubjectLocationResources != nil {
		for _, r := range *f.IncludedSubjectLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedPatientResources != nil {
		for _, r := range *f.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedAuthorPractitionerResources != nil {
		for _, r := range *f.IncludedAuthorPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedAuthorOrganizationResources != nil {
		for _, r := range *f.IncludedAuthorOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedAuthorDeviceResources != nil {
		for _, r := range *f.IncludedAuthorDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedAuthorPatientResources != nil {
		for _, r := range *f.IncludedAuthorPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedEncounterResources != nil {
		for _, r := range *f.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
