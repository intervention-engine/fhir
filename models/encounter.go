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

type Encounter struct {
	DomainResource   `bson:",inline"`
	Identifier       []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status           string                             `bson:"status,omitempty" json:"status,omitempty"`
	StatusHistory    []EncounterStatusHistoryComponent  `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Class            *Coding                            `bson:"class,omitempty" json:"class,omitempty"`
	ClassHistory     []EncounterClassHistoryComponent   `bson:"classHistory,omitempty" json:"classHistory,omitempty"`
	Type             []CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	Priority         *CodeableConcept                   `bson:"priority,omitempty" json:"priority,omitempty"`
	Subject          *Reference                         `bson:"subject,omitempty" json:"subject,omitempty"`
	EpisodeOfCare    []Reference                        `bson:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty"`
	IncomingReferral []Reference                        `bson:"incomingReferral,omitempty" json:"incomingReferral,omitempty"`
	Participant      []EncounterParticipantComponent    `bson:"participant,omitempty" json:"participant,omitempty"`
	Appointment      *Reference                         `bson:"appointment,omitempty" json:"appointment,omitempty"`
	Period           *Period                            `bson:"period,omitempty" json:"period,omitempty"`
	Length           *Quantity                          `bson:"length,omitempty" json:"length,omitempty"`
	Reason           []CodeableConcept                  `bson:"reason,omitempty" json:"reason,omitempty"`
	Diagnosis        []EncounterDiagnosisComponent      `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Account          []Reference                        `bson:"account,omitempty" json:"account,omitempty"`
	Hospitalization  *EncounterHospitalizationComponent `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Location         []EncounterLocationComponent       `bson:"location,omitempty" json:"location,omitempty"`
	ServiceProvider  *Reference                         `bson:"serviceProvider,omitempty" json:"serviceProvider,omitempty"`
	PartOf           *Reference                         `bson:"partOf,omitempty" json:"partOf,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Encounter) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Encounter"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Encounter), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Encounter) GetBSON() (interface{}, error) {
	x.ResourceType = "Encounter"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "encounter" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type encounter Encounter

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Encounter) UnmarshalJSON(data []byte) (err error) {
	x2 := encounter{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Encounter(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Encounter) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Encounter"
	} else if x.ResourceType != "Encounter" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Encounter, instead received %s", x.ResourceType))
	}
	return nil
}

type EncounterStatusHistoryComponent struct {
	BackboneElement `bson:",inline"`
	Status          string  `bson:"status,omitempty" json:"status,omitempty"`
	Period          *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type EncounterClassHistoryComponent struct {
	BackboneElement `bson:",inline"`
	Class           *Coding `bson:"class,omitempty" json:"class,omitempty"`
	Period          *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type EncounterParticipantComponent struct {
	BackboneElement `bson:",inline"`
	Type            []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Period          *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Individual      *Reference        `bson:"individual,omitempty" json:"individual,omitempty"`
}

type EncounterDiagnosisComponent struct {
	BackboneElement `bson:",inline"`
	Condition       *Reference       `bson:"condition,omitempty" json:"condition,omitempty"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Rank            *uint32          `bson:"rank,omitempty" json:"rank,omitempty"`
}

type EncounterHospitalizationComponent struct {
	BackboneElement        `bson:",inline"`
	PreAdmissionIdentifier *Identifier       `bson:"preAdmissionIdentifier,omitempty" json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference        `bson:"origin,omitempty" json:"origin,omitempty"`
	AdmitSource            *CodeableConcept  `bson:"admitSource,omitempty" json:"admitSource,omitempty"`
	ReAdmission            *CodeableConcept  `bson:"reAdmission,omitempty" json:"reAdmission,omitempty"`
	DietPreference         []CodeableConcept `bson:"dietPreference,omitempty" json:"dietPreference,omitempty"`
	SpecialCourtesy        []CodeableConcept `bson:"specialCourtesy,omitempty" json:"specialCourtesy,omitempty"`
	SpecialArrangement     []CodeableConcept `bson:"specialArrangement,omitempty" json:"specialArrangement,omitempty"`
	Destination            *Reference        `bson:"destination,omitempty" json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept  `bson:"dischargeDisposition,omitempty" json:"dischargeDisposition,omitempty"`
}

