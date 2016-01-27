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

type AuditEvent struct {
	DomainResource `bson:",inline"`
	Event          *AuditEventEventComponent        `bson:"event,omitempty" json:"event,omitempty"`
	Participant    []AuditEventParticipantComponent `bson:"participant,omitempty" json:"participant,omitempty"`
	Source         *AuditEventSourceComponent       `bson:"source,omitempty" json:"source,omitempty"`
	Object         []AuditEventObjectComponent      `bson:"object,omitempty" json:"object,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *AuditEvent) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		AuditEvent
	}{
		ResourceType: "AuditEvent",
		AuditEvent:   *resource,
	}
	return json.Marshal(x)
}

// The "auditEvent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type auditEvent AuditEvent

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *AuditEvent) UnmarshalJSON(data []byte) (err error) {
	x2 := auditEvent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = AuditEvent(x2)
	}
	return
}

type AuditEventEventComponent struct {
	Type           *Coding       `bson:"type,omitempty" json:"type,omitempty"`
	Subtype        []Coding      `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Action         string        `bson:"action,omitempty" json:"action,omitempty"`
	DateTime       *FHIRDateTime `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Outcome        string        `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OutcomeDesc    string        `bson:"outcomeDesc,omitempty" json:"outcomeDesc,omitempty"`
	PurposeOfEvent []Coding      `bson:"purposeOfEvent,omitempty" json:"purposeOfEvent,omitempty"`
}

type AuditEventParticipantComponent struct {
	Role         []CodeableConcept                      `bson:"role,omitempty" json:"role,omitempty"`
	Reference    *Reference                             `bson:"reference,omitempty" json:"reference,omitempty"`
	UserId       *Identifier                            `bson:"userId,omitempty" json:"userId,omitempty"`
	AltId        string                                 `bson:"altId,omitempty" json:"altId,omitempty"`
	Name         string                                 `bson:"name,omitempty" json:"name,omitempty"`
	Requestor    *bool                                  `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Location     *Reference                             `bson:"location,omitempty" json:"location,omitempty"`
	Policy       []string                               `bson:"policy,omitempty" json:"policy,omitempty"`
	Media        *Coding                                `bson:"media,omitempty" json:"media,omitempty"`
	Network      *AuditEventParticipantNetworkComponent `bson:"network,omitempty" json:"network,omitempty"`
	PurposeOfUse []Coding                               `bson:"purposeOfUse,omitempty" json:"purposeOfUse,omitempty"`
}

type AuditEventParticipantNetworkComponent struct {
	Address string `bson:"address,omitempty" json:"address,omitempty"`
	Type    string `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventSourceComponent struct {
	Site       string      `bson:"site,omitempty" json:"site,omitempty"`
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type       []Coding    `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventObjectComponent struct {
	Identifier    *Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Reference     *Reference                        `bson:"reference,omitempty" json:"reference,omitempty"`
	Type          *Coding                           `bson:"type,omitempty" json:"type,omitempty"`
	Role          *Coding                           `bson:"role,omitempty" json:"role,omitempty"`
	Lifecycle     *Coding                           `bson:"lifecycle,omitempty" json:"lifecycle,omitempty"`
	SecurityLabel []Coding                          `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Name          string                            `bson:"name,omitempty" json:"name,omitempty"`
	Description   string                            `bson:"description,omitempty" json:"description,omitempty"`
	Query         string                            `bson:"query,omitempty" json:"query,omitempty"`
	Detail        []AuditEventObjectDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type AuditEventObjectDetailComponent struct {
	Type  string `bson:"type,omitempty" json:"type,omitempty"`
	Value string `bson:"value,omitempty" json:"value,omitempty"`
}

type AuditEventPlus struct {
	AuditEvent             `bson:",inline"`
	AuditEventPlusIncludes `bson:",inline"`
}

type AuditEventPlusIncludes struct {
	IncludedParticipantPractitionerResources  *[]Practitioner  `bson:"_includedParticipantPractitionerResources,omitempty"`
	IncludedParticipantOrganizationResources  *[]Organization  `bson:"_includedParticipantOrganizationResources,omitempty"`
	IncludedParticipantDeviceResources        *[]Device        `bson:"_includedParticipantDeviceResources,omitempty"`
	IncludedParticipantPatientResources       *[]Patient       `bson:"_includedParticipantPatientResources,omitempty"`
	IncludedParticipantRelatedPersonResources *[]RelatedPerson `bson:"_includedParticipantRelatedPersonResources,omitempty"`
	IncludedPatientPath1Resources             *[]Patient       `bson:"_includedPatientPath1Resources,omitempty"`
	IncludedPatientPath2Resources             *[]Patient       `bson:"_includedPatientPath2Resources,omitempty"`
}

func (a *AuditEventPlusIncludes) GetIncludedParticipantPractitionerResource() (practitioner *Practitioner, err error) {
	if a.IncludedParticipantPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedParticipantPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedParticipantPractitionerResources))
	} else if len(*a.IncludedParticipantPractitionerResources) == 1 {
		practitioner = &(*a.IncludedParticipantPractitionerResources)[0]
	}
	return
}

