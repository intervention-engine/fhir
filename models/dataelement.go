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

type DataElement struct {
	Id                    string                        `json:"-" bson:"_id"`
	Identifier            Identifier                    `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Version               string                        `bson:"version,omitempty", json:"version,omitempty"`
	Publisher             string                        `bson:"publisher,omitempty", json:"publisher,omitempty"`
	Telecom               []ContactPoint                `bson:"telecom,omitempty", json:"telecom,omitempty"`
	Status                string                        `bson:"status,omitempty", json:"status,omitempty"`
	Date                  FHIRDateTime                  `bson:"date,omitempty", json:"date,omitempty"`
	Name                  string                        `bson:"name,omitempty", json:"name,omitempty"`
	Category              []CodeableConcept             `bson:"category,omitempty", json:"category,omitempty"`
	Code                  []Coding                      `bson:"code,omitempty", json:"code,omitempty"`
	Question              string                        `bson:"question,omitempty", json:"question,omitempty"`
	Definition            string                        `bson:"definition,omitempty", json:"definition,omitempty"`
	Comments              string                        `bson:"comments,omitempty", json:"comments,omitempty"`
	Requirements          string                        `bson:"requirements,omitempty", json:"requirements,omitempty"`
	Synonym               []string                      `bson:"synonym,omitempty", json:"synonym,omitempty"`
	Type                  string                        `bson:"type,omitempty", json:"type,omitempty"`
	ExampleString         string                        `bson:"examplestring,omitempty", json:"examplestring,omitempty"`
	ExampleInteger        int                           `bson:"exampleinteger,omitempty", json:"exampleinteger,omitempty"`
	ExampleDateTime       FHIRDateTime                  `bson:"exampledatetime,omitempty", json:"exampledatetime,omitempty"`
	ExampleBoolean        *bool                         `bson:"exampleboolean,omitempty", json:"exampleboolean,omitempty"`
	ExampleCodableConcept CodeableConcept               `bson:"examplecodableconcept,omitempty", json:"examplecodableconcept,omitempty"`
	MaxLength             float64                       `bson:"maxLength,omitempty", json:"maxLength,omitempty"`
	Units                 CodeableConcept               `bson:"units,omitempty", json:"units,omitempty"`
	Binding               DataElementBindingComponent   `bson:"binding,omitempty", json:"binding,omitempty"`
	Mapping               []DataElementMappingComponent `bson:"mapping,omitempty", json:"mapping,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec binding
type DataElementBindingComponent struct {
	IsExtensible *bool     `bson:"isExtensible,omitempty", json:"isExtensible,omitempty"`
	Conformance  string    `bson:"conformance,omitempty", json:"conformance,omitempty"`
	Description  string    `bson:"description,omitempty", json:"description,omitempty"`
	ValueSet     Reference `bson:"valueSet,omitempty", json:"valueSet,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec mapping
type DataElementMappingComponent struct {
	Uri      string `bson:"uri,omitempty", json:"uri,omitempty"`
	Name     string `bson:"name,omitempty", json:"name,omitempty"`
	Comments string `bson:"comments,omitempty", json:"comments,omitempty"`
	Map      string `bson:"map,omitempty", json:"map,omitempty"`
}
type DataElementBundle struct {
	Type         string
	Title        string
	Id           string
	Updated      time.Time
	TotalResults int
	Entries      []DataElement
	Category     DataElementCategory
}

type DataElementCategory struct {
	Term   string
	Label  string
	Scheme string
}
