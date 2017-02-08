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

type Linkage struct {
	DomainResource `bson:",inline"`
	Author         *Reference             `bson:"author,omitempty" json:"author,omitempty"`
	Item           []LinkageItemComponent `bson:"item,omitempty" json:"item,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Linkage) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Linkage"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Linkage), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Linkage) GetBSON() (interface{}, error) {
	x.ResourceType = "Linkage"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "linkage" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type linkage Linkage

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Linkage) UnmarshalJSON(data []byte) (err error) {
	x2 := linkage{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Linkage(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Linkage) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Linkage"
	} else if x.ResourceType != "Linkage" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Linkage, instead received %s", x.ResourceType))
	}
	return nil
}

type LinkageItemComponent struct {
	BackboneElement `bson:",inline"`
	Type            string     `bson:"type,omitempty" json:"type,omitempty"`
	Resource        *Reference `bson:"resource,omitempty" json:"resource,omitempty"`
}

type LinkagePlus struct {
	Linkage                     `bson:",inline"`
	LinkagePlusRelatedResources `bson:",inline"`
}

type LinkagePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByAuthor             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthor             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                  *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic               *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject              *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse        *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource  *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon         *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                    *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces    *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition  *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces     *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition   *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity             *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
}

func (l *LinkagePlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if l.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*l.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*l.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*l.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*l.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (l *LinkagePlusRelatedResources) GetIncludedOrganizationResourceReferencedByAuthor() (organization *Organization, err error) {
	if l.IncludedOrganizationResourcesReferencedByAuthor == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*l.IncludedOrganizationResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*l.IncludedOrganizationResourcesReferencedByAuthor))
	} else if len(*l.IncludedOrganizationResourcesReferencedByAuthor) == 1 {
		organization = &(*l.IncludedOrganizationResourcesReferencedByAuthor)[0]
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if l.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *l.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *l.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *l.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if l.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *l.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if l.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *l.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if l.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *l.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if l.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *l.RevIncludedListResourcesReferencingItem
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if l.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *l.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if l.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *l.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if l.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *l.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if l.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *l.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if l.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *l.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if l.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *l.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if l.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *l.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if l.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *l.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *l.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if l.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *l.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (l *LinkagePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *l.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*l.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *l.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*l.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (l *LinkagePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *l.RevIncludedConsentResourcesReferencingData {
			rsc := (*l.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*l.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingSubject {
			rsc := (*l.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingTopic {
			rsc := (*l.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *l.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*l.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *l.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*l.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*l.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*l.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*l.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedListResourcesReferencingItem {
			rsc := (*l.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*l.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*l.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*l.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*l.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*l.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*l.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*l.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (l *LinkagePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *l.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*l.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *l.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*l.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *l.RevIncludedConsentResourcesReferencingData {
			rsc := (*l.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*l.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingSubject {
			rsc := (*l.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingTopic {
			rsc := (*l.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *l.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*l.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *l.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*l.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*l.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*l.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*l.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedListResourcesReferencingItem {
			rsc := (*l.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*l.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*l.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*l.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*l.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*l.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*l.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*l.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
