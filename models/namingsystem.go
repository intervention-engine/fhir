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

type NamingSystem struct {
	DomainResource `bson:",inline"`
	Name           string                          `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                          `bson:"status,omitempty" json:"status,omitempty"`
	Kind           string                          `bson:"kind,omitempty" json:"kind,omitempty"`
	Publisher      string                          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []NamingSystemContactComponent  `bson:"contact,omitempty" json:"contact,omitempty"`
	Responsible    string                          `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Date           *FHIRDateTime                   `bson:"date,omitempty" json:"date,omitempty"`
	Type           *CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Description    string                          `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []CodeableConcept               `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Usage          string                          `bson:"usage,omitempty" json:"usage,omitempty"`
	UniqueId       []NamingSystemUniqueIdComponent `bson:"uniqueId,omitempty" json:"uniqueId,omitempty"`
	ReplacedBy     *Reference                      `bson:"replacedBy,omitempty" json:"replacedBy,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *NamingSystem) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		NamingSystem
	}{
		ResourceType: "NamingSystem",
		NamingSystem: *resource,
	}
	return json.Marshal(x)
}

// The "namingSystem" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type namingSystem NamingSystem

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *NamingSystem) UnmarshalJSON(data []byte) (err error) {
	x2 := namingSystem{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = NamingSystem(x2)
	}
	return
}

type NamingSystemContactComponent struct {
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type NamingSystemUniqueIdComponent struct {
	Type      string  `bson:"type,omitempty" json:"type,omitempty"`
	Value     string  `bson:"value,omitempty" json:"value,omitempty"`
	Preferred *bool   `bson:"preferred,omitempty" json:"preferred,omitempty"`
	Period    *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type NamingSystemPlus struct {
	NamingSystem             `bson:",inline"`
	NamingSystemPlusIncludes `bson:",inline"`
}

type NamingSystemPlusIncludes struct {
	IncludedReplacedbyResources *[]NamingSystem `bson:"_includedReplacedbyResources,omitempty"`
}

func (n *NamingSystemPlusIncludes) GetIncludedReplacedbyResource() (namingSystem *NamingSystem, err error) {
	if n.IncludedReplacedbyResources == nil {
		err = errors.New("Included namingsystems not requested")
	} else if len(*n.IncludedReplacedbyResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 namingSystem, but found %d", len(*n.IncludedReplacedbyResources))
	} else if len(*n.IncludedReplacedbyResources) == 1 {
		namingSystem = &(*n.IncludedReplacedbyResources)[0]
	}
	return
}

func (n *NamingSystemPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.IncludedReplacedbyResources != nil {
		for _, r := range *n.IncludedReplacedbyResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
