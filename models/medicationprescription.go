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

type MedicationPrescription struct {
	Id                    string                                             `json:"-" bson:"_id"`
	Identifier            []Identifier                                       `bson:"identifier,omitempty", json:"identifier,omitempty"`
	DateWritten           FHIRDateTime                                       `bson:"dateWritten,omitempty", json:"dateWritten,omitempty"`
	Status                string                                             `bson:"status,omitempty", json:"status,omitempty"`
	Patient               Reference                                          `bson:"patient,omitempty", json:"patient,omitempty"`
	Prescriber            Reference                                          `bson:"prescriber,omitempty", json:"prescriber,omitempty"`
	Encounter             Reference                                          `bson:"encounter,omitempty", json:"encounter,omitempty"`
	ReasonCodeableConcept CodeableConcept                                    `bson:"reasonCodeableConcept,omitempty", json:"reasonCodeableConcept,omitempty"`
	ReasonReference       Reference                                          `bson:"reasonReference,omitempty", json:"reasonReference,omitempty"`
	Medication            Reference                                          `bson:"medication,omitempty", json:"medication,omitempty"`
	DosageInstruction     []MedicationPrescriptionDosageInstructionComponent `bson:"dosageInstruction,omitempty", json:"dosageInstruction,omitempty"`
	Dispense              MedicationPrescriptionDispenseComponent            `bson:"dispense,omitempty", json:"dispense,omitempty"`
	Substitution          MedicationPrescriptionSubstitutionComponent        `bson:"substitution,omitempty", json:"substitution,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec dosageInstruction
type MedicationPrescriptionDosageInstructionComponent struct {
	Text                    string          `bson:"text,omitempty", json:"text,omitempty"`
	AdditionalInstructions  CodeableConcept `bson:"additionalInstructions,omitempty", json:"additionalInstructions,omitempty"`
	ScheduledDateTime       FHIRDateTime    `bson:"scheduledDateTime,omitempty", json:"scheduledDateTime,omitempty"`
	ScheduledPeriod         Period          `bson:"scheduledPeriod,omitempty", json:"scheduledPeriod,omitempty"`
	ScheduledTiming         Timing          `bson:"scheduledTiming,omitempty", json:"scheduledTiming,omitempty"`
	AsNeededBoolean         *bool           `bson:"asNeededBoolean,omitempty", json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept CodeableConcept `bson:"asNeededCodeableConcept,omitempty", json:"asNeededCodeableConcept,omitempty"`
	Site                    CodeableConcept `bson:"site,omitempty", json:"site,omitempty"`
	Route                   CodeableConcept `bson:"route,omitempty", json:"route,omitempty"`
	Method                  CodeableConcept `bson:"method,omitempty", json:"method,omitempty"`
	DoseQuantity            Quantity        `bson:"doseQuantity,omitempty", json:"doseQuantity,omitempty"`
	Rate                    Ratio           `bson:"rate,omitempty", json:"rate,omitempty"`
	MaxDosePerPeriod        Ratio           `bson:"maxDosePerPeriod,omitempty", json:"maxDosePerPeriod,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec dispense
type MedicationPrescriptionDispenseComponent struct {
	Medication             Reference `bson:"medication,omitempty", json:"medication,omitempty"`
	ValidityPeriod         Period    `bson:"validityPeriod,omitempty", json:"validityPeriod,omitempty"`
	NumberOfRepeatsAllowed float64   `bson:"numberOfRepeatsAllowed,omitempty", json:"numberOfRepeatsAllowed,omitempty"`
	Quantity               Quantity  `bson:"quantity,omitempty", json:"quantity,omitempty"`
	ExpectedSupplyDuration Quantity  `bson:"expectedSupplyDuration,omitempty", json:"expectedSupplyDuration,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec substitution
type MedicationPrescriptionSubstitutionComponent struct {
	Type   CodeableConcept `bson:"type,omitempty", json:"type,omitempty"`
	Reason CodeableConcept `bson:"reason,omitempty", json:"reason,omitempty"`
}
