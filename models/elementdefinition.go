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

type ElementDefinition struct {
	Id                          string                                 `json:"-" bson:"_id"`
	Path                        string                                 `bson:"path,omitempty" json:"path,omitempty"`
	Representation              []string                               `bson:"representation,omitempty" json:"representation,omitempty"`
	Name                        string                                 `bson:"name,omitempty" json:"name,omitempty"`
	Label                       string                                 `bson:"label,omitempty" json:"label,omitempty"`
	Code                        []Coding                               `bson:"code,omitempty" json:"code,omitempty"`
	Slicing                     *ElementDefinitionSlicingComponent     `bson:"slicing,omitempty" json:"slicing,omitempty"`
	Short                       string                                 `bson:"short,omitempty" json:"short,omitempty"`
	Definition                  string                                 `bson:"definition,omitempty" json:"definition,omitempty"`
	Comments                    string                                 `bson:"comments,omitempty" json:"comments,omitempty"`
	Requirements                string                                 `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Alias                       []string                               `bson:"alias,omitempty" json:"alias,omitempty"`
	Min                         *int32                                 `bson:"min,omitempty" json:"min,omitempty"`
	Max                         string                                 `bson:"max,omitempty" json:"max,omitempty"`
	Type                        []ElementDefinitionTypeRefComponent    `bson:"type,omitempty" json:"type,omitempty"`
	NameReference               string                                 `bson:"nameReference,omitempty" json:"nameReference,omitempty"`
	DefaultValueString          string                                 `bson:"defaultValueString,omitempty" json:"defaultValueString,omitempty"`
	DefaultValueInteger         *int32                                 `bson:"defaultValueInteger,omitempty" json:"defaultValueInteger,omitempty"`
	DefaultValueDateTime        *FHIRDateTime                          `bson:"defaultValueDateTime,omitempty" json:"defaultValueDateTime,omitempty"`
	DefaultValueBoolean         *bool                                  `bson:"defaultValueBoolean,omitempty" json:"defaultValueBoolean,omitempty"`
	DefaultValueCodeableConcept *CodeableConcept                       `bson:"defaultValueCodeableConcept,omitempty" json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueRange           *Range                                 `bson:"defaultValueRange,omitempty" json:"defaultValueRange,omitempty"`
	MeaningWhenMissing          string                                 `bson:"meaningWhenMissing,omitempty" json:"meaningWhenMissing,omitempty"`
	FixedString                 string                                 `bson:"fixedString,omitempty" json:"fixedString,omitempty"`
	FixedInteger                *int32                                 `bson:"fixedInteger,omitempty" json:"fixedInteger,omitempty"`
	FixedDateTime               *FHIRDateTime                          `bson:"fixedDateTime,omitempty" json:"fixedDateTime,omitempty"`
	FixedBoolean                *bool                                  `bson:"fixedBoolean,omitempty" json:"fixedBoolean,omitempty"`
	FixedCodeableConcept        *CodeableConcept                       `bson:"fixedCodeableConcept,omitempty" json:"fixedCodeableConcept,omitempty"`
	FixedRange                  *Range                                 `bson:"fixedRange,omitempty" json:"fixedRange,omitempty"`
	PatternString               string                                 `bson:"patternString,omitempty" json:"patternString,omitempty"`
	PatternInteger              *int32                                 `bson:"patternInteger,omitempty" json:"patternInteger,omitempty"`
	PatternDateTime             *FHIRDateTime                          `bson:"patternDateTime,omitempty" json:"patternDateTime,omitempty"`
	PatternBoolean              *bool                                  `bson:"patternBoolean,omitempty" json:"patternBoolean,omitempty"`
	PatternCodeableConcept      *CodeableConcept                       `bson:"patternCodeableConcept,omitempty" json:"patternCodeableConcept,omitempty"`
	PatternRange                *Range                                 `bson:"patternRange,omitempty" json:"patternRange,omitempty"`
	ExampleString               string                                 `bson:"exampleString,omitempty" json:"exampleString,omitempty"`
	ExampleInteger              *int32                                 `bson:"exampleInteger,omitempty" json:"exampleInteger,omitempty"`
	ExampleDateTime             *FHIRDateTime                          `bson:"exampleDateTime,omitempty" json:"exampleDateTime,omitempty"`
	ExampleBoolean              *bool                                  `bson:"exampleBoolean,omitempty" json:"exampleBoolean,omitempty"`
	ExampleCodeableConcept      *CodeableConcept                       `bson:"exampleCodeableConcept,omitempty" json:"exampleCodeableConcept,omitempty"`
	ExampleRange                *Range                                 `bson:"exampleRange,omitempty" json:"exampleRange,omitempty"`
	MaxLength                   *int32                                 `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	Condition                   []string                               `bson:"condition,omitempty" json:"condition,omitempty"`
	Constraint                  []ElementDefinitionConstraintComponent `bson:"constraint,omitempty" json:"constraint,omitempty"`
	MustSupport                 *bool                                  `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	IsModifier                  *bool                                  `bson:"isModifier,omitempty" json:"isModifier,omitempty"`
	IsSummary                   *bool                                  `bson:"isSummary,omitempty" json:"isSummary,omitempty"`
	Binding                     *ElementDefinitionBindingComponent     `bson:"binding,omitempty" json:"binding,omitempty"`
	Mapping                     []ElementDefinitionMappingComponent    `bson:"mapping,omitempty" json:"mapping,omitempty"`
}

type ElementDefinitionSlicingComponent struct {
	Discriminator []string `bson:"discriminator,omitempty" json:"discriminator,omitempty"`
	Description   string   `bson:"description,omitempty" json:"description,omitempty"`
	Ordered       *bool    `bson:"ordered,omitempty" json:"ordered,omitempty"`
	Rules         string   `bson:"rules,omitempty" json:"rules,omitempty"`
}

type ElementDefinitionTypeRefComponent struct {
	Code        string   `bson:"code,omitempty" json:"code,omitempty"`
	Profile     []string `bson:"profile,omitempty" json:"profile,omitempty"`
	Aggregation []string `bson:"aggregation,omitempty" json:"aggregation,omitempty"`
}

type ElementDefinitionConstraintComponent struct {
	Key      string `bson:"key,omitempty" json:"key,omitempty"`
	Name     string `bson:"name,omitempty" json:"name,omitempty"`
	Severity string `bson:"severity,omitempty" json:"severity,omitempty"`
	Human    string `bson:"human,omitempty" json:"human,omitempty"`
	Xpath    string `bson:"xpath,omitempty" json:"xpath,omitempty"`
}

type ElementDefinitionBindingComponent struct {
	Name              string     `bson:"name,omitempty" json:"name,omitempty"`
	Strength          string     `bson:"strength,omitempty" json:"strength,omitempty"`
	Description       string     `bson:"description,omitempty" json:"description,omitempty"`
	ValueSetUri       string     `bson:"valueSetUri,omitempty" json:"valueSetUri,omitempty"`
	ValueSetReference *Reference `bson:"valueSetReference,omitempty" json:"valueSetReference,omitempty"`
}

type ElementDefinitionMappingComponent struct {
	Identity string `bson:"identity,omitempty" json:"identity,omitempty"`
	Language string `bson:"language,omitempty" json:"language,omitempty"`
	Map      string `bson:"map,omitempty" json:"map,omitempty"`
}
