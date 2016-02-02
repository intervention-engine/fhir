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

type ReferralRequest struct {
	DomainResource        `bson:",inline"`
	Status                string            `bson:"status,omitempty" json:"status,omitempty"`
	Identifier            []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Date                  *FHIRDateTime     `bson:"date,omitempty" json:"date,omitempty"`
	Type                  *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Specialty             *CodeableConcept  `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Priority              *CodeableConcept  `bson:"priority,omitempty" json:"priority,omitempty"`
	Patient               *Reference        `bson:"patient,omitempty" json:"patient,omitempty"`
	Requester             *Reference        `bson:"requester,omitempty" json:"requester,omitempty"`
	Recipient             []Reference       `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Encounter             *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateSent              *FHIRDateTime     `bson:"dateSent,omitempty" json:"dateSent,omitempty"`
	Reason                *CodeableConcept  `bson:"reason,omitempty" json:"reason,omitempty"`
	Description           string            `bson:"description,omitempty" json:"description,omitempty"`
	ServiceRequested      []CodeableConcept `bson:"serviceRequested,omitempty" json:"serviceRequested,omitempty"`
	SupportingInformation []Reference       `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	FulfillmentTime       *Period           `bson:"fulfillmentTime,omitempty" json:"fulfillmentTime,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ReferralRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ReferralRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ReferralRequest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ReferralRequest) GetBSON() (interface{}, error) {
	x.ResourceType = "ReferralRequest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "referralRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type referralRequest ReferralRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ReferralRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := referralRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ReferralRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ReferralRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ReferralRequest"
	} else if x.ResourceType != "ReferralRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ReferralRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type ReferralRequestPlus struct {
	ReferralRequest                     `bson:",inline"`
	ReferralRequestPlusRelatedResources `bson:",inline"`
}

type ReferralRequestPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByRequester           *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRequester,omitempty"`
	IncludedOrganizationResourcesReferencedByRequester           *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRequester,omitempty"`
	IncludedPatientResourcesReferencedByRequester                *[]Patient               `bson:"_includedPatientResourcesReferencedByRequester,omitempty"`
	IncludedPatientResourcesReferencedByPatient                  *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByRecipient           *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRecipient,omitempty"`
	IncludedOrganizationResourcesReferencedByRecipient           *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRecipient,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference     *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral *[]EpisodeOfCare         `bson:"_revIncludedEpisodeOfCareResourcesReferencingIncomingreferral,omitempty"`
	RevIncludedListResourcesReferencingItem                      *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref   *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                   *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                  *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingRequest       *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingRequest,omitempty"`
	RevIncludedEncounterResourcesReferencingIncomingreferral     *[]Encounter             `bson:"_revIncludedEncounterResourcesReferencingIncomingreferral,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference           *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject            *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated       *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment      *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject  *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest        *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger     *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAction      *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPlan        *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPlan,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData             *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedPractitionerResourceReferencedByRequester() (practitioner *Practitioner, err error) {
	if r.IncludedPractitionerResourcesReferencedByRequester == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*r.IncludedPractitionerResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*r.IncludedPractitionerResourcesReferencedByRequester))
	} else if len(*r.IncludedPractitionerResourcesReferencedByRequester) == 1 {
		practitioner = &(*r.IncludedPractitionerResourcesReferencedByRequester)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedOrganizationResourceReferencedByRequester() (organization *Organization, err error) {
	if r.IncludedOrganizationResourcesReferencedByRequester == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*r.IncludedOrganizationResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*r.IncludedOrganizationResourcesReferencedByRequester))
	} else if len(*r.IncludedOrganizationResourcesReferencedByRequester) == 1 {
		organization = &(*r.IncludedOrganizationResourcesReferencedByRequester)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByRequester() (patient *Patient, err error) {
	if r.IncludedPatientResourcesReferencedByRequester == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResourcesReferencedByRequester))
	} else if len(*r.IncludedPatientResourcesReferencedByRequester) == 1 {
		patient = &(*r.IncludedPatientResourcesReferencedByRequester)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if r.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResourcesReferencedByPatient))
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*r.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByRecipient() (practitioners []Practitioner, err error) {
	if r.IncludedPractitionerResourcesReferencedByRecipient == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *r.IncludedPractitionerResourcesReferencedByRecipient
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByRecipient() (organizations []Organization, err error) {
	if r.IncludedOrganizationResourcesReferencedByRecipient == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *r.IncludedOrganizationResourcesReferencedByRecipient
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if r.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *r.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingIncomingreferral() (episodeOfCares []EpisodeOfCare, err error) {
	if r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if r.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *r.RevIncludedListResourcesReferencingItem
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if r.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *r.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if r.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *r.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingRequest() (diagnosticReports []DiagnosticReport, err error) {
	if r.RevIncludedDiagnosticReportResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *r.RevIncludedDiagnosticReportResourcesReferencingRequest
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingIncomingreferral() (encounters []Encounter, err error) {
	if r.RevIncludedEncounterResourcesReferencingIncomingreferral == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *r.RevIncludedEncounterResourcesReferencingIncomingreferral
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if r.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *r.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *r.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if r.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *r.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if r.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *r.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if r.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *r.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAction() (clinicalImpressions []ClinicalImpression, err error) {
	if r.RevIncludedClinicalImpressionResourcesReferencingAction == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *r.RevIncludedClinicalImpressionResourcesReferencingAction
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPlan() (clinicalImpressions []ClinicalImpression, err error) {
	if r.RevIncludedClinicalImpressionResourcesReferencingPlan == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *r.RevIncludedClinicalImpressionResourcesReferencingPlan
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if r.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *r.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPractitionerResourcesReferencedByRequester != nil {
		for _, r := range *r.IncludedPractitionerResourcesReferencedByRequester {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedOrganizationResourcesReferencedByRequester != nil {
		for _, r := range *r.IncludedOrganizationResourcesReferencedByRequester {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedPatientResourcesReferencedByRequester != nil {
		for _, r := range *r.IncludedPatientResourcesReferencedByRequester {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *r.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedPractitionerResourcesReferencedByRecipient != nil {
		for _, r := range *r.IncludedPractitionerResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedOrganizationResourcesReferencedByRecipient != nil {
		for _, r := range *r.IncludedOrganizationResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *r.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for _, r := range *r.RevIncludedCarePlanResourcesReferencingActivityreference {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral != nil {
		for _, r := range *r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *r.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *r.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *r.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedDiagnosticReportResourcesReferencingRequest != nil {
		for _, r := range *r.RevIncludedDiagnosticReportResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedEncounterResourcesReferencingIncomingreferral != nil {
		for _, r := range *r.RevIncludedEncounterResourcesReferencingIncomingreferral {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *r.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *r.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *r.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *r.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *r.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *r.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for _, r := range *r.RevIncludedClinicalImpressionResourcesReferencingAction {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for _, r := range *r.RevIncludedClinicalImpressionResourcesReferencingPlan {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *r.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPractitionerResourcesReferencedByRequester != nil {
		for _, r := range *r.IncludedPractitionerResourcesReferencedByRequester {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedOrganizationResourcesReferencedByRequester != nil {
		for _, r := range *r.IncludedOrganizationResourcesReferencedByRequester {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedPatientResourcesReferencedByRequester != nil {
		for _, r := range *r.IncludedPatientResourcesReferencedByRequester {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *r.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedPractitionerResourcesReferencedByRecipient != nil {
		for _, r := range *r.IncludedPractitionerResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if r.IncludedOrganizationResourcesReferencedByRecipient != nil {
		for _, r := range *r.IncludedOrganizationResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *r.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for _, r := range *r.RevIncludedCarePlanResourcesReferencingActivityreference {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral != nil {
		for _, r := range *r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *r.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *r.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *r.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedDiagnosticReportResourcesReferencingRequest != nil {
		for _, r := range *r.RevIncludedDiagnosticReportResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedEncounterResourcesReferencingIncomingreferral != nil {
		for _, r := range *r.RevIncludedEncounterResourcesReferencingIncomingreferral {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *r.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *r.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *r.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *r.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *r.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *r.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for _, r := range *r.RevIncludedClinicalImpressionResourcesReferencingAction {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for _, r := range *r.RevIncludedClinicalImpressionResourcesReferencingPlan {
			resourceMap[r.Id] = &r
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *r.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
