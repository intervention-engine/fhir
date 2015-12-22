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

type ImmunizationRecommendation struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient        *Reference                                          `bson:"patient,omitempty" json:"patient,omitempty"`
	Recommendation []ImmunizationRecommendationRecommendationComponent `bson:"recommendation,omitempty" json:"recommendation,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ImmunizationRecommendation) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		ImmunizationRecommendation
	}{
		ResourceType:               "ImmunizationRecommendation",
		ImmunizationRecommendation: *resource,
	}
	return json.Marshal(x)
}

// The "immunizationRecommendation" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type immunizationRecommendation ImmunizationRecommendation

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ImmunizationRecommendation) UnmarshalJSON(data []byte) (err error) {
	x2 := immunizationRecommendation{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ImmunizationRecommendation(x2)
	}
	return
}

type ImmunizationRecommendationRecommendationComponent struct {
	Date                         *FHIRDateTime                                                    `bson:"date,omitempty" json:"date,omitempty"`
	VaccineCode                  *CodeableConcept                                                 `bson:"vaccineCode,omitempty" json:"vaccineCode,omitempty"`
	DoseNumber                   *uint32                                                          `bson:"doseNumber,omitempty" json:"doseNumber,omitempty"`
	ForecastStatus               *CodeableConcept                                                 `bson:"forecastStatus,omitempty" json:"forecastStatus,omitempty"`
	DateCriterion                []ImmunizationRecommendationRecommendationDateCriterionComponent `bson:"dateCriterion,omitempty" json:"dateCriterion,omitempty"`
	Protocol                     *ImmunizationRecommendationRecommendationProtocolComponent       `bson:"protocol,omitempty" json:"protocol,omitempty"`
	SupportingImmunization       []Reference                                                      `bson:"supportingImmunization,omitempty" json:"supportingImmunization,omitempty"`
	SupportingPatientInformation []Reference                                                      `bson:"supportingPatientInformation,omitempty" json:"supportingPatientInformation,omitempty"`
}

type ImmunizationRecommendationRecommendationDateCriterionComponent struct {
	Code  *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Value *FHIRDateTime    `bson:"value,omitempty" json:"value,omitempty"`
}

type ImmunizationRecommendationRecommendationProtocolComponent struct {
	DoseSequence *int32     `bson:"doseSequence,omitempty" json:"doseSequence,omitempty"`
	Description  string     `bson:"description,omitempty" json:"description,omitempty"`
	Authority    *Reference `bson:"authority,omitempty" json:"authority,omitempty"`
	Series       string     `bson:"series,omitempty" json:"series,omitempty"`
}
