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
	DomainResource           `bson:",inline"`
	Identifier               []Identifier                                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                   string                                      `bson:"status,omitempty" json:"status,omitempty"`
	Code                     *CodeableConcept                            `bson:"code,omitempty" json:"code,omitempty"`
	Description              string                                      `bson:"description,omitempty" json:"description,omitempty"`
	Subject                  *Reference                                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Assessor                 *Reference                                  `bson:"assessor,omitempty" json:"assessor,omitempty"`
	Date                     *FHIRDateTime                               `bson:"date,omitempty" json:"date,omitempty"`
	EffectiveDateTime        *FHIRDateTime                               `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod          *Period                                     `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Context                  *Reference                                  `bson:"context,omitempty" json:"context,omitempty"`
	Previous                 *Reference                                  `bson:"previous,omitempty" json:"previous,omitempty"`
	Problem                  []Reference                                 `bson:"problem,omitempty" json:"problem,omitempty"`
	Investigations           []ClinicalImpressionInvestigationsComponent `bson:"investigations,omitempty" json:"investigations,omitempty"`
	Protocol                 []string                                    `bson:"protocol,omitempty" json:"protocol,omitempty"`
	Summary                  string                                      `bson:"summary,omitempty" json:"summary,omitempty"`
	Finding                  []ClinicalImpressionFindingComponent        `bson:"finding,omitempty" json:"finding,omitempty"`
	PrognosisCodeableConcept []CodeableConcept                           `bson:"prognosisCodeableConcept,omitempty" json:"prognosisCodeableConcept,omitempty"`
	PrognosisReference       []Reference                                 `bson:"prognosisReference,omitempty" json:"prognosisReference,omitempty"`
	Plan                     []Reference                                 `bson:"plan,omitempty" json:"plan,omitempty"`
	Action                   []Reference                                 `bson:"action,omitempty" json:"action,omitempty"`
	Note                     []Annotation                                `bson:"note,omitempty" json:"note,omitempty"`
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
	BackboneElement     `bson:",inline"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
	Cause               string           `bson:"cause,omitempty" json:"cause,omitempty"`
}

type ClinicalImpressionPlus struct {
	ClinicalImpression                     `bson:",inline"`
	ClinicalImpressionPlusRelatedResources `bson:",inline"`
}

type ClinicalImpressionPlusRelatedResources struct {
	IncludedClinicalImpressionResourcesReferencedByPrevious         *[]ClinicalImpression    `bson:"_includedClinicalImpressionResourcesReferencedByPrevious,omitempty"`
	IncludedPractitionerResourcesReferencedByAssessor               *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAssessor,omitempty"`
	IncludedGroupResourcesReferencedBySubject                       *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                     *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedConditionResourcesReferencedByFindingref                *[]Condition             `bson:"_includedConditionResourcesReferencedByFindingref,omitempty"`
	IncludedObservationResourcesReferencedByFindingref              *[]Observation           `bson:"_includedObservationResourcesReferencedByFindingref,omitempty"`
	IncludedConditionResourcesReferencedByProblem                   *[]Condition             `bson:"_includedConditionResourcesReferencedByProblem,omitempty"`
	IncludedAllergyIntoleranceResourcesReferencedByProblem          *[]AllergyIntolerance    `bson:"_includedAllergyIntoleranceResourcesReferencedByProblem,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext               *[]EpisodeOfCare         `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext                   *[]Encounter             `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
	IncludedRiskAssessmentResourcesReferencedByInvestigation        *[]RiskAssessment        `bson:"_includedRiskAssessmentResourcesReferencedByInvestigation,omitempty"`
	IncludedFamilyMemberHistoryResourcesReferencedByInvestigation   *[]FamilyMemberHistory   `bson:"_includedFamilyMemberHistoryResourcesReferencedByInvestigation,omitempty"`
	IncludedObservationResourcesReferencedByInvestigation           *[]Observation           `bson:"_includedObservationResourcesReferencedByInvestigation,omitempty"`
	IncludedDiagnosticReportResourcesReferencedByInvestigation      *[]DiagnosticReport      `bson:"_includedDiagnosticReportResourcesReferencedByInvestigation,omitempty"`
	IncludedImagingStudyResourcesReferencedByInvestigation          *[]ImagingStudy          `bson:"_includedImagingStudyResourcesReferencedByInvestigation,omitempty"`
	IncludedQuestionnaireResponseResourcesReferencedByInvestigation *[]QuestionnaireResponse `bson:"_includedQuestionnaireResponseResourcesReferencedByInvestigation,omitempty"`
	IncludedAppointmentResourcesReferencedByAction                  *[]Appointment           `bson:"_includedAppointmentResourcesReferencedByAction,omitempty"`
	IncludedReferralRequestResourcesReferencedByAction              *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByAction,omitempty"`
	IncludedNutritionRequestResourcesReferencedByAction             *[]NutritionRequest      `bson:"_includedNutritionRequestResourcesReferencedByAction,omitempty"`
	IncludedProcedureRequestResourcesReferencedByAction             *[]ProcedureRequest      `bson:"_includedProcedureRequestResourcesReferencedByAction,omitempty"`
	IncludedProcedureResourcesReferencedByAction                    *[]Procedure             `bson:"_includedProcedureResourcesReferencedByAction,omitempty"`
	IncludedMedicationOrderResourcesReferencedByAction              *[]MedicationOrder       `bson:"_includedMedicationOrderResourcesReferencedByAction,omitempty"`
	IncludedDiagnosticRequestResourcesReferencedByAction            *[]DiagnosticRequest     `bson:"_includedDiagnosticRequestResourcesReferencedByAction,omitempty"`
	IncludedSupplyRequestResourcesReferencedByAction                *[]SupplyRequest         `bson:"_includedSupplyRequestResourcesReferencedByAction,omitempty"`
	IncludedAppointmentResourcesReferencedByPlan                    *[]Appointment           `bson:"_includedAppointmentResourcesReferencedByPlan,omitempty"`
	IncludedReferralRequestResourcesReferencedByPlan                *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByPlan,omitempty"`
	IncludedCarePlanResourcesReferencedByPlan                       *[]CarePlan              `bson:"_includedCarePlanResourcesReferencedByPlan,omitempty"`
	IncludedNutritionRequestResourcesReferencedByPlan               *[]NutritionRequest      `bson:"_includedNutritionRequestResourcesReferencedByPlan,omitempty"`
	IncludedProcessRequestResourcesReferencedByPlan                 *[]ProcessRequest        `bson:"_includedProcessRequestResourcesReferencedByPlan,omitempty"`
	IncludedVisionPrescriptionResourcesReferencedByPlan             *[]VisionPrescription    `bson:"_includedVisionPrescriptionResourcesReferencedByPlan,omitempty"`
	IncludedProcedureRequestResourcesReferencedByPlan               *[]ProcedureRequest      `bson:"_includedProcedureRequestResourcesReferencedByPlan,omitempty"`
	IncludedDeviceUseRequestResourcesReferencedByPlan               *[]DeviceUseRequest      `bson:"_includedDeviceUseRequestResourcesReferencedByPlan,omitempty"`
	IncludedMedicationOrderResourcesReferencedByPlan                *[]MedicationOrder       `bson:"_includedMedicationOrderResourcesReferencedByPlan,omitempty"`
	IncludedDiagnosticRequestResourcesReferencedByPlan              *[]DiagnosticRequest     `bson:"_includedDiagnosticRequestResourcesReferencedByPlan,omitempty"`
	IncludedCommunicationRequestResourcesReferencedByPlan           *[]CommunicationRequest  `bson:"_includedCommunicationRequestResourcesReferencedByPlan,omitempty"`
	IncludedSupplyRequestResourcesReferencedByPlan                  *[]SupplyRequest         `bson:"_includedSupplyRequestResourcesReferencedByPlan,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                      *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref      *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                   *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                    *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequestreference    *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequestreference,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponsereference   *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponsereference,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource      *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon             *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                 *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                        *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces        *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon         *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition      *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces         *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon          *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition       *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                     *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                 *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject               *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                 *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated          *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject     *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestreference  *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequestreference,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPrevious       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPrevious,omitempty"`
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

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if c.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*c.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*c.IncludedGroupResourcesReferencedBySubject))
	} else if len(*c.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*c.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedBySubject))
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedConditionResourceReferencedByFindingref() (condition *Condition, err error) {
	if c.IncludedConditionResourcesReferencedByFindingref == nil {
		err = errors.New("Included conditions not requested")
	} else if len(*c.IncludedConditionResourcesReferencedByFindingref) > 1 {
		err = fmt.Errorf("Expected 0 or 1 condition, but found %d", len(*c.IncludedConditionResourcesReferencedByFindingref))
	} else if len(*c.IncludedConditionResourcesReferencedByFindingref) == 1 {
		condition = &(*c.IncludedConditionResourcesReferencedByFindingref)[0]
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedObservationResourceReferencedByFindingref() (observation *Observation, err error) {
	if c.IncludedObservationResourcesReferencedByFindingref == nil {
		err = errors.New("Included observations not requested")
	} else if len(*c.IncludedObservationResourcesReferencedByFindingref) > 1 {
		err = fmt.Errorf("Expected 0 or 1 observation, but found %d", len(*c.IncludedObservationResourcesReferencedByFindingref))
	} else if len(*c.IncludedObservationResourcesReferencedByFindingref) == 1 {
		observation = &(*c.IncludedObservationResourcesReferencedByFindingref)[0]
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

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByContext() (episodeOfCare *EpisodeOfCare, err error) {
	if c.IncludedEpisodeOfCareResourcesReferencedByContext == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*c.IncludedEpisodeOfCareResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*c.IncludedEpisodeOfCareResourcesReferencedByContext))
	} else if len(*c.IncludedEpisodeOfCareResourcesReferencedByContext) == 1 {
		episodeOfCare = &(*c.IncludedEpisodeOfCareResourcesReferencedByContext)[0]
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedEncounterResourceReferencedByContext() (encounter *Encounter, err error) {
	if c.IncludedEncounterResourcesReferencedByContext == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*c.IncludedEncounterResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*c.IncludedEncounterResourcesReferencedByContext))
	} else if len(*c.IncludedEncounterResourcesReferencedByContext) == 1 {
		encounter = &(*c.IncludedEncounterResourcesReferencedByContext)[0]
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedRiskAssessmentResourcesReferencedByInvestigation() (riskAssessments []RiskAssessment, err error) {
	if c.IncludedRiskAssessmentResourcesReferencedByInvestigation == nil {
		err = errors.New("Included riskAssessments not requested")
	} else {
		riskAssessments = *c.IncludedRiskAssessmentResourcesReferencedByInvestigation
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

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedImagingStudyResourcesReferencedByInvestigation() (imagingStudies []ImagingStudy, err error) {
	if c.IncludedImagingStudyResourcesReferencedByInvestigation == nil {
		err = errors.New("Included imagingStudies not requested")
	} else {
		imagingStudies = *c.IncludedImagingStudyResourcesReferencedByInvestigation
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

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedNutritionRequestResourcesReferencedByAction() (nutritionRequests []NutritionRequest, err error) {
	if c.IncludedNutritionRequestResourcesReferencedByAction == nil {
		err = errors.New("Included nutritionRequests not requested")
	} else {
		nutritionRequests = *c.IncludedNutritionRequestResourcesReferencedByAction
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

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedMedicationOrderResourcesReferencedByAction() (medicationOrders []MedicationOrder, err error) {
	if c.IncludedMedicationOrderResourcesReferencedByAction == nil {
		err = errors.New("Included medicationOrders not requested")
	} else {
		medicationOrders = *c.IncludedMedicationOrderResourcesReferencedByAction
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedDiagnosticRequestResourcesReferencedByAction() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.IncludedDiagnosticRequestResourcesReferencedByAction == nil {
		err = errors.New("Included diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.IncludedDiagnosticRequestResourcesReferencedByAction
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

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedReferralRequestResourcesReferencedByPlan() (referralRequests []ReferralRequest, err error) {
	if c.IncludedReferralRequestResourcesReferencedByPlan == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *c.IncludedReferralRequestResourcesReferencedByPlan
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

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedNutritionRequestResourcesReferencedByPlan() (nutritionRequests []NutritionRequest, err error) {
	if c.IncludedNutritionRequestResourcesReferencedByPlan == nil {
		err = errors.New("Included nutritionRequests not requested")
	} else {
		nutritionRequests = *c.IncludedNutritionRequestResourcesReferencedByPlan
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

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedMedicationOrderResourcesReferencedByPlan() (medicationOrders []MedicationOrder, err error) {
	if c.IncludedMedicationOrderResourcesReferencedByPlan == nil {
		err = errors.New("Included medicationOrders not requested")
	} else {
		medicationOrders = *c.IncludedMedicationOrderResourcesReferencedByPlan
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedDiagnosticRequestResourcesReferencedByPlan() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.IncludedDiagnosticRequestResourcesReferencedByPlan == nil {
		err = errors.New("Included diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.IncludedDiagnosticRequestResourcesReferencedByPlan
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

func (c *ClinicalImpressionPlusRelatedResources) GetIncludedSupplyRequestResourcesReferencedByPlan() (supplyRequests []SupplyRequest, err error) {
	if c.IncludedSupplyRequestResourcesReferencedByPlan == nil {
		err = errors.New("Included supplyRequests not requested")
	} else {
		supplyRequests = *c.IncludedSupplyRequestResourcesReferencedByPlan
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

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingData
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

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingBasedon
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

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingFocus
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

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingDefinition
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

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
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

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequestreference
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
	if c.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedGroupResourcesReferencedBySubject {
			rsc := (*c.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubject {
			rsc := (*c.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByFindingref != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByFindingref {
			rsc := (*c.IncludedConditionResourcesReferencedByFindingref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedObservationResourcesReferencedByFindingref != nil {
		for idx := range *c.IncludedObservationResourcesReferencedByFindingref {
			rsc := (*c.IncludedObservationResourcesReferencedByFindingref)[idx]
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
	if c.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *c.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*c.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByContext {
			rsc := (*c.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRiskAssessmentResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedRiskAssessmentResourcesReferencedByInvestigation {
			rsc := (*c.IncludedRiskAssessmentResourcesReferencedByInvestigation)[idx]
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
	if c.IncludedImagingStudyResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedImagingStudyResourcesReferencedByInvestigation {
			rsc := (*c.IncludedImagingStudyResourcesReferencedByInvestigation)[idx]
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
	if c.IncludedNutritionRequestResourcesReferencedByAction != nil {
		for idx := range *c.IncludedNutritionRequestResourcesReferencedByAction {
			rsc := (*c.IncludedNutritionRequestResourcesReferencedByAction)[idx]
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
	if c.IncludedMedicationOrderResourcesReferencedByAction != nil {
		for idx := range *c.IncludedMedicationOrderResourcesReferencedByAction {
			rsc := (*c.IncludedMedicationOrderResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticRequestResourcesReferencedByAction != nil {
		for idx := range *c.IncludedDiagnosticRequestResourcesReferencedByAction {
			rsc := (*c.IncludedDiagnosticRequestResourcesReferencedByAction)[idx]
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
	if c.IncludedReferralRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedReferralRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedReferralRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByPlan {
			rsc := (*c.IncludedCarePlanResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedNutritionRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedNutritionRequestResourcesReferencedByPlan)[idx]
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
	if c.IncludedMedicationOrderResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedMedicationOrderResourcesReferencedByPlan {
			rsc := (*c.IncludedMedicationOrderResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedDiagnosticRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedDiagnosticRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSupplyRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedSupplyRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedSupplyRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ClinicalImpressionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
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
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
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
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingPrevious != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingPrevious {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingPrevious)[idx]
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
	if c.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedGroupResourcesReferencedBySubject {
			rsc := (*c.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubject {
			rsc := (*c.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByFindingref != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByFindingref {
			rsc := (*c.IncludedConditionResourcesReferencedByFindingref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedObservationResourcesReferencedByFindingref != nil {
		for idx := range *c.IncludedObservationResourcesReferencedByFindingref {
			rsc := (*c.IncludedObservationResourcesReferencedByFindingref)[idx]
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
	if c.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *c.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*c.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByContext {
			rsc := (*c.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRiskAssessmentResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedRiskAssessmentResourcesReferencedByInvestigation {
			rsc := (*c.IncludedRiskAssessmentResourcesReferencedByInvestigation)[idx]
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
	if c.IncludedImagingStudyResourcesReferencedByInvestigation != nil {
		for idx := range *c.IncludedImagingStudyResourcesReferencedByInvestigation {
			rsc := (*c.IncludedImagingStudyResourcesReferencedByInvestigation)[idx]
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
	if c.IncludedNutritionRequestResourcesReferencedByAction != nil {
		for idx := range *c.IncludedNutritionRequestResourcesReferencedByAction {
			rsc := (*c.IncludedNutritionRequestResourcesReferencedByAction)[idx]
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
	if c.IncludedMedicationOrderResourcesReferencedByAction != nil {
		for idx := range *c.IncludedMedicationOrderResourcesReferencedByAction {
			rsc := (*c.IncludedMedicationOrderResourcesReferencedByAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticRequestResourcesReferencedByAction != nil {
		for idx := range *c.IncludedDiagnosticRequestResourcesReferencedByAction {
			rsc := (*c.IncludedDiagnosticRequestResourcesReferencedByAction)[idx]
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
	if c.IncludedReferralRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedReferralRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedReferralRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByPlan {
			rsc := (*c.IncludedCarePlanResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedNutritionRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedNutritionRequestResourcesReferencedByPlan)[idx]
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
	if c.IncludedMedicationOrderResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedMedicationOrderResourcesReferencedByPlan {
			rsc := (*c.IncludedMedicationOrderResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedDiagnosticRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedDiagnosticRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSupplyRequestResourcesReferencedByPlan != nil {
		for idx := range *c.IncludedSupplyRequestResourcesReferencedByPlan {
			rsc := (*c.IncludedSupplyRequestResourcesReferencedByPlan)[idx]
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
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
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
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingPrevious != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingPrevious {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingPrevious)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
