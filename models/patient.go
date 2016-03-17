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
		for idx := range *p.IncludedPatientResourcesReferencedByLink {
			rsc := (*p.IncludedPatientResourcesReferencedByLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPractitionerResourcesReferencedByCareprovider != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByCareprovider {
			rsc := (*p.IncludedPractitionerResourcesReferencedByCareprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByCareprovider != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByCareprovider {
			rsc := (*p.IncludedOrganizationResourcesReferencedByCareprovider)[idx]
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
	if p.RevIncludedAllergyIntoleranceResourcesReferencingReporter != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingReporter {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingReporter)[idx]
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
	if p.RevIncludedCarePlanResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingParticipant {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPatient {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPatient)[idx]
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
	if p.RevIncludedOrderResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedOrderResourcesReferencingSubject {
			rsc := (*p.RevIncludedOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedOrderResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *p.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*p.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingOrderer {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingOrderer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedRelatedPersonResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRelatedPersonResourcesReferencingPatient {
			rsc := (*p.RevIncludedRelatedPersonResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingSource {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingSource)[idx]
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
	if p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingActor {
			rsc := (*p.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingPatient {
			rsc := (*p.RevIncludedContractResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSigner {
			rsc := (*p.RevIncludedContractResourcesReferencingSigner)[idx]
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
	if p.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingMember {
			rsc := (*p.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingReceiver)[idx]
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
	if p.RevIncludedImagingStudyResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPatient {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			rsc := (*p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingObjectSelectionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImagingObjectSelectionResourcesReferencingPatient {
			rsc := (*p.RevIncludedImagingObjectSelectionResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient {
			rsc := (*p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingPatient {
			rsc := (*p.RevIncludedEncounterResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingParticipant {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if p.RevIncludedMedicationOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationOrderResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSubject {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSubject)[idx]
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
	if p.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			rsc := (*p.RevIncludedDiagnosticOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDiagnosticOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedDiagnosticOrderResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingLink {
			rsc := (*p.RevIncludedPatientResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *p.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*p.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if p.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingPatient {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPatient {
			rsc := (*p.RevIncludedClaimResourcesReferencingPatient)[idx]
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
	if p.IncludedPractitionerResourcesReferencedByCareprovider != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByCareprovider {
			rsc := (*p.IncludedPractitionerResourcesReferencedByCareprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByCareprovider != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByCareprovider {
			rsc := (*p.IncludedOrganizationResourcesReferencedByCareprovider)[idx]
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
	if p.RevIncludedAllergyIntoleranceResourcesReferencingReporter != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingReporter {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingReporter)[idx]
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
	if p.RevIncludedCarePlanResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingParticipant {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPatient {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPatient)[idx]
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
	if p.RevIncludedOrderResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedOrderResourcesReferencingSubject {
			rsc := (*p.RevIncludedOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedOrderResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *p.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*p.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingOrderer {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingOrderer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedRelatedPersonResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRelatedPersonResourcesReferencingPatient {
			rsc := (*p.RevIncludedRelatedPersonResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingSource {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingSource)[idx]
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
	if p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingActor {
			rsc := (*p.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingPatient {
			rsc := (*p.RevIncludedContractResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSigner {
			rsc := (*p.RevIncludedContractResourcesReferencingSigner)[idx]
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
	if p.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingMember {
			rsc := (*p.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingReceiver)[idx]
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
	if p.RevIncludedImagingStudyResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPatient {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			rsc := (*p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingObjectSelectionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImagingObjectSelectionResourcesReferencingPatient {
			rsc := (*p.RevIncludedImagingObjectSelectionResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient {
			rsc := (*p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingPatient {
			rsc := (*p.RevIncludedEncounterResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingParticipant {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if p.RevIncludedMedicationOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationOrderResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSubject {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSubject)[idx]
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
	if p.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			rsc := (*p.RevIncludedDiagnosticOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDiagnosticOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedDiagnosticOrderResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingLink {
			rsc := (*p.RevIncludedPatientResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *p.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*p.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if p.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingPatient {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPatient {
			rsc := (*p.RevIncludedClaimResourcesReferencingPatient)[idx]
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
	return resourceMap
}
