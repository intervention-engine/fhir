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

type Location struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                     `bson:"status,omitempty" json:"status,omitempty"`
	Name                 string                     `bson:"name,omitempty" json:"name,omitempty"`
	Description          string                     `bson:"description,omitempty" json:"description,omitempty"`
	Mode                 string                     `bson:"mode,omitempty" json:"mode,omitempty"`
	Type                 *CodeableConcept           `bson:"type,omitempty" json:"type,omitempty"`
	Telecom              []ContactPoint             `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address              *Address                   `bson:"address,omitempty" json:"address,omitempty"`
	PhysicalType         *CodeableConcept           `bson:"physicalType,omitempty" json:"physicalType,omitempty"`
	Position             *LocationPositionComponent `bson:"position,omitempty" json:"position,omitempty"`
	ManagingOrganization *Reference                 `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	PartOf               *Reference                 `bson:"partOf,omitempty" json:"partOf,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Location) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Location"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Location), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Location) GetBSON() (interface{}, error) {
	x.ResourceType = "Location"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "location" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type location Location

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Location) UnmarshalJSON(data []byte) (err error) {
	x2 := location{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Location(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Location) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Location"
	} else if x.ResourceType != "Location" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Location, instead received %s", x.ResourceType))
	}
	return nil
}

type LocationPositionComponent struct {
	Longitude *float64 `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude  *float64 `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Altitude  *float64 `bson:"altitude,omitempty" json:"altitude,omitempty"`
}

type LocationPlus struct {
	Location             `bson:",inline"`
	LocationPlusIncludes `bson:",inline"`
}

type LocationPlusIncludes struct {
	IncludedPartofResources       *[]Location     `bson:"_includedPartofResources,omitempty"`
	IncludedOrganizationResources *[]Organization `bson:"_includedOrganizationResources,omitempty"`
}

func (l *LocationPlusIncludes) GetIncludedPartofResource() (location *Location, err error) {
	if l.IncludedPartofResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*l.IncludedPartofResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*l.IncludedPartofResources))
	} else if len(*l.IncludedPartofResources) == 1 {
		location = &(*l.IncludedPartofResources)[0]
	}
	return
}

func (l *LocationPlusIncludes) GetIncludedOrganizationResource() (organization *Organization, err error) {
	if l.IncludedOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*l.IncludedOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*l.IncludedOrganizationResources))
	} else if len(*l.IncludedOrganizationResources) == 1 {
		organization = &(*l.IncludedOrganizationResources)[0]
	}
	return
}

func (l *LocationPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedPartofResources != nil {
		for _, r := range *l.IncludedPartofResources {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedOrganizationResources != nil {
		for _, r := range *l.IncludedOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
