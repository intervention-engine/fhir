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

type RelatedPerson struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient        *Reference       `bson:"patient,omitempty" json:"patient,omitempty"`
	Relationship   *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name           *HumanName       `bson:"name,omitempty" json:"name,omitempty"`
	Telecom        []ContactPoint   `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender         string           `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate      *FHIRDateTime    `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address        []Address        `bson:"address,omitempty" json:"address,omitempty"`
	Photo          []Attachment     `bson:"photo,omitempty" json:"photo,omitempty"`
	Period         *Period          `bson:"period,omitempty" json:"period,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *RelatedPerson) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "RelatedPerson"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to RelatedPerson), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *RelatedPerson) GetBSON() (interface{}, error) {
	x.ResourceType = "RelatedPerson"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "relatedPerson" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type relatedPerson RelatedPerson

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *RelatedPerson) UnmarshalJSON(data []byte) (err error) {
	x2 := relatedPerson{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = RelatedPerson(x2)
		return x.checkResourceType()
	}
	return
}

func (x *RelatedPerson) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "RelatedPerson"
	} else if x.ResourceType != "RelatedPerson" {
		return errors.New(fmt.Sprintf("Expected resourceType to be RelatedPerson, instead received %s", x.ResourceType))
	}
	return nil
}

type RelatedPersonPlus struct {
	RelatedPerson                     `bson:",inline"`
	RelatedPersonPlusRelatedResources `bson:",inline"`
}

type RelatedPersonPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                         *[]Patient                  `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                     *[]Appointment              `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                      *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                     *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref           *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor               *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref           *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRecipient            *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRecipient,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingReporter           *[]AllergyIntolerance       `bson:"_revIncludedAllergyIntoleranceResourcesReferencingReporter,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                    *[]CarePlan                 `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedCarePlanResourcesReferencingParticipant                  *[]CarePlan                 `bson:"_revIncludedCarePlanResourcesReferencingParticipant,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                   *[]Procedure                `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedListResourcesReferencingItem                             *[]List                     `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor              *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref          *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                          *[]Order                    `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingPerformer            *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingOrderer              *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingOrderer,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor             *[]AppointmentResponse      `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedObservationResourcesReferencingPerformer                 *[]Observation              `bson:"_revIncludedObservationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPractitioner *[]MedicationAdministration `bson:"_revIncludedMedicationAdministrationResourcesReferencingPractitioner,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSource            *[]MedicationStatement      `bson:"_revIncludedMedicationStatementResourcesReferencingSource,omitempty"`
	RevIncludedPersonResourcesReferencingLink                           *[]Person                   `bson:"_revIncludedPersonResourcesReferencingLink,omitempty"`
	RevIncludedPersonResourcesReferencingRelatedperson                  *[]Person                   `bson:"_revIncludedPersonResourcesReferencingRelatedperson,omitempty"`
	RevIncludedContractResourcesReferencingActor                        *[]Contract                 `bson:"_revIncludedContractResourcesReferencingActor,omitempty"`
	RevIncludedContractResourcesReferencingSigner                       *[]Contract                 `bson:"_revIncludedContractResourcesReferencingSigner,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRequester        *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender           *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient        *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                         *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingAuthor                          *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingAuthor,omitempty"`
	RevIncludedImagingObjectSelectionResourcesReferencingAuthor         *[]ImagingObjectSelection   `bson:"_revIncludedImagingObjectSelectionResourcesReferencingAuthor,omitempty"`
	RevIncludedEncounterResourcesReferencingParticipant                 *[]Encounter                `bson:"_revIncludedEncounterResourcesReferencingParticipant,omitempty"`
	RevIncludedAuditEventResourcesReferencingParticipant                *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingParticipant,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference                  *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender                  *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient               *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                   *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor                    *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                     *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated              *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment             *[]OrderResponse            `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject         *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor          *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSource          *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSource,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest               *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                        *[]Schedule                 `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger            *[]ClinicalImpression       `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                    *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (r *RelatedPersonPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if r.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResourcesReferencedByPatient))
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*r.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if r.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *r.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRecipient() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingRecipient
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingReporter() (allergyIntolerances []AllergyIntolerance, err error) {
	if r.RevIncludedAllergyIntoleranceResourcesReferencingReporter == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *r.RevIncludedAllergyIntoleranceResourcesReferencingReporter
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if r.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *r.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingParticipant() (carePlans []CarePlan, err error) {
	if r.RevIncludedCarePlanResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *r.RevIncludedCarePlanResourcesReferencingParticipant
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if r.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *r.RevIncludedProcedureResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if r.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *r.RevIncludedListResourcesReferencingItem
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if r.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if r.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *r.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingPerformer() (procedureRequests []ProcedureRequest, err error) {
	if r.RevIncludedProcedureRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *r.RevIncludedProcedureRequestResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingOrderer() (procedureRequests []ProcedureRequest, err error) {
	if r.RevIncludedProcedureRequestResourcesReferencingOrderer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *r.RevIncludedProcedureRequestResourcesReferencingOrderer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if r.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *r.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPerformer() (observations []Observation, err error) {
	if r.RevIncludedObservationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *r.RevIncludedObservationResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPractitioner() (medicationAdministrations []MedicationAdministration, err error) {
	if r.RevIncludedMedicationAdministrationResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *r.RevIncludedMedicationAdministrationResourcesReferencingPractitioner
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSource() (medicationStatements []MedicationStatement, err error) {
	if r.RevIncludedMedicationStatementResourcesReferencingSource == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *r.RevIncludedMedicationStatementResourcesReferencingSource
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPersonResourcesReferencingLink() (people []Person, err error) {
	if r.RevIncludedPersonResourcesReferencingLink == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *r.RevIncludedPersonResourcesReferencingLink
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPersonResourcesReferencingRelatedperson() (people []Person, err error) {
	if r.RevIncludedPersonResourcesReferencingRelatedperson == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *r.RevIncludedPersonResourcesReferencingRelatedperson
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedContractResourcesReferencingActor() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingActor == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingActor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedContractResourcesReferencingSigner() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingSigner == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingSigner
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRequester() (communicationRequests []CommunicationRequest, err error) {
	if r.RevIncludedCommunicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *r.RevIncludedCommunicationRequestResourcesReferencingRequester
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if r.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *r.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if r.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *r.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if r.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *r.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedBasicResourcesReferencingAuthor() (basics []Basic, err error) {
	if r.RevIncludedBasicResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *r.RevIncludedBasicResourcesReferencingAuthor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedImagingObjectSelectionResourcesReferencingAuthor() (imagingObjectSelections []ImagingObjectSelection, err error) {
	if r.RevIncludedImagingObjectSelectionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded imagingObjectSelections not requested")
	} else {
		imagingObjectSelections = *r.RevIncludedImagingObjectSelectionResourcesReferencingAuthor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingParticipant() (encounters []Encounter, err error) {
	if r.RevIncludedEncounterResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *r.RevIncludedEncounterResourcesReferencingParticipant
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingParticipant() (auditEvents []AuditEvent, err error) {
	if r.RevIncludedAuditEventResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *r.RevIncludedAuditEventResourcesReferencingParticipant
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if r.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *r.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAuthor() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingAuthor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *r.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if r.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *r.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSource() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSource == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSource
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if r.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *r.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if r.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *r.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if r.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *r.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if r.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *r.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByPatient {
			rsc := (*r.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAllergyIntoleranceResourcesReferencingReporter != nil {
		for idx := range *r.RevIncludedAllergyIntoleranceResourcesReferencingReporter {
			rsc := (*r.RevIncludedAllergyIntoleranceResourcesReferencingReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*r.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCarePlanResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedCarePlanResourcesReferencingParticipant {
			rsc := (*r.RevIncludedCarePlanResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*r.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *r.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*r.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*r.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for idx := range *r.RevIncludedProcedureRequestResourcesReferencingOrderer {
			rsc := (*r.RevIncludedProcedureRequestResourcesReferencingOrderer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*r.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*r.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationAdministrationResourcesReferencingPractitioner != nil {
		for idx := range *r.RevIncludedMedicationAdministrationResourcesReferencingPractitioner {
			rsc := (*r.RevIncludedMedicationAdministrationResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*r.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *r.RevIncludedPersonResourcesReferencingLink {
			rsc := (*r.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPersonResourcesReferencingRelatedperson != nil {
		for idx := range *r.RevIncludedPersonResourcesReferencingRelatedperson {
			rsc := (*r.RevIncludedPersonResourcesReferencingRelatedperson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingActor {
			rsc := (*r.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSigner {
			rsc := (*r.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*r.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			rsc := (*r.RevIncludedImagingObjectSelectionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedEncounterResourcesReferencingParticipant {
			rsc := (*r.RevIncludedEncounterResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingParticipant {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*r.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *r.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*r.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*r.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*r.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *RelatedPersonPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByPatient {
			rsc := (*r.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAllergyIntoleranceResourcesReferencingReporter != nil {
		for idx := range *r.RevIncludedAllergyIntoleranceResourcesReferencingReporter {
			rsc := (*r.RevIncludedAllergyIntoleranceResourcesReferencingReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*r.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCarePlanResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedCarePlanResourcesReferencingParticipant {
			rsc := (*r.RevIncludedCarePlanResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*r.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *r.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*r.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*r.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for idx := range *r.RevIncludedProcedureRequestResourcesReferencingOrderer {
			rsc := (*r.RevIncludedProcedureRequestResourcesReferencingOrderer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*r.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*r.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationAdministrationResourcesReferencingPractitioner != nil {
		for idx := range *r.RevIncludedMedicationAdministrationResourcesReferencingPractitioner {
			rsc := (*r.RevIncludedMedicationAdministrationResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*r.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *r.RevIncludedPersonResourcesReferencingLink {
			rsc := (*r.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPersonResourcesReferencingRelatedperson != nil {
		for idx := range *r.RevIncludedPersonResourcesReferencingRelatedperson {
			rsc := (*r.RevIncludedPersonResourcesReferencingRelatedperson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingActor {
			rsc := (*r.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSigner {
			rsc := (*r.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*r.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			rsc := (*r.RevIncludedImagingObjectSelectionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedEncounterResourcesReferencingParticipant {
			rsc := (*r.RevIncludedEncounterResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingParticipant {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*r.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *r.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*r.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*r.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*r.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
