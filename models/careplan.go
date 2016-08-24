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

type CarePlan struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject        *Reference                     `bson:"subject,omitempty" json:"subject,omitempty"`
	Status         string                         `bson:"status,omitempty" json:"status,omitempty"`
	Context        *Reference                     `bson:"context,omitempty" json:"context,omitempty"`
	Period         *Period                        `bson:"period,omitempty" json:"period,omitempty"`
	Author         []Reference                    `bson:"author,omitempty" json:"author,omitempty"`
	Modified       *FHIRDateTime                  `bson:"modified,omitempty" json:"modified,omitempty"`
	Category       []CodeableConcept              `bson:"category,omitempty" json:"category,omitempty"`
	Description    string                         `bson:"description,omitempty" json:"description,omitempty"`
	Addresses      []Reference                    `bson:"addresses,omitempty" json:"addresses,omitempty"`
	Support        []Reference                    `bson:"support,omitempty" json:"support,omitempty"`
	RelatedPlan    []CarePlanRelatedPlanComponent `bson:"relatedPlan,omitempty" json:"relatedPlan,omitempty"`
	CareTeam       []Reference                    `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	Goal           []Reference                    `bson:"goal,omitempty" json:"goal,omitempty"`
	Activity       []CarePlanActivityComponent    `bson:"activity,omitempty" json:"activity,omitempty"`
	Note           *Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *CarePlan) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "CarePlan"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to CarePlan), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *CarePlan) GetBSON() (interface{}, error) {
	x.ResourceType = "CarePlan"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "carePlan" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type carePlan CarePlan

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *CarePlan) UnmarshalJSON(data []byte) (err error) {
	x2 := carePlan{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = CarePlan(x2)
		return x.checkResourceType()
	}
	return
}

func (x *CarePlan) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "CarePlan"
	} else if x.ResourceType != "CarePlan" {
		return errors.New(fmt.Sprintf("Expected resourceType to be CarePlan, instead received %s", x.ResourceType))
	}
	return nil
}

type CarePlanRelatedPlanComponent struct {
	BackboneElement `bson:",inline"`
	Code            string     `bson:"code,omitempty" json:"code,omitempty"`
	Plan            *Reference `bson:"plan,omitempty" json:"plan,omitempty"`
}

type CarePlanActivityComponent struct {
	BackboneElement `bson:",inline"`
	ActionResulting []Reference                      `bson:"actionResulting,omitempty" json:"actionResulting,omitempty"`
	Outcome         *CodeableConcept                 `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Progress        []Annotation                     `bson:"progress,omitempty" json:"progress,omitempty"`
	Reference       *Reference                       `bson:"reference,omitempty" json:"reference,omitempty"`
	Detail          *CarePlanActivityDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type CarePlanActivityDetailComponent struct {
	BackboneElement        `bson:",inline"`
	Category               *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Definition             *Reference        `bson:"definition,omitempty" json:"definition,omitempty"`
	Code                   *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	ReasonCode             []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference        []Reference       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Goal                   []Reference       `bson:"goal,omitempty" json:"goal,omitempty"`
	Status                 string            `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason           *CodeableConcept  `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Prohibited             *bool             `bson:"prohibited,omitempty" json:"prohibited,omitempty"`
	ScheduledTiming        *Timing           `bson:"scheduledTiming,omitempty" json:"scheduledTiming,omitempty"`
	ScheduledPeriod        *Period           `bson:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	ScheduledString        string            `bson:"scheduledString,omitempty" json:"scheduledString,omitempty"`
	Location               *Reference        `bson:"location,omitempty" json:"location,omitempty"`
	Performer              []Reference       `bson:"performer,omitempty" json:"performer,omitempty"`
	ProductCodeableConcept *CodeableConcept  `bson:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty"`
	ProductReference       *Reference        `bson:"productReference,omitempty" json:"productReference,omitempty"`
	DailyAmount            *Quantity         `bson:"dailyAmount,omitempty" json:"dailyAmount,omitempty"`
	Quantity               *Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Description            string            `bson:"description,omitempty" json:"description,omitempty"`
}

type CarePlanPlus struct {
	CarePlan                     `bson:",inline"`
	CarePlanPlusRelatedResources `bson:",inline"`
}

