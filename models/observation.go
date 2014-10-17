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

type Observation struct {
	Id                   string                               `json:"-" bson:"_id"`
	Name                 CodeableConcept                      `bson:"name"`
	ValueQuantity        Quantity                             `bson:"valueQuantity"`
	ValueCodeableConcept CodeableConcept                      `bson:"valueCodeableConcept"`
	ValueAttachment      Attachment                           `bson:"valueAttachment"`
	ValueRatio           Ratio                                `bson:"valueRatio"`
	ValueDateTime        time.Time                            `bson:"valueDateTime"`
	ValuePeriod          Period                               `bson:"valuePeriod"`
	ValueSampledData     SampledData                          `bson:"valueSampledData"`
	ValueString          string                               `bson:"valueString"`
	ValueTime            time.Time                            `bson:"valueTime"`
	Interpretation       CodeableConcept                      `bson:"interpretation"`
	Comments             string                               `bson:"comments"`
	AppliesDateTime      time.Time                            `bson:"appliesDateTime"`
	AppliesPeriod        Period                               `bson:"appliesPeriod"`
	Issued               time.Time                            `bson:"issued"`
	Status               string                               `bson:"status"`
	Reliability          string                               `bson:"reliability"`
	BodySite             CodeableConcept                      `bson:"bodySite"`
	Method               CodeableConcept                      `bson:"method"`
	Identifier           Identifier                           `bson:"identifier"`
	Subject              Reference                            `bson:"subject"`
	Specimen             Reference                            `bson:"specimen"`
	Performer            []Reference                          `bson:"performer"`
	Encounter            Reference                            `bson:"encounter"`
	ReferenceRange       []ObservationReferenceRangeComponent `bson:"referenceRange"`
	Related              []ObservationRelatedComponent        `bson:"related"`
}

// This is an ugly hack to deal with embedded structures in the spec referenceRange
type ObservationReferenceRangeComponent struct {
	Low     Quantity        `bson:"low"`
	High    Quantity        `bson:"high"`
	Meaning CodeableConcept `bson:"meaning"`
	Age     Range           `bson:"age"`
	Text    string          `bson:"text"`
}

// This is an ugly hack to deal with embedded structures in the spec related
type ObservationRelatedComponent struct {
	FhirType string    `bson:"fhirType"`
	Target   Reference `bson:"target"`
}
