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

type NamingSystem struct {
	DomainResource `bson:",inline"`
	Name           string                          `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                          `bson:"status,omitempty" json:"status,omitempty"`
	Kind           string                          `bson:"kind,omitempty" json:"kind,omitempty"`
	Publisher      string                          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []NamingSystemContactComponent  `bson:"contact,omitempty" json:"contact,omitempty"`
	Responsible    string                          `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Date           *FHIRDateTime                   `bson:"date,omitempty" json:"date,omitempty"`
	Type           *CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Description    string                          `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []CodeableConcept               `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Usage          string                          `bson:"usage,omitempty" json:"usage,omitempty"`
	UniqueId       []NamingSystemUniqueIdComponent `bson:"uniqueId,omitempty" json:"uniqueId,omitempty"`
	ReplacedBy     *Reference                      `bson:"replacedBy,omitempty" json:"replacedBy,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *NamingSystem) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "NamingSystem"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to NamingSystem), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *NamingSystem) GetBSON() (interface{}, error) {
	x.ResourceType = "NamingSystem"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "namingSystem" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type namingSystem NamingSystem

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *NamingSystem) UnmarshalJSON(data []byte) (err error) {
	x2 := namingSystem{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = NamingSystem(x2)
		return x.checkResourceType()
	}
	return
}

func (x *NamingSystem) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "NamingSystem"
	} else if x.ResourceType != "NamingSystem" {
		return errors.New(fmt.Sprintf("Expected resourceType to be NamingSystem, instead received %s", x.ResourceType))
	}
	return nil
}

type NamingSystemContactComponent struct {
	BackboneElement `bson:",inline"`
	Name            string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type NamingSystemUniqueIdComponent struct {
	BackboneElement `bson:",inline"`
	Type            string  `bson:"type,omitempty" json:"type,omitempty"`
	Value           string  `bson:"value,omitempty" json:"value,omitempty"`
	Preferred       *bool   `bson:"preferred,omitempty" json:"preferred,omitempty"`
	Period          *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type NamingSystemPlus struct {
	NamingSystem                     `bson:",inline"`
	NamingSystemPlusRelatedResources `bson:",inline"`
}

type NamingSystemPlusRelatedResources struct {
	IncludedNamingSystemResourcesReferencedByReplacedby         *[]NamingSystem          `bson:"_includedNamingSystemResourcesReferencedByReplacedby,omitempty"`
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
	RevIncludedNamingSystemResourcesReferencingReplacedby       *[]NamingSystem          `bson:"_revIncludedNamingSystemResourcesReferencingReplacedby,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (n *NamingSystemPlusRelatedResources) GetIncludedNamingSystemResourceReferencedByReplacedby() (namingSystem *NamingSystem, err error) {
	if n.IncludedNamingSystemResourcesReferencedByReplacedby == nil {
		err = errors.New("Included namingsystems not requested")
	} else if len(*n.IncludedNamingSystemResourcesReferencedByReplacedby) > 1 {
		err = fmt.Errorf("Expected 0 or 1 namingSystem, but found %d", len(*n.IncludedNamingSystemResourcesReferencedByReplacedby))
	} else if len(*n.IncludedNamingSystemResourcesReferencedByReplacedby) == 1 {
		namingSystem = &(*n.IncludedNamingSystemResourcesReferencedByReplacedby)[0]
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if n.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *n.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if n.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *n.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *n.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if n.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *n.RevIncludedListResourcesReferencingItem
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if n.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *n.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if n.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *n.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if n.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *n.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if n.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *n.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if n.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *n.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if n.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *n.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *n.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if n.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *n.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if n.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *n.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedNamingSystemResourcesReferencingReplacedby() (namingSystems []NamingSystem, err error) {
	if n.RevIncludedNamingSystemResourcesReferencingReplacedby == nil {
		err = errors.New("RevIncluded namingSystems not requested")
	} else {
		namingSystems = *n.RevIncludedNamingSystemResourcesReferencingReplacedby
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if n.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *n.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if n.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *n.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (n *NamingSystemPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.IncludedNamingSystemResourcesReferencedByReplacedby != nil {
		for idx := range *n.IncludedNamingSystemResourcesReferencedByReplacedby {
			rsc := (*n.IncludedNamingSystemResourcesReferencedByReplacedby)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (n *NamingSystemPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedListResourcesReferencingItem {
			rsc := (*n.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *n.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*n.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*n.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *n.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*n.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*n.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*n.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *n.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*n.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *n.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*n.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*n.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *n.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*n.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedNamingSystemResourcesReferencingReplacedby != nil {
		for idx := range *n.RevIncludedNamingSystemResourcesReferencingReplacedby {
			rsc := (*n.RevIncludedNamingSystemResourcesReferencingReplacedby)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *n.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*n.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (n *NamingSystemPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.IncludedNamingSystemResourcesReferencedByReplacedby != nil {
		for idx := range *n.IncludedNamingSystemResourcesReferencedByReplacedby {
			rsc := (*n.IncludedNamingSystemResourcesReferencedByReplacedby)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedListResourcesReferencingItem {
			rsc := (*n.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *n.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*n.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*n.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *n.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*n.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*n.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*n.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *n.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*n.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *n.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*n.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*n.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *n.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*n.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedNamingSystemResourcesReferencingReplacedby != nil {
		for idx := range *n.RevIncludedNamingSystemResourcesReferencingReplacedby {
			rsc := (*n.RevIncludedNamingSystemResourcesReferencingReplacedby)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *n.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*n.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
