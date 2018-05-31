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

type Questionnaire struct {
	DomainResource  `bson:",inline"`
	Url             string                       `bson:"url,omitempty" json:"url,omitempty"`
	Identifier      []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version         string                       `bson:"version,omitempty" json:"version,omitempty"`
	Name            string                       `bson:"name,omitempty" json:"name,omitempty"`
	Title           string                       `bson:"title,omitempty" json:"title,omitempty"`
	Status          string                       `bson:"status,omitempty" json:"status,omitempty"`
	Experimental    *bool                        `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date            *FHIRDateTime                `bson:"date,omitempty" json:"date,omitempty"`
	Publisher       string                       `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Description     string                       `bson:"description,omitempty" json:"description,omitempty"`
	Purpose         string                       `bson:"purpose,omitempty" json:"purpose,omitempty"`
	ApprovalDate    *FHIRDateTime                `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate  *FHIRDateTime                `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod *Period                      `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext      []UsageContext               `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction    []CodeableConcept            `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Contact         []ContactDetail              `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright       string                       `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Code            []Coding                     `bson:"code,omitempty" json:"code,omitempty"`
	SubjectType     []string                     `bson:"subjectType,omitempty" json:"subjectType,omitempty"`
	Item            []QuestionnaireItemComponent `bson:"item,omitempty" json:"item,omitempty"`
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
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
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
	Code              []Coding                               `bson:"code,omitempty" json:"code,omitempty"`
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
	RevIncludedConsentResourcesReferencingDataPath1                   *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                   *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                   *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                 *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                 *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                  *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1              *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2              *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref        *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingSubject                    *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest               *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse              *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource        *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedFamilyMemberHistoryResourcesReferencingDefinition      *[]FamilyMemberHistory   `bson:"_revIncludedFamilyMemberHistoryResourcesReferencingDefinition,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor         *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom       *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor       *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof        *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson         *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon               *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor        *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom      *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor      *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof       *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1   *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2   *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition             *[]RequestGroup          `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon               *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest          *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                 *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref                *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                   *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                        *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                          *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                        *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingDefinition                 *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingDefinition,omitempty"`
	RevIncludedListResourcesReferencingItem                           *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces           *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon            *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                   *[]Library               `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                  *[]Library               `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                   *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon        *[]CommunicationRequest  `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                       *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                   *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail            *[]Condition             `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                 *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                   *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated            *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingQuestionnaire,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject       *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest             *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor            *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom          *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor          *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof           *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1       *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2       *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
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

func (q *QuestionnairePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if q.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *q.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if q.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *q.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDependsonPath2
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

func (q *QuestionnairePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if q.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *q.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if q.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *q.RevIncludedContractResourcesReferencingTermtopic
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

func (q *QuestionnairePlusRelatedResources) GetRevIncludedFamilyMemberHistoryResourcesReferencingDefinition() (familyMemberHistories []FamilyMemberHistory, err error) {
	if q.RevIncludedFamilyMemberHistoryResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded familyMemberHistories not requested")
	} else {
		familyMemberHistories = *q.RevIncludedFamilyMemberHistoryResourcesReferencingDefinition
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if q.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *q.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if q.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *q.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if q.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *q.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if q.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *q.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if q.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *q.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if q.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *q.RevIncludedCommunicationResourcesReferencingPartof
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

func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if q.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *q.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if q.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *q.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if q.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *q.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if q.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *q.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if q.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *q.RevIncludedProvenanceResourcesReferencingEntityref
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

func (q *QuestionnairePlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingDefinition() (carePlans []CarePlan, err error) {
	if q.RevIncludedCarePlanResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *q.RevIncludedCarePlanResourcesReferencingDefinition
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

func (q *QuestionnairePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if q.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *q.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if q.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *q.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if q.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *q.RevIncludedCommunicationRequestResourcesReferencingBasedon
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

func (q *QuestionnairePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if q.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *q.RevIncludedConditionResourcesReferencingEvidencedetail
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

func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
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
	if q.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*q.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*q.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*q.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingSubject {
			rsc := (*q.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*q.RevIncludedContractResourcesReferencingTermtopic)[idx]
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
	if q.RevIncludedFamilyMemberHistoryResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedFamilyMemberHistoryResourcesReferencingDefinition {
			rsc := (*q.RevIncludedFamilyMemberHistoryResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*q.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*q.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if q.RevIncludedCarePlanResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedCarePlanResourcesReferencingDefinition {
			rsc := (*q.RevIncludedCarePlanResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *q.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*q.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*q.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if q.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *q.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*q.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if q.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
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
	if q.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*q.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*q.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*q.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingSubject {
			rsc := (*q.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*q.RevIncludedContractResourcesReferencingTermtopic)[idx]
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
	if q.RevIncludedFamilyMemberHistoryResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedFamilyMemberHistoryResourcesReferencingDefinition {
			rsc := (*q.RevIncludedFamilyMemberHistoryResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*q.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*q.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if q.RevIncludedCarePlanResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedCarePlanResourcesReferencingDefinition {
			rsc := (*q.RevIncludedCarePlanResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *q.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*q.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*q.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if q.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *q.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*q.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if q.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
