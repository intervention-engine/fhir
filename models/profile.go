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

type Profile struct {
	Id            string                          `json:"-" bson:"_id"`
	Url           string                          `bson:"url,omitempty", json:"url,omitempty"`
	Identifier    []Identifier                    `bson:"identifier,omitempty", json:"identifier,omitempty"`
	Version       string                          `bson:"version,omitempty", json:"version,omitempty"`
	Name          string                          `bson:"name,omitempty", json:"name,omitempty"`
	Publisher     string                          `bson:"publisher,omitempty", json:"publisher,omitempty"`
	Telecom       []ContactPoint                  `bson:"telecom,omitempty", json:"telecom,omitempty"`
	Description   string                          `bson:"description,omitempty", json:"description,omitempty"`
	Code          []Coding                        `bson:"code,omitempty", json:"code,omitempty"`
	Status        string                          `bson:"status,omitempty", json:"status,omitempty"`
	Experimental  *bool                           `bson:"experimental,omitempty", json:"experimental,omitempty"`
	Date          FHIRDateTime                    `bson:"date,omitempty", json:"date,omitempty"`
	Requirements  string                          `bson:"requirements,omitempty", json:"requirements,omitempty"`
	FhirVersion   string                          `bson:"fhirVersion,omitempty", json:"fhirVersion,omitempty"`
	Mapping       []ProfileMappingComponent       `bson:"mapping,omitempty", json:"mapping,omitempty"`
	Structure     []ProfileStructureComponent     `bson:"structure,omitempty", json:"structure,omitempty"`
	ExtensionDefn []ProfileExtensionDefnComponent `bson:"extensionDefn,omitempty", json:"extensionDefn,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec mapping
type ProfileMappingComponent struct {
	Identity string `bson:"identity,omitempty", json:"identity,omitempty"`
	Uri      string `bson:"uri,omitempty", json:"uri,omitempty"`
	Name     string `bson:"name,omitempty", json:"name,omitempty"`
	Comments string `bson:"comments,omitempty", json:"comments,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec slicing
type ElementSlicingComponent struct {
	Discriminator string `bson:"discriminator,omitempty", json:"discriminator,omitempty"`
	Ordered       *bool  `bson:"ordered,omitempty", json:"ordered,omitempty"`
	Rules         string `bson:"rules,omitempty", json:"rules,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec type
type TypeRefComponent struct {
	Code        string   `bson:"code,omitempty", json:"code,omitempty"`
	Profile     string   `bson:"profile,omitempty", json:"profile,omitempty"`
	Aggregation []string `bson:"aggregation,omitempty", json:"aggregation,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec constraint
type ElementDefinitionConstraintComponent struct {
	Key      string `bson:"key,omitempty", json:"key,omitempty"`
	Name     string `bson:"name,omitempty", json:"name,omitempty"`
	Severity string `bson:"severity,omitempty", json:"severity,omitempty"`
	Human    string `bson:"human,omitempty", json:"human,omitempty"`
	Xpath    string `bson:"xpath,omitempty", json:"xpath,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec binding
type ElementDefinitionBindingComponent struct {
	Name               string    `bson:"name,omitempty", json:"name,omitempty"`
	IsExtensible       *bool     `bson:"isExtensible,omitempty", json:"isExtensible,omitempty"`
	Conformance        string    `bson:"conformance,omitempty", json:"conformance,omitempty"`
	Description        string    `bson:"description,omitempty", json:"description,omitempty"`
	ReferenceUri       string    `bson:"referenceUri,omitempty", json:"referenceUri,omitempty"`
	ReferenceReference Reference `bson:"referenceReference,omitempty", json:"referenceReference,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec mapping
type ElementDefinitionMappingComponent struct {
	Identity string `bson:"identity,omitempty", json:"identity,omitempty"`
	Map      string `bson:"map,omitempty", json:"map,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec definition
type ElementDefinitionComponent struct {
	Short                  string                                 `bson:"short,omitempty", json:"short,omitempty"`
	Formal                 string                                 `bson:"formal,omitempty", json:"formal,omitempty"`
	Comments               string                                 `bson:"comments,omitempty", json:"comments,omitempty"`
	Requirements           string                                 `bson:"requirements,omitempty", json:"requirements,omitempty"`
	Synonym                []string                               `bson:"synonym,omitempty", json:"synonym,omitempty"`
	Min                    float64                                `bson:"min,omitempty", json:"min,omitempty"`
	Max                    string                                 `bson:"max,omitempty", json:"max,omitempty"`
	Type                   []TypeRefComponent                     `bson:"type,omitempty", json:"type,omitempty"`
	NameReference          string                                 `bson:"nameReference,omitempty", json:"nameReference,omitempty"`
	ValueString            string                                 `bson:"valuestring,omitempty", json:"valuestring,omitempty"`
	ValueInteger           int                                    `bson:"valueinteger,omitempty", json:"valueinteger,omitempty"`
	ValueDateTime          FHIRDateTime                           `bson:"valuedatetime,omitempty", json:"valuedatetime,omitempty"`
	ValueBoolean           *bool                                  `bson:"valueboolean,omitempty", json:"valueboolean,omitempty"`
	ValueCodeableConcept   CodeableConcept                        `bson:"valuecodeableconcept,omitempty", json:"valuecodeableconcept,omitempty"`
	ValueRange             Range                                  `bson:"valuerange,omitempty", json:"valuerange,omitempty"`
	ExampleString          string                                 `bson:"examplestring,omitempty", json:"examplestring,omitempty"`
	ExampleInteger         int                                    `bson:"exampleinteger,omitempty", json:"exampleinteger,omitempty"`
	ExampleDateTime        FHIRDateTime                           `bson:"exampledatetime,omitempty", json:"exampledatetime,omitempty"`
	ExampleBoolean         *bool                                  `bson:"exampleboolean,omitempty", json:"exampleboolean,omitempty"`
	ExampleCodeableConcept CodeableConcept                        `bson:"examplecodeableconcept,omitempty", json:"examplecodeableconcept,omitempty"`
	ExampleRange           Range                                  `bson:"examplerange,omitempty", json:"examplerange,omitempty"`
	MaxLength              float64                                `bson:"maxLength,omitempty", json:"maxLength,omitempty"`
	Condition              []string                               `bson:"condition,omitempty", json:"condition,omitempty"`
	Constraint             []ElementDefinitionConstraintComponent `bson:"constraint,omitempty", json:"constraint,omitempty"`
	MustSupport            *bool                                  `bson:"mustSupport,omitempty", json:"mustSupport,omitempty"`
	IsModifier             *bool                                  `bson:"isModifier,omitempty", json:"isModifier,omitempty"`
	Binding                ElementDefinitionBindingComponent      `bson:"binding,omitempty", json:"binding,omitempty"`
	Mapping                []ElementDefinitionMappingComponent    `bson:"mapping,omitempty", json:"mapping,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec element
type ElementComponent struct {
	Path           string                     `bson:"path,omitempty", json:"path,omitempty"`
	Representation []string                   `bson:"representation,omitempty", json:"representation,omitempty"`
	Name           string                     `bson:"name,omitempty", json:"name,omitempty"`
	Slicing        ElementSlicingComponent    `bson:"slicing,omitempty", json:"slicing,omitempty"`
	Definition     ElementDefinitionComponent `bson:"definition,omitempty", json:"definition,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec snapshot
type ConstraintComponent struct {
	Element []ElementComponent `bson:"element,omitempty", json:"element,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec searchParam
type ProfileStructureSearchParamComponent struct {
	Name          string   `bson:"name,omitempty", json:"name,omitempty"`
	Type          string   `bson:"type,omitempty", json:"type,omitempty"`
	Documentation string   `bson:"documentation,omitempty", json:"documentation,omitempty"`
	Xpath         string   `bson:"xpath,omitempty", json:"xpath,omitempty"`
	Target        []string `bson:"target,omitempty", json:"target,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec structure
type ProfileStructureComponent struct {
	Type         string                                 `bson:"type,omitempty", json:"type,omitempty"`
	Base         string                                 `bson:"base,omitempty", json:"base,omitempty"`
	Name         string                                 `bson:"name,omitempty", json:"name,omitempty"`
	Publish      *bool                                  `bson:"publish,omitempty", json:"publish,omitempty"`
	Purpose      string                                 `bson:"purpose,omitempty", json:"purpose,omitempty"`
	Snapshot     ConstraintComponent                    `bson:"snapshot,omitempty", json:"snapshot,omitempty"`
	Differential ConstraintComponent                    `bson:"differential,omitempty", json:"differential,omitempty"`
	SearchParam  []ProfileStructureSearchParamComponent `bson:"searchParam,omitempty", json:"searchParam,omitempty"`
}

// This is an ugly hack to deal with embedded structures in the spec extensionDefn
type ProfileExtensionDefnComponent struct {
	Code        string             `bson:"code,omitempty", json:"code,omitempty"`
	Display     string             `bson:"display,omitempty", json:"display,omitempty"`
	ContextType string             `bson:"contextType,omitempty", json:"contextType,omitempty"`
	Context     []string           `bson:"context,omitempty", json:"context,omitempty"`
	Element     []ElementComponent `bson:"element,omitempty", json:"element,omitempty"`
}

type ProfileBundle struct {
	Type         string               `json:"resourceType,omitempty"`
	Title        string               `json:"title,omitempty"`
	Id           string               `json:"id,omitempty"`
	Updated      time.Time            `json:"updated,omitempty"`
	TotalResults int                  `json:"totalResults,omitempty"`
	Entry        []ProfileBundleEntry `json:"entry,omitempty"`
	Category     ProfileCategory      `json:"category,omitempty"`
}

type ProfileBundleEntry struct {
	Title    string          `json:"title,omitempty"`
	Id       string          `json:"id,omitempty"`
	Content  Profile         `json:"content,omitempty"`
	Category ProfileCategory `json:"category,omitempty"`
}

type ProfileCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
