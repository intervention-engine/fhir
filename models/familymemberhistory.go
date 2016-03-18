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

type FamilyMemberHistory struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier                            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient         *Reference                              `bson:"patient,omitempty" json:"patient,omitempty"`
	Date            *FHIRDateTime                           `bson:"date,omitempty" json:"date,omitempty"`
	Status          string                                  `bson:"status,omitempty" json:"status,omitempty"`
	Name            string                                  `bson:"name,omitempty" json:"name,omitempty"`
	Relationship    *CodeableConcept                        `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Gender          string                                  `bson:"gender,omitempty" json:"gender,omitempty"`
	BornPeriod      *Period                                 `bson:"bornPeriod,omitempty" json:"bornPeriod,omitempty"`
	BornDate        *FHIRDateTime                           `bson:"bornDate,omitempty" json:"bornDate,omitempty"`
	BornString      string                                  `bson:"bornString,omitempty" json:"bornString,omitempty"`
	AgeAge          *Quantity                               `bson:"ageAge,omitempty" json:"ageAge,omitempty"`
	AgeRange        *Range                                  `bson:"ageRange,omitempty" json:"ageRange,omitempty"`
	AgeString       string                                  `bson:"ageString,omitempty" json:"ageString,omitempty"`
	DeceasedBoolean *bool                                   `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedAge     *Quantity                               `bson:"deceasedAge,omitempty" json:"deceasedAge,omitempty"`
	DeceasedRange   *Range                                  `bson:"deceasedRange,omitempty" json:"deceasedRange,omitempty"`
	DeceasedDate    *FHIRDateTime                           `bson:"deceasedDate,omitempty" json:"deceasedDate,omitempty"`
	DeceasedString  string                                  `bson:"deceasedString,omitempty" json:"deceasedString,omitempty"`
	Note            *Annotation                             `bson:"note,omitempty" json:"note,omitempty"`
	Condition       []FamilyMemberHistoryConditionComponent `bson:"condition,omitempty" json:"condition,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "FamilyMemberHistory"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to FamilyMemberHistory), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *FamilyMemberHistory) GetBSON() (interface{}, error) {
	x.ResourceType = "FamilyMemberHistory"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "familyMemberHistory" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type familyMemberHistory FamilyMemberHistory

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *FamilyMemberHistory) UnmarshalJSON(data []byte) (err error) {
	x2 := familyMemberHistory{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = FamilyMemberHistory(x2)
		return x.checkResourceType()
	}
	return
}

func (x *FamilyMemberHistory) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "FamilyMemberHistory"
	} else if x.ResourceType != "FamilyMemberHistory" {
		return errors.New(fmt.Sprintf("Expected resourceType to be FamilyMemberHistory, instead received %s", x.ResourceType))
	}
	return nil
}

type FamilyMemberHistoryConditionComponent struct {
	BackboneElement `bson:",inline"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Outcome         *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OnsetAge        *Quantity        `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetRange      *Range           `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetPeriod     *Period          `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetString     string           `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	Note            *Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}

type FamilyMemberHistoryPlus struct {
	FamilyMemberHistory                     `bson:",inline"`
	FamilyMemberHistoryPlusRelatedResources `bson:",inline"`
}

type FamilyMemberHistoryPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                    *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                        *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref     *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                     *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                    *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference             *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated         *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment        *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject    *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest          *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingInvestigation *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingInvestigation,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if f.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedPatientResourcesReferencedByPatient))
	} else if len(*f.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*f.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if f.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *f.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if f.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *f.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *f.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if f.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *f.RevIncludedListResourcesReferencingItem
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if f.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *f.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if f.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *f.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if f.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *f.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if f.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *f.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if f.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *f.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if f.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *f.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *f.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if f.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *f.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if f.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *f.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if f.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *f.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingInvestigation() (clinicalImpressions []ClinicalImpression, err error) {
	if f.RevIncludedClinicalImpressionResourcesReferencingInvestigation == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *f.RevIncludedClinicalImpressionResourcesReferencingInvestigation
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if f.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *f.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByPatient {
			rsc := (*f.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedListResourcesReferencingItem {
			rsc := (*f.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *f.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*f.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*f.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *f.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*f.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*f.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*f.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *f.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*f.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *f.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*f.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*f.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *f.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*f.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *f.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*f.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *f.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*f.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *f.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*f.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (f *FamilyMemberHistoryPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByPatient {
			rsc := (*f.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedListResourcesReferencingItem {
			rsc := (*f.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *f.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*f.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*f.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *f.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*f.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*f.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*f.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *f.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*f.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *f.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*f.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*f.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *f.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*f.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *f.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*f.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *f.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*f.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *f.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*f.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
