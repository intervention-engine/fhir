// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
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
	Team                 []Reference                           `bson:"team,omitempty" json:"team,omitempty"`
	Account              []Reference                           `bson:"account,omitempty" json:"account,omitempty"`
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

type EpisodeOfCarePlus struct {
	EpisodeOfCare                     `bson:",inline"`
	EpisodeOfCarePlusRelatedResources `bson:",inline"`
}

type EpisodeOfCarePlusRelatedResources struct {
	IncludedConditionResourcesReferencedByCondition              *[]Condition             `bson:"_includedConditionResourcesReferencedByCondition,omitempty"`
	IncludedReferralRequestResourcesReferencedByIncomingreferral *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByIncomingreferral,omitempty"`
	IncludedPatientResourcesReferencedByPatient                  *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization        *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedPractitionerResourcesReferencedByCaremanager         *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByCaremanager,omitempty"`
	RevIncludedReferralRequestResourcesReferencingContext        *[]ReferralRequest       `bson:"_revIncludedReferralRequestResourcesReferencingContext,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                   *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref   *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject               *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest          *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource   *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingContext           *[]ImagingStudy          `bson:"_revIncludedImagingStudyResourcesReferencingContext,omitempty"`
	RevIncludedEncounterResourcesReferencingEpisodeofcare        *[]Encounter             `bson:"_revIncludedEncounterResourcesReferencingEpisodeofcare,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon          *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingContext          *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingContext,omitempty"`
	RevIncludedRequestGroupResourcesReferencingContext           *[]RequestGroup          `bson:"_revIncludedRequestGroupResourcesReferencingContext,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData             *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                   *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                     *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                   *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedTaskResourcesReferencingContext                   *[]Task                  `bson:"_revIncludedTaskResourcesReferencingContext,omitempty"`
	RevIncludedListResourcesReferencingItem                      *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingEncounter    *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon      *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition   *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingContext      *[]MedicationRequest     `bson:"_revIncludedMedicationRequestResourcesReferencingContext,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingEncounter     *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon       *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition    *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingContext   *[]CommunicationRequest  `bson:"_revIncludedCommunicationRequestResourcesReferencingContext,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                  *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity              *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingContext              *[]Condition             `bson:"_revIncludedConditionResourcesReferencingContext,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject            *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated       *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject  *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingContext  *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingContext,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest        *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingContext     *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingContext,omitempty"`
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

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingContext() (referralRequests []ReferralRequest, err error) {
	if e.RevIncludedReferralRequestResourcesReferencingContext == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *e.RevIncludedReferralRequestResourcesReferencingContext
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

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if e.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *e.RevIncludedConsentResourcesReferencingData
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

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if e.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *e.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingContext() (imagingStudies []ImagingStudy, err error) {
	if e.RevIncludedImagingStudyResourcesReferencingContext == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *e.RevIncludedImagingStudyResourcesReferencingContext
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

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingContext() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingContext == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingContext
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingContext() (requestGroups []RequestGroup, err error) {
	if e.RevIncludedRequestGroupResourcesReferencingContext == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *e.RevIncludedRequestGroupResourcesReferencingContext
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

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingEntity
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

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedTaskResourcesReferencingContext() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingContext == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingContext
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

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if e.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *e.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingEncounter() (diagnosticRequests []DiagnosticRequest, err error) {
	if e.RevIncludedDiagnosticRequestResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *e.RevIncludedDiagnosticRequestResourcesReferencingEncounter
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if e.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *e.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if e.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *e.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingContext() (medicationRequests []MedicationRequest, err error) {
	if e.RevIncludedMedicationRequestResourcesReferencingContext == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *e.RevIncludedMedicationRequestResourcesReferencingContext
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if e.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *e.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingEncounter() (deviceUseRequests []DeviceUseRequest, err error) {
	if e.RevIncludedDeviceUseRequestResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *e.RevIncludedDeviceUseRequestResourcesReferencingEncounter
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if e.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *e.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if e.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *e.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingContext() (communicationRequests []CommunicationRequest, err error) {
	if e.RevIncludedCommunicationRequestResourcesReferencingContext == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *e.RevIncludedCommunicationRequestResourcesReferencingContext
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

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedConditionResourcesReferencingContext() (conditions []Condition, err error) {
	if e.RevIncludedConditionResourcesReferencingContext == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *e.RevIncludedConditionResourcesReferencingContext
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

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingContext() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingContext == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingContext
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

func (e *EpisodeOfCarePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingContext() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingContext == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingContext
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
	if e.RevIncludedReferralRequestResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedReferralRequestResourcesReferencingContext {
			rsc := (*e.RevIncludedReferralRequestResourcesReferencingContext)[idx]
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
	if e.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingData {
			rsc := (*e.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImagingStudyResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedImagingStudyResourcesReferencingContext {
			rsc := (*e.RevIncludedImagingStudyResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEncounterResourcesReferencingEpisodeofcare != nil {
		for idx := range *e.RevIncludedEncounterResourcesReferencingEpisodeofcare {
			rsc := (*e.RevIncludedEncounterResourcesReferencingEpisodeofcare)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingContext {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingContext {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*e.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*e.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingContext {
			rsc := (*e.RevIncludedTaskResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationRequestResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationRequestResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationRequestResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingContext {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingContext {
			rsc := (*e.RevIncludedConditionResourcesReferencingContext)[idx]
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
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingContext {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*e.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedClinicalImpressionResourcesReferencingContext {
			rsc := (*e.RevIncludedClinicalImpressionResourcesReferencingContext)[idx]
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
	if e.IncludedPractitionerResourcesReferencedByCaremanager != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByCaremanager {
			rsc := (*e.IncludedPractitionerResourcesReferencedByCaremanager)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedReferralRequestResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedReferralRequestResourcesReferencingContext {
			rsc := (*e.RevIncludedReferralRequestResourcesReferencingContext)[idx]
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
	if e.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingData {
			rsc := (*e.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImagingStudyResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedImagingStudyResourcesReferencingContext {
			rsc := (*e.RevIncludedImagingStudyResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEncounterResourcesReferencingEpisodeofcare != nil {
		for idx := range *e.RevIncludedEncounterResourcesReferencingEpisodeofcare {
			rsc := (*e.RevIncludedEncounterResourcesReferencingEpisodeofcare)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingContext {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingContext {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*e.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*e.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingContext {
			rsc := (*e.RevIncludedTaskResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationRequestResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationRequestResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationRequestResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingContext {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingContext {
			rsc := (*e.RevIncludedConditionResourcesReferencingContext)[idx]
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
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingContext {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*e.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedClinicalImpressionResourcesReferencingContext {
			rsc := (*e.RevIncludedClinicalImpressionResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
