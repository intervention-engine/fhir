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

type Profile struct {
	Id            string                          `json:"-" bson:"_id"`
	Url           string                          `bson:"url"`
	Identifier    []Identifier                    `bson:"identifier"`
	Version       string                          `bson:"version"`
	Name          string                          `bson:"name"`
	Publisher     string                          `bson:"publisher"`
	Telecom       []ContactPoint                  `bson:"telecom"`
	Description   string                          `bson:"description"`
	Code          []Coding                        `bson:"code"`
	Status        string                          `bson:"status"`
	Experimental  bool                            `bson:"experimental"`
	Date          FHIRDateTime                    `bson:"date"`
	Requirements  string                          `bson:"requirements"`
	FhirVersion   string                          `bson:"fhirVersion"`
	Mapping       []ProfileMappingComponent       `bson:"mapping"`
	Structure     []ProfileStructureComponent     `bson:"structure"`
	ExtensionDefn []ProfileExtensionDefnComponent `bson:"extensionDefn"`
}

// This is an ugly hack to deal with embedded structures in the spec mapping
type ProfileMappingComponent struct {
	Identity string `bson:"identity"`
	Uri      string `bson:"uri"`
	Name     string `bson:"name"`
	Comments string `bson:"comments"`
}

// This is an ugly hack to deal with embedded structures in the spec slicing
type ElementSlicingComponent struct {
	Discriminator string `bson:"discriminator"`
	Ordered       bool   `bson:"ordered"`
	Rules         string `bson:"rules"`
}

// This is an ugly hack to deal with embedded structures in the spec type
type TypeRefComponent struct {
	Code        string   `bson:"code"`
	Profile     string   `bson:"profile"`
	Aggregation []string `bson:"aggregation"`
}

// This is an ugly hack to deal with embedded structures in the spec constraint
type ElementDefinitionConstraintComponent struct {
	Key      string `bson:"key"`
	Name     string `bson:"name"`
	Severity string `bson:"severity"`
	Human    string `bson:"human"`
	Xpath    string `bson:"xpath"`
}

// This is an ugly hack to deal with embedded structures in the spec binding
type ElementDefinitionBindingComponent struct {
	Name               string    `bson:"name"`
	IsExtensible       bool      `bson:"isExtensible"`
	Conformance        string    `bson:"conformance"`
	Description        string    `bson:"description"`
	ReferenceUri       string    `bson:"referenceUri"`
	ReferenceReference Reference `bson:"referenceReference"`
}

// This is an ugly hack to deal with embedded structures in the spec mapping
type ElementDefinitionMappingComponent struct {
	Identity string `bson:"identity"`
	Map      string `bson:"map"`
}

// This is an ugly hack to deal with embedded structures in the spec definition
type ElementDefinitionComponent struct {
	Short         string                                 `bson:"short"`
	Formal        string                                 `bson:"formal"`
	Comments      string                                 `bson:"comments"`
	Requirements  string                                 `bson:"requirements"`
	Synonym       []string                               `bson:"synonym"`
	Min           float64                                `bson:"min"`
	Max           string                                 `bson:"max"`
	Type          []TypeRefComponent                     `bson:"type"`
	NameReference string                                 `bson:"nameReference"`
	MaxLength     float64                                `bson:"maxLength"`
	Condition     []string                               `bson:"condition"`
	Constraint    []ElementDefinitionConstraintComponent `bson:"constraint"`
	MustSupport   bool                                   `bson:"mustSupport"`
	IsModifier    bool                                   `bson:"isModifier"`
	Binding       ElementDefinitionBindingComponent      `bson:"binding"`
	Mapping       []ElementDefinitionMappingComponent    `bson:"mapping"`
}

// This is an ugly hack to deal with embedded structures in the spec element
type ElementComponent struct {
	Path           string                     `bson:"path"`
	Representation []string                   `bson:"representation"`
	Name           string                     `bson:"name"`
	Slicing        ElementSlicingComponent    `bson:"slicing"`
	Definition     ElementDefinitionComponent `bson:"definition"`
}

// This is an ugly hack to deal with embedded structures in the spec snapshot
type ConstraintComponent struct {
	Element []ElementComponent `bson:"element"`
}

// This is an ugly hack to deal with embedded structures in the spec searchParam
type ProfileStructureSearchParamComponent struct {
	Name          string   `bson:"name"`
	Type          string   `bson:"type"`
	Documentation string   `bson:"documentation"`
	Xpath         string   `bson:"xpath"`
	Target        []string `bson:"target"`
}

// This is an ugly hack to deal with embedded structures in the spec structure
type ProfileStructureComponent struct {
	Type         string                                 `bson:"type"`
	Base         string                                 `bson:"base"`
	Name         string                                 `bson:"name"`
	Publish      bool                                   `bson:"publish"`
	Purpose      string                                 `bson:"purpose"`
	Snapshot     ConstraintComponent                    `bson:"snapshot"`
	Differential ConstraintComponent                    `bson:"differential"`
	SearchParam  []ProfileStructureSearchParamComponent `bson:"searchParam"`
}

// This is an ugly hack to deal with embedded structures in the spec extensionDefn
type ProfileExtensionDefnComponent struct {
	Code        string             `bson:"code"`
	Display     string             `bson:"display"`
	ContextType string             `bson:"contextType"`
	Context     []string           `bson:"context"`
	Element     []ElementComponent `bson:"element"`
}
