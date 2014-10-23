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

type Immunization struct {
	Id                  string                                     `json:"-" bson:"_id"`
	Identifier          []Identifier                               `bson:"identifier"`
	Date                FHIRDateTime                               `bson:"date"`
	VaccineType         CodeableConcept                            `bson:"vaccineType"`
	Subject             Reference                                  `bson:"subject"`
	RefusedIndicator    bool                                       `bson:"refusedIndicator"`
	Reported            bool                                       `bson:"reported"`
	Performer           Reference                                  `bson:"performer"`
	Requester           Reference                                  `bson:"requester"`
	Manufacturer        Reference                                  `bson:"manufacturer"`
	Location            Reference                                  `bson:"location"`
	LotNumber           string                                     `bson:"lotNumber"`
	ExpirationDate      FHIRDateTime                               `bson:"expirationDate"`
	Site                CodeableConcept                            `bson:"site"`
	Route               CodeableConcept                            `bson:"route"`
	DoseQuantity        Quantity                                   `bson:"doseQuantity"`
	Explanation         ImmunizationExplanationComponent           `bson:"explanation"`
	Reaction            []ImmunizationReactionComponent            `bson:"reaction"`
	VaccinationProtocol []ImmunizationVaccinationProtocolComponent `bson:"vaccinationProtocol"`
}

// This is an ugly hack to deal with embedded structures in the spec explanation
type ImmunizationExplanationComponent struct {
	Reason        []CodeableConcept `bson:"reason"`
	RefusalReason []CodeableConcept `bson:"refusalReason"`
}

// This is an ugly hack to deal with embedded structures in the spec reaction
type ImmunizationReactionComponent struct {
	Date     FHIRDateTime `bson:"date"`
	Detail   Reference    `bson:"detail"`
	Reported bool         `bson:"reported"`
}

// This is an ugly hack to deal with embedded structures in the spec vaccinationProtocol
type ImmunizationVaccinationProtocolComponent struct {
	DoseSequence     float64         `bson:"doseSequence"`
	Description      string          `bson:"description"`
	Authority        Reference       `bson:"authority"`
	Series           string          `bson:"series"`
	SeriesDoses      float64         `bson:"seriesDoses"`
	DoseTarget       CodeableConcept `bson:"doseTarget"`
	DoseStatus       CodeableConcept `bson:"doseStatus"`
	DoseStatusReason CodeableConcept `bson:"doseStatusReason"`
}
