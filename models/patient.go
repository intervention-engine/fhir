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

type Patient struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active               *bool                           `bson:"active,omitempty" json:"active,omitempty"`
	Name                 []HumanName                     `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []ContactPoint                  `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               string                          `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *FHIRDateTime                   `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                           `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *FHIRDateTime                   `bson:"deceasedDateTime,omitempty" json:"deceasedDateTime,omitempty"`
	Address              []Address                       `bson:"address,omitempty" json:"address,omitempty"`
	MaritalStatus        *CodeableConcept                `bson:"maritalStatus,omitempty" json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *bool                           `bson:"multipleBirthBoolean,omitempty" json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int32                          `bson:"multipleBirthInteger,omitempty" json:"multipleBirthInteger,omitempty"`
	Photo                []Attachment                    `bson:"photo,omitempty" json:"photo,omitempty"`
	Contact              []PatientContactComponent       `bson:"contact,omitempty" json:"contact,omitempty"`
	Animal               *PatientAnimalComponent         `bson:"animal,omitempty" json:"animal,omitempty"`
	Communication        []PatientCommunicationComponent `bson:"communication,omitempty" json:"communication,omitempty"`
	GeneralPractitioner  []Reference                     `bson:"generalPractitioner,omitempty" json:"generalPractitioner,omitempty"`
	ManagingOrganization *Reference                      `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Link                 []PatientLinkComponent          `bson:"link,omitempty" json:"link,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Patient) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Patient"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Patient), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Patient) GetBSON() (interface{}, error) {
	x.ResourceType = "Patient"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "patient" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type patient Patient

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Patient) UnmarshalJSON(data []byte) (err error) {
	x2 := patient{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Patient(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Patient) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Patient"
	} else if x.ResourceType != "Patient" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Patient, instead received %s", x.ResourceType))
	}
	return nil
}

type PatientContactComponent struct {
	BackboneElement `bson:",inline"`
	Relationship    []CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name            *HumanName        `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint    `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address         *Address          `bson:"address,omitempty" json:"address,omitempty"`
	Gender          string            `bson:"gender,omitempty" json:"gender,omitempty"`
	Organization    *Reference        `bson:"organization,omitempty" json:"organization,omitempty"`
	Period          *Period           `bson:"period,omitempty" json:"period,omitempty"`
}

type PatientAnimalComponent struct {
	BackboneElement `bson:",inline"`
	Species         *CodeableConcept `bson:"species,omitempty" json:"species,omitempty"`
	Breed           *CodeableConcept `bson:"breed,omitempty" json:"breed,omitempty"`
	GenderStatus    *CodeableConcept `bson:"genderStatus,omitempty" json:"genderStatus,omitempty"`
}

type PatientCommunicationComponent struct {
	BackboneElement `bson:",inline"`
	Language        *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	Preferred       *bool            `bson:"preferred,omitempty" json:"preferred,omitempty"`
}

type PatientLinkComponent struct {
	BackboneElement `bson:",inline"`
	Other           *Reference `bson:"other,omitempty" json:"other,omitempty"`
	Type            string     `bson:"type,omitempty" json:"type,omitempty"`
}

type PatientPlus struct {
	Patient                     `bson:",inline"`
	PatientPlusRelatedResources `bson:",inline"`
}

type PatientPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByLink                         *[]Patient                    `bson:"_includedPatientResourcesReferencedByLink,omitempty"`
	IncludedRelatedPersonResourcesReferencedByLink                   *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByLink,omitempty"`
	IncludedPractitionerResourcesReferencedByGeneralpractitioner     *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByGeneralpractitioner,omitempty"`
	IncludedOrganizationResourcesReferencedByGeneralpractitioner     *[]Organization               `bson:"_includedOrganizationResourcesReferencedByGeneralpractitioner,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization            *[]Organization               `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                  *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingPatient                *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingPatient,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRequester          *[]ReferralRequest            `bson:"_revIncludedReferralRequestResourcesReferencingRequester,omitempty"`
	RevIncludedReferralRequestResourcesReferencingSubject            *[]ReferralRequest            `bson:"_revIncludedReferralRequestResourcesReferencingSubject,omitempty"`
	RevIncludedReferralRequestResourcesReferencingPatient            *[]ReferralRequest            `bson:"_revIncludedReferralRequestResourcesReferencingPatient,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                    *[]Account                    `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedAccountResourcesReferencingPatient                    *[]Account                    `bson:"_revIncludedAccountResourcesReferencingPatient,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref        *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingSubject           *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor            *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref        *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingPatient           *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingPatient,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRecipient         *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRecipient,omitempty"`
	RevIncludedGoalResourcesReferencingPatient                       *[]Goal                       `bson:"_revIncludedGoalResourcesReferencingPatient,omitempty"`
	RevIncludedGoalResourcesReferencingSubject                       *[]Goal                       `bson:"_revIncludedGoalResourcesReferencingSubject,omitempty"`
	RevIncludedEnrollmentRequestResourcesReferencingSubject          *[]EnrollmentRequest          `bson:"_revIncludedEnrollmentRequestResourcesReferencingSubject,omitempty"`
	RevIncludedEnrollmentRequestResourcesReferencingPatient          *[]EnrollmentRequest          `bson:"_revIncludedEnrollmentRequestResourcesReferencingPatient,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                  *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                  *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedConsentResourcesReferencingActorPath1                 *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingActorPath1,omitempty"`
	RevIncludedConsentResourcesReferencingActorPath2                 *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingActorPath2,omitempty"`
	RevIncludedConsentResourcesReferencingPatient                    *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingPatient,omitempty"`
	RevIncludedConsentResourcesReferencingConsentor                  *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingConsentor,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                  *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                 *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1             *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2             *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedResearchSubjectResourcesReferencingIndividual         *[]ResearchSubject            `bson:"_revIncludedResearchSubjectResourcesReferencingIndividual,omitempty"`
	RevIncludedResearchSubjectResourcesReferencingPatient            *[]ResearchSubject            `bson:"_revIncludedResearchSubjectResourcesReferencingPatient,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingSubject          *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingPatient          *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingPatient,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor           *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref       *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedImagingManifestResourcesReferencingAuthor             *[]ImagingManifest            `bson:"_revIncludedImagingManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedImagingManifestResourcesReferencingPatient            *[]ImagingManifest            `bson:"_revIncludedImagingManifestResourcesReferencingPatient,omitempty"`
	RevIncludedMeasureReportResourcesReferencingPatient              *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingPatient,omitempty"`
	RevIncludedRelatedPersonResourcesReferencingPatient              *[]RelatedPerson              `bson:"_revIncludedRelatedPersonResourcesReferencingPatient,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingRequester            *[]SupplyRequest              `bson:"_revIncludedSupplyRequestResourcesReferencingRequester,omitempty"`
	RevIncludedPersonResourcesReferencingLink                        *[]Person                     `bson:"_revIncludedPersonResourcesReferencingLink,omitempty"`
	RevIncludedPersonResourcesReferencingPatient                     *[]Person                     `bson:"_revIncludedPersonResourcesReferencingPatient,omitempty"`
	RevIncludedContractResourcesReferencingAgent                     *[]Contract                   `bson:"_revIncludedContractResourcesReferencingAgent,omitempty"`
	RevIncludedContractResourcesReferencingPatient                   *[]Contract                   `bson:"_revIncludedContractResourcesReferencingPatient,omitempty"`
	RevIncludedContractResourcesReferencingSubject                   *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                 *[]Contract                   `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedContractResourcesReferencingSigner                    *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSigner,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingSubject             *[]RiskAssessment             `bson:"_revIncludedRiskAssessmentResourcesReferencingSubject,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingPatient             *[]RiskAssessment             `bson:"_revIncludedRiskAssessmentResourcesReferencingPatient,omitempty"`
	RevIncludedGroupResourcesReferencingMember                       *[]Group                      `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest              *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse             *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedCareTeamResourcesReferencingPatient                   *[]CareTeam                   `bson:"_revIncludedCareTeamResourcesReferencingPatient,omitempty"`
	RevIncludedCareTeamResourcesReferencingSubject                   *[]CareTeam                   `bson:"_revIncludedCareTeamResourcesReferencingSubject,omitempty"`
	RevIncludedCareTeamResourcesReferencingParticipant               *[]CareTeam                   `bson:"_revIncludedCareTeamResourcesReferencingParticipant,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource       *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingPatient               *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingPatient,omitempty"`
	RevIncludedFamilyMemberHistoryResourcesReferencingPatient        *[]FamilyMemberHistory        `bson:"_revIncludedFamilyMemberHistoryResourcesReferencingPatient,omitempty"`
	RevIncludedChargeItemResourcesReferencingSubject                 *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingSubject,omitempty"`
	RevIncludedChargeItemResourcesReferencingParticipantactor        *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingParticipantactor,omitempty"`
	RevIncludedChargeItemResourcesReferencingPatient                 *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingPatient,omitempty"`
	RevIncludedChargeItemResourcesReferencingEnterer                 *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingEnterer,omitempty"`
	RevIncludedEncounterResourcesReferencingSubject                  *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingSubject,omitempty"`
	RevIncludedEncounterResourcesReferencingPatient                  *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingPatient,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor        *[]ServiceDefinition          `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom      *[]ServiceDefinition          `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor      *[]ServiceDefinition          `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof       *[]ServiceDefinition          `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson        *[]ServiceDefinition          `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingSubject              *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof               *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon              *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender               *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingPatient              *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingPatient,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient            *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor       *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom     *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor     *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof      *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1  *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2  *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDeviceUseStatementResourcesReferencingSubject         *[]DeviceUseStatement         `bson:"_revIncludedDeviceUseStatementResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceUseStatementResourcesReferencingPatient         *[]DeviceUseStatement         `bson:"_revIncludedDeviceUseStatementResourcesReferencingPatient,omitempty"`
	RevIncludedRequestGroupResourcesReferencingSubject               *[]RequestGroup               `bson:"_revIncludedRequestGroupResourcesReferencingSubject,omitempty"`
	RevIncludedRequestGroupResourcesReferencingParticipant           *[]RequestGroup               `bson:"_revIncludedRequestGroupResourcesReferencingParticipant,omitempty"`
	RevIncludedRequestGroupResourcesReferencingPatient               *[]RequestGroup               `bson:"_revIncludedRequestGroupResourcesReferencingPatient,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition            *[]RequestGroup               `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPerformer            *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingSubject              *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon              *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest         *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPatient              *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPatient,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingPatient *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingPatient,omitempty"`
	RevIncludedBodySiteResourcesReferencingPatient                   *[]BodySite                   `bson:"_revIncludedBodySiteResourcesReferencingPatient,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref               *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                   *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingPatient                 *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingPatient,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                  *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingOwner                         *[]Task                       `bson:"_revIncludedTaskResourcesReferencingOwner,omitempty"`
	RevIncludedTaskResourcesReferencingRequester                     *[]Task                       `bson:"_revIncludedTaskResourcesReferencingRequester,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                       *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                         *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                       *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedTaskResourcesReferencingPatient                       *[]Task                       `bson:"_revIncludedTaskResourcesReferencingPatient,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingPayee         *[]ExplanationOfBenefit       `bson:"_revIncludedExplanationOfBenefitResourcesReferencingPayee,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingPatient       *[]ExplanationOfBenefit       `bson:"_revIncludedExplanationOfBenefitResourcesReferencingPatient,omitempty"`
	RevIncludedSpecimenResourcesReferencingSubject                   *[]Specimen                   `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedSpecimenResourcesReferencingPatient                   *[]Specimen                   `bson:"_revIncludedSpecimenResourcesReferencingPatient,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingRecorder        *[]AllergyIntolerance         `bson:"_revIncludedAllergyIntoleranceResourcesReferencingRecorder,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingAsserter        *[]AllergyIntolerance         `bson:"_revIncludedAllergyIntoleranceResourcesReferencingAsserter,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingPatient         *[]AllergyIntolerance         `bson:"_revIncludedAllergyIntoleranceResourcesReferencingPatient,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                 *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedCarePlanResourcesReferencingSubject                   *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingSubject,omitempty"`
	RevIncludedCarePlanResourcesReferencingPatient                   *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingPatient,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingPatient              *[]EpisodeOfCare              `bson:"_revIncludedEpisodeOfCareResourcesReferencingPatient,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureResourcesReferencingSubject                  *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureResourcesReferencingPatient                  *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingPatient,omitempty"`
	RevIncludedListResourcesReferencingItem                          *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                       *[]List                       `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingPatient                       *[]List                       `bson:"_revIncludedListResourcesReferencingPatient,omitempty"`
	RevIncludedListResourcesReferencingSource                        *[]List                       `bson:"_revIncludedListResourcesReferencingSource,omitempty"`
	RevIncludedImmunizationResourcesReferencingPatient               *[]Immunization               `bson:"_revIncludedImmunizationResourcesReferencingPatient,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingRequester        *[]MedicationRequest          `bson:"_revIncludedMedicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingSubject          *[]MedicationRequest          `bson:"_revIncludedMedicationRequestResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingPatient          *[]MedicationRequest          `bson:"_revIncludedMedicationRequestResourcesReferencingPatient,omitempty"`
	RevIncludedDeviceResourcesReferencingPatient                     *[]Device                     `bson:"_revIncludedDeviceResourcesReferencingPatient,omitempty"`
	RevIncludedVisionPrescriptionResourcesReferencingPatient         *[]VisionPrescription         `bson:"_revIncludedVisionPrescriptionResourcesReferencingPatient,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                      *[]Media                      `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedMediaResourcesReferencingPatient                      *[]Media                      `bson:"_revIncludedMediaResourcesReferencingPatient,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingPerformer         *[]ProcedureRequest           `bson:"_revIncludedProcedureRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces          *[]ProcedureRequest           `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingSubject           *[]ProcedureRequest           `bson:"_revIncludedProcedureRequestResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon           *[]ProcedureRequest           `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingPatient           *[]ProcedureRequest           `bson:"_revIncludedProcedureRequestResourcesReferencingPatient,omitempty"`
	RevIncludedSequenceResourcesReferencingPatient                   *[]Sequence                   `bson:"_revIncludedSequenceResourcesReferencingPatient,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                       *[]Flag                       `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedFlagResourcesReferencingPatient                       *[]Flag                       `bson:"_revIncludedFlagResourcesReferencingPatient,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                        *[]Flag                       `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor          *[]AppointmentResponse        `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingPatient        *[]AppointmentResponse        `bson:"_revIncludedAppointmentResponseResourcesReferencingPatient,omitempty"`
	RevIncludedAdverseEventResourcesReferencingRecorder              *[]AdverseEvent               `bson:"_revIncludedAdverseEventResourcesReferencingRecorder,omitempty"`
	RevIncludedAdverseEventResourcesReferencingSubject               *[]AdverseEvent               `bson:"_revIncludedAdverseEventResourcesReferencingSubject,omitempty"`
	RevIncludedGuidanceResponseResourcesReferencingPatient           *[]GuidanceResponse           `bson:"_revIncludedGuidanceResponseResourcesReferencingPatient,omitempty"`
	RevIncludedGuidanceResponseResourcesReferencingSubject           *[]GuidanceResponse           `bson:"_revIncludedGuidanceResponseResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingSubject                *[]Observation                `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingPatient                *[]Observation                `bson:"_revIncludedObservationResourcesReferencingPatient,omitempty"`
	RevIncludedObservationResourcesReferencingPerformer              *[]Observation                `bson:"_revIncludedObservationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPerformer *[]MedicationAdministration   `bson:"_revIncludedMedicationAdministrationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingSubject   *[]MedicationAdministration   `bson:"_revIncludedMedicationAdministrationResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPatient   *[]MedicationAdministration   `bson:"_revIncludedMedicationAdministrationResourcesReferencingPatient,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                  *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                 *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                  *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSubject        *[]MedicationStatement        `bson:"_revIncludedMedicationStatementResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingPatient        *[]MedicationStatement        `bson:"_revIncludedMedicationStatementResourcesReferencingPatient,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSource         *[]MedicationStatement        `bson:"_revIncludedMedicationStatementResourcesReferencingSource,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRequester     *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSubject       *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon       *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender        *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingPatient       *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingPatient,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient     *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                      *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingPatient                      *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingPatient,omitempty"`
	RevIncludedBasicResourcesReferencingAuthor                       *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingAuthor,omitempty"`
	RevIncludedClaimResponseResourcesReferencingPatient              *[]ClaimResponse              `bson:"_revIncludedClaimResponseResourcesReferencingPatient,omitempty"`
	RevIncludedEligibilityRequestResourcesReferencingPatient         *[]EligibilityRequest         `bson:"_revIncludedEligibilityRequestResourcesReferencingPatient,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPerformer       *[]MedicationDispense         `bson:"_revIncludedMedicationDispenseResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingReceiver        *[]MedicationDispense         `bson:"_revIncludedMedicationDispenseResourcesReferencingReceiver,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingSubject         *[]MedicationDispense         `bson:"_revIncludedMedicationDispenseResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPatient         *[]MedicationDispense         `bson:"_revIncludedMedicationDispenseResourcesReferencingPatient,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject           *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingPatient           *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingPatient,omitempty"`
	RevIncludedNutritionOrderResourcesReferencingPatient             *[]NutritionOrder             `bson:"_revIncludedNutritionOrderResourcesReferencingPatient,omitempty"`
	RevIncludedAuditEventResourcesReferencingAgent                   *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingAgent,omitempty"`
	RevIncludedAuditEventResourcesReferencingPatientPath1            *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingPatientPath1,omitempty"`
	RevIncludedAuditEventResourcesReferencingPatientPath2            *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingPatientPath2,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                  *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail           *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedConditionResourcesReferencingSubject                  *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingSubject,omitempty"`
	RevIncludedConditionResourcesReferencingAsserter                 *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingAsserter,omitempty"`
	RevIncludedConditionResourcesReferencingPatient                  *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingPatient,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor                 *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingAttester               *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingAttester,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                  *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedCompositionResourcesReferencingPatient                *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingPatient,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingPatient              *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingPatient,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated           *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedPatientResourcesReferencingLink                       *[]Patient                    `bson:"_revIncludedPatientResourcesReferencingLink,omitempty"`
	RevIncludedCoverageResourcesReferencingSubscriber                *[]Coverage                   `bson:"_revIncludedCoverageResourcesReferencingSubscriber,omitempty"`
	RevIncludedCoverageResourcesReferencingPayor                     *[]Coverage                   `bson:"_revIncludedCoverageResourcesReferencingPayor,omitempty"`
	RevIncludedCoverageResourcesReferencingBeneficiary               *[]Coverage                   `bson:"_revIncludedCoverageResourcesReferencingBeneficiary,omitempty"`
	RevIncludedCoverageResourcesReferencingPolicyholder              *[]Coverage                   `bson:"_revIncludedCoverageResourcesReferencingPolicyholder,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject      *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor       *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingPatient      *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingPatient,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSource       *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSource,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest            *[]ProcessResponse            `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                     *[]Schedule                   `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedSupplyDeliveryResourcesReferencingPatient             *[]SupplyDelivery             `bson:"_revIncludedSupplyDeliveryResourcesReferencingPatient,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSubject         *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSubject,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPatient         *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingPatient,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor           *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom         *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor         *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof          *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1      *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2      *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedClaimResourcesReferencingPayee                        *[]Claim                      `bson:"_revIncludedClaimResourcesReferencingPayee,omitempty"`
	RevIncludedClaimResourcesReferencingPatient                      *[]Claim                      `bson:"_revIncludedClaimResourcesReferencingPatient,omitempty"`
}

