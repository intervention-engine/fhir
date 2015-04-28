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

type Group struct {
	Id             string                         `json:"-" bson:"_id"`
	Identifier     *Identifier                    `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Type           string                         `bson:"type,omitempty", json:"type,omitempty"`
	Actual         *bool                          `bson:"actual,omitempty", json:"actual,omitempty"`
	Code           *CodeableConcept               `bson:"code,omitempty", json:"code,omitempty"`
	Name           string                         `bson:"name,omitempty", json:"name,omitempty"`
	Quantity       float64                        `bson:"quantity,omitempty", json:"quantity,omitempty"`
	Characteristic []GroupCharacteristicComponent `bson:"characteristic,omitempty", json:"characteristic,omitempty"`
	Member         []Reference                    `bson:"member,omitempty", json:"member,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec characteristic
type GroupCharacteristicComponent struct {
	Code                 *CodeableConcept `bson:"code,omitempty", json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty", json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty", json:"valueBoolean,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty", json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty", json:"valueRange,omitempty"`
	Exclude              *bool            `bson:"exclude,omitempty", json:"exclude,omitempty"`
}

type GroupBundle struct {
	Type         string             `json:"resourceType,omitempty"`
	Title        string             `json:"title,omitempty"`
	Id           string             `json:"id,omitempty"`
	Updated      time.Time          `json:"updated,omitempty"`
	TotalResults int                `json:"totalResults,omitempty"`
	Entry        []GroupBundleEntry `json:"entry,omitempty"`
	Category     GroupCategory      `json:"category,omitempty"`
}

type GroupBundleEntry struct {
	Title    string        `json:"title,omitempty"`
	Id       string        `json:"id,omitempty"`
	Content  Group         `json:"content,omitempty"`
	Category GroupCategory `json:"category,omitempty"`
}

type GroupCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
