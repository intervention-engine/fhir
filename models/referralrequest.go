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

type ReferralRequest struct {
	Id                    string            `json:"-" bson:"_id"`
	Status                string            `bson:"status"`
	Identifier            []Identifier      `bson:"identifier"`
	Type                  CodeableConcept   `bson:"type"`
	Specialty             CodeableConcept   `bson:"specialty"`
	Priority              CodeableConcept   `bson:"priority"`
	Subject               Reference         `bson:"subject"`
	Requester             Reference         `bson:"requester"`
	Recipient             []Reference       `bson:"recipient"`
	Encounter             Reference         `bson:"encounter"`
	DateSent              FHIRDateTime      `bson:"dateSent"`
	Reason                CodeableConcept   `bson:"reason"`
	Description           string            `bson:"description"`
	ServiceRequested      []CodeableConcept `bson:"serviceRequested"`
	SupportingInformation []Reference       `bson:"supportingInformation"`
	FulfillmentTime       Period            `bson:"fulfillmentTime"`
}
