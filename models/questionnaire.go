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

type Questionnaire struct {
	DomainResource `bson:",inline"`
	Url            string                       `bson:"url,omitempty" json:"url,omitempty"`
	Identifier     []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version        string                       `bson:"version,omitempty" json:"version,omitempty"`
	Status         string                       `bson:"status,omitempty" json:"status,omitempty"`
	Date           *FHIRDateTime                `bson:"date,omitempty" json:"date,omitempty"`
	Publisher      string                       `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Telecom        []ContactPoint               `bson:"telecom,omitempty" json:"telecom,omitempty"`
	UseContext     []CodeableConcept            `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Title          string                       `bson:"title,omitempty" json:"title,omitempty"`
	Concept        []Coding                     `bson:"concept,omitempty" json:"concept,omitempty"`
	SubjectType    []string                     `bson:"subjectType,omitempty" json:"subjectType,omitempty"`
	Item           []QuestionnaireItemComponent `bson:"item,omitempty" json:"item,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Questionnaire) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Questionnaire"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Questionnaire), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Questionnaire) GetBSON() (interface{}, error) {
	x.ResourceType = "Questionnaire"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "questionnaire" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type questionnaire Questionnaire

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Questionnaire) UnmarshalJSON(data []byte) (err error) {
	x2 := questionnaire{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Questionnaire(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Questionnaire) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Questionnaire"
	} else if x.ResourceType != "Questionnaire" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Questionnaire, instead received %s", x.ResourceType))
	}
	return nil
}

type QuestionnaireItemComponent struct {
	BackboneElement   `bson:",inline"`
	LinkId            string                                 `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Definition        string                                 `bson:"definition,omitempty" json:"definition,omitempty"`
	Concept           []Coding                               `bson:"concept,omitempty" json:"concept,omitempty"`
	Prefix            string                                 `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Text              string                                 `bson:"text,omitempty" json:"text,omitempty"`
	Type              string                                 `bson:"type,omitempty" json:"type,omitempty"`
	EnableWhen        []QuestionnaireItemEnableWhenComponent `bson:"enableWhen,omitempty" json:"enableWhen,omitempty"`
	Required          *bool                                  `bson:"required,omitempty" json:"required,omitempty"`
	Repeats           *bool                                  `bson:"repeats,omitempty" json:"repeats,omitempty"`
	ReadOnly          *bool                                  `bson:"readOnly,omitempty" json:"readOnly,omitempty"`
	MaxLength         *int32                                 `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	Options           *Reference                             `bson:"options,omitempty" json:"options,omitempty"`
	Option            []QuestionnaireItemOptionComponent     `bson:"option,omitempty" json:"option,omitempty"`
	InitialBoolean    *bool                                  `bson:"initialBoolean,omitempty" json:"initialBoolean,omitempty"`
	InitialDecimal    *float64                               `bson:"initialDecimal,omitempty" json:"initialDecimal,omitempty"`
	InitialInteger    *int32                                 `bson:"initialInteger,omitempty" json:"initialInteger,omitempty"`
	InitialDate       *FHIRDateTime                          `bson:"initialDate,omitempty" json:"initialDate,omitempty"`
	InitialDateTime   *FHIRDateTime                          `bson:"initialDateTime,omitempty" json:"initialDateTime,omitempty"`
	InitialInstant    *FHIRDateTime                          `bson:"initialInstant,omitempty" json:"initialInstant,omitempty"`
	InitialTime       *FHIRDateTime                          `bson:"initialTime,omitempty" json:"initialTime,omitempty"`
	InitialString     string                                 `bson:"initialString,omitempty" json:"initialString,omitempty"`
	InitialUri        string                                 `bson:"initialUri,omitempty" json:"initialUri,omitempty"`
	InitialAttachment *Attachment                            `bson:"initialAttachment,omitempty" json:"initialAttachment,omitempty"`
	InitialCoding     *Coding                                `bson:"initialCoding,omitempty" json:"initialCoding,omitempty"`
	InitialQuantity   *Quantity                              `bson:"initialQuantity,omitempty" json:"initialQuantity,omitempty"`
	InitialReference  *Reference                             `bson:"initialReference,omitempty" json:"initialReference,omitempty"`
	Item              []QuestionnaireItemComponent           `bson:"item,omitempty" json:"item,omitempty"`
}

