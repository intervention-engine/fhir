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

type Organization struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active         *bool                          `bson:"active,omitempty" json:"active,omitempty"`
	Type           *CodeableConcept               `bson:"type,omitempty" json:"type,omitempty"`
	Name           string                         `bson:"name,omitempty" json:"name,omitempty"`
	Alias          []string                       `bson:"alias,omitempty" json:"alias,omitempty"`
	Telecom        []ContactPoint                 `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address        []Address                      `bson:"address,omitempty" json:"address,omitempty"`
	PartOf         *Reference                     `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Contact        []OrganizationContactComponent `bson:"contact,omitempty" json:"contact,omitempty"`
	Endpoint       []Reference                    `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Organization) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Organization"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Organization), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Organization) GetBSON() (interface{}, error) {
	x.ResourceType = "Organization"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "organization" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type organization Organization

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Organization) UnmarshalJSON(data []byte) (err error) {
	x2 := organization{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Organization(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Organization) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Organization"
	} else if x.ResourceType != "Organization" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Organization, instead received %s", x.ResourceType))
	}
	return nil
}

type OrganizationContactComponent struct {
	BackboneElement `bson:",inline"`
	Purpose         *CodeableConcept `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Name            *HumanName       `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint   `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address         *Address         `bson:"address,omitempty" json:"address,omitempty"`
}

type OrganizationPlus struct {
	Organization                     `bson:",inline"`
	OrganizationPlusRelatedResources `bson:",inline"`
}

