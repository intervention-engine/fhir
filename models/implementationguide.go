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

type ImplementationGuide struct {
	DomainResource `bson:",inline"`
	Url            string                                   `bson:"url,omitempty" json:"url,omitempty"`
	Version        string                                   `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                                   `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                                   `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                                    `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                                   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ImplementationGuideContactComponent    `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                            `bson:"date,omitempty" json:"date,omitempty"`
	Description    string                                   `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []CodeableConcept                        `bson:"useContext,omitempty" json:"useContext,omitempty"`
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

type ImplementationGuideContactComponent struct {
	BackboneElement `bson:",inline"`
	Name            string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
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
	Name            string                             `bson:"name,omitempty" json:"name,omitempty"`
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
	RevIncludedDocumentManifestResourcesReferencingContentref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref     *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                   *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponsereference  *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponsereference,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequestreference   *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequestreference,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource     *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment        *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                     *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingItem                        *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                     *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                    *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated         *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject    *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestreference *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequestreference,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
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

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingRequestreference
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

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if i.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *i.RevIncludedOrderResponseResourcesReferencingFulfillment
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

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if i.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *i.RevIncludedListResourcesReferencingItem
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if i.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *i.RevIncludedOrderResourcesReferencingDetail
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

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if i.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *i.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if i.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *i.RevIncludedClinicalImpressionResourcesReferencingTrigger
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
	if i.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *i.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*i.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingData)[idx]
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
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *i.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*i.RevIncludedOrderResourcesReferencingDetail)[idx]
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
	if i.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *i.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*i.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *i.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*i.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
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
	if i.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *i.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*i.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingData)[idx]
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
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *i.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*i.RevIncludedOrderResourcesReferencingDetail)[idx]
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
	if i.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *i.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*i.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *i.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*i.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
