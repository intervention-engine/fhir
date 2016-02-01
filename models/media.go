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

type Media struct {
	DomainResource `bson:",inline"`
	Type           string           `bson:"type,omitempty" json:"type,omitempty"`
	Subtype        *CodeableConcept `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject        *Reference       `bson:"subject,omitempty" json:"subject,omitempty"`
	Operator       *Reference       `bson:"operator,omitempty" json:"operator,omitempty"`
	View           *CodeableConcept `bson:"view,omitempty" json:"view,omitempty"`
	DeviceName     string           `bson:"deviceName,omitempty" json:"deviceName,omitempty"`
	Height         *uint32          `bson:"height,omitempty" json:"height,omitempty"`
	Width          *uint32          `bson:"width,omitempty" json:"width,omitempty"`
	Frames         *uint32          `bson:"frames,omitempty" json:"frames,omitempty"`
	Duration       *uint32          `bson:"duration,omitempty" json:"duration,omitempty"`
	Content        *Attachment      `bson:"content,omitempty" json:"content,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Media) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Media"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Media), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Media) GetBSON() (interface{}, error) {
	x.ResourceType = "Media"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "media" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type media Media

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Media) UnmarshalJSON(data []byte) (err error) {
	x2 := media{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Media(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Media) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Media"
	} else if x.ResourceType != "Media" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Media, instead received %s", x.ResourceType))
	}
	return nil
}

type MediaPlus struct {
	Media             `bson:",inline"`
	MediaPlusIncludes `bson:",inline"`
}

type MediaPlusIncludes struct {
	IncludedSubjectPractitionerResources *[]Practitioner `bson:"_includedSubjectPractitionerResources,omitempty"`
	IncludedSubjectGroupResources        *[]Group        `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectSpecimenResources     *[]Specimen     `bson:"_includedSubjectSpecimenResources,omitempty"`
	IncludedSubjectDeviceResources       *[]Device       `bson:"_includedSubjectDeviceResources,omitempty"`
	IncludedSubjectPatientResources      *[]Patient      `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedPatientResources             *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedOperatorResources            *[]Practitioner `bson:"_includedOperatorResources,omitempty"`
}

func (m *MediaPlusIncludes) GetIncludedSubjectPractitionerResource() (practitioner *Practitioner, err error) {
	if m.IncludedSubjectPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedSubjectPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedSubjectPractitionerResources))
	} else if len(*m.IncludedSubjectPractitionerResources) == 1 {
		practitioner = &(*m.IncludedSubjectPractitionerResources)[0]
	}
	return
}

func (m *MediaPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if m.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*m.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*m.IncludedSubjectGroupResources))
	} else if len(*m.IncludedSubjectGroupResources) == 1 {
		group = &(*m.IncludedSubjectGroupResources)[0]
	}
	return
}

func (m *MediaPlusIncludes) GetIncludedSubjectSpecimenResource() (specimen *Specimen, err error) {
	if m.IncludedSubjectSpecimenResources == nil {
		err = errors.New("Included specimen not requested")
	} else if len(*m.IncludedSubjectSpecimenResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 specimen, but found %d", len(*m.IncludedSubjectSpecimenResources))
	} else if len(*m.IncludedSubjectSpecimenResources) == 1 {
		specimen = &(*m.IncludedSubjectSpecimenResources)[0]
	}
	return
}

func (m *MediaPlusIncludes) GetIncludedSubjectDeviceResource() (device *Device, err error) {
	if m.IncludedSubjectDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*m.IncludedSubjectDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*m.IncludedSubjectDeviceResources))
	} else if len(*m.IncludedSubjectDeviceResources) == 1 {
		device = &(*m.IncludedSubjectDeviceResources)[0]
	}
	return
}

func (m *MediaPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if m.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedSubjectPatientResources))
	} else if len(*m.IncludedSubjectPatientResources) == 1 {
		patient = &(*m.IncludedSubjectPatientResources)[0]
	}
	return
}

func (m *MediaPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if m.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResources))
	} else if len(*m.IncludedPatientResources) == 1 {
		patient = &(*m.IncludedPatientResources)[0]
	}
	return
}

func (m *MediaPlusIncludes) GetIncludedOperatorResource() (practitioner *Practitioner, err error) {
	if m.IncludedOperatorResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedOperatorResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedOperatorResources))
	} else if len(*m.IncludedOperatorResources) == 1 {
		practitioner = &(*m.IncludedOperatorResources)[0]
	}
	return
}

func (m *MediaPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedSubjectPractitionerResources != nil {
		for _, r := range *m.IncludedSubjectPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedSubjectGroupResources != nil {
		for _, r := range *m.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedSubjectSpecimenResources != nil {
		for _, r := range *m.IncludedSubjectSpecimenResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedSubjectDeviceResources != nil {
		for _, r := range *m.IncludedSubjectDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedSubjectPatientResources != nil {
		for _, r := range *m.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedPatientResources != nil {
		for _, r := range *m.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedOperatorResources != nil {
		for _, r := range *m.IncludedOperatorResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