type QuestionnaireItemEnableWhenComponent struct {
	BackboneElement  `bson:",inline"`
	Question         string        `bson:"question,omitempty" json:"question,omitempty"`
	HasAnswer        *bool         `bson:"hasAnswer,omitempty" json:"hasAnswer,omitempty"`
	AnswerBoolean    *bool         `bson:"answerBoolean,omitempty" json:"answerBoolean,omitempty"`
	AnswerDecimal    *float64      `bson:"answerDecimal,omitempty" json:"answerDecimal,omitempty"`
	AnswerInteger    *int32        `bson:"answerInteger,omitempty" json:"answerInteger,omitempty"`
	AnswerDate       *FHIRDateTime `bson:"answerDate,omitempty" json:"answerDate,omitempty"`
	AnswerDateTime   *FHIRDateTime `bson:"answerDateTime,omitempty" json:"answerDateTime,omitempty"`
	AnswerInstant    *FHIRDateTime `bson:"answerInstant,omitempty" json:"answerInstant,omitempty"`
	AnswerTime       *FHIRDateTime `bson:"answerTime,omitempty" json:"answerTime,omitempty"`
	AnswerString     string        `bson:"answerString,omitempty" json:"answerString,omitempty"`
	AnswerUri        string        `bson:"answerUri,omitempty" json:"answerUri,omitempty"`
	AnswerAttachment *Attachment   `bson:"answerAttachment,omitempty" json:"answerAttachment,omitempty"`
	AnswerCoding     *Coding       `bson:"answerCoding,omitempty" json:"answerCoding,omitempty"`
	AnswerQuantity   *Quantity     `bson:"answerQuantity,omitempty" json:"answerQuantity,omitempty"`
	AnswerReference  *Reference    `bson:"answerReference,omitempty" json:"answerReference,omitempty"`
}

type QuestionnaireItemOptionComponent struct {
	BackboneElement `bson:",inline"`
	ValueInteger    *int32        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate       *FHIRDateTime `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueTime       *FHIRDateTime `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString     string        `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueCoding     *Coding       `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
}

type QuestionnairePlus struct {
	Questionnaire                     `bson:",inline"`
	QuestionnairePlusRelatedResources `bson:",inline"`
}

type QuestionnairePlusRelatedResources struct {
	RevIncludedDocumentManifestResourcesReferencingContentref         *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref         *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                        *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref        *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                     *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                    *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                      *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest               *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse              *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource        *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon               *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                  *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                   *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                   *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                        *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                          *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                        *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                           *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces          *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon           *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition        *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces           *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon            *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition         *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                       *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                   *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                 *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                   *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated            *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingQuestionnaire,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject       *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest             *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if q.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *q.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if q.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *q.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if q.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *q.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if q.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *q.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if q.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *q.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if q.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *q.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if q.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *q.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if q.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *q.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if q.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *q.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if q.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *q.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if q.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *q.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if q.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *q.RevIncludedListResourcesReferencingItem
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if q.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *q.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if q.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *q.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if q.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *q.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if q.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *q.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if q.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *q.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if q.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *q.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if q.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *q.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if q.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *q.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *q.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire() (questionnaireResponses []QuestionnaireResponse, err error) {
	if q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if q.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *q.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingData {
			rsc := (*q.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*q.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingSubject {
			rsc := (*q.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingTopic {
			rsc := (*q.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *q.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*q.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*q.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*q.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*q.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *q.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*q.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*q.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *q.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*q.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*q.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*q.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*q.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*q.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*q.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *q.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*q.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*q.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (q *QuestionnairePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingData {
			rsc := (*q.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*q.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingSubject {
			rsc := (*q.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingTopic {
			rsc := (*q.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *q.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*q.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*q.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*q.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*q.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *q.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*q.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*q.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *q.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*q.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*q.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*q.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*q.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*q.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*q.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *q.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*q.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*q.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
