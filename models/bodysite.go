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

type BodySite struct {
	DomainResource `bson:",inline"`
	Patient        *Reference        `bson:"patient,omitempty" json:"patient,omitempty"`
	Identifier     []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code           *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	Modifier       []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Description    string            `bson:"description,omitempty" json:"description,omitempty"`
	Image          []Attachment      `bson:"image,omitempty" json:"image,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *BodySite) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "BodySite"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to BodySite), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *BodySite) GetBSON() (interface{}, error) {
	x.ResourceType = "BodySite"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "bodySite" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type bodySite BodySite

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *BodySite) UnmarshalJSON(data []byte) (err error) {
	x2 := bodySite{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = BodySite(x2)
		return x.checkResourceType()
	}
	return
}

func (x *BodySite) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "BodySite"
	} else if x.ResourceType != "BodySite" {
		return errors.New(fmt.Sprintf("Expected resourceType to be BodySite, instead received %s", x.ResourceType))
	}
	return nil
}

type BodySitePlus struct {
	BodySite                     `bson:",inline"`
	BodySitePlusRelatedResources `bson:",inline"`
}

type BodySitePlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
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

func (b *BodySitePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if b.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*b.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*b.IncludedPatientResourcesReferencedByPatient))
	} else if len(*b.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*b.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if b.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *b.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *b.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if b.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *b.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if b.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *b.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if b.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *b.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if b.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *b.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if b.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *b.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if b.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *b.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if b.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *b.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if b.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *b.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if b.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *b.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if b.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *b.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if b.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *b.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if b.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *b.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if b.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *b.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if b.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *b.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if b.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *b.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if b.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *b.RevIncludedListResourcesReferencingItem
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if b.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *b.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if b.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *b.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if b.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *b.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if b.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *b.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if b.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *b.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if b.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *b.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if b.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *b.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if b.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *b.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if b.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *b.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if b.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *b.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *b.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (b *BodySitePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if b.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *b.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (b *BodySitePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *b.IncludedPatientResourcesReferencedByPatient {
			rsc := (*b.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (b *BodySitePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *b.RevIncludedConsentResourcesReferencingData {
			rsc := (*b.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *b.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*b.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *b.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*b.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedContractResourcesReferencingSubject {
			rsc := (*b.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *b.RevIncludedContractResourcesReferencingTopic {
			rsc := (*b.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *b.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*b.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*b.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *b.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*b.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*b.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*b.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*b.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedListResourcesReferencingItem {
			rsc := (*b.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *b.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*b.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *b.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*b.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *b.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*b.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *b.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*b.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*b.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*b.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*b.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*b.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *b.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*b.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*b.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *b.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*b.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (b *BodySitePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *b.IncludedPatientResourcesReferencedByPatient {
			rsc := (*b.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *b.RevIncludedConsentResourcesReferencingData {
			rsc := (*b.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *b.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*b.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *b.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*b.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedContractResourcesReferencingSubject {
			rsc := (*b.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *b.RevIncludedContractResourcesReferencingTopic {
			rsc := (*b.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *b.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*b.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*b.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *b.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*b.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*b.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*b.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*b.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedListResourcesReferencingItem {
			rsc := (*b.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *b.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*b.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *b.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*b.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *b.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*b.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *b.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*b.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*b.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*b.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*b.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*b.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *b.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*b.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*b.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *b.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*b.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
