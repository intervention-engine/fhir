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

type ClinicalImpression struct {
	DomainResource         `bson:",inline"`
	Patient                *Reference                                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Assessor               *Reference                                  `bson:"assessor,omitempty" json:"assessor,omitempty"`
	Status                 string                                      `bson:"status,omitempty" json:"status,omitempty"`
	Date                   *FHIRDateTime                               `bson:"date,omitempty" json:"date,omitempty"`
	Description            string                                      `bson:"description,omitempty" json:"description,omitempty"`
	Previous               *Reference                                  `bson:"previous,omitempty" json:"previous,omitempty"`
	Problem                []Reference                                 `bson:"problem,omitempty" json:"problem,omitempty"`
	TriggerCodeableConcept *CodeableConcept                            `bson:"triggerCodeableConcept,omitempty" json:"triggerCodeableConcept,omitempty"`
	TriggerReference       *Reference                                  `bson:"triggerReference,omitempty" json:"triggerReference,omitempty"`
	Investigations         []ClinicalImpressionInvestigationsComponent `bson:"investigations,omitempty" json:"investigations,omitempty"`
	Protocol               string                                      `bson:"protocol,omitempty" json:"protocol,omitempty"`
	Summary                string                                      `bson:"summary,omitempty" json:"summary,omitempty"`
	Finding                []ClinicalImpressionFindingComponent        `bson:"finding,omitempty" json:"finding,omitempty"`
	Resolved               []CodeableConcept                           `bson:"resolved,omitempty" json:"resolved,omitempty"`
	RuledOut               []ClinicalImpressionRuledOutComponent       `bson:"ruledOut,omitempty" json:"ruledOut,omitempty"`
	Prognosis              string                                      `bson:"prognosis,omitempty" json:"prognosis,omitempty"`
	Plan                   []Reference                                 `bson:"plan,omitempty" json:"plan,omitempty"`
	Action                 []Reference                                 `bson:"action,omitempty" json:"action,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ClinicalImpression) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ClinicalImpression"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ClinicalImpression), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ClinicalImpression) GetBSON() (interface{}, error) {
	x.ResourceType = "ClinicalImpression"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "clinicalImpression" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type clinicalImpression ClinicalImpression

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ClinicalImpression) UnmarshalJSON(data []byte) (err error) {
	x2 := clinicalImpression{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ClinicalImpression(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ClinicalImpression) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ClinicalImpression"
	} else if x.ResourceType != "ClinicalImpression" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ClinicalImpression, instead received %s", x.ResourceType))
	}
	return nil
}

type ClinicalImpressionInvestigationsComponent struct {
	BackboneElement `bson:",inline"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Item            []Reference      `bson:"item,omitempty" json:"item,omitempty"`
}

type ClinicalImpressionFindingComponent struct {
	BackboneElement `bson:",inline"`
	Item            *CodeableConcept `bson:"item,omitempty" json:"item,omitempty"`
	Cause           string           `bson:"cause,omitempty" json:"cause,omitempty"`
}

