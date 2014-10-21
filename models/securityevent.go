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

type SecurityEvent struct {
	Id          string                              `json:"-" bson:"_id"`
	Event       SecurityEventEventComponent         `bson:"event"`
	Participant []SecurityEventParticipantComponent `bson:"participant"`
	Source      SecurityEventSourceComponent        `bson:"source"`
	Object      []SecurityEventObjectComponent      `bson:"object"`
}

// This is an ugly hack to deal with embedded structures in the spec event
type SecurityEventEventComponent struct {
	Type        CodeableConcept   `bson:"type"`
	Subtype     []CodeableConcept `bson:"subtype"`
	Action      string            `bson:"action"`
	DateTime    time.Time         `bson:"dateTime"`
	Outcome     string            `bson:"outcome"`
	OutcomeDesc string            `bson:"outcomeDesc"`
}

// This is an ugly hack to deal with embedded structures in the spec network
type SecurityEventParticipantNetworkComponent struct {
	Identifier string `bson:"identifier"`
	Type       string `bson:"type"`
}

// This is an ugly hack to deal with embedded structures in the spec participant
type SecurityEventParticipantComponent struct {
	Role      []CodeableConcept                        `bson:"role"`
	Reference Reference                                `bson:"reference"`
	UserId    string                                   `bson:"userId"`
	AltId     string                                   `bson:"altId"`
	Name      string                                   `bson:"name"`
	Requestor bool                                     `bson:"requestor"`
	Media     Coding                                   `bson:"media"`
	Network   SecurityEventParticipantNetworkComponent `bson:"network"`
}

// This is an ugly hack to deal with embedded structures in the spec source
type SecurityEventSourceComponent struct {
	Site       string   `bson:"site"`
	Identifier string   `bson:"identifier"`
	Type       []Coding `bson:"type"`
}

// This is an ugly hack to deal with embedded structures in the spec detail
type SecurityEventObjectDetailComponent struct {
	Type  string `bson:"type"`
	Value string `bson:"value"`
}

// This is an ugly hack to deal with embedded structures in the spec object
type SecurityEventObjectComponent struct {
	Identifier  Identifier                           `bson:"identifier"`
	Reference   Reference                            `bson:"reference"`
	Type        string                               `bson:"type"`
	Role        string                               `bson:"role"`
	Lifecycle   string                               `bson:"lifecycle"`
	Sensitivity CodeableConcept                      `bson:"sensitivity"`
	Name        string                               `bson:"name"`
	Description string                               `bson:"description"`
	Query       string                               `bson:"query"`
	Detail      []SecurityEventObjectDetailComponent `bson:"detail"`
}
