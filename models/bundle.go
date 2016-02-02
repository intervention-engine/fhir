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

type Bundle struct {
	Resource  `bson:",inline"`
	Type      string                 `bson:"type,omitempty" json:"type,omitempty"`
	Total     *uint32                `bson:"total,omitempty" json:"total,omitempty"`
	Link      []BundleLinkComponent  `bson:"link,omitempty" json:"link,omitempty"`
	Entry     []BundleEntryComponent `bson:"entry,omitempty" json:"entry,omitempty"`
	Signature *Signature             `bson:"signature,omitempty" json:"signature,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Bundle) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Bundle"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Bundle), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Bundle) GetBSON() (interface{}, error) {
	x.ResourceType = "Bundle"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "bundle" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type bundle Bundle

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Bundle) UnmarshalJSON(data []byte) (err error) {
	x2 := bundle{}
	if err = json.Unmarshal(data, &x2); err == nil {
		*x = Bundle(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Bundle) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Bundle"
	} else if x.ResourceType != "Bundle" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Bundle, instead received %s", x.ResourceType))
	}
	return nil
}

type BundleLinkComponent struct {
	Relation string `bson:"relation,omitempty" json:"relation,omitempty"`
	Url      string `bson:"url,omitempty" json:"url,omitempty"`
}

type BundleEntryComponent struct {
	Link     []BundleLinkComponent         `bson:"link,omitempty" json:"link,omitempty"`
	FullUrl  string                        `bson:"fullUrl,omitempty" json:"fullUrl,omitempty"`
	Resource interface{}                   `bson:"resource,omitempty" json:"resource,omitempty"`
	Search   *BundleEntrySearchComponent   `bson:"search,omitempty" json:"search,omitempty"`
	Request  *BundleEntryRequestComponent  `bson:"request,omitempty" json:"request,omitempty"`
	Response *BundleEntryResponseComponent `bson:"response,omitempty" json:"response,omitempty"`
}

// The "bundleEntryComponent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type bundleEntryComponent BundleEntryComponent

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *BundleEntryComponent) UnmarshalJSON(data []byte) (err error) {
	x2 := bundleEntryComponent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Resource != nil {
			x2.Resource = MapToResource(x2.Resource, true)
		}
		*x = BundleEntryComponent(x2)
	}
	return
}

type BundleEntrySearchComponent struct {
	Mode  string   `bson:"mode,omitempty" json:"mode,omitempty"`
	Score *float64 `bson:"score,omitempty" json:"score,omitempty"`
}

type BundleEntryRequestComponent struct {
	Method          string        `bson:"method,omitempty" json:"method,omitempty"`
	Url             string        `bson:"url,omitempty" json:"url,omitempty"`
	IfNoneMatch     string        `bson:"ifNoneMatch,omitempty" json:"ifNoneMatch,omitempty"`
	IfModifiedSince *FHIRDateTime `bson:"ifModifiedSince,omitempty" json:"ifModifiedSince,omitempty"`
	IfMatch         string        `bson:"ifMatch,omitempty" json:"ifMatch,omitempty"`
	IfNoneExist     string        `bson:"ifNoneExist,omitempty" json:"ifNoneExist,omitempty"`
}

type BundleEntryResponseComponent struct {
	Status       string        `bson:"status,omitempty" json:"status,omitempty"`
	Location     string        `bson:"location,omitempty" json:"location,omitempty"`
	Etag         string        `bson:"etag,omitempty" json:"etag,omitempty"`
	LastModified *FHIRDateTime `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
}

type BundlePlus struct {
	Bundle                     `bson:",inline"`
	BundlePlusRelatedResources `bson:",inline"`
}

type BundlePlusRelatedResources struct {
	IncludedCompositionResourcesReferencedByComposition         *[]Composition           `bson:"_includedCompositionResourcesReferencedByComposition,omitempty"`
	IncludedMessageHeaderResourcesReferencedByMessage           *[]MessageHeader         `bson:"_includedMessageHeaderResourcesReferencedByMessage,omitempty"`
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

func (b *BundlePlusRelatedResources) GetIncludedCompositionResourceReferencedByComposition() (composition *Composition, err error) {
	if b.IncludedCompositionResourcesReferencedByComposition == nil {
		err = errors.New("Included compositions not requested")
	} else if len(*b.IncludedCompositionResourcesReferencedByComposition) > 1 {
		err = fmt.Errorf("Expected 0 or 1 composition, but found %d", len(*b.IncludedCompositionResourcesReferencedByComposition))
	} else if len(*b.IncludedCompositionResourcesReferencedByComposition) == 1 {
		composition = &(*b.IncludedCompositionResourcesReferencedByComposition)[0]
	}
	return
}

func (b *BundlePlusRelatedResources) GetIncludedMessageHeaderResourceReferencedByMessage() (messageHeader *MessageHeader, err error) {
	if b.IncludedMessageHeaderResourcesReferencedByMessage == nil {
		err = errors.New("Included messageheaders not requested")
	} else if len(*b.IncludedMessageHeaderResourcesReferencedByMessage) > 1 {
		err = fmt.Errorf("Expected 0 or 1 messageHeader, but found %d", len(*b.IncludedMessageHeaderResourcesReferencedByMessage))
	} else if len(*b.IncludedMessageHeaderResourcesReferencedByMessage) == 1 {
		messageHeader = &(*b.IncludedMessageHeaderResourcesReferencedByMessage)[0]
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if b.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *b.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if b.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *b.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *b.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if b.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *b.RevIncludedListResourcesReferencingItem
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if b.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *b.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if b.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *b.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if b.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *b.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if b.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *b.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if b.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *b.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if b.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *b.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *b.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if b.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *b.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if b.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *b.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if b.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *b.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if b.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *b.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (b *BundlePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.IncludedCompositionResourcesReferencedByComposition != nil {
		for _, r := range *b.IncludedCompositionResourcesReferencedByComposition {
			resourceMap[r.Id] = &r
		}
	}
	if b.IncludedMessageHeaderResourcesReferencedByMessage != nil {
		for _, r := range *b.IncludedMessageHeaderResourcesReferencedByMessage {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (b *BundlePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *b.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *b.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *b.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *b.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *b.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *b.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *b.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *b.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *b.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *b.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *b.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *b.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *b.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *b.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *b.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (b *BundlePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.IncludedCompositionResourcesReferencedByComposition != nil {
		for _, r := range *b.IncludedCompositionResourcesReferencedByComposition {
			resourceMap[r.Id] = &r
		}
	}
	if b.IncludedMessageHeaderResourcesReferencedByMessage != nil {
		for _, r := range *b.IncludedMessageHeaderResourcesReferencedByMessage {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *b.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *b.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *b.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *b.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *b.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *b.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *b.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *b.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *b.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *b.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *b.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *b.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *b.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *b.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if b.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *b.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
