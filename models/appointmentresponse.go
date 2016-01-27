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

type AppointmentResponse struct {
	DomainResource    `bson:",inline"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Appointment       *Reference        `bson:"appointment,omitempty" json:"appointment,omitempty"`
	Start             *FHIRDateTime     `bson:"start,omitempty" json:"start,omitempty"`
	End               *FHIRDateTime     `bson:"end,omitempty" json:"end,omitempty"`
	ParticipantType   []CodeableConcept `bson:"participantType,omitempty" json:"participantType,omitempty"`
	Actor             *Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
	ParticipantStatus string            `bson:"participantStatus,omitempty" json:"participantStatus,omitempty"`
	Comment           string            `bson:"comment,omitempty" json:"comment,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *AppointmentResponse) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		AppointmentResponse
	}{
		ResourceType:        "AppointmentResponse",
		AppointmentResponse: *resource,
	}
	return json.Marshal(x)
}

// The "appointmentResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type appointmentResponse AppointmentResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *AppointmentResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := appointmentResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = AppointmentResponse(x2)
	}
	return
}

type AppointmentResponsePlus struct {
	AppointmentResponse             `bson:",inline"`
	AppointmentResponsePlusIncludes `bson:",inline"`
}

type AppointmentResponsePlusIncludes struct {
	IncludedActorPractitionerResources      *[]Practitioner      `bson:"_includedActorPractitionerResources,omitempty"`
	IncludedActorDeviceResources            *[]Device            `bson:"_includedActorDeviceResources,omitempty"`
	IncludedActorPatientResources           *[]Patient           `bson:"_includedActorPatientResources,omitempty"`
	IncludedActorHealthcareServiceResources *[]HealthcareService `bson:"_includedActorHealthcareServiceResources,omitempty"`
	IncludedActorRelatedPersonResources     *[]RelatedPerson     `bson:"_includedActorRelatedPersonResources,omitempty"`
	IncludedActorLocationResources          *[]Location          `bson:"_includedActorLocationResources,omitempty"`
	IncludedPractitionerResources           *[]Practitioner      `bson:"_includedPractitionerResources,omitempty"`
	IncludedPatientResources                *[]Patient           `bson:"_includedPatientResources,omitempty"`
	IncludedAppointmentResources            *[]Appointment       `bson:"_includedAppointmentResources,omitempty"`
	IncludedLocationResources               *[]Location          `bson:"_includedLocationResources,omitempty"`
}

func (a *AppointmentResponsePlusIncludes) GetIncludedActorPractitionerResource() (practitioner *Practitioner, err error) {
	if a.IncludedActorPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedActorPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedActorPractitionerResources))
	} else if len(*a.IncludedActorPractitionerResources) == 1 {
		practitioner = &(*a.IncludedActorPractitionerResources)[0]
	}
	return
}

func (a *AppointmentResponsePlusIncludes) GetIncludedActorDeviceResource() (device *Device, err error) {
	if a.IncludedActorDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedActorDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedActorDeviceResources))
	} else if len(*a.IncludedActorDeviceResources) == 1 {
		device = &(*a.IncludedActorDeviceResources)[0]
	}
	return
}

func (a *AppointmentResponsePlusIncludes) GetIncludedActorPatientResource() (patient *Patient, err error) {
	if a.IncludedActorPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedActorPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedActorPatientResources))
	} else if len(*a.IncludedActorPatientResources) == 1 {
		patient = &(*a.IncludedActorPatientResources)[0]
	}
	return
}

func (a *AppointmentResponsePlusIncludes) GetIncludedActorHealthcareServiceResource() (healthcareService *HealthcareService, err error) {
	if a.IncludedActorHealthcareServiceResources == nil {
		err = errors.New("Included healthcareservices not requested")
	} else if len(*a.IncludedActorHealthcareServiceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 healthcareService, but found %d", len(*a.IncludedActorHealthcareServiceResources))
	} else if len(*a.IncludedActorHealthcareServiceResources) == 1 {
		healthcareService = &(*a.IncludedActorHealthcareServiceResources)[0]
	}
	return
}

func (a *AppointmentResponsePlusIncludes) GetIncludedActorRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedActorRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedActorRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedActorRelatedPersonResources))
	} else if len(*a.IncludedActorRelatedPersonResources) == 1 {
		relatedPerson = &(*a.IncludedActorRelatedPersonResources)[0]
	}
	return
}

func (a *AppointmentResponsePlusIncludes) GetIncludedActorLocationResource() (location *Location, err error) {
	if a.IncludedActorLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*a.IncludedActorLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*a.IncludedActorLocationResources))
	} else if len(*a.IncludedActorLocationResources) == 1 {
		location = &(*a.IncludedActorLocationResources)[0]
	}
	return
}

func (a *AppointmentResponsePlusIncludes) GetIncludedPractitionerResource() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResources))
	} else if len(*a.IncludedPractitionerResources) == 1 {
		practitioner = &(*a.IncludedPractitionerResources)[0]
	}
	return
}

func (a *AppointmentResponsePlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if a.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResources))
	} else if len(*a.IncludedPatientResources) == 1 {
		patient = &(*a.IncludedPatientResources)[0]
	}
	return
}

func (a *AppointmentResponsePlusIncludes) GetIncludedAppointmentResource() (appointment *Appointment, err error) {
	if a.IncludedAppointmentResources == nil {
		err = errors.New("Included appointments not requested")
	} else if len(*a.IncludedAppointmentResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 appointment, but found %d", len(*a.IncludedAppointmentResources))
	} else if len(*a.IncludedAppointmentResources) == 1 {
		appointment = &(*a.IncludedAppointmentResources)[0]
	}
	return
}

func (a *AppointmentResponsePlusIncludes) GetIncludedLocationResource() (location *Location, err error) {
	if a.IncludedLocationResources == nil {
		err = errors.New("Included locations not requested")
	} else if len(*a.IncludedLocationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*a.IncludedLocationResources))
	} else if len(*a.IncludedLocationResources) == 1 {
		location = &(*a.IncludedLocationResources)[0]
	}
	return
}

func (a *AppointmentResponsePlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedActorPractitionerResources != nil {
		for _, r := range *a.IncludedActorPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedActorDeviceResources != nil {
		for _, r := range *a.IncludedActorDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedActorPatientResources != nil {
		for _, r := range *a.IncludedActorPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedActorHealthcareServiceResources != nil {
		for _, r := range *a.IncludedActorHealthcareServiceResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedActorRelatedPersonResources != nil {
		for _, r := range *a.IncludedActorRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedActorLocationResources != nil {
		for _, r := range *a.IncludedActorLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPractitionerResources != nil {
		for _, r := range *a.IncludedPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPatientResources != nil {
		for _, r := range *a.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedAppointmentResources != nil {
		for _, r := range *a.IncludedAppointmentResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedLocationResources != nil {
		for _, r := range *a.IncludedLocationResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
