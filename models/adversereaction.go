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

type AdverseReaction struct {
	Id              string                             `json:"-" bson:"_id"`
	Identifier      []Identifier                       `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Date            FHIRDateTime                       `bson:"date,omitempty", json:"date,omitempty"`
	Subject         Reference                          `bson:"subject,omitempty", json:"subject,omitempty"`
	DidNotOccurFlag *bool                              `bson:"didNotOccurFlag,omitempty", json:"didNotOccurFlag,omitempty"`
	Recorder        Reference                          `bson:"recorder,omitempty", json:"recorder,omitempty"`
	Symptom         []AdverseReactionSymptomComponent  `bson:"symptom,omitempty", json:"symptom,omitempty"`
	Exposure        []AdverseReactionExposureComponent `bson:"exposure,omitempty", json:"exposure,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec symptom
type AdverseReactionSymptomComponent struct {
	Code     CodeableConcept `bson:"code,omitempty", json:"code,omitempty"`
	Severity string          `bson:"severity,omitempty", json:"severity,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec exposure
type AdverseReactionExposureComponent struct {
	Date                 FHIRDateTime `bson:"date,omitempty", json:"date,omitempty"`
	Type                 string       `bson:"type,omitempty", json:"type,omitempty"`
	CausalityExpectation string       `bson:"causalityExpectation,omitempty", json:"causalityExpectation,omitempty"`
	Substance            Reference    `bson:"substance,omitempty", json:"substance,omitempty"`
}

type AdverseReactionBundle struct {
	Type         string
	Title        string
	Id           string
	Updated      time.Time
	TotalResults int
	Entries      []AdverseReaction
	Category     AdverseReactionCategory
}

type AdverseReactionCategory struct {
	Term   string
	Label  string
	Scheme string
}
