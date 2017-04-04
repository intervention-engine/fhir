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

type CarePlan struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition     []Reference                 `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn        []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces       []Reference                 `bson:"replaces,omitempty" json:"replaces,omitempty"`
	PartOf         []Reference                 `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status         string                      `bson:"status,omitempty" json:"status,omitempty"`
	Intent         string                      `bson:"intent,omitempty" json:"intent,omitempty"`
	Category       []CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Title          string                      `bson:"title,omitempty" json:"title,omitempty"`
	Description    string                      `bson:"description,omitempty" json:"description,omitempty"`
	Subject        *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Context        *Reference                  `bson:"context,omitempty" json:"context,omitempty"`
	Period         *Period                     `bson:"period,omitempty" json:"period,omitempty"`
	Author         []Reference                 `bson:"author,omitempty" json:"author,omitempty"`
	CareTeam       []Reference                 `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	Addresses      []Reference                 `bson:"addresses,omitempty" json:"addresses,omitempty"`
	SupportingInfo []Reference                 `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Goal           []Reference                 `bson:"goal,omitempty" json:"goal,omitempty"`
	Activity       []CarePlanActivityComponent `bson:"activity,omitempty" json:"activity,omitempty"`
	Note           []Annotation                `bson:"note,omitempty" json:"note,omitempty"`
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

