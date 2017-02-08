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

type PlanDefinition struct {
	DomainResource   `bson:",inline"`
	Url              string                                    `bson:"url,omitempty" json:"url,omitempty"`
	Identifier       []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version          string                                    `bson:"version,omitempty" json:"version,omitempty"`
	Name             string                                    `bson:"name,omitempty" json:"name,omitempty"`
	Title            string                                    `bson:"title,omitempty" json:"title,omitempty"`
	Type             *CodeableConcept                          `bson:"type,omitempty" json:"type,omitempty"`
	Status           string                                    `bson:"status,omitempty" json:"status,omitempty"`
	Experimental     *bool                                     `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date             *FHIRDateTime                             `bson:"date,omitempty" json:"date,omitempty"`
	Description      string                                    `bson:"description,omitempty" json:"description,omitempty"`
	Purpose          string                                    `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage            string                                    `bson:"usage,omitempty" json:"usage,omitempty"`
	ApprovalDate     *FHIRDateTime                             `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate   *FHIRDateTime                             `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod  *Period                                   `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext       []UsageContext                            `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction     []CodeableConcept                         `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Topic            []CodeableConcept                         `bson:"topic,omitempty" json:"topic,omitempty"`
	Contributor      []Contributor                             `bson:"contributor,omitempty" json:"contributor,omitempty"`
	Publisher        string                                    `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact          []ContactDetail                           `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright        string                                    `bson:"copyright,omitempty" json:"copyright,omitempty"`
	RelatedArtifact  []RelatedArtifact                         `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Library          []Reference                               `bson:"library,omitempty" json:"library,omitempty"`
	ActionDefinition []PlanDefinitionActionDefinitionComponent `bson:"actionDefinition,omitempty" json:"actionDefinition,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *PlanDefinition) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "PlanDefinition"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to PlanDefinition), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *PlanDefinition) GetBSON() (interface{}, error) {
	x.ResourceType = "PlanDefinition"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "planDefinition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type planDefinition PlanDefinition

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *PlanDefinition) UnmarshalJSON(data []byte) (err error) {
	x2 := planDefinition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = PlanDefinition(x2)
		return x.checkResourceType()
	}
	return
}

func (x *PlanDefinition) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "PlanDefinition"
	} else if x.ResourceType != "PlanDefinition" {
		return errors.New(fmt.Sprintf("Expected resourceType to be PlanDefinition, instead received %s", x.ResourceType))
	}
	return nil
}

type PlanDefinitionActionDefinitionComponent struct {
	BackboneElement     `bson:",inline"`
	ActionIdentifier    *Identifier                                            `bson:"actionIdentifier,omitempty" json:"actionIdentifier,omitempty"`
	Label               string                                                 `bson:"label,omitempty" json:"label,omitempty"`
	Title               string                                                 `bson:"title,omitempty" json:"title,omitempty"`
	Description         string                                                 `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent      string                                                 `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	Code                []CodeableConcept                                      `bson:"code,omitempty" json:"code,omitempty"`
	Documentation       []RelatedArtifact                                      `bson:"documentation,omitempty" json:"documentation,omitempty"`
	TriggerDefinition   []TriggerDefinition                                    `bson:"triggerDefinition,omitempty" json:"triggerDefinition,omitempty"`
	Condition           []PlanDefinitionActionDefinitionConditionComponent     `bson:"condition,omitempty" json:"condition,omitempty"`
	Input               []DataRequirement                                      `bson:"input,omitempty" json:"input,omitempty"`
	Output              []DataRequirement                                      `bson:"output,omitempty" json:"output,omitempty"`
	RelatedAction       []PlanDefinitionActionDefinitionRelatedActionComponent `bson:"relatedAction,omitempty" json:"relatedAction,omitempty"`
	TimingDateTime      *FHIRDateTime                                          `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	TimingPeriod        *Period                                                `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDuration      *Quantity                                              `bson:"timingDuration,omitempty" json:"timingDuration,omitempty"`
	TimingRange         *Range                                                 `bson:"timingRange,omitempty" json:"timingRange,omitempty"`
	TimingTiming        *Timing                                                `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	ParticipantType     []string                                               `bson:"participantType,omitempty" json:"participantType,omitempty"`
	Type                *Coding                                                `bson:"type,omitempty" json:"type,omitempty"`
	GroupingBehavior    string                                                 `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty"`
	SelectionBehavior   string                                                 `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty"`
	RequiredBehavior    string                                                 `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty"`
	PrecheckBehavior    string                                                 `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty"`
	CardinalityBehavior string                                                 `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty"`
	ActivityDefinition  *Reference                                             `bson:"activityDefinition,omitempty" json:"activityDefinition,omitempty"`
	Transform           *Reference                                             `bson:"transform,omitempty" json:"transform,omitempty"`
	DynamicValue        []PlanDefinitionActionDefinitionDynamicValueComponent  `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty"`
	ActionDefinition    []PlanDefinitionActionDefinitionComponent              `bson:"actionDefinition,omitempty" json:"actionDefinition,omitempty"`
}

