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
	resource.ResourceType = "MessageHeader"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to MessageHeader), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *MessageHeader) GetBSON() (interface{}, error) {
	x.ResourceType = "MessageHeader"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
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
		return x.checkResourceType()
	}
	return
}

func (x *MessageHeader) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "MessageHeader"
	} else if x.ResourceType != "MessageHeader" {
		return errors.New(fmt.Sprintf("Expected resourceType to be MessageHeader, instead received %s", x.ResourceType))
	}
	return nil
}

type MessageHeaderResponseComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      string     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code            string     `bson:"code,omitempty" json:"code,omitempty"`
	Details         *Reference `bson:"details,omitempty" json:"details,omitempty"`
}

type MessageHeaderMessageSourceComponent struct {
	BackboneElement `bson:",inline"`
	Name            string        `bson:"name,omitempty" json:"name,omitempty"`
	Software        string        `bson:"software,omitempty" json:"software,omitempty"`
	Version         string        `bson:"version,omitempty" json:"version,omitempty"`
	Contact         *ContactPoint `bson:"contact,omitempty" json:"contact,omitempty"`
	Endpoint        string        `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

type MessageHeaderMessageDestinationComponent struct {
	BackboneElement `bson:",inline"`
	Name            string     `bson:"name,omitempty" json:"name,omitempty"`
	Target          *Reference `bson:"target,omitempty" json:"target,omitempty"`
	Endpoint        string     `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

type MessageHeaderPlus struct {
	MessageHeader                     `bson:",inline"`
	MessageHeaderPlusRelatedResources `bson:",inline"`
}

type MessageHeaderPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByReceiver           *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByReceiver,omitempty"`
	IncludedOrganizationResourcesReferencedByReceiver           *[]Organization          `bson:"_includedOrganizationResourcesReferencedByReceiver,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthor             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByTarget                   *[]Device                `bson:"_includedDeviceResourcesReferencedByTarget,omitempty"`
	IncludedPractitionerResourcesReferencedByResponsible        *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByResponsible,omitempty"`
	IncludedOrganizationResourcesReferencedByResponsible        *[]Organization          `bson:"_includedOrganizationResourcesReferencedByResponsible,omitempty"`
	IncludedPractitionerResourcesReferencedByEnterer            *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByEnterer,omitempty"`
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
	RevIncludedBundleResourcesReferencingMessage                *[]Bundle                `bson:"_revIncludedBundleResourcesReferencingMessage,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (m *MessageHeaderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByReceiver() (practitioner *Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByReceiver == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedPractitionerResourcesReferencedByReceiver) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedPractitionerResourcesReferencedByReceiver))
	} else if len(*m.IncludedPractitionerResourcesReferencedByReceiver) == 1 {
		practitioner = &(*m.IncludedPractitionerResourcesReferencedByReceiver)[0]
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetIncludedOrganizationResourceReferencedByReceiver() (organization *Organization, err error) {
	if m.IncludedOrganizationResourcesReferencedByReceiver == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*m.IncludedOrganizationResourcesReferencedByReceiver) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*m.IncludedOrganizationResourcesReferencedByReceiver))
	} else if len(*m.IncludedOrganizationResourcesReferencedByReceiver) == 1 {
		organization = &(*m.IncludedOrganizationResourcesReferencedByReceiver)[0]
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*m.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*m.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetIncludedDeviceResourceReferencedByTarget() (device *Device, err error) {
	if m.IncludedDeviceResourcesReferencedByTarget == nil {
		err = errors.New("Included devices not requested")
	} else if len(*m.IncludedDeviceResourcesReferencedByTarget) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*m.IncludedDeviceResourcesReferencedByTarget))
	} else if len(*m.IncludedDeviceResourcesReferencedByTarget) == 1 {
		device = &(*m.IncludedDeviceResourcesReferencedByTarget)[0]
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByResponsible() (practitioner *Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByResponsible == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedPractitionerResourcesReferencedByResponsible) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedPractitionerResourcesReferencedByResponsible))
	} else if len(*m.IncludedPractitionerResourcesReferencedByResponsible) == 1 {
		practitioner = &(*m.IncludedPractitionerResourcesReferencedByResponsible)[0]
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetIncludedOrganizationResourceReferencedByResponsible() (organization *Organization, err error) {
	if m.IncludedOrganizationResourcesReferencedByResponsible == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*m.IncludedOrganizationResourcesReferencedByResponsible) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*m.IncludedOrganizationResourcesReferencedByResponsible))
	} else if len(*m.IncludedOrganizationResourcesReferencedByResponsible) == 1 {
		organization = &(*m.IncludedOrganizationResourcesReferencedByResponsible)[0]
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByEnterer() (practitioner *Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByEnterer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedPractitionerResourcesReferencedByEnterer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedPractitionerResourcesReferencedByEnterer))
	} else if len(*m.IncludedPractitionerResourcesReferencedByEnterer) == 1 {
		practitioner = &(*m.IncludedPractitionerResourcesReferencedByEnterer)[0]
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if m.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *m.RevIncludedListResourcesReferencingItem
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if m.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *m.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if m.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *m.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if m.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *m.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *m.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedBundleResourcesReferencingMessage() (bundles []Bundle, err error) {
	if m.RevIncludedBundleResourcesReferencingMessage == nil {
		err = errors.New("RevIncluded bundles not requested")
	} else {
		bundles = *m.RevIncludedBundleResourcesReferencingMessage
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *m.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if m.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *m.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *m.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if m.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *m.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (m *MessageHeaderPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedPractitionerResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByReceiver {
			rsc := (*m.IncludedPractitionerResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedByReceiver {
			rsc := (*m.IncludedOrganizationResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*m.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedDeviceResourcesReferencedByTarget != nil {
		for idx := range *m.IncludedDeviceResourcesReferencedByTarget {
			rsc := (*m.IncludedDeviceResourcesReferencedByTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByResponsible != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByResponsible {
			rsc := (*m.IncludedPractitionerResourcesReferencedByResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedByResponsible != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedByResponsible {
			rsc := (*m.IncludedOrganizationResourcesReferencedByResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByEnterer != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByEnterer {
			rsc := (*m.IncludedPractitionerResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MessageHeaderPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *m.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*m.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBundleResourcesReferencingMessage != nil {
		for idx := range *m.RevIncludedBundleResourcesReferencingMessage {
			rsc := (*m.RevIncludedBundleResourcesReferencingMessage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *m.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*m.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MessageHeaderPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedPractitionerResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByReceiver {
			rsc := (*m.IncludedPractitionerResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedByReceiver != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedByReceiver {
			rsc := (*m.IncludedOrganizationResourcesReferencedByReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*m.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedDeviceResourcesReferencedByTarget != nil {
		for idx := range *m.IncludedDeviceResourcesReferencedByTarget {
			rsc := (*m.IncludedDeviceResourcesReferencedByTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByResponsible != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByResponsible {
			rsc := (*m.IncludedPractitionerResourcesReferencedByResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedByResponsible != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedByResponsible {
			rsc := (*m.IncludedOrganizationResourcesReferencedByResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByEnterer != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByEnterer {
			rsc := (*m.IncludedPractitionerResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *m.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*m.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBundleResourcesReferencingMessage != nil {
		for idx := range *m.RevIncludedBundleResourcesReferencingMessage {
			rsc := (*m.RevIncludedBundleResourcesReferencingMessage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *m.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*m.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
