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
	Identifier      Identifier                     `bson:"identifier"`
	Date            time.Time                      `bson:"date"`
	FhirType        CodeableConcept                `bson:"fhirType"`
	Class           CodeableConcept                `bson:"class"`
	Title           string                         `bson:"title"`
	Status          string                         `bson:"status"`
	Confidentiality Coding                         `bson:"confidentiality"`
	Subject         Reference                      `bson:"subject"`
	Author          []Reference                    `bson:"author"`
	Attester        []CompositionAttesterComponent `bson:"attester"`
	Custodian       Reference                      `bson:"custodian"`
	Event           []CompositionEventComponent    `bson:"event"`
	Encounter       Reference                      `bson:"encounter"`
	Section         []SectionComponent             `bson:"section"`
}

// This is an ugly hack to deal with embedded structures in the spec attester
type CompositionAttesterComponent struct {
	Mode  string    `bson:"mode"`
	Time  time.Time `bson:"time"`
	Party Reference `bson:"party"`
}

// This is an ugly hack to deal with embedded structures in the spec event
type CompositionEventComponent struct {
	Code   []CodeableConcept `bson:"code"`
	Period Period            `bson:"period"`
	Detail []Reference       `bson:"detail"`
}

// This is an ugly hack to deal with embedded structures in the spec section
type SectionComponent struct {
	Title       string             `bson:"title"`
	Identifier  []Identifier       `bson:"identifier"`
	Code        CodeableConcept    `bson:"code"`
	Subject     Reference          `bson:"subject"`
	Text        Narrative          `bson:"text"`
	EmptyReason CodeableConcept    `bson:"emptyReason"`
	Order       CodeableConcept    `bson:"order"`
	Section     []SectionComponent `bson:"section"`
	Entry       []Reference        `bson:"entry"`
}
