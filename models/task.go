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

type Task struct {
	DomainResource      `bson:",inline"`
	Identifier          []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	DefinitionUri       string                    `bson:"definitionUri,omitempty" json:"definitionUri,omitempty"`
	DefinitionReference *Reference                `bson:"definitionReference,omitempty" json:"definitionReference,omitempty"`
	BasedOn             []Reference               `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	GroupIdentifier     *Identifier               `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	PartOf              []Reference               `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status              string                    `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason        *CodeableConcept          `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	BusinessStatus      *CodeableConcept          `bson:"businessStatus,omitempty" json:"businessStatus,omitempty"`
	Intent              string                    `bson:"intent,omitempty" json:"intent,omitempty"`
	Priority            string                    `bson:"priority,omitempty" json:"priority,omitempty"`
	Code                *CodeableConcept          `bson:"code,omitempty" json:"code,omitempty"`
	Description         string                    `bson:"description,omitempty" json:"description,omitempty"`
	Focus               *Reference                `bson:"focus,omitempty" json:"focus,omitempty"`
	For                 *Reference                `bson:"for,omitempty" json:"for,omitempty"`
	Context             *Reference                `bson:"context,omitempty" json:"context,omitempty"`
	ExecutionPeriod     *Period                   `bson:"executionPeriod,omitempty" json:"executionPeriod,omitempty"`
	AuthoredOn          *FHIRDateTime             `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	LastModified        *FHIRDateTime             `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	Requester           *TaskRequesterComponent   `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType       []CodeableConcept         `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Owner               *Reference                `bson:"owner,omitempty" json:"owner,omitempty"`
	Reason              *CodeableConcept          `bson:"reason,omitempty" json:"reason,omitempty"`
	Note                []Annotation              `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory     []Reference               `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
	Restriction         *TaskRestrictionComponent `bson:"restriction,omitempty" json:"restriction,omitempty"`
	Input               []TaskParameterComponent  `bson:"input,omitempty" json:"input,omitempty"`
	Output              []TaskOutputComponent     `bson:"output,omitempty" json:"output,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Task) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Task"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Task), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Task) GetBSON() (interface{}, error) {
	x.ResourceType = "Task"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "task" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type task Task

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Task) UnmarshalJSON(data []byte) (err error) {
	x2 := task{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = Task(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Task) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Task"
	} else if x.ResourceType != "Task" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Task, instead received %s", x.ResourceType))
	}
	return nil
}

type TaskRequesterComponent struct {
	BackboneElement `bson:",inline"`
	Agent           *Reference `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf      *Reference `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}

type TaskRestrictionComponent struct {
	BackboneElement `bson:",inline"`
	Repetitions     *uint32     `bson:"repetitions,omitempty" json:"repetitions,omitempty"`
	Period          *Period     `bson:"period,omitempty" json:"period,omitempty"`
	Recipient       []Reference `bson:"recipient,omitempty" json:"recipient,omitempty"`
}

type TaskParameterComponent struct {
	BackboneElement      `bson:",inline"`
	Type                 *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ValueAddress         *Address         `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAnnotation      *Annotation      `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueBase64Binary    string           `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCode            string           `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCoding          *Coding          `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint    *ContactPoint    `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueDate            *FHIRDateTime    `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime        *FHIRDateTime    `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal         *float64         `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueHumanName       *HumanName       `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueId              string           `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueIdentifier      *Identifier      `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueInstant         *FHIRDateTime    `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger         *int32           `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown        string           `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueMeta            *Meta            `bson:"valueMeta,omitempty" json:"valueMeta,omitempty"`
	ValueOid             string           `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePeriod          *Period          `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValuePositiveInt     *uint32          `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio           `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData     *SampledData     `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature       *Signature       `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueString          string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime            *FHIRDateTime    `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueTiming          *Timing          `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueUnsignedInt     *uint32          `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri             string           `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
}

type TaskOutputComponent struct {
	BackboneElement      `bson:",inline"`
	Type                 *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ValueAddress         *Address         `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAnnotation      *Annotation      `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueBase64Binary    string           `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCode            string           `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCoding          *Coding          `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint    *ContactPoint    `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueDate            *FHIRDateTime    `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime        *FHIRDateTime    `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal         *float64         `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueHumanName       *HumanName       `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueId              string           `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueIdentifier      *Identifier      `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueInstant         *FHIRDateTime    `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger         *int32           `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown        string           `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueMeta            *Meta            `bson:"valueMeta,omitempty" json:"valueMeta,omitempty"`
	ValueOid             string           `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePeriod          *Period          `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValuePositiveInt     *uint32          `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio           `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData     *SampledData     `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature       *Signature       `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueString          string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime            *FHIRDateTime    `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueTiming          *Timing          `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueUnsignedInt     *uint32          `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri             string           `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
}

