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
	"errors"
	"fmt"
)

type ConceptMap struct {
	DomainResource  `bson:",inline"`
	Url             string                             `bson:"url,omitempty" json:"url,omitempty"`
	Identifier      *Identifier                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version         string                             `bson:"version,omitempty" json:"version,omitempty"`
	Name            string                             `bson:"name,omitempty" json:"name,omitempty"`
	Status          string                             `bson:"status,omitempty" json:"status,omitempty"`
	Experimental    *bool                              `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher       string                             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact         []ConceptMapContactComponent       `bson:"contact,omitempty" json:"contact,omitempty"`
	Date            *FHIRDateTime                      `bson:"date,omitempty" json:"date,omitempty"`
	Description     string                             `bson:"description,omitempty" json:"description,omitempty"`
	UseContext      []CodeableConcept                  `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Requirements    string                             `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright       string                             `bson:"copyright,omitempty" json:"copyright,omitempty"`
	SourceUri       string                             `bson:"sourceUri,omitempty" json:"sourceUri,omitempty"`
	SourceReference *Reference                         `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	TargetUri       string                             `bson:"targetUri,omitempty" json:"targetUri,omitempty"`
	TargetReference *Reference                         `bson:"targetReference,omitempty" json:"targetReference,omitempty"`
	Element         []ConceptMapSourceElementComponent `bson:"element,omitempty" json:"element,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ConceptMap) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		ConceptMap
	}{
		ResourceType: "ConceptMap",
		ConceptMap:   *resource,
	}
	return json.Marshal(x)
}

// The "conceptMap" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type conceptMap ConceptMap

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ConceptMap) UnmarshalJSON(data []byte) (err error) {
	x2 := conceptMap{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ConceptMap(x2)
	}
	return
}

type ConceptMapContactComponent struct {
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type ConceptMapSourceElementComponent struct {
	CodeSystem string                             `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
	Code       string                             `bson:"code,omitempty" json:"code,omitempty"`
	Target     []ConceptMapTargetElementComponent `bson:"target,omitempty" json:"target,omitempty"`
}

type ConceptMapTargetElementComponent struct {
	CodeSystem  string                            `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
	Code        string                            `bson:"code,omitempty" json:"code,omitempty"`
	Equivalence string                            `bson:"equivalence,omitempty" json:"equivalence,omitempty"`
	Comments    string                            `bson:"comments,omitempty" json:"comments,omitempty"`
	DependsOn   []ConceptMapOtherElementComponent `bson:"dependsOn,omitempty" json:"dependsOn,omitempty"`
	Product     []ConceptMapOtherElementComponent `bson:"product,omitempty" json:"product,omitempty"`
}

type ConceptMapOtherElementComponent struct {
	Element    string `bson:"element,omitempty" json:"element,omitempty"`
	CodeSystem string `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
	Code       string `bson:"code,omitempty" json:"code,omitempty"`
}

type ConceptMapPlus struct {
	ConceptMap             `bson:",inline"`
	ConceptMapPlusIncludes `bson:",inline"`
}

type ConceptMapPlusIncludes struct {
	IncludedSourceStructureDefinitionResources    *[]StructureDefinition `bson:"_includedSourceStructureDefinitionResources,omitempty"`
	IncludedSourceValueSetResources               *[]ValueSet            `bson:"_includedSourceValueSetResources,omitempty"`
	IncludedTargetStructureDefinitionResources    *[]StructureDefinition `bson:"_includedTargetStructureDefinitionResources,omitempty"`
	IncludedTargetValueSetResources               *[]ValueSet            `bson:"_includedTargetValueSetResources,omitempty"`
	IncludedSourceuriStructureDefinitionResources *[]StructureDefinition `bson:"_includedSourceuriStructureDefinitionResources,omitempty"`
	IncludedSourceuriValueSetResources            *[]ValueSet            `bson:"_includedSourceuriValueSetResources,omitempty"`
}

