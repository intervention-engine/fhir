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

type Condition struct {
	Id                 string                       `json:"id" bson:"_id"`
	Identifier         []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient            *Reference                   `bson:"patient,omitempty" json:"patient,omitempty"`
	Encounter          *Reference                   `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Asserter           *Reference                   `bson:"asserter,omitempty" json:"asserter,omitempty"`
	DateRecorded       *FHIRDateTime                `bson:"dateRecorded,omitempty" json:"dateRecorded,omitempty"`
	Code               *CodeableConcept             `bson:"code,omitempty" json:"code,omitempty"`
	Category           *CodeableConcept             `bson:"category,omitempty" json:"category,omitempty"`
	ClinicalStatus     string                       `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus string                       `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Severity           *CodeableConcept             `bson:"severity,omitempty" json:"severity,omitempty"`
	OnsetDateTime      *FHIRDateTime                `bson:"onsetDateTime,omitempty" json:"onsetDateTime,omitempty"`
	OnsetAge           *Quantity                    `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                      `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                       `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetString        string                       `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	AbatementDateTime  *FHIRDateTime                `bson:"abatementDateTime,omitempty" json:"abatementDateTime,omitempty"`
	AbatementAge       *Quantity                    `bson:"abatementAge,omitempty" json:"abatementAge,omitempty"`
	AbatementBoolean   *bool                        `bson:"abatementBoolean,omitempty" json:"abatementBoolean,omitempty"`
	AbatementPeriod    *Period                      `bson:"abatementPeriod,omitempty" json:"abatementPeriod,omitempty"`
	AbatementRange     *Range                       `bson:"abatementRange,omitempty" json:"abatementRange,omitempty"`
	AbatementString    string                       `bson:"abatementString,omitempty" json:"abatementString,omitempty"`
	Stage              *ConditionStageComponent     `bson:"stage,omitempty" json:"stage,omitempty"`
	Evidence           []ConditionEvidenceComponent `bson:"evidence,omitempty" json:"evidence,omitempty"`
	BodySite           []CodeableConcept            `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Notes              string                       `bson:"notes,omitempty" json:"notes,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Condition) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Condition
	}{
		ResourceType: "Condition",
		Condition:    *resource,
	}
	return json.Marshal(x)
}

type ConditionStageComponent struct {
	Summary    *CodeableConcept `bson:"summary,omitempty" json:"summary,omitempty"`
	Assessment []Reference      `bson:"assessment,omitempty" json:"assessment,omitempty"`
}

type ConditionEvidenceComponent struct {
	Code   *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Detail []Reference      `bson:"detail,omitempty" json:"detail,omitempty"`
}
