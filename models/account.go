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

type Account struct {
	Id             string           `json:"id" bson:"_id"`
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