type CarePlanPlusRelatedResources struct {
	IncludedAppointmentResourcesReferencedByActivityreference          *[]Appointment           `bson:"_includedAppointmentResourcesReferencedByActivityreference,omitempty"`
	IncludedReferralRequestResourcesReferencedByActivityreference      *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedNutritionRequestResourcesReferencedByActivityreference     *[]NutritionRequest      `bson:"_includedNutritionRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedProcessRequestResourcesReferencedByActivityreference       *[]ProcessRequest        `bson:"_includedProcessRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedVisionPrescriptionResourcesReferencedByActivityreference   *[]VisionPrescription    `bson:"_includedVisionPrescriptionResourcesReferencedByActivityreference,omitempty"`
	IncludedProcedureRequestResourcesReferencedByActivityreference     *[]ProcedureRequest      `bson:"_includedProcedureRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedDeviceUseRequestResourcesReferencedByActivityreference     *[]DeviceUseRequest      `bson:"_includedDeviceUseRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedMedicationOrderResourcesReferencedByActivityreference      *[]MedicationOrder       `bson:"_includedMedicationOrderResourcesReferencedByActivityreference,omitempty"`
	IncludedDiagnosticRequestResourcesReferencedByActivityreference    *[]DiagnosticRequest     `bson:"_includedDiagnosticRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedCommunicationRequestResourcesReferencedByActivityreference *[]CommunicationRequest  `bson:"_includedCommunicationRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedSupplyRequestResourcesReferencedByActivityreference        *[]SupplyRequest         `bson:"_includedSupplyRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer                 *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer                 *[]Organization          `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                      *[]Patient               `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedGoalResourcesReferencedByGoal                              *[]Goal                  `bson:"_includedGoalResourcesReferencedByGoal,omitempty"`
	IncludedGroupResourcesReferencedBySubject                          *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                        *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedCareTeamResourcesReferencedByCareteam                      *[]CareTeam              `bson:"_includedCareTeamResourcesReferencedByCareteam,omitempty"`
	IncludedCarePlanResourcesReferencedByRelatedplan                   *[]CarePlan              `bson:"_includedCarePlanResourcesReferencedByRelatedplan,omitempty"`
	IncludedConditionResourcesReferencedByCondition                    *[]Condition             `bson:"_includedConditionResourcesReferencedByCondition,omitempty"`
	IncludedPatientResourcesReferencedByPatient                        *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	RevIncludedReferralRequestResourcesReferencingBasedon              *[]ReferralRequest       `bson:"_revIncludedReferralRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref          *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref          *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                         *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref         *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                      *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                     *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                       *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequestreference       *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequestreference,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponsereference      *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponsereference,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource         *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingBasedon                 *[]ImagingStudy          `bson:"_revIncludedImagingStudyResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                   *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                    *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                           *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedCarePlanResourcesReferencingRelatedplan                 *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingRelatedplan,omitempty"`
	RevIncludedListResourcesReferencingItem                            *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces           *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon            *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition         *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces            *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon             *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition          *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                        *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                    *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                  *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                    *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated             *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingBasedon        *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingBasedon,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject        *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestreference     *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequestreference,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPlan              *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPlan,omitempty"`
}

