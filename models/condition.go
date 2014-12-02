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

type Condition struct {
	Id               string                          `json:"-" bson:"_id"`
	Identifier       []Identifier                    `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Subject          Reference                       `bson:"subject,omitempty", json:"subject,omitempty"`
	Encounter        Reference                       `bson:"encounter,omitempty", json:"encounter,omitempty"`
	Asserter         Reference                       `bson:"asserter,omitempty", json:"asserter,omitempty"`
	DateAsserted     FHIRDateTime                    `bson:"dateAsserted,omitempty", json:"dateAsserted,omitempty"`
	Code             CodeableConcept                 `bson:"code,omitempty", json:"code,omitempty"`
	Category         CodeableConcept                 `bson:"category,omitempty", json:"category,omitempty"`
	Status           string                          `bson:"status,omitempty", json:"status,omitempty"`
	Certainty        CodeableConcept                 `bson:"certainty,omitempty", json:"certainty,omitempty"`
	Severity         CodeableConcept                 `bson:"severity,omitempty", json:"severity,omitempty"`
	OnsetDate        FHIRDateTime                    `bson:"onsetDate,omitempty", json:"onsetDate,omitempty"`
	OnsetAge         Quantity                        `bson:"onsetAge,omitempty", json:"onsetAge,omitempty"`
	AbatementDate    FHIRDateTime                    `bson:"abatementDate,omitempty", json:"abatementDate,omitempty"`
	AbatementAge     Quantity                        `bson:"abatementAge,omitempty", json:"abatementAge,omitempty"`
	AbatementBoolean *bool                           `bson:"abatementBoolean,omitempty", json:"abatementBoolean,omitempty"`
	Stage            ConditionStageComponent         `bson:"stage,omitempty", json:"stage,omitempty"`
	Evidence         []ConditionEvidenceComponent    `bson:"evidence,omitempty", json:"evidence,omitempty"`
	Location         []ConditionLocationComponent    `bson:"location,omitempty", json:"location,omitempty"`
	RelatedItem      []ConditionRelatedItemComponent `bson:"relatedItem,omitempty", json:"relatedItem,omitempty"`
	Notes            string                          `bson:"notes,omitempty", json:"notes,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec stage
type ConditionStageComponent struct {
	Summary    CodeableConcept `bson:"summary,omitempty", json:"summary,omitempty"`
	Assessment []Reference     `bson:"assessment,omitempty", json:"assessment,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec evidence
type ConditionEvidenceComponent struct {
	Code   CodeableConcept `bson:"code,omitempty", json:"code,omitempty"`
	Detail []Reference     `bson:"detail,omitempty", json:"detail,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec location
type ConditionLocationComponent struct {
	Code   CodeableConcept `bson:"code,omitempty", json:"code,omitempty"`
	Detail string          `bson:"detail,omitempty", json:"detail,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec relatedItem
type ConditionRelatedItemComponent struct {
	Type   string          `bson:"type,omitempty", json:"type,omitempty"`
	Code   CodeableConcept `bson:"code,omitempty", json:"code,omitempty"`
	Target Reference       `bson:"target,omitempty", json:"target,omitempty"`
}

type ConditionBundle struct {
	Type         string
	Title        string
	Id           string
	Updated      time.Time
	TotalResults int
	Entries      []Condition
	Category     ConditionCategory
}

type ConditionCategory struct {
	Term   string
	Label  string
	Scheme string
}
