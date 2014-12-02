// Copyright (c) 2011-2014, HL7, Inc & The MITRE Corporation
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

import "time"

type ReferralRequest struct {
	Id                    string            `json:"-" bson:"_id"`
	Status                string            `bson:"status,omitempty", json:"status,omitempty"`
	Identifier            []Identifier      `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Type                  CodeableConcept   `bson:"type,omitempty", json:"type,omitempty"`
	Specialty             CodeableConcept   `bson:"specialty,omitempty", json:"specialty,omitempty"`
	Priority              CodeableConcept   `bson:"priority,omitempty", json:"priority,omitempty"`
	Subject               Reference         `bson:"subject,omitempty", json:"subject,omitempty"`
	Requester             Reference         `bson:"requester,omitempty", json:"requester,omitempty"`
	Recipient             []Reference       `bson:"recipient,omitempty", json:"recipient,omitempty"`
	Encounter             Reference         `bson:"encounter,omitempty", json:"encounter,omitempty"`
	DateSent              FHIRDateTime      `bson:"dateSent,omitempty", json:"dateSent,omitempty"`
	Reason                CodeableConcept   `bson:"reason,omitempty", json:"reason,omitempty"`
	Description           string            `bson:"description,omitempty", json:"description,omitempty"`
	ServiceRequested      []CodeableConcept `bson:"serviceRequested,omitempty", json:"serviceRequested,omitempty"`
	SupportingInformation []Reference       `bson:"supportingInformation,omitempty", json:"supportingInformation,omitempty"`
	FulfillmentTime       Period            `bson:"fulfillmentTime,omitempty", json:"fulfillmentTime,omitempty"`
}

type ReferralRequestBundle struct {
	Type         string
	Title        string
	Id           string
	Updated      time.Time
	TotalResults int
	Entries      []ReferralRequest
	Category     ReferralRequestCategory
}

type ReferralRequestCategory struct {
	Term   string
	Label  string
	Scheme string
}
