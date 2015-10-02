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

type MedicationOrder struct {
	Id                        string                                      `json:"id" bson:"_id"`
	Identifier                []Identifier                                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	DateWritten               *FHIRDateTime                               `bson:"dateWritten,omitempty" json:"dateWritten,omitempty"`
	Status                    string                                      `bson:"status,omitempty" json:"status,omitempty"`
	DateEnded                 *FHIRDateTime                               `bson:"dateEnded,omitempty" json:"dateEnded,omitempty"`
	ReasonEnded               *CodeableConcept                            `bson:"reasonEnded,omitempty" json:"reasonEnded,omitempty"`
	Patient                   *Reference                                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Prescriber                *Reference                                  `bson:"prescriber,omitempty" json:"prescriber,omitempty"`
	Encounter                 *Reference                                  `bson:"encounter,omitempty" json:"encounter,omitempty"`
	ReasonCodeableConcept     *CodeableConcept                            `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference           *Reference                                  `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                      string                                      `bson:"note,omitempty" json:"note,omitempty"`
	MedicationCodeableConcept *CodeableConcept                            `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                                  `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	DosageInstruction         []MedicationOrderDosageInstructionComponent `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	DispenseRequest           *MedicationOrderDispenseRequestComponent    `bson:"dispenseRequest,omitempty" json:"dispenseRequest,omitempty"`
	Substitution              *MedicationOrderSubstitutionComponent       `bson:"substitution,omitempty" json:"substitution,omitempty"`
	PriorPrescription         *Reference                                  `bson:"priorPrescription,omitempty" json:"priorPrescription,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *MedicationOrder) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		MedicationOrder
	}{
		ResourceType:    "MedicationOrder",
		MedicationOrder: *resource,
	}
	return json.Marshal(x)
}

type MedicationOrderDosageInstructionComponent struct {
	Text                    string           `bson:"text,omitempty" json:"text,omitempty"`
	AdditionalInstructions  *CodeableConcept `bson:"additionalInstructions,omitempty" json:"additionalInstructions,omitempty"`
	Timing                  *Timing          `bson:"timing,omitempty" json:"timing,omitempty"`
	AsNeededBoolean         *bool            `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	SiteCodeableConcept     *CodeableConcept `bson:"siteCodeableConcept,omitempty" json:"siteCodeableConcept,omitempty"`
	SiteReference           *Reference       `bson:"siteReference,omitempty" json:"siteReference,omitempty"`
	Route                   *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	Method                  *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	DoseRange               *Range           `bson:"doseRange,omitempty" json:"doseRange,omitempty"`
	DoseSimpleQuantity      *Quantity        `bson:"doseSimpleQuantity,omitempty" json:"doseSimpleQuantity,omitempty"`
	RateRatio               *Ratio           `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateRange               *Range           `bson:"rateRange,omitempty" json:"rateRange,omitempty"`
	MaxDosePerPeriod        *Ratio           `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty"`
}

type MedicationOrderDispenseRequestComponent struct {
	MedicationCodeableConcept *CodeableConcept `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference       `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	ValidityPeriod            *Period          `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	NumberOfRepeatsAllowed    *uint32          `bson:"numberOfRepeatsAllowed,omitempty" json:"numberOfRepeatsAllowed,omitempty"`
	Quantity                  *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ExpectedSupplyDuration    *Quantity        `bson:"expectedSupplyDuration,omitempty" json:"expectedSupplyDuration,omitempty"`
}

type MedicationOrderSubstitutionComponent struct {
	Type   *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}
