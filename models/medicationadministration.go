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

type MedicationAdministration struct {
	Id                    string                                    `json:"-" bson:"_id"`
	Identifier            []Identifier                              `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Status                string                                    `bson:"status,omitempty", json:"status,omitempty"`
	Patient               Reference                                 `bson:"patient,omitempty", json:"patient,omitempty"`
	Practitioner          Reference                                 `bson:"practitioner,omitempty", json:"practitioner,omitempty"`
	Encounter             Reference                                 `bson:"encounter,omitempty", json:"encounter,omitempty"`
	Prescription          Reference                                 `bson:"prescription,omitempty", json:"prescription,omitempty"`
	WasNotGiven           *bool                                     `bson:"wasNotGiven,omitempty", json:"wasNotGiven,omitempty"`
	ReasonNotGiven        []CodeableConcept                         `bson:"reasonNotGiven,omitempty", json:"reasonNotGiven,omitempty"`
	EffectiveTimeDateTime FHIRDateTime                              `bson:"effectiveTimeDateTime,omitempty", json:"effectiveTimeDateTime,omitempty"`
	EffectiveTimePeriod   Period                                    `bson:"effectiveTimePeriod,omitempty", json:"effectiveTimePeriod,omitempty"`
	Medication            Reference                                 `bson:"medication,omitempty", json:"medication,omitempty"`
	Device                []Reference                               `bson:"device,omitempty", json:"device,omitempty"`
	Dosage                []MedicationAdministrationDosageComponent `bson:"dosage,omitempty", json:"dosage,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec dosage
type MedicationAdministrationDosageComponent struct {
	TimingDateTime          FHIRDateTime    `bson:"timingDateTime,omitempty", json:"timingDateTime,omitempty"`
	TimingPeriod            Period          `bson:"timingPeriod,omitempty", json:"timingPeriod,omitempty"`
	AsNeededBoolean         *bool           `bson:"asNeededBoolean,omitempty", json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept CodeableConcept `bson:"asNeededCodeableConcept,omitempty", json:"asNeededCodeableConcept,omitempty"`
	Site                    CodeableConcept `bson:"site,omitempty", json:"site,omitempty"`
	Route                   CodeableConcept `bson:"route,omitempty", json:"route,omitempty"`
	Method                  CodeableConcept `bson:"method,omitempty", json:"method,omitempty"`
	Quantity                Quantity        `bson:"quantity,omitempty", json:"quantity,omitempty"`
	Rate                    Ratio           `bson:"rate,omitempty", json:"rate,omitempty"`
	MaxDosePerPeriod        Ratio           `bson:"maxDosePerPeriod,omitempty", json:"maxDosePerPeriod,omitempty"`
}
type MedicationAdministrationBundle struct {
	Type         string
	Title        string
	Id           string
	Updated      time.Time
	TotalResults int
	Entries      []MedicationAdministration
	Category     MedicationAdministrationCategory
}

type MedicationAdministrationCategory struct {
	Term   string
	Label  string
	Scheme string
}