func (c *CarePlanPlusRelatedResources) GetIncludedAppointmentResourceReferencedByActivityreference() (appointment *Appointment, err error) {
	if c.IncludedAppointmentResourcesReferencedByActivityreference == nil {
		err = errors.New("Included appointments not requested")
	} else if len(*c.IncludedAppointmentResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 appointment, but found %d", len(*c.IncludedAppointmentResourcesReferencedByActivityreference))
	} else if len(*c.IncludedAppointmentResourcesReferencedByActivityreference) == 1 {
		appointment = &(*c.IncludedAppointmentResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedReferralRequestResourceReferencedByActivityreference() (referralRequest *ReferralRequest, err error) {
	if c.IncludedReferralRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included referralrequests not requested")
	} else if len(*c.IncludedReferralRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 referralRequest, but found %d", len(*c.IncludedReferralRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedReferralRequestResourcesReferencedByActivityreference) == 1 {
		referralRequest = &(*c.IncludedReferralRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedNutritionRequestResourceReferencedByActivityreference() (nutritionRequest *NutritionRequest, err error) {
	if c.IncludedNutritionRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included nutritionrequests not requested")
	} else if len(*c.IncludedNutritionRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 nutritionRequest, but found %d", len(*c.IncludedNutritionRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedNutritionRequestResourcesReferencedByActivityreference) == 1 {
		nutritionRequest = &(*c.IncludedNutritionRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedProcessRequestResourceReferencedByActivityreference() (processRequest *ProcessRequest, err error) {
	if c.IncludedProcessRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included processrequests not requested")
	} else if len(*c.IncludedProcessRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 processRequest, but found %d", len(*c.IncludedProcessRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedProcessRequestResourcesReferencedByActivityreference) == 1 {
		processRequest = &(*c.IncludedProcessRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedVisionPrescriptionResourceReferencedByActivityreference() (visionPrescription *VisionPrescription, err error) {
	if c.IncludedVisionPrescriptionResourcesReferencedByActivityreference == nil {
		err = errors.New("Included visionprescriptions not requested")
	} else if len(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 visionPrescription, but found %d", len(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference))
	} else if len(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference) == 1 {
		visionPrescription = &(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedProcedureRequestResourceReferencedByActivityreference() (procedureRequest *ProcedureRequest, err error) {
	if c.IncludedProcedureRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included procedurerequests not requested")
	} else if len(*c.IncludedProcedureRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 procedureRequest, but found %d", len(*c.IncludedProcedureRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedProcedureRequestResourcesReferencedByActivityreference) == 1 {
		procedureRequest = &(*c.IncludedProcedureRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedDeviceUseRequestResourceReferencedByActivityreference() (deviceUseRequest *DeviceUseRequest, err error) {
	if c.IncludedDeviceUseRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included deviceuserequests not requested")
	} else if len(*c.IncludedDeviceUseRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceUseRequest, but found %d", len(*c.IncludedDeviceUseRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedDeviceUseRequestResourcesReferencedByActivityreference) == 1 {
		deviceUseRequest = &(*c.IncludedDeviceUseRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedMedicationOrderResourceReferencedByActivityreference() (medicationOrder *MedicationOrder, err error) {
	if c.IncludedMedicationOrderResourcesReferencedByActivityreference == nil {
		err = errors.New("Included medicationorders not requested")
	} else if len(*c.IncludedMedicationOrderResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medicationOrder, but found %d", len(*c.IncludedMedicationOrderResourcesReferencedByActivityreference))
	} else if len(*c.IncludedMedicationOrderResourcesReferencedByActivityreference) == 1 {
		medicationOrder = &(*c.IncludedMedicationOrderResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedDiagnosticRequestResourceReferencedByActivityreference() (diagnosticRequest *DiagnosticRequest, err error) {
	if c.IncludedDiagnosticRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included diagnosticrequests not requested")
	} else if len(*c.IncludedDiagnosticRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 diagnosticRequest, but found %d", len(*c.IncludedDiagnosticRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedDiagnosticRequestResourcesReferencedByActivityreference) == 1 {
		diagnosticRequest = &(*c.IncludedDiagnosticRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedCommunicationRequestResourceReferencedByActivityreference() (communicationRequest *CommunicationRequest, err error) {
	if c.IncludedCommunicationRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included communicationrequests not requested")
	} else if len(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 communicationRequest, but found %d", len(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference) == 1 {
		communicationRequest = &(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedSupplyRequestResourceReferencedByActivityreference() (supplyRequest *SupplyRequest, err error) {
	if c.IncludedSupplyRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included supplyrequests not requested")
	} else if len(*c.IncludedSupplyRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 supplyRequest, but found %d", len(*c.IncludedSupplyRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedSupplyRequestResourcesReferencedByActivityreference) == 1 {
		supplyRequest = &(*c.IncludedSupplyRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByPerformer() (practitioners []Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *c.IncludedPractitionerResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByPerformer() (organizations []Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *c.IncludedOrganizationResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPatientResourcesReferencedByPerformer() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByPerformer() (relatedPeople []RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *c.IncludedRelatedPersonResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedGoalResourcesReferencedByGoal() (goals []Goal, err error) {
	if c.IncludedGoalResourcesReferencedByGoal == nil {
		err = errors.New("Included goals not requested")
	} else {
		goals = *c.IncludedGoalResourcesReferencedByGoal
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if c.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*c.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*c.IncludedGroupResourcesReferencedBySubject))
	} else if len(*c.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*c.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedBySubject))
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedCareTeamResourcesReferencedByCareteam() (careTeams []CareTeam, err error) {
	if c.IncludedCareTeamResourcesReferencedByCareteam == nil {
		err = errors.New("Included careTeams not requested")
	} else {
		careTeams = *c.IncludedCareTeamResourcesReferencedByCareteam
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedCarePlanResourceReferencedByRelatedplan() (carePlan *CarePlan, err error) {
	if c.IncludedCarePlanResourcesReferencedByRelatedplan == nil {
		err = errors.New("Included careplans not requested")
	} else if len(*c.IncludedCarePlanResourcesReferencedByRelatedplan) > 1 {
		err = fmt.Errorf("Expected 0 or 1 carePlan, but found %d", len(*c.IncludedCarePlanResourcesReferencedByRelatedplan))
	} else if len(*c.IncludedCarePlanResourcesReferencedByRelatedplan) == 1 {
		carePlan = &(*c.IncludedCarePlanResourcesReferencedByRelatedplan)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedConditionResourcesReferencedByCondition() (conditions []Condition, err error) {
	if c.IncludedConditionResourcesReferencedByCondition == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *c.IncludedConditionResourcesReferencedByCondition
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPatient))
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingBasedon() (referralRequests []ReferralRequest, err error) {
	if c.RevIncludedReferralRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *c.RevIncludedReferralRequestResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingBasedon() (imagingStudies []ImagingStudy, err error) {
	if c.RevIncludedImagingStudyResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *c.RevIncludedImagingStudyResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingRelatedplan() (carePlans []CarePlan, err error) {
	if c.RevIncludedCarePlanResourcesReferencingRelatedplan == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *c.RevIncludedCarePlanResourcesReferencingRelatedplan
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingBasedon() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPlan() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingPlan == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingPlan
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedAppointmentResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedAppointmentResourcesReferencedByActivityreference {
			rsc := (*c.IncludedAppointmentResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedReferralRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedReferralRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedReferralRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedNutritionRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedNutritionRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcessRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedProcessRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedProcessRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedVisionPrescriptionResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedVisionPrescriptionResourcesReferencedByActivityreference {
			rsc := (*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcedureRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedProcedureRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedProcedureRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceUseRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDeviceUseRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDeviceUseRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMedicationOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedMedicationOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedMedicationOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDiagnosticRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDiagnosticRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSupplyRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedSupplyRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedSupplyRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*c.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*c.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGoalResourcesReferencedByGoal != nil {
		for idx := range *c.IncludedGoalResourcesReferencedByGoal {
			rsc := (*c.IncludedGoalResourcesReferencedByGoal)[idx]
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
	if c.IncludedCareTeamResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByCareteam {
			rsc := (*c.IncludedCareTeamResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByRelatedplan != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByRelatedplan {
			rsc := (*c.IncludedCarePlanResourcesReferencedByRelatedplan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByCondition != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByCondition {
			rsc := (*c.IncludedConditionResourcesReferencedByCondition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedReferralRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedReferralRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedReferralRequestResourcesReferencingBasedon)[idx]
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
	if c.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*c.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
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
	if c.RevIncludedCarePlanResourcesReferencingRelatedplan != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingRelatedplan {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingRelatedplan)[idx]
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
	if c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
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
	if c.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *CarePlanPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedAppointmentResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedAppointmentResourcesReferencedByActivityreference {
			rsc := (*c.IncludedAppointmentResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedReferralRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedReferralRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedReferralRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedNutritionRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedNutritionRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcessRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedProcessRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedProcessRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedVisionPrescriptionResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedVisionPrescriptionResourcesReferencedByActivityreference {
			rsc := (*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcedureRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedProcedureRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedProcedureRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceUseRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDeviceUseRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDeviceUseRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMedicationOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedMedicationOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedMedicationOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDiagnosticRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDiagnosticRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSupplyRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedSupplyRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedSupplyRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*c.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*c.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGoalResourcesReferencedByGoal != nil {
		for idx := range *c.IncludedGoalResourcesReferencedByGoal {
			rsc := (*c.IncludedGoalResourcesReferencedByGoal)[idx]
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
	if c.IncludedCareTeamResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByCareteam {
			rsc := (*c.IncludedCareTeamResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByRelatedplan != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByRelatedplan {
			rsc := (*c.IncludedCarePlanResourcesReferencedByRelatedplan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByCondition != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByCondition {
			rsc := (*c.IncludedConditionResourcesReferencedByCondition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedReferralRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedReferralRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedReferralRequestResourcesReferencingBasedon)[idx]
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
	if c.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*c.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
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
	if c.RevIncludedCarePlanResourcesReferencingRelatedplan != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingRelatedplan {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingRelatedplan)[idx]
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
	if c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
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
	if c.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
