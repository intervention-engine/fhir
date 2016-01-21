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

type DeviceMetric struct {
	DomainResource    `bson:",inline"`
	Type              *CodeableConcept                   `bson:"type,omitempty" json:"type,omitempty"`
	Identifier        *Identifier                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Unit              *CodeableConcept                   `bson:"unit,omitempty" json:"unit,omitempty"`
	Source            *Reference                         `bson:"source,omitempty" json:"source,omitempty"`
	Parent            *Reference                         `bson:"parent,omitempty" json:"parent,omitempty"`
	OperationalStatus string                             `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	Color             string                             `bson:"color,omitempty" json:"color,omitempty"`
	Category          string                             `bson:"category,omitempty" json:"category,omitempty"`
	MeasurementPeriod *Timing                            `bson:"measurementPeriod,omitempty" json:"measurementPeriod,omitempty"`
	Calibration       []DeviceMetricCalibrationComponent `bson:"calibration,omitempty" json:"calibration,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DeviceMetric) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		DeviceMetric
	}{
		ResourceType: "DeviceMetric",
		DeviceMetric: *resource,
	}
	return json.Marshal(x)
}

// The "deviceMetric" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type deviceMetric DeviceMetric

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DeviceMetric) UnmarshalJSON(data []byte) (err error) {
	x2 := deviceMetric{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DeviceMetric(x2)
	}
	return
}

type DeviceMetricCalibrationComponent struct {
	Type  string        `bson:"type,omitempty" json:"type,omitempty"`
	State string        `bson:"state,omitempty" json:"state,omitempty"`
	Time  *FHIRDateTime `bson:"time,omitempty" json:"time,omitempty"`
}

type DeviceMetricPlus struct {
	DeviceMetric             `bson:",inline"`
	DeviceMetricPlusIncludes `bson:",inline"`
}

type DeviceMetricPlusIncludes struct {
	IncludedParentResources *[]DeviceComponent `bson:"_includedParentResources,omitempty"`
	IncludedSourceResources *[]Device          `bson:"_includedSourceResources,omitempty"`
}

func (d *DeviceMetricPlusIncludes) GetIncludedParentResource() (deviceComponent *DeviceComponent, err error) {
	if d.IncludedParentResources == nil {
		err = errors.New("Included devicecomponents not requested")
	} else if len(*d.IncludedParentResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceComponent, but found %d", len(*d.IncludedParentResources))
	} else if len(*d.IncludedParentResources) == 1 {
		deviceComponent = &(*d.IncludedParentResources)[0]
	}
	return
}

func (d *DeviceMetricPlusIncludes) GetIncludedSourceResource() (device *Device, err error) {
	if d.IncludedSourceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedSourceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedSourceResources))
	} else if len(*d.IncludedSourceResources) == 1 {
		device = &(*d.IncludedSourceResources)[0]
	}
	return
}

func (d *DeviceMetricPlusIncludes) GetIncludedResources() map[string]interface{} {
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
