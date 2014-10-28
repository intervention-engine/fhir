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

type Subscription struct {
	Id       string                       `json:"-" bson:"_id"`
	Criteria string                       `bson:"criteria,omitempty", json:"criteria,omitempty"`
	Contact  []ContactPoint               `bson:"contact,omitempty", json:"contact,omitempty"`
	Reason   string                       `bson:"reason,omitempty", json:"reason,omitempty"`
	Status   string                       `bson:"status,omitempty", json:"status,omitempty"`
	Error    string                       `bson:"error,omitempty", json:"error,omitempty"`
	Channel  SubscriptionChannelComponent `bson:"channel,omitempty", json:"channel,omitempty"`
	End      FHIRDateTime                 `bson:"end,omitempty", json:"end,omitempty"`
	Tag      []SubscriptionTagComponent   `bson:"tag,omitempty", json:"tag,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec channel
type SubscriptionChannelComponent struct {
	Type    string `bson:"type,omitempty", json:"type,omitempty"`
	Url     string `bson:"url,omitempty", json:"url,omitempty"`
	Payload string `bson:"payload,omitempty", json:"payload,omitempty"`
	Header  string `bson:"header,omitempty", json:"header,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec tag
type SubscriptionTagComponent struct {
	Term        string `bson:"term,omitempty", json:"term,omitempty"`
	Scheme      string `bson:"scheme,omitempty", json:"scheme,omitempty"`
	Description string `bson:"description,omitempty", json:"description,omitempty"`
}
type SubscriptionBundle struct {
	Type         string
	Title        string
	Id           string
	Updated      time.Time
	TotalResults int
	Entries      []Subscription
	Category     SubscriptionCategory
}

type SubscriptionCategory struct {
	Term   string
	Label  string
	Scheme string
}
