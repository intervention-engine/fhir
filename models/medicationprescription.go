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

type MedicationPrescription struct {
	Id                    string                                             `json:"-" bson:"_id"`
	Identifier            []Identifier                                       `bson:"identifier"`
	DateWritten           time.Time                                          `bson:"dateWritten"`
	Status                string                                             `bson:"status"`
	Patient               Reference                                          `bson:"patient"`
	Prescriber            Reference                                          `bson:"prescriber"`
	Encounter             Reference                                          `bson:"encounter"`
	ReasonCodeableConcept CodeableConcept                                    `bson:"reasonCodeableConcept"`
	ReasonReference       Reference                                          `bson:"reasonReference"`
	Medication            Reference                                          `bson:"medication"`
	DosageInstruction     []MedicationPrescriptionDosageInstructionComponent `bson:"dosageInstruction"`
	Dispense              MedicationPrescriptionDispenseComponent            `bson:"dispense"`
	Substitution          MedicationPrescriptionSubstitutionComponent        `bson:"substitution"`
}

// This is an ugly hack to deal with embedded structures in the spec dosageInstruction
type MedicationPrescriptionDosageInstructionComponent struct {
	Text                    string          `bson:"text"`
	AdditionalInstructions  CodeableConcept `bson:"additionalInstructions"`
	ScheduledDateTime       time.Time       `bson:"scheduledDateTime"`
	ScheduledPeriod         Period          `bson:"scheduledPeriod"`
	ScheduledTiming         Timing          `bson:"scheduledTiming"`
	AsNeededBoolean         bool            `bson:"asNeededBoolean"`
	AsNeededCodeableConcept CodeableConcept `bson:"asNeededCodeableConcept"`
	Site                    CodeableConcept `bson:"site"`
	Route                   CodeableConcept `bson:"route"`
	Method                  CodeableConcept `bson:"method"`
	DoseQuantity            Quantity        `bson:"doseQuantity"`
	Rate                    Ratio           `bson:"rate"`
	MaxDosePerPeriod        Ratio           `bson:"maxDosePerPeriod"`
}

// This is an ugly hack to deal with embedded structures in the spec dispense
type MedicationPrescriptionDispenseComponent struct {
	Medication             Reference `bson:"medication"`
	ValidityPeriod         Period    `bson:"validityPeriod"`
	NumberOfRepeatsAllowed float64   `bson:"numberOfRepeatsAllowed"`
	Quantity               Quantity  `bson:"quantity"`
	ExpectedSupplyDuration Quantity  `bson:"expectedSupplyDuration"`
}

// This is an ugly hack to deal with embedded structures in the spec substitution
type MedicationPrescriptionSubstitutionComponent struct {
	Type   CodeableConcept `bson:"type"`
	Reason CodeableConcept `bson:"reason"`
}
