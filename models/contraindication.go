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

type Contraindication struct {
	Id         string                                `json:"-" bson:"_id"`
	Patient    *Reference                            `bson:"patient,omitempty", json:"patient,omitempty"`
	Category   *CodeableConcept                      `bson:"category,omitempty", json:"category,omitempty"`
	Severity   string                                `bson:"severity,omitempty", json:"severity,omitempty"`
	Implicated []Reference                           `bson:"implicated,omitempty", json:"implicated,omitempty"`
	Detail     string                                `bson:"detail,omitempty", json:"detail,omitempty"`
	Date       *FHIRDateTime                         `bson:"date,omitempty", json:"date,omitempty"`
	Author     *Reference                            `bson:"author,omitempty", json:"author,omitempty"`
	Identifier *Identifier                           `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Reference  string                                `bson:"reference,omitempty", json:"reference,omitempty"`
	Mitigation []ContraindicationMitigationComponent `bson:"mitigation,omitempty", json:"mitigation,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec mitigation
type ContraindicationMitigationComponent struct {
	Action *CodeableConcept `bson:"action,omitempty", json:"action,omitempty"`
	Date   *FHIRDateTime    `bson:"date,omitempty", json:"date,omitempty"`
	Author *Reference       `bson:"author,omitempty", json:"author,omitempty"`
}

type ContraindicationBundle struct {
	Type         string                        `json:"resourceType,omitempty"`
	Title        string                        `json:"title,omitempty"`
	Id           string                        `json:"id,omitempty"`
	Updated      time.Time                     `json:"updated,omitempty"`
	TotalResults int                           `json:"totalResults,omitempty"`
	Entry        []ContraindicationBundleEntry `json:"entry,omitempty"`
	Category     ContraindicationCategory      `json:"category,omitempty"`
}

type ContraindicationBundleEntry struct {
	Title    string                   `json:"title,omitempty"`
	Id       string                   `json:"id,omitempty"`
	Content  Contraindication         `json:"content,omitempty"`
	Category ContraindicationCategory `json:"category,omitempty"`
}

type ContraindicationCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
