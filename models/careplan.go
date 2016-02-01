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
	Code string     `bson:"code,omitempty" json:"code,omitempty"`
	Plan *Reference `bson:"plan,omitempty" json:"plan,omitempty"`
}

type CarePlanParticipantComponent struct {
	Role   *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Member *Reference       `bson:"member,omitempty" json:"member,omitempty"`
}

type CarePlanActivityComponent struct {
	ActionResulting []Reference                      `bson:"actionResulting,omitempty" json:"actionResulting,omitempty"`
	Progress        []Annotation                     `bson:"progress,omitempty" json:"progress,omitempty"`
	Reference       *Reference                       `bson:"reference,omitempty" json:"reference,omitempty"`
	Detail          *CarePlanActivityDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type CarePlanActivityDetailComponent struct {
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
	CarePlan             `bson:",inline"`
	CarePlanPlusIncludes `bson:",inline"`
}

type CarePlanPlusIncludes struct {
	IncludedActivityreferenceAppointmentResources          *[]Appointment          `bson:"_includedActivityreferenceAppointmentResources,omitempty"`
	IncludedActivityreferenceOrderResources                *[]Order                `bson:"_includedActivityreferenceOrderResources,omitempty"`
	IncludedActivityreferenceReferralRequestResources      *[]ReferralRequest      `bson:"_includedActivityreferenceReferralRequestResources,omitempty"`
	IncludedActivityreferenceProcessRequestResources       *[]ProcessRequest       `bson:"_includedActivityreferenceProcessRequestResources,omitempty"`
	IncludedActivityreferenceNutritionOrderResources       *[]NutritionOrder       `bson:"_includedActivityreferenceNutritionOrderResources,omitempty"`
	IncludedActivityreferenceVisionPrescriptionResources   *[]VisionPrescription   `bson:"_includedActivityreferenceVisionPrescriptionResources,omitempty"`
	IncludedActivityreferenceDiagnosticOrderResources      *[]DiagnosticOrder      `bson:"_includedActivityreferenceDiagnosticOrderResources,omitempty"`
	IncludedActivityreferenceProcedureRequestResources     *[]ProcedureRequest     `bson:"_includedActivityreferenceProcedureRequestResources,omitempty"`
	IncludedActivityreferenceDeviceUseRequestResources     *[]DeviceUseRequest     `bson:"_includedActivityreferenceDeviceUseRequestResources,omitempty"`
	IncludedActivityreferenceMedicationOrderResources      *[]MedicationOrder      `bson:"_includedActivityreferenceMedicationOrderResources,omitempty"`
	IncludedActivityreferenceCommunicationRequestResources *[]CommunicationRequest `bson:"_includedActivityreferenceCommunicationRequestResources,omitempty"`
	IncludedActivityreferenceSupplyRequestResources        *[]SupplyRequest        `bson:"_includedActivityreferenceSupplyRequestResources,omitempty"`
	IncludedPerformerPractitionerResources                 *[]Practitioner         `bson:"_includedPerformerPractitionerResources,omitempty"`
	IncludedPerformerOrganizationResources                 *[]Organization         `bson:"_includedPerformerOrganizationResources,omitempty"`
	IncludedPerformerPatientResources                      *[]Patient              `bson:"_includedPerformerPatientResources,omitempty"`
	IncludedPerformerRelatedPersonResources                *[]RelatedPerson        `bson:"_includedPerformerRelatedPersonResources,omitempty"`
	IncludedGoalResources                                  *[]Goal                 `bson:"_includedGoalResources,omitempty"`
	IncludedSubjectGroupResources                          *[]Group                `bson:"_includedSubjectGroupResources,omitempty"`
	IncludedSubjectPatientResources                        *[]Patient              `bson:"_includedSubjectPatientResources,omitempty"`
	IncludedParticipantPractitionerResources               *[]Practitioner         `bson:"_includedParticipantPractitionerResources,omitempty"`
	IncludedParticipantOrganizationResources               *[]Organization         `bson:"_includedParticipantOrganizationResources,omitempty"`
	IncludedParticipantPatientResources                    *[]Patient              `bson:"_includedParticipantPatientResources,omitempty"`
	IncludedParticipantRelatedPersonResources              *[]RelatedPerson        `bson:"_includedParticipantRelatedPersonResources,omitempty"`
	IncludedRelatedplanResources                           *[]CarePlan             `bson:"_includedRelatedplanResources,omitempty"`
	IncludedConditionResources                             *[]Condition            `bson:"_includedConditionResources,omitempty"`
	IncludedPatientResources                               *[]Patient              `bson:"_includedPatientResources,omitempty"`
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceAppointmentResource() (appointment *Appointment, err error) {
	if c.IncludedActivityreferenceAppointmentResources == nil {
		err = errors.New("Included appointments not requested")
	} else if len(*c.IncludedActivityreferenceAppointmentResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 appointment, but found %d", len(*c.IncludedActivityreferenceAppointmentResources))
	} else if len(*c.IncludedActivityreferenceAppointmentResources) == 1 {
		appointment = &(*c.IncludedActivityreferenceAppointmentResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceOrderResource() (order *Order, err error) {
	if c.IncludedActivityreferenceOrderResources == nil {
		err = errors.New("Included orders not requested")
	} else if len(*c.IncludedActivityreferenceOrderResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 order, but found %d", len(*c.IncludedActivityreferenceOrderResources))
	} else if len(*c.IncludedActivityreferenceOrderResources) == 1 {
		order = &(*c.IncludedActivityreferenceOrderResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceReferralRequestResource() (referralRequest *ReferralRequest, err error) {
	if c.IncludedActivityreferenceReferralRequestResources == nil {
		err = errors.New("Included referralrequests not requested")
	} else if len(*c.IncludedActivityreferenceReferralRequestResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 referralRequest, but found %d", len(*c.IncludedActivityreferenceReferralRequestResources))
	} else if len(*c.IncludedActivityreferenceReferralRequestResources) == 1 {
		referralRequest = &(*c.IncludedActivityreferenceReferralRequestResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceProcessRequestResource() (processRequest *ProcessRequest, err error) {
	if c.IncludedActivityreferenceProcessRequestResources == nil {
		err = errors.New("Included processrequests not requested")
	} else if len(*c.IncludedActivityreferenceProcessRequestResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 processRequest, but found %d", len(*c.IncludedActivityreferenceProcessRequestResources))
	} else if len(*c.IncludedActivityreferenceProcessRequestResources) == 1 {
		processRequest = &(*c.IncludedActivityreferenceProcessRequestResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceNutritionOrderResource() (nutritionOrder *NutritionOrder, err error) {
	if c.IncludedActivityreferenceNutritionOrderResources == nil {
		err = errors.New("Included nutritionorders not requested")
	} else if len(*c.IncludedActivityreferenceNutritionOrderResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 nutritionOrder, but found %d", len(*c.IncludedActivityreferenceNutritionOrderResources))
	} else if len(*c.IncludedActivityreferenceNutritionOrderResources) == 1 {
		nutritionOrder = &(*c.IncludedActivityreferenceNutritionOrderResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceVisionPrescriptionResource() (visionPrescription *VisionPrescription, err error) {
	if c.IncludedActivityreferenceVisionPrescriptionResources == nil {
		err = errors.New("Included visionprescriptions not requested")
	} else if len(*c.IncludedActivityreferenceVisionPrescriptionResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 visionPrescription, but found %d", len(*c.IncludedActivityreferenceVisionPrescriptionResources))
	} else if len(*c.IncludedActivityreferenceVisionPrescriptionResources) == 1 {
		visionPrescription = &(*c.IncludedActivityreferenceVisionPrescriptionResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceDiagnosticOrderResource() (diagnosticOrder *DiagnosticOrder, err error) {
	if c.IncludedActivityreferenceDiagnosticOrderResources == nil {
		err = errors.New("Included diagnosticorders not requested")
	} else if len(*c.IncludedActivityreferenceDiagnosticOrderResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 diagnosticOrder, but found %d", len(*c.IncludedActivityreferenceDiagnosticOrderResources))
	} else if len(*c.IncludedActivityreferenceDiagnosticOrderResources) == 1 {
		diagnosticOrder = &(*c.IncludedActivityreferenceDiagnosticOrderResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceProcedureRequestResource() (procedureRequest *ProcedureRequest, err error) {
	if c.IncludedActivityreferenceProcedureRequestResources == nil {
		err = errors.New("Included procedurerequests not requested")
	} else if len(*c.IncludedActivityreferenceProcedureRequestResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 procedureRequest, but found %d", len(*c.IncludedActivityreferenceProcedureRequestResources))
	} else if len(*c.IncludedActivityreferenceProcedureRequestResources) == 1 {
		procedureRequest = &(*c.IncludedActivityreferenceProcedureRequestResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceDeviceUseRequestResource() (deviceUseRequest *DeviceUseRequest, err error) {
	if c.IncludedActivityreferenceDeviceUseRequestResources == nil {
		err = errors.New("Included deviceuserequests not requested")
	} else if len(*c.IncludedActivityreferenceDeviceUseRequestResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceUseRequest, but found %d", len(*c.IncludedActivityreferenceDeviceUseRequestResources))
	} else if len(*c.IncludedActivityreferenceDeviceUseRequestResources) == 1 {
		deviceUseRequest = &(*c.IncludedActivityreferenceDeviceUseRequestResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceMedicationOrderResource() (medicationOrder *MedicationOrder, err error) {
	if c.IncludedActivityreferenceMedicationOrderResources == nil {
		err = errors.New("Included medicationorders not requested")
	} else if len(*c.IncludedActivityreferenceMedicationOrderResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medicationOrder, but found %d", len(*c.IncludedActivityreferenceMedicationOrderResources))
	} else if len(*c.IncludedActivityreferenceMedicationOrderResources) == 1 {
		medicationOrder = &(*c.IncludedActivityreferenceMedicationOrderResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceCommunicationRequestResource() (communicationRequest *CommunicationRequest, err error) {
	if c.IncludedActivityreferenceCommunicationRequestResources == nil {
		err = errors.New("Included communicationrequests not requested")
	} else if len(*c.IncludedActivityreferenceCommunicationRequestResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 communicationRequest, but found %d", len(*c.IncludedActivityreferenceCommunicationRequestResources))
	} else if len(*c.IncludedActivityreferenceCommunicationRequestResources) == 1 {
		communicationRequest = &(*c.IncludedActivityreferenceCommunicationRequestResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedActivityreferenceSupplyRequestResource() (supplyRequest *SupplyRequest, err error) {
	if c.IncludedActivityreferenceSupplyRequestResources == nil {
		err = errors.New("Included supplyrequests not requested")
	} else if len(*c.IncludedActivityreferenceSupplyRequestResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 supplyRequest, but found %d", len(*c.IncludedActivityreferenceSupplyRequestResources))
	} else if len(*c.IncludedActivityreferenceSupplyRequestResources) == 1 {
		supplyRequest = &(*c.IncludedActivityreferenceSupplyRequestResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedPerformerPractitionerResources() (practitioners []Practitioner, err error) {
	if c.IncludedPerformerPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *c.IncludedPerformerPractitionerResources
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedPerformerOrganizationResources() (organizations []Organization, err error) {
	if c.IncludedPerformerOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *c.IncludedPerformerOrganizationResources
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedPerformerPatientResources() (patients []Patient, err error) {
	if c.IncludedPerformerPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPerformerPatientResources
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedPerformerRelatedPersonResources() (relatedPeople []RelatedPerson, err error) {
	if c.IncludedPerformerRelatedPersonResources == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *c.IncludedPerformerRelatedPersonResources
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedGoalResources() (goals []Goal, err error) {
	if c.IncludedGoalResources == nil {
		err = errors.New("Included goals not requested")
	} else {
		goals = *c.IncludedGoalResources
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedSubjectGroupResource() (group *Group, err error) {
	if c.IncludedSubjectGroupResources == nil {
		err = errors.New("Included groups not requested")
	} else if len(*c.IncludedSubjectGroupResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*c.IncludedSubjectGroupResources))
	} else if len(*c.IncludedSubjectGroupResources) == 1 {
		group = &(*c.IncludedSubjectGroupResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedSubjectPatientResource() (patient *Patient, err error) {
	if c.IncludedSubjectPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedSubjectPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedSubjectPatientResources))
	} else if len(*c.IncludedSubjectPatientResources) == 1 {
		patient = &(*c.IncludedSubjectPatientResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedParticipantPractitionerResource() (practitioner *Practitioner, err error) {
	if c.IncludedParticipantPractitionerResources == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedParticipantPractitionerResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedParticipantPractitionerResources))
	} else if len(*c.IncludedParticipantPractitionerResources) == 1 {
		practitioner = &(*c.IncludedParticipantPractitionerResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedParticipantOrganizationResource() (organization *Organization, err error) {
	if c.IncludedParticipantOrganizationResources == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedParticipantOrganizationResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedParticipantOrganizationResources))
	} else if len(*c.IncludedParticipantOrganizationResources) == 1 {
		organization = &(*c.IncludedParticipantOrganizationResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedParticipantPatientResource() (patient *Patient, err error) {
	if c.IncludedParticipantPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedParticipantPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedParticipantPatientResources))
	} else if len(*c.IncludedParticipantPatientResources) == 1 {
		patient = &(*c.IncludedParticipantPatientResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedParticipantRelatedPersonResource() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedParticipantRelatedPersonResources == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedParticipantRelatedPersonResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedParticipantRelatedPersonResources))
	} else if len(*c.IncludedParticipantRelatedPersonResources) == 1 {
		relatedPerson = &(*c.IncludedParticipantRelatedPersonResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedRelatedplanResource() (carePlan *CarePlan, err error) {
	if c.IncludedRelatedplanResources == nil {
		err = errors.New("Included careplans not requested")
	} else if len(*c.IncludedRelatedplanResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 carePlan, but found %d", len(*c.IncludedRelatedplanResources))
	} else if len(*c.IncludedRelatedplanResources) == 1 {
		carePlan = &(*c.IncludedRelatedplanResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedConditionResources() (conditions []Condition, err error) {
	if c.IncludedConditionResources == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *c.IncludedConditionResources
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if c.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResources))
	} else if len(*c.IncludedPatientResources) == 1 {
		patient = &(*c.IncludedPatientResources)[0]
	}
	return
}

func (c *CarePlanPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedActivityreferenceAppointmentResources != nil {
		for _, r := range *c.IncludedActivityreferenceAppointmentResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActivityreferenceOrderResources != nil {
		for _, r := range *c.IncludedActivityreferenceOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActivityreferenceReferralRequestResources != nil {
		for _, r := range *c.IncludedActivityreferenceReferralRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActivityreferenceProcessRequestResources != nil {
		for _, r := range *c.IncludedActivityreferenceProcessRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActivityreferenceNutritionOrderResources != nil {
		for _, r := range *c.IncludedActivityreferenceNutritionOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActivityreferenceVisionPrescriptionResources != nil {
		for _, r := range *c.IncludedActivityreferenceVisionPrescriptionResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActivityreferenceDiagnosticOrderResources != nil {
		for _, r := range *c.IncludedActivityreferenceDiagnosticOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActivityreferenceProcedureRequestResources != nil {
		for _, r := range *c.IncludedActivityreferenceProcedureRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActivityreferenceDeviceUseRequestResources != nil {
		for _, r := range *c.IncludedActivityreferenceDeviceUseRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActivityreferenceMedicationOrderResources != nil {
		for _, r := range *c.IncludedActivityreferenceMedicationOrderResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActivityreferenceCommunicationRequestResources != nil {
		for _, r := range *c.IncludedActivityreferenceCommunicationRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedActivityreferenceSupplyRequestResources != nil {
		for _, r := range *c.IncludedActivityreferenceSupplyRequestResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPerformerPractitionerResources != nil {
		for _, r := range *c.IncludedPerformerPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPerformerOrganizationResources != nil {
		for _, r := range *c.IncludedPerformerOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPerformerPatientResources != nil {
		for _, r := range *c.IncludedPerformerPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPerformerRelatedPersonResources != nil {
		for _, r := range *c.IncludedPerformerRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedGoalResources != nil {
		for _, r := range *c.IncludedGoalResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSubjectGroupResources != nil {
		for _, r := range *c.IncludedSubjectGroupResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSubjectPatientResources != nil {
		for _, r := range *c.IncludedSubjectPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedParticipantPractitionerResources != nil {
		for _, r := range *c.IncludedParticipantPractitionerResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedParticipantOrganizationResources != nil {
		for _, r := range *c.IncludedParticipantOrganizationResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedParticipantPatientResources != nil {
		for _, r := range *c.IncludedParticipantPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedParticipantRelatedPersonResources != nil {
		for _, r := range *c.IncludedParticipantRelatedPersonResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRelatedplanResources != nil {
		for _, r := range *c.IncludedRelatedplanResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedConditionResources != nil {
		for _, r := range *c.IncludedConditionResources {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResources != nil {
		for _, r := range *c.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
