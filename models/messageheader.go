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

type MessageHeader struct {
	Id          string                         `json:"-" bson:"_id"`
	Identifier  string                         `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Timestamp   FHIRDateTime                   `bson:"timestamp,omitempty", json:"timestamp,omitempty"`
	Event       Coding                         `bson:"event,omitempty", json:"event,omitempty"`
	Response    MessageHeaderResponseComponent `bson:"response,omitempty", json:"response,omitempty"`
	Source      MessageSourceComponent         `bson:"source,omitempty", json:"source,omitempty"`
	Destination []MessageDestinationComponent  `bson:"destination,omitempty", json:"destination,omitempty"`
	Enterer     Reference                      `bson:"enterer,omitempty", json:"enterer,omitempty"`
	Author      Reference                      `bson:"author,omitempty", json:"author,omitempty"`
	Receiver    Reference                      `bson:"receiver,omitempty", json:"receiver,omitempty"`
	Responsible Reference                      `bson:"responsible,omitempty", json:"responsible,omitempty"`
	Reason      CodeableConcept                `bson:"reason,omitempty", json:"reason,omitempty"`
	Data        []Reference                    `bson:"data,omitempty", json:"data,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec response
type MessageHeaderResponseComponent struct {
	Identifier string    `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Code       string    `bson:"code,omitempty", json:"code,omitempty"`
	Details    Reference `bson:"details,omitempty", json:"details,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec source
type MessageSourceComponent struct {
	Name     string       `bson:"name,omitempty", json:"name,omitempty"`
	Software string       `bson:"software,omitempty", json:"software,omitempty"`
	Version  string       `bson:"version,omitempty", json:"version,omitempty"`
	Contact  ContactPoint `bson:"contact,omitempty", json:"contact,omitempty"`
	Endpoint string       `bson:"endpoint,omitempty", json:"endpoint,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec destination
type MessageDestinationComponent struct {
	Name     string    `bson:"name,omitempty", json:"name,omitempty"`
	Target   Reference `bson:"target,omitempty", json:"target,omitempty"`
	Endpoint string    `bson:"endpoint,omitempty", json:"endpoint,omitempty"`
}

type MessageHeaderBundle struct {
	Type         string
	Title        string
	Id           string
	Updated      time.Time
	TotalResults int
	Entries      []MessageHeader
	Category     MessageHeaderCategory
}

type MessageHeaderCategory struct {
	Term   string
	Label  string
	Scheme string
}
