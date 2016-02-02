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
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type ImplementationGuideDependencyComponent struct {
	Type string `bson:"type,omitempty" json:"type,omitempty"`
	Uri  string `bson:"uri,omitempty" json:"uri,omitempty"`
}

type ImplementationGuidePackageComponent struct {
	Name        string                                        `bson:"name,omitempty" json:"name,omitempty"`
	Description string                                        `bson:"description,omitempty" json:"description,omitempty"`
	Resource    []ImplementationGuidePackageResourceComponent `bson:"resource,omitempty" json:"resource,omitempty"`
}

type ImplementationGuidePackageResourceComponent struct {
	Purpose         string     `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Name            string     `bson:"name,omitempty" json:"name,omitempty"`
	Description     string     `bson:"description,omitempty" json:"description,omitempty"`
	Acronym         string     `bson:"acronym,omitempty" json:"acronym,omitempty"`
	SourceUri       string     `bson:"sourceUri,omitempty" json:"sourceUri,omitempty"`
	SourceReference *Reference `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	ExampleFor      *Reference `bson:"exampleFor,omitempty" json:"exampleFor,omitempty"`
}

type ImplementationGuideGlobalComponent struct {
	Type    string     `bson:"type,omitempty" json:"type,omitempty"`
	Profile *Reference `bson:"profile,omitempty" json:"profile,omitempty"`
}

type ImplementationGuidePageComponent struct {
	Source  string                             `bson:"source,omitempty" json:"source,omitempty"`
	Name    string                             `bson:"name,omitempty" json:"name,omitempty"`
	Kind    string                             `bson:"kind,omitempty" json:"kind,omitempty"`
	Type    []string                           `bson:"type,omitempty" json:"type,omitempty"`
	Package []string                           `bson:"package,omitempty" json:"package,omitempty"`
	Format  string                             `bson:"format,omitempty" json:"format,omitempty"`
	Page    []ImplementationGuidePageComponent `bson:"page,omitempty" json:"page,omitempty"`
}

type ImplementationGuidePlus struct {
	ImplementationGuide                     `bson:",inline"`
	ImplementationGuidePlusRelatedResources `bson:",inline"`
}

type ImplementationGuidePlusRelatedResources struct {
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

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
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

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if i.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *i.RevIncludedListResourcesReferencingItem
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

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if i.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *i.RevIncludedAuditEventResourcesReferencingReference
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

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if i.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *i.RevIncludedOrderResponseResourcesReferencingFulfillment
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

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if i.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *i.RevIncludedClinicalImpressionResourcesReferencingTrigger
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

func (i *ImplementationGuidePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *i.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *i.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *i.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *i.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *i.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *i.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *i.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *i.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *i.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (i *ImplementationGuidePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *i.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *i.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *i.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *i.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *i.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *i.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *i.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *i.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *i.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
