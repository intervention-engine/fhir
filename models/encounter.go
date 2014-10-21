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

type Encounter struct {
	Id              string                            `json:"-" bson:"_id"`
	Identifier      []Identifier                      `bson:"identifier"`
	Status          string                            `bson:"status"`
	Class           string                            `bson:"class"`
	Type            []CodeableConcept                 `bson:"type"`
	Subject         Reference                         `bson:"subject"`
	Participant     []EncounterParticipantComponent   `bson:"participant"`
	Fulfills        Reference                         `bson:"fulfills"`
	Period          Period                            `bson:"period"`
	Length          Quantity                          `bson:"length"`
	Reason          CodeableConcept                   `bson:"reason"`
	Indication      Reference                         `bson:"indication"`
	Priority        CodeableConcept                   `bson:"priority"`
	Hospitalization EncounterHospitalizationComponent `bson:"hospitalization"`
	Location        []EncounterLocationComponent      `bson:"location"`
	ServiceProvider Reference                         `bson:"serviceProvider"`
	PartOf          Reference                         `bson:"partOf"`
}

// This is an ugly hack to deal with embedded structures in the spec participant
type EncounterParticipantComponent struct {
	Type       []CodeableConcept `bson:"type"`
	Individual Reference         `bson:"individual"`
}

// This is an ugly hack to deal with embedded structures in the spec accomodation
type EncounterHospitalizationAccomodationComponent struct {
	Bed    Reference `bson:"bed"`
	Period Period    `bson:"period"`
}

// This is an ugly hack to deal with embedded structures in the spec hospitalization
type EncounterHospitalizationComponent struct {
	PreAdmissionIdentifier Identifier                                      `bson:"preAdmissionIdentifier"`
	Origin                 Reference                                       `bson:"origin"`
	AdmitSource            CodeableConcept                                 `bson:"admitSource"`
	Period                 Period                                          `bson:"period"`
	Accomodation           []EncounterHospitalizationAccomodationComponent `bson:"accomodation"`
	Diet                   CodeableConcept                                 `bson:"diet"`
	SpecialCourtesy        []CodeableConcept                               `bson:"specialCourtesy"`
	SpecialArrangement     []CodeableConcept                               `bson:"specialArrangement"`
	Destination            Reference                                       `bson:"destination"`
	DischargeDisposition   CodeableConcept                                 `bson:"dischargeDisposition"`
	DischargeDiagnosis     Reference                                       `bson:"dischargeDiagnosis"`
	ReAdmission            bool                                            `bson:"reAdmission"`
}

// This is an ugly hack to deal with embedded structures in the spec location
type EncounterLocationComponent struct {
	Location Reference `bson:"location"`
	Period   Period    `bson:"period"`
}