type PlanDefinitionActionDefinitionConditionComponent struct {
	BackboneElement `bson:",inline"`
	Kind            string `bson:"kind,omitempty" json:"kind,omitempty"`
	Description     string `bson:"description,omitempty" json:"description,omitempty"`
	Language        string `bson:"language,omitempty" json:"language,omitempty"`
	Expression      string `bson:"expression,omitempty" json:"expression,omitempty"`
}

type PlanDefinitionActionDefinitionRelatedActionComponent struct {
	BackboneElement  `bson:",inline"`
	ActionIdentifier *Identifier `bson:"actionIdentifier,omitempty" json:"actionIdentifier,omitempty"`
	Relationship     string      `bson:"relationship,omitempty" json:"relationship,omitempty"`
	OffsetDuration   *Quantity   `bson:"offsetDuration,omitempty" json:"offsetDuration,omitempty"`
	OffsetRange      *Range      `bson:"offsetRange,omitempty" json:"offsetRange,omitempty"`
}

type PlanDefinitionActionDefinitionDynamicValueComponent struct {
	BackboneElement `bson:",inline"`
	Description     string `bson:"description,omitempty" json:"description,omitempty"`
	Path            string `bson:"path,omitempty" json:"path,omitempty"`
	Language        string `bson:"language,omitempty" json:"language,omitempty"`
	Expression      string `bson:"expression,omitempty" json:"expression,omitempty"`
}

type PlanDefinitionPlus struct {
	PlanDefinition                     `bson:",inline"`
	PlanDefinitionPlusRelatedResources `bson:",inline"`
}

type PlanDefinitionPlusRelatedResources struct {
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
	RevIncludedResearchStudyResourcesReferencingProtocol        *[]ResearchStudy         `bson:"_revIncludedResearchStudyResourcesReferencingProtocol,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces    *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition  *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces     *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition   *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                  *[]Flag                  `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity             *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if p.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *p.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedResearchStudyResourcesReferencingProtocol() (researchStudies []ResearchStudy, err error) {
	if p.RevIncludedResearchStudyResourcesReferencingProtocol == nil {
		err = errors.New("RevIncluded researchStudies not requested")
	} else {
		researchStudies = *p.RevIncludedResearchStudyResourcesReferencingProtocol
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingItem
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if p.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *p.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if p.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *p.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if p.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *p.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if p.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *p.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if p.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *p.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if p.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *p.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if p.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *p.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (p *PlanDefinitionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (p *PlanDefinitionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingData {
			rsc := (*p.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchStudyResourcesReferencingProtocol != nil {
		for idx := range *p.RevIncludedResearchStudyResourcesReferencingProtocol {
			rsc := (*p.RevIncludedResearchStudyResourcesReferencingProtocol)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*p.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *PlanDefinitionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingData {
			rsc := (*p.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchStudyResourcesReferencingProtocol != nil {
		for idx := range *p.RevIncludedResearchStudyResourcesReferencingProtocol {
			rsc := (*p.RevIncludedResearchStudyResourcesReferencingProtocol)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*p.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
