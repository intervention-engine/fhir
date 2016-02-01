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

type Schedule struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type            []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Actor           *Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
	PlanningHorizon *Period           `bson:"planningHorizon,omitempty" json:"planningHorizon,omitempty"`
	Comment         string            `bson:"comment,omitempty" json:"comment,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Schedule) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Schedule"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Schedule), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Schedule) GetBSON() (interface{}, error) {
	x.ResourceType = "Schedule"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "schedule" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type schedule Schedule

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Schedule) UnmarshalJSON(data []byte) (err error) {
	x2 := schedule{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Schedule(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Schedule) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Schedule"
	} else if x.ResourceType != "Schedule" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Schedule, instead received %s", x.ResourceType))
	}
	return nil
}

type SchedulePlus struct {
	Schedule             `bson:",inline"`
	SchedulePlusIncludes `bson:",inline"`
}

type SchedulePlusIncludes struct {
	IncludedActorPractitionerResources      *[]Practitioner      `bson:"_includedActorPractitionerResources,omitempty"`
	IncludedActorDeviceResources            *[]Device            `bson:"_includedActorDeviceResources,omitempty"`
	IncludedActorPatientResources           *[]Patient           `bson:"_includedActorPatientResources,omitempty"`
	IncludedActorHealthcareServiceResources *[]HealthcareService `bson:"_includedActorHealthcareServiceResources,omitempty"`
	IncludedActorRelatedPersonResources     *[]RelatedPerson     `bson:"_includedActorRelatedPersonResources,omitempty"`
	IncludedActorLocationResources          *[]Location          `bson:"_includedActorLocationResources,omitempty"`
}

func (s *SchedulePlusIncludes) GetIncludedActorPractitionerResource() (practitioner *Practitioner, err error) {
	if s.IncludedActorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*s.IncludedActorPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*s.IncludedActorPractitionerResources))
	} else if len(*s.IncludedActorPractitionerResources) == 1 {
		practitioner = &(*s.IncludedActorPractitionerResources)[0]
	}
	return
}

func (s *SchedulePlusIncludes) GetIncludedActorDeviceResource() (device *Device, err error) {
	if s.IncludedActorDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*s.IncludedActorDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*s.IncludedActorDeviceResources))
	} else if len(*s.IncludedActorDeviceResources) == 1 {
		device = &(*s.IncludedActorDeviceResources)[0]
	}
	return
}

func (s *SchedulePlusIncludes) GetIncludedActorPatientResource() (patient *Patient, err error) {
	if s.IncludedActorPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedActorPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedActorPatientResources))
	} else if len(*s.IncludedActorPatientResources) == 1 {
		patient = &(*s.IncludedActorPatientResources)[0]
	}
	return
}

func (s *SchedulePlusIncludes) GetIncludedActorHealthcareServiceResource() (healthcareService *HealthcareService, err error) {
	if s.IncludedActorHealthcareServiceResources == nil {
		err = errors.New("Included healthcareservices not requested")
	} else if len(*s.IncludedActorHealthcareServiceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 healthcareService, but found %d", len(*s.IncludedActorHealthcareServiceResources))
	} else if len(*s.IncludedActorHealthcareServiceResources) == 1 {
		healthcareService = &(*s.IncludedActorHealthcareServiceResources)[0]
	}
	return
}

func (s *SchedulePlusIncludes) GetIncludedActorRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if s.IncludedActorRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*s.IncludedActorRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*s.IncludedActorRelatedPersonResources))
	} else if len(*s.IncludedActorRelatedPersonResources) == 1 {
		relatedPerson = &(*s.IncludedActorRelatedPersonResources)[0]
	}
	return
}

func (s *SchedulePlusIncludes) GetIncludedActorLocationResource() (location *Location, err error) {
	if s.IncludedActorLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*s.IncludedActorLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*s.IncludedActorLocationResources))
	} else if len(*s.IncludedActorLocationResources) == 1 {
		location = &(*s.IncludedActorLocationResources)[0]
	}
	return
}

func (s *SchedulePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedActorPractitionerResources != nil {
		for _, r := range *s.IncludedActorPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedActorDeviceResources != nil {
		for _, r := range *s.IncludedActorDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedActorPatientResources != nil {
		for _, r := range *s.IncludedActorPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedActorHealthcareServiceResources != nil {
		for _, r := range *s.IncludedActorHealthcareServiceResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedActorRelatedPersonResources != nil {
		for _, r := range *s.IncludedActorRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedActorLocationResources != nil {
		for _, r := range *s.IncludedActorLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
