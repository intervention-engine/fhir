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

type MessageHeader struct {
	Id          string                                     `json:"-" bson:"_id"`
	Identifier  string                                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Timestamp   *FHIRDateTime                              `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Event       *Coding                                    `bson:"event,omitempty" json:"event,omitempty"`
	Response    *MessageHeaderResponseComponent            `bson:"response,omitempty" json:"response,omitempty"`
	Source      *MessageHeaderMessageSourceComponent       `bson:"source,omitempty" json:"source,omitempty"`
	Destination []MessageHeaderMessageDestinationComponent `bson:"destination,omitempty" json:"destination,omitempty"`
	Enterer     *Reference                                 `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Author      *Reference                                 `bson:"author,omitempty" json:"author,omitempty"`
	Receiver    *Reference                                 `bson:"receiver,omitempty" json:"receiver,omitempty"`
	Responsible *Reference                                 `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Reason      *CodeableConcept                           `bson:"reason,omitempty" json:"reason,omitempty"`
	Data        []Reference                                `bson:"data,omitempty" json:"data,omitempty"`
}

type MessageHeaderResponseComponent struct {
	Identifier string     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code       string     `bson:"code,omitempty" json:"code,omitempty"`
	Details    *Reference `bson:"details,omitempty" json:"details,omitempty"`
}

type MessageHeaderMessageSourceComponent struct {
	Name     string        `bson:"name,omitempty" json:"name,omitempty"`
	Software string        `bson:"software,omitempty" json:"software,omitempty"`
	Version  string        `bson:"version,omitempty" json:"version,omitempty"`
	Contact  *ContactPoint `bson:"contact,omitempty" json:"contact,omitempty"`
	Endpoint string        `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

type MessageHeaderMessageDestinationComponent struct {
	Name     string     `bson:"name,omitempty" json:"name,omitempty"`
	Target   *Reference `bson:"target,omitempty" json:"target,omitempty"`
	Endpoint string     `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

type MessageHeaderBundle struct {
	Id    string                     `json:"id,omitempty"`
	Type  string                     `json:"resourceType,omitempty"`
	Base  string                     `json:"base,omitempty"`
	Total int                        `json:"total,omitempty"`
	Link  []BundleLinkComponent      `json:"link,omitempty"`
	Entry []MessageHeaderBundleEntry `json:"entry,omitempty"`
}

type MessageHeaderBundleEntry struct {
	Id       string                `json:"id,omitempty"`
	Base     string                `json:"base,omitempty"`
	Link     []BundleLinkComponent `json:"link,omitempty"`
	Resource MessageHeader         `json:"resource,omitempty"`
}

func (resource *MessageHeader) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		MessageHeader
	}{
		ResourceType:  "MessageHeader",
		MessageHeader: *resource,
	}
	return json.Marshal(x)
}
