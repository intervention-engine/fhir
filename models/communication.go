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
	BackboneElement   `bson:",inline"`
	ContentString     string      `bson:"contentString,omitempty" json:"contentString,omitempty"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type CommunicationPlus struct {
	Communication                     `bson:",inline"`
	CommunicationPlusRelatedResources `bson:",inline"`
}

type CommunicationPlusRelatedResources struct {
	IncludedCommunicationRequestResourcesReferencedByRequest    *[]CommunicationRequest  `bson:"_includedCommunicationRequestResourcesReferencedByRequest,omitempty"`
	IncludedPractitionerResourcesReferencedBySender             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySender,omitempty"`
	IncludedOrganizationResourcesReferencedBySender             *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySender,omitempty"`
	IncludedDeviceResourcesReferencedBySender                   *[]Device                `bson:"_includedDeviceResourcesReferencedBySender,omitempty"`
	IncludedPatientResourcesReferencedBySender                  *[]Patient               `bson:"_includedPatientResourcesReferencedBySender,omitempty"`
	IncludedRelatedPersonResourcesReferencedBySender            *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedBySender,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByRecipient          *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRecipient,omitempty"`
	IncludedGroupResourcesReferencedByRecipient                 *[]Group                 `bson:"_includedGroupResourcesReferencedByRecipient,omitempty"`
	IncludedOrganizationResourcesReferencedByRecipient          *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRecipient,omitempty"`
	IncludedDeviceResourcesReferencedByRecipient                *[]Device                `bson:"_includedDeviceResourcesReferencedByRecipient,omitempty"`
	IncludedPatientResourcesReferencedByRecipient               *[]Patient               `bson:"_includedPatientResourcesReferencedByRecipient,omitempty"`
	IncludedRelatedPersonResourcesReferencedByRecipient         *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByRecipient,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter             *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (c *CommunicationPlusRelatedResources) GetIncludedCommunicationRequestResourceReferencedByRequest() (communicationRequest *CommunicationRequest, err error) {
	if c.IncludedCommunicationRequestResourcesReferencedByRequest == nil {
		err = errors.New("Included communicationrequests not requested")
	} else if len(*c.IncludedCommunicationRequestResourcesReferencedByRequest) > 1 {
		err = fmt.Errorf("Expected 0 or 1 communicationRequest, but found %d", len(*c.IncludedCommunicationRequestResourcesReferencedByRequest))
	} else if len(*c.IncludedCommunicationRequestResourcesReferencedByRequest) == 1 {
		communicationRequest = &(*c.IncludedCommunicationRequestResourcesReferencedByRequest)[0]
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySender() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedBySender == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedBySender) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedBySender))
	} else if len(*c.IncludedPractitionerResourcesReferencedBySender) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedBySender)[0]
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySender() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedBySender == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedBySender) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedBySender))
	} else if len(*c.IncludedOrganizationResourcesReferencedBySender) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedBySender)[0]
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedDeviceResourceReferencedBySender() (device *Device, err error) {
	if c.IncludedDeviceResourcesReferencedBySender == nil {
		err = errors.New("Included devices not requested")
	} else if len(*c.IncludedDeviceResourcesReferencedBySender) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*c.IncludedDeviceResourcesReferencedBySender))
	} else if len(*c.IncludedDeviceResourcesReferencedBySender) == 1 {
		device = &(*c.IncludedDeviceResourcesReferencedBySender)[0]
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedPatientResourceReferencedBySender() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySender == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedBySender) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedBySender))
	} else if len(*c.IncludedPatientResourcesReferencedBySender) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedBySender)[0]
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedBySender() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedBySender == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedBySender) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedBySender))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedBySender) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedBySender)[0]
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedBySubject))
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPatient))
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByRecipient() (practitioners []Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByRecipient == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *c.IncludedPractitionerResourcesReferencedByRecipient
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedGroupResourcesReferencedByRecipient() (groups []Group, err error) {
	if c.IncludedGroupResourcesReferencedByRecipient == nil {
		err = errors.New("Included groups not requested")
	} else {
		groups = *c.IncludedGroupResourcesReferencedByRecipient
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByRecipient() (organizations []Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByRecipient == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *c.IncludedOrganizationResourcesReferencedByRecipient
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedDeviceResourcesReferencedByRecipient() (devices []Device, err error) {
	if c.IncludedDeviceResourcesReferencedByRecipient == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *c.IncludedDeviceResourcesReferencedByRecipient
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedPatientResourcesReferencedByRecipient() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedByRecipient == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedByRecipient
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByRecipient() (relatedPeople []RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByRecipient == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *c.IncludedRelatedPersonResourcesReferencedByRecipient
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if c.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*c.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*c.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*c.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*c.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *CommunicationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedCommunicationRequestResourcesReferencedByRequest != nil {
		for _, r := range *c.IncludedCommunicationRequestResourcesReferencedByRequest {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPractitionerResourcesReferencedBySender != nil {
		for _, r := range *c.IncludedPractitionerResourcesReferencedBySender {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedOrganizationResourcesReferencedBySender != nil {
		for _, r := range *c.IncludedOrganizationResourcesReferencedBySender {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedDeviceResourcesReferencedBySender != nil {
		for _, r := range *c.IncludedDeviceResourcesReferencedBySender {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedBySender != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedBySender {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedBySender != nil {
		for _, r := range *c.IncludedRelatedPersonResourcesReferencedBySender {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPractitionerResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedPractitionerResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedGroupResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedGroupResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedOrganizationResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedOrganizationResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedDeviceResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedDeviceResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedRelatedPersonResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *c.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (c *CommunicationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *c.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *c.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *c.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *c.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (c *CommunicationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedCommunicationRequestResourcesReferencedByRequest != nil {
		for _, r := range *c.IncludedCommunicationRequestResourcesReferencedByRequest {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPractitionerResourcesReferencedBySender != nil {
		for _, r := range *c.IncludedPractitionerResourcesReferencedBySender {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedOrganizationResourcesReferencedBySender != nil {
		for _, r := range *c.IncludedOrganizationResourcesReferencedBySender {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedDeviceResourcesReferencedBySender != nil {
		for _, r := range *c.IncludedDeviceResourcesReferencedBySender {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedBySender != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedBySender {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedBySender != nil {
		for _, r := range *c.IncludedRelatedPersonResourcesReferencedBySender {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPractitionerResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedPractitionerResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedGroupResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedGroupResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedOrganizationResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedOrganizationResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedDeviceResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedDeviceResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByRecipient != nil {
		for _, r := range *c.IncludedRelatedPersonResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *c.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *c.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *c.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *c.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *c.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
