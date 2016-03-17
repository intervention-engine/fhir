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

type EpisodeOfCare struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                                `bson:"status,omitempty" json:"status,omitempty"`
	StatusHistory        []EpisodeOfCareStatusHistoryComponent `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Type                 []CodeableConcept                     `bson:"type,omitempty" json:"type,omitempty"`
	Condition            []Reference                           `bson:"condition,omitempty" json:"condition,omitempty"`
	Patient              *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	ManagingOrganization *Reference                            `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Period               *Period                               `bson:"period,omitempty" json:"period,omitempty"`
	ReferralRequest      []Reference                           `bson:"referralRequest,omitempty" json:"referralRequest,omitempty"`
	CareManager          *Reference                            `bson:"careManager,omitempty" json:"careManager,omitempty"`
	CareTeam             []EpisodeOfCareCareTeamComponent      `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *EpisodeOfCare) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "EpisodeOfCare"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to EpisodeOfCare), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *EpisodeOfCare) GetBSON() (interface{}, error) {
	x.ResourceType = "EpisodeOfCare"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "episodeOfCare" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type episodeOfCare EpisodeOfCare

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *EpisodeOfCare) UnmarshalJSON(data []byte) (err error) {
	x2 := episodeOfCare{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = EpisodeOfCare(x2)
		return x.checkResourceType()
	}
	return
}

func (x *EpisodeOfCare) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "EpisodeOfCare"
	} else if x.ResourceType != "EpisodeOfCare" {
		return errors.New(fmt.Sprintf("Expected resourceType to be EpisodeOfCare, instead received %s", x.ResourceType))
	}
	return nil
}

type EpisodeOfCareStatusHistoryComponent struct {
	BackboneElement `bson:",inline"`
	Status          string  `bson:"status,omitempty" json:"status,omitempty"`
	Period          *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type EpisodeOfCareCareTeamComponent struct {
	BackboneElement `bson:",inline"`
	Role            []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Period          *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Member          *Reference        `bson:"member,omitempty" json:"member,omitempty"`
}

type EpisodeOfCarePlus struct {
	EpisodeOfCare                     `bson:",inline"`
	EpisodeOfCarePlusRelatedResources `bson:",inline"`
}

type EpisodeOfCarePlusRelatedResources struct {
	IncludedConditionResourcesReferencedByCondition              *[]Condition             `bson:"_includedConditionResourcesReferencedByCondition,omitempty"`
	IncludedReferralRequestResourcesReferencedByIncomingreferral *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByIncomingreferral,omitempty"`
	IncludedPatientResourcesReferencedByPatient                  *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization        *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedPractitionerResourcesReferencedByTeammember          *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByTeammember,omitempty"`
	IncludedOrganizationResourcesReferencedByTeammember          *[]Organization          `bson:"_includedOrganizationResourcesReferencedByTeammember,omitempty"`
	IncludedPractitionerResourcesReferencedByCaremanager         *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByCaremanager,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                      *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref   *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                   *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                  *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedEncounterResourcesReferencingEpisodeofcare        *[]Encounter             `bson:"_revIncludedEncounterResourcesReferencingEpisodeofcare,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference           *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject            *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated       *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment      *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject  *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest        *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger     *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData             *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (e *EpisodeOfCarePlusRelatedResources) GetIncludedConditionResourcesReferencedByCondition() (conditions []Condition, err error) {
	if e.IncludedConditionResourcesReferencedByCondition == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *e.IncludedConditionResourcesReferencedByCondition
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetIncludedReferralRequestResourcesReferencedByIncomingreferral() (referralRequests []ReferralRequest, err error) {
	if e.IncludedReferralRequestResourcesReferencedByIncomingreferral == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *e.IncludedReferralRequestResourcesReferencedByIncomingreferral
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedByPatient))
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetIncludedPractitionerResourceReferencedByTeammember() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByTeammember == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByTeammember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByTeammember))
	} else if len(*e.IncludedPractitionerResourcesReferencedByTeammember) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByTeammember)[0]
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetIncludedOrganizationResourceReferencedByTeammember() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByTeammember == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByTeammember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByTeammember))
	} else if len(*e.IncludedOrganizationResourcesReferencedByTeammember) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByTeammember)[0]
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetIncludedPractitionerResourceReferencedByCaremanager() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByCaremanager == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByCaremanager) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByCaremanager))
	} else if len(*e.IncludedPractitionerResourcesReferencedByCaremanager) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByCaremanager)[0]
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingItem
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if e.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *e.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if e.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *e.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedEncounterResourcesReferencingEpisodeofcare() (encounters []Encounter, err error) {
	if e.RevIncludedEncounterResourcesReferencingEpisodeofcare == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *e.RevIncludedEncounterResourcesReferencingEpisodeofcare
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *e.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *e.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if e.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *e.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedConditionResourcesReferencedByCondition != nil {
		for idx := range *e.IncludedConditionResourcesReferencedByCondition {
			rsc := (*e.IncludedConditionResourcesReferencedByCondition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedReferralRequestResourcesReferencedByIncomingreferral != nil {
		for idx := range *e.IncludedReferralRequestResourcesReferencedByIncomingreferral {
			rsc := (*e.IncludedReferralRequestResourcesReferencedByIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatient {
			rsc := (*e.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*e.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByTeammember != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByTeammember {
			rsc := (*e.IncludedPractitionerResourcesReferencedByTeammember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByTeammember != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByTeammember {
			rsc := (*e.IncludedOrganizationResourcesReferencedByTeammember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByCaremanager != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByCaremanager {
			rsc := (*e.IncludedPractitionerResourcesReferencedByCaremanager)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *e.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*e.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEncounterResourcesReferencingEpisodeofcare != nil {
		for idx := range *e.RevIncludedEncounterResourcesReferencingEpisodeofcare {
			rsc := (*e.RevIncludedEncounterResourcesReferencingEpisodeofcare)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*e.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*e.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *e.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*e.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*e.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *e.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*e.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *EpisodeOfCarePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedConditionResourcesReferencedByCondition != nil {
		for idx := range *e.IncludedConditionResourcesReferencedByCondition {
			rsc := (*e.IncludedConditionResourcesReferencedByCondition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedReferralRequestResourcesReferencedByIncomingreferral != nil {
		for idx := range *e.IncludedReferralRequestResourcesReferencedByIncomingreferral {
			rsc := (*e.IncludedReferralRequestResourcesReferencedByIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatient {
			rsc := (*e.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*e.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByTeammember != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByTeammember {
			rsc := (*e.IncludedPractitionerResourcesReferencedByTeammember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByTeammember != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByTeammember {
			rsc := (*e.IncludedOrganizationResourcesReferencedByTeammember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByCaremanager != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByCaremanager {
			rsc := (*e.IncludedPractitionerResourcesReferencedByCaremanager)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *e.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*e.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEncounterResourcesReferencingEpisodeofcare != nil {
		for idx := range *e.RevIncludedEncounterResourcesReferencingEpisodeofcare {
			rsc := (*e.RevIncludedEncounterResourcesReferencingEpisodeofcare)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*e.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*e.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *e.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*e.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*e.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *e.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*e.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
