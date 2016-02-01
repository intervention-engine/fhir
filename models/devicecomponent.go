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

type DeviceComponent struct {
	DomainResource          `bson:",inline"`
	Type                    *CodeableConcept                                  `bson:"type,omitempty" json:"type,omitempty"`
	Identifier              *Identifier                                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	LastSystemChange        *FHIRDateTime                                     `bson:"lastSystemChange,omitempty" json:"lastSystemChange,omitempty"`
	Source                  *Reference                                        `bson:"source,omitempty" json:"source,omitempty"`
	Parent                  *Reference                                        `bson:"parent,omitempty" json:"parent,omitempty"`
	OperationalStatus       []CodeableConcept                                 `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	ParameterGroup          *CodeableConcept                                  `bson:"parameterGroup,omitempty" json:"parameterGroup,omitempty"`
	MeasurementPrinciple    string                                            `bson:"measurementPrinciple,omitempty" json:"measurementPrinciple,omitempty"`
	ProductionSpecification []DeviceComponentProductionSpecificationComponent `bson:"productionSpecification,omitempty" json:"productionSpecification,omitempty"`
	LanguageCode            *CodeableConcept                                  `bson:"languageCode,omitempty" json:"languageCode,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DeviceComponent) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DeviceComponent"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DeviceComponent), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DeviceComponent) GetBSON() (interface{}, error) {
	x.ResourceType = "DeviceComponent"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "deviceComponent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type deviceComponent DeviceComponent

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DeviceComponent) UnmarshalJSON(data []byte) (err error) {
	x2 := deviceComponent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DeviceComponent(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DeviceComponent) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DeviceComponent"
	} else if x.ResourceType != "DeviceComponent" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DeviceComponent, instead received %s", x.ResourceType))
	}
	return nil
}

type DeviceComponentProductionSpecificationComponent struct {
	SpecType       *CodeableConcept `bson:"specType,omitempty" json:"specType,omitempty"`
	ComponentId    *Identifier      `bson:"componentId,omitempty" json:"componentId,omitempty"`
	ProductionSpec string           `bson:"productionSpec,omitempty" json:"productionSpec,omitempty"`
}

type DeviceComponentPlus struct {
	DeviceComponent             `bson:",inline"`
	DeviceComponentPlusIncludes `bson:",inline"`
}

type DeviceComponentPlusIncludes struct {
	IncludedParentResources *[]DeviceComponent `bson:"_includedParentResources,omitempty"`
	IncludedSourceResources *[]Device          `bson:"_includedSourceResources,omitempty"`
}

func (d *DeviceComponentPlusIncludes) GetIncludedParentResource() (deviceComponent *DeviceComponent, err error) {
	if d.IncludedParentResources == nil {
		err = errors.New("Included devicecomponents not requested")
	} else if len(*d.IncludedParentResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceComponent, but found %d", len(*d.IncludedParentResources))
	} else if len(*d.IncludedParentResources) == 1 {
		deviceComponent = &(*d.IncludedParentResources)[0]
	}
	return
}

func (d *DeviceComponentPlusIncludes) GetIncludedSourceResource() (device *Device, err error) {
	if d.IncludedSourceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedSourceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedSourceResources))
	} else if len(*d.IncludedSourceResources) == 1 {
		device = &(*d.IncludedSourceResources)[0]
	}
	return
}

func (d *DeviceComponentPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedParentResources != nil {
		for _, r := range *d.IncludedParentResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSourceResources != nil {
		for _, r := range *d.IncludedSourceResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