type EncounterLocationComponent struct {
	BackboneElement `bson:",inline"`
	Location        *Reference `bson:"location,omitempty" json:"location,omitempty"`
	Status          string     `bson:"status,omitempty" json:"status,omitempty"`
	Period          *Period    `bson:"period,omitempty" json:"period,omitempty"`
}

type EncounterPlus struct {
	Encounter                     `bson:",inline"`
	EncounterPlusRelatedResources `bson:",inline"`
}

type EncounterPlusRelatedResources struct {
	IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare         *[]EpisodeOfCare            `bson:"_includedEpisodeOfCareResourcesReferencedByEpisodeofcare,omitempty"`
	IncludedReferralRequestResourcesReferencedByIncomingreferral    *[]ReferralRequest          `bson:"_includedReferralRequestResourcesReferencedByIncomingreferral,omitempty"`
	IncludedPractitionerResourcesReferencedByPractitioner           *[]Practitioner             `bson:"_includedPractitionerResourcesReferencedByPractitioner,omitempty"`
	IncludedGroupResourcesReferencedBySubject                       *[]Group                    `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                     *[]Patient                  `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedConditionResourcesReferencedByDiagnosis                 *[]Condition                `bson:"_includedConditionResourcesReferencedByDiagnosis,omitempty"`
	IncludedProcedureResourcesReferencedByDiagnosis                 *[]Procedure                `bson:"_includedProcedureResourcesReferencedByDiagnosis,omitempty"`
	IncludedAppointmentResourcesReferencedByAppointment             *[]Appointment              `bson:"_includedAppointmentResourcesReferencedByAppointment,omitempty"`
	IncludedEncounterResourcesReferencedByPartof                    *[]Encounter                `bson:"_includedEncounterResourcesReferencedByPartof,omitempty"`
	IncludedPractitionerResourcesReferencedByParticipant            *[]Practitioner             `bson:"_includedPractitionerResourcesReferencedByParticipant,omitempty"`
	IncludedRelatedPersonResourcesReferencedByParticipant           *[]RelatedPerson            `bson:"_includedRelatedPersonResourcesReferencedByParticipant,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient                  `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedLocationResourcesReferencedByLocation                   *[]Location                 `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	IncludedOrganizationResourcesReferencedByServiceprovider        *[]Organization             `bson:"_includedOrganizationResourcesReferencedByServiceprovider,omitempty"`
	RevIncludedReferralRequestResourcesReferencingEncounter         *[]ReferralRequest          `bson:"_revIncludedReferralRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedReferralRequestResourcesReferencingContext           *[]ReferralRequest          `bson:"_revIncludedReferralRequestResourcesReferencingContext,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref       *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref       *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                 *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                 *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                 *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom               *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor               *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1            *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2            *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingEncounter       *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingEncounter,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref      *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingSubject                  *[]Contract                 `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                *[]Contract                 `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingEncounter          *[]RiskAssessment           `bson:"_revIncludedRiskAssessmentResourcesReferencingEncounter,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest             *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse            *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedCareTeamResourcesReferencingContext                  *[]CareTeam                 `bson:"_revIncludedCareTeamResourcesReferencingContext,omitempty"`
	RevIncludedCareTeamResourcesReferencingEncounter                *[]CareTeam                 `bson:"_revIncludedCareTeamResourcesReferencingEncounter,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource      *[]ImplementationGuide      `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingContext              *[]ImagingStudy             `bson:"_revIncludedImagingStudyResourcesReferencingContext,omitempty"`
	RevIncludedChargeItemResourcesReferencingContext                *[]ChargeItem               `bson:"_revIncludedChargeItemResourcesReferencingContext,omitempty"`
	RevIncludedEncounterResourcesReferencingPartof                  *[]Encounter                `bson:"_revIncludedEncounterResourcesReferencingPartof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor       *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom     *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor     *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof      *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson       *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof              *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingEncounter           *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingEncounter,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon             *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingContext             *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingContext,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor      *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom    *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor    *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof     *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedRequestGroupResourcesReferencingEncounter            *[]RequestGroup             `bson:"_revIncludedRequestGroupResourcesReferencingEncounter,omitempty"`
	RevIncludedRequestGroupResourcesReferencingContext              *[]RequestGroup             `bson:"_revIncludedRequestGroupResourcesReferencingContext,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition           *[]RequestGroup             `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingEncounter           *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon             *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest        *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus               *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref              *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                 *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                      *[]Task                     `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                        *[]Task                     `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                      *[]Task                     `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedTaskResourcesReferencingContext                      *[]Task                     `bson:"_revIncludedTaskResourcesReferencingContext,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingEncounter    *[]ExplanationOfBenefit     `bson:"_revIncludedExplanationOfBenefitResourcesReferencingEncounter,omitempty"`
	RevIncludedCarePlanResourcesReferencingEncounter                *[]CarePlan                 `bson:"_revIncludedCarePlanResourcesReferencingEncounter,omitempty"`
	RevIncludedCarePlanResourcesReferencingContext                  *[]CarePlan                 `bson:"_revIncludedCarePlanResourcesReferencingContext,omitempty"`
	RevIncludedProcedureResourcesReferencingEncounter               *[]Procedure                `bson:"_revIncludedProcedureResourcesReferencingEncounter,omitempty"`
	RevIncludedProcedureResourcesReferencingContext                 *[]Procedure                `bson:"_revIncludedProcedureResourcesReferencingContext,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                     `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingEncounter                    *[]List                     `bson:"_revIncludedListResourcesReferencingEncounter,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingContext         *[]MedicationRequest        `bson:"_revIncludedMedicationRequestResourcesReferencingContext,omitempty"`
	RevIncludedVisionPrescriptionResourcesReferencingEncounter      *[]VisionPrescription       `bson:"_revIncludedVisionPrescriptionResourcesReferencingEncounter,omitempty"`
	RevIncludedMediaResourcesReferencingContext                     *[]Media                    `bson:"_revIncludedMediaResourcesReferencingContext,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces         *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingEncounter        *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon          *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingContext          *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingContext,omitempty"`
	RevIncludedFlagResourcesReferencingEncounter                    *[]Flag                     `bson:"_revIncludedFlagResourcesReferencingEncounter,omitempty"`
	RevIncludedObservationResourcesReferencingContext               *[]Observation              `bson:"_revIncludedObservationResourcesReferencingContext,omitempty"`
	RevIncludedObservationResourcesReferencingEncounter             *[]Observation              `bson:"_revIncludedObservationResourcesReferencingEncounter,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingContext  *[]MedicationAdministration `bson:"_revIncludedMedicationAdministrationResourcesReferencingContext,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                 *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom               *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor               *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                 *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingContext       *[]MedicationStatement      `bson:"_revIncludedMedicationStatementResourcesReferencingContext,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingEncounter    *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon      *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingContext      *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingContext,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                     *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingContext        *[]MedicationDispense       `bson:"_revIncludedMedicationDispenseResourcesReferencingContext,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingEncounter        *[]DiagnosticReport         `bson:"_revIncludedDiagnosticReportResourcesReferencingEncounter,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingContext          *[]DiagnosticReport         `bson:"_revIncludedDiagnosticReportResourcesReferencingContext,omitempty"`
	RevIncludedNutritionOrderResourcesReferencingEncounter          *[]NutritionOrder           `bson:"_revIncludedNutritionOrderResourcesReferencingEncounter,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                 *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail          *[]Condition                `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedConditionResourcesReferencingEncounter               *[]Condition                `bson:"_revIncludedConditionResourcesReferencingEncounter,omitempty"`
	RevIncludedConditionResourcesReferencingContext                 *[]Condition                `bson:"_revIncludedConditionResourcesReferencingContext,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject               *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEncounter             *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEncounter,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                 *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated          *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject     *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingContext     *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingContext,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest           *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingContext        *[]ClinicalImpression       `bson:"_revIncludedClinicalImpressionResourcesReferencingContext,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor          *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom        *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor        *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof         *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1     *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2     *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedClaimResourcesReferencingEncounter                   *[]Claim                    `bson:"_revIncludedClaimResourcesReferencingEncounter,omitempty"`
}