func (a *AuditEventPlusIncludes) GetIncludedParticipantOrganizationResource() (organization *Organization, err error) {
	if a.IncludedParticipantOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*a.IncludedParticipantOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*a.IncludedParticipantOrganizationResources))
	} else if len(*a.IncludedParticipantOrganizationResources) == 1 {
		organization = &(*a.IncludedParticipantOrganizationResources)[0]
	}
	return
}

func (a *AuditEventPlusIncludes) GetIncludedParticipantDeviceResource() (device *Device, err error) {
	if a.IncludedParticipantDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedParticipantDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedParticipantDeviceResources))
	} else if len(*a.IncludedParticipantDeviceResources) == 1 {
		device = &(*a.IncludedParticipantDeviceResources)[0]
	}
	return
}

func (a *AuditEventPlusIncludes) GetIncludedParticipantPatientResource() (patient *Patient, err error) {
	if a.IncludedParticipantPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedParticipantPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedParticipantPatientResources))
	} else if len(*a.IncludedParticipantPatientResources) == 1 {
		patient = &(*a.IncludedParticipantPatientResources)[0]
	}
	return
}

func (a *AuditEventPlusIncludes) GetIncludedParticipantRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedParticipantRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedParticipantRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedParticipantRelatedPersonResources))
	} else if len(*a.IncludedParticipantRelatedPersonResources) == 1 {
		relatedPerson = &(*a.IncludedParticipantRelatedPersonResources)[0]
	}
	return
}

func (a *AuditEventPlusIncludes) GetIncludedPatientPath1Resource() (patient *Patient, err error) {
	if a.IncludedPatientPath1Resources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientPath1Resources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientPath1Resources))
	} else if len(*a.IncludedPatientPath1Resources) == 1 {
		patient = &(*a.IncludedPatientPath1Resources)[0]
	}
	return
}

func (a *AuditEventPlusIncludes) GetIncludedPatientPath2Resource() (patient *Patient, err error) {
	if a.IncludedPatientPath2Resources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientPath2Resources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientPath2Resources))
	} else if len(*a.IncludedPatientPath2Resources) == 1 {
		patient = &(*a.IncludedPatientPath2Resources)[0]
	}
	return
}

func (a *AuditEventPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedParticipantPractitionerResources != nil {
		for _, r := range *a.IncludedParticipantPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedParticipantOrganizationResources != nil {
		for _, r := range *a.IncludedParticipantOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedParticipantDeviceResources != nil {
		for _, r := range *a.IncludedParticipantDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedParticipantPatientResources != nil {
		for _, r := range *a.IncludedParticipantPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedParticipantRelatedPersonResources != nil {
		for _, r := range *a.IncludedParticipantRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPatientPath1Resources != nil {
		for _, r := range *a.IncludedPatientPath1Resources {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPatientPath2Resources != nil {
		for _, r := range *a.IncludedPatientPath2Resources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
