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

type MessageHeader struct {
	DomainResource `bson:",inline"`
	Timestamp      *FHIRDateTime                              `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Event          *Coding                                    `bson:"event,omitempty" json:"event,omitempty"`
	Response       *MessageHeaderResponseComponent            `bson:"response,omitempty" json:"response,omitempty"`
	Source         *MessageHeaderMessageSourceComponent       `bson:"source,omitempty" json:"source,omitempty"`
	Destination    []MessageHeaderMessageDestinationComponent `bson:"destination,omitempty" json:"destination,omitempty"`
	Enterer        *Reference                                 `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Author         *Reference                                 `bson:"author,omitempty" json:"author,omitempty"`
	Receiver       *Reference                                 `bson:"receiver,omitempty" json:"receiver,omitempty"`
	Responsible    *Reference                                 `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Reason         *CodeableConcept                           `bson:"reason,omitempty" json:"reason,omitempty"`
	Data           []Reference                                `bson:"data,omitempty" json:"data,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *MessageHeader) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		MessageHeader
	}{
		ResourceType:  "MessageHeader",
		MessageHeader: *resource,
	}
	return json.Marshal(x)
}

// The "messageHeader" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type messageHeader MessageHeader

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *MessageHeader) UnmarshalJSON(data []byte) (err error) {
	x2 := messageHeader{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MessageHeader(x2)
	}
	return
}

type MessageHeaderResponseComponent struct {
	Identifier string     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code       string     `bson:"code,omitempty" json:"code,omitempty"`
	Details    *Reference `bson:"details,omitempty" json:"details,omitempty"`
}

type MessageHeaderMessageSourceComponent struct {
	Name     string        `bson:"name,omitempty" json:"name,omitempty"`
	Software string        `bson:"software,omitempty" json:"software,omitempty"`
	Version  string        `bson:"version,omitempty" json:"version,omitempty"`
	Contact  *ContactPoint `bson:"contact,omitempty" json:"contact,omitempty"`
	Endpoint string        `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

type MessageHeaderMessageDestinationComponent struct {
	Name     string     `bson:"name,omitempty" json:"name,omitempty"`
	Target   *Reference `bson:"target,omitempty" json:"target,omitempty"`
	Endpoint string     `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

type MessageHeaderPlus struct {
	MessageHeader             `bson:",inline"`
	MessageHeaderPlusIncludes `bson:",inline"`
}

type MessageHeaderPlusIncludes struct {
	IncludedReceiverPractitionerResources    *[]Practitioner `bson:"_includedReceiverPractitionerResources,omitempty"`
	IncludedReceiverOrganizationResources    *[]Organization `bson:"_includedReceiverOrganizationResources,omitempty"`
	IncludedAuthorResources                  *[]Practitioner `bson:"_includedAuthorResources,omitempty"`
	IncludedTargetResources                  *[]Device       `bson:"_includedTargetResources,omitempty"`
	IncludedResponsiblePractitionerResources *[]Practitioner `bson:"_includedResponsiblePractitionerResources,omitempty"`
	IncludedResponsibleOrganizationResources *[]Organization `bson:"_includedResponsibleOrganizationResources,omitempty"`
	IncludedEntererResources                 *[]Practitioner `bson:"_includedEntererResources,omitempty"`
}

func (m *MessageHeaderPlusIncludes) GetIncludedReceiverPractitionerResource() (practitioner *Practitioner, err error) {
	if m.IncludedReceiverPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedReceiverPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedReceiverPractitionerResources))
	} else if len(*m.IncludedReceiverPractitionerResources) == 1 {
		practitioner = &(*m.IncludedReceiverPractitionerResources)[0]
	}
	return
}

func (m *MessageHeaderPlusIncludes) GetIncludedReceiverOrganizationResource() (organization *Organization, err error) {
	if m.IncludedReceiverOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*m.IncludedReceiverOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*m.IncludedReceiverOrganizationResources))
	} else if len(*m.IncludedReceiverOrganizationResources) == 1 {
		organization = &(*m.IncludedReceiverOrganizationResources)[0]
	}
	return
}

func (m *MessageHeaderPlusIncludes) GetIncludedAuthorResource() (practitioner *Practitioner, err error) {
	if m.IncludedAuthorResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedAuthorResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedAuthorResources))
	} else if len(*m.IncludedAuthorResources) == 1 {
		practitioner = &(*m.IncludedAuthorResources)[0]
	}
	return
}

func (m *MessageHeaderPlusIncludes) GetIncludedTargetResource() (device *Device, err error) {
	if m.IncludedTargetResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*m.IncludedTargetResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*m.IncludedTargetResources))
	} else if len(*m.IncludedTargetResources) == 1 {
		device = &(*m.IncludedTargetResources)[0]
	}
	return
}

func (m *MessageHeaderPlusIncludes) GetIncludedResponsiblePractitionerResource() (practitioner *Practitioner, err error) {
	if m.IncludedResponsiblePractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedResponsiblePractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedResponsiblePractitionerResources))
	} else if len(*m.IncludedResponsiblePractitionerResources) == 1 {
		practitioner = &(*m.IncludedResponsiblePractitionerResources)[0]
	}
	return
}

func (m *MessageHeaderPlusIncludes) GetIncludedResponsibleOrganizationResource() (organization *Organization, err error) {
	if m.IncludedResponsibleOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*m.IncludedResponsibleOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*m.IncludedResponsibleOrganizationResources))
	} else if len(*m.IncludedResponsibleOrganizationResources) == 1 {
		organization = &(*m.IncludedResponsibleOrganizationResources)[0]
	}
	return
}

func (m *MessageHeaderPlusIncludes) GetIncludedEntererResource() (practitioner *Practitioner, err error) {
	if m.IncludedEntererResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedEntererResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedEntererResources))
	} else if len(*m.IncludedEntererResources) == 1 {
		practitioner = &(*m.IncludedEntererResources)[0]
	}
	return
}

func (m *MessageHeaderPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedReceiverPractitionerResources != nil {
		for _, r := range *m.IncludedReceiverPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedReceiverOrganizationResources != nil {
		for _, r := range *m.IncludedReceiverOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedAuthorResources != nil {
		for _, r := range *m.IncludedAuthorResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedTargetResources != nil {
		for _, r := range *m.IncludedTargetResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedResponsiblePractitionerResources != nil {
		for _, r := range *m.IncludedResponsiblePractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedResponsibleOrganizationResources != nil {
		for _, r := range *m.IncludedResponsibleOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if m.IncludedEntererResources != nil {
		for _, r := range *m.IncludedEntererResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
