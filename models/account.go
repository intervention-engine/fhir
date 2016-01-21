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

type Account struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name           string           `bson:"name,omitempty" json:"name,omitempty"`
	Type           *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Status         string           `bson:"status,omitempty" json:"status,omitempty"`
	ActivePeriod   *Period          `bson:"activePeriod,omitempty" json:"activePeriod,omitempty"`
	Currency       *Coding          `bson:"currency,omitempty" json:"currency,omitempty"`
	Balance        *Quantity        `bson:"balance,omitempty" json:"balance,omitempty"`
	CoveragePeriod *Period          `bson:"coveragePeriod,omitempty" json:"coveragePeriod,omitempty"`
	Subject        *Reference       `bson:"subject,omitempty" json:"subject,omitempty"`
	Owner          *Reference       `bson:"owner,omitempty" json:"owner,omitempty"`
	Description    string           `bson:"description,omitempty" json:"description,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Account) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Account
	}{
		ResourceType: "Account",
		Account:      *resource,
	}
	return json.Marshal(x)
}

// The "account" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type account Account

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Account) UnmarshalJSON(data []byte) (err error) {
	x2 := account{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Account(x2)
	}
	return
}

type AccountPlus struct {
	Account             `bson:",inline"`
	AccountPlusIncludes `bson:",inline"`
}

type AccountPlusIncludes struct {
	IncludedOwnerResources                    *[]Organization      `bson:"_includedOwnerResources,omitempty"`
	IncludedSubjectPractitionerResources      *[]Practitioner      `bson:"_includedSubjectPractitionerResources,omitempty"`
	IncludedSubjectOrganizationResources      *[]Organization      `bson:"_includedSubjectOrganizationResources,omitempty"`
	IncludedSubjectDeviceResources            *[]Device            `bson:"_includedSubjectDeviceResources,omitempty"`
	IncludedSubjectPatientResources           *[]Patient           `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedSubjectHealthcareServiceResources *[]HealthcareService `bson:"_includedSubjectHealthcareServiceResources,omitempty"`
	IncludedSubjectLocationResources          *[]Location          `bson:"_includedSubjectLocationResources,omitempty"`
	IncludedPatientResources                  *[]Patient           `bson:"_includedPatientResources,omitempty"`
}

func (a *AccountPlusIncludes) GetIncludedOwnerResource() (organization *Organization, err error) {
	if a.IncludedOwnerResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*a.IncludedOwnerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*a.IncludedOwnerResources))
	} else if len(*a.IncludedOwnerResources) == 1 {
		organization = &(*a.IncludedOwnerResources)[0]
	}
	return
}

func (a *AccountPlusIncludes) GetIncludedSubjectPractitionerResource() (practitioner *Practitioner, err error) {
	if a.IncludedSubjectPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedSubjectPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedSubjectPractitionerResources))
	} else if len(*a.IncludedSubjectPractitionerResources) == 1 {
		practitioner = &(*a.IncludedSubjectPractitionerResources)[0]
	}
	return
}

func (a *AccountPlusIncludes) GetIncludedSubjectOrganizationResource() (organization *Organization, err error) {
	if a.IncludedSubjectOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*a.IncludedSubjectOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*a.IncludedSubjectOrganizationResources))
	} else if len(*a.IncludedSubjectOrganizationResources) == 1 {
		organization = &(*a.IncludedSubjectOrganizationResources)[0]
	}
	return
}

func (a *AccountPlusIncludes) GetIncludedSubjectDeviceResource() (device *Device, err error) {
	if a.IncludedSubjectDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedSubjectDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedSubjectDeviceResources))
	} else if len(*a.IncludedSubjectDeviceResources) == 1 {
		device = &(*a.IncludedSubjectDeviceResources)[0]
	}
	return
}

func (a *AccountPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if a.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedSubjectPatientResources))
	} else if len(*a.IncludedSubjectPatientResources) == 1 {
		patient = &(*a.IncludedSubjectPatientResources)[0]
	}
	return
}

func (a *AccountPlusIncludes) GetIncludedSubjectHealthcareServiceResource() (healthcareService *HealthcareService, err error) {
	if a.IncludedSubjectHealthcareServiceResources == nil {
		err = errors.New("Included healthcareservices not requested")
	} else if len(*a.IncludedSubjectHealthcareServiceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 healthcareService, but found %d", len(*a.IncludedSubjectHealthcareServiceResources))
	} else if len(*a.IncludedSubjectHealthcareServiceResources) == 1 {
		healthcareService = &(*a.IncludedSubjectHealthcareServiceResources)[0]
	}
	return
}

func (a *AccountPlusIncludes) GetIncludedSubjectLocationResource() (location *Location, err error) {
	if a.IncludedSubjectLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*a.IncludedSubjectLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*a.IncludedSubjectLocationResources))
	} else if len(*a.IncludedSubjectLocationResources) == 1 {
		location = &(*a.IncludedSubjectLocationResources)[0]
	}
	return
}

func (a *AccountPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if a.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResources))
	} else if len(*a.IncludedPatientResources) == 1 {
		patient = &(*a.IncludedPatientResources)[0]
	}
	return
}

func (a *AccountPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedOwnerResources != nil {
		for _, r := range *a.IncludedOwnerResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedSubjectPractitionerResources != nil {
		for _, r := range *a.IncludedSubjectPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedSubjectOrganizationResources != nil {
		for _, r := range *a.IncludedSubjectOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedSubjectDeviceResources != nil {
		for _, r := range *a.IncludedSubjectDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedSubjectPatientResources != nil {
		for _, r := range *a.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedSubjectHealthcareServiceResources != nil {
		for _, r := range *a.IncludedSubjectHealthcareServiceResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedSubjectLocationResources != nil {
		for _, r := range *a.IncludedSubjectLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPatientResources != nil {
		for _, r := range *a.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