type CarePlanActivityComponent struct {
	BackboneElement        `bson:",inline"`
	OutcomeCodeableConcept []CodeableConcept                `bson:"outcomeCodeableConcept,omitempty" json:"outcomeCodeableConcept,omitempty"`
	OutcomeReference       []Reference                      `bson:"outcomeReference,omitempty" json:"outcomeReference,omitempty"`
	Progress               []Annotation                     `bson:"progress,omitempty" json:"progress,omitempty"`
	Reference              *Reference                       `bson:"reference,omitempty" json:"reference,omitempty"`
	Detail                 *CarePlanActivityDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
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
	StatusReason           string            `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
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
	IncludedCareTeamResourcesReferencedByCareteam                      *[]CareTeam              `bson:"_includedCareTeamResourcesReferencedByCareteam,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer                 *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer                 *[]Organization          `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedCareTeamResourcesReferencedByPerformer                     *[]CareTeam              `bson:"_includedCareTeamResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                      *[]Patient               `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedGoalResourcesReferencedByGoal                              *[]Goal                  `bson:"_includedGoalResourcesReferencedByGoal,omitempty"`
	IncludedGroupResourcesReferencedBySubject                          *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                        *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedCarePlanResourcesReferencedByReplaces                      *[]CarePlan              `bson:"_includedCarePlanResourcesReferencedByReplaces,omitempty"`
	IncludedCarePlanResourcesReferencedByPartof                        *[]CarePlan              `bson:"_includedCarePlanResourcesReferencedByPartof,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                    *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedAppointmentResourcesReferencedByActivityreference          *[]Appointment           `bson:"_includedAppointmentResourcesReferencedByActivityreference,omitempty"`
	IncludedReferralRequestResourcesReferencedByActivityreference      *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedMedicationRequestResourcesReferencedByActivityreference    *[]MedicationRequest     `bson:"_includedMedicationRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedTaskResourcesReferencedByActivityreference                 *[]Task                  `bson:"_includedTaskResourcesReferencedByActivityreference,omitempty"`
	IncludedNutritionOrderResourcesReferencedByActivityreference       *[]NutritionOrder        `bson:"_includedNutritionOrderResourcesReferencedByActivityreference,omitempty"`
	IncludedRequestGroupResourcesReferencedByActivityreference         *[]RequestGroup          `bson:"_includedRequestGroupResourcesReferencedByActivityreference,omitempty"`
	IncludedVisionPrescriptionResourcesReferencedByActivityreference   *[]VisionPrescription    `bson:"_includedVisionPrescriptionResourcesReferencedByActivityreference,omitempty"`
	IncludedProcedureRequestResourcesReferencedByActivityreference     *[]ProcedureRequest      `bson:"_includedProcedureRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedDeviceRequestResourcesReferencedByActivityreference        *[]DeviceRequest         `bson:"_includedDeviceRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedCommunicationRequestResourcesReferencedByActivityreference *[]CommunicationRequest  `bson:"_includedCommunicationRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedConditionResourcesReferencedByCondition                    *[]Condition             `bson:"_includedConditionResourcesReferencedByCondition,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                       *[]CarePlan              `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedPatientResourcesReferencedByPatient                        *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext                  *[]EpisodeOfCare         `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext                      *[]Encounter             `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
	IncludedQuestionnaireResourcesReferencedByDefinition               *[]Questionnaire         `bson:"_includedQuestionnaireResourcesReferencedByDefinition,omitempty"`
	IncludedPlanDefinitionResourcesReferencedByDefinition              *[]PlanDefinition        `bson:"_includedPlanDefinitionResourcesReferencedByDefinition,omitempty"`
	RevIncludedReferralRequestResourcesReferencingBasedon              *[]ReferralRequest       `bson:"_revIncludedReferralRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref          *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref          *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                    *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                    *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                    *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                  *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                  *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                   *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref         *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingSubject                     *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                   *[]Contract              `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse               *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource         *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingBasedon                 *[]ImagingStudy          `bson:"_revIncludedImagingStudyResourcesReferencingBasedon,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor          *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom        *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor        *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof         *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson          *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                 *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor         *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom       *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor       *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof        *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1    *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2    *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition              *[]RequestGroup          `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest           *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                  *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref                 *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                    *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                         *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                           *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                         *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingReplaces                    *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingReplaces,omitempty"`
	RevIncludedCarePlanResourcesReferencingPartof                      *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingPartof,omitempty"`
	RevIncludedCarePlanResourcesReferencingBasedon                     *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingBasedon,omitempty"`
	RevIncludedProcedureResourcesReferencingBasedon                    *[]Procedure             `bson:"_revIncludedProcedureResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                            *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces            *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon             *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedObservationResourcesReferencingBasedon                  *[]Observation           `bson:"_revIncludedObservationResourcesReferencingBasedon,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                    *[]Library               `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                  *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                  *[]Library               `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                   *[]Library               `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                    *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon         *[]CommunicationRequest  `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                        *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingBasedon             *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingBasedon,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                    *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail             *[]Condition             `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                  *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                    *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated             *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingBasedon        *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingBasedon,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject        *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest              *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor             *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom           *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor           *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof            *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (c *CarePlanPlusRelatedResources) GetIncludedCareTeamResourcesReferencedByCareteam() (careTeams []CareTeam, err error) {
	if c.IncludedCareTeamResourcesReferencedByCareteam == nil {
		err = errors.New("Included careTeams not requested")
	} else {
		careTeams = *c.IncludedCareTeamResourcesReferencedByCareteam
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

func (c *CarePlanPlusRelatedResources) GetIncludedCareTeamResourcesReferencedByPerformer() (careTeams []CareTeam, err error) {
	if c.IncludedCareTeamResourcesReferencedByPerformer == nil {
		err = errors.New("Included careTeams not requested")
	} else {
		careTeams = *c.IncludedCareTeamResourcesReferencedByPerformer
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

func (c *CarePlanPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByReplaces() (carePlans []CarePlan, err error) {
	if c.IncludedCarePlanResourcesReferencedByReplaces == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *c.IncludedCarePlanResourcesReferencedByReplaces
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByPartof() (carePlans []CarePlan, err error) {
	if c.IncludedCarePlanResourcesReferencedByPartof == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *c.IncludedCarePlanResourcesReferencedByPartof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if c.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*c.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*c.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*c.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*c.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
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

func (c *CarePlanPlusRelatedResources) GetIncludedMedicationRequestResourceReferencedByActivityreference() (medicationRequest *MedicationRequest, err error) {
	if c.IncludedMedicationRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included medicationrequests not requested")
	} else if len(*c.IncludedMedicationRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medicationRequest, but found %d", len(*c.IncludedMedicationRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedMedicationRequestResourcesReferencedByActivityreference) == 1 {
		medicationRequest = &(*c.IncludedMedicationRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedTaskResourceReferencedByActivityreference() (task *Task, err error) {
	if c.IncludedTaskResourcesReferencedByActivityreference == nil {
		err = errors.New("Included tasks not requested")
	} else if len(*c.IncludedTaskResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 task, but found %d", len(*c.IncludedTaskResourcesReferencedByActivityreference))
	} else if len(*c.IncludedTaskResourcesReferencedByActivityreference) == 1 {
		task = &(*c.IncludedTaskResourcesReferencedByActivityreference)[0]
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

func (c *CarePlanPlusRelatedResources) GetIncludedRequestGroupResourceReferencedByActivityreference() (requestGroup *RequestGroup, err error) {
	if c.IncludedRequestGroupResourcesReferencedByActivityreference == nil {
		err = errors.New("Included requestgroups not requested")
	} else if len(*c.IncludedRequestGroupResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 requestGroup, but found %d", len(*c.IncludedRequestGroupResourcesReferencedByActivityreference))
	} else if len(*c.IncludedRequestGroupResourcesReferencedByActivityreference) == 1 {
		requestGroup = &(*c.IncludedRequestGroupResourcesReferencedByActivityreference)[0]
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

func (c *CarePlanPlusRelatedResources) GetIncludedDeviceRequestResourceReferencedByActivityreference() (deviceRequest *DeviceRequest, err error) {
	if c.IncludedDeviceRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included devicerequests not requested")
	} else if len(*c.IncludedDeviceRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceRequest, but found %d", len(*c.IncludedDeviceRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedDeviceRequestResourcesReferencedByActivityreference) == 1 {
		deviceRequest = &(*c.IncludedDeviceRequestResourcesReferencedByActivityreference)[0]
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

func (c *CarePlanPlusRelatedResources) GetIncludedConditionResourcesReferencedByCondition() (conditions []Condition, err error) {
	if c.IncludedConditionResourcesReferencedByCondition == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *c.IncludedConditionResourcesReferencedByCondition
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if c.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *c.IncludedCarePlanResourcesReferencedByBasedon
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

func (c *CarePlanPlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByContext() (episodeOfCare *EpisodeOfCare, err error) {
	if c.IncludedEpisodeOfCareResourcesReferencedByContext == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*c.IncludedEpisodeOfCareResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*c.IncludedEpisodeOfCareResourcesReferencedByContext))
	} else if len(*c.IncludedEpisodeOfCareResourcesReferencedByContext) == 1 {
		episodeOfCare = &(*c.IncludedEpisodeOfCareResourcesReferencedByContext)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedEncounterResourceReferencedByContext() (encounter *Encounter, err error) {
	if c.IncludedEncounterResourcesReferencedByContext == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*c.IncludedEncounterResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*c.IncludedEncounterResourcesReferencedByContext))
	} else if len(*c.IncludedEncounterResourcesReferencedByContext) == 1 {
		encounter = &(*c.IncludedEncounterResourcesReferencedByContext)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedQuestionnaireResourcesReferencedByDefinition() (questionnaires []Questionnaire, err error) {
	if c.IncludedQuestionnaireResourcesReferencedByDefinition == nil {
		err = errors.New("Included questionnaires not requested")
	} else {
		questionnaires = *c.IncludedQuestionnaireResourcesReferencedByDefinition
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPlanDefinitionResourcesReferencedByDefinition() (planDefinitions []PlanDefinition, err error) {
	if c.IncludedPlanDefinitionResourcesReferencedByDefinition == nil {
		err = errors.New("Included planDefinitions not requested")
	} else {
		planDefinitions = *c.IncludedPlanDefinitionResourcesReferencedByDefinition
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath2
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponse
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingPartof
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if c.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *c.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingEntityref
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingSubject
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingReplaces() (carePlans []CarePlan, err error) {
	if c.RevIncludedCarePlanResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *c.RevIncludedCarePlanResourcesReferencingReplaces
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPartof() (carePlans []CarePlan, err error) {
	if c.RevIncludedCarePlanResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *c.RevIncludedCarePlanResourcesReferencingPartof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingBasedon() (carePlans []CarePlan, err error) {
	if c.RevIncludedCarePlanResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *c.RevIncludedCarePlanResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingBasedon() (procedures []Procedure, err error) {
	if c.RevIncludedProcedureResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *c.RevIncludedProcedureResourcesReferencingBasedon
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if c.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *c.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if c.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *c.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedObservationResourcesReferencingBasedon() (observations []Observation, err error) {
	if c.RevIncludedObservationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *c.RevIncludedObservationResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *c.RevIncludedCommunicationRequestResourcesReferencingBasedon
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingBasedon() (diagnosticReports []DiagnosticReport, err error) {
	if c.RevIncludedDiagnosticReportResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *c.RevIncludedDiagnosticReportResourcesReferencingBasedon
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if c.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *c.RevIncludedConditionResourcesReferencingEvidencedetail
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

func (c *CarePlanPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedCareTeamResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByCareteam {
			rsc := (*c.IncludedCareTeamResourcesReferencedByCareteam)[idx]
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
	if c.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*c.IncludedCareTeamResourcesReferencedByPerformer)[idx]
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
	if c.IncludedCarePlanResourcesReferencedByReplaces != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByReplaces {
			rsc := (*c.IncludedCarePlanResourcesReferencedByReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByPartof != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByPartof {
			rsc := (*c.IncludedCarePlanResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*c.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
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
	if c.IncludedMedicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedMedicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedMedicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedTaskResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedTaskResourcesReferencedByActivityreference {
			rsc := (*c.IncludedTaskResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRequestGroupResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedRequestGroupResourcesReferencedByActivityreference {
			rsc := (*c.IncludedRequestGroupResourcesReferencedByActivityreference)[idx]
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
	if c.IncludedDeviceRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDeviceRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDeviceRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByCondition != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByCondition {
			rsc := (*c.IncludedConditionResourcesReferencedByCondition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*c.IncludedCarePlanResourcesReferencedByBasedon)[idx]
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
	if c.IncludedQuestionnaireResourcesReferencedByDefinition != nil {
		for idx := range *c.IncludedQuestionnaireResourcesReferencedByDefinition {
			rsc := (*c.IncludedQuestionnaireResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPlanDefinitionResourcesReferencedByDefinition != nil {
		for idx := range *c.IncludedPlanDefinitionResourcesReferencedByDefinition {
			rsc := (*c.IncludedPlanDefinitionResourcesReferencedByDefinition)[idx]
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
	if c.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
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
	if c.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*c.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingReplaces {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingPartof {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *CarePlanPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedCareTeamResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByCareteam {
			rsc := (*c.IncludedCareTeamResourcesReferencedByCareteam)[idx]
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
	if c.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*c.IncludedCareTeamResourcesReferencedByPerformer)[idx]
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
	if c.IncludedCarePlanResourcesReferencedByReplaces != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByReplaces {
			rsc := (*c.IncludedCarePlanResourcesReferencedByReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByPartof != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByPartof {
			rsc := (*c.IncludedCarePlanResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*c.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
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
	if c.IncludedMedicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedMedicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedMedicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedTaskResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedTaskResourcesReferencedByActivityreference {
			rsc := (*c.IncludedTaskResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRequestGroupResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedRequestGroupResourcesReferencedByActivityreference {
			rsc := (*c.IncludedRequestGroupResourcesReferencedByActivityreference)[idx]
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
	if c.IncludedDeviceRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDeviceRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDeviceRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByCondition != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByCondition {
			rsc := (*c.IncludedConditionResourcesReferencedByCondition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*c.IncludedCarePlanResourcesReferencedByBasedon)[idx]
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
	if c.IncludedQuestionnaireResourcesReferencedByDefinition != nil {
		for idx := range *c.IncludedQuestionnaireResourcesReferencedByDefinition {
			rsc := (*c.IncludedQuestionnaireResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPlanDefinitionResourcesReferencedByDefinition != nil {
		for idx := range *c.IncludedPlanDefinitionResourcesReferencedByDefinition {
			rsc := (*c.IncludedPlanDefinitionResourcesReferencedByDefinition)[idx]
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
	if c.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
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
	if c.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*c.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingReplaces {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingPartof {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
