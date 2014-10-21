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

type ValueSet struct {
	Id           string                     `json:"-" bson:"_id"`
	Identifier   string                     `bson:"identifier"`
	Version      string                     `bson:"version"`
	Name         string                     `bson:"name"`
	Purpose      string                     `bson:"purpose"`
	Immutable    bool                       `bson:"immutable"`
	Publisher    string                     `bson:"publisher"`
	Telecom      []ContactPoint             `bson:"telecom"`
	Description  string                     `bson:"description"`
	Copyright    string                     `bson:"copyright"`
	Status       string                     `bson:"status"`
	Experimental bool                       `bson:"experimental"`
	Extensible   bool                       `bson:"extensible"`
	Date         time.Time                  `bson:"date"`
	StableDate   time.Time                  `bson:"stableDate"`
	Define       ValueSetDefineComponent    `bson:"define"`
	Compose      ValueSetComposeComponent   `bson:"compose"`
	Expansion    ValueSetExpansionComponent `bson:"expansion"`
}

// This is an ugly hack to deal with embedded structures in the spec designation
type ConceptDefinitionDesignationComponent struct {
	Language string `bson:"language"`
	Use      Coding `bson:"use"`
	Value    string `bson:"value"`
}

// This is an ugly hack to deal with embedded structures in the spec concept
type ConceptDefinitionComponent struct {
	Code        string                                  `bson:"code"`
	Abstract    bool                                    `bson:"abstract"`
	Display     string                                  `bson:"display"`
	Definition  string                                  `bson:"definition"`
	Designation []ConceptDefinitionDesignationComponent `bson:"designation"`
	Concept     []ConceptDefinitionComponent            `bson:"concept"`
}

// This is an ugly hack to deal with embedded structures in the spec define
type ValueSetDefineComponent struct {
	System        string                       `bson:"system"`
	Version       string                       `bson:"version"`
	CaseSensitive bool                         `bson:"caseSensitive"`
	Concept       []ConceptDefinitionComponent `bson:"concept"`
}

// This is an ugly hack to deal with embedded structures in the spec concept
type ConceptReferenceComponent struct {
	Code        string                                  `bson:"code"`
	Display     string                                  `bson:"display"`
	Designation []ConceptDefinitionDesignationComponent `bson:"designation"`
}

// This is an ugly hack to deal with embedded structures in the spec filter
type ConceptSetFilterComponent struct {
	Property string `bson:"property"`
	Op       string `bson:"op"`
	Value    string `bson:"value"`
}

// This is an ugly hack to deal with embedded structures in the spec include
type ConceptSetComponent struct {
	System  string                      `bson:"system"`
	Version string                      `bson:"version"`
	Concept []ConceptReferenceComponent `bson:"concept"`
	Filter  []ConceptSetFilterComponent `bson:"filter"`
}

// This is an ugly hack to deal with embedded structures in the spec compose
type ValueSetComposeComponent struct {
	Import  []string              `bson:"import"`
	Include []ConceptSetComponent `bson:"include"`
	Exclude []ConceptSetComponent `bson:"exclude"`
}

// This is an ugly hack to deal with embedded structures in the spec contains
type ValueSetExpansionContainsComponent struct {
	System   string                               `bson:"system"`
	Abstract bool                                 `bson:"abstract"`
	Version  string                               `bson:"version"`
	Code     string                               `bson:"code"`
	Display  string                               `bson:"display"`
	Contains []ValueSetExpansionContainsComponent `bson:"contains"`
}

// This is an ugly hack to deal with embedded structures in the spec expansion
type ValueSetExpansionComponent struct {
	Identifier Identifier                           `bson:"identifier"`
	Timestamp  time.Time                            `bson:"timestamp"`
	Contains   []ValueSetExpansionContainsComponent `bson:"contains"`
}
