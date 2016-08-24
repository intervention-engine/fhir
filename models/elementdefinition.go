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
	Base                        *ElementDefinitionBaseComponent        `bson:"base,omitempty" json:"base,omitempty"`
	ContentReference            string                                 `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
	Type                        []ElementDefinitionTypeRefComponent    `bson:"type,omitempty" json:"type,omitempty"`
	DefaultValueAddress         *Address                               `bson:"defaultValueAddress,omitempty" json:"defaultValueAddress,omitempty"`
	DefaultValueAnnotation      *Annotation                            `bson:"defaultValueAnnotation,omitempty" json:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment      *Attachment                            `bson:"defaultValueAttachment,omitempty" json:"defaultValueAttachment,omitempty"`
	DefaultValueBase64Binary    string                                 `bson:"defaultValueBase64Binary,omitempty" json:"defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean         *bool                                  `bson:"defaultValueBoolean,omitempty" json:"defaultValueBoolean,omitempty"`
	DefaultValueCode            string                                 `bson:"defaultValueCode,omitempty" json:"defaultValueCode,omitempty"`
	DefaultValueCodeableConcept *CodeableConcept                       `bson:"defaultValueCodeableConcept,omitempty" json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCoding          *Coding                                `bson:"defaultValueCoding,omitempty" json:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint    *ContactPoint                          `bson:"defaultValueContactPoint,omitempty" json:"defaultValueContactPoint,omitempty"`
	DefaultValueDate            *FHIRDateTime                          `bson:"defaultValueDate,omitempty" json:"defaultValueDate,omitempty"`
	DefaultValueDateTime        *FHIRDateTime                          `bson:"defaultValueDateTime,omitempty" json:"defaultValueDateTime,omitempty"`
	DefaultValueDecimal         *float64                               `bson:"defaultValueDecimal,omitempty" json:"defaultValueDecimal,omitempty"`
	DefaultValueHumanName       *HumanName                             `bson:"defaultValueHumanName,omitempty" json:"defaultValueHumanName,omitempty"`
	DefaultValueId              string                                 `bson:"defaultValueId,omitempty" json:"defaultValueId,omitempty"`
	DefaultValueIdentifier      *Identifier                            `bson:"defaultValueIdentifier,omitempty" json:"defaultValueIdentifier,omitempty"`
	DefaultValueInstant         *FHIRDateTime                          `bson:"defaultValueInstant,omitempty" json:"defaultValueInstant,omitempty"`
	DefaultValueInteger         *int32                                 `bson:"defaultValueInteger,omitempty" json:"defaultValueInteger,omitempty"`
	DefaultValueMarkdown        string                                 `bson:"defaultValueMarkdown,omitempty" json:"defaultValueMarkdown,omitempty"`
	DefaultValueMeta            *Meta                                  `bson:"defaultValueMeta,omitempty" json:"defaultValueMeta,omitempty"`
	DefaultValueOid             string                                 `bson:"defaultValueOid,omitempty" json:"defaultValueOid,omitempty"`
	DefaultValuePeriod          *Period                                `bson:"defaultValuePeriod,omitempty" json:"defaultValuePeriod,omitempty"`
	DefaultValuePositiveInt     *uint32                                `bson:"defaultValuePositiveInt,omitempty" json:"defaultValuePositiveInt,omitempty"`
	DefaultValueQuantity        *Quantity                              `bson:"defaultValueQuantity,omitempty" json:"defaultValueQuantity,omitempty"`
	DefaultValueRange           *Range                                 `bson:"defaultValueRange,omitempty" json:"defaultValueRange,omitempty"`
	DefaultValueRatio           *Ratio                                 `bson:"defaultValueRatio,omitempty" json:"defaultValueRatio,omitempty"`
	DefaultValueReference       *Reference                             `bson:"defaultValueReference,omitempty" json:"defaultValueReference,omitempty"`
	DefaultValueSampledData     *SampledData                           `bson:"defaultValueSampledData,omitempty" json:"defaultValueSampledData,omitempty"`
	DefaultValueSignature       *Signature                             `bson:"defaultValueSignature,omitempty" json:"defaultValueSignature,omitempty"`
	DefaultValueString          string                                 `bson:"defaultValueString,omitempty" json:"defaultValueString,omitempty"`
	DefaultValueTime            *FHIRDateTime                          `bson:"defaultValueTime,omitempty" json:"defaultValueTime,omitempty"`
	DefaultValueTiming          *Timing                                `bson:"defaultValueTiming,omitempty" json:"defaultValueTiming,omitempty"`
	DefaultValueUnsignedInt     *uint32                                `bson:"defaultValueUnsignedInt,omitempty" json:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUri             string                                 `bson:"defaultValueUri,omitempty" json:"defaultValueUri,omitempty"`
	MeaningWhenMissing          string                                 `bson:"meaningWhenMissing,omitempty" json:"meaningWhenMissing,omitempty"`
	FixedAddress                *Address                               `bson:"fixedAddress,omitempty" json:"fixedAddress,omitempty"`
	FixedAnnotation             *Annotation                            `bson:"fixedAnnotation,omitempty" json:"fixedAnnotation,omitempty"`
	FixedAttachment             *Attachment                            `bson:"fixedAttachment,omitempty" json:"fixedAttachment,omitempty"`
	FixedBase64Binary           string                                 `bson:"fixedBase64Binary,omitempty" json:"fixedBase64Binary,omitempty"`
	FixedBoolean                *bool                                  `bson:"fixedBoolean,omitempty" json:"fixedBoolean,omitempty"`
	FixedCode                   string                                 `bson:"fixedCode,omitempty" json:"fixedCode,omitempty"`
	FixedCodeableConcept        *CodeableConcept                       `bson:"fixedCodeableConcept,omitempty" json:"fixedCodeableConcept,omitempty"`
	FixedCoding                 *Coding                                `bson:"fixedCoding,omitempty" json:"fixedCoding,omitempty"`
	FixedContactPoint           *ContactPoint                          `bson:"fixedContactPoint,omitempty" json:"fixedContactPoint,omitempty"`
	FixedDate                   *FHIRDateTime                          `bson:"fixedDate,omitempty" json:"fixedDate,omitempty"`
	FixedDateTime               *FHIRDateTime                          `bson:"fixedDateTime,omitempty" json:"fixedDateTime,omitempty"`
	FixedDecimal                *float64                               `bson:"fixedDecimal,omitempty" json:"fixedDecimal,omitempty"`
	FixedHumanName              *HumanName                             `bson:"fixedHumanName,omitempty" json:"fixedHumanName,omitempty"`
	FixedId                     string                                 `bson:"fixedId,omitempty" json:"fixedId,omitempty"`
	FixedIdentifier             *Identifier                            `bson:"fixedIdentifier,omitempty" json:"fixedIdentifier,omitempty"`
	FixedInstant                *FHIRDateTime                          `bson:"fixedInstant,omitempty" json:"fixedInstant,omitempty"`
	FixedInteger                *int32                                 `bson:"fixedInteger,omitempty" json:"fixedInteger,omitempty"`
	FixedMarkdown               string                                 `bson:"fixedMarkdown,omitempty" json:"fixedMarkdown,omitempty"`
	FixedMeta                   *Meta                                  `bson:"fixedMeta,omitempty" json:"fixedMeta,omitempty"`
	FixedOid                    string                                 `bson:"fixedOid,omitempty" json:"fixedOid,omitempty"`
	FixedPeriod                 *Period                                `bson:"fixedPeriod,omitempty" json:"fixedPeriod,omitempty"`
	FixedPositiveInt            *uint32                                `bson:"fixedPositiveInt,omitempty" json:"fixedPositiveInt,omitempty"`
	FixedQuantity               *Quantity                              `bson:"fixedQuantity,omitempty" json:"fixedQuantity,omitempty"`
	FixedRange                  *Range                                 `bson:"fixedRange,omitempty" json:"fixedRange,omitempty"`
	FixedRatio                  *Ratio                                 `bson:"fixedRatio,omitempty" json:"fixedRatio,omitempty"`
	FixedReference              *Reference                             `bson:"fixedReference,omitempty" json:"fixedReference,omitempty"`
	FixedSampledData            *SampledData                           `bson:"fixedSampledData,omitempty" json:"fixedSampledData,omitempty"`
	FixedSignature              *Signature                             `bson:"fixedSignature,omitempty" json:"fixedSignature,omitempty"`
	FixedString                 string                                 `bson:"fixedString,omitempty" json:"fixedString,omitempty"`
	FixedTime                   *FHIRDateTime                          `bson:"fixedTime,omitempty" json:"fixedTime,omitempty"`
	FixedTiming                 *Timing                                `bson:"fixedTiming,omitempty" json:"fixedTiming,omitempty"`
	FixedUnsignedInt            *uint32                                `bson:"fixedUnsignedInt,omitempty" json:"fixedUnsignedInt,omitempty"`
	FixedUri                    string                                 `bson:"fixedUri,omitempty" json:"fixedUri,omitempty"`
	PatternAddress              *Address                               `bson:"patternAddress,omitempty" json:"patternAddress,omitempty"`
	PatternAnnotation           *Annotation                            `bson:"patternAnnotation,omitempty" json:"patternAnnotation,omitempty"`
	PatternAttachment           *Attachment                            `bson:"patternAttachment,omitempty" json:"patternAttachment,omitempty"`
	PatternBase64Binary         string                                 `bson:"patternBase64Binary,omitempty" json:"patternBase64Binary,omitempty"`
	PatternBoolean              *bool                                  `bson:"patternBoolean,omitempty" json:"patternBoolean,omitempty"`
	PatternCode                 string                                 `bson:"patternCode,omitempty" json:"patternCode,omitempty"`
	PatternCodeableConcept      *CodeableConcept                       `bson:"patternCodeableConcept,omitempty" json:"patternCodeableConcept,omitempty"`
	PatternCoding               *Coding                                `bson:"patternCoding,omitempty" json:"patternCoding,omitempty"`
	PatternContactPoint         *ContactPoint                          `bson:"patternContactPoint,omitempty" json:"patternContactPoint,omitempty"`
	PatternDate                 *FHIRDateTime                          `bson:"patternDate,omitempty" json:"patternDate,omitempty"`
	PatternDateTime             *FHIRDateTime                          `bson:"patternDateTime,omitempty" json:"patternDateTime,omitempty"`
	PatternDecimal              *float64                               `bson:"patternDecimal,omitempty" json:"patternDecimal,omitempty"`
	PatternHumanName            *HumanName                             `bson:"patternHumanName,omitempty" json:"patternHumanName,omitempty"`
	PatternId                   string                                 `bson:"patternId,omitempty" json:"patternId,omitempty"`
	PatternIdentifier           *Identifier                            `bson:"patternIdentifier,omitempty" json:"patternIdentifier,omitempty"`
	PatternInstant              *FHIRDateTime                          `bson:"patternInstant,omitempty" json:"patternInstant,omitempty"`
	PatternInteger              *int32                                 `bson:"patternInteger,omitempty" json:"patternInteger,omitempty"`
	PatternMarkdown             string                                 `bson:"patternMarkdown,omitempty" json:"patternMarkdown,omitempty"`
	PatternMeta                 *Meta                                  `bson:"patternMeta,omitempty" json:"patternMeta,omitempty"`
	PatternOid                  string                                 `bson:"patternOid,omitempty" json:"patternOid,omitempty"`
	PatternPeriod               *Period                                `bson:"patternPeriod,omitempty" json:"patternPeriod,omitempty"`
	PatternPositiveInt          *uint32                                `bson:"patternPositiveInt,omitempty" json:"patternPositiveInt,omitempty"`
	PatternQuantity             *Quantity                              `bson:"patternQuantity,omitempty" json:"patternQuantity,omitempty"`
	PatternRange                *Range                                 `bson:"patternRange,omitempty" json:"patternRange,omitempty"`
	PatternRatio                *Ratio                                 `bson:"patternRatio,omitempty" json:"patternRatio,omitempty"`
	PatternReference            *Reference                             `bson:"patternReference,omitempty" json:"patternReference,omitempty"`
	PatternSampledData          *SampledData                           `bson:"patternSampledData,omitempty" json:"patternSampledData,omitempty"`
	PatternSignature            *Signature                             `bson:"patternSignature,omitempty" json:"patternSignature,omitempty"`
	PatternString               string                                 `bson:"patternString,omitempty" json:"patternString,omitempty"`
	PatternTime                 *FHIRDateTime                          `bson:"patternTime,omitempty" json:"patternTime,omitempty"`
	PatternTiming               *Timing                                `bson:"patternTiming,omitempty" json:"patternTiming,omitempty"`
	PatternUnsignedInt          *uint32                                `bson:"patternUnsignedInt,omitempty" json:"patternUnsignedInt,omitempty"`
	PatternUri                  string                                 `bson:"patternUri,omitempty" json:"patternUri,omitempty"`
	ExampleAddress              *Address                               `bson:"exampleAddress,omitempty" json:"exampleAddress,omitempty"`
	ExampleAnnotation           *Annotation                            `bson:"exampleAnnotation,omitempty" json:"exampleAnnotation,omitempty"`
	ExampleAttachment           *Attachment                            `bson:"exampleAttachment,omitempty" json:"exampleAttachment,omitempty"`
	ExampleBase64Binary         string                                 `bson:"exampleBase64Binary,omitempty" json:"exampleBase64Binary,omitempty"`
	ExampleBoolean              *bool                                  `bson:"exampleBoolean,omitempty" json:"exampleBoolean,omitempty"`
	ExampleCode                 string                                 `bson:"exampleCode,omitempty" json:"exampleCode,omitempty"`
	ExampleCodeableConcept      *CodeableConcept                       `bson:"exampleCodeableConcept,omitempty" json:"exampleCodeableConcept,omitempty"`
	ExampleCoding               *Coding                                `bson:"exampleCoding,omitempty" json:"exampleCoding,omitempty"`
	ExampleContactPoint         *ContactPoint                          `bson:"exampleContactPoint,omitempty" json:"exampleContactPoint,omitempty"`
	ExampleDate                 *FHIRDateTime                          `bson:"exampleDate,omitempty" json:"exampleDate,omitempty"`
	ExampleDateTime             *FHIRDateTime                          `bson:"exampleDateTime,omitempty" json:"exampleDateTime,omitempty"`
	ExampleDecimal              *float64                               `bson:"exampleDecimal,omitempty" json:"exampleDecimal,omitempty"`
	ExampleHumanName            *HumanName                             `bson:"exampleHumanName,omitempty" json:"exampleHumanName,omitempty"`
	ExampleId                   string                                 `bson:"exampleId,omitempty" json:"exampleId,omitempty"`
	ExampleIdentifier           *Identifier                            `bson:"exampleIdentifier,omitempty" json:"exampleIdentifier,omitempty"`
	ExampleInstant              *FHIRDateTime                          `bson:"exampleInstant,omitempty" json:"exampleInstant,omitempty"`
	ExampleInteger              *int32                                 `bson:"exampleInteger,omitempty" json:"exampleInteger,omitempty"`
	ExampleMarkdown             string                                 `bson:"exampleMarkdown,omitempty" json:"exampleMarkdown,omitempty"`
	ExampleMeta                 *Meta                                  `bson:"exampleMeta,omitempty" json:"exampleMeta,omitempty"`
	ExampleOid                  string                                 `bson:"exampleOid,omitempty" json:"exampleOid,omitempty"`
	ExamplePeriod               *Period                                `bson:"examplePeriod,omitempty" json:"examplePeriod,omitempty"`
	ExamplePositiveInt          *uint32                                `bson:"examplePositiveInt,omitempty" json:"examplePositiveInt,omitempty"`
	ExampleQuantity             *Quantity                              `bson:"exampleQuantity,omitempty" json:"exampleQuantity,omitempty"`
	ExampleRange                *Range                                 `bson:"exampleRange,omitempty" json:"exampleRange,omitempty"`
	ExampleRatio                *Ratio                                 `bson:"exampleRatio,omitempty" json:"exampleRatio,omitempty"`
	ExampleReference            *Reference                             `bson:"exampleReference,omitempty" json:"exampleReference,omitempty"`
	ExampleSampledData          *SampledData                           `bson:"exampleSampledData,omitempty" json:"exampleSampledData,omitempty"`
	ExampleSignature            *Signature                             `bson:"exampleSignature,omitempty" json:"exampleSignature,omitempty"`
	ExampleString               string                                 `bson:"exampleString,omitempty" json:"exampleString,omitempty"`
	ExampleTime                 *FHIRDateTime                          `bson:"exampleTime,omitempty" json:"exampleTime,omitempty"`
	ExampleTiming               *Timing                                `bson:"exampleTiming,omitempty" json:"exampleTiming,omitempty"`
	ExampleUnsignedInt          *uint32                                `bson:"exampleUnsignedInt,omitempty" json:"exampleUnsignedInt,omitempty"`
	ExampleUri                  string                                 `bson:"exampleUri,omitempty" json:"exampleUri,omitempty"`
	MinValueDate                *FHIRDateTime                          `bson:"minValueDate,omitempty" json:"minValueDate,omitempty"`
	MinValueDateTime            *FHIRDateTime                          `bson:"minValueDateTime,omitempty" json:"minValueDateTime,omitempty"`
	MinValueInstant             *FHIRDateTime                          `bson:"minValueInstant,omitempty" json:"minValueInstant,omitempty"`
	MinValueTime                *FHIRDateTime                          `bson:"minValueTime,omitempty" json:"minValueTime,omitempty"`
	MinValueDecimal             *float64                               `bson:"minValueDecimal,omitempty" json:"minValueDecimal,omitempty"`
	MinValueInteger             *int32                                 `bson:"minValueInteger,omitempty" json:"minValueInteger,omitempty"`
	MinValuePositiveInt         *uint32                                `bson:"minValuePositiveInt,omitempty" json:"minValuePositiveInt,omitempty"`
	MinValueUnsignedInt         *uint32                                `bson:"minValueUnsignedInt,omitempty" json:"minValueUnsignedInt,omitempty"`
	MinValueQuantity            *Quantity                              `bson:"minValueQuantity,omitempty" json:"minValueQuantity,omitempty"`
	MaxValueDate                *FHIRDateTime                          `bson:"maxValueDate,omitempty" json:"maxValueDate,omitempty"`
	MaxValueDateTime            *FHIRDateTime                          `bson:"maxValueDateTime,omitempty" json:"maxValueDateTime,omitempty"`
	MaxValueInstant             *FHIRDateTime                          `bson:"maxValueInstant,omitempty" json:"maxValueInstant,omitempty"`
	MaxValueTime                *FHIRDateTime                          `bson:"maxValueTime,omitempty" json:"maxValueTime,omitempty"`
	MaxValueDecimal             *float64                               `bson:"maxValueDecimal,omitempty" json:"maxValueDecimal,omitempty"`
	MaxValueInteger             *int32                                 `bson:"maxValueInteger,omitempty" json:"maxValueInteger,omitempty"`
	MaxValuePositiveInt         *uint32                                `bson:"maxValuePositiveInt,omitempty" json:"maxValuePositiveInt,omitempty"`
	MaxValueUnsignedInt         *uint32                                `bson:"maxValueUnsignedInt,omitempty" json:"maxValueUnsignedInt,omitempty"`
	MaxValueQuantity            *Quantity                              `bson:"maxValueQuantity,omitempty" json:"maxValueQuantity,omitempty"`
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
	BackboneElement `bson:",inline"`
	Discriminator   []string `bson:"discriminator,omitempty" json:"discriminator,omitempty"`
	Description     string   `bson:"description,omitempty" json:"description,omitempty"`
	Ordered         *bool    `bson:"ordered,omitempty" json:"ordered,omitempty"`
	Rules           string   `bson:"rules,omitempty" json:"rules,omitempty"`
}

