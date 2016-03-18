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

type AllergyIntolerance struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Onset          *FHIRDateTime                         `bson:"onset,omitempty" json:"onset,omitempty"`
	RecordedDate   *FHIRDateTime                         `bson:"recordedDate,omitempty" json:"recordedDate,omitempty"`
	Recorder       *Reference                            `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Patient        *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	Reporter       *Reference                            `bson:"reporter,omitempty" json:"reporter,omitempty"`
	Substance      *CodeableConcept                      `bson:"substance,omitempty" json:"substance,omitempty"`
	Status         string                                `bson:"status,omitempty" json:"status,omitempty"`
	Criticality    string                                `bson:"criticality,omitempty" json:"criticality,omitempty"`
	Type           string                                `bson:"type,omitempty" json:"type,omitempty"`
	Category       string                                `bson:"category,omitempty" json:"category,omitempty"`
	LastOccurence  *FHIRDateTime                         `bson:"lastOccurence,omitempty" json:"lastOccurence,omitempty"`
	Note           *Annotation                           `bson:"note,omitempty" json:"note,omitempty"`
	Reaction       []AllergyIntoleranceReactionComponent `bson:"reaction,omitempty" json:"reaction,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *AllergyIntolerance) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "AllergyIntolerance"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to AllergyIntolerance), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *AllergyIntolerance) GetBSON() (interface{}, error) {
	x.ResourceType = "AllergyIntolerance"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "allergyIntolerance" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type allergyIntolerance AllergyIntolerance

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *AllergyIntolerance) UnmarshalJSON(data []byte) (err error) {
	x2 := allergyIntolerance{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = AllergyIntolerance(x2)
		return x.checkResourceType()
	}
	return
}

func (x *AllergyIntolerance) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "AllergyIntolerance"
	} else if x.ResourceType != "AllergyIntolerance" {
		return errors.New(fmt.Sprintf("Expected resourceType to be AllergyIntolerance, instead received %s", x.ResourceType))
	}
	return nil
}

type AllergyIntoleranceReactionComponent struct {
	BackboneElement `bson:",inline"`
	Substance       *CodeableConcept  `bson:"substance,omitempty" json:"substance,omitempty"`
	Certainty       string            `bson:"certainty,omitempty" json:"certainty,omitempty"`
	Manifestation   []CodeableConcept `bson:"manifestation,omitempty" json:"manifestation,omitempty"`
	Description     string            `bson:"description,omitempty" json:"description,omitempty"`
	Onset           *FHIRDateTime     `bson:"onset,omitempty" json:"onset,omitempty"`
	Severity        string            `bson:"severity,omitempty" json:"severity,omitempty"`
	ExposureRoute   *CodeableConcept  `bson:"exposureRoute,omitempty" json:"exposureRoute,omitempty"`
	Note            *Annotation       `bson:"note,omitempty" json:"note,omitempty"`
}

type AllergyIntolerancePlus struct {
	AllergyIntolerance                     `bson:",inline"`
	AllergyIntolerancePlusRelatedResources `bson:",inline"`
}

type AllergyIntolerancePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByRecorder                    *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByRecorder,omitempty"`
	IncludedPatientResourcesReferencedByRecorder                         *[]Patient                    `bson:"_includedPatientResourcesReferencedByRecorder,omitempty"`
	IncludedPractitionerResourcesReferencedByReporter                    *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByReporter,omitempty"`
	IncludedPatientResourcesReferencedByReporter                         *[]Patient                    `bson:"_includedPatientResourcesReferencedByReporter,omitempty"`
	IncludedRelatedPersonResourcesReferencedByReporter                   *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByReporter,omitempty"`
	IncludedPatientResourcesReferencedByPatient                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                      *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref            *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref            *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                              *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref           *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                           *[]Order                      `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                          *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference                   *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                    *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                      *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated               *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment              *[]OrderResponse              `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject          *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest                *[]ProcessResponse            `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger             *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingProblem             *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingProblem,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                     *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPractitionerResourceReferencedByRecorder() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByRecorder == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByRecorder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByRecorder))
	} else if len(*a.IncludedPractitionerResourcesReferencedByRecorder) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByRecorder)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPatientResourceReferencedByRecorder() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByRecorder == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByRecorder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByRecorder))
	} else if len(*a.IncludedPatientResourcesReferencedByRecorder) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByRecorder)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPractitionerResourceReferencedByReporter() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByReporter == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByReporter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByReporter))
	} else if len(*a.IncludedPractitionerResourcesReferencedByReporter) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByReporter)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPatientResourceReferencedByReporter() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByReporter == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByReporter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByReporter))
	} else if len(*a.IncludedPatientResourcesReferencedByReporter) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByReporter)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByReporter() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedRelatedPersonResourcesReferencedByReporter == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByReporter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedRelatedPersonResourcesReferencedByReporter))
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByReporter) == 1 {
		relatedPerson = &(*a.IncludedRelatedPersonResourcesReferencedByReporter)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatient))
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if a.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *a.RevIncludedListResourcesReferencingItem
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if a.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *a.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if a.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *a.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if a.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *a.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *a.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *a.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if a.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *a.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *a.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingProblem() (clinicalImpressions []ClinicalImpression, err error) {
	if a.RevIncludedClinicalImpressionResourcesReferencingProblem == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *a.RevIncludedClinicalImpressionResourcesReferencingProblem
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if a.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *a.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByRecorder {
			rsc := (*a.IncludedPractitionerResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByRecorder {
			rsc := (*a.IncludedPatientResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedByReporter != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByReporter {
			rsc := (*a.IncludedPractitionerResourcesReferencedByReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByReporter != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByReporter {
			rsc := (*a.IncludedPatientResourcesReferencedByReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByReporter != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByReporter {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *a.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*a.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *a.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*a.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*a.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingProblem != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingProblem {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingProblem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*a.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByRecorder {
			rsc := (*a.IncludedPractitionerResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByRecorder {
			rsc := (*a.IncludedPatientResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedByReporter != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByReporter {
			rsc := (*a.IncludedPractitionerResourcesReferencedByReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByReporter != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByReporter {
			rsc := (*a.IncludedPatientResourcesReferencedByReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByReporter != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByReporter {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *a.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*a.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *a.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*a.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*a.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingProblem != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingProblem {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingProblem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*a.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
