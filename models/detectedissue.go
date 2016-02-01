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

type DetectedIssue struct {
	DomainResource `bson:",inline"`
	Patient        *Reference                         `bson:"patient,omitempty" json:"patient,omitempty"`
	Category       *CodeableConcept                   `bson:"category,omitempty" json:"category,omitempty"`
	Severity       string                             `bson:"severity,omitempty" json:"severity,omitempty"`
	Implicated     []Reference                        `bson:"implicated,omitempty" json:"implicated,omitempty"`
	Detail         string                             `bson:"detail,omitempty" json:"detail,omitempty"`
	Date           *FHIRDateTime                      `bson:"date,omitempty" json:"date,omitempty"`
	Author         *Reference                         `bson:"author,omitempty" json:"author,omitempty"`
	Identifier     *Identifier                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Reference      string                             `bson:"reference,omitempty" json:"reference,omitempty"`
	Mitigation     []DetectedIssueMitigationComponent `bson:"mitigation,omitempty" json:"mitigation,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DetectedIssue) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DetectedIssue"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DetectedIssue), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DetectedIssue) GetBSON() (interface{}, error) {
	x.ResourceType = "DetectedIssue"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "detectedIssue" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type detectedIssue DetectedIssue

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DetectedIssue) UnmarshalJSON(data []byte) (err error) {
	x2 := detectedIssue{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DetectedIssue(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DetectedIssue) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DetectedIssue"
	} else if x.ResourceType != "DetectedIssue" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DetectedIssue, instead received %s", x.ResourceType))
	}
	return nil
}

type DetectedIssueMitigationComponent struct {
	Action *CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Date   *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Author *Reference       `bson:"author,omitempty" json:"author,omitempty"`
}

type DetectedIssuePlus struct {
	DetectedIssue             `bson:",inline"`
	DetectedIssuePlusIncludes `bson:",inline"`
}

type DetectedIssuePlusIncludes struct {
	IncludedPatientResources            *[]Patient      `bson:"_includedPatientResources,omitempty"`
	IncludedAuthorPractitionerResources *[]Practitioner `bson:"_includedAuthorPractitionerResources,omitempty"`
	IncludedAuthorDeviceResources       *[]Device       `bson:"_includedAuthorDeviceResources,omitempty"`
}

func (d *DetectedIssuePlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if d.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResources))
	} else if len(*d.IncludedPatientResources) == 1 {
		patient = &(*d.IncludedPatientResources)[0]
	}
	return
}

func (d *DetectedIssuePlusIncludes) GetIncludedAuthorPractitionerResource() (practitioner *Practitioner, err error) {
	if d.IncludedAuthorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedAuthorPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedAuthorPractitionerResources))
	} else if len(*d.IncludedAuthorPractitionerResources) == 1 {
		practitioner = &(*d.IncludedAuthorPractitionerResources)[0]
	}
	return
}

func (d *DetectedIssuePlusIncludes) GetIncludedAuthorDeviceResource() (device *Device, err error) {
	if d.IncludedAuthorDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedAuthorDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedAuthorDeviceResources))
	} else if len(*d.IncludedAuthorDeviceResources) == 1 {
		device = &(*d.IncludedAuthorDeviceResources)[0]
	}
	return
}

func (d *DetectedIssuePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPatientResources != nil {
		for _, r := range *d.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedAuthorPractitionerResources != nil {
		for _, r := range *d.IncludedAuthorPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedAuthorDeviceResources != nil {
		for _, r := range *d.IncludedAuthorDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
