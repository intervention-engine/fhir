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

type Composition struct {
	Id              string                         `json:"-" bson:"_id"`
	Identifier      Identifier                     `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Date            FHIRDateTime                   `bson:"date,omitempty", json:"date,omitempty"`
	Type            CodeableConcept                `bson:"type,omitempty", json:"type,omitempty"`
	Class           CodeableConcept                `bson:"class,omitempty", json:"class,omitempty"`
	Title           string                         `bson:"title,omitempty", json:"title,omitempty"`
	Status          string                         `bson:"status,omitempty", json:"status,omitempty"`
	Confidentiality Coding                         `bson:"confidentiality,omitempty", json:"confidentiality,omitempty"`
	Subject         Reference                      `bson:"subject,omitempty", json:"subject,omitempty"`
	Author          []Reference                    `bson:"author,omitempty", json:"author,omitempty"`
	Attester        []CompositionAttesterComponent `bson:"attester,omitempty", json:"attester,omitempty"`
	Custodian       Reference                      `bson:"custodian,omitempty", json:"custodian,omitempty"`
	Event           []CompositionEventComponent    `bson:"event,omitempty", json:"event,omitempty"`
	Encounter       Reference                      `bson:"encounter,omitempty", json:"encounter,omitempty"`
	Section         []SectionComponent             `bson:"section,omitempty", json:"section,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec attester
type CompositionAttesterComponent struct {
	Mode  []string     `bson:"mode,omitempty", json:"mode,omitempty"`
	Time  FHIRDateTime `bson:"time,omitempty", json:"time,omitempty"`
	Party Reference    `bson:"party,omitempty", json:"party,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec event
type CompositionEventComponent struct {
	Code   []CodeableConcept `bson:"code,omitempty", json:"code,omitempty"`
	Period Period            `bson:"period,omitempty", json:"period,omitempty"`
	Detail []Reference       `bson:"detail,omitempty", json:"detail,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec section
type SectionComponent struct {
	Title       string             `bson:"title,omitempty", json:"title,omitempty"`
	Identifier  []Identifier       `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Code        CodeableConcept    `bson:"code,omitempty", json:"code,omitempty"`
	Subject     Reference          `bson:"subject,omitempty", json:"subject,omitempty"`
	Text        Narrative          `bson:"text,omitempty", json:"text,omitempty"`
	EmptyReason CodeableConcept    `bson:"emptyReason,omitempty", json:"emptyReason,omitempty"`
	Order       CodeableConcept    `bson:"order,omitempty", json:"order,omitempty"`
	Section     []SectionComponent `bson:"section,omitempty", json:"section,omitempty"`
	Entry       []Reference        `bson:"entry,omitempty", json:"entry,omitempty"`
}
type CompositionBundle struct {
	Type         string
	Title        string
	Id           string
	Updated      time.Time
	TotalResults int
	Entries      []Composition
	Category     CompositionCategory
}

type CompositionCategory struct {
	Term   string
	Label  string
	Scheme string
}
