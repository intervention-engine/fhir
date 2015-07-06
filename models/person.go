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

type Person struct {
	Id                   string                `json:"-" bson:"_id"`
	Identifier           []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name                 []HumanName           `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []ContactPoint        `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               string                `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *FHIRDateTime         `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address              []Address             `bson:"address,omitempty" json:"address,omitempty"`
	Photo                *Attachment           `bson:"photo,omitempty" json:"photo,omitempty"`
	ManagingOrganization *Reference            `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Active               *bool                 `bson:"active,omitempty" json:"active,omitempty"`
	Link                 []PersonLinkComponent `bson:"link,omitempty" json:"link,omitempty"`
}

type PersonLinkComponent struct {
	Target    *Reference `bson:"target,omitempty" json:"target,omitempty"`
	Assurance string     `bson:"assurance,omitempty" json:"assurance,omitempty"`
}

type PersonBundle struct {
	Type         string              `json:"resourceType,omitempty"`
	Title        string              `json:"title,omitempty"`
	Id           string              `json:"id,omitempty"`
	Updated      time.Time           `json:"updated,omitempty"`
	TotalResults int                 `json:"totalResults,omitempty"`
	Entry        []PersonBundleEntry `json:"entry,omitempty"`
	Category     PersonCategory      `json:"category,omitempty"`
}

type PersonBundleEntry struct {
	Title    string         `json:"title,omitempty"`
	Id       string         `json:"id,omitempty"`
	Content  Person         `json:"content,omitempty"`
	Category PersonCategory `json:"category,omitempty"`
}

type PersonCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}

func (resource *Person) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Person
	}{
		ResourceType: "Person",
		Person:       *resource,
	}
	return json.Marshal(x)
}