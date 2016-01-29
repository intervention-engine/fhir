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
	Code                 *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	Exclude              *bool            `bson:"exclude,omitempty" json:"exclude,omitempty"`
	Period               *Period          `bson:"period,omitempty" json:"period,omitempty"`
}

type GroupMemberComponent struct {
	Entity   *Reference `bson:"entity,omitempty" json:"entity,omitempty"`
	Period   *Period    `bson:"period,omitempty" json:"period,omitempty"`
	Inactive *bool      `bson:"inactive,omitempty" json:"inactive,omitempty"`
}

type GroupPlus struct {
	Group                     `bson:",inline"`
	GroupPlusRelatedResources `bson:",inline"`
}

type GroupPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByMember             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByMember,omitempty"`
	IncludedDeviceResourcesReferencedByMember                   *[]Device                `bson:"_includedDeviceResourcesReferencedByMember,omitempty"`
	IncludedMedicationResourcesReferencedByMember               *[]Medication            `bson:"_includedMedicationResourcesReferencedByMember,omitempty"`
	IncludedPatientResourcesReferencedByMember                  *[]Patient               `bson:"_includedPatientResourcesReferencedByMember,omitempty"`
	IncludedSubstanceResourcesReferencedByMember                *[]Substance             `bson:"_includedSubstanceResourcesReferencedByMember,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingSubject      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedSpecimenResourcesReferencingSubject              *[]Specimen              `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedCarePlanResourcesReferencingSubject              *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingSubject,omitempty"`
	RevIncludedGoalResourcesReferencingSubject                  *[]Goal                  `bson:"_revIncludedGoalResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureResourcesReferencingSubject             *[]Procedure             `bson:"_revIncludedProcedureResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                  *[]List                  `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingSubject     *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingSubject                 *[]Order                 `bson:"_revIncludedOrderResourcesReferencingSubject,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                 *[]Media                 `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingSubject      *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingSubject,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                  *[]Flag                  `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingSubject           *[]Observation           `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingActor                *[]Contract              `bson:"_revIncludedContractResourcesReferencingActor,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingSubject        *[]RiskAssessment        `bson:"_revIncludedRiskAssessmentResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject      *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient       *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingSubject       *[]DiagnosticOrder       `bson:"_revIncludedDiagnosticOrderResourcesReferencingSubject,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
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

func (g *GroupPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if g.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *g.RevIncludedProvenanceResourcesReferencingTarget
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

func (g *GroupPlusRelatedResources) GetRevIncludedGoalResourcesReferencingSubject() (goals []Goal, err error) {
	if g.RevIncludedGoalResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded goals not requested")
	} else {
		goals = *g.RevIncludedGoalResourcesReferencingSubject
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

func (g *GroupPlusRelatedResources) GetRevIncludedOrderResourcesReferencingSubject() (orders []Order, err error) {
	if g.RevIncludedOrderResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *g.RevIncludedOrderResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if g.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *g.RevIncludedOrderResourcesReferencingDetail
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

func (g *GroupPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if g.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *g.RevIncludedFlagResourcesReferencingSubject
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

func (g *GroupPlusRelatedResources) GetRevIncludedContractResourcesReferencingActor() (contracts []Contract, err error) {
	if g.RevIncludedContractResourcesReferencingActor == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *g.RevIncludedContractResourcesReferencingActor
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

func (g *GroupPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if g.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *g.RevIncludedAuditEventResourcesReferencingReference
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

func (g *GroupPlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingSubject() (diagnosticOrders []DiagnosticOrder, err error) {
	if g.RevIncludedDiagnosticOrderResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *g.RevIncludedDiagnosticOrderResourcesReferencingSubject
	}
	return
}

func (g *GroupPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if g.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *g.RevIncludedOrderResponseResourcesReferencingFulfillment
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

func (g *GroupPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if g.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *g.RevIncludedClinicalImpressionResourcesReferencingTrigger
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

func (g *GroupPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.IncludedPractitionerResourcesReferencedByMember != nil {
		for _, r := range *g.IncludedPractitionerResourcesReferencedByMember {
			resourceMap[r.Id] = &r
		}
	}
	if g.IncludedDeviceResourcesReferencedByMember != nil {
		for _, r := range *g.IncludedDeviceResourcesReferencedByMember {
			resourceMap[r.Id] = &r
		}
	}
	if g.IncludedMedicationResourcesReferencedByMember != nil {
		for _, r := range *g.IncludedMedicationResourcesReferencedByMember {
			resourceMap[r.Id] = &r
		}
	}
	if g.IncludedPatientResourcesReferencedByMember != nil {
		for _, r := range *g.IncludedPatientResourcesReferencedByMember {
			resourceMap[r.Id] = &r
		}
	}
	if g.IncludedSubstanceResourcesReferencedByMember != nil {
		for _, r := range *g.IncludedSubstanceResourcesReferencedByMember {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (g *GroupPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *g.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *g.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedDocumentManifestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *g.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedSpecimenResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedCarePlanResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedGoalResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedGoalResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedProcedureResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedProcedureResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *g.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedListResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedListResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedDocumentReferenceResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedOrderResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *g.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedMediaResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedMediaResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedProcedureRequestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedFlagResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedFlagResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedObservationResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedObservationResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedContractResourcesReferencingActor != nil {
		for _, r := range *g.RevIncludedContractResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedRiskAssessmentResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedDiagnosticReportResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *g.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for _, r := range *g.RevIncludedCommunicationResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *g.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *g.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *g.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *g.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *g.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *g.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (g *GroupPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.IncludedPractitionerResourcesReferencedByMember != nil {
		for _, r := range *g.IncludedPractitionerResourcesReferencedByMember {
			resourceMap[r.Id] = &r
		}
	}
	if g.IncludedDeviceResourcesReferencedByMember != nil {
		for _, r := range *g.IncludedDeviceResourcesReferencedByMember {
			resourceMap[r.Id] = &r
		}
	}
	if g.IncludedMedicationResourcesReferencedByMember != nil {
		for _, r := range *g.IncludedMedicationResourcesReferencedByMember {
			resourceMap[r.Id] = &r
		}
	}
	if g.IncludedPatientResourcesReferencedByMember != nil {
		for _, r := range *g.IncludedPatientResourcesReferencedByMember {
			resourceMap[r.Id] = &r
		}
	}
	if g.IncludedSubstanceResourcesReferencedByMember != nil {
		for _, r := range *g.IncludedSubstanceResourcesReferencedByMember {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *g.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *g.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedDocumentManifestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *g.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedSpecimenResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedCarePlanResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedGoalResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedGoalResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedProcedureResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedProcedureResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *g.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedListResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedListResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedDocumentReferenceResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedOrderResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *g.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedMediaResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedMediaResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedProcedureRequestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedFlagResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedFlagResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedObservationResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedObservationResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedContractResourcesReferencingActor != nil {
		for _, r := range *g.RevIncludedContractResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedRiskAssessmentResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedDiagnosticReportResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *g.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for _, r := range *g.RevIncludedCommunicationResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *g.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *g.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *g.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *g.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *g.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *g.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
