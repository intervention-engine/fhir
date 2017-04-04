// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
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

type Dosage struct {
	Sequence                 *int32            `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Text                     string            `bson:"text,omitempty" json:"text,omitempty"`
	AdditionalInstruction    []CodeableConcept `bson:"additionalInstruction,omitempty" json:"additionalInstruction,omitempty"`
	PatientInstruction       string            `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	Timing                   *Timing           `bson:"timing,omitempty" json:"timing,omitempty"`
	AsNeededBoolean          *bool             `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept  *CodeableConcept  `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	Site                     *CodeableConcept  `bson:"site,omitempty" json:"site,omitempty"`
	Route                    *CodeableConcept  `bson:"route,omitempty" json:"route,omitempty"`
	Method                   *CodeableConcept  `bson:"method,omitempty" json:"method,omitempty"`
	DoseRange                *Range            `bson:"doseRange,omitempty" json:"doseRange,omitempty"`
	DoseSimpleQuantity       *Quantity         `bson:"doseSimpleQuantity,omitempty" json:"doseSimpleQuantity,omitempty"`
	MaxDosePerPeriod         *Ratio            `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty"`
	MaxDosePerAdministration *Quantity         `bson:"maxDosePerAdministration,omitempty" json:"maxDosePerAdministration,omitempty"`
	MaxDosePerLifetime       *Quantity         `bson:"maxDosePerLifetime,omitempty" json:"maxDosePerLifetime,omitempty"`
	RateRatio                *Ratio            `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateRange                *Range            `bson:"rateRange,omitempty" json:"rateRange,omitempty"`
	RateSimpleQuantity       *Quantity         `bson:"rateSimpleQuantity,omitempty" json:"rateSimpleQuantity,omitempty"`
}
