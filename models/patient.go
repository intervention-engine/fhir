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

type Patient struct {
	Id                   string                 `json:"-" bson:"_id"`
	Identifier           []Identifier           `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Name                 []HumanName            `bson:"name,omitempty", json:"name,omitempty"`
	Telecom              []ContactPoint         `bson:"telecom,omitempty", json:"telecom,omitempty"`
	Gender               CodeableConcept        `bson:"gender,omitempty", json:"gender,omitempty"`
	BirthDate            FHIRDateTime           `bson:"birthDate,omitempty", json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                  `bson:"deceasedBoolean,omitempty", json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     FHIRDateTime           `bson:"deceasedDateTime,omitempty", json:"deceasedDateTime,omitempty"`
	Address              []Address              `bson:"address,omitempty", json:"address,omitempty"`
	MaritalStatus        CodeableConcept        `bson:"maritalStatus,omitempty", json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *bool                  `bson:"multipleBirthBoolean,omitempty", json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger float64                `bson:"multipleBirthInteger,omitempty", json:"multipleBirthInteger,omitempty"`
	Photo                []Attachment           `bson:"photo,omitempty", json:"photo,omitempty"`
	Contact              []ContactComponent     `bson:"contact,omitempty", json:"contact,omitempty"`
	Animal               AnimalComponent        `bson:"animal,omitempty", json:"animal,omitempty"`
	Communication        []CodeableConcept      `bson:"communication,omitempty", json:"communication,omitempty"`
	CareProvider         []Reference            `bson:"careProvider,omitempty", json:"careProvider,omitempty"`
	ManagingOrganization Reference              `bson:"managingOrganization,omitempty", json:"managingOrganization,omitempty"`
	Link                 []PatientLinkComponent `bson:"link,omitempty", json:"link,omitempty"`
	Active               *bool                  `bson:"active,omitempty", json:"active,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec contact
type ContactComponent struct {
	Relationship []CodeableConcept `bson:"relationship,omitempty", json:"relationship,omitempty"`
	Name         HumanName         `bson:"name,omitempty", json:"name,omitempty"`
	Telecom      []ContactPoint    `bson:"telecom,omitempty", json:"telecom,omitempty"`
	Address      Address           `bson:"address,omitempty", json:"address,omitempty"`
	Gender       CodeableConcept   `bson:"gender,omitempty", json:"gender,omitempty"`
	Organization Reference         `bson:"organization,omitempty", json:"organization,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec animal
type AnimalComponent struct {
	Species      CodeableConcept `bson:"species,omitempty", json:"species,omitempty"`
	Breed        CodeableConcept `bson:"breed,omitempty", json:"breed,omitempty"`
	GenderStatus CodeableConcept `bson:"genderStatus,omitempty", json:"genderStatus,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec link
type PatientLinkComponent struct {
	Other Reference `bson:"other,omitempty", json:"other,omitempty"`
	Type  string    `bson:"type,omitempty", json:"type,omitempty"`
}
type PatientBundle struct {
	Type         string
	Title        string
	Id           string
	Updated      time.Time
	TotalResults int
	Entries      []Patient
	Category     PatientCategory
}

type PatientCategory struct {
	Term   string
	Label  string
	Scheme string
}
