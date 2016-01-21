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

type SupplyDelivery struct {
	DomainResource `bson:",inline"`
	Identifier     *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status         string           `bson:"status,omitempty" json:"status,omitempty"`
	Patient        *Reference       `bson:"patient,omitempty" json:"patient,omitempty"`
	Type           *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Quantity       *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	SuppliedItem   *Reference       `bson:"suppliedItem,omitempty" json:"suppliedItem,omitempty"`
	Supplier       *Reference       `bson:"supplier,omitempty" json:"supplier,omitempty"`
	WhenPrepared   *Period          `bson:"whenPrepared,omitempty" json:"whenPrepared,omitempty"`
	Time           *FHIRDateTime    `bson:"time,omitempty" json:"time,omitempty"`
	Destination    *Reference       `bson:"destination,omitempty" json:"destination,omitempty"`
	Receiver       []Reference      `bson:"receiver,omitempty" json:"receiver,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *SupplyDelivery) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		SupplyDelivery
	}{
		ResourceType:   "SupplyDelivery",
		SupplyDelivery: *resource,
	}
	return json.Marshal(x)
}

// The "supplyDelivery" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type supplyDelivery SupplyDelivery

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *SupplyDelivery) UnmarshalJSON(data []byte) (err error) {
	x2 := supplyDelivery{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = SupplyDelivery(x2)
	}
	return
}

type SupplyDeliveryPlus struct {
	SupplyDelivery             `bson:",inline"`
	SupplyDeliveryPlusIncludes `bson:",inline"`
}

type SupplyDeliveryPlusIncludes struct {
	IncludedReceiverResources *[]Practitioner `bson:"_includedReceiverResources,omitempty"`
	IncludedPatientResources  *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedSupplierResources *[]Practitioner `bson:"_includedSupplierResources,omitempty"`
}

func (s *SupplyDeliveryPlusIncludes) GetIncludedReceiverResources() (practitioners []Practitioner, err error) {
	if s.IncludedReceiverResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *s.IncludedReceiverResources
	}
	return
}

func (s *SupplyDeliveryPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if s.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResources))
	} else if len(*s.IncludedPatientResources) == 1 {
		patient = &(*s.IncludedPatientResources)[0]
	}
	return
}

func (s *SupplyDeliveryPlusIncludes) GetIncludedSupplierResource() (practitioner *Practitioner, err error) {
	if s.IncludedSupplierResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*s.IncludedSupplierResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*s.IncludedSupplierResources))
	} else if len(*s.IncludedSupplierResources) == 1 {
		practitioner = &(*s.IncludedSupplierResources)[0]
	}
	return
}

func (s *SupplyDeliveryPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedReceiverResources != nil {
		for _, r := range *s.IncludedReceiverResources {
			resourceMap[r.Id] = &r
		}
	}
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
	return resourceMap
}
