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

type Goal struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                 `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept      `bson:"category,omitempty" json:"category,omitempty"`
	Priority             *CodeableConcept       `bson:"priority,omitempty" json:"priority,omitempty"`
	Description          *CodeableConcept       `bson:"description,omitempty" json:"description,omitempty"`
	Subject              *Reference             `bson:"subject,omitempty" json:"subject,omitempty"`
	StartDate            *FHIRDateTime          `bson:"startDate,omitempty" json:"startDate,omitempty"`
	StartCodeableConcept *CodeableConcept       `bson:"startCodeableConcept,omitempty" json:"startCodeableConcept,omitempty"`
	TargetDate           *FHIRDateTime          `bson:"targetDate,omitempty" json:"targetDate,omitempty"`
	TargetDuration       *Quantity              `bson:"targetDuration,omitempty" json:"targetDuration,omitempty"`
	StatusDate           *FHIRDateTime          `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	StatusReason         []CodeableConcept      `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	ExpressedBy          *Reference             `bson:"expressedBy,omitempty" json:"expressedBy,omitempty"`
	Addresses            []Reference            `bson:"addresses,omitempty" json:"addresses,omitempty"`
	Note                 []Annotation           `bson:"note,omitempty" json:"note,omitempty"`
	Outcome              []GoalOutcomeComponent `bson:"outcome,omitempty" json:"outcome,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Goal) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Goal"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Goal), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Goal) GetBSON() (interface{}, error) {
	x.ResourceType = "Goal"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "goal" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type goal Goal

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Goal) UnmarshalJSON(data []byte) (err error) {
	x2 := goal{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Goal(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Goal) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Goal"
	} else if x.ResourceType != "Goal" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Goal, instead received %s", x.ResourceType))
	}
	return nil
}

type GoalOutcomeComponent struct {
	BackboneElement       `bson:",inline"`
	ResultCodeableConcept *CodeableConcept `bson:"resultCodeableConcept,omitempty" json:"resultCodeableConcept,omitempty"`
	ResultReference       *Reference       `bson:"resultReference,omitempty" json:"resultReference,omitempty"`
}

type GoalPlus struct {
	Goal                     `bson:",inline"`
	GoalPlusRelatedResources `bson:",inline"`
}

type GoalPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedGroupResourcesReferencedBySubject                   *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedOrganizationResourcesReferencedBySubject            *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                  *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic               *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject              *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse        *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource  *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon         *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                    *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingGoal                 *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingGoal,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces    *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition  *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces     *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition   *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity             *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
}

func (g *GoalPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if g.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*g.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*g.IncludedPatientResourcesReferencedByPatient))
	} else if len(*g.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*g.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (g *GoalPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if g.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*g.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*g.IncludedGroupResourcesReferencedBySubject))
	} else if len(*g.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*g.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (g *GoalPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySubject() (organization *Organization, err error) {
	if g.IncludedOrganizationResourcesReferencedBySubject == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*g.IncludedOrganizationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*g.IncludedOrganizationResourcesReferencedBySubject))
	} else if len(*g.IncludedOrganizationResourcesReferencedBySubject) == 1 {
		organization = &(*g.IncludedOrganizationResourcesReferencedBySubject)[0]
	}
	return
}

func (g *GoalPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if g.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*g.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*g.IncludedPatientResourcesReferencedBySubject))
	} else if len(*g.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*g.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if g.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *g.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *g.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if g.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *g.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if g.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *g.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if g.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *g.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if g.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *g.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if g.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *g.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if g.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *g.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if g.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *g.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if g.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *g.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if g.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *g.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if g.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *g.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if g.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *g.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if g.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *g.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if g.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *g.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if g.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *g.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingGoal() (carePlans []CarePlan, err error) {
	if g.RevIncludedCarePlanResourcesReferencingGoal == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *g.RevIncludedCarePlanResourcesReferencingGoal
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if g.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *g.RevIncludedListResourcesReferencingItem
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if g.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *g.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if g.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *g.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if g.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *g.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if g.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *g.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if g.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *g.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if g.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *g.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if g.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *g.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if g.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *g.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if g.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *g.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if g.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *g.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *g.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if g.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *g.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (g *GoalPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *g.IncludedPatientResourcesReferencedByPatient {
			rsc := (*g.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *g.IncludedGroupResourcesReferencedBySubject {
			rsc := (*g.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedOrganizationResourcesReferencedBySubject != nil {
		for idx := range *g.IncludedOrganizationResourcesReferencedBySubject {
			rsc := (*g.IncludedOrganizationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *g.IncludedPatientResourcesReferencedBySubject {
			rsc := (*g.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (g *GoalPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingData {
			rsc := (*g.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
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
	if g.RevIncludedCarePlanResourcesReferencingGoal != nil {
		for idx := range *g.RevIncludedCarePlanResourcesReferencingGoal {
			rsc := (*g.RevIncludedCarePlanResourcesReferencingGoal)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *g.RevIncludedListResourcesReferencingItem {
			rsc := (*g.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *g.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*g.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
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
	if g.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *g.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*g.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
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
	if g.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*g.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *g.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*g.RevIncludedAuditEventResourcesReferencingEntity)[idx]
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
	return resourceMap
}

func (g *GoalPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *g.IncludedPatientResourcesReferencedByPatient {
			rsc := (*g.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *g.IncludedGroupResourcesReferencedBySubject {
			rsc := (*g.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedOrganizationResourcesReferencedBySubject != nil {
		for idx := range *g.IncludedOrganizationResourcesReferencedBySubject {
			rsc := (*g.IncludedOrganizationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *g.IncludedPatientResourcesReferencedBySubject {
			rsc := (*g.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingData {
			rsc := (*g.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
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
	if g.RevIncludedCarePlanResourcesReferencingGoal != nil {
		for idx := range *g.RevIncludedCarePlanResourcesReferencingGoal {
			rsc := (*g.RevIncludedCarePlanResourcesReferencingGoal)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *g.RevIncludedListResourcesReferencingItem {
			rsc := (*g.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *g.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*g.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
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
	if g.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *g.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*g.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
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
	if g.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*g.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *g.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*g.RevIncludedAuditEventResourcesReferencingEntity)[idx]
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
	return resourceMap
}
