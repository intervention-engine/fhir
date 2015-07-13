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

type AuditEvent struct {
	Id          string                           `json:"-" bson:"_id"`
	Event       *AuditEventEventComponent        `bson:"event,omitempty" json:"event,omitempty"`
	Participant []AuditEventParticipantComponent `bson:"participant,omitempty" json:"participant,omitempty"`
	Source      *AuditEventSourceComponent       `bson:"source,omitempty" json:"source,omitempty"`
	Object      []AuditEventObjectComponent      `bson:"object,omitempty" json:"object,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *AuditEvent) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		AuditEvent
	}{
		ResourceType: "AuditEvent",
		AuditEvent:   *resource,
	}
	return json.Marshal(x)
}

type AuditEventEventComponent struct {
	Type           *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Subtype        []CodeableConcept `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Action         string            `bson:"action,omitempty" json:"action,omitempty"`
	DateTime       *FHIRDateTime     `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Outcome        string            `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OutcomeDesc    string            `bson:"outcomeDesc,omitempty" json:"outcomeDesc,omitempty"`
	PurposeOfEvent []Coding          `bson:"purposeOfEvent,omitempty" json:"purposeOfEvent,omitempty"`
}

type AuditEventParticipantComponent struct {
	Role         []CodeableConcept                      `bson:"role,omitempty" json:"role,omitempty"`
	Reference    *Reference                             `bson:"reference,omitempty" json:"reference,omitempty"`
	UserId       string                                 `bson:"userId,omitempty" json:"userId,omitempty"`
	AltId        string                                 `bson:"altId,omitempty" json:"altId,omitempty"`
	Name         string                                 `bson:"name,omitempty" json:"name,omitempty"`
	Requestor    *bool                                  `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Location     *Reference                             `bson:"location,omitempty" json:"location,omitempty"`
	Policy       []string                               `bson:"policy,omitempty" json:"policy,omitempty"`
	Media        *Coding                                `bson:"media,omitempty" json:"media,omitempty"`
	Network      *AuditEventParticipantNetworkComponent `bson:"network,omitempty" json:"network,omitempty"`
	PurposeOfUse []Coding                               `bson:"purposeOfUse,omitempty" json:"purposeOfUse,omitempty"`
}

type AuditEventParticipantNetworkComponent struct {
	Identifier string `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type       string `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventSourceComponent struct {
	Site       string   `bson:"site,omitempty" json:"site,omitempty"`
	Identifier string   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type       []Coding `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventObjectComponent struct {
	Identifier  *Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Reference   *Reference                        `bson:"reference,omitempty" json:"reference,omitempty"`
	Type        string                            `bson:"type,omitempty" json:"type,omitempty"`
	Role        string                            `bson:"role,omitempty" json:"role,omitempty"`
	Lifecycle   string                            `bson:"lifecycle,omitempty" json:"lifecycle,omitempty"`
	Sensitivity *CodeableConcept                  `bson:"sensitivity,omitempty" json:"sensitivity,omitempty"`
	Name        string                            `bson:"name,omitempty" json:"name,omitempty"`
	Description string                            `bson:"description,omitempty" json:"description,omitempty"`
	Query       string                            `bson:"query,omitempty" json:"query,omitempty"`
	Detail      []AuditEventObjectDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type AuditEventObjectDetailComponent struct {
	Type  string `bson:"type,omitempty" json:"type,omitempty"`
	Value string `bson:"value,omitempty" json:"value,omitempty"`
}
