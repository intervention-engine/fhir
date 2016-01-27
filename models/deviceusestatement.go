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

type DeviceUseStatement struct {
	DomainResource          `bson:",inline"`
	BodySiteCodeableConcept *CodeableConcept  `bson:"bodySiteCodeableConcept,omitempty" json:"bodySiteCodeableConcept,omitempty"`
	BodySiteReference       *Reference        `bson:"bodySiteReference,omitempty" json:"bodySiteReference,omitempty"`
	WhenUsed                *Period           `bson:"whenUsed,omitempty" json:"whenUsed,omitempty"`
	Device                  *Reference        `bson:"device,omitempty" json:"device,omitempty"`
	Identifier              []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Indication              []CodeableConcept `bson:"indication,omitempty" json:"indication,omitempty"`
	Notes                   []string          `bson:"notes,omitempty" json:"notes,omitempty"`
	RecordedOn              *FHIRDateTime     `bson:"recordedOn,omitempty" json:"recordedOn,omitempty"`
	Subject                 *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	TimingTiming            *Timing           `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingPeriod            *Period           `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDateTime          *FHIRDateTime     `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DeviceUseStatement) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		DeviceUseStatement
	}{
		ResourceType:       "DeviceUseStatement",
		DeviceUseStatement: *resource,
	}
	return json.Marshal(x)
}

// The "deviceUseStatement" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type deviceUseStatement DeviceUseStatement

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DeviceUseStatement) UnmarshalJSON(data []byte) (err error) {
	x2 := deviceUseStatement{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DeviceUseStatement(x2)
	}
	return
}

type DeviceUseStatementPlus struct {
	DeviceUseStatement             `bson:",inline"`
	DeviceUseStatementPlusIncludes `bson:",inline"`
}

type DeviceUseStatementPlusIncludes struct {
	IncludedSubjectResources *[]Patient `bson:"_includedSubjectResources,omitempty"`
	IncludedPatientResources *[]Patient `bson:"_includedPatientResources,omitempty"`
	IncludedDeviceResources  *[]Device  `bson:"_includedDeviceResources,omitempty"`
}

func (d *DeviceUseStatementPlusIncludes) GetIncludedSubjectResource() (patient *Patient, err error) {
	if d.IncludedSubjectResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedSubjectResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedSubjectResources))
	} else if len(*d.IncludedSubjectResources) == 1 {
		patient = &(*d.IncludedSubjectResources)[0]
	}
	return
}

func (d *DeviceUseStatementPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if d.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResources))
	} else if len(*d.IncludedPatientResources) == 1 {
		patient = &(*d.IncludedPatientResources)[0]
	}
	return
}

func (d *DeviceUseStatementPlusIncludes) GetIncludedDeviceResource() (device *Device, err error) {
	if d.IncludedDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResources))
	} else if len(*d.IncludedDeviceResources) == 1 {
		device = &(*d.IncludedDeviceResources)[0]
	}
	return
}

func (d *DeviceUseStatementPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedSubjectResources != nil {
		for _, r := range *d.IncludedSubjectResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResources != nil {
		for _, r := range *d.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResources != nil {
		for _, r := range *d.IncludedDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
