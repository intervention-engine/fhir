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

type Conformance struct {
	DomainResource `bson:",inline"`
	Url            string                              `bson:"url,omitempty" json:"url,omitempty"`
	Version        string                              `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                              `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                              `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                               `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                              `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ConformanceContactComponent       `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                       `bson:"date,omitempty" json:"date,omitempty"`
	Description    string                              `bson:"description,omitempty" json:"description,omitempty"`
	Requirements   string                              `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright      string                              `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Kind           string                              `bson:"kind,omitempty" json:"kind,omitempty"`
	Software       *ConformanceSoftwareComponent       `bson:"software,omitempty" json:"software,omitempty"`
	Implementation *ConformanceImplementationComponent `bson:"implementation,omitempty" json:"implementation,omitempty"`
	FhirVersion    string                              `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	AcceptUnknown  string                              `bson:"acceptUnknown,omitempty" json:"acceptUnknown,omitempty"`
	Format         []string                            `bson:"format,omitempty" json:"format,omitempty"`
	Profile        []Reference                         `bson:"profile,omitempty" json:"profile,omitempty"`
	Rest           []ConformanceRestComponent          `bson:"rest,omitempty" json:"rest,omitempty"`
	Messaging      []ConformanceMessagingComponent     `bson:"messaging,omitempty" json:"messaging,omitempty"`
	Document       []ConformanceDocumentComponent      `bson:"document,omitempty" json:"document,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Conformance) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Conformance"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Conformance), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Conformance) GetBSON() (interface{}, error) {
	x.ResourceType = "Conformance"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "conformance" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type conformance Conformance

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Conformance) UnmarshalJSON(data []byte) (err error) {
	x2 := conformance{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Conformance(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Conformance) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Conformance"
	} else if x.ResourceType != "Conformance" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Conformance, instead received %s", x.ResourceType))
	}
	return nil
}

