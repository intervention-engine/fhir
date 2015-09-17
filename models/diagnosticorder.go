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

type DiagnosticOrder struct {
	Id                    string                          `json:"id" bson:"_id"`
	Subject               *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Orderer               *Reference                      `bson:"orderer,omitempty" json:"orderer,omitempty"`
	Identifier            []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Encounter             *Reference                      `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Reason                []CodeableConcept               `bson:"reason,omitempty" json:"reason,omitempty"`
	SupportingInformation []Reference                     `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Specimen              []Reference                     `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Status                string                          `bson:"status,omitempty" json:"status,omitempty"`
	Priority              string                          `bson:"priority,omitempty" json:"priority,omitempty"`
	Event                 []DiagnosticOrderEventComponent `bson:"event,omitempty" json:"event,omitempty"`
	Item                  []DiagnosticOrderItemComponent  `bson:"item,omitempty" json:"item,omitempty"`
	Note                  []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DiagnosticOrder) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		DiagnosticOrder
	}{
		ResourceType:    "DiagnosticOrder",
		DiagnosticOrder: *resource,
	}
	return json.Marshal(x)
}

type DiagnosticOrderEventComponent struct {
	Status      string           `bson:"status,omitempty" json:"status,omitempty"`
	Description *CodeableConcept `bson:"description,omitempty" json:"description,omitempty"`
	DateTime    *FHIRDateTime    `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Actor       *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
}

type DiagnosticOrderItemComponent struct {
	Code     *CodeableConcept                `bson:"code,omitempty" json:"code,omitempty"`
	Specimen []Reference                     `bson:"specimen,omitempty" json:"specimen,omitempty"`
	BodySite *CodeableConcept                `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Status   string                          `bson:"status,omitempty" json:"status,omitempty"`
	Event    []DiagnosticOrderEventComponent `bson:"event,omitempty" json:"event,omitempty"`
}
