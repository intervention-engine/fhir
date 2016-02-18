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

type Practitioner struct {
	DomainResource   `bson:",inline"`
	Identifier       []Identifier                            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active           *bool                                   `bson:"active,omitempty" json:"active,omitempty"`
	Name             *HumanName                              `bson:"name,omitempty" json:"name,omitempty"`
	Telecom          []ContactPoint                          `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address          []Address                               `bson:"address,omitempty" json:"address,omitempty"`
	Gender           string                                  `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate        *FHIRDateTime                           `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Photo            []Attachment                            `bson:"photo,omitempty" json:"photo,omitempty"`
	PractitionerRole []PractitionerPractitionerRoleComponent `bson:"practitionerRole,omitempty" json:"practitionerRole,omitempty"`
	Qualification    []PractitionerQualificationComponent    `bson:"qualification,omitempty" json:"qualification,omitempty"`
	Communication    []CodeableConcept                       `bson:"communication,omitempty" json:"communication,omitempty"`
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
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
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

type PractitionerPractitionerRoleComponent struct {
	BackboneElement      `bson:",inline"`
	ManagingOrganization *Reference        `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Role                 *CodeableConcept  `bson:"role,omitempty" json:"role,omitempty"`
	Specialty            []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Period               *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Location             []Reference       `bson:"location,omitempty" json:"location,omitempty"`
	HealthcareService    []Reference       `bson:"healthcareService,omitempty" json:"healthcareService,omitempty"`
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
	IncludedOrganizationResourcesReferencedByOrganization               *[]Organization             `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedLocationResourcesReferencedByLocation                       *[]Location                 `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                     *[]Appointment              `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingPractitioner              *[]Appointment              `bson:"_revIncludedAppointmentResourcesReferencingPractitioner,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRequester             *[]ReferralRequest          `bson:"_revIncludedReferralRequestResourcesReferencingRequester,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRecipient             *[]ReferralRequest          `bson:"_revIncludedReferralRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                       *[]Account                  `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                      *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                     *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref           *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingSubject              *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor               *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref           *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRecipient            *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRecipient,omitempty"`
	RevIncludedSpecimenResourcesReferencingCollector                    *[]Specimen                 `bson:"_revIncludedSpecimenResourcesReferencingCollector,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingRecorder           *[]AllergyIntolerance       `bson:"_revIncludedAllergyIntoleranceResourcesReferencingRecorder,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingReporter           *[]AllergyIntolerance       `bson:"_revIncludedAllergyIntoleranceResourcesReferencingReporter,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                    *[]CarePlan                 `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedCarePlanResourcesReferencingParticipant                  *[]CarePlan                 `bson:"_revIncludedCarePlanResourcesReferencingParticipant,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingTeammember              *[]EpisodeOfCare            `bson:"_revIncludedEpisodeOfCareResourcesReferencingTeammember,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingCaremanager             *[]EpisodeOfCare            `bson:"_revIncludedEpisodeOfCareResourcesReferencingCaremanager,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                   *[]Procedure                `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedListResourcesReferencingItem                             *[]List                     `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSource                           *[]List                     `bson:"_revIncludedListResourcesReferencingSource,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingSubject             *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthenticator       *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthenticator,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor              *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref          *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingSource                          *[]Order                    `bson:"_revIncludedOrderResourcesReferencingSource,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                          *[]Order                    `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedOrderResourcesReferencingTarget                          *[]Order                    `bson:"_revIncludedOrderResourcesReferencingTarget,omitempty"`
	RevIncludedImmunizationResourcesReferencingRequester                *[]Immunization             `bson:"_revIncludedImmunizationResourcesReferencingRequester,omitempty"`
	RevIncludedImmunizationResourcesReferencingPerformer                *[]Immunization             `bson:"_revIncludedImmunizationResourcesReferencingPerformer,omitempty"`
	RevIncludedVisionPrescriptionResourcesReferencingPrescriber         *[]VisionPrescription       `bson:"_revIncludedVisionPrescriptionResourcesReferencingPrescriber,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                         *[]Media                    `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedMediaResourcesReferencingOperator                        *[]Media                    `bson:"_revIncludedMediaResourcesReferencingOperator,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingPerformer            *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingOrderer              *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingOrderer,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                          *[]Flag                     `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                           *[]Flag                     `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingSource                  *[]SupplyRequest            `bson:"_revIncludedSupplyRequestResourcesReferencingSource,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor             *[]AppointmentResponse      `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingPractitioner      *[]AppointmentResponse      `bson:"_revIncludedAppointmentResponseResourcesReferencingPractitioner,omitempty"`
	RevIncludedObservationResourcesReferencingPerformer                 *[]Observation              `bson:"_revIncludedObservationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPractitioner *[]MedicationAdministration `bson:"_revIncludedMedicationAdministrationResourcesReferencingPractitioner,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSource            *[]MedicationStatement      `bson:"_revIncludedMedicationStatementResourcesReferencingSource,omitempty"`
	RevIncludedPersonResourcesReferencingPractitioner                   *[]Person                   `bson:"_revIncludedPersonResourcesReferencingPractitioner,omitempty"`
	RevIncludedPersonResourcesReferencingLink                           *[]Person                   `bson:"_revIncludedPersonResourcesReferencingLink,omitempty"`
	RevIncludedContractResourcesReferencingActor                        *[]Contract                 `bson:"_revIncludedContractResourcesReferencingActor,omitempty"`
	RevIncludedContractResourcesReferencingSigner                       *[]Contract                 `bson:"_revIncludedContractResourcesReferencingSigner,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRequester        *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender           *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient        *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingPerformer              *[]RiskAssessment           `bson:"_revIncludedRiskAssessmentResourcesReferencingPerformer,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                         *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingAuthor                          *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingAuthor,omitempty"`
	RevIncludedGroupResourcesReferencingMember                          *[]Group                    `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedProcessRequestResourcesReferencingProvider               *[]ProcessRequest           `bson:"_revIncludedProcessRequestResourcesReferencingProvider,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingReceiver           *[]MedicationDispense       `bson:"_revIncludedMedicationDispenseResourcesReferencingReceiver,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingResponsibleparty   *[]MedicationDispense       `bson:"_revIncludedMedicationDispenseResourcesReferencingResponsibleparty,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingDispenser          *[]MedicationDispense       `bson:"_revIncludedMedicationDispenseResourcesReferencingDispenser,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingPerformer            *[]DiagnosticReport         `bson:"_revIncludedDiagnosticReportResourcesReferencingPerformer,omitempty"`
	RevIncludedImagingObjectSelectionResourcesReferencingAuthor         *[]ImagingObjectSelection   `bson:"_revIncludedImagingObjectSelectionResourcesReferencingAuthor,omitempty"`
	RevIncludedNutritionOrderResourcesReferencingProvider               *[]NutritionOrder           `bson:"_revIncludedNutritionOrderResourcesReferencingProvider,omitempty"`
	RevIncludedEncounterResourcesReferencingPractitioner                *[]Encounter                `bson:"_revIncludedEncounterResourcesReferencingPractitioner,omitempty"`
	RevIncludedEncounterResourcesReferencingParticipant                 *[]Encounter                `bson:"_revIncludedEncounterResourcesReferencingParticipant,omitempty"`
	RevIncludedAuditEventResourcesReferencingParticipant                *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingParticipant,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference                  *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedMedicationOrderResourcesReferencingPrescriber            *[]MedicationOrder          `bson:"_revIncludedMedicationOrderResourcesReferencingPrescriber,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender                  *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient               *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedConditionResourcesReferencingAsserter                    *[]Condition                `bson:"_revIncludedConditionResourcesReferencingAsserter,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                   *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor                    *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingAttester                  *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingAttester,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                     *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingAuthor                  *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingAuthor,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated              *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingActorPath1            *[]DiagnosticOrder          `bson:"_revIncludedDiagnosticOrderResourcesReferencingActorPath1,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingActorPath2            *[]DiagnosticOrder          `bson:"_revIncludedDiagnosticOrderResourcesReferencingActorPath2,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingOrderer               *[]DiagnosticOrder          `bson:"_revIncludedDiagnosticOrderResourcesReferencingOrderer,omitempty"`
	RevIncludedPatientResourcesReferencingCareprovider                  *[]Patient                  `bson:"_revIncludedPatientResourcesReferencingCareprovider,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment             *[]OrderResponse            `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedOrderResponseResourcesReferencingWho                     *[]OrderResponse            `bson:"_revIncludedOrderResponseResourcesReferencingWho,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject         *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor          *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSource          *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSource,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest               *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestprovider       *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequestprovider,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                        *[]Schedule                 `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedSupplyDeliveryResourcesReferencingReceiver               *[]SupplyDelivery           `bson:"_revIncludedSupplyDeliveryResourcesReferencingReceiver,omitempty"`
	RevIncludedSupplyDeliveryResourcesReferencingSupplier               *[]SupplyDelivery           `bson:"_revIncludedSupplyDeliveryResourcesReferencingSupplier,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAssessor           *[]ClinicalImpression       `bson:"_revIncludedClinicalImpressionResourcesReferencingAssessor,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger            *[]ClinicalImpression       `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                    *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingReceiver                *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingReceiver,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingAuthor                  *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingAuthor,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingResponsible             *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingResponsible,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingEnterer                 *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingEnterer,omitempty"`
	RevIncludedClaimResourcesReferencingProvider                        *[]Claim                    `bson:"_revIncludedClaimResourcesReferencingProvider,omitempty"`
}

func (p *PractitionerPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if p.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*p.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*p.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetIncludedLocationResourcesReferencedByLocation() (locations []Location, err error) {
	if p.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else {
		locations = *p.IncludedLocationResourcesReferencedByLocation
	}
	return
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingReporter() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingReporter == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingReporter
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingParticipant() (carePlans []CarePlan, err error) {
	if p.RevIncludedCarePlanResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *p.RevIncludedCarePlanResourcesReferencingParticipant
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingTeammember() (episodeOfCares []EpisodeOfCare, err error) {
	if p.RevIncludedEpisodeOfCareResourcesReferencingTeammember == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *p.RevIncludedEpisodeOfCareResourcesReferencingTeammember
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedOrderResourcesReferencingSource() (orders []Order, err error) {
	if p.RevIncludedOrderResourcesReferencingSource == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *p.RevIncludedOrderResourcesReferencingSource
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if p.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *p.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedOrderResourcesReferencingTarget() (orders []Order, err error) {
	if p.RevIncludedOrderResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *p.RevIncludedOrderResourcesReferencingTarget
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingRequester() (immunizations []Immunization, err error) {
	if p.RevIncludedImmunizationResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *p.RevIncludedImmunizationResourcesReferencingRequester
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingPerformer() (immunizations []Immunization, err error) {
	if p.RevIncludedImmunizationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *p.RevIncludedImmunizationResourcesReferencingPerformer
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingPerformer() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingOrderer() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingOrderer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingOrderer
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingSource() (supplyRequests []SupplyRequest, err error) {
	if p.RevIncludedSupplyRequestResourcesReferencingSource == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *p.RevIncludedSupplyRequestResourcesReferencingSource
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPerformer() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPractitioner() (medicationAdministrations []MedicationAdministration, err error) {
	if p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.RevIncludedMedicationAdministrationResourcesReferencingPractitioner
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedContractResourcesReferencingActor() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingActor == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingActor
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRequester() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingRequester
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingPerformer() (riskAssessments []RiskAssessment, err error) {
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *p.RevIncludedRiskAssessmentResourcesReferencingPerformer
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if p.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *p.RevIncludedGroupResourcesReferencingMember
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingDispenser() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingDispenser == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingDispenser
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedImagingObjectSelectionResourcesReferencingAuthor() (imagingObjectSelections []ImagingObjectSelection, err error) {
	if p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded imagingObjectSelections not requested")
	} else {
		imagingObjectSelections = *p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingParticipant() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingParticipant
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationOrderResourcesReferencingPrescriber() (medicationOrders []MedicationOrder, err error) {
	if p.RevIncludedMedicationOrderResourcesReferencingPrescriber == nil {
		err = errors.New("RevIncluded medicationOrders not requested")
	} else {
		medicationOrders = *p.RevIncludedMedicationOrderResourcesReferencingPrescriber
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingActorPath1() (diagnosticOrders []DiagnosticOrder, err error) {
	if p.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *p.RevIncludedDiagnosticOrderResourcesReferencingActorPath1
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingActorPath2() (diagnosticOrders []DiagnosticOrder, err error) {
	if p.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *p.RevIncludedDiagnosticOrderResourcesReferencingActorPath2
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingOrderer() (diagnosticOrders []DiagnosticOrder, err error) {
	if p.RevIncludedDiagnosticOrderResourcesReferencingOrderer == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *p.RevIncludedDiagnosticOrderResourcesReferencingOrderer
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedPatientResourcesReferencingCareprovider() (patients []Patient, err error) {
	if p.RevIncludedPatientResourcesReferencingCareprovider == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *p.RevIncludedPatientResourcesReferencingCareprovider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if p.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *p.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingWho() (orderResponses []OrderResponse, err error) {
	if p.RevIncludedOrderResponseResourcesReferencingWho == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *p.RevIncludedOrderResponseResourcesReferencingWho
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingData
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

func (p *PractitionerPlusRelatedResources) GetRevIncludedClaimResourcesReferencingProvider() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingProvider
	}
	return
}

func (p *PractitionerPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *p.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedLocationResourcesReferencedByLocation != nil {
		for _, r := range *p.IncludedLocationResourcesReferencedByLocation {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (p *PractitionerPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedAppointmentResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingPractitioner != nil {
		for _, r := range *p.RevIncludedAppointmentResourcesReferencingPractitioner {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for _, r := range *p.RevIncludedReferralRequestResourcesReferencingRequester {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedReferralRequestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedAccountResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for _, r := range *p.RevIncludedProvenanceResourcesReferencingAgent {
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
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingCollector != nil {
		for _, r := range *p.RevIncludedSpecimenResourcesReferencingCollector {
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
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingParticipant != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingTeammember != nil {
		for _, r := range *p.RevIncludedEpisodeOfCareResourcesReferencingTeammember {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager != nil {
		for _, r := range *p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingItem {
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
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
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
	if p.RevIncludedOrderResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingTarget != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingRequester != nil {
		for _, r := range *p.RevIncludedImmunizationResourcesReferencingRequester {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedImmunizationResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber != nil {
		for _, r := range *p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedMediaResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMediaResourcesReferencingOperator != nil {
		for _, r := range *p.RevIncludedMediaResourcesReferencingOperator {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingOrderer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedFlagResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFlagResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedFlagResourcesReferencingAuthor {
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
	if p.RevIncludedAppointmentResponseResourcesReferencingPractitioner != nil {
		for _, r := range *p.RevIncludedAppointmentResponseResourcesReferencingPractitioner {
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
	if p.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedMedicationStatementResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPersonResourcesReferencingPractitioner != nil {
		for _, r := range *p.RevIncludedPersonResourcesReferencingPractitioner {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for _, r := range *p.RevIncludedPersonResourcesReferencingLink {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedContractResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedContractResourcesReferencingActor {
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
	if p.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingSender {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedBasicResourcesReferencingSubject {
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
	if p.RevIncludedProcessRequestResourcesReferencingProvider != nil {
		for _, r := range *p.RevIncludedProcessRequestResourcesReferencingProvider {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for _, r := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty != nil {
		for _, r := range *p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingDispenser != nil {
		for _, r := range *p.RevIncludedMedicationDispenseResourcesReferencingDispenser {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingProvider != nil {
		for _, r := range *p.RevIncludedNutritionOrderResourcesReferencingProvider {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPractitioner != nil {
		for _, r := range *p.RevIncludedEncounterResourcesReferencingPractitioner {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for _, r := range *p.RevIncludedEncounterResourcesReferencingParticipant {
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
	if p.RevIncludedMedicationOrderResourcesReferencingPrescriber != nil {
		for _, r := range *p.RevIncludedMedicationOrderResourcesReferencingPrescriber {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for _, r := range *p.RevIncludedCommunicationResourcesReferencingSender {
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
	if p.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedDetectedIssueResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 != nil {
		for _, r := range *p.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 != nil {
		for _, r := range *p.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingOrderer != nil {
		for _, r := range *p.RevIncludedDiagnosticOrderResourcesReferencingOrderer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPatientResourcesReferencingCareprovider != nil {
		for _, r := range *p.RevIncludedPatientResourcesReferencingCareprovider {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *p.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResponseResourcesReferencingWho != nil {
		for _, r := range *p.RevIncludedOrderResponseResourcesReferencingWho {
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
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequestprovider != nil {
		for _, r := range *p.RevIncludedProcessResponseResourcesReferencingRequestprovider {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedScheduleResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingReceiver != nil {
		for _, r := range *p.RevIncludedSupplyDeliveryResourcesReferencingReceiver {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingSupplier != nil {
		for _, r := range *p.RevIncludedSupplyDeliveryResourcesReferencingSupplier {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingAssessor != nil {
		for _, r := range *p.RevIncludedClinicalImpressionResourcesReferencingAssessor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *p.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingReceiver {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingResponsible {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingEnterer != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingEnterer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClaimResourcesReferencingProvider != nil {
		for _, r := range *p.RevIncludedClaimResourcesReferencingProvider {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (p *PractitionerPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *p.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedLocationResourcesReferencedByLocation != nil {
		for _, r := range *p.IncludedLocationResourcesReferencedByLocation {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedAppointmentResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingPractitioner != nil {
		for _, r := range *p.RevIncludedAppointmentResourcesReferencingPractitioner {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for _, r := range *p.RevIncludedReferralRequestResourcesReferencingRequester {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedReferralRequestResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedReferralRequestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedAccountResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for _, r := range *p.RevIncludedProvenanceResourcesReferencingAgent {
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
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingCollector != nil {
		for _, r := range *p.RevIncludedSpecimenResourcesReferencingCollector {
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
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingParticipant != nil {
		for _, r := range *p.RevIncludedCarePlanResourcesReferencingParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingTeammember != nil {
		for _, r := range *p.RevIncludedEpisodeOfCareResourcesReferencingTeammember {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager != nil {
		for _, r := range *p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingItem {
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
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
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
	if p.RevIncludedOrderResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingTarget != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingRequester != nil {
		for _, r := range *p.RevIncludedImmunizationResourcesReferencingRequester {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedImmunizationResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber != nil {
		for _, r := range *p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedMediaResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMediaResourcesReferencingOperator != nil {
		for _, r := range *p.RevIncludedMediaResourcesReferencingOperator {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for _, r := range *p.RevIncludedProcedureRequestResourcesReferencingOrderer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedFlagResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedFlagResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedFlagResourcesReferencingAuthor {
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
	if p.RevIncludedAppointmentResponseResourcesReferencingPractitioner != nil {
		for _, r := range *p.RevIncludedAppointmentResponseResourcesReferencingPractitioner {
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
	if p.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedMedicationStatementResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPersonResourcesReferencingPractitioner != nil {
		for _, r := range *p.RevIncludedPersonResourcesReferencingPractitioner {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for _, r := range *p.RevIncludedPersonResourcesReferencingLink {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedContractResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedContractResourcesReferencingActor {
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
	if p.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingSender {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for _, r := range *p.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedBasicResourcesReferencingSubject {
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
	if p.RevIncludedProcessRequestResourcesReferencingProvider != nil {
		for _, r := range *p.RevIncludedProcessRequestResourcesReferencingProvider {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for _, r := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty != nil {
		for _, r := range *p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingDispenser != nil {
		for _, r := range *p.RevIncludedMedicationDispenseResourcesReferencingDispenser {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for _, r := range *p.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingProvider != nil {
		for _, r := range *p.RevIncludedNutritionOrderResourcesReferencingProvider {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPractitioner != nil {
		for _, r := range *p.RevIncludedEncounterResourcesReferencingPractitioner {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for _, r := range *p.RevIncludedEncounterResourcesReferencingParticipant {
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
	if p.RevIncludedMedicationOrderResourcesReferencingPrescriber != nil {
		for _, r := range *p.RevIncludedMedicationOrderResourcesReferencingPrescriber {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for _, r := range *p.RevIncludedCommunicationResourcesReferencingSender {
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
	if p.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedDetectedIssueResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 != nil {
		for _, r := range *p.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 != nil {
		for _, r := range *p.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDiagnosticOrderResourcesReferencingOrderer != nil {
		for _, r := range *p.RevIncludedDiagnosticOrderResourcesReferencingOrderer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPatientResourcesReferencingCareprovider != nil {
		for _, r := range *p.RevIncludedPatientResourcesReferencingCareprovider {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *p.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResponseResourcesReferencingWho != nil {
		for _, r := range *p.RevIncludedOrderResponseResourcesReferencingWho {
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
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequestprovider != nil {
		for _, r := range *p.RevIncludedProcessResponseResourcesReferencingRequestprovider {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for _, r := range *p.RevIncludedScheduleResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingReceiver != nil {
		for _, r := range *p.RevIncludedSupplyDeliveryResourcesReferencingReceiver {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingSupplier != nil {
		for _, r := range *p.RevIncludedSupplyDeliveryResourcesReferencingSupplier {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingAssessor != nil {
		for _, r := range *p.RevIncludedClinicalImpressionResourcesReferencingAssessor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *p.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingReceiver {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingAuthor != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingResponsible {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingEnterer != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingEnterer {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClaimResourcesReferencingProvider != nil {
		for _, r := range *p.RevIncludedClaimResourcesReferencingProvider {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