func (c *ConceptMapPlusIncludes) GetIncludedSourceStructureDefinitionResource() (structureDefinition *StructureDefinition, err error) {
	if c.IncludedSourceStructureDefinitionResources == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*c.IncludedSourceStructureDefinitionResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*c.IncludedSourceStructureDefinitionResources))
	} else if len(*c.IncludedSourceStructureDefinitionResources) == 1 {
		structureDefinition = &(*c.IncludedSourceStructureDefinitionResources)[0]
	}
	return
}

func (c *ConceptMapPlusIncludes) GetIncludedSourceValueSetResource() (valueSet *ValueSet, err error) {
	if c.IncludedSourceValueSetResources == nil {
		err = errors.New("Included valuesets not requested")
	} else if len(*c.IncludedSourceValueSetResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 valueSet, but found %d", len(*c.IncludedSourceValueSetResources))
	} else if len(*c.IncludedSourceValueSetResources) == 1 {
		valueSet = &(*c.IncludedSourceValueSetResources)[0]
	}
	return
}

func (c *ConceptMapPlusIncludes) GetIncludedTargetStructureDefinitionResource() (structureDefinition *StructureDefinition, err error) {
	if c.IncludedTargetStructureDefinitionResources == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*c.IncludedTargetStructureDefinitionResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*c.IncludedTargetStructureDefinitionResources))
	} else if len(*c.IncludedTargetStructureDefinitionResources) == 1 {
		structureDefinition = &(*c.IncludedTargetStructureDefinitionResources)[0]
	}
	return
}

func (c *ConceptMapPlusIncludes) GetIncludedTargetValueSetResource() (valueSet *ValueSet, err error) {
	if c.IncludedTargetValueSetResources == nil {
		err = errors.New("Included valuesets not requested")
	} else if len(*c.IncludedTargetValueSetResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 valueSet, but found %d", len(*c.IncludedTargetValueSetResources))
	} else if len(*c.IncludedTargetValueSetResources) == 1 {
		valueSet = &(*c.IncludedTargetValueSetResources)[0]
	}
	return
}

func (c *ConceptMapPlusIncludes) GetIncludedSourceuriStructureDefinitionResource() (structureDefinition *StructureDefinition, err error) {
	if c.IncludedSourceuriStructureDefinitionResources == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*c.IncludedSourceuriStructureDefinitionResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*c.IncludedSourceuriStructureDefinitionResources))
	} else if len(*c.IncludedSourceuriStructureDefinitionResources) == 1 {
		structureDefinition = &(*c.IncludedSourceuriStructureDefinitionResources)[0]
	}
	return
}

func (c *ConceptMapPlusIncludes) GetIncludedSourceuriValueSetResource() (valueSet *ValueSet, err error) {
	if c.IncludedSourceuriValueSetResources == nil {
		err = errors.New("Included valuesets not requested")
	} else if len(*c.IncludedSourceuriValueSetResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 valueSet, but found %d", len(*c.IncludedSourceuriValueSetResources))
	} else if len(*c.IncludedSourceuriValueSetResources) == 1 {
		valueSet = &(*c.IncludedSourceuriValueSetResources)[0]
	}
	return
}

func (c *ConceptMapPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedSourceStructureDefinitionResources != nil {
		for _, r := range *c.IncludedSourceStructureDefinitionResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSourceValueSetResources != nil {
		for _, r := range *c.IncludedSourceValueSetResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedTargetStructureDefinitionResources != nil {
		for _, r := range *c.IncludedTargetStructureDefinitionResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedTargetValueSetResources != nil {
		for _, r := range *c.IncludedTargetValueSetResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSourceuriStructureDefinitionResources != nil {
		for _, r := range *c.IncludedSourceuriStructureDefinitionResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSourceuriValueSetResources != nil {
		for _, r := range *c.IncludedSourceuriValueSetResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