type TaskPlus struct {
	Task                     `bson:",inline"`
	TaskPlusRelatedResources `bson:",inline"`
}

type TaskPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByOwner                  *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByOwner,omitempty"`
	IncludedOrganizationResourcesReferencedByOwner                  *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOwner,omitempty"`
	IncludedDeviceResourcesReferencedByOwner                        *[]Device                `bson:"_includedDeviceResourcesReferencedByOwner,omitempty"`
	IncludedPatientResourcesReferencedByOwner                       *[]Patient               `bson:"_includedPatientResourcesReferencedByOwner,omitempty"`
	IncludedRelatedPersonResourcesReferencedByOwner                 *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByOwner,omitempty"`
	IncludedPractitionerResourcesReferencedByRequester              *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRequester,omitempty"`
	IncludedOrganizationResourcesReferencedByRequester              *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRequester,omitempty"`
	IncludedDeviceResourcesReferencedByRequester                    *[]Device                `bson:"_includedDeviceResourcesReferencedByRequester,omitempty"`
	IncludedPatientResourcesReferencedByRequester                   *[]Patient               `bson:"_includedPatientResourcesReferencedByRequester,omitempty"`
	IncludedRelatedPersonResourcesReferencedByRequester             *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByRequester,omitempty"`
	IncludedTaskResourcesReferencedByPartof                         *[]Task                  `bson:"_includedTaskResourcesReferencedByPartof,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization           *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext               *[]EpisodeOfCare         `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext                   *[]Encounter             `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
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
	RevIncludedTaskResourcesReferencingPartof                       *[]Task                  `bson:"_revIncludedTaskResourcesReferencingPartof,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference        *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
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

