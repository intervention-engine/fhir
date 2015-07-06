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

import (
	"encoding/json"
	"time"
)

type PaymentReconciliation struct {
	Id                  string                                  `json:"-" bson:"_id"`
	Identifier          []Identifier                            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Request             *Reference                              `bson:"request,omitempty" json:"request,omitempty"`
	Outcome             string                                  `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition         string                                  `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Ruleset             *Coding                                 `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset     *Coding                                 `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created             *FHIRDateTime                           `bson:"created,omitempty" json:"created,omitempty"`
	Period              *Period                                 `bson:"period,omitempty" json:"period,omitempty"`
	Organization        *Reference                              `bson:"organization,omitempty" json:"organization,omitempty"`
	RequestProvider     *Reference                              `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization *Reference                              `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Detail              []PaymentReconciliationDetailsComponent `bson:"detail,omitempty" json:"detail,omitempty"`
	Form                *Coding                                 `bson:"form,omitempty" json:"form,omitempty"`
	Total               *Quantity                               `bson:"total,omitempty" json:"total,omitempty"`
	Note                []PaymentReconciliationNotesComponent   `bson:"note,omitempty" json:"note,omitempty"`
}

type PaymentReconciliationDetailsComponent struct {
	Type      *Coding       `bson:"type,omitempty" json:"type,omitempty"`
	Request   *Reference    `bson:"request,omitempty" json:"request,omitempty"`
	Responce  *Reference    `bson:"responce,omitempty" json:"responce,omitempty"`
	Submitter *Reference    `bson:"submitter,omitempty" json:"submitter,omitempty"`
	Payee     *Reference    `bson:"payee,omitempty" json:"payee,omitempty"`
	Date      *FHIRDateTime `bson:"date,omitempty" json:"date,omitempty"`
	Amount    *Quantity     `bson:"amount,omitempty" json:"amount,omitempty"`
}

type PaymentReconciliationNotesComponent struct {
	Type *Coding `bson:"type,omitempty" json:"type,omitempty"`
	Text string  `bson:"text,omitempty" json:"text,omitempty"`
}

type PaymentReconciliationBundle struct {
	Type         string                             `json:"resourceType,omitempty"`
	Title        string                             `json:"title,omitempty"`
	Id           string                             `json:"id,omitempty"`
	Updated      time.Time                          `json:"updated,omitempty"`
	TotalResults int                                `json:"totalResults,omitempty"`
	Entry        []PaymentReconciliationBundleEntry `json:"entry,omitempty"`
	Category     PaymentReconciliationCategory      `json:"category,omitempty"`
}

type PaymentReconciliationBundleEntry struct {
	Title    string                        `json:"title,omitempty"`
	Id       string                        `json:"id,omitempty"`
	Content  PaymentReconciliation         `json:"content,omitempty"`
	Category PaymentReconciliationCategory `json:"category,omitempty"`
}

type PaymentReconciliationCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}

func (resource *PaymentReconciliation) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		PaymentReconciliation
	}{
		ResourceType:          "PaymentReconciliation",
		PaymentReconciliation: *resource,
	}
	return json.Marshal(x)
}
