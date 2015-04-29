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

type MedicationStatement struct {
	Id             string                               `json:"-" bson:"_id"`
	Identifier     []Identifier                         `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Patient        *Reference                           `bson:"patient,omitempty", json:"patient,omitempty"`
	WasNotGiven    *bool                                `bson:"wasNotGiven,omitempty", json:"wasNotGiven,omitempty"`
	ReasonNotGiven []CodeableConcept                    `bson:"reasonNotGiven,omitempty", json:"reasonNotGiven,omitempty"`
	WhenGiven      *Period                              `bson:"whenGiven,omitempty", json:"whenGiven,omitempty"`
	Medication     *Reference                           `bson:"medication,omitempty", json:"medication,omitempty"`
	Device         []Reference                          `bson:"device,omitempty", json:"device,omitempty"`
	Dosage         []MedicationStatementDosageComponent `bson:"dosage,omitempty", json:"dosage,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec dosage
type MedicationStatementDosageComponent struct {
	Schedule                *Timing          `bson:"schedule,omitempty", json:"schedule,omitempty"`
	AsNeededBoolean         *bool            `bson:"asNeededBoolean,omitempty", json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept `bson:"asNeededCodeableConcept,omitempty", json:"asNeededCodeableConcept,omitempty"`
	Site                    *CodeableConcept `bson:"site,omitempty", json:"site,omitempty"`
	Route                   *CodeableConcept `bson:"route,omitempty", json:"route,omitempty"`
	Method                  *CodeableConcept `bson:"method,omitempty", json:"method,omitempty"`
	Quantity                *Quantity        `bson:"quantity,omitempty", json:"quantity,omitempty"`
	Rate                    *Ratio           `bson:"rate,omitempty", json:"rate,omitempty"`
	MaxDosePerPeriod        *Ratio           `bson:"maxDosePerPeriod,omitempty", json:"maxDosePerPeriod,omitempty"`
}

type MedicationStatementBundle struct {
	Type         string                           `json:"resourceType,omitempty"`
	Title        string                           `json:"title,omitempty"`
	Id           string                           `json:"id,omitempty"`
	Updated      time.Time                        `json:"updated,omitempty"`
	TotalResults int                              `json:"totalResults,omitempty"`
	Entry        []MedicationStatementBundleEntry `json:"entry,omitempty"`
	Category     MedicationStatementCategory      `json:"category,omitempty"`
}

type MedicationStatementBundleEntry struct {
	Title    string                      `json:"title,omitempty"`
	Id       string                      `json:"id,omitempty"`
	Content  MedicationStatement         `json:"content,omitempty"`
	Category MedicationStatementCategory `json:"category,omitempty"`
}

type MedicationStatementCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
