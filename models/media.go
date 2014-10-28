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

type Media struct {
	Id         string          `json:"-" bson:"_id"`
	Type       string          `bson:"type,omitempty", json:"type,omitempty"`
	Subtype    CodeableConcept `bson:"subtype,omitempty", json:"subtype,omitempty"`
	Identifier []Identifier    `bson:"identifier,omitempty", json:"identifier,omitempty"`
	DateTime   FHIRDateTime    `bson:"dateTime,omitempty", json:"dateTime,omitempty"`
	Subject    Reference       `bson:"subject,omitempty", json:"subject,omitempty"`
	Operator   Reference       `bson:"operator,omitempty", json:"operator,omitempty"`
	View       CodeableConcept `bson:"view,omitempty", json:"view,omitempty"`
	DeviceName string          `bson:"deviceName,omitempty", json:"deviceName,omitempty"`
	Height     float64         `bson:"height,omitempty", json:"height,omitempty"`
	Width      float64         `bson:"width,omitempty", json:"width,omitempty"`
	Frames     float64         `bson:"frames,omitempty", json:"frames,omitempty"`
	Length     float64         `bson:"length,omitempty", json:"length,omitempty"`
	Content    Attachment      `bson:"content,omitempty", json:"content,omitempty"`
}
type MediaBundle struct {
	Type         string
	Title        string
	Id           string
	Updated      time.Time
	TotalResults int
	Entries      []Media
	Category     MediaCategory
}

type MediaCategory struct {
	Term   string
	Label  string
	Scheme string
}
