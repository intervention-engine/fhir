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

type Query struct {
	Id         string                  `json:"-" bson:"_id"`
	Identifier string                  `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Parameter  []Extension             `bson:"parameter,omitempty", json:"parameter,omitempty"`
	Response   *QueryResponseComponent `bson:"response,omitempty", json:"response,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec response
type QueryResponseComponent struct {
	Identifier string      `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Outcome    string      `bson:"outcome,omitempty", json:"outcome,omitempty"`
	Total      float64     `bson:"total,omitempty", json:"total,omitempty"`
	Parameter  []Extension `bson:"parameter,omitempty", json:"parameter,omitempty"`
	First      []Extension `bson:"first,omitempty", json:"first,omitempty"`
	Previous   []Extension `bson:"previous,omitempty", json:"previous,omitempty"`
	Next       []Extension `bson:"next,omitempty", json:"next,omitempty"`
	Last       []Extension `bson:"last,omitempty", json:"last,omitempty"`
	Reference  []Reference `bson:"reference,omitempty", json:"reference,omitempty"`
}

type QueryBundle struct {
	Type         string             `json:"resourceType,omitempty"`
	Title        string             `json:"title,omitempty"`
	Id           string             `json:"id,omitempty"`
	Updated      time.Time          `json:"updated,omitempty"`
	TotalResults int                `json:"totalResults,omitempty"`
	Entry        []QueryBundleEntry `json:"entry,omitempty"`
	Category     QueryCategory      `json:"category,omitempty"`
}

type QueryBundleEntry struct {
	Title    string        `json:"title,omitempty"`
	Id       string        `json:"id,omitempty"`
	Content  Query         `json:"content,omitempty"`
	Category QueryCategory `json:"category,omitempty"`
}

type QueryCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
