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

import "time"

type Encounter struct {
	Id                      string                                      `json:"-" bson:"_id"`
	Identifier              []Identifier                                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                  string                                      `bson:"status,omitempty" json:"status,omitempty"`
	StatusHistory           []EncounterEncounterStatusHistoryComponent  `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Class                   string                                      `bson:"class,omitempty" json:"class,omitempty"`
	Type                    []CodeableConcept                           `bson:"type,omitempty" json:"type,omitempty"`
	Patient                 *Reference                                  `bson:"patient,omitempty" json:"patient,omitempty"`
	EpisodeOfCare           *Reference                                  `bson:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty"`
	IncomingReferralRequest []Reference                                 `bson:"incomingReferralRequest,omitempty" json:"incomingReferralRequest,omitempty"`
	Participant             []EncounterEncounterParticipantComponent    `bson:"participant,omitempty" json:"participant,omitempty"`
	Fulfills                *Reference                                  `bson:"fulfills,omitempty" json:"fulfills,omitempty"`
	Period                  *Period                                     `bson:"period,omitempty" json:"period,omitempty"`
	Length                  *Quantity                                   `bson:"length,omitempty" json:"length,omitempty"`
	Reason                  []CodeableConcept                           `bson:"reason,omitempty" json:"reason,omitempty"`
	Indication              []Reference                                 `bson:"indication,omitempty" json:"indication,omitempty"`
	Priority                *CodeableConcept                            `bson:"priority,omitempty" json:"priority,omitempty"`
	Hospitalization         *EncounterEncounterHospitalizationComponent `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Location                []EncounterEncounterLocationComponent       `bson:"location,omitempty" json:"location,omitempty"`
	ServiceProvider         *Reference                                  `bson:"serviceProvider,omitempty" json:"serviceProvider,omitempty"`
	PartOf                  *Reference                                  `bson:"partOf,omitempty" json:"partOf,omitempty"`
}
type EncounterEncounterStatusHistoryComponent struct {
	Status string  `bson:"status,omitempty" json:"status,omitempty"`
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}
type EncounterEncounterParticipantComponent struct {
	Type       []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Period     *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Individual *Reference        `bson:"individual,omitempty" json:"individual,omitempty"`
}
type EncounterEncounterHospitalizationComponent struct {
	PreAdmissionIdentifier *Identifier       `bson:"preAdmissionIdentifier,omitempty" json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference        `bson:"origin,omitempty" json:"origin,omitempty"`
	AdmitSource            *CodeableConcept  `bson:"admitSource,omitempty" json:"admitSource,omitempty"`
	DietPreference         *CodeableConcept  `bson:"dietPreference,omitempty" json:"dietPreference,omitempty"`
	SpecialCourtesy        []CodeableConcept `bson:"specialCourtesy,omitempty" json:"specialCourtesy,omitempty"`
	SpecialArrangement     []CodeableConcept `bson:"specialArrangement,omitempty" json:"specialArrangement,omitempty"`
	Destination            *Reference        `bson:"destination,omitempty" json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept  `bson:"dischargeDisposition,omitempty" json:"dischargeDisposition,omitempty"`
	DischargeDiagnosis     *Reference        `bson:"dischargeDiagnosis,omitempty" json:"dischargeDiagnosis,omitempty"`
	ReAdmission            *bool             `bson:"reAdmission,omitempty" json:"reAdmission,omitempty"`
}
type EncounterEncounterLocationComponent struct {
	Location *Reference `bson:"location,omitempty" json:"location,omitempty"`
	Status   string     `bson:"status,omitempty" json:"status,omitempty"`
	Period   *Period    `bson:"period,omitempty" json:"period,omitempty"`
}

type EncounterBundle struct {
	Type         string                 `json:"resourceType,omitempty"`
	Title        string                 `json:"title,omitempty"`
	Id           string                 `json:"id,omitempty"`
	Updated      time.Time              `json:"updated,omitempty"`
	TotalResults int                    `json:"totalResults,omitempty"`
	Entry        []EncounterBundleEntry `json:"entry,omitempty"`
	Category     EncounterCategory      `json:"category,omitempty"`
}

type EncounterBundleEntry struct {
	Title    string            `json:"title,omitempty"`
	Id       string            `json:"id,omitempty"`
	Content  Encounter         `json:"content,omitempty"`
	Category EncounterCategory `json:"category,omitempty"`
}

type EncounterCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
