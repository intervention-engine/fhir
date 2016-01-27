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

type OrderResponse struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Request        *Reference    `bson:"request,omitempty" json:"request,omitempty"`
	Date           *FHIRDateTime `bson:"date,omitempty" json:"date,omitempty"`
	Who            *Reference    `bson:"who,omitempty" json:"who,omitempty"`
	OrderStatus    string        `bson:"orderStatus,omitempty" json:"orderStatus,omitempty"`
	Description    string        `bson:"description,omitempty" json:"description,omitempty"`
	Fulfillment    []Reference   `bson:"fulfillment,omitempty" json:"fulfillment,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *OrderResponse) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		OrderResponse
	}{
		ResourceType:  "OrderResponse",
		OrderResponse: *resource,
	}
	return json.Marshal(x)
}

// The "orderResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type orderResponse OrderResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *OrderResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := orderResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = OrderResponse(x2)
	}
	return
}

type OrderResponsePlus struct {
	OrderResponse             `bson:",inline"`
	OrderResponsePlusIncludes `bson:",inline"`
}

type OrderResponsePlusIncludes struct {
	IncludedRequestResources         *[]Order        `bson:"_includedRequestResources,omitempty"`
	IncludedWhoPractitionerResources *[]Practitioner `bson:"_includedWhoPractitionerResources,omitempty"`
	IncludedWhoOrganizationResources *[]Organization `bson:"_includedWhoOrganizationResources,omitempty"`
	IncludedWhoDeviceResources       *[]Device       `bson:"_includedWhoDeviceResources,omitempty"`
}

func (o *OrderResponsePlusIncludes) GetIncludedRequestResource() (order *Order, err error) {
	if o.IncludedRequestResources == nil {
		err = errors.New("Included orders not requested")
	} else if len(*o.IncludedRequestResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 order, but found %d", len(*o.IncludedRequestResources))
	} else if len(*o.IncludedRequestResources) == 1 {
		order = &(*o.IncludedRequestResources)[0]
	}
	return
}

func (o *OrderResponsePlusIncludes) GetIncludedWhoPractitionerResource() (practitioner *Practitioner, err error) {
	if o.IncludedWhoPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*o.IncludedWhoPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*o.IncludedWhoPractitionerResources))
	} else if len(*o.IncludedWhoPractitionerResources) == 1 {
		practitioner = &(*o.IncludedWhoPractitionerResources)[0]
	}
	return
}

func (o *OrderResponsePlusIncludes) GetIncludedWhoOrganizationResource() (organization *Organization, err error) {
	if o.IncludedWhoOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*o.IncludedWhoOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*o.IncludedWhoOrganizationResources))
	} else if len(*o.IncludedWhoOrganizationResources) == 1 {
		organization = &(*o.IncludedWhoOrganizationResources)[0]
	}
	return
}

func (o *OrderResponsePlusIncludes) GetIncludedWhoDeviceResource() (device *Device, err error) {
	if o.IncludedWhoDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedWhoDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedWhoDeviceResources))
	} else if len(*o.IncludedWhoDeviceResources) == 1 {
		device = &(*o.IncludedWhoDeviceResources)[0]
	}
	return
}

func (o *OrderResponsePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedRequestResources != nil {
		for _, r := range *o.IncludedRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedWhoPractitionerResources != nil {
		for _, r := range *o.IncludedWhoPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedWhoOrganizationResources != nil {
		for _, r := range *o.IncludedWhoOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedWhoDeviceResources != nil {
		for _, r := range *o.IncludedWhoDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
