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

type ProcessResponse struct {
	DomainResource      `bson:",inline"`
	Identifier          []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Request             *Reference                      `bson:"request,omitempty" json:"request,omitempty"`
	Outcome             *Coding                         `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition         string                          `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Ruleset             *Coding                         `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset     *Coding                         `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created             *FHIRDateTime                   `bson:"created,omitempty" json:"created,omitempty"`
	Organization        *Reference                      `bson:"organization,omitempty" json:"organization,omitempty"`
	RequestProvider     *Reference                      `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization *Reference                      `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Form                *Coding                         `bson:"form,omitempty" json:"form,omitempty"`
	Notes               []ProcessResponseNotesComponent `bson:"notes,omitempty" json:"notes,omitempty"`
	Error               []Coding                        `bson:"error,omitempty" json:"error,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ProcessResponse) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		ProcessResponse
	}{
		ResourceType:    "ProcessResponse",
		ProcessResponse: *resource,
	}
	return json.Marshal(x)
}

// The "processResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type processResponse ProcessResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ProcessResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := processResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ProcessResponse(x2)
	}
	return
}

type ProcessResponseNotesComponent struct {
	Type *Coding `bson:"type,omitempty" json:"type,omitempty"`
	Text string  `bson:"text,omitempty" json:"text,omitempty"`
}

type ProcessResponsePlus struct {
	ProcessResponse             `bson:",inline"`
	ProcessResponsePlusIncludes `bson:",inline"`
}

type ProcessResponsePlusIncludes struct {
	IncludedOrganizationResources        *[]Organization `bson:"_includedOrganizationResources,omitempty"`
	IncludedRequestproviderResources     *[]Practitioner `bson:"_includedRequestproviderResources,omitempty"`
	IncludedRequestorganizationResources *[]Organization `bson:"_includedRequestorganizationResources,omitempty"`
}

func (p *ProcessResponsePlusIncludes) GetIncludedOrganizationResource() (organization *Organization, err error) {
	if p.IncludedOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResources))
	} else if len(*p.IncludedOrganizationResources) == 1 {
		organization = &(*p.IncludedOrganizationResources)[0]
	}
	return
}

func (p *ProcessResponsePlusIncludes) GetIncludedRequestproviderResource() (practitioner *Practitioner, err error) {
	if p.IncludedRequestproviderResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedRequestproviderResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedRequestproviderResources))
	} else if len(*p.IncludedRequestproviderResources) == 1 {
		practitioner = &(*p.IncludedRequestproviderResources)[0]
	}
	return
}

func (p *ProcessResponsePlusIncludes) GetIncludedRequestorganizationResource() (organization *Organization, err error) {
	if p.IncludedRequestorganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedRequestorganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedRequestorganizationResources))
	} else if len(*p.IncludedRequestorganizationResources) == 1 {
		organization = &(*p.IncludedRequestorganizationResources)[0]
	}
	return
}

func (p *ProcessResponsePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedOrganizationResources != nil {
		for _, r := range *p.IncludedOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedRequestproviderResources != nil {
		for _, r := range *p.IncludedRequestproviderResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedRequestorganizationResources != nil {
		for _, r := range *p.IncludedRequestorganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
