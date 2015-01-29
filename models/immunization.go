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

type Immunization struct {
	Id                  string                                     `json:"-" bson:"_id"`
	Identifier          []Identifier                               `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Date                FHIRDateTime                               `bson:"date,omitempty", json:"date,omitempty"`
	VaccineType         CodeableConcept                            `bson:"vaccineType,omitempty", json:"vaccineType,omitempty"`
	Subject             Reference                                  `bson:"subject,omitempty", json:"subject,omitempty"`
	RefusedIndicator    *bool                                      `bson:"refusedIndicator,omitempty", json:"refusedIndicator,omitempty"`
	Reported            *bool                                      `bson:"reported,omitempty", json:"reported,omitempty"`
	Performer           Reference                                  `bson:"performer,omitempty", json:"performer,omitempty"`
	Requester           Reference                                  `bson:"requester,omitempty", json:"requester,omitempty"`
	Manufacturer        Reference                                  `bson:"manufacturer,omitempty", json:"manufacturer,omitempty"`
	Location            Reference                                  `bson:"location,omitempty", json:"location,omitempty"`
	LotNumber           string                                     `bson:"lotNumber,omitempty", json:"lotNumber,omitempty"`
	ExpirationDate      FHIRDateTime                               `bson:"expirationDate,omitempty", json:"expirationDate,omitempty"`
	Site                CodeableConcept                            `bson:"site,omitempty", json:"site,omitempty"`
	Route               CodeableConcept                            `bson:"route,omitempty", json:"route,omitempty"`
	DoseQuantity        Quantity                                   `bson:"doseQuantity,omitempty", json:"doseQuantity,omitempty"`
	Explanation         ImmunizationExplanationComponent           `bson:"explanation,omitempty", json:"explanation,omitempty"`
	Reaction            []ImmunizationReactionComponent            `bson:"reaction,omitempty", json:"reaction,omitempty"`
	VaccinationProtocol []ImmunizationVaccinationProtocolComponent `bson:"vaccinationProtocol,omitempty", json:"vaccinationProtocol,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec explanation
type ImmunizationExplanationComponent struct {
	Reason        []CodeableConcept `bson:"reason,omitempty", json:"reason,omitempty"`
	RefusalReason []CodeableConcept `bson:"refusalReason,omitempty", json:"refusalReason,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec reaction
type ImmunizationReactionComponent struct {
	Date     FHIRDateTime `bson:"date,omitempty", json:"date,omitempty"`
	Detail   Reference    `bson:"detail,omitempty", json:"detail,omitempty"`
	Reported *bool        `bson:"reported,omitempty", json:"reported,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec vaccinationProtocol
type ImmunizationVaccinationProtocolComponent struct {
	DoseSequence     float64         `bson:"doseSequence,omitempty", json:"doseSequence,omitempty"`
	Description      string          `bson:"description,omitempty", json:"description,omitempty"`
	Authority        Reference       `bson:"authority,omitempty", json:"authority,omitempty"`
	Series           string          `bson:"series,omitempty", json:"series,omitempty"`
	SeriesDoses      float64         `bson:"seriesDoses,omitempty", json:"seriesDoses,omitempty"`
	DoseTarget       CodeableConcept `bson:"doseTarget,omitempty", json:"doseTarget,omitempty"`
	DoseStatus       CodeableConcept `bson:"doseStatus,omitempty", json:"doseStatus,omitempty"`
	DoseStatusReason CodeableConcept `bson:"doseStatusReason,omitempty", json:"doseStatusReason,omitempty"`
}

type ImmunizationBundle struct {
	Type         string                    `json:"resourceType,omitempty"`
	Title        string                    `json:"title,omitempty"`
	Id           string                    `json:"id,omitempty"`
	Updated      time.Time                 `json:"updated,omitempty"`
	TotalResults int                       `json:"totalResults,omitempty"`
	Entry        []ImmunizationBundleEntry `json:"entry,omitempty"`
	Category     ImmunizationCategory      `json:"category,omitempty"`
}

type ImmunizationBundleEntry struct {
	Title    string               `json:"title,omitempty"`
	Id       string               `json:"id,omitempty"`
	Content  Immunization         `json:"content,omitempty"`
	Category ImmunizationCategory `json:"category,omitempty"`
}

type ImmunizationCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
