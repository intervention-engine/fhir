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

type DataRequirement struct {
	Type        string                               `bson:"type,omitempty" json:"type,omitempty"`
	Profile     *Reference                           `bson:"profile,omitempty" json:"profile,omitempty"`
	MustSupport []string                             `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	CodeFilter  []DataRequirementCodeFilterComponent `bson:"codeFilter,omitempty" json:"codeFilter,omitempty"`
	DateFilter  []DataRequirementDateFilterComponent `bson:"dateFilter,omitempty" json:"dateFilter,omitempty"`
}

type DataRequirementCodeFilterComponent struct {
	BackboneElement      `bson:",inline"`
	Path                 string            `bson:"path,omitempty" json:"path,omitempty"`
	ValueSetString       string            `bson:"valueSetString,omitempty" json:"valueSetString,omitempty"`
	ValueSetReference    *Reference        `bson:"valueSetReference,omitempty" json:"valueSetReference,omitempty"`
	ValueCode            []string          `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCoding          []Coding          `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueCodeableConcept []CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
}

type DataRequirementDateFilterComponent struct {
	BackboneElement `bson:",inline"`
	Path            string        `bson:"path,omitempty" json:"path,omitempty"`
	ValueDateTime   *FHIRDateTime `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod     *Period       `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
}
