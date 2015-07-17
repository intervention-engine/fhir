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

type Observation struct {
	Id                      string                               `json:"id" bson:"_id"`
	Code                    *CodeableConcept                     `bson:"code,omitempty" json:"code,omitempty"`
	Category                *CodeableConcept                     `bson:"category,omitempty" json:"category,omitempty"`
	ValueQuantity           *Quantity                            `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept    *CodeableConcept                     `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString             string                               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueRange              *Range                               `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio              *Ratio                               `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData        *SampledData                         `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueAttachment         *Attachment                          `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueTime               *FHIRDateTime                        `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime           *FHIRDateTime                        `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod             *Period                              `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason        *CodeableConcept                     `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation          *CodeableConcept                     `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	Comments                string                               `bson:"comments,omitempty" json:"comments,omitempty"`
	EffectiveDateTime       *FHIRDateTime                        `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod         *Period                              `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued                  *FHIRDateTime                        `bson:"issued,omitempty" json:"issued,omitempty"`
	Status                  string                               `bson:"status,omitempty" json:"status,omitempty"`
	Reliability             string                               `bson:"reliability,omitempty" json:"reliability,omitempty"`
	BodySiteCodeableConcept *CodeableConcept                     `bson:"bodySiteCodeableConcept,omitempty" json:"bodySiteCodeableConcept,omitempty"`
	BodySiteReference       *Reference                           `bson:"bodySiteReference,omitempty" json:"bodySiteReference,omitempty"`
	Method                  *CodeableConcept                     `bson:"method,omitempty" json:"method,omitempty"`
	Identifier              []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject                 *Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Specimen                *Reference                           `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Performer               []Reference                          `bson:"performer,omitempty" json:"performer,omitempty"`
	Device                  *Reference                           `bson:"device,omitempty" json:"device,omitempty"`
	Encounter               *Reference                           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	ReferenceRange          []ObservationReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
	DerivedFrom             []Reference                          `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Related                 []ObservationRelatedComponent        `bson:"related,omitempty" json:"related,omitempty"`
	Component               []ObservationComponentComponent      `bson:"component,omitempty" json:"component,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Observation) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Observation
	}{
		ResourceType: "Observation",
		Observation:  *resource,
	}
	return json.Marshal(x)
}

type ObservationReferenceRangeComponent struct {
	Low     *Quantity        `bson:"low,omitempty" json:"low,omitempty"`
	High    *Quantity        `bson:"high,omitempty" json:"high,omitempty"`
	Meaning *CodeableConcept `bson:"meaning,omitempty" json:"meaning,omitempty"`
	Age     *Range           `bson:"age,omitempty" json:"age,omitempty"`
	Text    string           `bson:"text,omitempty" json:"text,omitempty"`
}

type ObservationRelatedComponent struct {
	Type   string     `bson:"type,omitempty" json:"type,omitempty"`
	Target *Reference `bson:"target,omitempty" json:"target,omitempty"`
}

type ObservationComponentComponent struct {
	Code                 *CodeableConcept                     `bson:"code,omitempty" json:"code,omitempty"`
	ValueQuantity        *Quantity                            `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept                     `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          string                               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueRange           *Range                               `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                               `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                         `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueAttachment      *Attachment                          `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueTime            *FHIRDateTime                        `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime        *FHIRDateTime                        `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                              `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept                     `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	ReferenceRange       []ObservationReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}
