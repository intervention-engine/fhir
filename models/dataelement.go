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

type DataElement struct {
	Id           string                        `json:"-" bson:"_id"`
	Identifier   Identifier                    `bson:"identifier"`
	Version      string                        `bson:"version"`
	Publisher    string                        `bson:"publisher"`
	Telecom      []ContactPoint                `bson:"telecom"`
	Status       string                        `bson:"status"`
	Date         FHIRDateTime                  `bson:"date"`
	Name         string                        `bson:"name"`
	Category     []CodeableConcept             `bson:"category"`
	Code         []Coding                      `bson:"code"`
	Question     string                        `bson:"question"`
	Definition   string                        `bson:"definition"`
	Comments     string                        `bson:"comments"`
	Requirements string                        `bson:"requirements"`
	Synonym      []string                      `bson:"synonym"`
	Type         string                        `bson:"type"`
	MaxLength    float64                       `bson:"maxLength"`
	Units        CodeableConcept               `bson:"units"`
	Binding      DataElementBindingComponent   `bson:"binding"`
	Mapping      []DataElementMappingComponent `bson:"mapping"`
}

// This is an ugly hack to deal with embedded structures in the spec binding
type DataElementBindingComponent struct {
	IsExtensible bool      `bson:"isExtensible"`
	Conformance  string    `bson:"conformance"`
	Description  string    `bson:"description"`
	ValueSet     Reference `bson:"valueSet"`
}

// This is an ugly hack to deal with embedded structures in the spec mapping
type DataElementMappingComponent struct {
	Uri      string `bson:"uri"`
	Name     string `bson:"name"`
	Comments string `bson:"comments"`
	Map      string `bson:"map"`
}
