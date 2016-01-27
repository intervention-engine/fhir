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

type OperationDefinition struct {
	DomainResource `bson:",inline"`
	Url            string                                  `bson:"url,omitempty" json:"url,omitempty"`
	Version        string                                  `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                                  `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                                  `bson:"status,omitempty" json:"status,omitempty"`
	Kind           string                                  `bson:"kind,omitempty" json:"kind,omitempty"`
	Experimental   *bool                                   `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                                  `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []OperationDefinitionContactComponent   `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                           `bson:"date,omitempty" json:"date,omitempty"`
	Description    string                                  `bson:"description,omitempty" json:"description,omitempty"`
	Requirements   string                                  `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Idempotent     *bool                                   `bson:"idempotent,omitempty" json:"idempotent,omitempty"`
	Code           string                                  `bson:"code,omitempty" json:"code,omitempty"`
	Notes          string                                  `bson:"notes,omitempty" json:"notes,omitempty"`
	Base           *Reference                              `bson:"base,omitempty" json:"base,omitempty"`
	System         *bool                                   `bson:"system,omitempty" json:"system,omitempty"`
	Type           []string                                `bson:"type,omitempty" json:"type,omitempty"`
	Instance       *bool                                   `bson:"instance,omitempty" json:"instance,omitempty"`
	Parameter      []OperationDefinitionParameterComponent `bson:"parameter,omitempty" json:"parameter,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *OperationDefinition) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		OperationDefinition
	}{
		ResourceType:        "OperationDefinition",
		OperationDefinition: *resource,
	}
	return json.Marshal(x)
}

// The "operationDefinition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type operationDefinition OperationDefinition

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *OperationDefinition) UnmarshalJSON(data []byte) (err error) {
	x2 := operationDefinition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = OperationDefinition(x2)
	}
	return
}

type OperationDefinitionContactComponent struct {
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type OperationDefinitionParameterComponent struct {
	Name          string                                        `bson:"name,omitempty" json:"name,omitempty"`
	Use           string                                        `bson:"use,omitempty" json:"use,omitempty"`
	Min           *int32                                        `bson:"min,omitempty" json:"min,omitempty"`
	Max           string                                        `bson:"max,omitempty" json:"max,omitempty"`
	Documentation string                                        `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Type          string                                        `bson:"type,omitempty" json:"type,omitempty"`
	Profile       *Reference                                    `bson:"profile,omitempty" json:"profile,omitempty"`
	Binding       *OperationDefinitionParameterBindingComponent `bson:"binding,omitempty" json:"binding,omitempty"`
	Part          []OperationDefinitionParameterComponent       `bson:"part,omitempty" json:"part,omitempty"`
}

type OperationDefinitionParameterBindingComponent struct {
	Strength          string     `bson:"strength,omitempty" json:"strength,omitempty"`
	ValueSetUri       string     `bson:"valueSetUri,omitempty" json:"valueSetUri,omitempty"`
	ValueSetReference *Reference `bson:"valueSetReference,omitempty" json:"valueSetReference,omitempty"`
}

type OperationDefinitionPlus struct {
	OperationDefinition             `bson:",inline"`
	OperationDefinitionPlusIncludes `bson:",inline"`
}

type OperationDefinitionPlusIncludes struct {
	IncludedProfileResources *[]StructureDefinition `bson:"_includedProfileResources,omitempty"`
	IncludedBaseResources    *[]OperationDefinition `bson:"_includedBaseResources,omitempty"`
}

func (o *OperationDefinitionPlusIncludes) GetIncludedProfileResource() (structureDefinition *StructureDefinition, err error) {
	if o.IncludedProfileResources == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*o.IncludedProfileResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*o.IncludedProfileResources))
	} else if len(*o.IncludedProfileResources) == 1 {
		structureDefinition = &(*o.IncludedProfileResources)[0]
	}
	return
}

func (o *OperationDefinitionPlusIncludes) GetIncludedBaseResource() (operationDefinition *OperationDefinition, err error) {
	if o.IncludedBaseResources == nil {
		err = errors.New("Included operationdefinitions not requested")
	} else if len(*o.IncludedBaseResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 operationDefinition, but found %d", len(*o.IncludedBaseResources))
	} else if len(*o.IncludedBaseResources) == 1 {
		operationDefinition = &(*o.IncludedBaseResources)[0]
	}
	return
}

func (o *OperationDefinitionPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedProfileResources != nil {
		for _, r := range *o.IncludedProfileResources {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedBaseResources != nil {
		for _, r := range *o.IncludedBaseResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
