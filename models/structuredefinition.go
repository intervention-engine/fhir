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

type StructureDefinition struct {
	DomainResource  `bson:",inline"`
	Url             string                                    `bson:"url,omitempty" json:"url,omitempty"`
	Identifier      []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version         string                                    `bson:"version,omitempty" json:"version,omitempty"`
	Name            string                                    `bson:"name,omitempty" json:"name,omitempty"`
	Display         string                                    `bson:"display,omitempty" json:"display,omitempty"`
	Status          string                                    `bson:"status,omitempty" json:"status,omitempty"`
	Experimental    *bool                                     `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher       string                                    `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact         []StructureDefinitionContactComponent     `bson:"contact,omitempty" json:"contact,omitempty"`
	Date            *FHIRDateTime                             `bson:"date,omitempty" json:"date,omitempty"`
	Description     string                                    `bson:"description,omitempty" json:"description,omitempty"`
	UseContext      []CodeableConcept                         `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Requirements    string                                    `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright       string                                    `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Code            []Coding                                  `bson:"code,omitempty" json:"code,omitempty"`
	FhirVersion     string                                    `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Mapping         []StructureDefinitionMappingComponent     `bson:"mapping,omitempty" json:"mapping,omitempty"`
	Kind            string                                    `bson:"kind,omitempty" json:"kind,omitempty"`
	ConstrainedType string                                    `bson:"constrainedType,omitempty" json:"constrainedType,omitempty"`
	Abstract        *bool                                     `bson:"abstract,omitempty" json:"abstract,omitempty"`
	ContextType     string                                    `bson:"contextType,omitempty" json:"contextType,omitempty"`
	Context         []string                                  `bson:"context,omitempty" json:"context,omitempty"`
	Base            string                                    `bson:"base,omitempty" json:"base,omitempty"`
	Snapshot        *StructureDefinitionSnapshotComponent     `bson:"snapshot,omitempty" json:"snapshot,omitempty"`
	Differential    *StructureDefinitionDifferentialComponent `bson:"differential,omitempty" json:"differential,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *StructureDefinition) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "StructureDefinition"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to StructureDefinition), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *StructureDefinition) GetBSON() (interface{}, error) {
	x.ResourceType = "StructureDefinition"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "structureDefinition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type structureDefinition StructureDefinition

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *StructureDefinition) UnmarshalJSON(data []byte) (err error) {
	x2 := structureDefinition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = StructureDefinition(x2)
		return x.checkResourceType()
	}
	return
}

func (x *StructureDefinition) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "StructureDefinition"
	} else if x.ResourceType != "StructureDefinition" {
		return errors.New(fmt.Sprintf("Expected resourceType to be StructureDefinition, instead received %s", x.ResourceType))
	}
	return nil
}

type StructureDefinitionContactComponent struct {
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type StructureDefinitionMappingComponent struct {
	Identity string `bson:"identity,omitempty" json:"identity,omitempty"`
	Uri      string `bson:"uri,omitempty" json:"uri,omitempty"`
	Name     string `bson:"name,omitempty" json:"name,omitempty"`
	Comments string `bson:"comments,omitempty" json:"comments,omitempty"`
}

type StructureDefinitionSnapshotComponent struct {
	Element []ElementDefinition `bson:"element,omitempty" json:"element,omitempty"`
}

type StructureDefinitionDifferentialComponent struct {
	Element []ElementDefinition `bson:"element,omitempty" json:"element,omitempty"`
}

type StructureDefinitionPlus struct {
	StructureDefinition             `bson:",inline"`
	StructureDefinitionPlusIncludes `bson:",inline"`
}

type StructureDefinitionPlusIncludes struct {
	IncludedValuesetResources *[]ValueSet `bson:"_includedValuesetResources,omitempty"`
}

func (s *StructureDefinitionPlusIncludes) GetIncludedValuesetResource() (valueSet *ValueSet, err error) {
	if s.IncludedValuesetResources == nil {
		err = errors.New("Included valuesets not requested")
	} else if len(*s.IncludedValuesetResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 valueSet, but found %d", len(*s.IncludedValuesetResources))
	} else if len(*s.IncludedValuesetResources) == 1 {
		valueSet = &(*s.IncludedValuesetResources)[0]
	}
	return
}

func (s *StructureDefinitionPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedValuesetResources != nil {
		for _, r := range *s.IncludedValuesetResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
