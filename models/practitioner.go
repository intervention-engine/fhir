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

type Practitioner struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active         *bool                                `bson:"active,omitempty" json:"active,omitempty"`
	Name           []HumanName                          `bson:"name,omitempty" json:"name,omitempty"`
	Telecom        []ContactPoint                       `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address        []Address                            `bson:"address,omitempty" json:"address,omitempty"`
	Gender         string                               `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate      *FHIRDateTime                        `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Photo          []Attachment                         `bson:"photo,omitempty" json:"photo,omitempty"`
	Qualification  []PractitionerQualificationComponent `bson:"qualification,omitempty" json:"qualification,omitempty"`
	Communication  []CodeableConcept                    `bson:"communication,omitempty" json:"communication,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Practitioner) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Practitioner"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Practitioner), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Practitioner) GetBSON() (interface{}, error) {
	x.ResourceType = "Practitioner"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "practitioner" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type practitioner Practitioner

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Practitioner) UnmarshalJSON(data []byte) (err error) {
	x2 := practitioner{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = Practitioner(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Practitioner) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Practitioner"
	} else if x.ResourceType != "Practitioner" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Practitioner, instead received %s", x.ResourceType))
	}
	return nil
}

type PractitionerQualificationComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Period          *Period          `bson:"period,omitempty" json:"period,omitempty"`
	Issuer          *Reference       `bson:"issuer,omitempty" json:"issuer,omitempty"`
}

type PractitionerPlus struct {
	Practitioner                     `bson:",inline"`
	PractitionerPlusRelatedResources `bson:",inline"`
}

