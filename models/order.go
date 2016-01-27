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

type Order struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Date                  *FHIRDateTime       `bson:"date,omitempty" json:"date,omitempty"`
	Subject               *Reference          `bson:"subject,omitempty" json:"subject,omitempty"`
	Source                *Reference          `bson:"source,omitempty" json:"source,omitempty"`
	Target                *Reference          `bson:"target,omitempty" json:"target,omitempty"`
	ReasonCodeableConcept *CodeableConcept    `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference          `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	When                  *OrderWhenComponent `bson:"when,omitempty" json:"when,omitempty"`
	Detail                []Reference         `bson:"detail,omitempty" json:"detail,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Order) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Order
	}{
		ResourceType: "Order",
		Order:        *resource,
	}
	return json.Marshal(x)
}

// The "order" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type order Order

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Order) UnmarshalJSON(data []byte) (err error) {
	x2 := order{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Order(x2)
	}
	return
}

type OrderWhenComponent struct {
	Code     *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Schedule *Timing          `bson:"schedule,omitempty" json:"schedule,omitempty"`
}

type OrderPlus struct {
	Order             `bson:",inline"`
	OrderPlusIncludes `bson:",inline"`
}

type OrderPlusIncludes struct {
	IncludedSubjectGroupResources       *[]Group        `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectDeviceResources      *[]Device       `bson:"_includedSubjectDeviceResources,omitempty"`
	IncludedSubjectPatientResources     *[]Patient      `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedSubjectSubstanceResources   *[]Substance    `bson:"_includedSubjectSubstanceResources,omitempty"`
	IncludedPatientResources            *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedSourcePractitionerResources *[]Practitioner `bson:"_includedSourcePractitionerResources,omitempty"`
	IncludedSourceOrganizationResources *[]Organization `bson:"_includedSourceOrganizationResources,omitempty"`
	IncludedTargetPractitionerResources *[]Practitioner `bson:"_includedTargetPractitionerResources,omitempty"`
	IncludedTargetOrganizationResources *[]Organization `bson:"_includedTargetOrganizationResources,omitempty"`
	IncludedTargetDeviceResources       *[]Device       `bson:"_includedTargetDeviceResources,omitempty"`
}

func (o *OrderPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if o.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*o.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*o.IncludedSubjectGroupResources))
	} else if len(*o.IncludedSubjectGroupResources) == 1 {
		group = &(*o.IncludedSubjectGroupResources)[0]
	}
	return
}

func (o *OrderPlusIncludes) GetIncludedSubjectDeviceResource() (device *Device, err error) {
	if o.IncludedSubjectDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedSubjectDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedSubjectDeviceResources))
	} else if len(*o.IncludedSubjectDeviceResources) == 1 {
		device = &(*o.IncludedSubjectDeviceResources)[0]
	}
	return
}

func (o *OrderPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if o.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedSubjectPatientResources))
	} else if len(*o.IncludedSubjectPatientResources) == 1 {
		patient = &(*o.IncludedSubjectPatientResources)[0]
	}
	return
}

func (o *OrderPlusIncludes) GetIncludedSubjectSubstanceResource() (substance *Substance, err error) {
	if o.IncludedSubjectSubstanceResources == nil {
		err = errors.New("Included substances not requested")
	} else if len(*o.IncludedSubjectSubstanceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*o.IncludedSubjectSubstanceResources))
	} else if len(*o.IncludedSubjectSubstanceResources) == 1 {
		substance = &(*o.IncludedSubjectSubstanceResources)[0]
	}
	return
}

func (o *OrderPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if o.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedPatientResources))
	} else if len(*o.IncludedPatientResources) == 1 {
		patient = &(*o.IncludedPatientResources)[0]
	}
	return
}

func (o *OrderPlusIncludes) GetIncludedSourcePractitionerResource() (practitioner *Practitioner, err error) {
	if o.IncludedSourcePractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*o.IncludedSourcePractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*o.IncludedSourcePractitionerResources))
	} else if len(*o.IncludedSourcePractitionerResources) == 1 {
		practitioner = &(*o.IncludedSourcePractitionerResources)[0]
	}
	return
}

func (o *OrderPlusIncludes) GetIncludedSourceOrganizationResource() (organization *Organization, err error) {
	if o.IncludedSourceOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*o.IncludedSourceOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*o.IncludedSourceOrganizationResources))
	} else if len(*o.IncludedSourceOrganizationResources) == 1 {
		organization = &(*o.IncludedSourceOrganizationResources)[0]
	}
	return
}

func (o *OrderPlusIncludes) GetIncludedTargetPractitionerResource() (practitioner *Practitioner, err error) {
	if o.IncludedTargetPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*o.IncludedTargetPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*o.IncludedTargetPractitionerResources))
	} else if len(*o.IncludedTargetPractitionerResources) == 1 {
		practitioner = &(*o.IncludedTargetPractitionerResources)[0]
	}
	return
}

func (o *OrderPlusIncludes) GetIncludedTargetOrganizationResource() (organization *Organization, err error) {
	if o.IncludedTargetOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*o.IncludedTargetOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*o.IncludedTargetOrganizationResources))
	} else if len(*o.IncludedTargetOrganizationResources) == 1 {
		organization = &(*o.IncludedTargetOrganizationResources)[0]
	}
	return
}

func (o *OrderPlusIncludes) GetIncludedTargetDeviceResource() (device *Device, err error) {
	if o.IncludedTargetDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedTargetDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedTargetDeviceResources))
	} else if len(*o.IncludedTargetDeviceResources) == 1 {
		device = &(*o.IncludedTargetDeviceResources)[0]
	}
	return
}

func (o *OrderPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedSubjectGroupResources != nil {
		for _, r := range *o.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedSubjectDeviceResources != nil {
		for _, r := range *o.IncludedSubjectDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedSubjectPatientResources != nil {
		for _, r := range *o.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedSubjectSubstanceResources != nil {
		for _, r := range *o.IncludedSubjectSubstanceResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPatientResources != nil {
		for _, r := range *o.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedSourcePractitionerResources != nil {
		for _, r := range *o.IncludedSourcePractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedSourceOrganizationResources != nil {
		for _, r := range *o.IncludedSourceOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedTargetPractitionerResources != nil {
		for _, r := range *o.IncludedTargetPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedTargetOrganizationResources != nil {
		for _, r := range *o.IncludedTargetOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedTargetDeviceResources != nil {
		for _, r := range *o.IncludedTargetDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
