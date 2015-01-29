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

type ImmunizationRecommendation struct {
	Id             string                                              `json:"-" bson:"_id"`
	Identifier     []Identifier                                        `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Subject        Reference                                           `bson:"subject,omitempty", json:"subject,omitempty"`
	Recommendation []ImmunizationRecommendationRecommendationComponent `bson:"recommendation,omitempty", json:"recommendation,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec dateCriterion
type ImmunizationRecommendationRecommendationDateCriterionComponent struct {
	Code  CodeableConcept `bson:"code,omitempty", json:"code,omitempty"`
	Value FHIRDateTime    `bson:"value,omitempty", json:"value,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec protocol
type ImmunizationRecommendationRecommendationProtocolComponent struct {
	DoseSequence float64   `bson:"doseSequence,omitempty", json:"doseSequence,omitempty"`
	Description  string    `bson:"description,omitempty", json:"description,omitempty"`
	Authority    Reference `bson:"authority,omitempty", json:"authority,omitempty"`
	Series       string    `bson:"series,omitempty", json:"series,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec recommendation
type ImmunizationRecommendationRecommendationComponent struct {
	Date                         FHIRDateTime                                                     `bson:"date,omitempty", json:"date,omitempty"`
	VaccineType                  CodeableConcept                                                  `bson:"vaccineType,omitempty", json:"vaccineType,omitempty"`
	DoseNumber                   float64                                                          `bson:"doseNumber,omitempty", json:"doseNumber,omitempty"`
	ForecastStatus               CodeableConcept                                                  `bson:"forecastStatus,omitempty", json:"forecastStatus,omitempty"`
	DateCriterion                []ImmunizationRecommendationRecommendationDateCriterionComponent `bson:"dateCriterion,omitempty", json:"dateCriterion,omitempty"`
	Protocol                     ImmunizationRecommendationRecommendationProtocolComponent        `bson:"protocol,omitempty", json:"protocol,omitempty"`
	SupportingImmunization       []Reference                                                      `bson:"supportingImmunization,omitempty", json:"supportingImmunization,omitempty"`
	SupportingPatientInformation []Reference                                                      `bson:"supportingPatientInformation,omitempty", json:"supportingPatientInformation,omitempty"`
}

type ImmunizationRecommendationBundle struct {
	Type         string                                  `json:"resourceType,omitempty"`
	Title        string                                  `json:"title,omitempty"`
	Id           string                                  `json:"id,omitempty"`
	Updated      time.Time                               `json:"updated,omitempty"`
	TotalResults int                                     `json:"totalResults,omitempty"`
	Entry        []ImmunizationRecommendationBundleEntry `json:"entry,omitempty"`
	Category     ImmunizationRecommendationCategory      `json:"category,omitempty"`
}

type ImmunizationRecommendationBundleEntry struct {
	Title    string                             `json:"title,omitempty"`
	Id       string                             `json:"id,omitempty"`
	Content  ImmunizationRecommendation         `json:"content,omitempty"`
	Category ImmunizationRecommendationCategory `json:"category,omitempty"`
}

type ImmunizationRecommendationCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
