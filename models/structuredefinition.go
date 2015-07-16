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

import "encoding/json"

type StructureDefinition struct {
	Id           string                                    `json:"id" bson:"_id"`
	Url          string                                    `bson:"url,omitempty" json:"url,omitempty"`
	Identifier   []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version      string                                    `bson:"version,omitempty" json:"version,omitempty"`
	Name         string                                    `bson:"name,omitempty" json:"name,omitempty"`
	UseContext   []CodeableConcept                         `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Display      string                                    `bson:"display,omitempty" json:"display,omitempty"`
	Publisher    string                                    `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact      []StructureDefinitionContactComponent     `bson:"contact,omitempty" json:"contact,omitempty"`
	Description  string                                    `bson:"description,omitempty" json:"description,omitempty"`
	Requirements string                                    `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright    string                                    `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Code         []Coding                                  `bson:"code,omitempty" json:"code,omitempty"`
	Status       string                                    `bson:"status,omitempty" json:"status,omitempty"`
	Experimental *bool                                     `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date         *FHIRDateTime                             `bson:"date,omitempty" json:"date,omitempty"`
	FhirVersion  string                                    `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Mapping      []StructureDefinitionMappingComponent     `bson:"mapping,omitempty" json:"mapping,omitempty"`
	Type         string                                    `bson:"type,omitempty" json:"type,omitempty"`
	Abstract     *bool                                     `bson:"abstract,omitempty" json:"abstract,omitempty"`
	ContextType  string                                    `bson:"contextType,omitempty" json:"contextType,omitempty"`
	Context      []string                                  `bson:"context,omitempty" json:"context,omitempty"`
	Base         string                                    `bson:"base,omitempty" json:"base,omitempty"`
	Snapshot     *StructureDefinitionSnapshotComponent     `bson:"snapshot,omitempty" json:"snapshot,omitempty"`
	Differential *StructureDefinitionDifferentialComponent `bson:"differential,omitempty" json:"differential,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *StructureDefinition) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		StructureDefinition
	}{
		ResourceType:        "StructureDefinition",
		StructureDefinition: *resource,
	}
	return json.Marshal(x)
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
