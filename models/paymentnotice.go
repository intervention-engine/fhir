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

type PaymentNotice struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ruleset         *Coding       `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset *Coding       `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created         *FHIRDateTime `bson:"created,omitempty" json:"created,omitempty"`
	Target          *Reference    `bson:"target,omitempty" json:"target,omitempty"`
	Provider        *Reference    `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization    *Reference    `bson:"organization,omitempty" json:"organization,omitempty"`
	Request         *Reference    `bson:"request,omitempty" json:"request,omitempty"`
	Response        *Reference    `bson:"response,omitempty" json:"response,omitempty"`
	PaymentStatus   *Coding       `bson:"paymentStatus,omitempty" json:"paymentStatus,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *PaymentNotice) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		PaymentNotice
	}{
		ResourceType:  "PaymentNotice",
		PaymentNotice: *resource,
	}
	return json.Marshal(x)
}

// The "paymentNotice" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type paymentNotice PaymentNotice

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *PaymentNotice) UnmarshalJSON(data []byte) (err error) {
	x2 := paymentNotice{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = PaymentNotice(x2)
	}
	return
}
