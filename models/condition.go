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

type Condition struct {
	Id               string                          `json:"-" bson:"_id"`
	Identifier       []Identifier                    `bson:"identifier"`
	Subject          Reference                       `bson:"subject"`
	Encounter        Reference                       `bson:"encounter"`
	Asserter         Reference                       `bson:"asserter"`
	DateAsserted     FHIRDateTime                    `bson:"dateAsserted"`
	Code             CodeableConcept                 `bson:"code"`
	Category         CodeableConcept                 `bson:"category"`
	Status           string                          `bson:"status"`
	Certainty        CodeableConcept                 `bson:"certainty"`
	Severity         CodeableConcept                 `bson:"severity"`
	OnsetDate        FHIRDateTime                    `bson:"onsetDate"`
	OnsetAge         Quantity                        `bson:"onsetAge"`
	AbatementDate    FHIRDateTime                    `bson:"abatementDate"`
	AbatementAge     Quantity                        `bson:"abatementAge"`
	AbatementBoolean bool                            `bson:"abatementBoolean"`
	Stage            ConditionStageComponent         `bson:"stage"`
	Evidence         []ConditionEvidenceComponent    `bson:"evidence"`
	Location         []ConditionLocationComponent    `bson:"location"`
	RelatedItem      []ConditionRelatedItemComponent `bson:"relatedItem"`
	Notes            string                          `bson:"notes"`
}

// This is an ugly hack to deal with embedded structures in the spec stage
type ConditionStageComponent struct {
	Summary    CodeableConcept `bson:"summary"`
	Assessment []Reference     `bson:"assessment"`
}

// This is an ugly hack to deal with embedded structures in the spec evidence
type ConditionEvidenceComponent struct {
	Code   CodeableConcept `bson:"code"`
	Detail []Reference     `bson:"detail"`
}

// This is an ugly hack to deal with embedded structures in the spec location
type ConditionLocationComponent struct {
	Code   CodeableConcept `bson:"code"`
	Detail string          `bson:"detail"`
}

// This is an ugly hack to deal with embedded structures in the spec relatedItem
type ConditionRelatedItemComponent struct {
	Type   string          `bson:"type"`
	Code   CodeableConcept `bson:"code"`
	Target Reference       `bson:"target"`
}
