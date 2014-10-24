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

type OperationDefinition struct {
	Id           string                                  `json:"-" bson:"_id"`
	Identifier   string                                  `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Version      string                                  `bson:"version,omitempty", json:"version,omitempty"`
	Title        string                                  `bson:"title,omitempty", json:"title,omitempty"`
	Publisher    string                                  `bson:"publisher,omitempty", json:"publisher,omitempty"`
	Telecom      []ContactPoint                          `bson:"telecom,omitempty", json:"telecom,omitempty"`
	Description  string                                  `bson:"description,omitempty", json:"description,omitempty"`
	Code         []Coding                                `bson:"code,omitempty", json:"code,omitempty"`
	Status       string                                  `bson:"status,omitempty", json:"status,omitempty"`
	Experimental bool                                    `bson:"experimental,omitempty", json:"experimental,omitempty"`
	Date         FHIRDateTime                            `bson:"date,omitempty", json:"date,omitempty"`
	Kind         string                                  `bson:"kind,omitempty", json:"kind,omitempty"`
	Name         string                                  `bson:"name,omitempty", json:"name,omitempty"`
	Notes        string                                  `bson:"notes,omitempty", json:"notes,omitempty"`
	Base         Reference                               `bson:"base,omitempty", json:"base,omitempty"`
	System       bool                                    `bson:"system,omitempty", json:"system,omitempty"`
	Type         []string                                `bson:"type,omitempty", json:"type,omitempty"`
	Instance     bool                                    `bson:"instance,omitempty", json:"instance,omitempty"`
	Parameter    []OperationDefinitionParameterComponent `bson:"parameter,omitempty", json:"parameter,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec parameter
type OperationDefinitionParameterComponent struct {
	Name          string    `bson:"name,omitempty", json:"name,omitempty"`
	Use           string    `bson:"use,omitempty", json:"use,omitempty"`
	Min           float64   `bson:"min,omitempty", json:"min,omitempty"`
	Max           string    `bson:"max,omitempty", json:"max,omitempty"`
	Documentation string    `bson:"documentation,omitempty", json:"documentation,omitempty"`
	Type          Coding    `bson:"type,omitempty", json:"type,omitempty"`
	Profile       Reference `bson:"profile,omitempty", json:"profile,omitempty"`
}