type OrganizationPlusRelatedResources struct {
	IncludedOrganizationResourcesReferencedByPartof                         *[]Organization          `bson:"_includedOrganizationResourcesReferencedByPartof,omitempty"`
	IncludedEndpointResourcesReferencedByEndpoint                           *[]Endpoint              `bson:"_includedEndpointResourcesReferencedByEndpoint,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRequester                 *[]ReferralRequest       `bson:"_revIncludedReferralRequestResourcesReferencingRequester,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRecipient                 *[]ReferralRequest       `bson:"_revIncludedReferralRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedAccountResourcesReferencingOwner                             *[]Account               `bson:"_revIncludedAccountResourcesReferencingOwner,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                           *[]Account               `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref               *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor                   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref               *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRecipient                *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRecipient,omitempty"`
	RevIncludedGoalResourcesReferencingSubject                              *[]Goal                  `bson:"_revIncludedGoalResourcesReferencingSubject,omitempty"`
	RevIncludedEndpointResourcesReferencingOrganization                     *[]Endpoint              `bson:"_revIncludedEndpointResourcesReferencingOrganization,omitempty"`
	RevIncludedEnrollmentRequestResourcesReferencingOrganization            *[]EnrollmentRequest     `bson:"_revIncludedEnrollmentRequestResourcesReferencingOrganization,omitempty"`
	RevIncludedConsentResourcesReferencingData                              *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedConsentResourcesReferencingActor                             *[]Consent               `bson:"_revIncludedConsentResourcesReferencingActor,omitempty"`
	RevIncludedConsentResourcesReferencingOrganization                      *[]Consent               `bson:"_revIncludedConsentResourcesReferencingOrganization,omitempty"`
	RevIncludedConsentResourcesReferencingRecipient                         *[]Consent               `bson:"_revIncludedConsentResourcesReferencingRecipient,omitempty"`
	RevIncludedConsentResourcesReferencingConsentor                         *[]Consent               `bson:"_revIncludedConsentResourcesReferencingConsentor,omitempty"`
	RevIncludedMedicationResourcesReferencingManufacturer                   *[]Medication            `bson:"_revIncludedMedicationResourcesReferencingManufacturer,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthenticator           *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthenticator,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingCustodian               *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingCustodian,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor                  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref              *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedImagingManifestResourcesReferencingAuthor                    *[]ImagingManifest       `bson:"_revIncludedImagingManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedPractitionerRoleResourcesReferencingOrganization             *[]PractitionerRole      `bson:"_revIncludedPractitionerRoleResourcesReferencingOrganization,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingSupplier                    *[]SupplyRequest         `bson:"_revIncludedSupplyRequestResourcesReferencingSupplier,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingSource                      *[]SupplyRequest         `bson:"_revIncludedSupplyRequestResourcesReferencingSource,omitempty"`
	RevIncludedPractitionerResourcesReferencingOrganization                 *[]Practitioner          `bson:"_revIncludedPractitionerResourcesReferencingOrganization,omitempty"`
	RevIncludedPersonResourcesReferencingOrganization                       *[]Person                `bson:"_revIncludedPersonResourcesReferencingOrganization,omitempty"`
	RevIncludedContractResourcesReferencingAgent                            *[]Contract              `bson:"_revIncludedContractResourcesReferencingAgent,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                           *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                          *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingAuthority                        *[]Contract              `bson:"_revIncludedContractResourcesReferencingAuthority,omitempty"`
	RevIncludedContractResourcesReferencingTopic                            *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedContractResourcesReferencingSigner                           *[]Contract              `bson:"_revIncludedContractResourcesReferencingSigner,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                     *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                    *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingOrganization                *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingOrganization,omitempty"`
	RevIncludedOrganizationResourcesReferencingPartof                       *[]Organization          `bson:"_revIncludedOrganizationResourcesReferencingPartof,omitempty"`
	RevIncludedCareTeamResourcesReferencingParticipant                      *[]CareTeam              `bson:"_revIncludedCareTeamResourcesReferencingParticipant,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource              *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedEncounterResourcesReferencingServiceprovider                 *[]Encounter             `bson:"_revIncludedEncounterResourcesReferencingServiceprovider,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                     *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender                      *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient                   *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedLinkageResourcesReferencingAuthor                            *[]Linkage               `bson:"_revIncludedLinkageResourcesReferencingAuthor,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                        *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingReceiver                    *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingReceiver,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingResponsible                 *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingResponsible,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                          *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                         *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                         *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingOwner                                *[]Task                  `bson:"_revIncludedTaskResourcesReferencingOwner,omitempty"`
	RevIncludedTaskResourcesReferencingRequester                            *[]Task                  `bson:"_revIncludedTaskResourcesReferencingRequester,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                              *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                                *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                              *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedTaskResourcesReferencingOrganization                         *[]Task                  `bson:"_revIncludedTaskResourcesReferencingOrganization,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingOrganization         *[]ExplanationOfBenefit  `bson:"_revIncludedExplanationOfBenefitResourcesReferencingOrganization,omitempty"`
	RevIncludedResearchStudyResourcesReferencingSponsor                     *[]ResearchStudy         `bson:"_revIncludedResearchStudyResourcesReferencingSponsor,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                        *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingOrganization                *[]EpisodeOfCare         `bson:"_revIncludedEpisodeOfCareResourcesReferencingOrganization,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                       *[]Procedure             `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedListResourcesReferencingItem                                 *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingRequester               *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingRequester,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces                *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon                 *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingFiller                  *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingFiller,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition              *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedImmunizationResourcesReferencingManufacturer                 *[]Immunization          `bson:"_revIncludedImmunizationResourcesReferencingManufacturer,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingRequester               *[]MedicationRequest     `bson:"_revIncludedMedicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingIntendeddispenser       *[]MedicationRequest     `bson:"_revIncludedMedicationRequestResourcesReferencingIntendeddispenser,omitempty"`
	RevIncludedDeviceResourcesReferencingOrganization                       *[]Device                `bson:"_revIncludedDeviceResourcesReferencingOrganization,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingPerformer                *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedEligibilityResponseResourcesReferencingInsurer               *[]EligibilityResponse   `bson:"_revIncludedEligibilityResponseResourcesReferencingInsurer,omitempty"`
	RevIncludedEligibilityResponseResourcesReferencingRequestorganization   *[]EligibilityResponse   `bson:"_revIncludedEligibilityResponseResourcesReferencingRequestorganization,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingRequester                *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingRequester,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces                 *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon                  *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingFiller                   *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingFiller,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition               *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                              *[]Flag                  `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                               *[]Flag                  `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedObservationResourcesReferencingPerformer                     *[]Observation           `bson:"_revIncludedObservationResourcesReferencingPerformer,omitempty"`
	RevIncludedEnrollmentResponseResourcesReferencingOrganization           *[]EnrollmentResponse    `bson:"_revIncludedEnrollmentResponseResourcesReferencingOrganization,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSource                *[]MedicationStatement   `bson:"_revIncludedMedicationStatementResourcesReferencingSource,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender               *[]CommunicationRequest  `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient            *[]CommunicationRequest  `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                             *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedClaimResponseResourcesReferencingInsurer                     *[]ClaimResponse         `bson:"_revIncludedClaimResponseResourcesReferencingInsurer,omitempty"`
	RevIncludedEligibilityRequestResourcesReferencingOrganization           *[]EligibilityRequest    `bson:"_revIncludedEligibilityRequestResourcesReferencingOrganization,omitempty"`
	RevIncludedProcessRequestResourcesReferencingOrganization               *[]ProcessRequest        `bson:"_revIncludedProcessRequestResourcesReferencingOrganization,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingPerformer                *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingPerformer,omitempty"`
	RevIncludedHealthcareServiceResourcesReferencingOrganization            *[]HealthcareService     `bson:"_revIncludedHealthcareServiceResourcesReferencingOrganization,omitempty"`
	RevIncludedAuditEventResourcesReferencingAgent                          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingAgent,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                         *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedPaymentReconciliationResourcesReferencingOrganization        *[]PaymentReconciliation `bson:"_revIncludedPaymentReconciliationResourcesReferencingOrganization,omitempty"`
	RevIncludedPaymentReconciliationResourcesReferencingRequestorganization *[]PaymentReconciliation `bson:"_revIncludedPaymentReconciliationResourcesReferencingRequestorganization,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                       *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAttester                      *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingAttester,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                         *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                  *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedPatientResourcesReferencingGeneralpractitioner               *[]Patient               `bson:"_revIncludedPatientResourcesReferencingGeneralpractitioner,omitempty"`
	RevIncludedPatientResourcesReferencingOrganization                      *[]Patient               `bson:"_revIncludedPatientResourcesReferencingOrganization,omitempty"`
	RevIncludedCoverageResourcesReferencingPolicyholder                     *[]Coverage              `bson:"_revIncludedCoverageResourcesReferencingPolicyholder,omitempty"`
	RevIncludedCoverageResourcesReferencingPayor                            *[]Coverage              `bson:"_revIncludedCoverageResourcesReferencingPayor,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject             *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest                   *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedProcessResponseResourcesReferencingOrganization              *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingOrganization,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestorganization       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequestorganization,omitempty"`
	RevIncludedClaimResourcesReferencingInsurer                             *[]Claim                 `bson:"_revIncludedClaimResourcesReferencingInsurer,omitempty"`
	RevIncludedClaimResourcesReferencingOrganization                        *[]Claim                 `bson:"_revIncludedClaimResourcesReferencingOrganization,omitempty"`
	RevIncludedLocationResourcesReferencingOrganization                     *[]Location              `bson:"_revIncludedLocationResourcesReferencingOrganization,omitempty"`
}

func (o *OrganizationPlusRelatedResources) GetIncludedOrganizationResourceReferencedByPartof() (organization *Organization, err error) {
	if o.IncludedOrganizationResourcesReferencedByPartof == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*o.IncludedOrganizationResourcesReferencedByPartof) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*o.IncludedOrganizationResourcesReferencedByPartof))
	} else if len(*o.IncludedOrganizationResourcesReferencedByPartof) == 1 {
		organization = &(*o.IncludedOrganizationResourcesReferencedByPartof)[0]
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetIncludedEndpointResourcesReferencedByEndpoint() (endpoints []Endpoint, err error) {
	if o.IncludedEndpointResourcesReferencedByEndpoint == nil {
		err = errors.New("Included endpoints not requested")
	} else {
		endpoints = *o.IncludedEndpointResourcesReferencedByEndpoint
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingRequester() (referralRequests []ReferralRequest, err error) {
	if o.RevIncludedReferralRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *o.RevIncludedReferralRequestResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingRecipient() (referralRequests []ReferralRequest, err error) {
	if o.RevIncludedReferralRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *o.RevIncludedReferralRequestResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAccountResourcesReferencingOwner() (accounts []Account, err error) {
	if o.RevIncludedAccountResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *o.RevIncludedAccountResourcesReferencingOwner
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if o.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *o.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRecipient() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedGoalResourcesReferencingSubject() (goals []Goal, err error) {
	if o.RevIncludedGoalResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded goals not requested")
	} else {
		goals = *o.RevIncludedGoalResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEndpointResourcesReferencingOrganization() (endpoints []Endpoint, err error) {
	if o.RevIncludedEndpointResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded endpoints not requested")
	} else {
		endpoints = *o.RevIncludedEndpointResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEnrollmentRequestResourcesReferencingOrganization() (enrollmentRequests []EnrollmentRequest, err error) {
	if o.RevIncludedEnrollmentRequestResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded enrollmentRequests not requested")
	} else {
		enrollmentRequests = *o.RevIncludedEnrollmentRequestResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingActor() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingActor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingOrganization() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingRecipient() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingConsentor() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingConsentor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingConsentor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationResourcesReferencingManufacturer() (medications []Medication, err error) {
	if o.RevIncludedMedicationResourcesReferencingManufacturer == nil {
		err = errors.New("RevIncluded medications not requested")
	} else {
		medications = *o.RevIncludedMedicationResourcesReferencingManufacturer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthenticator() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingCustodian() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingCustodian == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingCustodian
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedImagingManifestResourcesReferencingAuthor() (imagingManifests []ImagingManifest, err error) {
	if o.RevIncludedImagingManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded imagingManifests not requested")
	} else {
		imagingManifests = *o.RevIncludedImagingManifestResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPractitionerRoleResourcesReferencingOrganization() (practitionerRoles []PractitionerRole, err error) {
	if o.RevIncludedPractitionerRoleResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded practitionerRoles not requested")
	} else {
		practitionerRoles = *o.RevIncludedPractitionerRoleResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingSupplier() (supplyRequests []SupplyRequest, err error) {
	if o.RevIncludedSupplyRequestResourcesReferencingSupplier == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *o.RevIncludedSupplyRequestResourcesReferencingSupplier
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingSource() (supplyRequests []SupplyRequest, err error) {
	if o.RevIncludedSupplyRequestResourcesReferencingSource == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *o.RevIncludedSupplyRequestResourcesReferencingSource
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPractitionerResourcesReferencingOrganization() (practitioners []Practitioner, err error) {
	if o.RevIncludedPractitionerResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded practitioners not requested")
	} else {
		practitioners = *o.RevIncludedPractitionerResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPersonResourcesReferencingOrganization() (people []Person, err error) {
	if o.RevIncludedPersonResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *o.RevIncludedPersonResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingAgent() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingAgent
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingAuthority() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingAuthority == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingAuthority
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSigner() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingSigner == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingSigner
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingOrganization() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedOrganizationResourcesReferencingPartof() (organizations []Organization, err error) {
	if o.RevIncludedOrganizationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded organizations not requested")
	} else {
		organizations = *o.RevIncludedOrganizationResourcesReferencingPartof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingParticipant() (careTeams []CareTeam, err error) {
	if o.RevIncludedCareTeamResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *o.RevIncludedCareTeamResourcesReferencingParticipant
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if o.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *o.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingServiceprovider() (encounters []Encounter, err error) {
	if o.RevIncludedEncounterResourcesReferencingServiceprovider == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *o.RevIncludedEncounterResourcesReferencingServiceprovider
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingAuthor() (linkages []Linkage, err error) {
	if o.RevIncludedLinkageResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *o.RevIncludedLinkageResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingReceiver() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingReceiver
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingResponsible() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingResponsible == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingResponsible
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingOwner() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingOwner
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingRequester() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingOrganization() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingOrganization() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if o.RevIncludedExplanationOfBenefitResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *o.RevIncludedExplanationOfBenefitResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchStudyResourcesReferencingSponsor() (researchStudies []ResearchStudy, err error) {
	if o.RevIncludedResearchStudyResourcesReferencingSponsor == nil {
		err = errors.New("RevIncluded researchStudies not requested")
	} else {
		researchStudies = *o.RevIncludedResearchStudyResourcesReferencingSponsor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if o.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *o.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingOrganization() (episodeOfCares []EpisodeOfCare, err error) {
	if o.RevIncludedEpisodeOfCareResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *o.RevIncludedEpisodeOfCareResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if o.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *o.RevIncludedProcedureResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if o.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *o.RevIncludedListResourcesReferencingItem
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingRequester() (diagnosticRequests []DiagnosticRequest, err error) {
	if o.RevIncludedDiagnosticRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *o.RevIncludedDiagnosticRequestResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if o.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *o.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if o.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *o.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingFiller() (diagnosticRequests []DiagnosticRequest, err error) {
	if o.RevIncludedDiagnosticRequestResourcesReferencingFiller == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *o.RevIncludedDiagnosticRequestResourcesReferencingFiller
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if o.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *o.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingManufacturer() (immunizations []Immunization, err error) {
	if o.RevIncludedImmunizationResourcesReferencingManufacturer == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *o.RevIncludedImmunizationResourcesReferencingManufacturer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingRequester() (medicationRequests []MedicationRequest, err error) {
	if o.RevIncludedMedicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *o.RevIncludedMedicationRequestResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingIntendeddispenser() (medicationRequests []MedicationRequest, err error) {
	if o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceResourcesReferencingOrganization() (devices []Device, err error) {
	if o.RevIncludedDeviceResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded devices not requested")
	} else {
		devices = *o.RevIncludedDeviceResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingPerformer() (procedureRequests []ProcedureRequest, err error) {
	if o.RevIncludedProcedureRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *o.RevIncludedProcedureRequestResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEligibilityResponseResourcesReferencingInsurer() (eligibilityResponses []EligibilityResponse, err error) {
	if o.RevIncludedEligibilityResponseResourcesReferencingInsurer == nil {
		err = errors.New("RevIncluded eligibilityResponses not requested")
	} else {
		eligibilityResponses = *o.RevIncludedEligibilityResponseResourcesReferencingInsurer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEligibilityResponseResourcesReferencingRequestorganization() (eligibilityResponses []EligibilityResponse, err error) {
	if o.RevIncludedEligibilityResponseResourcesReferencingRequestorganization == nil {
		err = errors.New("RevIncluded eligibilityResponses not requested")
	} else {
		eligibilityResponses = *o.RevIncludedEligibilityResponseResourcesReferencingRequestorganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingRequester() (deviceUseRequests []DeviceUseRequest, err error) {
	if o.RevIncludedDeviceUseRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *o.RevIncludedDeviceUseRequestResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if o.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *o.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if o.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *o.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingFiller() (deviceUseRequests []DeviceUseRequest, err error) {
	if o.RevIncludedDeviceUseRequestResourcesReferencingFiller == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *o.RevIncludedDeviceUseRequestResourcesReferencingFiller
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if o.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *o.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if o.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *o.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedFlagResourcesReferencingAuthor() (flags []Flag, err error) {
	if o.RevIncludedFlagResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *o.RevIncludedFlagResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPerformer() (observations []Observation, err error) {
	if o.RevIncludedObservationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *o.RevIncludedObservationResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEnrollmentResponseResourcesReferencingOrganization() (enrollmentResponses []EnrollmentResponse, err error) {
	if o.RevIncludedEnrollmentResponseResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded enrollmentResponses not requested")
	} else {
		enrollmentResponses = *o.RevIncludedEnrollmentResponseResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSource() (medicationStatements []MedicationStatement, err error) {
	if o.RevIncludedMedicationStatementResourcesReferencingSource == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *o.RevIncludedMedicationStatementResourcesReferencingSource
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if o.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *o.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if o.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *o.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedClaimResponseResourcesReferencingInsurer() (claimResponses []ClaimResponse, err error) {
	if o.RevIncludedClaimResponseResourcesReferencingInsurer == nil {
		err = errors.New("RevIncluded claimResponses not requested")
	} else {
		claimResponses = *o.RevIncludedClaimResponseResourcesReferencingInsurer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEligibilityRequestResourcesReferencingOrganization() (eligibilityRequests []EligibilityRequest, err error) {
	if o.RevIncludedEligibilityRequestResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded eligibilityRequests not requested")
	} else {
		eligibilityRequests = *o.RevIncludedEligibilityRequestResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcessRequestResourcesReferencingOrganization() (processRequests []ProcessRequest, err error) {
	if o.RevIncludedProcessRequestResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded processRequests not requested")
	} else {
		processRequests = *o.RevIncludedProcessRequestResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingPerformer() (diagnosticReports []DiagnosticReport, err error) {
	if o.RevIncludedDiagnosticReportResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *o.RevIncludedDiagnosticReportResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedHealthcareServiceResourcesReferencingOrganization() (healthcareServices []HealthcareService, err error) {
	if o.RevIncludedHealthcareServiceResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded healthcareServices not requested")
	} else {
		healthcareServices = *o.RevIncludedHealthcareServiceResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingAgent() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingAgent
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPaymentReconciliationResourcesReferencingOrganization() (paymentReconciliations []PaymentReconciliation, err error) {
	if o.RevIncludedPaymentReconciliationResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded paymentReconciliations not requested")
	} else {
		paymentReconciliations = *o.RevIncludedPaymentReconciliationResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPaymentReconciliationResourcesReferencingRequestorganization() (paymentReconciliations []PaymentReconciliation, err error) {
	if o.RevIncludedPaymentReconciliationResourcesReferencingRequestorganization == nil {
		err = errors.New("RevIncluded paymentReconciliations not requested")
	} else {
		paymentReconciliations = *o.RevIncludedPaymentReconciliationResourcesReferencingRequestorganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAttester() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingAttester == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingAttester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *o.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPatientResourcesReferencingGeneralpractitioner() (patients []Patient, err error) {
	if o.RevIncludedPatientResourcesReferencingGeneralpractitioner == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *o.RevIncludedPatientResourcesReferencingGeneralpractitioner
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPatientResourcesReferencingOrganization() (patients []Patient, err error) {
	if o.RevIncludedPatientResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *o.RevIncludedPatientResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingPolicyholder() (coverages []Coverage, err error) {
	if o.RevIncludedCoverageResourcesReferencingPolicyholder == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *o.RevIncludedCoverageResourcesReferencingPolicyholder
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingPayor() (coverages []Coverage, err error) {
	if o.RevIncludedCoverageResourcesReferencingPayor == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *o.RevIncludedCoverageResourcesReferencingPayor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingOrganization() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestorganization() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingRequestorganization == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingRequestorganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedClaimResourcesReferencingInsurer() (claims []Claim, err error) {
	if o.RevIncludedClaimResourcesReferencingInsurer == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *o.RevIncludedClaimResourcesReferencingInsurer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedClaimResourcesReferencingOrganization() (claims []Claim, err error) {
	if o.RevIncludedClaimResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *o.RevIncludedClaimResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLocationResourcesReferencingOrganization() (locations []Location, err error) {
	if o.RevIncludedLocationResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded locations not requested")
	} else {
		locations = *o.RevIncludedLocationResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedOrganizationResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByPartof {
			rsc := (*o.IncludedOrganizationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *o.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*o.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedReferralRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedReferralRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedReferralRequestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedReferralRequestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedReferralRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingOwner != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingOwner {
			rsc := (*o.RevIncludedAccountResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*o.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*o.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEndpointResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEndpointResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEndpointResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEnrollmentRequestResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEnrollmentRequestResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEnrollmentRequestResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingData {
			rsc := (*o.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingActor {
			rsc := (*o.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingOrganization {
			rsc := (*o.RevIncludedConsentResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingRecipient {
			rsc := (*o.RevIncludedConsentResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*o.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedMedicationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedMedicationResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingCustodian != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingCustodian {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingCustodian)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImagingManifestResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedImagingManifestResourcesReferencingAuthor {
			rsc := (*o.RevIncludedImagingManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPractitionerRoleResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPractitionerRoleResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPractitionerRoleResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSupplier != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSupplier {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSource {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPractitionerResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPractitionerResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPractitionerResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPersonResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPersonResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPersonResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingAgent {
			rsc := (*o.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*o.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingAuthority != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingAuthority {
			rsc := (*o.RevIncludedContractResourcesReferencingAuthority)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingTopic {
			rsc := (*o.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSigner {
			rsc := (*o.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrganizationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedOrganizationResourcesReferencingPartof {
			rsc := (*o.RevIncludedOrganizationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *o.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*o.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEncounterResourcesReferencingServiceprovider != nil {
		for idx := range *o.RevIncludedEncounterResourcesReferencingServiceprovider {
			rsc := (*o.RevIncludedEncounterResourcesReferencingServiceprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingAuthor {
			rsc := (*o.RevIncludedLinkageResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*o.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*o.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*o.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*o.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*o.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingOrganization {
			rsc := (*o.RevIncludedTaskResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedExplanationOfBenefitResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedExplanationOfBenefitResourcesReferencingOrganization {
			rsc := (*o.RevIncludedExplanationOfBenefitResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchStudyResourcesReferencingSponsor != nil {
		for idx := range *o.RevIncludedResearchStudyResourcesReferencingSponsor {
			rsc := (*o.RevIncludedResearchStudyResourcesReferencingSponsor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*o.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEpisodeOfCareResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEpisodeOfCareResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEpisodeOfCareResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*o.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingFiller != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingFiller {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser != nil {
		for idx := range *o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser {
			rsc := (*o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedDeviceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedDeviceResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*o.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEligibilityResponseResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedEligibilityResponseResourcesReferencingInsurer {
			rsc := (*o.RevIncludedEligibilityResponseResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEligibilityResponseResourcesReferencingRequestorganization != nil {
		for idx := range *o.RevIncludedEligibilityResponseResourcesReferencingRequestorganization {
			rsc := (*o.RevIncludedEligibilityResponseResourcesReferencingRequestorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingFiller != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingFiller {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*o.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*o.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*o.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEnrollmentResponseResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEnrollmentResponseResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEnrollmentResponseResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*o.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResponseResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedClaimResponseResourcesReferencingInsurer {
			rsc := (*o.RevIncludedClaimResponseResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEligibilityRequestResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEligibilityRequestResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEligibilityRequestResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessRequestResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedProcessRequestResourcesReferencingOrganization {
			rsc := (*o.RevIncludedProcessRequestResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedHealthcareServiceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedHealthcareServiceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedHealthcareServiceResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentReconciliationResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPaymentReconciliationResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPaymentReconciliationResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentReconciliationResourcesReferencingRequestorganization != nil {
		for idx := range *o.RevIncludedPaymentReconciliationResourcesReferencingRequestorganization {
			rsc := (*o.RevIncludedPaymentReconciliationResourcesReferencingRequestorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*o.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingGeneralpractitioner != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingGeneralpractitioner {
			rsc := (*o.RevIncludedPatientResourcesReferencingGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPatientResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageResourcesReferencingPolicyholder != nil {
		for idx := range *o.RevIncludedCoverageResourcesReferencingPolicyholder {
			rsc := (*o.RevIncludedCoverageResourcesReferencingPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageResourcesReferencingPayor != nil {
		for idx := range *o.RevIncludedCoverageResourcesReferencingPayor {
			rsc := (*o.RevIncludedCoverageResourcesReferencingPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingOrganization {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequestorganization != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequestorganization {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequestorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingInsurer {
			rsc := (*o.RevIncludedClaimResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingOrganization {
			rsc := (*o.RevIncludedClaimResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLocationResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedLocationResourcesReferencingOrganization {
			rsc := (*o.RevIncludedLocationResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *OrganizationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedOrganizationResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByPartof {
			rsc := (*o.IncludedOrganizationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *o.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*o.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedReferralRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedReferralRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedReferralRequestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedReferralRequestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedReferralRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingOwner != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingOwner {
			rsc := (*o.RevIncludedAccountResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*o.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*o.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEndpointResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEndpointResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEndpointResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEnrollmentRequestResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEnrollmentRequestResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEnrollmentRequestResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingData {
			rsc := (*o.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingActor {
			rsc := (*o.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingOrganization {
			rsc := (*o.RevIncludedConsentResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingRecipient {
			rsc := (*o.RevIncludedConsentResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*o.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedMedicationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedMedicationResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingCustodian != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingCustodian {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingCustodian)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImagingManifestResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedImagingManifestResourcesReferencingAuthor {
			rsc := (*o.RevIncludedImagingManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPractitionerRoleResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPractitionerRoleResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPractitionerRoleResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSupplier != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSupplier {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSource {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPractitionerResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPractitionerResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPractitionerResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPersonResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPersonResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPersonResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingAgent {
			rsc := (*o.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*o.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingAuthority != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingAuthority {
			rsc := (*o.RevIncludedContractResourcesReferencingAuthority)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingTopic {
			rsc := (*o.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSigner {
			rsc := (*o.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrganizationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedOrganizationResourcesReferencingPartof {
			rsc := (*o.RevIncludedOrganizationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *o.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*o.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEncounterResourcesReferencingServiceprovider != nil {
		for idx := range *o.RevIncludedEncounterResourcesReferencingServiceprovider {
			rsc := (*o.RevIncludedEncounterResourcesReferencingServiceprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingAuthor {
			rsc := (*o.RevIncludedLinkageResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*o.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*o.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*o.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*o.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*o.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingOrganization {
			rsc := (*o.RevIncludedTaskResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedExplanationOfBenefitResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedExplanationOfBenefitResourcesReferencingOrganization {
			rsc := (*o.RevIncludedExplanationOfBenefitResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchStudyResourcesReferencingSponsor != nil {
		for idx := range *o.RevIncludedResearchStudyResourcesReferencingSponsor {
			rsc := (*o.RevIncludedResearchStudyResourcesReferencingSponsor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*o.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEpisodeOfCareResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEpisodeOfCareResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEpisodeOfCareResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*o.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingFiller != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingFiller {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser != nil {
		for idx := range *o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser {
			rsc := (*o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedDeviceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedDeviceResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*o.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEligibilityResponseResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedEligibilityResponseResourcesReferencingInsurer {
			rsc := (*o.RevIncludedEligibilityResponseResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEligibilityResponseResourcesReferencingRequestorganization != nil {
		for idx := range *o.RevIncludedEligibilityResponseResourcesReferencingRequestorganization {
			rsc := (*o.RevIncludedEligibilityResponseResourcesReferencingRequestorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingFiller != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingFiller {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingFiller)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*o.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*o.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*o.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEnrollmentResponseResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEnrollmentResponseResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEnrollmentResponseResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*o.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResponseResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedClaimResponseResourcesReferencingInsurer {
			rsc := (*o.RevIncludedClaimResponseResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEligibilityRequestResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEligibilityRequestResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEligibilityRequestResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessRequestResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedProcessRequestResourcesReferencingOrganization {
			rsc := (*o.RevIncludedProcessRequestResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedHealthcareServiceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedHealthcareServiceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedHealthcareServiceResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentReconciliationResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPaymentReconciliationResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPaymentReconciliationResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentReconciliationResourcesReferencingRequestorganization != nil {
		for idx := range *o.RevIncludedPaymentReconciliationResourcesReferencingRequestorganization {
			rsc := (*o.RevIncludedPaymentReconciliationResourcesReferencingRequestorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*o.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingGeneralpractitioner != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingGeneralpractitioner {
			rsc := (*o.RevIncludedPatientResourcesReferencingGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPatientResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageResourcesReferencingPolicyholder != nil {
		for idx := range *o.RevIncludedCoverageResourcesReferencingPolicyholder {
			rsc := (*o.RevIncludedCoverageResourcesReferencingPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageResourcesReferencingPayor != nil {
		for idx := range *o.RevIncludedCoverageResourcesReferencingPayor {
			rsc := (*o.RevIncludedCoverageResourcesReferencingPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingOrganization {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequestorganization != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequestorganization {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequestorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingInsurer {
			rsc := (*o.RevIncludedClaimResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingOrganization {
			rsc := (*o.RevIncludedClaimResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLocationResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedLocationResourcesReferencingOrganization {
			rsc := (*o.RevIncludedLocationResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
