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
	x := struct {
		ResourceType string `json:"resourceType"`
		DeviceComponent
	}{
		ResourceType:    "DeviceComponent",
		DeviceComponent: *resource,
	}
	return json.Marshal(x)
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
	}
	return
}

type DeviceComponentProductionSpecificationComponent struct {
	SpecType       *CodeableConcept `bson:"specType,omitempty" json:"specType,omitempty"`
	ComponentId    *Identifier      `bson:"componentId,omitempty" json:"componentId,omitempty"`
	ProductionSpec string           `bson:"productionSpec,omitempty" json:"productionSpec,omitempty"`
}
