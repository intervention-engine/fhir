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

type DiagnosticReport struct {
	Id                string                           `json:"id" bson:"_id"`
	Code              *CodeableConcept                 `bson:"code,omitempty" json:"code,omitempty"`
	Status            string                           `bson:"status,omitempty" json:"status,omitempty"`
	Issued            *FHIRDateTime                    `bson:"issued,omitempty" json:"issued,omitempty"`
	Subject           *Reference                       `bson:"subject,omitempty" json:"subject,omitempty"`
	Performer         *Reference                       `bson:"performer,omitempty" json:"performer,omitempty"`
	Encounter         *Reference                       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Identifier        []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	RequestDetail     []Reference                      `bson:"requestDetail,omitempty" json:"requestDetail,omitempty"`
	ServiceCategory   *CodeableConcept                 `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	EffectiveDateTime *FHIRDateTime                    `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period                          `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Specimen          []Reference                      `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Result            []Reference                      `bson:"result,omitempty" json:"result,omitempty"`
	ImagingStudy      []Reference                      `bson:"imagingStudy,omitempty" json:"imagingStudy,omitempty"`
	Image             []DiagnosticReportImageComponent `bson:"image,omitempty" json:"image,omitempty"`
	Conclusion        string                           `bson:"conclusion,omitempty" json:"conclusion,omitempty"`
	CodedDiagnosis    []CodeableConcept                `bson:"codedDiagnosis,omitempty" json:"codedDiagnosis,omitempty"`
	PresentedForm     []Attachment                     `bson:"presentedForm,omitempty" json:"presentedForm,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DiagnosticReport) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		DiagnosticReport
	}{
		ResourceType:     "DiagnosticReport",
		DiagnosticReport: *resource,
	}
	return json.Marshal(x)
}

type DiagnosticReportImageComponent struct {
	Comment string     `bson:"comment,omitempty" json:"comment,omitempty"`
	Link    *Reference `bson:"link,omitempty" json:"link,omitempty"`
}
