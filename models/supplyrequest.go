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

type SupplyRequest struct {
	DomainResource        `bson:",inline"`
	Patient               *Reference                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Source                *Reference                  `bson:"source,omitempty" json:"source,omitempty"`
	Date                  *FHIRDateTime               `bson:"date,omitempty" json:"date,omitempty"`
	Identifier            *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                string                      `bson:"status,omitempty" json:"status,omitempty"`
	Kind                  *CodeableConcept            `bson:"kind,omitempty" json:"kind,omitempty"`
	OrderedItem           *Reference                  `bson:"orderedItem,omitempty" json:"orderedItem,omitempty"`
	Supplier              []Reference                 `bson:"supplier,omitempty" json:"supplier,omitempty"`
	ReasonCodeableConcept *CodeableConcept            `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                  `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	When                  *SupplyRequestWhenComponent `bson:"when,omitempty" json:"when,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *SupplyRequest) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		SupplyRequest
	}{
		ResourceType:  "SupplyRequest",
		SupplyRequest: *resource,
	}
	return json.Marshal(x)
}

// The "supplyRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type supplyRequest SupplyRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *SupplyRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := supplyRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = SupplyRequest(x2)
	}
	return
}

type SupplyRequestWhenComponent struct {
	Code     *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Schedule *Timing          `bson:"schedule,omitempty" json:"schedule,omitempty"`
}

type SupplyRequestPlus struct {
	SupplyRequest             `bson:",inline"`
	SupplyRequestPlusIncludes `bson:",inline"`
}

type SupplyRequestPlusIncludes struct {
	IncludedPatientResources            *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedSupplierResources           *[]Organization `bson:"_includedSupplierResources,omitempty"`
	IncludedSourcePractitionerResources *[]Practitioner `bson:"_includedSourcePractitionerResources,omitempty"`
	IncludedSourceOrganizationResources *[]Organization `bson:"_includedSourceOrganizationResources,omitempty"`
	IncludedSourcePatientResources      *[]Patient      `bson:"_includedSourcePatientResources,omitempty"`
}

func (s *SupplyRequestPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if s.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResources))
	} else if len(*s.IncludedPatientResources) == 1 {
		patient = &(*s.IncludedPatientResources)[0]
	}
	return
}

func (s *SupplyRequestPlusIncludes) GetIncludedSupplierResources() (organizations []Organization, err error) {
	if s.IncludedSupplierResources == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *s.IncludedSupplierResources
	}
	return
}

func (s *SupplyRequestPlusIncludes) GetIncludedSourcePractitionerResource() (practitioner *Practitioner, err error) {
	if s.IncludedSourcePractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*s.IncludedSourcePractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*s.IncludedSourcePractitionerResources))
	} else if len(*s.IncludedSourcePractitionerResources) == 1 {
		practitioner = &(*s.IncludedSourcePractitionerResources)[0]
	}
	return
}

func (s *SupplyRequestPlusIncludes) GetIncludedSourceOrganizationResource() (organization *Organization, err error) {
	if s.IncludedSourceOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*s.IncludedSourceOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*s.IncludedSourceOrganizationResources))
	} else if len(*s.IncludedSourceOrganizationResources) == 1 {
		organization = &(*s.IncludedSourceOrganizationResources)[0]
	}
	return
}

func (s *SupplyRequestPlusIncludes) GetIncludedSourcePatientResource() (patient *Patient, err error) {
	if s.IncludedSourcePatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedSourcePatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedSourcePatientResources))
	} else if len(*s.IncludedSourcePatientResources) == 1 {
		patient = &(*s.IncludedSourcePatientResources)[0]
	}
	return
}

func (s *SupplyRequestPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedPatientResources != nil {
		for _, r := range *s.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedSupplierResources != nil {
		for _, r := range *s.IncludedSupplierResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedSourcePractitionerResources != nil {
		for _, r := range *s.IncludedSourcePractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedSourceOrganizationResources != nil {
		for _, r := range *s.IncludedSourceOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedSourcePatientResources != nil {
		for _, r := range *s.IncludedSourcePatientResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
