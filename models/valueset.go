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

type ValueSet struct {
	Id           string                      `json:"-" bson:"_id"`
	Url          string                      `bson:"url,omitempty" json:"url,omitempty"`
	Identifier   *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version      string                      `bson:"version,omitempty" json:"version,omitempty"`
	Name         string                      `bson:"name,omitempty" json:"name,omitempty"`
	UseContext   []CodeableConcept           `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Immutable    *bool                       `bson:"immutable,omitempty" json:"immutable,omitempty"`
	Publisher    string                      `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact      []ValueSetContactComponent  `bson:"contact,omitempty" json:"contact,omitempty"`
	Description  string                      `bson:"description,omitempty" json:"description,omitempty"`
	Requirements string                      `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright    string                      `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Status       string                      `bson:"status,omitempty" json:"status,omitempty"`
	Experimental *bool                       `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Extensible   *bool                       `bson:"extensible,omitempty" json:"extensible,omitempty"`
	Date         *FHIRDateTime               `bson:"date,omitempty" json:"date,omitempty"`
	LockedDate   *FHIRDateTime               `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	Define       *ValueSetDefineComponent    `bson:"define,omitempty" json:"define,omitempty"`
	Compose      *ValueSetComposeComponent   `bson:"compose,omitempty" json:"compose,omitempty"`
	Expansion    *ValueSetExpansionComponent `bson:"expansion,omitempty" json:"expansion,omitempty"`
}

type ValueSetContactComponent struct {
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type ValueSetDefineComponent struct {
	System        string                               `bson:"system,omitempty" json:"system,omitempty"`
	Version       string                               `bson:"version,omitempty" json:"version,omitempty"`
	CaseSensitive *bool                                `bson:"caseSensitive,omitempty" json:"caseSensitive,omitempty"`
	Concept       []ValueSetConceptDefinitionComponent `bson:"concept,omitempty" json:"concept,omitempty"`
}

type ValueSetConceptDefinitionComponent struct {
	Code        string                                          `bson:"code,omitempty" json:"code,omitempty"`
	Abstract    *bool                                           `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Display     string                                          `bson:"display,omitempty" json:"display,omitempty"`
	Definition  string                                          `bson:"definition,omitempty" json:"definition,omitempty"`
	Designation []ValueSetConceptDefinitionDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
	Concept     []ValueSetConceptDefinitionComponent            `bson:"concept,omitempty" json:"concept,omitempty"`
}

type ValueSetConceptDefinitionDesignationComponent struct {
	Language string  `bson:"language,omitempty" json:"language,omitempty"`
	Use      *Coding `bson:"use,omitempty" json:"use,omitempty"`
	Value    string  `bson:"value,omitempty" json:"value,omitempty"`
}

type ValueSetComposeComponent struct {
	Import  []string                      `bson:"import,omitempty" json:"import,omitempty"`
	Include []ValueSetConceptSetComponent `bson:"include,omitempty" json:"include,omitempty"`
	Exclude []ValueSetConceptSetComponent `bson:"exclude,omitempty" json:"exclude,omitempty"`
}

type ValueSetConceptSetComponent struct {
	System  string                              `bson:"system,omitempty" json:"system,omitempty"`
	Version string                              `bson:"version,omitempty" json:"version,omitempty"`
	Concept []ValueSetConceptReferenceComponent `bson:"concept,omitempty" json:"concept,omitempty"`
	Filter  []ValueSetConceptSetFilterComponent `bson:"filter,omitempty" json:"filter,omitempty"`
}

type ValueSetConceptReferenceComponent struct {
	Code        string                                          `bson:"code,omitempty" json:"code,omitempty"`
	Display     string                                          `bson:"display,omitempty" json:"display,omitempty"`
	Designation []ValueSetConceptDefinitionDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
}

type ValueSetConceptSetFilterComponent struct {
	Property string `bson:"property,omitempty" json:"property,omitempty"`
	Op       string `bson:"op,omitempty" json:"op,omitempty"`
	Value    string `bson:"value,omitempty" json:"value,omitempty"`
}

type ValueSetExpansionComponent struct {
	Identifier string                                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Timestamp  *FHIRDateTime                         `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Parameter  []ValueSetExpansionParameterComponent `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Contains   []ValueSetExpansionContainsComponent  `bson:"contains,omitempty" json:"contains,omitempty"`
}

type ValueSetExpansionParameterComponent struct {
	Name         string   `bson:"name,omitempty" json:"name,omitempty"`
	ValueString  string   `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean *bool    `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger *int32   `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDecimal *float64 `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueUri     string   `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueCode    string   `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
}

type ValueSetExpansionContainsComponent struct {
	System   string                               `bson:"system,omitempty" json:"system,omitempty"`
	Abstract *bool                                `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Version  string                               `bson:"version,omitempty" json:"version,omitempty"`
	Code     string                               `bson:"code,omitempty" json:"code,omitempty"`
	Display  string                               `bson:"display,omitempty" json:"display,omitempty"`
	Contains []ValueSetExpansionContainsComponent `bson:"contains,omitempty" json:"contains,omitempty"`
}

type ValueSetBundle struct {
	Type         string                `json:"resourceType,omitempty"`
	Title        string                `json:"title,omitempty"`
	Id           string                `json:"id,omitempty"`
	Updated      time.Time             `json:"updated,omitempty"`
	TotalResults int                   `json:"totalResults,omitempty"`
	Entry        []ValueSetBundleEntry `json:"entry,omitempty"`
	Category     ValueSetCategory      `json:"category,omitempty"`
}

type ValueSetBundleEntry struct {
	Title    string           `json:"title,omitempty"`
	Id       string           `json:"id,omitempty"`
	Content  ValueSet         `json:"content,omitempty"`
	Category ValueSetCategory `json:"category,omitempty"`
}

type ValueSetCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}

func (resource *ValueSet) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		ValueSet
	}{
		ResourceType: "ValueSet",
		ValueSet:     *resource,
	}
	return json.Marshal(x)
}
