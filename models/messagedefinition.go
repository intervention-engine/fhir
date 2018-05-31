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

type MessageDefinition struct {
	DomainResource   `bson:",inline"`
	Url              string                                      `bson:"url,omitempty" json:"url,omitempty"`
	Identifier       *Identifier                                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version          string                                      `bson:"version,omitempty" json:"version,omitempty"`
	Name             string                                      `bson:"name,omitempty" json:"name,omitempty"`
	Title            string                                      `bson:"title,omitempty" json:"title,omitempty"`
	Status           string                                      `bson:"status,omitempty" json:"status,omitempty"`
	Experimental     *bool                                       `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date             *FHIRDateTime                               `bson:"date,omitempty" json:"date,omitempty"`
	Publisher        string                                      `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact          []ContactDetail                             `bson:"contact,omitempty" json:"contact,omitempty"`
	Description      string                                      `bson:"description,omitempty" json:"description,omitempty"`
	UseContext       []UsageContext                              `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction     []CodeableConcept                           `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose          string                                      `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright        string                                      `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Base             *Reference                                  `bson:"base,omitempty" json:"base,omitempty"`
	Parent           []Reference                                 `bson:"parent,omitempty" json:"parent,omitempty"`
	Replaces         []Reference                                 `bson:"replaces,omitempty" json:"replaces,omitempty"`
	Event            *Coding                                     `bson:"event,omitempty" json:"event,omitempty"`
	Category         string                                      `bson:"category,omitempty" json:"category,omitempty"`
	Focus            []MessageDefinitionFocusComponent           `bson:"focus,omitempty" json:"focus,omitempty"`
	ResponseRequired *bool                                       `bson:"responseRequired,omitempty" json:"responseRequired,omitempty"`
	AllowedResponse  []MessageDefinitionAllowedResponseComponent `bson:"allowedResponse,omitempty" json:"allowedResponse,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *MessageDefinition) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "MessageDefinition"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to MessageDefinition), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *MessageDefinition) GetBSON() (interface{}, error) {
	x.ResourceType = "MessageDefinition"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "messageDefinition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type messageDefinition MessageDefinition

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *MessageDefinition) UnmarshalJSON(data []byte) (err error) {
	x2 := messageDefinition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = MessageDefinition(x2)
		return x.checkResourceType()
	}
	return
}

func (x *MessageDefinition) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "MessageDefinition"
	} else if x.ResourceType != "MessageDefinition" {
		return errors.New(fmt.Sprintf("Expected resourceType to be MessageDefinition, instead received %s", x.ResourceType))
	}
	return nil
}

type MessageDefinitionFocusComponent struct {
	BackboneElement `bson:",inline"`
	Code            string     `bson:"code,omitempty" json:"code,omitempty"`
	Profile         *Reference `bson:"profile,omitempty" json:"profile,omitempty"`
	Min             *uint32    `bson:"min,omitempty" json:"min,omitempty"`
	Max             string     `bson:"max,omitempty" json:"max,omitempty"`
}

type MessageDefinitionAllowedResponseComponent struct {
	BackboneElement `bson:",inline"`
	Message         *Reference `bson:"message,omitempty" json:"message,omitempty"`
	Situation       string     `bson:"situation,omitempty" json:"situation,omitempty"`
}

type MessageDefinitionPlus struct {
	MessageDefinition                     `bson:",inline"`
	MessageDefinitionPlusRelatedResources `bson:",inline"`
}

type MessageDefinitionPlusRelatedResources struct {
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

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if m.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *m.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if m.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *m.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if m.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *m.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if m.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *m.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if m.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *m.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if m.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *m.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if m.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *m.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if m.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *m.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if m.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *m.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if m.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *m.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if m.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *m.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *m.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if m.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *m.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if m.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *m.RevIncludedListResourcesReferencingItem
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if m.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *m.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if m.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *m.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *m.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if m.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *m.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if m.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *m.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if m.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *m.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *m.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if m.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *m.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (m *MessageDefinitionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (m *MessageDefinitionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*m.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*m.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*m.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*m.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *m.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*m.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *m.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*m.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*m.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *m.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*m.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MessageDefinitionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*m.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*m.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*m.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*m.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *m.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*m.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *m.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*m.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*m.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *m.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*m.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
