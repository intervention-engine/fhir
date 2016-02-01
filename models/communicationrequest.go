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

type CommunicationRequest struct {
	DomainResource    `bson:",inline"`
	Identifier        []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Category          *CodeableConcept                       `bson:"category,omitempty" json:"category,omitempty"`
	Sender            *Reference                             `bson:"sender,omitempty" json:"sender,omitempty"`
	Recipient         []Reference                            `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Payload           []CommunicationRequestPayloadComponent `bson:"payload,omitempty" json:"payload,omitempty"`
	Medium            []CodeableConcept                      `bson:"medium,omitempty" json:"medium,omitempty"`
	Requester         *Reference                             `bson:"requester,omitempty" json:"requester,omitempty"`
	Status            string                                 `bson:"status,omitempty" json:"status,omitempty"`
	Encounter         *Reference                             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	ScheduledDateTime *FHIRDateTime                          `bson:"scheduledDateTime,omitempty" json:"scheduledDateTime,omitempty"`
	ScheduledPeriod   *Period                                `bson:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	Reason            []CodeableConcept                      `bson:"reason,omitempty" json:"reason,omitempty"`
	RequestedOn       *FHIRDateTime                          `bson:"requestedOn,omitempty" json:"requestedOn,omitempty"`
	Subject           *Reference                             `bson:"subject,omitempty" json:"subject,omitempty"`
	Priority          *CodeableConcept                       `bson:"priority,omitempty" json:"priority,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *CommunicationRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "CommunicationRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to CommunicationRequest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *CommunicationRequest) GetBSON() (interface{}, error) {
	x.ResourceType = "CommunicationRequest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "communicationRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type communicationRequest CommunicationRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *CommunicationRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := communicationRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = CommunicationRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *CommunicationRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "CommunicationRequest"
	} else if x.ResourceType != "CommunicationRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be CommunicationRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type CommunicationRequestPayloadComponent struct {
	ContentString     string      `bson:"contentString,omitempty" json:"contentString,omitempty"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type CommunicationRequestPlus struct {
	CommunicationRequest             `bson:",inline"`
	CommunicationRequestPlusIncludes `bson:",inline"`
}

type CommunicationRequestPlusIncludes struct {
	IncludedRequesterPractitionerResources  *[]Practitioner  `bson:"_includedRequesterPractitionerResources,omitempty"`
	IncludedRequesterPatientResources       *[]Patient       `bson:"_includedRequesterPatientResources,omitempty"`
	IncludedRequesterRelatedPersonResources *[]RelatedPerson `bson:"_includedRequesterRelatedPersonResources,omitempty"`
	IncludedSubjectResources                *[]Patient       `bson:"_includedSubjectResources,omitempty"`
	IncludedEncounterResources              *[]Encounter     `bson:"_includedEncounterResources,omitempty"`
	IncludedSenderPractitionerResources     *[]Practitioner  `bson:"_includedSenderPractitionerResources,omitempty"`
	IncludedSenderOrganizationResources     *[]Organization  `bson:"_includedSenderOrganizationResources,omitempty"`
	IncludedSenderDeviceResources           *[]Device        `bson:"_includedSenderDeviceResources,omitempty"`
	IncludedSenderPatientResources          *[]Patient       `bson:"_includedSenderPatientResources,omitempty"`
	IncludedSenderRelatedPersonResources    *[]RelatedPerson `bson:"_includedSenderRelatedPersonResources,omitempty"`
	IncludedPatientResources                *[]Patient       `bson:"_includedPatientResources,omitempty"`
	IncludedRecipientPractitionerResources  *[]Practitioner  `bson:"_includedRecipientPractitionerResources,omitempty"`
	IncludedRecipientOrganizationResources  *[]Organization  `bson:"_includedRecipientOrganizationResources,omitempty"`
	IncludedRecipientDeviceResources        *[]Device        `bson:"_includedRecipientDeviceResources,omitempty"`
	IncludedRecipientPatientResources       *[]Patient       `bson:"_includedRecipientPatientResources,omitempty"`
	IncludedRecipientRelatedPersonResources *[]RelatedPerson `bson:"_includedRecipientRelatedPersonResources,omitempty"`
}

func (c *CommunicationRequestPlusIncludes) GetIncludedRequesterPractitionerResource() (practitioner *Practitioner, err error) {
	if c.IncludedRequesterPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedRequesterPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedRequesterPractitionerResources))
	} else if len(*c.IncludedRequesterPractitionerResources) == 1 {
		practitioner = &(*c.IncludedRequesterPractitionerResources)[0]
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedRequesterPatientResource() (patient *Patient, err error) {
	if c.IncludedRequesterPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedRequesterPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedRequesterPatientResources))
	} else if len(*c.IncludedRequesterPatientResources) == 1 {
		patient = &(*c.IncludedRequesterPatientResources)[0]
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedRequesterRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRequesterRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRequesterRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRequesterRelatedPersonResources))
	} else if len(*c.IncludedRequesterRelatedPersonResources) == 1 {
		relatedPerson = &(*c.IncludedRequesterRelatedPersonResources)[0]
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedSubjectResource() (patient *Patient, err error) {
	if c.IncludedSubjectResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedSubjectResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedSubjectResources))
	} else if len(*c.IncludedSubjectResources) == 1 {
		patient = &(*c.IncludedSubjectResources)[0]
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if c.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*c.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*c.IncludedEncounterResources))
	} else if len(*c.IncludedEncounterResources) == 1 {
		encounter = &(*c.IncludedEncounterResources)[0]
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedSenderPractitionerResource() (practitioner *Practitioner, err error) {
	if c.IncludedSenderPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedSenderPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedSenderPractitionerResources))
	} else if len(*c.IncludedSenderPractitionerResources) == 1 {
		practitioner = &(*c.IncludedSenderPractitionerResources)[0]
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedSenderOrganizationResource() (organization *Organization, err error) {
	if c.IncludedSenderOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedSenderOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedSenderOrganizationResources))
	} else if len(*c.IncludedSenderOrganizationResources) == 1 {
		organization = &(*c.IncludedSenderOrganizationResources)[0]
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedSenderDeviceResource() (device *Device, err error) {
	if c.IncludedSenderDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*c.IncludedSenderDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*c.IncludedSenderDeviceResources))
	} else if len(*c.IncludedSenderDeviceResources) == 1 {
		device = &(*c.IncludedSenderDeviceResources)[0]
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedSenderPatientResource() (patient *Patient, err error) {
	if c.IncludedSenderPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedSenderPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedSenderPatientResources))
	} else if len(*c.IncludedSenderPatientResources) == 1 {
		patient = &(*c.IncludedSenderPatientResources)[0]
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedSenderRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedSenderRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedSenderRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedSenderRelatedPersonResources))
	} else if len(*c.IncludedSenderRelatedPersonResources) == 1 {
		relatedPerson = &(*c.IncludedSenderRelatedPersonResources)[0]
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if c.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResources))
	} else if len(*c.IncludedPatientResources) == 1 {
		patient = &(*c.IncludedPatientResources)[0]
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedRecipientPractitionerResources() (practitioners []Practitioner, err error) {
	if c.IncludedRecipientPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *c.IncludedRecipientPractitionerResources
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedRecipientOrganizationResources() (organizations []Organization, err error) {
	if c.IncludedRecipientOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *c.IncludedRecipientOrganizationResources
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedRecipientDeviceResources() (devices []Device, err error) {
	if c.IncludedRecipientDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *c.IncludedRecipientDeviceResources
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedRecipientPatientResources() (patients []Patient, err error) {
	if c.IncludedRecipientPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedRecipientPatientResources
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedRecipientRelatedPersonResources() (relatedPeople []RelatedPerson, err error) {
	if c.IncludedRecipientRelatedPersonResources == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *c.IncludedRecipientRelatedPersonResources
	}
	return
}

func (c *CommunicationRequestPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedRequesterPractitionerResources != nil {
		for _, r := range *c.IncludedRequesterPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRequesterPatientResources != nil {
		for _, r := range *c.IncludedRequesterPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRequesterRelatedPersonResources != nil {
		for _, r := range *c.IncludedRequesterRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSubjectResources != nil {
		for _, r := range *c.IncludedSubjectResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedEncounterResources != nil {
		for _, r := range *c.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSenderPractitionerResources != nil {
		for _, r := range *c.IncludedSenderPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSenderOrganizationResources != nil {
		for _, r := range *c.IncludedSenderOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSenderDeviceResources != nil {
		for _, r := range *c.IncludedSenderDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSenderPatientResources != nil {
		for _, r := range *c.IncludedSenderPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSenderRelatedPersonResources != nil {
		for _, r := range *c.IncludedSenderRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResources != nil {
		for _, r := range *c.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRecipientPractitionerResources != nil {
		for _, r := range *c.IncludedRecipientPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRecipientOrganizationResources != nil {
		for _, r := range *c.IncludedRecipientOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRecipientDeviceResources != nil {
		for _, r := range *c.IncludedRecipientDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRecipientPatientResources != nil {
		for _, r := range *c.IncludedRecipientPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRecipientRelatedPersonResources != nil {
		for _, r := range *c.IncludedRecipientRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