func (t *TaskPlusRelatedResources) GetIncludedPractitionerResourceReferencedByOwner() (practitioner *Practitioner, err error) {
	if t.IncludedPractitionerResourcesReferencedByOwner == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*t.IncludedPractitionerResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*t.IncludedPractitionerResourcesReferencedByOwner))
	} else if len(*t.IncludedPractitionerResourcesReferencedByOwner) == 1 {
		practitioner = &(*t.IncludedPractitionerResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOwner() (organization *Organization, err error) {
	if t.IncludedOrganizationResourcesReferencedByOwner == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*t.IncludedOrganizationResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*t.IncludedOrganizationResourcesReferencedByOwner))
	} else if len(*t.IncludedOrganizationResourcesReferencedByOwner) == 1 {
		organization = &(*t.IncludedOrganizationResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedDeviceResourceReferencedByOwner() (device *Device, err error) {
	if t.IncludedDeviceResourcesReferencedByOwner == nil {
		err = errors.New("Included devices not requested")
	} else if len(*t.IncludedDeviceResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*t.IncludedDeviceResourcesReferencedByOwner))
	} else if len(*t.IncludedDeviceResourcesReferencedByOwner) == 1 {
		device = &(*t.IncludedDeviceResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPatientResourceReferencedByOwner() (patient *Patient, err error) {
	if t.IncludedPatientResourcesReferencedByOwner == nil {
		err = errors.New("Included patients not requested")
	} else if len(*t.IncludedPatientResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*t.IncludedPatientResourcesReferencedByOwner))
	} else if len(*t.IncludedPatientResourcesReferencedByOwner) == 1 {
		patient = &(*t.IncludedPatientResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByOwner() (relatedPerson *RelatedPerson, err error) {
	if t.IncludedRelatedPersonResourcesReferencedByOwner == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*t.IncludedRelatedPersonResourcesReferencedByOwner))
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByOwner) == 1 {
		relatedPerson = &(*t.IncludedRelatedPersonResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPractitionerResourceReferencedByRequester() (practitioner *Practitioner, err error) {
	if t.IncludedPractitionerResourcesReferencedByRequester == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*t.IncludedPractitionerResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*t.IncludedPractitionerResourcesReferencedByRequester))
	} else if len(*t.IncludedPractitionerResourcesReferencedByRequester) == 1 {
		practitioner = &(*t.IncludedPractitionerResourcesReferencedByRequester)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedOrganizationResourceReferencedByRequester() (organization *Organization, err error) {
	if t.IncludedOrganizationResourcesReferencedByRequester == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*t.IncludedOrganizationResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*t.IncludedOrganizationResourcesReferencedByRequester))
	} else if len(*t.IncludedOrganizationResourcesReferencedByRequester) == 1 {
		organization = &(*t.IncludedOrganizationResourcesReferencedByRequester)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedDeviceResourceReferencedByRequester() (device *Device, err error) {
	if t.IncludedDeviceResourcesReferencedByRequester == nil {
		err = errors.New("Included devices not requested")
	} else if len(*t.IncludedDeviceResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*t.IncludedDeviceResourcesReferencedByRequester))
	} else if len(*t.IncludedDeviceResourcesReferencedByRequester) == 1 {
		device = &(*t.IncludedDeviceResourcesReferencedByRequester)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPatientResourceReferencedByRequester() (patient *Patient, err error) {
	if t.IncludedPatientResourcesReferencedByRequester == nil {
		err = errors.New("Included patients not requested")
	} else if len(*t.IncludedPatientResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*t.IncludedPatientResourcesReferencedByRequester))
	} else if len(*t.IncludedPatientResourcesReferencedByRequester) == 1 {
		patient = &(*t.IncludedPatientResourcesReferencedByRequester)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByRequester() (relatedPerson *RelatedPerson, err error) {
	if t.IncludedRelatedPersonResourcesReferencedByRequester == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*t.IncludedRelatedPersonResourcesReferencedByRequester))
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByRequester) == 1 {
		relatedPerson = &(*t.IncludedRelatedPersonResourcesReferencedByRequester)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedTaskResourcesReferencedByPartof() (tasks []Task, err error) {
	if t.IncludedTaskResourcesReferencedByPartof == nil {
		err = errors.New("Included tasks not requested")
	} else {
		tasks = *t.IncludedTaskResourcesReferencedByPartof
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if t.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*t.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*t.IncludedPatientResourcesReferencedByPatient))
	} else if len(*t.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*t.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if t.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*t.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*t.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*t.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*t.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByContext() (episodeOfCare *EpisodeOfCare, err error) {
	if t.IncludedEpisodeOfCareResourcesReferencedByContext == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*t.IncludedEpisodeOfCareResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*t.IncludedEpisodeOfCareResourcesReferencedByContext))
	} else if len(*t.IncludedEpisodeOfCareResourcesReferencedByContext) == 1 {
		episodeOfCare = &(*t.IncludedEpisodeOfCareResourcesReferencedByContext)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedEncounterResourceReferencedByContext() (encounter *Encounter, err error) {
	if t.IncludedEncounterResourcesReferencedByContext == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*t.IncludedEncounterResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*t.IncludedEncounterResourcesReferencedByContext))
	} else if len(*t.IncludedEncounterResourcesReferencedByContext) == 1 {
		encounter = &(*t.IncludedEncounterResourcesReferencedByContext)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if t.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *t.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if t.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *t.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if t.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *t.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if t.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *t.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if t.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *t.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if t.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *t.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if t.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *t.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if t.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *t.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if t.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *t.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if t.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *t.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if t.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *t.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if t.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *t.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if t.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *t.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *t.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if t.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *t.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedTaskResourcesReferencingPartof() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingPartof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if t.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *t.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if t.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *t.RevIncludedListResourcesReferencingItem
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if t.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *t.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if t.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *t.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *t.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if t.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *t.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if t.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *t.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if t.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *t.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *t.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if t.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *t.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.IncludedPractitionerResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByOwner {
			rsc := (*t.IncludedPractitionerResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByOwner {
			rsc := (*t.IncludedOrganizationResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByOwner {
			rsc := (*t.IncludedDeviceResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByOwner {
			rsc := (*t.IncludedPatientResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByOwner {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*t.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*t.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*t.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByRequester {
			rsc := (*t.IncludedPatientResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByRequester {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedTaskResourcesReferencedByPartof != nil {
		for idx := range *t.IncludedTaskResourcesReferencedByPartof {
			rsc := (*t.IncludedTaskResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByPatient {
			rsc := (*t.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*t.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *t.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*t.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *t.IncludedEncounterResourcesReferencedByContext {
			rsc := (*t.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (t *TaskPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*t.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*t.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*t.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*t.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*t.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *t.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*t.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*t.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingPartof {
			rsc := (*t.RevIncludedTaskResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*t.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *t.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*t.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *t.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*t.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*t.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *t.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*t.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*t.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (t *TaskPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.IncludedPractitionerResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByOwner {
			rsc := (*t.IncludedPractitionerResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByOwner {
			rsc := (*t.IncludedOrganizationResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByOwner {
			rsc := (*t.IncludedDeviceResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByOwner {
			rsc := (*t.IncludedPatientResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByOwner {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*t.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*t.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*t.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByRequester {
			rsc := (*t.IncludedPatientResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByRequester {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedTaskResourcesReferencedByPartof != nil {
		for idx := range *t.IncludedTaskResourcesReferencedByPartof {
			rsc := (*t.IncludedTaskResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByPatient {
			rsc := (*t.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*t.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *t.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*t.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *t.IncludedEncounterResourcesReferencedByContext {
			rsc := (*t.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*t.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*t.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*t.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*t.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*t.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *t.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*t.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*t.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingPartof {
			rsc := (*t.RevIncludedTaskResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*t.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *t.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*t.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *t.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*t.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*t.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *t.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*t.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*t.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
