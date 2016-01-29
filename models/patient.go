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
	CareProvider         []Reference                     `bson:"careProvider,omitempty" json:"careProvider,omitempty"`
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
	Relationship []CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name         *HumanName        `bson:"name,omitempty" json:"name,omitempty"`
	Telecom      []ContactPoint    `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address      *Address          `bson:"address,omitempty" json:"address,omitempty"`
	Gender       string            `bson:"gender,omitempty" json:"gender,omitempty"`
	Organization *Reference        `bson:"organization,omitempty" json:"organization,omitempty"`
	Period       *Period           `bson:"period,omitempty" json:"period,omitempty"`
}

type PatientAnimalComponent struct {
	Species      *CodeableConcept `bson:"species,omitempty" json:"species,omitempty"`
	Breed        *CodeableConcept `bson:"breed,omitempty" json:"breed,omitempty"`
	GenderStatus *CodeableConcept `bson:"genderStatus,omitempty" json:"genderStatus,omitempty"`
}

type PatientCommunicationComponent struct {
	Language  *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	Preferred *bool            `bson:"preferred,omitempty" json:"preferred,omitempty"`
}

type PatientLinkComponent struct {
	Other *Reference `bson:"other,omitempty" json:"other,omitempty"`
	Type  string     `bson:"type,omitempty" json:"type,omitempty"`
}

type PatientPlus struct {
	Patient                     `bson:",inline"`
	PatientPlusRelatedResources `bson:",inline"`
}

type PatientPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByLink                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByLink,omitempty"`
	IncludedPractitionerResourcesReferencedByCareprovider               *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByCareprovider,omitempty"`
	IncludedOrganizationResourcesReferencedByCareprovider               *[]Organization               `bson:"_includedOrganizationResourcesReferencedByCareprovider,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization               *[]Organization               `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                     *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingPatient                   *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingPatient,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRequester             *[]ReferralRequest            `bson:"_revIncludedReferralRequestResourcesReferencingRequester,omitempty"`
	RevIncludedReferralRequestResourcesReferencingPatient               *[]ReferralRequest            `bson:"_revIncludedReferralRequestResourcesReferencingPatient,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                       *[]Account                    `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedAccountResourcesReferencingPatient                       *[]Account                    `bson:"_revIncludedAccountResourcesReferencingPatient,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                      *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingPatient                    *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingPatient,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                     *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref           *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingSubject              *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor               *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref           *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingPatient              *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingPatient,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRecipient            *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRecipient,omitempty"`
	RevIncludedSpecimenResourcesReferencingSubject                      *[]Specimen                   `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedSpecimenResourcesReferencingPatient                      *[]Specimen                   `bson:"_revIncludedSpecimenResourcesReferencingPatient,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingRecorder           *[]AllergyIntolerance         `bson:"_revIncludedAllergyIntoleranceResourcesReferencingRecorder,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingReporter           *[]AllergyIntolerance         `bson:"_revIncludedAllergyIntoleranceResourcesReferencingReporter,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingPatient            *[]AllergyIntolerance         `bson:"_revIncludedAllergyIntoleranceResourcesReferencingPatient,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                    *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedCarePlanResourcesReferencingSubject                      *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingSubject,omitempty"`
	RevIncludedCarePlanResourcesReferencingParticipant                  *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingParticipant,omitempty"`
	RevIncludedCarePlanResourcesReferencingPatient                      *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingPatient,omitempty"`
	RevIncludedGoalResourcesReferencingPatient                          *[]Goal                       `bson:"_revIncludedGoalResourcesReferencingPatient,omitempty"`
	RevIncludedGoalResourcesReferencingSubject                          *[]Goal                       `bson:"_revIncludedGoalResourcesReferencingSubject,omitempty"`
	RevIncludedEnrollmentRequestResourcesReferencingSubject             *[]EnrollmentRequest          `bson:"_revIncludedEnrollmentRequestResourcesReferencingSubject,omitempty"`
	RevIncludedEnrollmentRequestResourcesReferencingPatient             *[]EnrollmentRequest          `bson:"_revIncludedEnrollmentRequestResourcesReferencingPatient,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingPatient                 *[]EpisodeOfCare              `bson:"_revIncludedEpisodeOfCareResourcesReferencingPatient,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                   *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureResourcesReferencingSubject                     *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureResourcesReferencingPatient                     *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingPatient,omitempty"`
	RevIncludedListResourcesReferencingItem                             *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                          *[]List                       `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingPatient                          *[]List                       `bson:"_revIncludedListResourcesReferencingPatient,omitempty"`
	RevIncludedListResourcesReferencingSource                           *[]List                       `bson:"_revIncludedListResourcesReferencingSource,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingSubject             *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingPatient             *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingPatient,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor              *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref          *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingSubject                         *[]Order                      `bson:"_revIncludedOrderResourcesReferencingSubject,omitempty"`
	RevIncludedOrderResourcesReferencingPatient                         *[]Order                      `bson:"_revIncludedOrderResourcesReferencingPatient,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                          *[]Order                      `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedImmunizationResourcesReferencingPatient                  *[]Immunization               `bson:"_revIncludedImmunizationResourcesReferencingPatient,omitempty"`
	RevIncludedDeviceResourcesReferencingPatient                        *[]Device                     `bson:"_revIncludedDeviceResourcesReferencingPatient,omitempty"`
	RevIncludedVisionPrescriptionResourcesReferencingPatient            *[]VisionPrescription         `bson:"_revIncludedVisionPrescriptionResourcesReferencingPatient,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                         *[]Media                      `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedMediaResourcesReferencingPatient                         *[]Media                      `bson:"_revIncludedMediaResourcesReferencingPatient,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingPerformer            *[]ProcedureRequest           `bson:"_revIncludedProcedureRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingSubject              *[]ProcedureRequest           `bson:"_revIncludedProcedureRequestResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingPatient              *[]ProcedureRequest           `bson:"_revIncludedProcedureRequestResourcesReferencingPatient,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingOrderer              *[]ProcedureRequest           `bson:"_revIncludedProcedureRequestResourcesReferencingOrderer,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingSubject              *[]DeviceUseRequest           `bson:"_revIncludedDeviceUseRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingPatient              *[]DeviceUseRequest           `bson:"_revIncludedDeviceUseRequestResourcesReferencingPatient,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                          *[]Flag                       `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedFlagResourcesReferencingPatient                          *[]Flag                       `bson:"_revIncludedFlagResourcesReferencingPatient,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                           *[]Flag                       `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedRelatedPersonResourcesReferencingPatient                 *[]RelatedPerson              `bson:"_revIncludedRelatedPersonResourcesReferencingPatient,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingPatient                 *[]SupplyRequest              `bson:"_revIncludedSupplyRequestResourcesReferencingPatient,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingSource                  *[]SupplyRequest              `bson:"_revIncludedSupplyRequestResourcesReferencingSource,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor             *[]AppointmentResponse        `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingPatient           *[]AppointmentResponse        `bson:"_revIncludedAppointmentResponseResourcesReferencingPatient,omitempty"`
	RevIncludedObservationResourcesReferencingSubject                   *[]Observation                `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingPatient                   *[]Observation                `bson:"_revIncludedObservationResourcesReferencingPatient,omitempty"`
	RevIncludedObservationResourcesReferencingPerformer                 *[]Observation                `bson:"_revIncludedObservationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPractitioner *[]MedicationAdministration   `bson:"_revIncludedMedicationAdministrationResourcesReferencingPractitioner,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPatient      *[]MedicationAdministration   `bson:"_revIncludedMedicationAdministrationResourcesReferencingPatient,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingPatient           *[]MedicationStatement        `bson:"_revIncludedMedicationStatementResourcesReferencingPatient,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSource            *[]MedicationStatement        `bson:"_revIncludedMedicationStatementResourcesReferencingSource,omitempty"`
	RevIncludedPersonResourcesReferencingLink                           *[]Person                     `bson:"_revIncludedPersonResourcesReferencingLink,omitempty"`
	RevIncludedPersonResourcesReferencingPatient                        *[]Person                     `bson:"_revIncludedPersonResourcesReferencingPatient,omitempty"`
	RevIncludedContractResourcesReferencingActor                        *[]Contract                   `bson:"_revIncludedContractResourcesReferencingActor,omitempty"`
	RevIncludedContractResourcesReferencingSubject                      *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingPatient                      *[]Contract                   `bson:"_revIncludedContractResourcesReferencingPatient,omitempty"`
	RevIncludedContractResourcesReferencingSigner                       *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSigner,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRequester        *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSubject          *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender           *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingPatient          *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingPatient,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient        *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingSubject                *[]RiskAssessment             `bson:"_revIncludedRiskAssessmentResourcesReferencingSubject,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingPatient                *[]RiskAssessment             `bson:"_revIncludedRiskAssessmentResourcesReferencingPatient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                         *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingPatient                         *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingPatient,omitempty"`
	RevIncludedBasicResourcesReferencingAuthor                          *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingAuthor,omitempty"`
	RevIncludedGroupResourcesReferencingMember                          *[]Group                      `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingReceiver           *[]MedicationDispense         `bson:"_revIncludedMedicationDispenseResourcesReferencingReceiver,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPatient            *[]MedicationDispense         `bson:"_revIncludedMedicationDispenseResourcesReferencingPatient,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject              *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingPatient              *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingPatient,omitempty"`
	RevIncludedImagingStudyResourcesReferencingPatient                  *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingPatient,omitempty"`
	RevIncludedImagingObjectSelectionResourcesReferencingAuthor         *[]ImagingObjectSelection     `bson:"_revIncludedImagingObjectSelectionResourcesReferencingAuthor,omitempty"`
	RevIncludedImagingObjectSelectionResourcesReferencingPatient        *[]ImagingObjectSelection     `bson:"_revIncludedImagingObjectSelectionResourcesReferencingPatient,omitempty"`
	RevIncludedFamilyMemberHistoryResourcesReferencingPatient           *[]FamilyMemberHistory        `bson:"_revIncludedFamilyMemberHistoryResourcesReferencingPatient,omitempty"`
	RevIncludedNutritionOrderResourcesReferencingPatient                *[]NutritionOrder             `bson:"_revIncludedNutritionOrderResourcesReferencingPatient,omitempty"`
	RevIncludedEncounterResourcesReferencingPatient                     *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingPatient,omitempty"`
	RevIncludedAuditEventResourcesReferencingParticipant                *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingParticipant,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference                  *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedAuditEventResourcesReferencingPatientPath1               *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingPatientPath1,omitempty"`
	RevIncludedAuditEventResourcesReferencingPatientPath2               *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingPatientPath2,omitempty"`
	RevIncludedMedicationOrderResourcesReferencingPatient               *[]MedicationOrder            `bson:"_revIncludedMedicationOrderResourcesReferencingPatient,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender                  *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingSubject                 *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationResourcesReferencingPatient                 *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingPatient,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient               *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedConditionResourcesReferencingAsserter                    *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingAsserter,omitempty"`
	RevIncludedConditionResourcesReferencingPatient                     *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingPatient,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                   *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor                    *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingAttester                  *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingAttester,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                     *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedCompositionResourcesReferencingPatient                   *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingPatient,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingPatient                 *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingPatient,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated              *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingSubject               *[]DiagnosticOrder            `bson:"_revIncludedDiagnosticOrderResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingPatient               *[]DiagnosticOrder            `bson:"_revIncludedDiagnosticOrderResourcesReferencingPatient,omitempty"`
	RevIncludedPatientResourcesReferencingLink                          *[]Patient                    `bson:"_revIncludedPatientResourcesReferencingLink,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment             *[]OrderResponse              `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject         *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor          *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingPatient         *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingPatient,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSource          *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSource,omitempty"`
	RevIncludedDeviceUseStatementResourcesReferencingSubject            *[]DeviceUseStatement         `bson:"_revIncludedDeviceUseStatementResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceUseStatementResourcesReferencingPatient            *[]DeviceUseStatement         `bson:"_revIncludedDeviceUseStatementResourcesReferencingPatient,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest               *[]ProcessResponse            `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                        *[]Schedule                   `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedSupplyDeliveryResourcesReferencingPatient                *[]SupplyDelivery             `bson:"_revIncludedSupplyDeliveryResourcesReferencingPatient,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger            *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPatient            *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingPatient,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                    *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedClaimResourcesReferencingPatient                         *[]Claim                      `bson:"_revIncludedClaimResourcesReferencingPatient,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingPatient    *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingPatient,omitempty"`
	RevIncludedBodySiteResourcesReferencingPatient                      *[]BodySite                   `bson:"_revIncludedBodySiteResourcesReferencingPatient,omitempty"`
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

func (p *PatientPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByCareprovider() (practitioners []Practitioner, err error) {
	if p.IncludedPractitionerResourcesReferencedByCareprovider == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *p.IncludedPractitionerResourcesReferencedByCareprovider
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByCareprovider() (organizations []Organization, err error) {
	if p.IncludedOrganizationResourcesReferencedByCareprovider == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *p.IncludedOrganizationResourcesReferencedByCareprovider
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

func (p *PatientPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingReporter() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingReporter == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingReporter
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

func (p *PatientPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingParticipant() (carePlans []CarePlan, err error) {
	if p.RevIncludedCarePlanResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *p.RevIncludedCarePlanResourcesReferencingParticipant
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

func (p *PatientPlusRelatedResources) GetRevIncludedOrderResourcesReferencingSubject() (orders []Order, err error) {
	if p.RevIncludedOrderResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *p.RevIncludedOrderResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedOrderResourcesReferencingPatient() (orders []Order, err error) {
	if p.RevIncludedOrderResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *p.RevIncludedOrderResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if p.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *p.RevIncludedOrderResourcesReferencingDetail
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

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingSubject() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingSubject
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

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingOrderer() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingOrderer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingOrderer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingSubject() (deviceUseRequests []DeviceUseRequest, err error) {
	if p.RevIncludedDeviceUseRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *p.RevIncludedDeviceUseRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingPatient() (deviceUseRequests []DeviceUseRequest, err error) {
	if p.RevIncludedDeviceUseRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *p.RevIncludedDeviceUseRequestResourcesReferencingPatient
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

func (p *PatientPlusRelatedResources) GetRevIncludedRelatedPersonResourcesReferencingPatient() (relatedPeople []RelatedPerson, err error) {
	if p.RevIncludedRelatedPersonResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded relatedPeople not requested")
	} else {
		relatedPeople = *p.RevIncludedRelatedPersonResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingPatient() (supplyRequests []SupplyRequest, err error) {
	if p.RevIncludedSupplyRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *p.RevIncludedSupplyRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingSource() (supplyRequests []SupplyRequest, err error) {
	if p.RevIncludedSupplyRequestResourcesReferencingSource == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *p.RevIncludedSupplyRequestResourcesReferencingSource
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

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPractitioner() (medicationAdministrations []MedicationAdministration, err error) {
	if p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner
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

func (p *PatientPlusRelatedResources) GetRevIncludedContractResourcesReferencingActor() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingActor == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingActor
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

func (p *PatientPlusRelatedResources) GetRevIncludedContractResourcesReferencingPatient() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingPatient
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

func (p *PatientPlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if p.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *p.RevIncludedGroupResourcesReferencingMember
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

func (p *PatientPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingPatient() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImagingObjectSelectionResourcesReferencingAuthor() (imagingObjectSelections []ImagingObjectSelection, err error) {
	if p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded imagingObjectSelections not requested")
	} else {
		imagingObjectSelections = *p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImagingObjectSelectionResourcesReferencingPatient() (imagingObjectSelections []ImagingObjectSelection, err error) {
	if p.RevIncludedImagingObjectSelectionResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded imagingObjectSelections not requested")
	} else {
		imagingObjectSelections = *p.RevIncludedImagingObjectSelectionResourcesReferencingPatient
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

func (p *PatientPlusRelatedResources) GetRevIncludedNutritionOrderResourcesReferencingPatient() (nutritionOrders []NutritionOrder, err error) {
	if p.RevIncludedNutritionOrderResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded nutritionOrders not requested")
	} else {
		nutritionOrders = *p.RevIncludedNutritionOrderResourcesReferencingPatient
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

func (p *PatientPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingParticipant() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingParticipant
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingReference
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

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationOrderResourcesReferencingPatient() (medicationOrders []MedicationOrder, err error) {
	if p.RevIncludedMedicationOrderResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded medicationOrders not requested")
	} else {
		medicationOrders = *p.RevIncludedMedicationOrderResourcesReferencingPatient
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

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSubject() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingSubject
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

func (p *PatientPlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingSubject() (diagnosticOrders []DiagnosticOrder, err error) {
	if p.RevIncludedDiagnosticOrderResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *p.RevIncludedDiagnosticOrderResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingPatient() (diagnosticOrders []DiagnosticOrder, err error) {
	if p.RevIncludedDiagnosticOrderResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *p.RevIncludedDiagnosticOrderResourcesReferencingPatient
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

func (p *PatientPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if p.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *p.RevIncludedOrderResponseResourcesReferencingFulfillment
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

func (p *PatientPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingTrigger
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

func (p *PatientPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingData
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

func (p *PatientPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPatientResourcesReferencedByLink != nil {
		for _, r := range *p.IncludedPatientResourcesReferencedByLink {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedPractitionerResourcesReferencedByCareprovider != nil {
		for _, r := range *p.IncludedPractitionerResourcesReferencedByCareprovider {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedOrganizationResourcesReferencedByCareprovider != nil {
		for _, r := range *p.IncludedOrganizationResourcesReferencedByCareprovider {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *p.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (p *PatientPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedAppointmentResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedAppointmentResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for _, r := range *p.RevIncludedReferralRequestResourcesReferencingRequester {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedReferralRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedAccountResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAccountResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedAccountResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for _, r := range *p.RevIncludedProvenanceResourcesReferencingAgent {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedProvenanceResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedSpecimenResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedSpecimenResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder != nil {
		for _, r := range *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingReporter != nil {
		for _, r := range *p.RevIncludedAllergyIntoleranceResourcesReferencingReporter {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedAllergyIntoleranceResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingParticipant != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedGoalResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedGoalResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedGoalResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedGoalResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedEnrollmentRequestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedEnrollmentRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedEpisodeOfCareResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedProcedureResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedProcedureResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedImmunizationResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDeviceResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDeviceResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedVisionPrescriptionResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedMediaResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMediaResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedMediaResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingOrderer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDeviceUseRequestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDeviceUseRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedFlagResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFlagResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedFlagResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFlagResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedFlagResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedRelatedPersonResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedRelatedPersonResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedSupplyRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedSupplyRequestResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedAppointmentResponseResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedAppointmentResponseResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedObservationResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedObservationResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedObservationResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedObservationResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedObservationResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedObservationResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner != nil {
		for _, r := range *p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedMedicationAdministrationResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedMedicationStatementResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedMedicationStatementResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for _, r := range *p.RevIncludedPersonResourcesReferencingLink {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPersonResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedPersonResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedContractResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedContractResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedContractResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedContractResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedContractResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedContractResourcesReferencingSigner != nil {
		for _, r := range *p.RevIncludedContractResourcesReferencingSigner {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingRequester {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingSender {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedRiskAssessmentResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedRiskAssessmentResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBasicResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedBasicResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedBasicResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedGroupResourcesReferencingMember != nil {
		for _, r := range *p.RevIncludedGroupResourcesReferencingMember {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for _, r := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedMedicationDispenseResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDiagnosticReportResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDiagnosticReportResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedImagingStudyResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImagingObjectSelectionResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedImagingObjectSelectionResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedNutritionOrderResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedEncounterResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for _, r := range *p.RevIncludedAuditEventResourcesReferencingParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *p.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingPatientPath1 != nil {
		for _, r := range *p.RevIncludedAuditEventResourcesReferencingPatientPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingPatientPath2 != nil {
		for _, r := range *p.RevIncludedAuditEventResourcesReferencingPatientPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationOrderResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedMedicationOrderResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for _, r := range *p.RevIncludedCommunicationResourcesReferencingSender {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedCommunicationResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedCommunicationResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedCommunicationResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for _, r := range *p.RevIncludedConditionResourcesReferencingAsserter {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedConditionResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedConditionResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAttester != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingAttester {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDetectedIssueResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDiagnosticOrderResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPatientResourcesReferencingLink != nil {
		for _, r := range *p.RevIncludedPatientResourcesReferencingLink {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *p.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDeviceUseStatementResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDeviceUseStatementResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDeviceUseStatementResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDeviceUseStatementResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedScheduleResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedSupplyDeliveryResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *p.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedClinicalImpressionResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClaimResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedClaimResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedImmunizationRecommendationResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBodySiteResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedBodySiteResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (p *PatientPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPatientResourcesReferencedByLink != nil {
		for _, r := range *p.IncludedPatientResourcesReferencedByLink {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedPractitionerResourcesReferencedByCareprovider != nil {
		for _, r := range *p.IncludedPractitionerResourcesReferencedByCareprovider {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedOrganizationResourcesReferencedByCareprovider != nil {
		for _, r := range *p.IncludedOrganizationResourcesReferencedByCareprovider {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *p.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedAppointmentResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedAppointmentResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for _, r := range *p.RevIncludedReferralRequestResourcesReferencingRequester {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedReferralRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedAccountResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAccountResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedAccountResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for _, r := range *p.RevIncludedProvenanceResourcesReferencingAgent {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedProvenanceResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedSpecimenResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedSpecimenResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder != nil {
		for _, r := range *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingReporter != nil {
		for _, r := range *p.RevIncludedAllergyIntoleranceResourcesReferencingReporter {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedAllergyIntoleranceResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingParticipant != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedGoalResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedGoalResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedGoalResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedGoalResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedEnrollmentRequestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedEnrollmentRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedEpisodeOfCareResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedProcedureResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedProcedureResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedImmunizationResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDeviceResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDeviceResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedVisionPrescriptionResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedMediaResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMediaResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedMediaResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingOrderer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDeviceUseRequestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDeviceUseRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedFlagResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFlagResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedFlagResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFlagResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedFlagResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedRelatedPersonResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedRelatedPersonResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedSupplyRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedSupplyRequestResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedAppointmentResponseResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedAppointmentResponseResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedObservationResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedObservationResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedObservationResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedObservationResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedObservationResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedObservationResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner != nil {
		for _, r := range *p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedMedicationAdministrationResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedMedicationStatementResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedMedicationStatementResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for _, r := range *p.RevIncludedPersonResourcesReferencingLink {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPersonResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedPersonResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedContractResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedContractResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedContractResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedContractResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedContractResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedContractResourcesReferencingSigner != nil {
		for _, r := range *p.RevIncludedContractResourcesReferencingSigner {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingRequester {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingSender {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedRiskAssessmentResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedRiskAssessmentResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBasicResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedBasicResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedBasicResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedGroupResourcesReferencingMember != nil {
		for _, r := range *p.RevIncludedGroupResourcesReferencingMember {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for _, r := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedMedicationDispenseResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDiagnosticReportResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDiagnosticReportResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedImagingStudyResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImagingObjectSelectionResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedImagingObjectSelectionResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedNutritionOrderResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedEncounterResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for _, r := range *p.RevIncludedAuditEventResourcesReferencingParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *p.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingPatientPath1 != nil {
		for _, r := range *p.RevIncludedAuditEventResourcesReferencingPatientPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingPatientPath2 != nil {
		for _, r := range *p.RevIncludedAuditEventResourcesReferencingPatientPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationOrderResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedMedicationOrderResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for _, r := range *p.RevIncludedCommunicationResourcesReferencingSender {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedCommunicationResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedCommunicationResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedCommunicationResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for _, r := range *p.RevIncludedConditionResourcesReferencingAsserter {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedConditionResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedConditionResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAttester != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingAttester {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDetectedIssueResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDiagnosticOrderResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPatientResourcesReferencingLink != nil {
		for _, r := range *p.RevIncludedPatientResourcesReferencingLink {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *p.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDeviceUseStatementResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedDeviceUseStatementResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDeviceUseStatementResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedDeviceUseStatementResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedScheduleResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedSupplyDeliveryResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *p.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedClinicalImpressionResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClaimResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedClaimResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedImmunizationRecommendationResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBodySiteResourcesReferencingPatient != nil {
		for _, r := range *p.RevIncludedBodySiteResourcesReferencingPatient {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