type PractitionerPlusRelatedResources struct {
	RevIncludedAppointmentResourcesReferencingActor                     *[]Appointment              `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingPractitioner              *[]Appointment              `bson:"_revIncludedAppointmentResourcesReferencingPractitioner,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRequester             *[]ReferralRequest          `bson:"_revIncludedReferralRequestResourcesReferencingRequester,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRecipient             *[]ReferralRequest          `bson:"_revIncludedReferralRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                       *[]Account                  `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref           *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingSubject              *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor               *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref           *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRecipient            *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRecipient,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                     *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                     *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedConsentResourcesReferencingActorPath1                    *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingActorPath1,omitempty"`
	RevIncludedConsentResourcesReferencingActorPath2                    *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingActorPath2,omitempty"`
	RevIncludedConsentResourcesReferencingConsentor                     *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingConsentor,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                     *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                   *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                   *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                    *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingSubject             *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthenticator       *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthenticator,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor              *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref          *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedImagingManifestResourcesReferencingAuthor                *[]ImagingManifest          `bson:"_revIncludedImagingManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedPractitionerRoleResourcesReferencingPractitioner         *[]PractitionerRole         `bson:"_revIncludedPractitionerRoleResourcesReferencingPractitioner,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingRequester               *[]SupplyRequest            `bson:"_revIncludedSupplyRequestResourcesReferencingRequester,omitempty"`
	RevIncludedPersonResourcesReferencingPractitioner                   *[]Person                   `bson:"_revIncludedPersonResourcesReferencingPractitioner,omitempty"`
	RevIncludedPersonResourcesReferencingLink                           *[]Person                   `bson:"_revIncludedPersonResourcesReferencingLink,omitempty"`
	RevIncludedContractResourcesReferencingAgent                        *[]Contract                 `bson:"_revIncludedContractResourcesReferencingAgent,omitempty"`
	RevIncludedContractResourcesReferencingSubject                      *[]Contract                 `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                    *[]Contract                 `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedContractResourcesReferencingSigner                       *[]Contract                 `bson:"_revIncludedContractResourcesReferencingSigner,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingPerformer              *[]RiskAssessment           `bson:"_revIncludedRiskAssessmentResourcesReferencingPerformer,omitempty"`
	RevIncludedGroupResourcesReferencingMember                          *[]Group                    `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                 *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingProvider                *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingProvider,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedCareTeamResourcesReferencingParticipant                  *[]CareTeam                 `bson:"_revIncludedCareTeamResourcesReferencingParticipant,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource          *[]ImplementationGuide      `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingPerformer                *[]ImagingStudy             `bson:"_revIncludedImagingStudyResourcesReferencingPerformer,omitempty"`
	RevIncludedChargeItemResourcesReferencingParticipantactor           *[]ChargeItem               `bson:"_revIncludedChargeItemResourcesReferencingParticipantactor,omitempty"`
	RevIncludedChargeItemResourcesReferencingEnterer                    *[]ChargeItem               `bson:"_revIncludedChargeItemResourcesReferencingEnterer,omitempty"`
	RevIncludedEncounterResourcesReferencingPractitioner                *[]Encounter                `bson:"_revIncludedEncounterResourcesReferencingPractitioner,omitempty"`
	RevIncludedEncounterResourcesReferencingParticipant                 *[]Encounter                `bson:"_revIncludedEncounterResourcesReferencingParticipant,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor           *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom         *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor         *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof          *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson           *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                  *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                 *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender                  *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient               *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor          *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom        *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor        *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof         *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1     *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2     *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedLinkageResourcesReferencingAuthor                        *[]Linkage                  `bson:"_revIncludedLinkageResourcesReferencingAuthor,omitempty"`
	RevIncludedRequestGroupResourcesReferencingAuthor                   *[]RequestGroup             `bson:"_revIncludedRequestGroupResourcesReferencingAuthor,omitempty"`
	RevIncludedRequestGroupResourcesReferencingParticipant              *[]RequestGroup             `bson:"_revIncludedRequestGroupResourcesReferencingParticipant,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition               *[]RequestGroup             `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingRequester               *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingRequester,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPerformer               *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                 *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest            *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingReceiver                *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingReceiver,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingAuthor                  *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingAuthor,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                   *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingSender                  *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingSender,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingResponsible             *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingResponsible,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingEnterer                 *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingEnterer,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref                  *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                      *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                     *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingOwner                            *[]Task                     `bson:"_revIncludedTaskResourcesReferencingOwner,omitempty"`
	RevIncludedTaskResourcesReferencingRequester                        *[]Task                     `bson:"_revIncludedTaskResourcesReferencingRequester,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                          *[]Task                     `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                            *[]Task                     `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                          *[]Task                     `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingCareteam         *[]ExplanationOfBenefit     `bson:"_revIncludedExplanationOfBenefitResourcesReferencingCareteam,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingPayee            *[]ExplanationOfBenefit     `bson:"_revIncludedExplanationOfBenefitResourcesReferencingPayee,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingProvider         *[]ExplanationOfBenefit     `bson:"_revIncludedExplanationOfBenefitResourcesReferencingProvider,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingEnterer          *[]ExplanationOfBenefit     `bson:"_revIncludedExplanationOfBenefitResourcesReferencingEnterer,omitempty"`
	RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator   *[]ResearchStudy            `bson:"_revIncludedResearchStudyResourcesReferencingPrincipalinvestigator,omitempty"`
	RevIncludedSpecimenResourcesReferencingCollector                    *[]Specimen                 `bson:"_revIncludedSpecimenResourcesReferencingCollector,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingRecorder           *[]AllergyIntolerance       `bson:"_revIncludedAllergyIntoleranceResourcesReferencingRecorder,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingAsserter           *[]AllergyIntolerance       `bson:"_revIncludedAllergyIntoleranceResourcesReferencingAsserter,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                    *[]CarePlan                 `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingCaremanager             *[]EpisodeOfCare            `bson:"_revIncludedEpisodeOfCareResourcesReferencingCaremanager,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                   *[]Procedure                `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedListResourcesReferencingItem                             *[]List                     `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSource                           *[]List                     `bson:"_revIncludedListResourcesReferencingSource,omitempty"`
	RevIncludedImmunizationResourcesReferencingPractitioner             *[]Immunization             `bson:"_revIncludedImmunizationResourcesReferencingPractitioner,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingRequester           *[]MedicationRequest        `bson:"_revIncludedMedicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedVisionPrescriptionResourcesReferencingPrescriber         *[]VisionPrescription       `bson:"_revIncludedVisionPrescriptionResourcesReferencingPrescriber,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                         *[]Media                    `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedMediaResourcesReferencingOperator                        *[]Media                    `bson:"_revIncludedMediaResourcesReferencingOperator,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingRequester            *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingRequester,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingPerformer            *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces             *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon              *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedEligibilityResponseResourcesReferencingRequestprovider   *[]EligibilityResponse      `bson:"_revIncludedEligibilityResponseResourcesReferencingRequestprovider,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                          *[]Flag                     `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                           *[]Flag                     `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor             *[]AppointmentResponse      `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingPractitioner      *[]AppointmentResponse      `bson:"_revIncludedAppointmentResponseResourcesReferencingPractitioner,omitempty"`
	RevIncludedAdverseEventResourcesReferencingRecorder                 *[]AdverseEvent             `bson:"_revIncludedAdverseEventResourcesReferencingRecorder,omitempty"`
	RevIncludedObservationResourcesReferencingPerformer                 *[]Observation              `bson:"_revIncludedObservationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPerformer    *[]MedicationAdministration `bson:"_revIncludedMedicationAdministrationResourcesReferencingPerformer,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                     *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                   *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                   *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                    *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                     *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSource            *[]MedicationStatement      `bson:"_revIncludedMedicationStatementResourcesReferencingSource,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRequester        *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon          *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender           *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient        *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                         *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingAuthor                          *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingAuthor,omitempty"`
	RevIncludedClaimResponseResourcesReferencingRequestprovider         *[]ClaimResponse            `bson:"_revIncludedClaimResponseResourcesReferencingRequestprovider,omitempty"`
	RevIncludedEligibilityRequestResourcesReferencingProvider           *[]EligibilityRequest       `bson:"_revIncludedEligibilityRequestResourcesReferencingProvider,omitempty"`
	RevIncludedEligibilityRequestResourcesReferencingEnterer            *[]EligibilityRequest       `bson:"_revIncludedEligibilityRequestResourcesReferencingEnterer,omitempty"`
	RevIncludedProcessRequestResourcesReferencingProvider               *[]ProcessRequest           `bson:"_revIncludedProcessRequestResourcesReferencingProvider,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPerformer          *[]MedicationDispense       `bson:"_revIncludedMedicationDispenseResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingReceiver           *[]MedicationDispense       `bson:"_revIncludedMedicationDispenseResourcesReferencingReceiver,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingResponsibleparty   *[]MedicationDispense       `bson:"_revIncludedMedicationDispenseResourcesReferencingResponsibleparty,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingPerformer            *[]DiagnosticReport         `bson:"_revIncludedDiagnosticReportResourcesReferencingPerformer,omitempty"`
	RevIncludedNutritionOrderResourcesReferencingProvider               *[]NutritionOrder           `bson:"_revIncludedNutritionOrderResourcesReferencingProvider,omitempty"`
	RevIncludedAuditEventResourcesReferencingAgent                      *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingAgent,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                     *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedPaymentReconciliationResourcesReferencingRequestprovider *[]PaymentReconciliation    `bson:"_revIncludedPaymentReconciliationResourcesReferencingRequestprovider,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail              *[]Condition                `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedConditionResourcesReferencingAsserter                    *[]Condition                `bson:"_revIncludedConditionResourcesReferencingAsserter,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                   *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor                    *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingAttester                  *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingAttester,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                     *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingAuthor                  *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingAuthor,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated              *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedPatientResourcesReferencingGeneralpractitioner           *[]Patient                  `bson:"_revIncludedPatientResourcesReferencingGeneralpractitioner,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject         *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor          *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSource          *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSource,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest               *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestprovider       *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequestprovider,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                        *[]Schedule                 `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedSupplyDeliveryResourcesReferencingReceiver               *[]SupplyDelivery           `bson:"_revIncludedSupplyDeliveryResourcesReferencingReceiver,omitempty"`
	RevIncludedSupplyDeliveryResourcesReferencingSupplier               *[]SupplyDelivery           `bson:"_revIncludedSupplyDeliveryResourcesReferencingSupplier,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAssessor           *[]ClinicalImpression       `bson:"_revIncludedClinicalImpressionResourcesReferencingAssessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor              *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom            *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor            *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof             *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1         *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2         *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedClaimResourcesReferencingCareteam                        *[]Claim                    `bson:"_revIncludedClaimResourcesReferencingCareteam,omitempty"`
	RevIncludedClaimResourcesReferencingPayee                           *[]Claim                    `bson:"_revIncludedClaimResourcesReferencingPayee,omitempty"`
	RevIncludedClaimResourcesReferencingProvider                        *[]Claim                    `bson:"_revIncludedClaimResourcesReferencingProvider,omitempty"`
	RevIncludedClaimResourcesReferencingEnterer                         *[]Claim                    `bson:"_revIncludedClaimResourcesReferencingEnterer,omitempty"`
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingPractitioner() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingPractitioner
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingRequester() (referralRequests []ReferralRequest, err error) {
	if p.RevIncludedReferralRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *p.RevIncludedReferralRequestResourcesReferencingRequester
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingRecipient() (referralRequests []ReferralRequest, err error) {
	if p.RevIncludedReferralRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *p.RevIncludedReferralRequestResourcesReferencingRecipient
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if p.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *p.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingSubject() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingSubject
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRecipient() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRecipient
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedConsentResourcesReferencingActorPath1() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingActorPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingActorPath1
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedConsentResourcesReferencingActorPath2() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingActorPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingActorPath2
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedConsentResourcesReferencingConsentor() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingConsentor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingConsentor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingSubject() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingSubject
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthenticator() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedImagingManifestResourcesReferencingAuthor() (imagingManifests []ImagingManifest, err error) {
	if p.RevIncludedImagingManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded imagingManifests not requested")
	} else {
		imagingManifests = *p.RevIncludedImagingManifestResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPractitionerRoleResourcesReferencingPractitioner() (practitionerRoles []PractitionerRole, err error) {
	if p.RevIncludedPractitionerRoleResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded practitionerRoles not requested")
	} else {
		practitionerRoles = *p.RevIncludedPractitionerRoleResourcesReferencingPractitioner
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingRequester() (supplyRequests []SupplyRequest, err error) {
	if p.RevIncludedSupplyRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *p.RevIncludedSupplyRequestResourcesReferencingRequester
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPersonResourcesReferencingPractitioner() (people []Person, err error) {
	if p.RevIncludedPersonResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *p.RevIncludedPersonResourcesReferencingPractitioner
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPersonResourcesReferencingLink() (people []Person, err error) {
	if p.RevIncludedPersonResourcesReferencingLink == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *p.RevIncludedPersonResourcesReferencingLink
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedContractResourcesReferencingAgent() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingAgent
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedContractResourcesReferencingSigner() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSigner == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSigner
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingPerformer() (riskAssessments []RiskAssessment, err error) {
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *p.RevIncludedRiskAssessmentResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if p.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *p.RevIncludedGroupResourcesReferencingMember
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingProvider() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingProvider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingParticipant() (careTeams []CareTeam, err error) {
	if p.RevIncludedCareTeamResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *p.RevIncludedCareTeamResourcesReferencingParticipant
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if p.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *p.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingPerformer() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingParticipantactor() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingParticipantactor == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingParticipantactor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingEnterer() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingEnterer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingPractitioner() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingPractitioner
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingParticipant() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingParticipant
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingAuthor() (linkages []Linkage, err error) {
	if p.RevIncludedLinkageResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *p.RevIncludedLinkageResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingAuthor() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingParticipant() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingParticipant
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingRequester() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingRequester
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPerformer() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingReceiver() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingReceiver
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingAuthor() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingSender() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingSender == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingSender
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingResponsible() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingResponsible == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingResponsible
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingEnterer() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingEnterer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedTaskResourcesReferencingOwner() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingOwner
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedTaskResourcesReferencingRequester() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingRequester
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingCareteam() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingPayee() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingProvider() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingProvider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingEnterer() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchStudyResourcesReferencingPrincipalinvestigator() (researchStudies []ResearchStudy, err error) {
	if p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator == nil {
		err = errors.New("RevIncluded researchStudies not requested")
	} else {
		researchStudies = *p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingCollector() (specimen []Specimen, err error) {
	if p.RevIncludedSpecimenResourcesReferencingCollector == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *p.RevIncludedSpecimenResourcesReferencingCollector
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingRecorder() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingAsserter() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if p.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *p.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingCaremanager() (episodeOfCares []EpisodeOfCare, err error) {
	if p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingItem
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedListResourcesReferencingSource() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingSource == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingSource
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingPractitioner() (immunizations []Immunization, err error) {
	if p.RevIncludedImmunizationResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *p.RevIncludedImmunizationResourcesReferencingPractitioner
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingRequester() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingRequester
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedVisionPrescriptionResourcesReferencingPrescriber() (visionPrescriptions []VisionPrescription, err error) {
	if p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber == nil {
		err = errors.New("RevIncluded visionPrescriptions not requested")
	} else {
		visionPrescriptions = *p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if p.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *p.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMediaResourcesReferencingOperator() (media []Media, err error) {
	if p.RevIncludedMediaResourcesReferencingOperator == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *p.RevIncludedMediaResourcesReferencingOperator
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingRequester() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingRequester
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingPerformer() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedEligibilityResponseResourcesReferencingRequestprovider() (eligibilityResponses []EligibilityResponse, err error) {
	if p.RevIncludedEligibilityResponseResourcesReferencingRequestprovider == nil {
		err = errors.New("RevIncluded eligibilityResponses not requested")
	} else {
		eligibilityResponses = *p.RevIncludedEligibilityResponseResourcesReferencingRequestprovider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedFlagResourcesReferencingAuthor() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if p.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *p.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingPractitioner() (appointmentResponses []AppointmentResponse, err error) {
	if p.RevIncludedAppointmentResponseResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *p.RevIncludedAppointmentResponseResourcesReferencingPractitioner
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingRecorder() (adverseEvents []AdverseEvent, err error) {
	if p.RevIncludedAdverseEventResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *p.RevIncludedAdverseEventResourcesReferencingRecorder
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPerformer() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPerformer() (medicationAdministrations []MedicationAdministration, err error) {
	if p.RevIncludedMedicationAdministrationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.RevIncludedMedicationAdministrationResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSource() (medicationStatements []MedicationStatement, err error) {
	if p.RevIncludedMedicationStatementResourcesReferencingSource == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *p.RevIncludedMedicationStatementResourcesReferencingSource
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRequester() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingRequester
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedBasicResourcesReferencingAuthor() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedClaimResponseResourcesReferencingRequestprovider() (claimResponses []ClaimResponse, err error) {
	if p.RevIncludedClaimResponseResourcesReferencingRequestprovider == nil {
		err = errors.New("RevIncluded claimResponses not requested")
	} else {
		claimResponses = *p.RevIncludedClaimResponseResourcesReferencingRequestprovider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedEligibilityRequestResourcesReferencingProvider() (eligibilityRequests []EligibilityRequest, err error) {
	if p.RevIncludedEligibilityRequestResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded eligibilityRequests not requested")
	} else {
		eligibilityRequests = *p.RevIncludedEligibilityRequestResourcesReferencingProvider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedEligibilityRequestResourcesReferencingEnterer() (eligibilityRequests []EligibilityRequest, err error) {
	if p.RevIncludedEligibilityRequestResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded eligibilityRequests not requested")
	} else {
		eligibilityRequests = *p.RevIncludedEligibilityRequestResourcesReferencingEnterer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProcessRequestResourcesReferencingProvider() (processRequests []ProcessRequest, err error) {
	if p.RevIncludedProcessRequestResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded processRequests not requested")
	} else {
		processRequests = *p.RevIncludedProcessRequestResourcesReferencingProvider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPerformer() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingReceiver() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingReceiver
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingResponsibleparty() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingPerformer() (diagnosticReports []DiagnosticReport, err error) {
	if p.RevIncludedDiagnosticReportResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *p.RevIncludedDiagnosticReportResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedNutritionOrderResourcesReferencingProvider() (nutritionOrders []NutritionOrder, err error) {
	if p.RevIncludedNutritionOrderResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded nutritionOrders not requested")
	} else {
		nutritionOrders = *p.RevIncludedNutritionOrderResourcesReferencingProvider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingAgent() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingAgent
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPaymentReconciliationResourcesReferencingRequestprovider() (paymentReconciliations []PaymentReconciliation, err error) {
	if p.RevIncludedPaymentReconciliationResourcesReferencingRequestprovider == nil {
		err = errors.New("RevIncluded paymentReconciliations not requested")
	} else {
		paymentReconciliations = *p.RevIncludedPaymentReconciliationResourcesReferencingRequestprovider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedConditionResourcesReferencingAsserter() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingAsserter
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAuthor() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAttester() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingAttester == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingAttester
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingAuthor() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPatientResourcesReferencingGeneralpractitioner() (patients []Patient, err error) {
	if p.RevIncludedPatientResourcesReferencingGeneralpractitioner == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *p.RevIncludedPatientResourcesReferencingGeneralpractitioner
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSource() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSource
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if p.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *p.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestprovider() (processResponses []ProcessResponse, err error) {
	if p.RevIncludedProcessResponseResourcesReferencingRequestprovider == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *p.RevIncludedProcessResponseResourcesReferencingRequestprovider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if p.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *p.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedSupplyDeliveryResourcesReferencingReceiver() (supplyDeliveries []SupplyDelivery, err error) {
	if p.RevIncludedSupplyDeliveryResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded supplyDeliveries not requested")
	} else {
		supplyDeliveries = *p.RevIncludedSupplyDeliveryResourcesReferencingReceiver
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedSupplyDeliveryResourcesReferencingSupplier() (supplyDeliveries []SupplyDelivery, err error) {
	if p.RevIncludedSupplyDeliveryResourcesReferencingSupplier == nil {
		err = errors.New("RevIncluded supplyDeliveries not requested")
	} else {
		supplyDeliveries = *p.RevIncludedSupplyDeliveryResourcesReferencingSupplier
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAssessor() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingAssessor == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingAssessor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedClaimResourcesReferencingCareteam() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingCareteam == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingCareteam
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedClaimResourcesReferencingPayee() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingPayee
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedClaimResourcesReferencingProvider() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingProvider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedClaimResourcesReferencingEnterer() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingEnterer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedReferralRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedReferralRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedReferralRequestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedReferralRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*p.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*p.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*p.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingActorPath1 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingActorPath1 {
			rsc := (*p.RevIncludedConsentResourcesReferencingActorPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingActorPath2 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingActorPath2 {
			rsc := (*p.RevIncludedConsentResourcesReferencingActorPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*p.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*p.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingManifestResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedImagingManifestResourcesReferencingAuthor {
			rsc := (*p.RevIncludedImagingManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPractitionerRoleResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedPractitionerRoleResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedPractitionerRoleResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedPersonResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingLink {
			rsc := (*p.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingAgent {
			rsc := (*p.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSigner {
			rsc := (*p.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingMember {
			rsc := (*p.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingProvider {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingParticipantactor != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingParticipantactor {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingParticipantactor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedEncounterResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingParticipant {
			rsc := (*p.RevIncludedEncounterResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingAuthor {
			rsc := (*p.RevIncludedLinkageResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingAuthor {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingAuthor {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingSender {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingEnterer {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*p.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*p.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingProvider {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator != nil {
		for idx := range *p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator {
			rsc := (*p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingCollector != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingCollector {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingCollector)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager != nil {
		for idx := range *p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager {
			rsc := (*p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSource {
			rsc := (*p.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber != nil {
		for idx := range *p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber {
			rsc := (*p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*p.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*p.RevIncludedMediaResourcesReferencingOperator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEligibilityResponseResourcesReferencingRequestprovider != nil {
		for idx := range *p.RevIncludedEligibilityResponseResourcesReferencingRequestprovider {
			rsc := (*p.RevIncludedEligibilityResponseResourcesReferencingRequestprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*p.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*p.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*p.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResponseResourcesReferencingRequestprovider != nil {
		for idx := range *p.RevIncludedClaimResponseResourcesReferencingRequestprovider {
			rsc := (*p.RevIncludedClaimResponseResourcesReferencingRequestprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEligibilityRequestResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedEligibilityRequestResourcesReferencingProvider {
			rsc := (*p.RevIncludedEligibilityRequestResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEligibilityRequestResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedEligibilityRequestResourcesReferencingEnterer {
			rsc := (*p.RevIncludedEligibilityRequestResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcessRequestResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedProcessRequestResourcesReferencingProvider {
			rsc := (*p.RevIncludedProcessRequestResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingProvider {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentReconciliationResourcesReferencingRequestprovider != nil {
		for idx := range *p.RevIncludedPaymentReconciliationResourcesReferencingRequestprovider {
			rsc := (*p.RevIncludedPaymentReconciliationResourcesReferencingRequestprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*p.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingGeneralpractitioner != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingGeneralpractitioner {
			rsc := (*p.RevIncludedPatientResourcesReferencingGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*p.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequestprovider != nil {
		for idx := range *p.RevIncludedProcessResponseResourcesReferencingRequestprovider {
			rsc := (*p.RevIncludedProcessResponseResourcesReferencingRequestprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*p.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingReceiver {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingSupplier != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingSupplier {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingAssessor != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingAssessor {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingAssessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingCareteam {
			rsc := (*p.RevIncludedClaimResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*p.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingProvider {
			rsc := (*p.RevIncludedClaimResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingEnterer {
			rsc := (*p.RevIncludedClaimResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *PractitionerPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedReferralRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedReferralRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedReferralRequestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedReferralRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*p.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*p.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*p.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingActorPath1 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingActorPath1 {
			rsc := (*p.RevIncludedConsentResourcesReferencingActorPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingActorPath2 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingActorPath2 {
			rsc := (*p.RevIncludedConsentResourcesReferencingActorPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*p.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*p.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingManifestResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedImagingManifestResourcesReferencingAuthor {
			rsc := (*p.RevIncludedImagingManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPractitionerRoleResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedPractitionerRoleResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedPractitionerRoleResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedPersonResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingLink {
			rsc := (*p.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingAgent {
			rsc := (*p.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSigner {
			rsc := (*p.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingMember {
			rsc := (*p.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingProvider {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingParticipantactor != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingParticipantactor {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingParticipantactor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedEncounterResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingParticipant {
			rsc := (*p.RevIncludedEncounterResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingAuthor {
			rsc := (*p.RevIncludedLinkageResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingAuthor {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingAuthor {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingSender {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingEnterer {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*p.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*p.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingProvider {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator != nil {
		for idx := range *p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator {
			rsc := (*p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingCollector != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingCollector {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingCollector)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager != nil {
		for idx := range *p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager {
			rsc := (*p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSource {
			rsc := (*p.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber != nil {
		for idx := range *p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber {
			rsc := (*p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*p.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*p.RevIncludedMediaResourcesReferencingOperator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEligibilityResponseResourcesReferencingRequestprovider != nil {
		for idx := range *p.RevIncludedEligibilityResponseResourcesReferencingRequestprovider {
			rsc := (*p.RevIncludedEligibilityResponseResourcesReferencingRequestprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*p.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*p.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*p.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResponseResourcesReferencingRequestprovider != nil {
		for idx := range *p.RevIncludedClaimResponseResourcesReferencingRequestprovider {
			rsc := (*p.RevIncludedClaimResponseResourcesReferencingRequestprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEligibilityRequestResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedEligibilityRequestResourcesReferencingProvider {
			rsc := (*p.RevIncludedEligibilityRequestResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEligibilityRequestResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedEligibilityRequestResourcesReferencingEnterer {
			rsc := (*p.RevIncludedEligibilityRequestResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcessRequestResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedProcessRequestResourcesReferencingProvider {
			rsc := (*p.RevIncludedProcessRequestResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingProvider {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentReconciliationResourcesReferencingRequestprovider != nil {
		for idx := range *p.RevIncludedPaymentReconciliationResourcesReferencingRequestprovider {
			rsc := (*p.RevIncludedPaymentReconciliationResourcesReferencingRequestprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*p.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingGeneralpractitioner != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingGeneralpractitioner {
			rsc := (*p.RevIncludedPatientResourcesReferencingGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*p.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequestprovider != nil {
		for idx := range *p.RevIncludedProcessResponseResourcesReferencingRequestprovider {
			rsc := (*p.RevIncludedProcessResponseResourcesReferencingRequestprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*p.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingReceiver {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingSupplier != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingSupplier {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingAssessor != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingAssessor {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingAssessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingCareteam {
			rsc := (*p.RevIncludedClaimResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*p.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingProvider {
			rsc := (*p.RevIncludedClaimResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingEnterer {
			rsc := (*p.RevIncludedClaimResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
