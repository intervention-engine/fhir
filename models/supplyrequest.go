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

type SupplyRequest struct {
	DomainResource        `bson:",inline"`
	Patient               *Reference                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Source                *Reference                  `bson:"source,omitempty" json:"source,omitempty"`
	Date                  *FHIRDateTime               `bson:"date,omitempty" json:"date,omitempty"`
	Identifier            *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                string                      `bson:"status,omitempty" json:"status,omitempty"`
	Kind                  *CodeableConcept            `bson:"kind,omitempty" json:"kind,omitempty"`
	OrderedItem           *Reference                  `bson:"orderedItem,omitempty" json:"orderedItem,omitempty"`
	Supplier              []Reference                 `bson:"supplier,omitempty" json:"supplier,omitempty"`
	ReasonCodeableConcept *CodeableConcept            `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                  `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	When                  *SupplyRequestWhenComponent `bson:"when,omitempty" json:"when,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *SupplyRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "SupplyRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to SupplyRequest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *SupplyRequest) GetBSON() (interface{}, error) {
	x.ResourceType = "SupplyRequest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "supplyRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type supplyRequest SupplyRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *SupplyRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := supplyRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = SupplyRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *SupplyRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "SupplyRequest"
	} else if x.ResourceType != "SupplyRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be SupplyRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type SupplyRequestWhenComponent struct {
	Code     *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Schedule *Timing          `bson:"schedule,omitempty" json:"schedule,omitempty"`
}

type SupplyRequestPlus struct {
	SupplyRequest                     `bson:",inline"`
	SupplyRequestPlusRelatedResources `bson:",inline"`
}

type SupplyRequestPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedBySupplier           *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySupplier,omitempty"`
	IncludedPractitionerResourcesReferencedBySource             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySource,omitempty"`
	IncludedOrganizationResourcesReferencedBySource             *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySource,omitempty"`
	IncludedPatientResourcesReferencedBySource                  *[]Patient               `bson:"_includedPatientResourcesReferencedBySource,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference    *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
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
	RevIncludedClinicalImpressionResourcesReferencingAction     *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPlan       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPlan,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (s *SupplyRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if s.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResourcesReferencedByPatient))
	} else if len(*s.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*s.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetIncludedOrganizationResourcesReferencedBySupplier() (organizations []Organization, err error) {
	if s.IncludedOrganizationResourcesReferencedBySupplier == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *s.IncludedOrganizationResourcesReferencedBySupplier
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySource() (practitioner *Practitioner, err error) {
	if s.IncludedPractitionerResourcesReferencedBySource == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*s.IncludedPractitionerResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*s.IncludedPractitionerResourcesReferencedBySource))
	} else if len(*s.IncludedPractitionerResourcesReferencedBySource) == 1 {
		practitioner = &(*s.IncludedPractitionerResourcesReferencedBySource)[0]
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySource() (organization *Organization, err error) {
	if s.IncludedOrganizationResourcesReferencedBySource == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*s.IncludedOrganizationResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*s.IncludedOrganizationResourcesReferencedBySource))
	} else if len(*s.IncludedOrganizationResourcesReferencedBySource) == 1 {
		organization = &(*s.IncludedOrganizationResourcesReferencedBySource)[0]
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetIncludedPatientResourceReferencedBySource() (patient *Patient, err error) {
	if s.IncludedPatientResourcesReferencedBySource == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResourcesReferencedBySource))
	} else if len(*s.IncludedPatientResourcesReferencedBySource) == 1 {
		patient = &(*s.IncludedPatientResourcesReferencedBySource)[0]
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if s.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *s.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if s.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *s.RevIncludedListResourcesReferencingItem
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if s.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *s.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if s.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *s.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if s.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *s.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *s.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *s.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if s.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *s.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *s.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAction() (clinicalImpressions []ClinicalImpression, err error) {
	if s.RevIncludedClinicalImpressionResourcesReferencingAction == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *s.RevIncludedClinicalImpressionResourcesReferencingAction
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPlan() (clinicalImpressions []ClinicalImpression, err error) {
	if s.RevIncludedClinicalImpressionResourcesReferencingPlan == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *s.RevIncludedClinicalImpressionResourcesReferencingPlan
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if s.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *s.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (s *SupplyRequestPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *s.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedOrganizationResourcesReferencedBySupplier != nil {
		for _, r := range *s.IncludedOrganizationResourcesReferencedBySupplier {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedPractitionerResourcesReferencedBySource != nil {
		for _, r := range *s.IncludedPractitionerResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedOrganizationResourcesReferencedBySource != nil {
		for _, r := range *s.IncludedOrganizationResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedPatientResourcesReferencedBySource != nil {
		for _, r := range *s.IncludedPatientResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (s *SupplyRequestPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for _, r := range *s.RevIncludedCarePlanResourcesReferencingActivityreference {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *s.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *s.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *s.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *s.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *s.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *s.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for _, r := range *s.RevIncludedClinicalImpressionResourcesReferencingAction {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for _, r := range *s.RevIncludedClinicalImpressionResourcesReferencingPlan {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *s.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (s *SupplyRequestPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *s.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedOrganizationResourcesReferencedBySupplier != nil {
		for _, r := range *s.IncludedOrganizationResourcesReferencedBySupplier {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedPractitionerResourcesReferencedBySource != nil {
		for _, r := range *s.IncludedPractitionerResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedOrganizationResourcesReferencedBySource != nil {
		for _, r := range *s.IncludedOrganizationResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedPatientResourcesReferencedBySource != nil {
		for _, r := range *s.IncludedPatientResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for _, r := range *s.RevIncludedCarePlanResourcesReferencingActivityreference {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *s.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *s.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *s.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *s.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *s.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *s.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for _, r := range *s.RevIncludedClinicalImpressionResourcesReferencingAction {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for _, r := range *s.RevIncludedClinicalImpressionResourcesReferencingPlan {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *s.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
