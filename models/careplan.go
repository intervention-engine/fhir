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
	Participant    []CarePlanParticipantComponent `bson:"participant,omitempty" json:"participant,omitempty"`
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

type CarePlanParticipantComponent struct {
	BackboneElement `bson:",inline"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Member          *Reference       `bson:"member,omitempty" json:"member,omitempty"`
}

type CarePlanActivityComponent struct {
	BackboneElement `bson:",inline"`
	ActionResulting []Reference                      `bson:"actionResulting,omitempty" json:"actionResulting,omitempty"`
	Progress        []Annotation                     `bson:"progress,omitempty" json:"progress,omitempty"`
	Reference       *Reference                       `bson:"reference,omitempty" json:"reference,omitempty"`
	Detail          *CarePlanActivityDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type CarePlanActivityDetailComponent struct {
	BackboneElement        `bson:",inline"`
	Category               *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
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
	IncludedOrderResourcesReferencedByActivityreference                *[]Order                 `bson:"_includedOrderResourcesReferencedByActivityreference,omitempty"`
	IncludedReferralRequestResourcesReferencedByActivityreference      *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedProcessRequestResourcesReferencedByActivityreference       *[]ProcessRequest        `bson:"_includedProcessRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedNutritionOrderResourcesReferencedByActivityreference       *[]NutritionOrder        `bson:"_includedNutritionOrderResourcesReferencedByActivityreference,omitempty"`
	IncludedVisionPrescriptionResourcesReferencedByActivityreference   *[]VisionPrescription    `bson:"_includedVisionPrescriptionResourcesReferencedByActivityreference,omitempty"`
	IncludedDiagnosticOrderResourcesReferencedByActivityreference      *[]DiagnosticOrder       `bson:"_includedDiagnosticOrderResourcesReferencedByActivityreference,omitempty"`
	IncludedProcedureRequestResourcesReferencedByActivityreference     *[]ProcedureRequest      `bson:"_includedProcedureRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedDeviceUseRequestResourcesReferencedByActivityreference     *[]DeviceUseRequest      `bson:"_includedDeviceUseRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedMedicationOrderResourcesReferencedByActivityreference      *[]MedicationOrder       `bson:"_includedMedicationOrderResourcesReferencedByActivityreference,omitempty"`
	IncludedCommunicationRequestResourcesReferencedByActivityreference *[]CommunicationRequest  `bson:"_includedCommunicationRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedSupplyRequestResourcesReferencedByActivityreference        *[]SupplyRequest         `bson:"_includedSupplyRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer                 *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer                 *[]Organization          `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                      *[]Patient               `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedGoalResourcesReferencedByGoal                              *[]Goal                  `bson:"_includedGoalResourcesReferencedByGoal,omitempty"`
	IncludedGroupResourcesReferencedBySubject                          *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                        *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPractitionerResourcesReferencedByParticipant               *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByParticipant,omitempty"`
	IncludedOrganizationResourcesReferencedByParticipant               *[]Organization          `bson:"_includedOrganizationResourcesReferencedByParticipant,omitempty"`
	IncludedPatientResourcesReferencedByParticipant                    *[]Patient               `bson:"_includedPatientResourcesReferencedByParticipant,omitempty"`
	IncludedRelatedPersonResourcesReferencedByParticipant              *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByParticipant,omitempty"`
	IncludedCarePlanResourcesReferencedByRelatedplan                   *[]CarePlan              `bson:"_includedCarePlanResourcesReferencedByRelatedplan,omitempty"`
	IncludedConditionResourcesReferencedByCondition                    *[]Condition             `bson:"_includedConditionResourcesReferencedByCondition,omitempty"`
	IncludedPatientResourcesReferencedByPatient                        *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                    *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref          *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref          *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedCarePlanResourcesReferencingRelatedplan                 *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingRelatedplan,omitempty"`
	RevIncludedListResourcesReferencingItem                            *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref         *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                         *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                        *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference                 *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                  *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                    *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated             *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment            *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject        *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest              *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger           *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPlan              *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPlan,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                   *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
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

func (c *CarePlanPlusRelatedResources) GetIncludedOrderResourceReferencedByActivityreference() (order *Order, err error) {
	if c.IncludedOrderResourcesReferencedByActivityreference == nil {
		err = errors.New("Included orders not requested")
	} else if len(*c.IncludedOrderResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 order, but found %d", len(*c.IncludedOrderResourcesReferencedByActivityreference))
	} else if len(*c.IncludedOrderResourcesReferencedByActivityreference) == 1 {
		order = &(*c.IncludedOrderResourcesReferencedByActivityreference)[0]
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

func (c *CarePlanPlusRelatedResources) GetIncludedNutritionOrderResourceReferencedByActivityreference() (nutritionOrder *NutritionOrder, err error) {
	if c.IncludedNutritionOrderResourcesReferencedByActivityreference == nil {
		err = errors.New("Included nutritionorders not requested")
	} else if len(*c.IncludedNutritionOrderResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 nutritionOrder, but found %d", len(*c.IncludedNutritionOrderResourcesReferencedByActivityreference))
	} else if len(*c.IncludedNutritionOrderResourcesReferencedByActivityreference) == 1 {
		nutritionOrder = &(*c.IncludedNutritionOrderResourcesReferencedByActivityreference)[0]
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

func (c *CarePlanPlusRelatedResources) GetIncludedDiagnosticOrderResourceReferencedByActivityreference() (diagnosticOrder *DiagnosticOrder, err error) {
	if c.IncludedDiagnosticOrderResourcesReferencedByActivityreference == nil {
		err = errors.New("Included diagnosticorders not requested")
	} else if len(*c.IncludedDiagnosticOrderResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 diagnosticOrder, but found %d", len(*c.IncludedDiagnosticOrderResourcesReferencedByActivityreference))
	} else if len(*c.IncludedDiagnosticOrderResourcesReferencedByActivityreference) == 1 {
		diagnosticOrder = &(*c.IncludedDiagnosticOrderResourcesReferencedByActivityreference)[0]
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

func (c *CarePlanPlusRelatedResources) GetIncludedPractitionerResourceReferencedByParticipant() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByParticipant == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByParticipant))
	} else if len(*c.IncludedPractitionerResourcesReferencedByParticipant) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByParticipant)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedOrganizationResourceReferencedByParticipant() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByParticipant == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByParticipant))
	} else if len(*c.IncludedOrganizationResourcesReferencedByParticipant) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByParticipant)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPatientResourceReferencedByParticipant() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByParticipant == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByParticipant))
	} else if len(*c.IncludedPatientResourcesReferencedByParticipant) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByParticipant)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByParticipant() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByParticipant == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedByParticipant))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByParticipant) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedByParticipant)[0]
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingReference
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
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
	if c.IncludedOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedReferralRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedReferralRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedReferralRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcessRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedProcessRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedProcessRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedVisionPrescriptionResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedVisionPrescriptionResourcesReferencedByActivityreference {
			rsc := (*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDiagnosticOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDiagnosticOrderResourcesReferencedByActivityreference)[idx]
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
	if c.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByParticipant {
			rsc := (*c.IncludedPractitionerResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByParticipant {
			rsc := (*c.IncludedOrganizationResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByParticipant {
			rsc := (*c.IncludedPatientResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByParticipant {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByParticipant)[idx]
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
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
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

func (c *CarePlanPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedAppointmentResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedAppointmentResourcesReferencedByActivityreference {
			rsc := (*c.IncludedAppointmentResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedReferralRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedReferralRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedReferralRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedProcessRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedProcessRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedProcessRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedVisionPrescriptionResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedVisionPrescriptionResourcesReferencedByActivityreference {
			rsc := (*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDiagnosticOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDiagnosticOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDiagnosticOrderResourcesReferencedByActivityreference)[idx]
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
	if c.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByParticipant {
			rsc := (*c.IncludedPractitionerResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByParticipant {
			rsc := (*c.IncludedOrganizationResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByParticipant {
			rsc := (*c.IncludedPatientResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByParticipant {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByParticipant)[idx]
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
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
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