type ElementDefinitionBaseComponent struct {
	BackboneElement `bson:",inline"`
	Path            string `bson:"path,omitempty" json:"path,omitempty"`
	Min             *int32 `bson:"min,omitempty" json:"min,omitempty"`
	Max             string `bson:"max,omitempty" json:"max,omitempty"`
}

type ElementDefinitionTypeRefComponent struct {
	BackboneElement `bson:",inline"`
	Code            string   `bson:"code,omitempty" json:"code,omitempty"`
	Profile         string   `bson:"profile,omitempty" json:"profile,omitempty"`
	TargetProfile   string   `bson:"targetProfile,omitempty" json:"targetProfile,omitempty"`
	Aggregation     []string `bson:"aggregation,omitempty" json:"aggregation,omitempty"`
	Versioning      string   `bson:"versioning,omitempty" json:"versioning,omitempty"`
}

type ElementDefinitionConstraintComponent struct {
	BackboneElement `bson:",inline"`
	Key             string `bson:"key,omitempty" json:"key,omitempty"`
	Requirements    string `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Severity        string `bson:"severity,omitempty" json:"severity,omitempty"`
	Human           string `bson:"human,omitempty" json:"human,omitempty"`
	Expression      string `bson:"expression,omitempty" json:"expression,omitempty"`
	Xpath           string `bson:"xpath,omitempty" json:"xpath,omitempty"`
}

type ElementDefinitionBindingComponent struct {
	BackboneElement   `bson:",inline"`
	Strength          string     `bson:"strength,omitempty" json:"strength,omitempty"`
	Description       string     `bson:"description,omitempty" json:"description,omitempty"`
	ValueSetUri       string     `bson:"valueSetUri,omitempty" json:"valueSetUri,omitempty"`
	ValueSetReference *Reference `bson:"valueSetReference,omitempty" json:"valueSetReference,omitempty"`
}

type ElementDefinitionMappingComponent struct {
	BackboneElement `bson:",inline"`
	Identity        string `bson:"identity,omitempty" json:"identity,omitempty"`
	Language        string `bson:"language,omitempty" json:"language,omitempty"`
	Map             string `bson:"map,omitempty" json:"map,omitempty"`
}
