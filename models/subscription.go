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

import "time"

type Subscription struct {
	Id       string                        `json:"-" bson:"_id"`
	Criteria string                        `bson:"criteria,omitempty" json:"criteria,omitempty"`
	Contact  []ContactPoint                `bson:"contact,omitempty" json:"contact,omitempty"`
	Reason   string                        `bson:"reason,omitempty" json:"reason,omitempty"`
	Status   string                        `bson:"status,omitempty" json:"status,omitempty"`
	Error    string                        `bson:"error,omitempty" json:"error,omitempty"`
	Channel  *SubscriptionChannelComponent `bson:"channel,omitempty" json:"channel,omitempty"`
	End      *FHIRDateTime                 `bson:"end,omitempty" json:"end,omitempty"`
	Tag      []Coding                      `bson:"tag,omitempty" json:"tag,omitempty"`
}
type SubscriptionChannelComponent struct {
	Type     string `bson:"type,omitempty" json:"type,omitempty"`
	Endpoint string `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	Payload  string `bson:"payload,omitempty" json:"payload,omitempty"`
	Header   string `bson:"header,omitempty" json:"header,omitempty"`
}

type SubscriptionBundle struct {
	Type         string                    `json:"resourceType,omitempty"`
	Title        string                    `json:"title,omitempty"`
	Id           string                    `json:"id,omitempty"`
	Updated      time.Time                 `json:"updated,omitempty"`
	TotalResults int                       `json:"totalResults,omitempty"`
	Entry        []SubscriptionBundleEntry `json:"entry,omitempty"`
	Category     SubscriptionCategory      `json:"category,omitempty"`
}

type SubscriptionBundleEntry struct {
	Title    string               `json:"title,omitempty"`
	Id       string               `json:"id,omitempty"`
	Content  Subscription         `json:"content,omitempty"`
	Category SubscriptionCategory `json:"category,omitempty"`
}

type SubscriptionCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
