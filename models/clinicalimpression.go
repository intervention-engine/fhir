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
	x := struct {
		ResourceType string `json:"resourceType"`
		ClinicalImpression
	}{
		ResourceType:       "ClinicalImpression",
		ClinicalImpression: *resource,
	}
	return json.Marshal(x)
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
	}
	return
}

type ClinicalImpressionInvestigationsComponent struct {
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Item []Reference      `bson:"item,omitempty" json:"item,omitempty"`
}

type ClinicalImpressionFindingComponent struct {
	Item  *CodeableConcept `bson:"item,omitempty" json:"item,omitempty"`
	Cause string           `bson:"cause,omitempty" json:"cause,omitempty"`
}

type ClinicalImpressionRuledOutComponent struct {
	Item   *CodeableConcept `bson:"item,omitempty" json:"item,omitempty"`
	Reason string           `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ClinicalImpressionPlus struct {
	ClinicalImpression             `bson:",inline"`
	ClinicalImpressionPlusIncludes `bson:",inline"`
}

type ClinicalImpressionPlusIncludes struct {
	IncludedPreviousResources                           *[]ClinicalImpression    `bson:"_includedPreviousResources,omitempty"`
	IncludedAssessorResources                           *[]Practitioner          `bson:"_includedAssessorResources,omitempty"`
	IncludedProblemConditionResources                   *[]Condition             `bson:"_includedProblemConditionResources,omitempty"`
	IncludedProblemAllergyIntoleranceResources          *[]AllergyIntolerance    `bson:"_includedProblemAllergyIntoleranceResources,omitempty"`
	IncludedPatientResources                            *[]Patient               `bson:"_includedPatientResources,omitempty"`
	IncludedInvestigationFamilyMemberHistoryResources   *[]FamilyMemberHistory   `bson:"_includedInvestigationFamilyMemberHistoryResources,omitempty"`
	IncludedInvestigationObservationResources           *[]Observation           `bson:"_includedInvestigationObservationResources,omitempty"`
	IncludedInvestigationDiagnosticReportResources      *[]DiagnosticReport      `bson:"_includedInvestigationDiagnosticReportResources,omitempty"`
	IncludedInvestigationQuestionnaireResponseResources *[]QuestionnaireResponse `bson:"_includedInvestigationQuestionnaireResponseResources,omitempty"`
	IncludedActionAppointmentResources                  *[]Appointment           `bson:"_includedActionAppointmentResources,omitempty"`
	IncludedActionReferralRequestResources              *[]ReferralRequest       `bson:"_includedActionReferralRequestResources,omitempty"`
	IncludedActionNutritionOrderResources               *[]NutritionOrder        `bson:"_includedActionNutritionOrderResources,omitempty"`
	IncludedActionProcedureRequestResources             *[]ProcedureRequest      `bson:"_includedActionProcedureRequestResources,omitempty"`
	IncludedActionProcedureResources                    *[]Procedure             `bson:"_includedActionProcedureResources,omitempty"`
	IncludedActionDiagnosticOrderResources              *[]DiagnosticOrder       `bson:"_includedActionDiagnosticOrderResources,omitempty"`
	IncludedActionMedicationOrderResources              *[]MedicationOrder       `bson:"_includedActionMedicationOrderResources,omitempty"`
	IncludedActionSupplyRequestResources                *[]SupplyRequest         `bson:"_includedActionSupplyRequestResources,omitempty"`
	IncludedPlanAppointmentResources                    *[]Appointment           `bson:"_includedPlanAppointmentResources,omitempty"`
	IncludedPlanOrderResources                          *[]Order                 `bson:"_includedPlanOrderResources,omitempty"`
	IncludedPlanReferralRequestResources                *[]ReferralRequest       `bson:"_includedPlanReferralRequestResources,omitempty"`
	IncludedPlanProcessRequestResources                 *[]ProcessRequest        `bson:"_includedPlanProcessRequestResources,omitempty"`
	IncludedPlanVisionPrescriptionResources             *[]VisionPrescription    `bson:"_includedPlanVisionPrescriptionResources,omitempty"`
	IncludedPlanDiagnosticOrderResources                *[]DiagnosticOrder       `bson:"_includedPlanDiagnosticOrderResources,omitempty"`
	IncludedPlanProcedureRequestResources               *[]ProcedureRequest      `bson:"_includedPlanProcedureRequestResources,omitempty"`
	IncludedPlanDeviceUseRequestResources               *[]DeviceUseRequest      `bson:"_includedPlanDeviceUseRequestResources,omitempty"`
	IncludedPlanSupplyRequestResources                  *[]SupplyRequest         `bson:"_includedPlanSupplyRequestResources,omitempty"`
	IncludedPlanCarePlanResources                       *[]CarePlan              `bson:"_includedPlanCarePlanResources,omitempty"`
	IncludedPlanNutritionOrderResources                 *[]NutritionOrder        `bson:"_includedPlanNutritionOrderResources,omitempty"`
	IncludedPlanMedicationOrderResources                *[]MedicationOrder       `bson:"_includedPlanMedicationOrderResources,omitempty"`
	IncludedPlanCommunicationRequestResources           *[]CommunicationRequest  `bson:"_includedPlanCommunicationRequestResources,omitempty"`
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPreviousResource() (clinicalImpression *ClinicalImpression, err error) {
	if c.IncludedPreviousResources == nil {
		err = errors.New("Included clinicalimpressions not requested")
	} else if len(*c.IncludedPreviousResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 clinicalImpression, but found %d", len(*c.IncludedPreviousResources))
	} else if len(*c.IncludedPreviousResources) == 1 {
		clinicalImpression = &(*c.IncludedPreviousResources)[0]
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedAssessorResource() (practitioner *Practitioner, err error) {
	if c.IncludedAssessorResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedAssessorResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedAssessorResources))
	} else if len(*c.IncludedAssessorResources) == 1 {
		practitioner = &(*c.IncludedAssessorResources)[0]
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedProblemConditionResources() (conditions []Condition, err error) {
	if c.IncludedProblemConditionResources == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *c.IncludedProblemConditionResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedProblemAllergyIntoleranceResources() (allergyIntolerances []AllergyIntolerance, err error) {
	if c.IncludedProblemAllergyIntoleranceResources == nil {
		err = errors.New("Included allergyIntolerances not requested")
	} else {
		allergyIntolerances = *c.IncludedProblemAllergyIntoleranceResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if c.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResources))
	} else if len(*c.IncludedPatientResources) == 1 {
		patient = &(*c.IncludedPatientResources)[0]
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedInvestigationFamilyMemberHistoryResources() (familyMemberHistories []FamilyMemberHistory, err error) {
	if c.IncludedInvestigationFamilyMemberHistoryResources == nil {
		err = errors.New("Included familyMemberHistories not requested")
	} else {
		familyMemberHistories = *c.IncludedInvestigationFamilyMemberHistoryResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedInvestigationObservationResources() (observations []Observation, err error) {
	if c.IncludedInvestigationObservationResources == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *c.IncludedInvestigationObservationResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedInvestigationDiagnosticReportResources() (diagnosticReports []DiagnosticReport, err error) {
	if c.IncludedInvestigationDiagnosticReportResources == nil {
		err = errors.New("Included diagnosticReports not requested")
	} else {
		diagnosticReports = *c.IncludedInvestigationDiagnosticReportResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedInvestigationQuestionnaireResponseResources() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.IncludedInvestigationQuestionnaireResponseResources == nil {
		err = errors.New("Included questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.IncludedInvestigationQuestionnaireResponseResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedActionAppointmentResources() (appointments []Appointment, err error) {
	if c.IncludedActionAppointmentResources == nil {
		err = errors.New("Included appointments not requested")
	} else {
		appointments = *c.IncludedActionAppointmentResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedActionReferralRequestResources() (referralRequests []ReferralRequest, err error) {
	if c.IncludedActionReferralRequestResources == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *c.IncludedActionReferralRequestResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedActionNutritionOrderResources() (nutritionOrders []NutritionOrder, err error) {
	if c.IncludedActionNutritionOrderResources == nil {
		err = errors.New("Included nutritionOrders not requested")
	} else {
		nutritionOrders = *c.IncludedActionNutritionOrderResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedActionProcedureRequestResources() (procedureRequests []ProcedureRequest, err error) {
	if c.IncludedActionProcedureRequestResources == nil {
		err = errors.New("Included procedureRequests not requested")
	} else {
		procedureRequests = *c.IncludedActionProcedureRequestResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedActionProcedureResources() (procedures []Procedure, err error) {
	if c.IncludedActionProcedureResources == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *c.IncludedActionProcedureResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedActionDiagnosticOrderResources() (diagnosticOrders []DiagnosticOrder, err error) {
	if c.IncludedActionDiagnosticOrderResources == nil {
		err = errors.New("Included diagnosticOrders not requested")
	} else {
		diagnosticOrders = *c.IncludedActionDiagnosticOrderResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedActionMedicationOrderResources() (medicationOrders []MedicationOrder, err error) {
	if c.IncludedActionMedicationOrderResources == nil {
		err = errors.New("Included medicationOrders not requested")
	} else {
		medicationOrders = *c.IncludedActionMedicationOrderResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedActionSupplyRequestResources() (supplyRequests []SupplyRequest, err error) {
	if c.IncludedActionSupplyRequestResources == nil {
		err = errors.New("Included supplyRequests not requested")
	} else {
		supplyRequests = *c.IncludedActionSupplyRequestResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanAppointmentResources() (appointments []Appointment, err error) {
	if c.IncludedPlanAppointmentResources == nil {
		err = errors.New("Included appointments not requested")
	} else {
		appointments = *c.IncludedPlanAppointmentResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanOrderResources() (orders []Order, err error) {
	if c.IncludedPlanOrderResources == nil {
		err = errors.New("Included orders not requested")
	} else {
		orders = *c.IncludedPlanOrderResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanReferralRequestResources() (referralRequests []ReferralRequest, err error) {
	if c.IncludedPlanReferralRequestResources == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *c.IncludedPlanReferralRequestResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanProcessRequestResources() (processRequests []ProcessRequest, err error) {
	if c.IncludedPlanProcessRequestResources == nil {
		err = errors.New("Included processRequests not requested")
	} else {
		processRequests = *c.IncludedPlanProcessRequestResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanVisionPrescriptionResources() (visionPrescriptions []VisionPrescription, err error) {
	if c.IncludedPlanVisionPrescriptionResources == nil {
		err = errors.New("Included visionPrescriptions not requested")
	} else {
		visionPrescriptions = *c.IncludedPlanVisionPrescriptionResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanDiagnosticOrderResources() (diagnosticOrders []DiagnosticOrder, err error) {
	if c.IncludedPlanDiagnosticOrderResources == nil {
		err = errors.New("Included diagnosticOrders not requested")
	} else {
		diagnosticOrders = *c.IncludedPlanDiagnosticOrderResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanProcedureRequestResources() (procedureRequests []ProcedureRequest, err error) {
	if c.IncludedPlanProcedureRequestResources == nil {
		err = errors.New("Included procedureRequests not requested")
	} else {
		procedureRequests = *c.IncludedPlanProcedureRequestResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanDeviceUseRequestResources() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.IncludedPlanDeviceUseRequestResources == nil {
		err = errors.New("Included deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.IncludedPlanDeviceUseRequestResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanSupplyRequestResources() (supplyRequests []SupplyRequest, err error) {
	if c.IncludedPlanSupplyRequestResources == nil {
		err = errors.New("Included supplyRequests not requested")
	} else {
		supplyRequests = *c.IncludedPlanSupplyRequestResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanCarePlanResources() (carePlans []CarePlan, err error) {
	if c.IncludedPlanCarePlanResources == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *c.IncludedPlanCarePlanResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanNutritionOrderResources() (nutritionOrders []NutritionOrder, err error) {
	if c.IncludedPlanNutritionOrderResources == nil {
		err = errors.New("Included nutritionOrders not requested")
	} else {
		nutritionOrders = *c.IncludedPlanNutritionOrderResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanMedicationOrderResources() (medicationOrders []MedicationOrder, err error) {
	if c.IncludedPlanMedicationOrderResources == nil {
		err = errors.New("Included medicationOrders not requested")
	} else {
		medicationOrders = *c.IncludedPlanMedicationOrderResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedPlanCommunicationRequestResources() (communicationRequests []CommunicationRequest, err error) {
	if c.IncludedPlanCommunicationRequestResources == nil {
		err = errors.New("Included communicationRequests not requested")
	} else {
		communicationRequests = *c.IncludedPlanCommunicationRequestResources
	}
	return
}

func (c *ClinicalImpressionPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPreviousResources != nil {
		for _, r := range *c.IncludedPreviousResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedAssessorResources != nil {
		for _, r := range *c.IncludedAssessorResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedProblemConditionResources != nil {
		for _, r := range *c.IncludedProblemConditionResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedProblemAllergyIntoleranceResources != nil {
		for _, r := range *c.IncludedProblemAllergyIntoleranceResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResources != nil {
		for _, r := range *c.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedInvestigationFamilyMemberHistoryResources != nil {
		for _, r := range *c.IncludedInvestigationFamilyMemberHistoryResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedInvestigationObservationResources != nil {
		for _, r := range *c.IncludedInvestigationObservationResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedInvestigationDiagnosticReportResources != nil {
		for _, r := range *c.IncludedInvestigationDiagnosticReportResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedInvestigationQuestionnaireResponseResources != nil {
		for _, r := range *c.IncludedInvestigationQuestionnaireResponseResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActionAppointmentResources != nil {
		for _, r := range *c.IncludedActionAppointmentResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActionReferralRequestResources != nil {
		for _, r := range *c.IncludedActionReferralRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActionNutritionOrderResources != nil {
		for _, r := range *c.IncludedActionNutritionOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActionProcedureRequestResources != nil {
		for _, r := range *c.IncludedActionProcedureRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActionProcedureResources != nil {
		for _, r := range *c.IncludedActionProcedureResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActionDiagnosticOrderResources != nil {
		for _, r := range *c.IncludedActionDiagnosticOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActionMedicationOrderResources != nil {
		for _, r := range *c.IncludedActionMedicationOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActionSupplyRequestResources != nil {
		for _, r := range *c.IncludedActionSupplyRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanAppointmentResources != nil {
		for _, r := range *c.IncludedPlanAppointmentResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanOrderResources != nil {
		for _, r := range *c.IncludedPlanOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanReferralRequestResources != nil {
		for _, r := range *c.IncludedPlanReferralRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanProcessRequestResources != nil {
		for _, r := range *c.IncludedPlanProcessRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanVisionPrescriptionResources != nil {
		for _, r := range *c.IncludedPlanVisionPrescriptionResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanDiagnosticOrderResources != nil {
		for _, r := range *c.IncludedPlanDiagnosticOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanProcedureRequestResources != nil {
		for _, r := range *c.IncludedPlanProcedureRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanDeviceUseRequestResources != nil {
		for _, r := range *c.IncludedPlanDeviceUseRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanSupplyRequestResources != nil {
		for _, r := range *c.IncludedPlanSupplyRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanCarePlanResources != nil {
		for _, r := range *c.IncludedPlanCarePlanResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanNutritionOrderResources != nil {
		for _, r := range *c.IncludedPlanNutritionOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanMedicationOrderResources != nil {
		for _, r := range *c.IncludedPlanMedicationOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPlanCommunicationRequestResources != nil {
		for _, r := range *c.IncludedPlanCommunicationRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
