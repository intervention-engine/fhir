// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
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

type ImplementationGuide struct {
	DomainResource `bson:",inline"`
	Url            string                                   `bson:"url,omitempty" json:"url,omitempty"`
	Version        string                                   `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                                   `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                                   `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                                    `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                                   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ContactDetail                          `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                            `bson:"date,omitempty" json:"date,omitempty"`
	Description    string                                   `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []UsageContext                           `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction   []CodeableConcept                        `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright      string                                   `bson:"copyright,omitempty" json:"copyright,omitempty"`
	FhirVersion    string                                   `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Dependency     []ImplementationGuideDependencyComponent `bson:"dependency,omitempty" json:"dependency,omitempty"`
	Package        []ImplementationGuidePackageComponent    `bson:"package,omitempty" json:"package,omitempty"`
	Global         []ImplementationGuideGlobalComponent     `bson:"global,omitempty" json:"global,omitempty"`
	Binary         []string                                 `bson:"binary,omitempty" json:"binary,omitempty"`
	Page           *ImplementationGuidePageComponent        `bson:"page,omitempty" json:"page,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ImplementationGuide) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ImplementationGuide"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ImplementationGuide), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ImplementationGuide) GetBSON() (interface{}, error) {
	x.ResourceType = "ImplementationGuide"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "implementationGuide" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type implementationGuide ImplementationGuide

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ImplementationGuide) UnmarshalJSON(data []byte) (err error) {
	x2 := implementationGuide{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ImplementationGuide(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ImplementationGuide) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ImplementationGuide"
	} else if x.ResourceType != "ImplementationGuide" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ImplementationGuide, instead received %s", x.ResourceType))
	}
	return nil
}

type ImplementationGuideDependencyComponent struct {
	BackboneElement `bson:",inline"`
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
	Uri             string `bson:"uri,omitempty" json:"uri,omitempty"`
}

type ImplementationGuidePackageComponent struct {
	BackboneElement `bson:",inline"`
	Name            string                                        `bson:"name,omitempty" json:"name,omitempty"`
	Description     string                                        `bson:"description,omitempty" json:"description,omitempty"`
	Resource        []ImplementationGuidePackageResourceComponent `bson:"resource,omitempty" json:"resource,omitempty"`
}

type ImplementationGuidePackageResourceComponent struct {
	BackboneElement `bson:",inline"`
	Example         *bool      `bson:"example,omitempty" json:"example,omitempty"`
	Name            string     `bson:"name,omitempty" json:"name,omitempty"`
	Description     string     `bson:"description,omitempty" json:"description,omitempty"`
	Acronym         string     `bson:"acronym,omitempty" json:"acronym,omitempty"`
	SourceUri       string     `bson:"sourceUri,omitempty" json:"sourceUri,omitempty"`
	SourceReference *Reference `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	ExampleFor      *Reference `bson:"exampleFor,omitempty" json:"exampleFor,omitempty"`
}

type ImplementationGuideGlobalComponent struct {
	BackboneElement `bson:",inline"`
	Type            string     `bson:"type,omitempty" json:"type,omitempty"`
	Profile         *Reference `bson:"profile,omitempty" json:"profile,omitempty"`
}

type ImplementationGuidePageComponent struct {
	BackboneElement `bson:",inline"`
	Source          string                             `bson:"source,omitempty" json:"source,omitempty"`
	Title           string                             `bson:"title,omitempty" json:"title,omitempty"`
	Kind            string                             `bson:"kind,omitempty" json:"kind,omitempty"`
	Type            []string                           `bson:"type,omitempty" json:"type,omitempty"`
	Package         []string                           `bson:"package,omitempty" json:"package,omitempty"`
	Format          string                             `bson:"format,omitempty" json:"format,omitempty"`
	Page            []ImplementationGuidePageComponent `bson:"page,omitempty" json:"page,omitempty"`
}

type ImplementationGuidePlus struct {
	ImplementationGuide                     `bson:",inline"`
	ImplementationGuidePlusRelatedResources `bson:",inline"`
}

type ImplementationGuidePlusRelatedResources struct {
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

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if i.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *i.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if i.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *i.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if i.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *i.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if i.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *i.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if i.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *i.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if i.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *i.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if i.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *i.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if i.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *i.RevIncludedListResourcesReferencingItem
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if i.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *i.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if i.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *i.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if i.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *i.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if i.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *i.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if i.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *i.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if i.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *i.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if i.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *i.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if i.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *i.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *i.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if i.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *i.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *i.RevIncludedConsentResourcesReferencingData {
			rsc := (*i.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*i.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingSubject {
			rsc := (*i.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingTopic {
			rsc := (*i.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*i.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*i.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*i.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*i.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*i.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*i.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*i.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*i.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*i.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*i.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (i *ImplementationGuidePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *i.RevIncludedConsentResourcesReferencingData {
			rsc := (*i.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*i.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingSubject {
			rsc := (*i.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingTopic {
			rsc := (*i.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*i.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*i.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*i.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*i.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*i.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*i.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*i.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*i.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*i.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*i.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