func (p *PatientPlusRelatedResources) GetIncludedPatientResourceReferencedByLink() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedByLink == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedByLink) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedByLink))
	} else if len(*p.IncludedPatientResourcesReferencedByLink) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedByLink)[0]
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByLink() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedRelatedPersonResourcesReferencedByLink == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByLink) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedRelatedPersonResourcesReferencedByLink))
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByLink) == 1 {
		relatedPerson = &(*p.IncludedRelatedPersonResourcesReferencedByLink)[0]
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByGeneralpractitioner() (practitioners []Practitioner, err error) {
	if p.IncludedPractitionerResourcesReferencedByGeneralpractitioner == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *p.IncludedPractitionerResourcesReferencedByGeneralpractitioner
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByGeneralpractitioner() (organizations []Organization, err error) {
	if p.IncludedOrganizationResourcesReferencedByGeneralpractitioner == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *p.IncludedOrganizationResourcesReferencedByGeneralpractitioner
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if p.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*p.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*p.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingPatient() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingRequester() (referralRequests []ReferralRequest, err error) {
	if p.RevIncludedReferralRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *p.RevIncludedReferralRequestResourcesReferencingRequester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingSubject() (referralRequests []ReferralRequest, err error) {
	if p.RevIncludedReferralRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *p.RevIncludedReferralRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingPatient() (referralRequests []ReferralRequest, err error) {
	if p.RevIncludedReferralRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *p.RevIncludedReferralRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if p.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *p.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAccountResourcesReferencingPatient() (accounts []Account, err error) {
	if p.RevIncludedAccountResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *p.RevIncludedAccountResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingSubject() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingPatient() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRecipient() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRecipient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedGoalResourcesReferencingPatient() (goals []Goal, err error) {
	if p.RevIncludedGoalResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded goals not requested")
	} else {
		goals = *p.RevIncludedGoalResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedGoalResourcesReferencingSubject() (goals []Goal, err error) {
	if p.RevIncludedGoalResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded goals not requested")
	} else {
		goals = *p.RevIncludedGoalResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEnrollmentRequestResourcesReferencingSubject() (enrollmentRequests []EnrollmentRequest, err error) {
	if p.RevIncludedEnrollmentRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded enrollmentRequests not requested")
	} else {
		enrollmentRequests = *p.RevIncludedEnrollmentRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEnrollmentRequestResourcesReferencingPatient() (enrollmentRequests []EnrollmentRequest, err error) {
	if p.RevIncludedEnrollmentRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded enrollmentRequests not requested")
	} else {
		enrollmentRequests = *p.RevIncludedEnrollmentRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConsentResourcesReferencingActorPath1() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingActorPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingActorPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConsentResourcesReferencingActorPath2() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingActorPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingActorPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConsentResourcesReferencingPatient() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConsentResourcesReferencingConsentor() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingConsentor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingConsentor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchSubjectResourcesReferencingIndividual() (researchSubjects []ResearchSubject, err error) {
	if p.RevIncludedResearchSubjectResourcesReferencingIndividual == nil {
		err = errors.New("RevIncluded researchSubjects not requested")
	} else {
		researchSubjects = *p.RevIncludedResearchSubjectResourcesReferencingIndividual
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchSubjectResourcesReferencingPatient() (researchSubjects []ResearchSubject, err error) {
	if p.RevIncludedResearchSubjectResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded researchSubjects not requested")
	} else {
		researchSubjects = *p.RevIncludedResearchSubjectResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingSubject() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingPatient() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImagingManifestResourcesReferencingAuthor() (imagingManifests []ImagingManifest, err error) {
	if p.RevIncludedImagingManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded imagingManifests not requested")
	} else {
		imagingManifests = *p.RevIncludedImagingManifestResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImagingManifestResourcesReferencingPatient() (imagingManifests []ImagingManifest, err error) {
	if p.RevIncludedImagingManifestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded imagingManifests not requested")
	} else {
		imagingManifests = *p.RevIncludedImagingManifestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingPatient() (measureReports []MeasureReport, err error) {
	if p.RevIncludedMeasureReportResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *p.RevIncludedMeasureReportResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRelatedPersonResourcesReferencingPatient() (relatedPeople []RelatedPerson, err error) {
	if p.RevIncludedRelatedPersonResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded relatedPeople not requested")
	} else {
		relatedPeople = *p.RevIncludedRelatedPersonResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingRequester() (supplyRequests []SupplyRequest, err error) {
	if p.RevIncludedSupplyRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *p.RevIncludedSupplyRequestResourcesReferencingRequester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPersonResourcesReferencingLink() (people []Person, err error) {
	if p.RevIncludedPersonResourcesReferencingLink == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *p.RevIncludedPersonResourcesReferencingLink
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPersonResourcesReferencingPatient() (people []Person, err error) {
	if p.RevIncludedPersonResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *p.RevIncludedPersonResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedContractResourcesReferencingAgent() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingAgent
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedContractResourcesReferencingPatient() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedContractResourcesReferencingSigner() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSigner == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSigner
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingSubject() (riskAssessments []RiskAssessment, err error) {
	if p.RevIncludedRiskAssessmentResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *p.RevIncludedRiskAssessmentResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingPatient() (riskAssessments []RiskAssessment, err error) {
	if p.RevIncludedRiskAssessmentResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *p.RevIncludedRiskAssessmentResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if p.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *p.RevIncludedGroupResourcesReferencingMember
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingPatient() (careTeams []CareTeam, err error) {
	if p.RevIncludedCareTeamResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *p.RevIncludedCareTeamResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingSubject() (careTeams []CareTeam, err error) {
	if p.RevIncludedCareTeamResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *p.RevIncludedCareTeamResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingParticipant() (careTeams []CareTeam, err error) {
	if p.RevIncludedCareTeamResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *p.RevIncludedCareTeamResourcesReferencingParticipant
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if p.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *p.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingPatient() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedFamilyMemberHistoryResourcesReferencingPatient() (familyMemberHistories []FamilyMemberHistory, err error) {
	if p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded familyMemberHistories not requested")
	} else {
		familyMemberHistories = *p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingSubject() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingParticipantactor() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingParticipantactor == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingParticipantactor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingPatient() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingEnterer() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingEnterer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingSubject() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingPatient() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSubject() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPatient() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceUseStatementResourcesReferencingSubject() (deviceUseStatements []DeviceUseStatement, err error) {
	if p.RevIncludedDeviceUseStatementResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceUseStatements not requested")
	} else {
		deviceUseStatements = *p.RevIncludedDeviceUseStatementResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceUseStatementResourcesReferencingPatient() (deviceUseStatements []DeviceUseStatement, err error) {
	if p.RevIncludedDeviceUseStatementResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded deviceUseStatements not requested")
	} else {
		deviceUseStatements = *p.RevIncludedDeviceUseStatementResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingSubject() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingParticipant() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingParticipant
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingPatient() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPerformer() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingSubject() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPatient() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingPatient() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if p.RevIncludedImmunizationRecommendationResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *p.RevIncludedImmunizationRecommendationResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedBodySiteResourcesReferencingPatient() (bodySites []BodySite, err error) {
	if p.RevIncludedBodySiteResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded bodySites not requested")
	} else {
		bodySites = *p.RevIncludedBodySiteResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingPatient() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingOwner() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingOwner
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingRequester() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingRequester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingPatient() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingPayee() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingPatient() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingSubject() (specimen []Specimen, err error) {
	if p.RevIncludedSpecimenResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *p.RevIncludedSpecimenResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingPatient() (specimen []Specimen, err error) {
	if p.RevIncludedSpecimenResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *p.RevIncludedSpecimenResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingRecorder() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingAsserter() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingPatient() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if p.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *p.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingSubject() (carePlans []CarePlan, err error) {
	if p.RevIncludedCarePlanResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *p.RevIncludedCarePlanResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPatient() (carePlans []CarePlan, err error) {
	if p.RevIncludedCarePlanResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *p.RevIncludedCarePlanResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingPatient() (episodeOfCares []EpisodeOfCare, err error) {
	if p.RevIncludedEpisodeOfCareResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *p.RevIncludedEpisodeOfCareResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingSubject() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPatient() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingItem
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedListResourcesReferencingSubject() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedListResourcesReferencingPatient() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedListResourcesReferencingSource() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingSource == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingSource
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingPatient() (immunizations []Immunization, err error) {
	if p.RevIncludedImmunizationResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *p.RevIncludedImmunizationResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingRequester() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingRequester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingSubject() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingPatient() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceResourcesReferencingPatient() (devices []Device, err error) {
	if p.RevIncludedDeviceResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded devices not requested")
	} else {
		devices = *p.RevIncludedDeviceResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedVisionPrescriptionResourcesReferencingPatient() (visionPrescriptions []VisionPrescription, err error) {
	if p.RevIncludedVisionPrescriptionResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded visionPrescriptions not requested")
	} else {
		visionPrescriptions = *p.RevIncludedVisionPrescriptionResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if p.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *p.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMediaResourcesReferencingPatient() (media []Media, err error) {
	if p.RevIncludedMediaResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *p.RevIncludedMediaResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingPerformer() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingSubject() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingPatient() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSequenceResourcesReferencingPatient() (sequences []Sequence, err error) {
	if p.RevIncludedSequenceResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded sequences not requested")
	} else {
		sequences = *p.RevIncludedSequenceResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedFlagResourcesReferencingPatient() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedFlagResourcesReferencingAuthor() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if p.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *p.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingPatient() (appointmentResponses []AppointmentResponse, err error) {
	if p.RevIncludedAppointmentResponseResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *p.RevIncludedAppointmentResponseResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingRecorder() (adverseEvents []AdverseEvent, err error) {
	if p.RevIncludedAdverseEventResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *p.RevIncludedAdverseEventResourcesReferencingRecorder
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubject() (adverseEvents []AdverseEvent, err error) {
	if p.RevIncludedAdverseEventResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *p.RevIncludedAdverseEventResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedGuidanceResponseResourcesReferencingPatient() (guidanceResponses []GuidanceResponse, err error) {
	if p.RevIncludedGuidanceResponseResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded guidanceResponses not requested")
	} else {
		guidanceResponses = *p.RevIncludedGuidanceResponseResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedGuidanceResponseResourcesReferencingSubject() (guidanceResponses []GuidanceResponse, err error) {
	if p.RevIncludedGuidanceResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded guidanceResponses not requested")
	} else {
		guidanceResponses = *p.RevIncludedGuidanceResponseResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedObservationResourcesReferencingSubject() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPatient() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPerformer() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPerformer() (medicationAdministrations []MedicationAdministration, err error) {
	if p.RevIncludedMedicationAdministrationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.RevIncludedMedicationAdministrationResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingSubject() (medicationAdministrations []MedicationAdministration, err error) {
	if p.RevIncludedMedicationAdministrationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.RevIncludedMedicationAdministrationResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPatient() (medicationAdministrations []MedicationAdministration, err error) {
	if p.RevIncludedMedicationAdministrationResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.RevIncludedMedicationAdministrationResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSubject() (medicationStatements []MedicationStatement, err error) {
	if p.RevIncludedMedicationStatementResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *p.RevIncludedMedicationStatementResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingPatient() (medicationStatements []MedicationStatement, err error) {
	if p.RevIncludedMedicationStatementResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *p.RevIncludedMedicationStatementResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSource() (medicationStatements []MedicationStatement, err error) {
	if p.RevIncludedMedicationStatementResourcesReferencingSource == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *p.RevIncludedMedicationStatementResourcesReferencingSource
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRequester() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingRequester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSubject() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingPatient() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedBasicResourcesReferencingPatient() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedBasicResourcesReferencingAuthor() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedClaimResponseResourcesReferencingPatient() (claimResponses []ClaimResponse, err error) {
	if p.RevIncludedClaimResponseResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded claimResponses not requested")
	} else {
		claimResponses = *p.RevIncludedClaimResponseResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEligibilityRequestResourcesReferencingPatient() (eligibilityRequests []EligibilityRequest, err error) {
	if p.RevIncludedEligibilityRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded eligibilityRequests not requested")
	} else {
		eligibilityRequests = *p.RevIncludedEligibilityRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPerformer() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingReceiver() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingReceiver
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingSubject() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPatient() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingSubject() (diagnosticReports []DiagnosticReport, err error) {
	if p.RevIncludedDiagnosticReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *p.RevIncludedDiagnosticReportResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingPatient() (diagnosticReports []DiagnosticReport, err error) {
	if p.RevIncludedDiagnosticReportResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *p.RevIncludedDiagnosticReportResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedNutritionOrderResourcesReferencingPatient() (nutritionOrders []NutritionOrder, err error) {
	if p.RevIncludedNutritionOrderResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded nutritionOrders not requested")
	} else {
		nutritionOrders = *p.RevIncludedNutritionOrderResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingAgent() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingAgent
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingPatientPath1() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingPatientPath1 == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingPatientPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingPatientPath2() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingPatientPath2 == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingPatientPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConditionResourcesReferencingSubject() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConditionResourcesReferencingAsserter() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingAsserter
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConditionResourcesReferencingPatient() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAuthor() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAttester() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingAttester == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingAttester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingPatient() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingPatient() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPatientResourcesReferencingLink() (patients []Patient, err error) {
	if p.RevIncludedPatientResourcesReferencingLink == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *p.RevIncludedPatientResourcesReferencingLink
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingSubscriber() (coverages []Coverage, err error) {
	if p.RevIncludedCoverageResourcesReferencingSubscriber == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *p.RevIncludedCoverageResourcesReferencingSubscriber
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingPayor() (coverages []Coverage, err error) {
	if p.RevIncludedCoverageResourcesReferencingPayor == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *p.RevIncludedCoverageResourcesReferencingPayor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingBeneficiary() (coverages []Coverage, err error) {
	if p.RevIncludedCoverageResourcesReferencingBeneficiary == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *p.RevIncludedCoverageResourcesReferencingBeneficiary
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingPolicyholder() (coverages []Coverage, err error) {
	if p.RevIncludedCoverageResourcesReferencingPolicyholder == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *p.RevIncludedCoverageResourcesReferencingPolicyholder
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingPatient() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSource() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSource
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if p.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *p.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if p.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *p.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSupplyDeliveryResourcesReferencingPatient() (supplyDeliveries []SupplyDelivery, err error) {
	if p.RevIncludedSupplyDeliveryResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded supplyDeliveries not requested")
	} else {
		supplyDeliveries = *p.RevIncludedSupplyDeliveryResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSubject() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPatient() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedClaimResourcesReferencingPayee() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingPayee
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedClaimResourcesReferencingPatient() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPatientResourcesReferencedByLink != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByLink {
			rsc := (*p.IncludedPatientResourcesReferencedByLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByLink != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByLink {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPractitionerResourcesReferencedByGeneralpractitioner != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByGeneralpractitioner {
			rsc := (*p.IncludedPractitionerResourcesReferencedByGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByGeneralpractitioner != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByGeneralpractitioner {
			rsc := (*p.IncludedOrganizationResourcesReferencedByGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*p.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *PatientPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingPatient {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedReferralRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedReferralRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedReferralRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedReferralRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedReferralRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedReferralRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*p.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAccountResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingPatient {
			rsc := (*p.RevIncludedAccountResourcesReferencingPatient)[idx]
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
	if p.RevIncludedDocumentManifestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingPatient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGoalResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedGoalResourcesReferencingPatient {
			rsc := (*p.RevIncludedGoalResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*p.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedEnrollmentRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedEnrollmentRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEnrollmentRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedEnrollmentRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedConsentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingPatient {
			rsc := (*p.RevIncludedConsentResourcesReferencingPatient)[idx]
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
	if p.RevIncludedResearchSubjectResourcesReferencingIndividual != nil {
		for idx := range *p.RevIncludedResearchSubjectResourcesReferencingIndividual {
			rsc := (*p.RevIncludedResearchSubjectResourcesReferencingIndividual)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchSubjectResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedResearchSubjectResourcesReferencingPatient {
			rsc := (*p.RevIncludedResearchSubjectResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingPatient {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingPatient)[idx]
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
	if p.RevIncludedImagingManifestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImagingManifestResourcesReferencingPatient {
			rsc := (*p.RevIncludedImagingManifestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingPatient {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRelatedPersonResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRelatedPersonResourcesReferencingPatient {
			rsc := (*p.RevIncludedRelatedPersonResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingLink {
			rsc := (*p.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingPatient {
			rsc := (*p.RevIncludedPersonResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingAgent {
			rsc := (*p.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingPatient {
			rsc := (*p.RevIncludedContractResourcesReferencingPatient)[idx]
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
	if p.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingSubject {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingPatient {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingPatient)[idx]
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
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingPatient {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingSubject {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingSubject)[idx]
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
	if p.RevIncludedImagingStudyResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPatient {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient {
			rsc := (*p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingSubject {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingParticipantactor != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingParticipantactor {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingParticipantactor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingPatient {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingSubject {
			rsc := (*p.RevIncludedEncounterResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingPatient {
			rsc := (*p.RevIncludedEncounterResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSubject {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSubject)[idx]
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
	if p.RevIncludedCommunicationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPatient {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedDeviceUseStatementResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDeviceUseStatementResourcesReferencingSubject {
			rsc := (*p.RevIncludedDeviceUseStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseStatementResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceUseStatementResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceUseStatementResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingSubject {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingPatient {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
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
	if p.RevIncludedDeviceRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationRecommendationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationRecommendationResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBodySiteResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedBodySiteResourcesReferencingPatient {
			rsc := (*p.RevIncludedBodySiteResourcesReferencingPatient)[idx]
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
	if p.RevIncludedProvenanceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingPatient {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingPatient)[idx]
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
	if p.RevIncludedTaskResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingPatient {
			rsc := (*p.RevIncludedTaskResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPatient {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingPatient {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAllergyIntoleranceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingPatient {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingSubject {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPatient {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEpisodeOfCareResourcesReferencingPatient {
			rsc := (*p.RevIncludedEpisodeOfCareResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingSubject {
			rsc := (*p.RevIncludedProcedureResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPatient {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSubject {
			rsc := (*p.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedListResourcesReferencingPatient {
			rsc := (*p.RevIncludedListResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSource {
			rsc := (*p.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedVisionPrescriptionResourcesReferencingPatient {
			rsc := (*p.RevIncludedVisionPrescriptionResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*p.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingPatient {
			rsc := (*p.RevIncludedMediaResourcesReferencingPatient)[idx]
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
	if p.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSequenceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSequenceResourcesReferencingPatient {
			rsc := (*p.RevIncludedSequenceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingPatient {
			rsc := (*p.RevIncludedFlagResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAppointmentResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGuidanceResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedGuidanceResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedGuidanceResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGuidanceResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedGuidanceResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedGuidanceResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*p.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPatient {
			rsc := (*p.RevIncludedObservationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMedicationAdministrationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMedicationStatementResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCommunicationRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingSubject)[idx]
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
	if p.RevIncludedCommunicationRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedBasicResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingPatient {
			rsc := (*p.RevIncludedBasicResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*p.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClaimResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedClaimResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEligibilityRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEligibilityRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedEligibilityRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMedicationDispenseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingPatient {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingPatientPath1 != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingPatientPath1 {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingPatientPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingPatientPath2 != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingPatientPath2 {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingPatientPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingSubject {
			rsc := (*p.RevIncludedConditionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*p.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingPatient {
			rsc := (*p.RevIncludedConditionResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCompositionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingPatient {
			rsc := (*p.RevIncludedCompositionResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingPatient {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingLink {
			rsc := (*p.RevIncludedPatientResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingSubscriber != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingSubscriber {
			rsc := (*p.RevIncludedCoverageResourcesReferencingSubscriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingPayor != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingPayor {
			rsc := (*p.RevIncludedCoverageResourcesReferencingPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingBeneficiary != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingBeneficiary {
			rsc := (*p.RevIncludedCoverageResourcesReferencingBeneficiary)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingPolicyholder != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingPolicyholder {
			rsc := (*p.RevIncludedCoverageResourcesReferencingPolicyholder)[idx]
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
	if p.RevIncludedQuestionnaireResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingPatient)[idx]
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
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*p.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingPatient {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingSubject {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingPatient {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingPatient)[idx]
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
	if p.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*p.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPatient {
			rsc := (*p.RevIncludedClaimResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *PatientPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPatientResourcesReferencedByLink != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByLink {
			rsc := (*p.IncludedPatientResourcesReferencedByLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByLink != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByLink {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPractitionerResourcesReferencedByGeneralpractitioner != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByGeneralpractitioner {
			rsc := (*p.IncludedPractitionerResourcesReferencedByGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByGeneralpractitioner != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByGeneralpractitioner {
			rsc := (*p.IncludedOrganizationResourcesReferencedByGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*p.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingPatient {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedReferralRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedReferralRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedReferralRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedReferralRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedReferralRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedReferralRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*p.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAccountResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingPatient {
			rsc := (*p.RevIncludedAccountResourcesReferencingPatient)[idx]
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
	if p.RevIncludedDocumentManifestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingPatient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGoalResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedGoalResourcesReferencingPatient {
			rsc := (*p.RevIncludedGoalResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*p.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedEnrollmentRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedEnrollmentRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEnrollmentRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedEnrollmentRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedConsentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingPatient {
			rsc := (*p.RevIncludedConsentResourcesReferencingPatient)[idx]
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
	if p.RevIncludedResearchSubjectResourcesReferencingIndividual != nil {
		for idx := range *p.RevIncludedResearchSubjectResourcesReferencingIndividual {
			rsc := (*p.RevIncludedResearchSubjectResourcesReferencingIndividual)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchSubjectResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedResearchSubjectResourcesReferencingPatient {
			rsc := (*p.RevIncludedResearchSubjectResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingPatient {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingPatient)[idx]
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
	if p.RevIncludedImagingManifestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImagingManifestResourcesReferencingPatient {
			rsc := (*p.RevIncludedImagingManifestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingPatient {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRelatedPersonResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRelatedPersonResourcesReferencingPatient {
			rsc := (*p.RevIncludedRelatedPersonResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingLink {
			rsc := (*p.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingPatient {
			rsc := (*p.RevIncludedPersonResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingAgent {
			rsc := (*p.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingPatient {
			rsc := (*p.RevIncludedContractResourcesReferencingPatient)[idx]
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
	if p.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingSubject {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingPatient {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingPatient)[idx]
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
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingPatient {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingSubject {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingSubject)[idx]
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
	if p.RevIncludedImagingStudyResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPatient {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient {
			rsc := (*p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingSubject {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingParticipantactor != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingParticipantactor {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingParticipantactor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingPatient {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingSubject {
			rsc := (*p.RevIncludedEncounterResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingPatient {
			rsc := (*p.RevIncludedEncounterResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSubject {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSubject)[idx]
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
	if p.RevIncludedCommunicationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPatient {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedDeviceUseStatementResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDeviceUseStatementResourcesReferencingSubject {
			rsc := (*p.RevIncludedDeviceUseStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseStatementResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceUseStatementResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceUseStatementResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingSubject {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingPatient {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
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
	if p.RevIncludedDeviceRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationRecommendationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationRecommendationResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBodySiteResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedBodySiteResourcesReferencingPatient {
			rsc := (*p.RevIncludedBodySiteResourcesReferencingPatient)[idx]
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
	if p.RevIncludedProvenanceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingPatient {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingPatient)[idx]
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
	if p.RevIncludedTaskResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingPatient {
			rsc := (*p.RevIncludedTaskResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPatient {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingPatient {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAllergyIntoleranceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingPatient {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingSubject {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPatient {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEpisodeOfCareResourcesReferencingPatient {
			rsc := (*p.RevIncludedEpisodeOfCareResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingSubject {
			rsc := (*p.RevIncludedProcedureResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPatient {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSubject {
			rsc := (*p.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedListResourcesReferencingPatient {
			rsc := (*p.RevIncludedListResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSource {
			rsc := (*p.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedVisionPrescriptionResourcesReferencingPatient {
			rsc := (*p.RevIncludedVisionPrescriptionResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*p.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingPatient {
			rsc := (*p.RevIncludedMediaResourcesReferencingPatient)[idx]
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
	if p.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSequenceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSequenceResourcesReferencingPatient {
			rsc := (*p.RevIncludedSequenceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingPatient {
			rsc := (*p.RevIncludedFlagResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAppointmentResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGuidanceResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedGuidanceResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedGuidanceResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGuidanceResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedGuidanceResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedGuidanceResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*p.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPatient {
			rsc := (*p.RevIncludedObservationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMedicationAdministrationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMedicationStatementResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCommunicationRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingSubject)[idx]
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
	if p.RevIncludedCommunicationRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedBasicResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingPatient {
			rsc := (*p.RevIncludedBasicResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*p.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClaimResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedClaimResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEligibilityRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEligibilityRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedEligibilityRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMedicationDispenseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingPatient {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingPatientPath1 != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingPatientPath1 {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingPatientPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingPatientPath2 != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingPatientPath2 {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingPatientPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingSubject {
			rsc := (*p.RevIncludedConditionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*p.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingPatient {
			rsc := (*p.RevIncludedConditionResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCompositionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingPatient {
			rsc := (*p.RevIncludedCompositionResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingPatient {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingLink {
			rsc := (*p.RevIncludedPatientResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingSubscriber != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingSubscriber {
			rsc := (*p.RevIncludedCoverageResourcesReferencingSubscriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingPayor != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingPayor {
			rsc := (*p.RevIncludedCoverageResourcesReferencingPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingBeneficiary != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingBeneficiary {
			rsc := (*p.RevIncludedCoverageResourcesReferencingBeneficiary)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingPolicyholder != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingPolicyholder {
			rsc := (*p.RevIncludedCoverageResourcesReferencingPolicyholder)[idx]
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
	if p.RevIncludedQuestionnaireResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingPatient)[idx]
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
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*p.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingPatient {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingSubject {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingPatient {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingPatient)[idx]
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
	if p.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*p.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPatient {
			rsc := (*p.RevIncludedClaimResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
