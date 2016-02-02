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

type Encounter struct {
	DomainResource   `bson:",inline"`
	Identifier       []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status           string                             `bson:"status,omitempty" json:"status,omitempty"`
	StatusHistory    []EncounterStatusHistoryComponent  `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Class            string                             `bson:"class,omitempty" json:"class,omitempty"`
	Type             []CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	Priority         *CodeableConcept                   `bson:"priority,omitempty" json:"priority,omitempty"`
	Patient          *Reference                         `bson:"patient,omitempty" json:"patient,omitempty"`
	EpisodeOfCare    []Reference                        `bson:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty"`
	IncomingReferral []Reference                        `bson:"incomingReferral,omitempty" json:"incomingReferral,omitempty"`
	Participant      []EncounterParticipantComponent    `bson:"participant,omitempty" json:"participant,omitempty"`
	Appointment      *Reference                         `bson:"appointment,omitempty" json:"appointment,omitempty"`
	Period           *Period                            `bson:"period,omitempty" json:"period,omitempty"`
	Length           *Quantity                          `bson:"length,omitempty" json:"length,omitempty"`
	Reason           []CodeableConcept                  `bson:"reason,omitempty" json:"reason,omitempty"`
	Indication       []Reference                        `bson:"indication,omitempty" json:"indication,omitempty"`
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
	Status string  `bson:"status,omitempty" json:"status,omitempty"`
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type EncounterParticipantComponent struct {
	Type       []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Period     *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Individual *Reference        `bson:"individual,omitempty" json:"individual,omitempty"`
}

type EncounterHospitalizationComponent struct {
	PreAdmissionIdentifier *Identifier       `bson:"preAdmissionIdentifier,omitempty" json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference        `bson:"origin,omitempty" json:"origin,omitempty"`
	AdmitSource            *CodeableConcept  `bson:"admitSource,omitempty" json:"admitSource,omitempty"`
	AdmittingDiagnosis     []Reference       `bson:"admittingDiagnosis,omitempty" json:"admittingDiagnosis,omitempty"`
	ReAdmission            *CodeableConcept  `bson:"reAdmission,omitempty" json:"reAdmission,omitempty"`
	DietPreference         []CodeableConcept `bson:"dietPreference,omitempty" json:"dietPreference,omitempty"`
	SpecialCourtesy        []CodeableConcept `bson:"specialCourtesy,omitempty" json:"specialCourtesy,omitempty"`
	SpecialArrangement     []CodeableConcept `bson:"specialArrangement,omitempty" json:"specialArrangement,omitempty"`
	Destination            *Reference        `bson:"destination,omitempty" json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept  `bson:"dischargeDisposition,omitempty" json:"dischargeDisposition,omitempty"`
	DischargeDiagnosis     []Reference       `bson:"dischargeDiagnosis,omitempty" json:"dischargeDiagnosis,omitempty"`
}

type EncounterLocationComponent struct {
	Location *Reference `bson:"location,omitempty" json:"location,omitempty"`
	Status   string     `bson:"status,omitempty" json:"status,omitempty"`
	Period   *Period    `bson:"period,omitempty" json:"period,omitempty"`
}

type EncounterPlus struct {
	Encounter                     `bson:",inline"`
	EncounterPlusRelatedResources `bson:",inline"`
}

type EncounterPlusRelatedResources struct {
	IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare          *[]EpisodeOfCare            `bson:"_includedEpisodeOfCareResourcesReferencedByEpisodeofcare,omitempty"`
	IncludedReferralRequestResourcesReferencedByIncomingreferral     *[]ReferralRequest          `bson:"_includedReferralRequestResourcesReferencedByIncomingreferral,omitempty"`
	IncludedPractitionerResourcesReferencedByPractitioner            *[]Practitioner             `bson:"_includedPractitionerResourcesReferencedByPractitioner,omitempty"`
	IncludedAppointmentResourcesReferencedByAppointment              *[]Appointment              `bson:"_includedAppointmentResourcesReferencedByAppointment,omitempty"`
	IncludedEncounterResourcesReferencedByPartof                     *[]Encounter                `bson:"_includedEncounterResourcesReferencedByPartof,omitempty"`
	IncludedProcedureResourcesReferencedByProcedure                  *[]Procedure                `bson:"_includedProcedureResourcesReferencedByProcedure,omitempty"`
	IncludedPractitionerResourcesReferencedByParticipant             *[]Practitioner             `bson:"_includedPractitionerResourcesReferencedByParticipant,omitempty"`
	IncludedRelatedPersonResourcesReferencedByParticipant            *[]RelatedPerson            `bson:"_includedRelatedPersonResourcesReferencedByParticipant,omitempty"`
	IncludedConditionResourcesReferencedByCondition                  *[]Condition                `bson:"_includedConditionResourcesReferencedByCondition,omitempty"`
	IncludedPatientResourcesReferencedByPatient                      *[]Patient                  `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedLocationResourcesReferencedByLocation                    *[]Location                 `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	IncludedConditionResourcesReferencedByIndication                 *[]Condition                `bson:"_includedConditionResourcesReferencedByIndication,omitempty"`
	IncludedProcedureResourcesReferencedByIndication                 *[]Procedure                `bson:"_includedProcedureResourcesReferencedByIndication,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                  *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref        *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref        *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedProcedureResourcesReferencingEncounter                *[]Procedure                `bson:"_revIncludedProcedureResourcesReferencingEncounter,omitempty"`
	RevIncludedListResourcesReferencingItem                          *[]List                     `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingEncounter                     *[]List                     `bson:"_revIncludedListResourcesReferencingEncounter,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingEncounter        *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingEncounter,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref       *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                       *[]Order                    `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedVisionPrescriptionResourcesReferencingEncounter       *[]VisionPrescription       `bson:"_revIncludedVisionPrescriptionResourcesReferencingEncounter,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingEncounter         *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedFlagResourcesReferencingEncounter                     *[]Flag                     `bson:"_revIncludedFlagResourcesReferencingEncounter,omitempty"`
	RevIncludedObservationResourcesReferencingEncounter              *[]Observation              `bson:"_revIncludedObservationResourcesReferencingEncounter,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingEncounter *[]MedicationAdministration `bson:"_revIncludedMedicationAdministrationResourcesReferencingEncounter,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingEncounter     *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingEncounter           *[]RiskAssessment           `bson:"_revIncludedRiskAssessmentResourcesReferencingEncounter,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                      *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingEncounter         *[]DiagnosticReport         `bson:"_revIncludedDiagnosticReportResourcesReferencingEncounter,omitempty"`
	RevIncludedNutritionOrderResourcesReferencingEncounter           *[]NutritionOrder           `bson:"_revIncludedNutritionOrderResourcesReferencingEncounter,omitempty"`
	RevIncludedEncounterResourcesReferencingPartof                   *[]Encounter                `bson:"_revIncludedEncounterResourcesReferencingPartof,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference               *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedMedicationOrderResourcesReferencingEncounter          *[]MedicationOrder          `bson:"_revIncludedMedicationOrderResourcesReferencingEncounter,omitempty"`
	RevIncludedCommunicationResourcesReferencingEncounter            *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingEncounter,omitempty"`
	RevIncludedConditionResourcesReferencingEncounter                *[]Condition                `bson:"_revIncludedConditionResourcesReferencingEncounter,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEncounter              *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEncounter,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                  *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated           *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingEncounter          *[]DiagnosticOrder          `bson:"_revIncludedDiagnosticOrderResourcesReferencingEncounter,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment          *[]OrderResponse            `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject      *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingEncounter    *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingEncounter,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest            *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger         *[]ClinicalImpression       `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                 *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
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

func (e *EncounterPlusRelatedResources) GetIncludedProcedureResourcesReferencedByProcedure() (procedures []Procedure, err error) {
	if e.IncludedProcedureResourcesReferencedByProcedure == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *e.IncludedProcedureResourcesReferencedByProcedure
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

func (e *EncounterPlusRelatedResources) GetIncludedConditionResourcesReferencedByCondition() (conditions []Condition, err error) {
	if e.IncludedConditionResourcesReferencedByCondition == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *e.IncludedConditionResourcesReferencedByCondition
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

func (e *EncounterPlusRelatedResources) GetIncludedConditionResourcesReferencedByIndication() (conditions []Condition, err error) {
	if e.IncludedConditionResourcesReferencedByIndication == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *e.IncludedConditionResourcesReferencedByIndication
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedProcedureResourcesReferencedByIndication() (procedures []Procedure, err error) {
	if e.IncludedProcedureResourcesReferencedByIndication == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *e.IncludedProcedureResourcesReferencedByIndication
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

func (e *EncounterPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingEncounter() (procedures []Procedure, err error) {
	if e.RevIncludedProcedureResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *e.RevIncludedProcedureResourcesReferencingEncounter
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

func (e *EncounterPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if e.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *e.RevIncludedOrderResourcesReferencingDetail
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

func (e *EncounterPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingEncounter() (procedureRequests []ProcedureRequest, err error) {
	if e.RevIncludedProcedureRequestResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *e.RevIncludedProcedureRequestResourcesReferencingEncounter
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

func (e *EncounterPlusRelatedResources) GetRevIncludedObservationResourcesReferencingEncounter() (observations []Observation, err error) {
	if e.RevIncludedObservationResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *e.RevIncludedObservationResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingEncounter() (medicationAdministrations []MedicationAdministration, err error) {
	if e.RevIncludedMedicationAdministrationResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *e.RevIncludedMedicationAdministrationResourcesReferencingEncounter
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

func (e *EncounterPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingEncounter() (riskAssessments []RiskAssessment, err error) {
	if e.RevIncludedRiskAssessmentResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *e.RevIncludedRiskAssessmentResourcesReferencingEncounter
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

func (e *EncounterPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingEncounter() (diagnosticReports []DiagnosticReport, err error) {
	if e.RevIncludedDiagnosticReportResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *e.RevIncludedDiagnosticReportResourcesReferencingEncounter
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

func (e *EncounterPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingPartof() (encounters []Encounter, err error) {
	if e.RevIncludedEncounterResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *e.RevIncludedEncounterResourcesReferencingPartof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMedicationOrderResourcesReferencingEncounter() (medicationOrders []MedicationOrder, err error) {
	if e.RevIncludedMedicationOrderResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded medicationOrders not requested")
	} else {
		medicationOrders = *e.RevIncludedMedicationOrderResourcesReferencingEncounter
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

func (e *EncounterPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEncounter() (conditions []Condition, err error) {
	if e.RevIncludedConditionResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *e.RevIncludedConditionResourcesReferencingEncounter
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

func (e *EncounterPlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingEncounter() (diagnosticOrders []DiagnosticOrder, err error) {
	if e.RevIncludedDiagnosticOrderResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *e.RevIncludedDiagnosticOrderResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *e.RevIncludedOrderResponseResourcesReferencingFulfillment
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

func (e *EncounterPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingEncounter() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter
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

func (e *EncounterPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare != nil {
		for _, r := range *e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedReferralRequestResourcesReferencedByIncomingreferral != nil {
		for _, r := range *e.IncludedReferralRequestResourcesReferencedByIncomingreferral {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for _, r := range *e.IncludedPractitionerResourcesReferencedByPractitioner {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedAppointmentResourcesReferencedByAppointment != nil {
		for _, r := range *e.IncludedAppointmentResourcesReferencedByAppointment {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedEncounterResourcesReferencedByPartof != nil {
		for _, r := range *e.IncludedEncounterResourcesReferencedByPartof {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedProcedureResourcesReferencedByProcedure != nil {
		for _, r := range *e.IncludedProcedureResourcesReferencedByProcedure {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for _, r := range *e.IncludedPractitionerResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for _, r := range *e.IncludedRelatedPersonResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedConditionResourcesReferencedByCondition != nil {
		for _, r := range *e.IncludedConditionResourcesReferencedByCondition {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *e.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedLocationResourcesReferencedByLocation != nil {
		for _, r := range *e.IncludedLocationResourcesReferencedByLocation {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedConditionResourcesReferencedByIndication != nil {
		for _, r := range *e.IncludedConditionResourcesReferencedByIndication {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedProcedureResourcesReferencedByIndication != nil {
		for _, r := range *e.IncludedProcedureResourcesReferencedByIndication {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProcedureResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedProcedureResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *e.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedListResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedListResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedDocumentReferenceResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *e.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedVisionPrescriptionResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedVisionPrescriptionResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedProcedureRequestResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedFlagResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedFlagResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedObservationResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedObservationResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedMedicationAdministrationResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedMedicationAdministrationResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedCommunicationRequestResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedRiskAssessmentResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedRiskAssessmentResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDiagnosticReportResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedDiagnosticReportResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedNutritionOrderResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedNutritionOrderResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedEncounterResourcesReferencingPartof != nil {
		for _, r := range *e.RevIncludedEncounterResourcesReferencingPartof {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *e.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedMedicationOrderResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedMedicationOrderResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedCommunicationResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedConditionResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedConditionResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDiagnosticOrderResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedDiagnosticOrderResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *e.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *e.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (e *EncounterPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare != nil {
		for _, r := range *e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedReferralRequestResourcesReferencedByIncomingreferral != nil {
		for _, r := range *e.IncludedReferralRequestResourcesReferencedByIncomingreferral {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for _, r := range *e.IncludedPractitionerResourcesReferencedByPractitioner {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedAppointmentResourcesReferencedByAppointment != nil {
		for _, r := range *e.IncludedAppointmentResourcesReferencedByAppointment {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedEncounterResourcesReferencedByPartof != nil {
		for _, r := range *e.IncludedEncounterResourcesReferencedByPartof {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedProcedureResourcesReferencedByProcedure != nil {
		for _, r := range *e.IncludedProcedureResourcesReferencedByProcedure {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for _, r := range *e.IncludedPractitionerResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for _, r := range *e.IncludedRelatedPersonResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedConditionResourcesReferencedByCondition != nil {
		for _, r := range *e.IncludedConditionResourcesReferencedByCondition {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *e.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedLocationResourcesReferencedByLocation != nil {
		for _, r := range *e.IncludedLocationResourcesReferencedByLocation {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedConditionResourcesReferencedByIndication != nil {
		for _, r := range *e.IncludedConditionResourcesReferencedByIndication {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedProcedureResourcesReferencedByIndication != nil {
		for _, r := range *e.IncludedProcedureResourcesReferencedByIndication {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProcedureResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedProcedureResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *e.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedListResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedListResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedDocumentReferenceResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *e.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedVisionPrescriptionResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedVisionPrescriptionResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedProcedureRequestResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedFlagResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedFlagResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedObservationResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedObservationResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedMedicationAdministrationResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedMedicationAdministrationResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedCommunicationRequestResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedRiskAssessmentResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedRiskAssessmentResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDiagnosticReportResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedDiagnosticReportResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedNutritionOrderResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedNutritionOrderResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedEncounterResourcesReferencingPartof != nil {
		for _, r := range *e.RevIncludedEncounterResourcesReferencingPartof {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *e.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedMedicationOrderResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedMedicationOrderResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedCommunicationResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedConditionResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedConditionResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDiagnosticOrderResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedDiagnosticOrderResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *e.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter != nil {
		for _, r := range *e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *e.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
