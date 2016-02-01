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

type ProcessRequest struct {
	DomainResource  `bson:",inline"`
	Action          string                         `bson:"action,omitempty" json:"action,omitempty"`
	Identifier      []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ruleset         *Coding                        `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset *Coding                        `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created         *FHIRDateTime                  `bson:"created,omitempty" json:"created,omitempty"`
	Target          *Reference                     `bson:"target,omitempty" json:"target,omitempty"`
	Provider        *Reference                     `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization    *Reference                     `bson:"organization,omitempty" json:"organization,omitempty"`
	Request         *Reference                     `bson:"request,omitempty" json:"request,omitempty"`
	Response        *Reference                     `bson:"response,omitempty" json:"response,omitempty"`
	Nullify         *bool                          `bson:"nullify,omitempty" json:"nullify,omitempty"`
	Reference       string                         `bson:"reference,omitempty" json:"reference,omitempty"`
	Item            []ProcessRequestItemsComponent `bson:"item,omitempty" json:"item,omitempty"`
	Include         []string                       `bson:"include,omitempty" json:"include,omitempty"`
	Exclude         []string                       `bson:"exclude,omitempty" json:"exclude,omitempty"`
	Period          *Period                        `bson:"period,omitempty" json:"period,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ProcessRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ProcessRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ProcessRequest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ProcessRequest) GetBSON() (interface{}, error) {
	x.ResourceType = "ProcessRequest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "processRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type processRequest ProcessRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ProcessRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := processRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ProcessRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ProcessRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ProcessRequest"
	} else if x.ResourceType != "ProcessRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ProcessRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type ProcessRequestItemsComponent struct {
	SequenceLinkId *int32 `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
}

type ProcessRequestPlus struct {
	ProcessRequest             `bson:",inline"`
	ProcessRequestPlusIncludes `bson:",inline"`
}

type ProcessRequestPlusIncludes struct {
	IncludedProviderResources     *[]Practitioner `bson:"_includedProviderResources,omitempty"`
	IncludedOrganizationResources *[]Organization `bson:"_includedOrganizationResources,omitempty"`
}

func (p *ProcessRequestPlusIncludes) GetIncludedProviderResource() (practitioner *Practitioner, err error) {
	if p.IncludedProviderResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedProviderResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedProviderResources))
	} else if len(*p.IncludedProviderResources) == 1 {
		practitioner = &(*p.IncludedProviderResources)[0]
	}
	return
}

func (p *ProcessRequestPlusIncludes) GetIncludedOrganizationResource() (organization *Organization, err error) {
	if p.IncludedOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResources))
	} else if len(*p.IncludedOrganizationResources) == 1 {
		organization = &(*p.IncludedOrganizationResources)[0]
	}
	return
}

func (p *ProcessRequestPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedProviderResources != nil {
		for _, r := range *p.IncludedProviderResources {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedOrganizationResources != nil {
		for _, r := range *p.IncludedOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
