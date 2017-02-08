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

type Group struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type           string                         `bson:"type,omitempty" json:"type,omitempty"`
	Actual         *bool                          `bson:"actual,omitempty" json:"actual,omitempty"`
	Active         *bool                          `bson:"active,omitempty" json:"active,omitempty"`
	Code           *CodeableConcept               `bson:"code,omitempty" json:"code,omitempty"`
	Name           string                         `bson:"name,omitempty" json:"name,omitempty"`
	Quantity       *uint32                        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Characteristic []GroupCharacteristicComponent `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	Member         []GroupMemberComponent         `bson:"member,omitempty" json:"member,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Group) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Group"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Group), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Group) GetBSON() (interface{}, error) {
	x.ResourceType = "Group"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "group" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type group Group

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Group) UnmarshalJSON(data []byte) (err error) {
	x2 := group{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Group(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Group) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Group"
	} else if x.ResourceType != "Group" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Group, instead received %s", x.ResourceType))
	}
	return nil
}

type GroupCharacteristicComponent struct {
	BackboneElement      `bson:",inline"`
	Code                 *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	Exclude              *bool            `bson:"exclude,omitempty" json:"exclude,omitempty"`
	Period               *Period          `bson:"period,omitempty" json:"period,omitempty"`
}

type GroupMemberComponent struct {
	BackboneElement `bson:",inline"`
	Entity          *Reference `bson:"entity,omitempty" json:"entity,omitempty"`
	Period          *Period    `bson:"period,omitempty" json:"period,omitempty"`
	Inactive        *bool      `bson:"inactive,omitempty" json:"inactive,omitempty"`
}

type GroupPlus struct {
	Group                     `bson:",inline"`
	GroupPlusRelatedResources `bson:",inline"`
}

type GroupPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByMember              *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByMember,omitempty"`
	IncludedDeviceResourcesReferencedByMember                    *[]Device                `bson:"_includedDeviceResourcesReferencedByMember,omitempty"`
	IncludedMedicationResourcesReferencedByMember                *[]Medication            `bson:"_includedMedicationResourcesReferencedByMember,omitempty"`
	IncludedPatientResourcesReferencedByMember                   *[]Patient               `bson:"_includedPatientResourcesReferencedByMember,omitempty"`
	IncludedSubstanceResourcesReferencedByMember                 *[]Substance             `bson:"_includedSubstanceResourcesReferencedByMember,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingSubject       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedGoalResourcesReferencingSubject                   *[]Goal                  `bson:"_revIncludedGoalResourcesReferencingSubject,omitempty"`
	RevIncludedConsentResourcesReferencingData                   *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedConsentResourcesReferencingActor                  *[]Consent               `bson:"_revIncludedConsentResourcesReferencingActor,omitempty"`
	RevIncludedConsentResourcesReferencingRecipient              *[]Consent               `bson:"_revIncludedConsentResourcesReferencingRecipient,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingSubject      *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref   *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingAgent                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingAgent,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject               *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingSubject         *[]RiskAssessment        `bson:"_revIncludedRiskAssessmentResourcesReferencingSubject,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest          *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedCareTeamResourcesReferencingSubject               *[]CareTeam              `bson:"_revIncludedCareTeamResourcesReferencingSubject,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource   *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon          *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingSubject          *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient        *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedRequestGroupResourcesReferencingSubject           *[]RequestGroup          `bson:"_revIncludedRequestGroupResourcesReferencingSubject,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData             *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                   *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                     *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                   *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedSpecimenResourcesReferencingSubject               *[]Specimen              `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedCarePlanResourcesReferencingSubject               *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureResourcesReferencingSubject              *[]Procedure             `bson:"_revIncludedProcedureResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingItem                      *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                   *[]List                  `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingSubject      *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon      *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition   *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                  *[]Media                 `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingSubject       *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingSubject       *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon       *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition    *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                   *[]Flag                  `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedGuidanceResponseResourcesReferencingSubject       *[]GuidanceResponse      `bson:"_revIncludedGuidanceResponseResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingSubject            *[]Observation           `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSubject    *[]MedicationStatement   `bson:"_revIncludedMedicationStatementResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient *[]CommunicationRequest  `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                  *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject       *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity              *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingSubject              *[]Condition             `bson:"_revIncludedConditionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject            *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated       *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject  *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest        *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSubject     *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingSubject,omitempty"`
}

func (g *GroupPlusRelatedResources) GetIncludedPractitionerResourceReferencedByMember() (practitioner *Practitioner, err error) {
	if g.IncludedPractitionerResourcesReferencedByMember == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*g.IncludedPractitionerResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*g.IncludedPractitionerResourcesReferencedByMember))
	} else if len(*g.IncludedPractitionerResourcesReferencedByMember) == 1 {
		practitioner = &(*g.IncludedPractitionerResourcesReferencedByMember)[0]
	}
	return
}

func (g *GroupPlusRelatedResources) GetIncludedDeviceResourceReferencedByMember() (device *Device, err error) {
	if g.IncludedDeviceResourcesReferencedByMember == nil {
		err = errors.New("Included devices not requested")
	} else if len(*g.IncludedDeviceResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*g.IncludedDeviceResourcesReferencedByMember))
	} else if len(*g.IncludedDeviceResourcesReferencedByMember) == 1 {
		device = &(*g.IncludedDeviceResourcesReferencedByMember)[0]
	}
	return
}

func (g *GroupPlusRelatedResources) GetIncludedMedicationResourceReferencedByMember() (medication *Medication, err error) {
	if g.IncludedMedicationResourcesReferencedByMember == nil {
		err = errors.New("Included medications not requested")
	} else if len(*g.IncludedMedicationResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*g.IncludedMedicationResourcesReferencedByMember))
	} else if len(*g.IncludedMedicationResourcesReferencedByMember) == 1 {
		medication = &(*g.IncludedMedicationResourcesReferencedByMember)[0]
	}
	return
}

func (g *GroupPlusRelatedResources) GetIncludedPatientResourceReferencedByMember() (patient *Patient, err error) {
	if g.IncludedPatientResourcesReferencedByMember == nil {
		err = errors.New("Included patients not requested")
	} else if len(*g.IncludedPatientResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*g.IncludedPatientResourcesReferencedByMember))
	} else if len(*g.IncludedPatientResourcesReferencedByMember) == 1 {
		patient = &(*g.IncludedPatientResourcesReferencedByMember)[0]
	}
	return
}

func (g *GroupPlusRelatedResources) GetIncludedSubstanceResourceReferencedByMember() (substance *Substance, err error) {
	if g.IncludedSubstanceResourcesReferencedByMember == nil {
		err = errors.New("Included substances not requested")
	} else if len(*g.IncludedSubstanceResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*g.IncludedSubstanceResourcesReferencedByMember))
	} else if len(*g.IncludedSubstanceResourcesReferencedByMember) == 1 {
		substance = &(*g.IncludedSubstanceResourcesReferencedByMember)[0]
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if g.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *g.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingSubject() (documentManifests []DocumentManifest, err error) {
	if g.RevIncludedDocumentManifestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *g.RevIncludedDocumentManifestResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *g.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedGoalResourcesReferencingSubject() (goals []Goal, err error) {
	if g.RevIncludedGoalResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded goals not requested")
	} else {
		goals = *g.RevIncludedGoalResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if g.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *g.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedConsentResourcesReferencingActor() (consents []Consent, err error) {
	if g.RevIncludedConsentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *g.RevIncludedConsentResourcesReferencingActor
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedConsentResourcesReferencingRecipient() (consents []Consent, err error) {
	if g.RevIncludedConsentResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *g.RevIncludedConsentResourcesReferencingRecipient
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingSubject() (documentReferences []DocumentReference, err error) {
	if g.RevIncludedDocumentReferenceResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *g.RevIncludedDocumentReferenceResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedContractResourcesReferencingAgent() (contracts []Contract, err error) {
	if g.RevIncludedContractResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *g.RevIncludedContractResourcesReferencingAgent
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if g.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *g.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if g.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *g.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if g.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *g.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingSubject() (riskAssessments []RiskAssessment, err error) {
	if g.RevIncludedRiskAssessmentResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *g.RevIncludedRiskAssessmentResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if g.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *g.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if g.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *g.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingSubject() (careTeams []CareTeam, err error) {
	if g.RevIncludedCareTeamResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *g.RevIncludedCareTeamResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if g.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *g.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if g.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *g.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSubject() (communications []Communication, err error) {
	if g.RevIncludedCommunicationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *g.RevIncludedCommunicationResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if g.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *g.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingSubject() (requestGroups []RequestGroup, err error) {
	if g.RevIncludedRequestGroupResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *g.RevIncludedRequestGroupResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if g.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *g.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if g.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *g.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if g.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *g.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if g.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *g.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if g.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *g.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if g.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *g.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingSubject() (specimen []Specimen, err error) {
	if g.RevIncludedSpecimenResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *g.RevIncludedSpecimenResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingSubject() (carePlans []CarePlan, err error) {
	if g.RevIncludedCarePlanResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *g.RevIncludedCarePlanResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingSubject() (procedures []Procedure, err error) {
	if g.RevIncludedProcedureResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *g.RevIncludedProcedureResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if g.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *g.RevIncludedListResourcesReferencingItem
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedListResourcesReferencingSubject() (lists []List, err error) {
	if g.RevIncludedListResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *g.RevIncludedListResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if g.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *g.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingSubject() (diagnosticRequests []DiagnosticRequest, err error) {
	if g.RevIncludedDiagnosticRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *g.RevIncludedDiagnosticRequestResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if g.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *g.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if g.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *g.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if g.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *g.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingSubject() (procedureRequests []ProcedureRequest, err error) {
	if g.RevIncludedProcedureRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *g.RevIncludedProcedureRequestResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if g.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *g.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingSubject() (deviceUseRequests []DeviceUseRequest, err error) {
	if g.RevIncludedDeviceUseRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *g.RevIncludedDeviceUseRequestResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if g.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *g.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if g.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *g.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if g.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *g.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedGuidanceResponseResourcesReferencingSubject() (guidanceResponses []GuidanceResponse, err error) {
	if g.RevIncludedGuidanceResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded guidanceResponses not requested")
	} else {
		guidanceResponses = *g.RevIncludedGuidanceResponseResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedObservationResourcesReferencingSubject() (observations []Observation, err error) {
	if g.RevIncludedObservationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *g.RevIncludedObservationResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSubject() (medicationStatements []MedicationStatement, err error) {
	if g.RevIncludedMedicationStatementResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *g.RevIncludedMedicationStatementResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if g.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *g.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if g.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *g.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingSubject() (diagnosticReports []DiagnosticReport, err error) {
	if g.RevIncludedDiagnosticReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *g.RevIncludedDiagnosticReportResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if g.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *g.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedConditionResourcesReferencingSubject() (conditions []Condition, err error) {
	if g.RevIncludedConditionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *g.RevIncludedConditionResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if g.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *g.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if g.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *g.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *g.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if g.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *g.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSubject() (clinicalImpressions []ClinicalImpression, err error) {
	if g.RevIncludedClinicalImpressionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *g.RevIncludedClinicalImpressionResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.IncludedPractitionerResourcesReferencedByMember != nil {
		for idx := range *g.IncludedPractitionerResourcesReferencedByMember {
			rsc := (*g.IncludedPractitionerResourcesReferencedByMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedDeviceResourcesReferencedByMember != nil {
		for idx := range *g.IncludedDeviceResourcesReferencedByMember {
			rsc := (*g.IncludedDeviceResourcesReferencedByMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedMedicationResourcesReferencedByMember != nil {
		for idx := range *g.IncludedMedicationResourcesReferencedByMember {
			rsc := (*g.IncludedMedicationResourcesReferencedByMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedPatientResourcesReferencedByMember != nil {
		for idx := range *g.IncludedPatientResourcesReferencedByMember {
			rsc := (*g.IncludedPatientResourcesReferencedByMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedSubstanceResourcesReferencedByMember != nil {
		for idx := range *g.IncludedSubstanceResourcesReferencedByMember {
			rsc := (*g.IncludedSubstanceResourcesReferencedByMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (g *GroupPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*g.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingData {
			rsc := (*g.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingActor {
			rsc := (*g.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingRecipient != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingRecipient {
			rsc := (*g.RevIncludedConsentResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingAgent {
			rsc := (*g.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*g.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingSubject {
			rsc := (*g.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingTopic {
			rsc := (*g.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedRiskAssessmentResourcesReferencingSubject {
			rsc := (*g.RevIncludedRiskAssessmentResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *g.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*g.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *g.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*g.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCareTeamResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCareTeamResourcesReferencingSubject {
			rsc := (*g.RevIncludedCareTeamResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *g.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*g.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingSubject {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedRequestGroupResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedRequestGroupResourcesReferencingSubject {
			rsc := (*g.RevIncludedRequestGroupResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *g.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*g.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *g.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*g.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *g.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*g.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*g.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*g.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*g.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*g.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCarePlanResourcesReferencingSubject {
			rsc := (*g.RevIncludedCarePlanResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProcedureResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedProcedureResourcesReferencingSubject {
			rsc := (*g.RevIncludedProcedureResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *g.RevIncludedListResourcesReferencingItem {
			rsc := (*g.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedListResourcesReferencingSubject {
			rsc := (*g.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *g.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*g.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDiagnosticRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedDiagnosticRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *g.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*g.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*g.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedProcedureRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedProcedureRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *g.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*g.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceUseRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDeviceUseRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedDeviceUseRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *g.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*g.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*g.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedGuidanceResponseResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedGuidanceResponseResourcesReferencingSubject {
			rsc := (*g.RevIncludedGuidanceResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*g.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMedicationStatementResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMedicationStatementResourcesReferencingSubject {
			rsc := (*g.RevIncludedMedicationStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *g.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*g.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*g.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*g.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *g.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*g.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConditionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedConditionResourcesReferencingSubject {
			rsc := (*g.RevIncludedConditionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*g.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *g.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*g.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *g.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*g.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*g.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *g.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*g.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedClinicalImpressionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedClinicalImpressionResourcesReferencingSubject {
			rsc := (*g.RevIncludedClinicalImpressionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (g *GroupPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.IncludedPractitionerResourcesReferencedByMember != nil {
		for idx := range *g.IncludedPractitionerResourcesReferencedByMember {
			rsc := (*g.IncludedPractitionerResourcesReferencedByMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedDeviceResourcesReferencedByMember != nil {
		for idx := range *g.IncludedDeviceResourcesReferencedByMember {
			rsc := (*g.IncludedDeviceResourcesReferencedByMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedMedicationResourcesReferencedByMember != nil {
		for idx := range *g.IncludedMedicationResourcesReferencedByMember {
			rsc := (*g.IncludedMedicationResourcesReferencedByMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedPatientResourcesReferencedByMember != nil {
		for idx := range *g.IncludedPatientResourcesReferencedByMember {
			rsc := (*g.IncludedPatientResourcesReferencedByMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedSubstanceResourcesReferencedByMember != nil {
		for idx := range *g.IncludedSubstanceResourcesReferencedByMember {
			rsc := (*g.IncludedSubstanceResourcesReferencedByMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*g.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingData {
			rsc := (*g.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingActor {
			rsc := (*g.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingRecipient != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingRecipient {
			rsc := (*g.RevIncludedConsentResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingAgent {
			rsc := (*g.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*g.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingSubject {
			rsc := (*g.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingTopic {
			rsc := (*g.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedRiskAssessmentResourcesReferencingSubject {
			rsc := (*g.RevIncludedRiskAssessmentResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *g.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*g.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *g.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*g.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCareTeamResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCareTeamResourcesReferencingSubject {
			rsc := (*g.RevIncludedCareTeamResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *g.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*g.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingSubject {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedRequestGroupResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedRequestGroupResourcesReferencingSubject {
			rsc := (*g.RevIncludedRequestGroupResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *g.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*g.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *g.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*g.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *g.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*g.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*g.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*g.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*g.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*g.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCarePlanResourcesReferencingSubject {
			rsc := (*g.RevIncludedCarePlanResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProcedureResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedProcedureResourcesReferencingSubject {
			rsc := (*g.RevIncludedProcedureResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *g.RevIncludedListResourcesReferencingItem {
			rsc := (*g.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedListResourcesReferencingSubject {
			rsc := (*g.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *g.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*g.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDiagnosticRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedDiagnosticRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *g.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*g.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*g.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedProcedureRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedProcedureRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *g.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*g.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceUseRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDeviceUseRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedDeviceUseRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *g.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*g.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*g.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedGuidanceResponseResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedGuidanceResponseResourcesReferencingSubject {
			rsc := (*g.RevIncludedGuidanceResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*g.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMedicationStatementResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMedicationStatementResourcesReferencingSubject {
			rsc := (*g.RevIncludedMedicationStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *g.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*g.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*g.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*g.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *g.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*g.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConditionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedConditionResourcesReferencingSubject {
			rsc := (*g.RevIncludedConditionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*g.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *g.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*g.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *g.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*g.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*g.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *g.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*g.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedClinicalImpressionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedClinicalImpressionResourcesReferencingSubject {
			rsc := (*g.RevIncludedClinicalImpressionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
