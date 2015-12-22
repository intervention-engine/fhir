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

type ImplementationGuide struct {
	DomainResource `bson:",inline"`
	Url            string                                   `bson:"url,omitempty" json:"url,omitempty"`
	Version        string                                   `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                                   `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                                   `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                                    `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                                   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ImplementationGuideContactComponent    `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                            `bson:"date,omitempty" json:"date,omitempty"`
	Description    string                                   `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []CodeableConcept                        `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Copyright      string                                   `bson:"copyright,omitempty" json:"copyright,omitempty"`
	FhirVersion    string                                   `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Dependency     []ImplementationGuideDependencyComponent `bson:"dependency,omitempty" json:"dependency,omitempty"`
	Package        []ImplementationGuidePackageComponent    `bson:"package,omitempty" json:"package,omitempty"`
	Global         []ImplementationGuideGlobalComponent     `bson:"global,omitempty" json:"global,omitempty"`
	Binary         []string                                 `bson:"binary,omitempty" json:"binary,omitempty"`
	Page           *ImplementationGuidePageComponent        `bson:"page,omitempty" json:"page,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ImplementationGuide) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		ImplementationGuide
	}{
		ResourceType:        "ImplementationGuide",
		ImplementationGuide: *resource,
	}
	return json.Marshal(x)
}

// The "implementationGuide" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type implementationGuide ImplementationGuide

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ImplementationGuide) UnmarshalJSON(data []byte) (err error) {
	x2 := implementationGuide{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ImplementationGuide(x2)
	}
	return
}

type ImplementationGuideContactComponent struct {
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type ImplementationGuideDependencyComponent struct {
	Type string `bson:"type,omitempty" json:"type,omitempty"`
	Uri  string `bson:"uri,omitempty" json:"uri,omitempty"`
}

type ImplementationGuidePackageComponent struct {
	Name        string                                        `bson:"name,omitempty" json:"name,omitempty"`
	Description string                                        `bson:"description,omitempty" json:"description,omitempty"`
	Resource    []ImplementationGuidePackageResourceComponent `bson:"resource,omitempty" json:"resource,omitempty"`
}

type ImplementationGuidePackageResourceComponent struct {
	Purpose         string     `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Name            string     `bson:"name,omitempty" json:"name,omitempty"`
	Description     string     `bson:"description,omitempty" json:"description,omitempty"`
	Acronym         string     `bson:"acronym,omitempty" json:"acronym,omitempty"`
	SourceUri       string     `bson:"sourceUri,omitempty" json:"sourceUri,omitempty"`
	SourceReference *Reference `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	ExampleFor      *Reference `bson:"exampleFor,omitempty" json:"exampleFor,omitempty"`
}

type ImplementationGuideGlobalComponent struct {
	Type    string     `bson:"type,omitempty" json:"type,omitempty"`
	Profile *Reference `bson:"profile,omitempty" json:"profile,omitempty"`
}

type ImplementationGuidePageComponent struct {
	Source  string                             `bson:"source,omitempty" json:"source,omitempty"`
	Name    string                             `bson:"name,omitempty" json:"name,omitempty"`
	Kind    string                             `bson:"kind,omitempty" json:"kind,omitempty"`
	Type    []string                           `bson:"type,omitempty" json:"type,omitempty"`
	Package []string                           `bson:"package,omitempty" json:"package,omitempty"`
	Format  string                             `bson:"format,omitempty" json:"format,omitempty"`
	Page    []ImplementationGuidePageComponent `bson:"page,omitempty" json:"page,omitempty"`
}
