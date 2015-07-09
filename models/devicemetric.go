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

type DeviceMetric struct {
	Id                string                             `json:"-" bson:"_id"`
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

type DeviceMetricCalibrationComponent struct {
	Type  string        `bson:"type,omitempty" json:"type,omitempty"`
	State string        `bson:"state,omitempty" json:"state,omitempty"`
	Time  *FHIRDateTime `bson:"time,omitempty" json:"time,omitempty"`
}

type DeviceMetricBundle struct {
	Id    string                    `json:"id,omitempty"`
	Type  string                    `json:"resourceType,omitempty"`
	Base  string                    `json:"base,omitempty"`
	Total int                       `json:"total,omitempty"`
	Link  []BundleLinkComponent     `json:"link,omitempty"`
	Entry []DeviceMetricBundleEntry `json:"entry,omitempty"`
}

type DeviceMetricBundleEntry struct {
	Id       string                `json:"id,omitempty"`
	Base     string                `json:"base,omitempty"`
	Link     []BundleLinkComponent `json:"link,omitempty"`
	Resource DeviceMetric          `json:"resource,omitempty"`
}

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