type ConformanceContactComponent struct {
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type ConformanceSoftwareComponent struct {
	Name        string        `bson:"name,omitempty" json:"name,omitempty"`
	Version     string        `bson:"version,omitempty" json:"version,omitempty"`
	ReleaseDate *FHIRDateTime `bson:"releaseDate,omitempty" json:"releaseDate,omitempty"`
}

type ConformanceImplementationComponent struct {
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	Url         string `bson:"url,omitempty" json:"url,omitempty"`
}

type ConformanceRestComponent struct {
	Mode            string                                        `bson:"mode,omitempty" json:"mode,omitempty"`
	Documentation   string                                        `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Security        *ConformanceRestSecurityComponent             `bson:"security,omitempty" json:"security,omitempty"`
	Resource        []ConformanceRestResourceComponent            `bson:"resource,omitempty" json:"resource,omitempty"`
	Interaction     []ConformanceSystemInteractionComponent       `bson:"interaction,omitempty" json:"interaction,omitempty"`
	TransactionMode string                                        `bson:"transactionMode,omitempty" json:"transactionMode,omitempty"`
	SearchParam     []ConformanceRestResourceSearchParamComponent `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	Operation       []ConformanceRestOperationComponent           `bson:"operation,omitempty" json:"operation,omitempty"`
	Compartment     []string                                      `bson:"compartment,omitempty" json:"compartment,omitempty"`
}

type ConformanceRestSecurityComponent struct {
	Cors        *bool                                         `bson:"cors,omitempty" json:"cors,omitempty"`
	Service     []CodeableConcept                             `bson:"service,omitempty" json:"service,omitempty"`
	Description string                                        `bson:"description,omitempty" json:"description,omitempty"`
	Certificate []ConformanceRestSecurityCertificateComponent `bson:"certificate,omitempty" json:"certificate,omitempty"`
}

type ConformanceRestSecurityCertificateComponent struct {
	Type string `bson:"type,omitempty" json:"type,omitempty"`
	Blob string `bson:"blob,omitempty" json:"blob,omitempty"`
}

type ConformanceRestResourceComponent struct {
	Type              string                                        `bson:"type,omitempty" json:"type,omitempty"`
	Profile           *Reference                                    `bson:"profile,omitempty" json:"profile,omitempty"`
	Interaction       []ConformanceResourceInteractionComponent     `bson:"interaction,omitempty" json:"interaction,omitempty"`
	Versioning        string                                        `bson:"versioning,omitempty" json:"versioning,omitempty"`
	ReadHistory       *bool                                         `bson:"readHistory,omitempty" json:"readHistory,omitempty"`
	UpdateCreate      *bool                                         `bson:"updateCreate,omitempty" json:"updateCreate,omitempty"`
	ConditionalCreate *bool                                         `bson:"conditionalCreate,omitempty" json:"conditionalCreate,omitempty"`
	ConditionalUpdate *bool                                         `bson:"conditionalUpdate,omitempty" json:"conditionalUpdate,omitempty"`
	ConditionalDelete string                                        `bson:"conditionalDelete,omitempty" json:"conditionalDelete,omitempty"`
	SearchInclude     []string                                      `bson:"searchInclude,omitempty" json:"searchInclude,omitempty"`
	SearchRevInclude  []string                                      `bson:"searchRevInclude,omitempty" json:"searchRevInclude,omitempty"`
	SearchParam       []ConformanceRestResourceSearchParamComponent `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
}

type ConformanceResourceInteractionComponent struct {
	Code          string `bson:"code,omitempty" json:"code,omitempty"`
	Documentation string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type ConformanceRestResourceSearchParamComponent struct {
	Name          string   `bson:"name,omitempty" json:"name,omitempty"`
	Definition    string   `bson:"definition,omitempty" json:"definition,omitempty"`
	Type          string   `bson:"type,omitempty" json:"type,omitempty"`
	Documentation string   `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Target        []string `bson:"target,omitempty" json:"target,omitempty"`
	Modifier      []string `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Chain         []string `bson:"chain,omitempty" json:"chain,omitempty"`
}

type ConformanceSystemInteractionComponent struct {
	Code          string `bson:"code,omitempty" json:"code,omitempty"`
	Documentation string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type ConformanceRestOperationComponent struct {
	Name       string     `bson:"name,omitempty" json:"name,omitempty"`
	Definition *Reference `bson:"definition,omitempty" json:"definition,omitempty"`
}

type ConformanceMessagingComponent struct {
	Endpoint      []ConformanceMessagingEndpointComponent `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	ReliableCache *uint32                                 `bson:"reliableCache,omitempty" json:"reliableCache,omitempty"`
	Documentation string                                  `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Event         []ConformanceMessagingEventComponent    `bson:"event,omitempty" json:"event,omitempty"`
}

type ConformanceMessagingEndpointComponent struct {
	Protocol *Coding `bson:"protocol,omitempty" json:"protocol,omitempty"`
	Address  string  `bson:"address,omitempty" json:"address,omitempty"`
}

type ConformanceMessagingEventComponent struct {
	Code          *Coding    `bson:"code,omitempty" json:"code,omitempty"`
	Category      string     `bson:"category,omitempty" json:"category,omitempty"`
	Mode          string     `bson:"mode,omitempty" json:"mode,omitempty"`
	Focus         string     `bson:"focus,omitempty" json:"focus,omitempty"`
	Request       *Reference `bson:"request,omitempty" json:"request,omitempty"`
	Response      *Reference `bson:"response,omitempty" json:"response,omitempty"`
	Documentation string     `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type ConformanceDocumentComponent struct {
	Mode          string     `bson:"mode,omitempty" json:"mode,omitempty"`
	Documentation string     `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Profile       *Reference `bson:"profile,omitempty" json:"profile,omitempty"`
}

type ConformancePlus struct {
	Conformance             `bson:",inline"`
	ConformancePlusIncludes `bson:",inline"`
}

type ConformancePlusIncludes struct {
	IncludedProfileResources          *[]StructureDefinition `bson:"_includedProfileResources,omitempty"`
	IncludedSupportedprofileResources *[]StructureDefinition `bson:"_includedSupportedprofileResources,omitempty"`
}

func (c *ConformancePlusIncludes) GetIncludedProfileResource() (structureDefinition *StructureDefinition, err error) {
	if c.IncludedProfileResources == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*c.IncludedProfileResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*c.IncludedProfileResources))
	} else if len(*c.IncludedProfileResources) == 1 {
		structureDefinition = &(*c.IncludedProfileResources)[0]
	}
	return
}

func (c *ConformancePlusIncludes) GetIncludedSupportedprofileResources() (structureDefinitions []StructureDefinition, err error) {
	if c.IncludedSupportedprofileResources == nil {
		err = errors.New("Included structureDefinitions not requested")
	} else {
		structureDefinitions = *c.IncludedSupportedprofileResources
	}
	return
}

func (c *ConformancePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedProfileResources != nil {
		for _, r := range *c.IncludedProfileResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSupportedprofileResources != nil {
		for _, r := range *c.IncludedSupportedprofileResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
