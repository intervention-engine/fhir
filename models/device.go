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

import "encoding/json"

type Device struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Note            []Annotation     `bson:"note,omitempty" json:"note,omitempty"`
	Status          string           `bson:"status,omitempty" json:"status,omitempty"`
	Manufacturer    string           `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Model           string           `bson:"model,omitempty" json:"model,omitempty"`
	Version         string           `bson:"version,omitempty" json:"version,omitempty"`
	ManufactureDate *FHIRDateTime    `bson:"manufactureDate,omitempty" json:"manufactureDate,omitempty"`
	Expiry          *FHIRDateTime    `bson:"expiry,omitempty" json:"expiry,omitempty"`
	Udi             string           `bson:"udi,omitempty" json:"udi,omitempty"`
	LotNumber       string           `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	Owner           *Reference       `bson:"owner,omitempty" json:"owner,omitempty"`
	Location        *Reference       `bson:"location,omitempty" json:"location,omitempty"`
	Patient         *Reference       `bson:"patient,omitempty" json:"patient,omitempty"`
	Contact         []ContactPoint   `bson:"contact,omitempty" json:"contact,omitempty"`
	Url             string           `bson:"url,omitempty" json:"url,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Device) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Device
	}{
		ResourceType: "Device",
		Device:       *resource,
	}
	return json.Marshal(x)
}

// The "device" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type device Device

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Device) UnmarshalJSON(data []byte) (err error) {
	x2 := device{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Device(x2)
	}
	return
}
