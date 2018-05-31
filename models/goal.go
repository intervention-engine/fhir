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
	Identifier           []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string               `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept    `bson:"category,omitempty" json:"category,omitempty"`
	Priority             *CodeableConcept     `bson:"priority,omitempty" json:"priority,omitempty"`
	Description          *CodeableConcept     `bson:"description,omitempty" json:"description,omitempty"`
	Subject              *Reference           `bson:"subject,omitempty" json:"subject,omitempty"`
	StartDate            *FHIRDateTime        `bson:"startDate,omitempty" json:"startDate,omitempty"`
	StartCodeableConcept *CodeableConcept     `bson:"startCodeableConcept,omitempty" json:"startCodeableConcept,omitempty"`
	Target               *GoalTargetComponent `bson:"target,omitempty" json:"target,omitempty"`
	StatusDate           *FHIRDateTime        `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	StatusReason         string               `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	ExpressedBy          *Reference           `bson:"expressedBy,omitempty" json:"expressedBy,omitempty"`
	Addresses            []Reference          `bson:"addresses,omitempty" json:"addresses,omitempty"`
	Note                 []Annotation         `bson:"note,omitempty" json:"note,omitempty"`
	OutcomeCode          []CodeableConcept    `bson:"outcomeCode,omitempty" json:"outcomeCode,omitempty"`
	OutcomeReference     []Reference          `bson:"outcomeReference,omitempty" json:"outcomeReference,omitempty"`
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
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
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

type GoalTargetComponent struct {
	BackboneElement       `bson:",inline"`
	Measure               *CodeableConcept `bson:"measure,omitempty" json:"measure,omitempty"`
	DetailQuantity        *Quantity        `bson:"detailQuantity,omitempty" json:"detailQuantity,omitempty"`
	DetailRange           *Range           `bson:"detailRange,omitempty" json:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `bson:"detailCodeableConcept,omitempty" json:"detailCodeableConcept,omitempty"`
	DueDate               *FHIRDateTime    `bson:"dueDate,omitempty" json:"dueDate,omitempty"`
	DueDuration           *Quantity        `bson:"dueDuration,omitempty" json:"dueDuration,omitempty"`
}

type GoalPlus struct {
	Goal                     `bson:",inline"`
	GoalPlusRelatedResources `bson:",inline"`
}

type GoalPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedGroupResourcesReferencedBySubject                       *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedOrganizationResourcesReferencedBySubject                *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                     *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                 *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1            *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2            *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref      *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingSubject                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest             *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse            *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource      *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor       *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom     *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor     *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof      *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson       *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof              *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon             *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor      *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom    *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor    *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof     *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition           *[]RequestGroup          `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon             *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest        *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                 *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                        *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingGoal                     *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingGoal,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces         *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon          *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom               *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor               *[]Library               `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                *[]Library               `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon      *[]CommunicationRequest  `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                     *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                 *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail          *[]Condition             `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject               *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                 *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated          *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject     *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest           *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor          *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof         *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
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

func (g *GoalPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if g.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *g.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if g.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *g.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingDependsonPath2
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

func (g *GoalPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if g.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *g.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if g.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *g.RevIncludedContractResourcesReferencingTermtopic
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

func (g *GoalPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if g.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *g.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if g.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *g.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if g.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *g.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if g.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *g.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if g.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *g.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if g.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *g.RevIncludedCommunicationResourcesReferencingPartof
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

func (g *GoalPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if g.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *g.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if g.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *g.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if g.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *g.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if g.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *g.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if g.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *g.RevIncludedProvenanceResourcesReferencingEntityref
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

func (g *GoalPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if g.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *g.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if g.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *g.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if g.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *g.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if g.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *g.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if g.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *g.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if g.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *g.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if g.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *g.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if g.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *g.RevIncludedCommunicationRequestResourcesReferencingBasedon
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

func (g *GoalPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if g.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *g.RevIncludedConditionResourcesReferencingEvidencedetail
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

func (g *GoalPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (g *GoalPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
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
	if g.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*g.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*g.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*g.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingSubject {
			rsc := (*g.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*g.RevIncludedContractResourcesReferencingTermtopic)[idx]
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
	if g.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*g.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *g.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*g.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *g.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*g.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *g.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*g.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *g.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*g.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if g.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *g.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*g.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*g.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*g.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if g.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *g.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*g.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if g.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
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
	if g.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*g.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*g.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*g.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingSubject {
			rsc := (*g.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*g.RevIncludedContractResourcesReferencingTermtopic)[idx]
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
	if g.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*g.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *g.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*g.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *g.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*g.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *g.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*g.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *g.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*g.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if g.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *g.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*g.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*g.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*g.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if g.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *g.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*g.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if g.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
