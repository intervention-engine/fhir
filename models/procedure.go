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

type Procedure struct {
	Id           string                          `json:"-" bson:"_id"`
	Identifier   []Identifier                    `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Subject      *Reference                      `bson:"subject,omitempty", json:"subject,omitempty"`
	Type         *CodeableConcept                `bson:"type,omitempty", json:"type,omitempty"`
	BodySite     []CodeableConcept               `bson:"bodySite,omitempty", json:"bodySite,omitempty"`
	Indication   []CodeableConcept               `bson:"indication,omitempty", json:"indication,omitempty"`
	Performer    []ProcedurePerformerComponent   `bson:"performer,omitempty", json:"performer,omitempty"`
	Date         *Period                         `bson:"date,omitempty", json:"date,omitempty"`
	Encounter    *Reference                      `bson:"encounter,omitempty", json:"encounter,omitempty"`
	Outcome      string                          `bson:"outcome,omitempty", json:"outcome,omitempty"`
	Report       []Reference                     `bson:"report,omitempty", json:"report,omitempty"`
	Complication []CodeableConcept               `bson:"complication,omitempty", json:"complication,omitempty"`
	FollowUp     string                          `bson:"followUp,omitempty", json:"followUp,omitempty"`
	RelatedItem  []ProcedureRelatedItemComponent `bson:"relatedItem,omitempty", json:"relatedItem,omitempty"`
	Notes        string                          `bson:"notes,omitempty", json:"notes,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec performer
type ProcedurePerformerComponent struct {
	Person *Reference       `bson:"person,omitempty", json:"person,omitempty"`
	Role   *CodeableConcept `bson:"role,omitempty", json:"role,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec relatedItem
type ProcedureRelatedItemComponent struct {
	Type   string     `bson:"type,omitempty", json:"type,omitempty"`
	Target *Reference `bson:"target,omitempty", json:"target,omitempty"`
}

type ProcedureBundle struct {
	Type         string                 `json:"resourceType,omitempty"`
	Title        string                 `json:"title,omitempty"`
	Id           string                 `json:"id,omitempty"`
	Updated      time.Time              `json:"updated,omitempty"`
	TotalResults int                    `json:"totalResults,omitempty"`
	Entry        []ProcedureBundleEntry `json:"entry,omitempty"`
	Category     ProcedureCategory      `json:"category,omitempty"`
}

type ProcedureBundleEntry struct {
	Title    string            `json:"title,omitempty"`
	Id       string            `json:"id,omitempty"`
	Content  Procedure         `json:"content,omitempty"`
	Category ProcedureCategory `json:"category,omitempty"`
}

type ProcedureCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
