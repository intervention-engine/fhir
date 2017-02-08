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

type ResearchSubject struct {
	DomainResource `bson:",inline"`
	Identifier     *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status         string      `bson:"status,omitempty" json:"status,omitempty"`
	Period         *Period     `bson:"period,omitempty" json:"period,omitempty"`
	Study          *Reference  `bson:"study,omitempty" json:"study,omitempty"`
	Individual     *Reference  `bson:"individual,omitempty" json:"individual,omitempty"`
	AssignedArm    string      `bson:"assignedArm,omitempty" json:"assignedArm,omitempty"`
	ActualArm      string      `bson:"actualArm,omitempty" json:"actualArm,omitempty"`
	Consent        *Reference  `bson:"consent,omitempty" json:"consent,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ResearchSubject) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ResearchSubject"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ResearchSubject), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ResearchSubject) GetBSON() (interface{}, error) {
	x.ResourceType = "ResearchSubject"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "researchSubject" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type researchSubject ResearchSubject

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ResearchSubject) UnmarshalJSON(data []byte) (err error) {
	x2 := researchSubject{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ResearchSubject(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ResearchSubject) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ResearchSubject"
	} else if x.ResourceType != "ResearchSubject" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ResearchSubject, instead received %s", x.ResourceType))
	}
	return nil
}

type ResearchSubjectPlus struct {
	ResearchSubject                     `bson:",inline"`
	ResearchSubjectPlusRelatedResources `bson:",inline"`
}

type ResearchSubjectPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByIndividual              *[]Patient               `bson:"_includedPatientResourcesReferencedByIndividual,omitempty"`
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

func (r *ResearchSubjectPlusRelatedResources) GetIncludedPatientResourceReferencedByIndividual() (patient *Patient, err error) {
	if r.IncludedPatientResourcesReferencedByIndividual == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResourcesReferencedByIndividual) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResourcesReferencedByIndividual))
	} else if len(*r.IncludedPatientResourcesReferencedByIndividual) == 1 {
		patient = &(*r.IncludedPatientResourcesReferencedByIndividual)[0]
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if r.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResourcesReferencedByPatient))
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*r.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if r.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *r.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if r.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *r.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if r.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *r.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if r.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *r.RevIncludedListResourcesReferencingItem
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if r.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *r.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if r.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *r.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if r.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *r.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if r.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *r.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if r.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *r.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if r.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *r.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if r.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *r.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if r.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *r.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *r.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if r.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *r.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (r *ResearchSubjectPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPatientResourcesReferencedByIndividual != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByIndividual {
			rsc := (*r.IncludedPatientResourcesReferencedByIndividual)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByPatient {
			rsc := (*r.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *ResearchSubjectPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingData {
			rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*r.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingTopic {
			rsc := (*r.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*r.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *ResearchSubjectPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPatientResourcesReferencedByIndividual != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByIndividual {
			rsc := (*r.IncludedPatientResourcesReferencedByIndividual)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByPatient {
			rsc := (*r.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingData {
			rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*r.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingTopic {
			rsc := (*r.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*r.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
