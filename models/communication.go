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

type Communication struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Category       *CodeableConcept                `bson:"category,omitempty" json:"category,omitempty"`
	Sender         *Reference                      `bson:"sender,omitempty" json:"sender,omitempty"`
	Recipient      []Reference                     `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Payload        []CommunicationPayloadComponent `bson:"payload,omitempty" json:"payload,omitempty"`
	Medium         []CodeableConcept               `bson:"medium,omitempty" json:"medium,omitempty"`
	Status         string                          `bson:"status,omitempty" json:"status,omitempty"`
	Encounter      *Reference                      `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Sent           *FHIRDateTime                   `bson:"sent,omitempty" json:"sent,omitempty"`
	Received       *FHIRDateTime                   `bson:"received,omitempty" json:"received,omitempty"`
	Reason         []CodeableConcept               `bson:"reason,omitempty" json:"reason,omitempty"`
	Subject        *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	RequestDetail  *Reference                      `bson:"requestDetail,omitempty" json:"requestDetail,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Communication) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Communication"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Communication), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Communication) GetBSON() (interface{}, error) {
	x.ResourceType = "Communication"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "communication" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type communication Communication

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Communication) UnmarshalJSON(data []byte) (err error) {
	x2 := communication{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Communication(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Communication) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Communication"
	} else if x.ResourceType != "Communication" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Communication, instead received %s", x.ResourceType))
	}
	return nil
}

type CommunicationPayloadComponent struct {
	ContentString     string      `bson:"contentString,omitempty" json:"contentString,omitempty"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type CommunicationPlus struct {
	Communication             `bson:",inline"`
	CommunicationPlusIncludes `bson:",inline"`
}

type CommunicationPlusIncludes struct {
	IncludedRequestResources                *[]CommunicationRequest `bson:"_includedRequestResources,omitempty"`
	IncludedSenderPractitionerResources     *[]Practitioner         `bson:"_includedSenderPractitionerResources,omitempty"`
	IncludedSenderOrganizationResources     *[]Organization         `bson:"_includedSenderOrganizationResources,omitempty"`
	IncludedSenderDeviceResources           *[]Device               `bson:"_includedSenderDeviceResources,omitempty"`
	IncludedSenderPatientResources          *[]Patient              `bson:"_includedSenderPatientResources,omitempty"`
	IncludedSenderRelatedPersonResources    *[]RelatedPerson        `bson:"_includedSenderRelatedPersonResources,omitempty"`
	IncludedSubjectResources                *[]Patient              `bson:"_includedSubjectResources,omitempty"`
	IncludedPatientResources                *[]Patient              `bson:"_includedPatientResources,omitempty"`
	IncludedRecipientPractitionerResources  *[]Practitioner         `bson:"_includedRecipientPractitionerResources,omitempty"`
	IncludedRecipientGroupResources         *[]Group                `bson:"_includedRecipientGroupResources,omitempty"`
	IncludedRecipientOrganizationResources  *[]Organization         `bson:"_includedRecipientOrganizationResources,omitempty"`
	IncludedRecipientDeviceResources        *[]Device               `bson:"_includedRecipientDeviceResources,omitempty"`
	IncludedRecipientPatientResources       *[]Patient              `bson:"_includedRecipientPatientResources,omitempty"`
	IncludedRecipientRelatedPersonResources *[]RelatedPerson        `bson:"_includedRecipientRelatedPersonResources,omitempty"`
	IncludedEncounterResources              *[]Encounter            `bson:"_includedEncounterResources,omitempty"`
}

func (c *CommunicationPlusIncludes) GetIncludedRequestResource() (communicationRequest *CommunicationRequest, err error) {
	if c.IncludedRequestResources == nil {
		err = errors.New("Included communicationrequests not requested")
	} else if len(*c.IncludedRequestResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 communicationRequest, but found %d", len(*c.IncludedRequestResources))
	} else if len(*c.IncludedRequestResources) == 1 {
		communicationRequest = &(*c.IncludedRequestResources)[0]
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedSenderPractitionerResource() (practitioner *Practitioner, err error) {
	if c.IncludedSenderPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedSenderPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedSenderPractitionerResources))
	} else if len(*c.IncludedSenderPractitionerResources) == 1 {
		practitioner = &(*c.IncludedSenderPractitionerResources)[0]
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedSenderOrganizationResource() (organization *Organization, err error) {
	if c.IncludedSenderOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedSenderOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedSenderOrganizationResources))
	} else if len(*c.IncludedSenderOrganizationResources) == 1 {
		organization = &(*c.IncludedSenderOrganizationResources)[0]
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedSenderDeviceResource() (device *Device, err error) {
	if c.IncludedSenderDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*c.IncludedSenderDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*c.IncludedSenderDeviceResources))
	} else if len(*c.IncludedSenderDeviceResources) == 1 {
		device = &(*c.IncludedSenderDeviceResources)[0]
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedSenderPatientResource() (patient *Patient, err error) {
	if c.IncludedSenderPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedSenderPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedSenderPatientResources))
	} else if len(*c.IncludedSenderPatientResources) == 1 {
		patient = &(*c.IncludedSenderPatientResources)[0]
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedSenderRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedSenderRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedSenderRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedSenderRelatedPersonResources))
	} else if len(*c.IncludedSenderRelatedPersonResources) == 1 {
		relatedPerson = &(*c.IncludedSenderRelatedPersonResources)[0]
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedSubjectResource() (patient *Patient, err error) {
	if c.IncludedSubjectResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedSubjectResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedSubjectResources))
	} else if len(*c.IncludedSubjectResources) == 1 {
		patient = &(*c.IncludedSubjectResources)[0]
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if c.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResources))
	} else if len(*c.IncludedPatientResources) == 1 {
		patient = &(*c.IncludedPatientResources)[0]
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedRecipientPractitionerResources() (practitioners []Practitioner, err error) {
	if c.IncludedRecipientPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *c.IncludedRecipientPractitionerResources
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedRecipientGroupResources() (groups []Group, err error) {
	if c.IncludedRecipientGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else {
		groups = *c.IncludedRecipientGroupResources
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedRecipientOrganizationResources() (organizations []Organization, err error) {
	if c.IncludedRecipientOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *c.IncludedRecipientOrganizationResources
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedRecipientDeviceResources() (devices []Device, err error) {
	if c.IncludedRecipientDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *c.IncludedRecipientDeviceResources
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedRecipientPatientResources() (patients []Patient, err error) {
	if c.IncludedRecipientPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedRecipientPatientResources
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedRecipientRelatedPersonResources() (relatedPeople []RelatedPerson, err error) {
	if c.IncludedRecipientRelatedPersonResources == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *c.IncludedRecipientRelatedPersonResources
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedEncounterResource() (encounter *Encounter, err error) {
	if c.IncludedEncounterResources == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*c.IncludedEncounterResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*c.IncludedEncounterResources))
	} else if len(*c.IncludedEncounterResources) == 1 {
		encounter = &(*c.IncludedEncounterResources)[0]
	}
	return
}

func (c *CommunicationPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedRequestResources != nil {
		for _, r := range *c.IncludedRequestResources {
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
	if c.IncludedSubjectResources != nil {
		for _, r := range *c.IncludedSubjectResources {
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
	if c.IncludedRecipientGroupResources != nil {
		for _, r := range *c.IncludedRecipientGroupResources {
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
	if c.IncludedEncounterResources != nil {
		for _, r := range *c.IncludedEncounterResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
