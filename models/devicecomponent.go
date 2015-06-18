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

import "time"

type DeviceComponent struct {
	Id                      string                                                           `json:"-" bson:"_id"`
	Type                    *CodeableConcept                                                 `bson:"type,omitempty" json:"type,omitempty"`
	Identifier              *Identifier                                                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	LastSystemChange        *FHIRDateTime                                                    `bson:"lastSystemChange,omitempty" json:"lastSystemChange,omitempty"`
	Source                  *Reference                                                       `bson:"source,omitempty" json:"source,omitempty"`
	Parent                  *Reference                                                       `bson:"parent,omitempty" json:"parent,omitempty"`
	OperationalStatus       []CodeableConcept                                                `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	ParameterGroup          *CodeableConcept                                                 `bson:"parameterGroup,omitempty" json:"parameterGroup,omitempty"`
	MeasurementPrinciple    string                                                           `bson:"measurementPrinciple,omitempty" json:"measurementPrinciple,omitempty"`
	ProductionSpecification []DeviceComponentDeviceComponentProductionSpecificationComponent `bson:"productionSpecification,omitempty" json:"productionSpecification,omitempty"`
	LanguageCode            *CodeableConcept                                                 `bson:"languageCode,omitempty" json:"languageCode,omitempty"`
}
type DeviceComponentDeviceComponentProductionSpecificationComponent struct {
	SpecType       *CodeableConcept `bson:"specType,omitempty" json:"specType,omitempty"`
	ComponentId    *Identifier      `bson:"componentId,omitempty" json:"componentId,omitempty"`
	ProductionSpec string           `bson:"productionSpec,omitempty" json:"productionSpec,omitempty"`
}

type DeviceComponentBundle struct {
	Type         string                       `json:"resourceType,omitempty"`
	Title        string                       `json:"title,omitempty"`
	Id           string                       `json:"id,omitempty"`
	Updated      time.Time                    `json:"updated,omitempty"`
	TotalResults int                          `json:"totalResults,omitempty"`
	Entry        []DeviceComponentBundleEntry `json:"entry,omitempty"`
	Category     DeviceComponentCategory      `json:"category,omitempty"`
}

type DeviceComponentBundleEntry struct {
	Title    string                  `json:"title,omitempty"`
	Id       string                  `json:"id,omitempty"`
	Content  DeviceComponent         `json:"content,omitempty"`
	Category DeviceComponentCategory `json:"category,omitempty"`
}

type DeviceComponentCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}