type ClinicalImpressionRuledOutComponent struct {
	BackboneElement `bson:",inline"`
	Item            *CodeableConcept `bson:"item,omitempty" json:"item,omitempty"`
	Reason          string           `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ClinicalImpressionPlus struct {
	ClinicalImpression                     `bson:",inline"`
	ClinicalImpressionPlusRelatedResources `bson:",inline"`
}

type ClinicalImpressionPlusRelatedResources struct {
	IncludedClinicalImpressionResourcesReferencedByPrevious         *[]ClinicalImpression    `bson:"_includedClinicalImpressionResourcesReferencedByPrevious,omitempty"`
	IncludedPractitionerResourcesReferencedByAssessor               *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAssessor,omitempty"`
	IncludedConditionResourcesReferencedByProblem                   *[]Condition             `bson:"_includedConditionResourcesReferencedByProblem,omitempty"`
	IncludedAllergyIntoleranceResourcesReferencedByProblem          *[]AllergyIntolerance    `bson:"_includedAllergyIntoleranceResourcesReferencedByProblem,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedFamilyMemberHistoryResourcesReferencedByInvestigation   *[]FamilyMemberHistory   `bson:"_includedFamilyMemberHistoryResourcesReferencedByInvestigation,omitempty"`
	IncludedObservationResourcesReferencedByInvestigation           *[]Observation           `bson:"_includedObservationResourcesReferencedByInvestigation,omitempty"`
	IncludedDiagnosticReportResourcesReferencedByInvestigation      *[]DiagnosticReport      `bson:"_includedDiagnosticReportResourcesReferencedByInvestigation,omitempty"`
	IncludedQuestionnaireResponseResourcesReferencedByInvestigation *[]QuestionnaireResponse `bson:"_includedQuestionnaireResponseResourcesReferencedByInvestigation,omitempty"`
	IncludedAppointmentResourcesReferencedByAction                  *[]Appointment           `bson:"_includedAppointmentResourcesReferencedByAction,omitempty"`
	IncludedReferralRequestResourcesReferencedByAction              *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByAction,omitempty"`
	IncludedNutritionOrderResourcesReferencedByAction               *[]NutritionOrder        `bson:"_includedNutritionOrderResourcesReferencedByAction,omitempty"`
	IncludedProcedureRequestResourcesReferencedByAction             *[]ProcedureRequest      `bson:"_includedProcedureRequestResourcesReferencedByAction,omitempty"`
	IncludedProcedureResourcesReferencedByAction                    *[]Procedure             `bson:"_includedProcedureResourcesReferencedByAction,omitempty"`
	IncludedDiagnosticOrderResourcesReferencedByAction              *[]DiagnosticOrder       `bson:"_includedDiagnosticOrderResourcesReferencedByAction,omitempty"`
	IncludedMedicationOrderResourcesReferencedByAction              *[]MedicationOrder       `bson:"_includedMedicationOrderResourcesReferencedByAction,omitempty"`
	IncludedSupplyRequestResourcesReferencedByAction                *[]SupplyRequest         `bson:"_includedSupplyRequestResourcesReferencedByAction,omitempty"`
	IncludedAppointmentResourcesReferencedByPlan                    *[]Appointment           `bson:"_includedAppointmentResourcesReferencedByPlan,omitempty"`
	IncludedOrderResourcesReferencedByPlan                          *[]Order                 `bson:"_includedOrderResourcesReferencedByPlan,omitempty"`
	IncludedReferralRequestResourcesReferencedByPlan                *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByPlan,omitempty"`
	IncludedProcessRequestResourcesReferencedByPlan                 *[]ProcessRequest        `bson:"_includedProcessRequestResourcesReferencedByPlan,omitempty"`
	IncludedVisionPrescriptionResourcesReferencedByPlan             *[]VisionPrescription    `bson:"_includedVisionPrescriptionResourcesReferencedByPlan,omitempty"`
	IncludedDiagnosticOrderResourcesReferencedByPlan                *[]DiagnosticOrder       `bson:"_includedDiagnosticOrderResourcesReferencedByPlan,omitempty"`
	IncludedProcedureRequestResourcesReferencedByPlan               *[]ProcedureRequest      `bson:"_includedProcedureRequestResourcesReferencedByPlan,omitempty"`
	IncludedDeviceUseRequestResourcesReferencedByPlan               *[]DeviceUseRequest      `bson:"_includedDeviceUseRequestResourcesReferencedByPlan,omitempty"`
	IncludedSupplyRequestResourcesReferencedByPlan                  *[]SupplyRequest         `bson:"_includedSupplyRequestResourcesReferencedByPlan,omitempty"`
	IncludedCarePlanResourcesReferencedByPlan                       *[]CarePlan              `bson:"_includedCarePlanResourcesReferencedByPlan,omitempty"`
	IncludedNutritionOrderResourcesReferencedByPlan                 *[]NutritionOrder        `bson:"_includedNutritionOrderResourcesReferencedByPlan,omitempty"`
	IncludedMedicationOrderResourcesReferencedByPlan                *[]MedicationOrder       `bson:"_includedMedicationOrderResourcesReferencedByPlan,omitempty"`
	IncludedCommunicationRequestResourcesReferencedByPlan           *[]CommunicationRequest  `bson:"_includedCommunicationRequestResourcesReferencedByPlan,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                 *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref      *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                      *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                     *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference              *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject               *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                 *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated          *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment         *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject     *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest           *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPrevious       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPrevious,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger        *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedClinicalImpressionResourceReferencedByPrevious() (clinicalImpression *ClinicalImpression, err error) {
	if c.IncludedClinicalImpressionResourcesReferencedByPrevious == nil {
		err = errors.New("Included clinicalimpressions not requested")
	} else if len(*c.IncludedClinicalImpressionResourcesReferencedByPrevious) > 1 {
		err = fmt.Errorf("Expected 0 or 1 clinicalImpression, but found %d", len(*c.IncludedClinicalImpressionResourcesReferencedByPrevious))
	} else if len(*c.IncludedClinicalImpressionResourcesReferencedByPrevious) == 1 {
		clinicalImpression = &(*c.IncludedClinicalImpressionResourcesReferencedByPrevious)[0]
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedPractitionerResourceReferencedByAssessor() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByAssessor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByAssessor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByAssessor))
	} else if len(*c.IncludedPractitionerResourcesReferencedByAssessor) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByAssessor)[0]
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedConditionResourcesReferencedByProblem() (conditions []Condition, err error) {
	if c.IncludedConditionResourcesReferencedByProblem == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *c.IncludedConditionResourcesReferencedByProblem
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedAllergyIntoleranceResourcesReferencedByProblem() (allergyIntolerances []AllergyIntolerance, err error) {
	if c.IncludedAllergyIntoleranceResourcesReferencedByProblem == nil {
		err = errors.New("Included allergyIntolerances not requested")
	} else {
		allergyIntolerances = *c.IncludedAllergyIntoleranceResourcesReferencedByProblem
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPatient))
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedFamilyMemberHistoryResourcesReferencedByInvestigation() (familyMemberHistories []FamilyMemberHistory, err error) {
	if c.IncludedFamilyMemberHistoryResourcesReferencedByInvestigation == nil {
		err = errors.New("Included familyMemberHistories not requested")
	} else {
		familyMemberHistories = *c.IncludedFamilyMemberHistoryResourcesReferencedByInvestigation
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedObservationResourcesReferencedByInvestigation() (observations []Observation, err error) {
	if c.IncludedObservationResourcesReferencedByInvestigation == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *c.IncludedObservationResourcesReferencedByInvestigation
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedDiagnosticReportResourcesReferencedByInvestigation() (diagnosticReports []DiagnosticReport, err error) {
	if c.IncludedDiagnosticReportResourcesReferencedByInvestigation == nil {
		err = errors.New("Included diagnosticReports not requested")
	} else {
		diagnosticReports = *c.IncludedDiagnosticReportResourcesReferencedByInvestigation
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedQuestionnaireResponseResourcesReferencedByInvestigation() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.IncludedQuestionnaireResponseResourcesReferencedByInvestigation == nil {
		err = errors.New("Included questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.IncludedQuestionnaireResponseResourcesReferencedByInvestigation
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedAppointmentResourcesReferencedByAction() (appointments []Appointment, err error) {
	if c.IncludedAppointmentResourcesReferencedByAction == nil {
		err = errors.New("Included appointments not requested")
	} else {
		appointments = *c.IncludedAppointmentResourcesReferencedByAction
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedReferralRequestResourcesReferencedByAction() (referralRequests []ReferralRequest, err error) {
	if c.IncludedReferralRequestResourcesReferencedByAction == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *c.IncludedReferralRequestResourcesReferencedByAction
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedNutritionOrderResourcesReferencedByAction() (nutritionOrders []NutritionOrder, err error) {
	if c.IncludedNutritionOrderResourcesReferencedByAction == nil {
		err = errors.New("Included nutritionOrders not requested")
	} else {
		nutritionOrders = *c.IncludedNutritionOrderResourcesReferencedByAction
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedProcedureRequestResourcesReferencedByAction() (procedureRequests []ProcedureRequest, err error) {
	if c.IncludedProcedureRequestResourcesReferencedByAction == nil {
		err = errors.New("Included procedureRequests not requested")
	} else {
		procedureRequests = *c.IncludedProcedureRequestResourcesReferencedByAction
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedProcedureResourcesReferencedByAction() (procedures []Procedure, err error) {
	if c.IncludedProcedureResourcesReferencedByAction == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *c.IncludedProcedureResourcesReferencedByAction
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedDiagnosticOrderResourcesReferencedByAction() (diagnosticOrders []DiagnosticOrder, err error) {
	if c.IncludedDiagnosticOrderResourcesReferencedByAction == nil {
		err = errors.New("Included diagnosticOrders not requested")
	} else {
		diagnosticOrders = *c.IncludedDiagnosticOrderResourcesReferencedByAction
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedMedicationOrderResourcesReferencedByAction() (medicationOrders []MedicationOrder, err error) {
	if c.IncludedMedicationOrderResourcesReferencedByAction == nil {
		err = errors.New("Included medicationOrders not requested")
	} else {
		medicationOrders = *c.IncludedMedicationOrderResourcesReferencedByAction
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedSupplyRequestResourcesReferencedByAction() (supplyRequests []SupplyRequest, err error) {
	if c.IncludedSupplyRequestResourcesReferencedByAction == nil {
		err = errors.New("Included supplyRequests not requested")
	} else {
		supplyRequests = *c.IncludedSupplyRequestResourcesReferencedByAction
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedAppointmentResourcesReferencedByPlan() (appointments []Appointment, err error) {
	if c.IncludedAppointmentResourcesReferencedByPlan == nil {
		err = errors.New("Included appointments not requested")
	} else {
		appointments = *c.IncludedAppointmentResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedOrderResourcesReferencedByPlan() (orders []Order, err error) {
	if c.IncludedOrderResourcesReferencedByPlan == nil {
		err = errors.New("Included orders not requested")
	} else {
		orders = *c.IncludedOrderResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedReferralRequestResourcesReferencedByPlan() (referralRequests []ReferralRequest, err error) {
	if c.IncludedReferralRequestResourcesReferencedByPlan == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *c.IncludedReferralRequestResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedProcessRequestResourcesReferencedByPlan() (processRequests []ProcessRequest, err error) {
	if c.IncludedProcessRequestResourcesReferencedByPlan == nil {
		err = errors.New("Included processRequests not requested")
	} else {
		processRequests = *c.IncludedProcessRequestResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedVisionPrescriptionResourcesReferencedByPlan() (visionPrescriptions []VisionPrescription, err error) {
	if c.IncludedVisionPrescriptionResourcesReferencedByPlan == nil {
		err = errors.New("Included visionPrescriptions not requested")
	} else {
		visionPrescriptions = *c.IncludedVisionPrescriptionResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedDiagnosticOrderResourcesReferencedByPlan() (diagnosticOrders []DiagnosticOrder, err error) {
	if c.IncludedDiagnosticOrderResourcesReferencedByPlan == nil {
		err = errors.New("Included diagnosticOrders not requested")
	} else {
		diagnosticOrders = *c.IncludedDiagnosticOrderResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedProcedureRequestResourcesReferencedByPlan() (procedureRequests []ProcedureRequest, err error) {
	if c.IncludedProcedureRequestResourcesReferencedByPlan == nil {
		err = errors.New("Included procedureRequests not requested")
	} else {
		procedureRequests = *c.IncludedProcedureRequestResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedDeviceUseRequestResourcesReferencedByPlan() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.IncludedDeviceUseRequestResourcesReferencedByPlan == nil {
		err = errors.New("Included deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.IncludedDeviceUseRequestResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedSupplyRequestResourcesReferencedByPlan() (supplyRequests []SupplyRequest, err error) {
	if c.IncludedSupplyRequestResourcesReferencedByPlan == nil {
		err = errors.New("Included supplyRequests not requested")
	} else {
		supplyRequests = *c.IncludedSupplyRequestResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByPlan() (carePlans []CarePlan, err error) {
	if c.IncludedCarePlanResourcesReferencedByPlan == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *c.IncludedCarePlanResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedNutritionOrderResourcesReferencedByPlan() (nutritionOrders []NutritionOrder, err error) {
	if c.IncludedNutritionOrderResourcesReferencedByPlan == nil {
		err = errors.New("Included nutritionOrders not requested")
	} else {
		nutritionOrders = *c.IncludedNutritionOrderResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedMedicationOrderResourcesReferencedByPlan() (medicationOrders []MedicationOrder, err error) {
	if c.IncludedMedicationOrderResourcesReferencedByPlan == nil {
		err = errors.New("Included medicationOrders not requested")
	} else {
		medicationOrders = *c.IncludedMedicationOrderResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedCommunicationRequestResourcesReferencedByPlan() (communicationRequests []CommunicationRequest, err error) {
	if c.IncludedCommunicationRequestResourcesReferencedByPlan == nil {
		err = errors.New("Included communicationRequests not requested")
	} else {
		communicationRequests = *c.IncludedCommunicationRequestResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPrevious() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingPrevious == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingPrevious
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedClinicalImpressionResourcesReferencedByPrevious != nil {
		for idx := range *c.IncludedClinicalImpressionResourcesReferencedByPrevious {
			rsc := (*c.IncludedClinicalImpressionResourcesReferencedByPrevious)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByAssessor != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByAssessor {
			rsc := (*c.IncludedPractitionerResourcesReferencedByAssessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByProblem != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByProblem {
			rsc := (*c.IncludedConditionResourcesReferencedByProblem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedAllergyIntoleranceResourcesReferencedByProblem != nil {
		for idx := range *c.IncludedAllergyIntoleranceResourcesReferencedByProblem {
			rsc := (*c.IncludedAllergyIntoleranceResourcesReferencedByProblem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedFamilyMemberHistoryResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedFamilyMemberHistoryResourcesReferencedByInvestigation {
			rsc := (*c.IncludedFamilyMemberHistoryResourcesReferencedByInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedObservationResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedObservationResourcesReferencedByInvestigation {
			rsc := (*c.IncludedObservationResourcesReferencedByInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticReportResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedDiagnosticReportResourcesReferencedByInvestigation {
			rsc := (*c.IncludedDiagnosticReportResourcesReferencedByInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedQuestionnaireResponseResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedQuestionnaireResponseResourcesReferencedByInvestigation {
			rsc := (*c.IncludedQuestionnaireResponseResourcesReferencedByInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedAppointmentResourcesReferencedByAction != nil {
		for idx := range *c.IncludedAppointmentResourcesReferencedByAction {
			rsc := (*c.IncludedAppointmentResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedReferralRequestResourcesReferencedByAction != nil {
		for idx := range *c.IncludedReferralRequestResourcesReferencedByAction {
			rsc := (*c.IncludedReferralRequestResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByAction != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByAction {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcedureRequestResourcesReferencedByAction != nil {
		for idx := range *c.IncludedProcedureRequestResourcesReferencedByAction {
			rsc := (*c.IncludedProcedureRequestResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcedureResourcesReferencedByAction != nil {
		for idx := range *c.IncludedProcedureResourcesReferencedByAction {
			rsc := (*c.IncludedProcedureResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticOrderResourcesReferencedByAction != nil {
		for idx := range *c.IncludedDiagnosticOrderResourcesReferencedByAction {
			rsc := (*c.IncludedDiagnosticOrderResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMedicationOrderResourcesReferencedByAction != nil {
		for idx := range *c.IncludedMedicationOrderResourcesReferencedByAction {
			rsc := (*c.IncludedMedicationOrderResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSupplyRequestResourcesReferencedByAction != nil {
		for idx := range *c.IncludedSupplyRequestResourcesReferencedByAction {
			rsc := (*c.IncludedSupplyRequestResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedAppointmentResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedAppointmentResourcesReferencedByPlan {
			rsc := (*c.IncludedAppointmentResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrderResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedOrderResourcesReferencedByPlan {
			rsc := (*c.IncludedOrderResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedReferralRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedReferralRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedReferralRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcessRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedProcessRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedProcessRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedVisionPrescriptionResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedVisionPrescriptionResourcesReferencedByPlan {
			rsc := (*c.IncludedVisionPrescriptionResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticOrderResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedDiagnosticOrderResourcesReferencedByPlan {
			rsc := (*c.IncludedDiagnosticOrderResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcedureRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedProcedureRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedProcedureRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceUseRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedDeviceUseRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedDeviceUseRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSupplyRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedSupplyRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedSupplyRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByPlan {
			rsc := (*c.IncludedCarePlanResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByPlan {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMedicationOrderResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedMedicationOrderResourcesReferencedByPlan {
			rsc := (*c.IncludedMedicationOrderResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *c.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*c.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*c.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingPrevious != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingPrevious {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingPrevious)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedClinicalImpressionResourcesReferencedByPrevious != nil {
		for idx := range *c.IncludedClinicalImpressionResourcesReferencedByPrevious {
			rsc := (*c.IncludedClinicalImpressionResourcesReferencedByPrevious)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByAssessor != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByAssessor {
			rsc := (*c.IncludedPractitionerResourcesReferencedByAssessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByProblem != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByProblem {
			rsc := (*c.IncludedConditionResourcesReferencedByProblem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedAllergyIntoleranceResourcesReferencedByProblem != nil {
		for idx := range *c.IncludedAllergyIntoleranceResourcesReferencedByProblem {
			rsc := (*c.IncludedAllergyIntoleranceResourcesReferencedByProblem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedFamilyMemberHistoryResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedFamilyMemberHistoryResourcesReferencedByInvestigation {
			rsc := (*c.IncludedFamilyMemberHistoryResourcesReferencedByInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedObservationResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedObservationResourcesReferencedByInvestigation {
			rsc := (*c.IncludedObservationResourcesReferencedByInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticReportResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedDiagnosticReportResourcesReferencedByInvestigation {
			rsc := (*c.IncludedDiagnosticReportResourcesReferencedByInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedQuestionnaireResponseResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedQuestionnaireResponseResourcesReferencedByInvestigation {
			rsc := (*c.IncludedQuestionnaireResponseResourcesReferencedByInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedAppointmentResourcesReferencedByAction != nil {
		for idx := range *c.IncludedAppointmentResourcesReferencedByAction {
			rsc := (*c.IncludedAppointmentResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedReferralRequestResourcesReferencedByAction != nil {
		for idx := range *c.IncludedReferralRequestResourcesReferencedByAction {
			rsc := (*c.IncludedReferralRequestResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByAction != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByAction {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcedureRequestResourcesReferencedByAction != nil {
		for idx := range *c.IncludedProcedureRequestResourcesReferencedByAction {
			rsc := (*c.IncludedProcedureRequestResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcedureResourcesReferencedByAction != nil {
		for idx := range *c.IncludedProcedureResourcesReferencedByAction {
			rsc := (*c.IncludedProcedureResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticOrderResourcesReferencedByAction != nil {
		for idx := range *c.IncludedDiagnosticOrderResourcesReferencedByAction {
			rsc := (*c.IncludedDiagnosticOrderResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMedicationOrderResourcesReferencedByAction != nil {
		for idx := range *c.IncludedMedicationOrderResourcesReferencedByAction {
			rsc := (*c.IncludedMedicationOrderResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSupplyRequestResourcesReferencedByAction != nil {
		for idx := range *c.IncludedSupplyRequestResourcesReferencedByAction {
			rsc := (*c.IncludedSupplyRequestResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedAppointmentResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedAppointmentResourcesReferencedByPlan {
			rsc := (*c.IncludedAppointmentResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrderResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedOrderResourcesReferencedByPlan {
			rsc := (*c.IncludedOrderResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedReferralRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedReferralRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedReferralRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcessRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedProcessRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedProcessRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedVisionPrescriptionResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedVisionPrescriptionResourcesReferencedByPlan {
			rsc := (*c.IncludedVisionPrescriptionResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticOrderResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedDiagnosticOrderResourcesReferencedByPlan {
			rsc := (*c.IncludedDiagnosticOrderResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcedureRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedProcedureRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedProcedureRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceUseRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedDeviceUseRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedDeviceUseRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSupplyRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedSupplyRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedSupplyRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByPlan {
			rsc := (*c.IncludedCarePlanResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByPlan {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMedicationOrderResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedMedicationOrderResourcesReferencedByPlan {
			rsc := (*c.IncludedMedicationOrderResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *c.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*c.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*c.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingPrevious != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingPrevious {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingPrevious)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