func (e *EncounterPlusRelatedResources) GetIncludedEpisodeOfCareResourcesReferencedByEpisodeofcare() (episodeOfCares []EpisodeOfCare, err error) {
	if e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare == nil {
		err = errors.New("Included episodeOfCares not requested")
	} else {
		episodeOfCares = *e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedReferralRequestResourcesReferencedByIncomingreferral() (referralRequests []ReferralRequest, err error) {
	if e.IncludedReferralRequestResourcesReferencedByIncomingreferral == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *e.IncludedReferralRequestResourcesReferencedByIncomingreferral
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPractitioner() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByPractitioner == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByPractitioner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByPractitioner))
	} else if len(*e.IncludedPractitionerResourcesReferencedByPractitioner) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByPractitioner)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if e.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*e.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*e.IncludedGroupResourcesReferencedBySubject))
	} else if len(*e.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*e.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedBySubject))
	} else if len(*e.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedConditionResourceReferencedByDiagnosis() (condition *Condition, err error) {
	if e.IncludedConditionResourcesReferencedByDiagnosis == nil {
		err = errors.New("Included conditions not requested")
	} else if len(*e.IncludedConditionResourcesReferencedByDiagnosis) > 1 {
		err = fmt.Errorf("Expected 0 or 1 condition, but found %d", len(*e.IncludedConditionResourcesReferencedByDiagnosis))
	} else if len(*e.IncludedConditionResourcesReferencedByDiagnosis) == 1 {
		condition = &(*e.IncludedConditionResourcesReferencedByDiagnosis)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedProcedureResourceReferencedByDiagnosis() (procedure *Procedure, err error) {
	if e.IncludedProcedureResourcesReferencedByDiagnosis == nil {
		err = errors.New("Included procedures not requested")
	} else if len(*e.IncludedProcedureResourcesReferencedByDiagnosis) > 1 {
		err = fmt.Errorf("Expected 0 or 1 procedure, but found %d", len(*e.IncludedProcedureResourcesReferencedByDiagnosis))
	} else if len(*e.IncludedProcedureResourcesReferencedByDiagnosis) == 1 {
		procedure = &(*e.IncludedProcedureResourcesReferencedByDiagnosis)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedAppointmentResourceReferencedByAppointment() (appointment *Appointment, err error) {
	if e.IncludedAppointmentResourcesReferencedByAppointment == nil {
		err = errors.New("Included appointments not requested")
	} else if len(*e.IncludedAppointmentResourcesReferencedByAppointment) > 1 {
		err = fmt.Errorf("Expected 0 or 1 appointment, but found %d", len(*e.IncludedAppointmentResourcesReferencedByAppointment))
	} else if len(*e.IncludedAppointmentResourcesReferencedByAppointment) == 1 {
		appointment = &(*e.IncludedAppointmentResourcesReferencedByAppointment)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedEncounterResourceReferencedByPartof() (encounter *Encounter, err error) {
	if e.IncludedEncounterResourcesReferencedByPartof == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*e.IncludedEncounterResourcesReferencedByPartof) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*e.IncludedEncounterResourcesReferencedByPartof))
	} else if len(*e.IncludedEncounterResourcesReferencedByPartof) == 1 {
		encounter = &(*e.IncludedEncounterResourcesReferencedByPartof)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedPractitionerResourceReferencedByParticipant() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByParticipant == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByParticipant))
	} else if len(*e.IncludedPractitionerResourcesReferencedByParticipant) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByParticipant)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByParticipant() (relatedPerson *RelatedPerson, err error) {
	if e.IncludedRelatedPersonResourcesReferencedByParticipant == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*e.IncludedRelatedPersonResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*e.IncludedRelatedPersonResourcesReferencedByParticipant))
	} else if len(*e.IncludedRelatedPersonResourcesReferencedByParticipant) == 1 {
		relatedPerson = &(*e.IncludedRelatedPersonResourcesReferencedByParticipant)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedByPatient))
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if e.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*e.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*e.IncludedLocationResourcesReferencedByLocation))
	} else if len(*e.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*e.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedOrganizationResourceReferencedByServiceprovider() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByServiceprovider == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByServiceprovider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByServiceprovider))
	} else if len(*e.IncludedOrganizationResourcesReferencedByServiceprovider) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByServiceprovider)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingEncounter() (referralRequests []ReferralRequest, err error) {
	if e.RevIncludedReferralRequestResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *e.RevIncludedReferralRequestResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingContext() (referralRequests []ReferralRequest, err error) {
	if e.RevIncludedReferralRequestResourcesReferencingContext == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *e.RevIncludedReferralRequestResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if e.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *e.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if e.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *e.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingEncounter() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingEncounter() (riskAssessments []RiskAssessment, err error) {
	if e.RevIncludedRiskAssessmentResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *e.RevIncludedRiskAssessmentResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingContext() (careTeams []CareTeam, err error) {
	if e.RevIncludedCareTeamResourcesReferencingContext == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *e.RevIncludedCareTeamResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingEncounter() (careTeams []CareTeam, err error) {
	if e.RevIncludedCareTeamResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *e.RevIncludedCareTeamResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if e.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *e.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingContext() (imagingStudies []ImagingStudy, err error) {
	if e.RevIncludedImagingStudyResourcesReferencingContext == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *e.RevIncludedImagingStudyResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingContext() (chargeItems []ChargeItem, err error) {
	if e.RevIncludedChargeItemResourcesReferencingContext == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *e.RevIncludedChargeItemResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingPartof() (encounters []Encounter, err error) {
	if e.RevIncludedEncounterResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *e.RevIncludedEncounterResourcesReferencingPartof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if e.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *e.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if e.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *e.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if e.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *e.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if e.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *e.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingEncounter() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingContext() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingContext == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingEncounter() (requestGroups []RequestGroup, err error) {
	if e.RevIncludedRequestGroupResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *e.RevIncludedRequestGroupResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingContext() (requestGroups []RequestGroup, err error) {
	if e.RevIncludedRequestGroupResourcesReferencingContext == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *e.RevIncludedRequestGroupResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if e.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *e.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingEncounter() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedTaskResourcesReferencingContext() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingContext == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingEncounter() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingEncounter() (carePlans []CarePlan, err error) {
	if e.RevIncludedCarePlanResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *e.RevIncludedCarePlanResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingContext() (carePlans []CarePlan, err error) {
	if e.RevIncludedCarePlanResourcesReferencingContext == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *e.RevIncludedCarePlanResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingEncounter() (procedures []Procedure, err error) {
	if e.RevIncludedProcedureResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *e.RevIncludedProcedureResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingContext() (procedures []Procedure, err error) {
	if e.RevIncludedProcedureResourcesReferencingContext == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *e.RevIncludedProcedureResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingItem
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedListResourcesReferencingEncounter() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingContext() (medicationRequests []MedicationRequest, err error) {
	if e.RevIncludedMedicationRequestResourcesReferencingContext == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *e.RevIncludedMedicationRequestResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedVisionPrescriptionResourcesReferencingEncounter() (visionPrescriptions []VisionPrescription, err error) {
	if e.RevIncludedVisionPrescriptionResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded visionPrescriptions not requested")
	} else {
		visionPrescriptions = *e.RevIncludedVisionPrescriptionResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMediaResourcesReferencingContext() (media []Media, err error) {
	if e.RevIncludedMediaResourcesReferencingContext == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *e.RevIncludedMediaResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if e.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *e.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingEncounter() (procedureRequests []ProcedureRequest, err error) {
	if e.RevIncludedProcedureRequestResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *e.RevIncludedProcedureRequestResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if e.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *e.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingContext() (procedureRequests []ProcedureRequest, err error) {
	if e.RevIncludedProcedureRequestResourcesReferencingContext == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *e.RevIncludedProcedureRequestResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedFlagResourcesReferencingEncounter() (flags []Flag, err error) {
	if e.RevIncludedFlagResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *e.RevIncludedFlagResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedObservationResourcesReferencingContext() (observations []Observation, err error) {
	if e.RevIncludedObservationResourcesReferencingContext == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *e.RevIncludedObservationResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedObservationResourcesReferencingEncounter() (observations []Observation, err error) {
	if e.RevIncludedObservationResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *e.RevIncludedObservationResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingContext() (medicationAdministrations []MedicationAdministration, err error) {
	if e.RevIncludedMedicationAdministrationResourcesReferencingContext == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *e.RevIncludedMedicationAdministrationResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingContext() (medicationStatements []MedicationStatement, err error) {
	if e.RevIncludedMedicationStatementResourcesReferencingContext == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *e.RevIncludedMedicationStatementResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingEncounter() (communicationRequests []CommunicationRequest, err error) {
	if e.RevIncludedCommunicationRequestResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *e.RevIncludedCommunicationRequestResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *e.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingContext() (communicationRequests []CommunicationRequest, err error) {
	if e.RevIncludedCommunicationRequestResourcesReferencingContext == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *e.RevIncludedCommunicationRequestResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if e.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *e.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingContext() (medicationDispenses []MedicationDispense, err error) {
	if e.RevIncludedMedicationDispenseResourcesReferencingContext == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *e.RevIncludedMedicationDispenseResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingEncounter() (diagnosticReports []DiagnosticReport, err error) {
	if e.RevIncludedDiagnosticReportResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *e.RevIncludedDiagnosticReportResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingContext() (diagnosticReports []DiagnosticReport, err error) {
	if e.RevIncludedDiagnosticReportResourcesReferencingContext == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *e.RevIncludedDiagnosticReportResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedNutritionOrderResourcesReferencingEncounter() (nutritionOrders []NutritionOrder, err error) {
	if e.RevIncludedNutritionOrderResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded nutritionOrders not requested")
	} else {
		nutritionOrders = *e.RevIncludedNutritionOrderResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if e.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *e.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEncounter() (conditions []Condition, err error) {
	if e.RevIncludedConditionResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *e.RevIncludedConditionResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedConditionResourcesReferencingContext() (conditions []Condition, err error) {
	if e.RevIncludedConditionResourcesReferencingContext == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *e.RevIncludedConditionResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEncounter() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *e.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingContext() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingContext == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if e.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *e.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingContext() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingContext == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedClaimResourcesReferencingEncounter() (claims []Claim, err error) {
	if e.RevIncludedClaimResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *e.RevIncludedClaimResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare != nil {
		for idx := range *e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare {
			rsc := (*e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedReferralRequestResourcesReferencedByIncomingreferral != nil {
		for idx := range *e.IncludedReferralRequestResourcesReferencedByIncomingreferral {
			rsc := (*e.IncludedReferralRequestResourcesReferencedByIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByPractitioner {
			rsc := (*e.IncludedPractitionerResourcesReferencedByPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *e.IncludedGroupResourcesReferencedBySubject {
			rsc := (*e.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *e.IncludedPatientResourcesReferencedBySubject {
			rsc := (*e.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedConditionResourcesReferencedByDiagnosis != nil {
		for idx := range *e.IncludedConditionResourcesReferencedByDiagnosis {
			rsc := (*e.IncludedConditionResourcesReferencedByDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedProcedureResourcesReferencedByDiagnosis != nil {
		for idx := range *e.IncludedProcedureResourcesReferencedByDiagnosis {
			rsc := (*e.IncludedProcedureResourcesReferencedByDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedAppointmentResourcesReferencedByAppointment != nil {
		for idx := range *e.IncludedAppointmentResourcesReferencedByAppointment {
			rsc := (*e.IncludedAppointmentResourcesReferencedByAppointment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedEncounterResourcesReferencedByPartof != nil {
		for idx := range *e.IncludedEncounterResourcesReferencedByPartof {
			rsc := (*e.IncludedEncounterResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByParticipant {
			rsc := (*e.IncludedPractitionerResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for idx := range *e.IncludedRelatedPersonResourcesReferencedByParticipant {
			rsc := (*e.IncludedRelatedPersonResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatient {
			rsc := (*e.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *e.IncludedLocationResourcesReferencedByLocation {
			rsc := (*e.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByServiceprovider != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByServiceprovider {
			rsc := (*e.IncludedOrganizationResourcesReferencedByServiceprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.RevIncludedReferralRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedReferralRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedReferralRequestResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*e.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*e.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*e.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRiskAssessmentResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedRiskAssessmentResourcesReferencingEncounter {
			rsc := (*e.RevIncludedRiskAssessmentResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedCareTeamResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedCareTeamResourcesReferencingContext {
			rsc := (*e.RevIncludedCareTeamResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCareTeamResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCareTeamResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCareTeamResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedChargeItemResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedChargeItemResourcesReferencingContext {
			rsc := (*e.RevIncludedChargeItemResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEncounterResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedEncounterResourcesReferencingPartof {
			rsc := (*e.RevIncludedEncounterResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingEncounter {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingContext {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter {
			rsc := (*e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCarePlanResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCarePlanResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCarePlanResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCarePlanResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedCarePlanResourcesReferencingContext {
			rsc := (*e.RevIncludedCarePlanResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedProcedureResourcesReferencingEncounter {
			rsc := (*e.RevIncludedProcedureResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedProcedureResourcesReferencingContext {
			rsc := (*e.RevIncludedProcedureResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedListResourcesReferencingEncounter {
			rsc := (*e.RevIncludedListResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationRequestResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationRequestResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationRequestResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedVisionPrescriptionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedVisionPrescriptionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedVisionPrescriptionResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMediaResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMediaResourcesReferencingContext {
			rsc := (*e.RevIncludedMediaResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingContext {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedFlagResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedFlagResourcesReferencingEncounter {
			rsc := (*e.RevIncludedFlagResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedObservationResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedObservationResourcesReferencingContext {
			rsc := (*e.RevIncludedObservationResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedObservationResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedObservationResourcesReferencingEncounter {
			rsc := (*e.RevIncludedObservationResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationAdministrationResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationAdministrationResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationAdministrationResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*e.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationStatementResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationStatementResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationStatementResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if e.RevIncludedMedicationDispenseResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationDispenseResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationDispenseResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticReportResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDiagnosticReportResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDiagnosticReportResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticReportResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedDiagnosticReportResourcesReferencingContext {
			rsc := (*e.RevIncludedDiagnosticReportResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedNutritionOrderResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedNutritionOrderResourcesReferencingEncounter {
			rsc := (*e.RevIncludedNutritionOrderResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*e.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedConditionResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedCompositionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClaimResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedClaimResourcesReferencingEncounter {
			rsc := (*e.RevIncludedClaimResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *EncounterPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare != nil {
		for idx := range *e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare {
			rsc := (*e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedReferralRequestResourcesReferencedByIncomingreferral != nil {
		for idx := range *e.IncludedReferralRequestResourcesReferencedByIncomingreferral {
			rsc := (*e.IncludedReferralRequestResourcesReferencedByIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByPractitioner {
			rsc := (*e.IncludedPractitionerResourcesReferencedByPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *e.IncludedGroupResourcesReferencedBySubject {
			rsc := (*e.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *e.IncludedPatientResourcesReferencedBySubject {
			rsc := (*e.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedConditionResourcesReferencedByDiagnosis != nil {
		for idx := range *e.IncludedConditionResourcesReferencedByDiagnosis {
			rsc := (*e.IncludedConditionResourcesReferencedByDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedProcedureResourcesReferencedByDiagnosis != nil {
		for idx := range *e.IncludedProcedureResourcesReferencedByDiagnosis {
			rsc := (*e.IncludedProcedureResourcesReferencedByDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedAppointmentResourcesReferencedByAppointment != nil {
		for idx := range *e.IncludedAppointmentResourcesReferencedByAppointment {
			rsc := (*e.IncludedAppointmentResourcesReferencedByAppointment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedEncounterResourcesReferencedByPartof != nil {
		for idx := range *e.IncludedEncounterResourcesReferencedByPartof {
			rsc := (*e.IncludedEncounterResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByParticipant {
			rsc := (*e.IncludedPractitionerResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for idx := range *e.IncludedRelatedPersonResourcesReferencedByParticipant {
			rsc := (*e.IncludedRelatedPersonResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatient {
			rsc := (*e.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *e.IncludedLocationResourcesReferencedByLocation {
			rsc := (*e.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByServiceprovider != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByServiceprovider {
			rsc := (*e.IncludedOrganizationResourcesReferencedByServiceprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedReferralRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedReferralRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedReferralRequestResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*e.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*e.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*e.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRiskAssessmentResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedRiskAssessmentResourcesReferencingEncounter {
			rsc := (*e.RevIncludedRiskAssessmentResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedCareTeamResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedCareTeamResourcesReferencingContext {
			rsc := (*e.RevIncludedCareTeamResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCareTeamResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCareTeamResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCareTeamResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedChargeItemResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedChargeItemResourcesReferencingContext {
			rsc := (*e.RevIncludedChargeItemResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEncounterResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedEncounterResourcesReferencingPartof {
			rsc := (*e.RevIncludedEncounterResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingEncounter {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingContext {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter {
			rsc := (*e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCarePlanResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCarePlanResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCarePlanResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCarePlanResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedCarePlanResourcesReferencingContext {
			rsc := (*e.RevIncludedCarePlanResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedProcedureResourcesReferencingEncounter {
			rsc := (*e.RevIncludedProcedureResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedProcedureResourcesReferencingContext {
			rsc := (*e.RevIncludedProcedureResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedListResourcesReferencingEncounter {
			rsc := (*e.RevIncludedListResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationRequestResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationRequestResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationRequestResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedVisionPrescriptionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedVisionPrescriptionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedVisionPrescriptionResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMediaResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMediaResourcesReferencingContext {
			rsc := (*e.RevIncludedMediaResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingContext {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedFlagResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedFlagResourcesReferencingEncounter {
			rsc := (*e.RevIncludedFlagResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedObservationResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedObservationResourcesReferencingContext {
			rsc := (*e.RevIncludedObservationResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedObservationResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedObservationResourcesReferencingEncounter {
			rsc := (*e.RevIncludedObservationResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationAdministrationResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationAdministrationResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationAdministrationResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*e.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationStatementResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationStatementResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationStatementResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if e.RevIncludedMedicationDispenseResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationDispenseResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationDispenseResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticReportResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDiagnosticReportResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDiagnosticReportResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticReportResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedDiagnosticReportResourcesReferencingContext {
			rsc := (*e.RevIncludedDiagnosticReportResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedNutritionOrderResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedNutritionOrderResourcesReferencingEncounter {
			rsc := (*e.RevIncludedNutritionOrderResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*e.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedConditionResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedCompositionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEncounter)[idx]
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
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClaimResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedClaimResourcesReferencingEncounter {
			rsc := (*e.RevIncludedClaimResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
