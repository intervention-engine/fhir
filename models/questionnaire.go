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

type Questionnaire struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version        string                       `bson:"version,omitempty" json:"version,omitempty"`
	Status         string                       `bson:"status,omitempty" json:"status,omitempty"`
	Date           *FHIRDateTime                `bson:"date,omitempty" json:"date,omitempty"`
	Publisher      string                       `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Telecom        []ContactPoint               `bson:"telecom,omitempty" json:"telecom,omitempty"`
	SubjectType    []string                     `bson:"subjectType,omitempty" json:"subjectType,omitempty"`
	Group          *QuestionnaireGroupComponent `bson:"group,omitempty" json:"group,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Questionnaire) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Questionnaire"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Questionnaire), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Questionnaire) GetBSON() (interface{}, error) {
	x.ResourceType = "Questionnaire"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "questionnaire" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type questionnaire Questionnaire

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Questionnaire) UnmarshalJSON(data []byte) (err error) {
	x2 := questionnaire{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Questionnaire(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Questionnaire) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Questionnaire"
	} else if x.ResourceType != "Questionnaire" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Questionnaire, instead received %s", x.ResourceType))
	}
	return nil
}

type QuestionnaireGroupComponent struct {
	BackboneElement `bson:",inline"`
	LinkId          string                           `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Title           string                           `bson:"title,omitempty" json:"title,omitempty"`
	Concept         []Coding                         `bson:"concept,omitempty" json:"concept,omitempty"`
	Text            string                           `bson:"text,omitempty" json:"text,omitempty"`
	Required        *bool                            `bson:"required,omitempty" json:"required,omitempty"`
	Repeats         *bool                            `bson:"repeats,omitempty" json:"repeats,omitempty"`
	Group           []QuestionnaireGroupComponent    `bson:"group,omitempty" json:"group,omitempty"`
	Question        []QuestionnaireQuestionComponent `bson:"question,omitempty" json:"question,omitempty"`
}

type QuestionnaireQuestionComponent struct {
	BackboneElement `bson:",inline"`
	LinkId          string                        `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Concept         []Coding                      `bson:"concept,omitempty" json:"concept,omitempty"`
	Text            string                        `bson:"text,omitempty" json:"text,omitempty"`
	Type            string                        `bson:"type,omitempty" json:"type,omitempty"`
	Required        *bool                         `bson:"required,omitempty" json:"required,omitempty"`
	Repeats         *bool                         `bson:"repeats,omitempty" json:"repeats,omitempty"`
	Options         *Reference                    `bson:"options,omitempty" json:"options,omitempty"`
	Option          []Coding                      `bson:"option,omitempty" json:"option,omitempty"`
	Group           []QuestionnaireGroupComponent `bson:"group,omitempty" json:"group,omitempty"`
}

type QuestionnairePlus struct {
	Questionnaire                     `bson:",inline"`
	QuestionnairePlusRelatedResources `bson:",inline"`
}

type QuestionnairePlusRelatedResources struct {
	RevIncludedProvenanceResourcesReferencingTarget                   *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref         *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref         *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                           *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref        *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                        *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                       *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference                *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                 *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                   *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated            *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment           *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingQuestionnaire,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject       *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest             *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger          *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                  *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if q.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *q.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if q.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *q.RevIncludedListResourcesReferencingItem
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if q.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *q.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if q.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *q.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if q.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *q.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *q.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if q.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *q.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire() (questionnaireResponses []QuestionnaireResponse, err error) {
	if q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if q.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *q.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if q.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *q.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if q.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *q.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *q.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*q.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*q.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *q.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*q.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*q.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*q.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *q.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*q.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *q.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*q.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*q.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (q *QuestionnairePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *q.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*q.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*q.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *q.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*q.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*q.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*q.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *q.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*q.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *q.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*q.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*q.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
