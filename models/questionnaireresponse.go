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

type QuestionnaireResponse struct {
	DomainResource `bson:",inline"`
	Identifier     *Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn        []Reference                          `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Parent         []Reference                          `bson:"parent,omitempty" json:"parent,omitempty"`
	Questionnaire  *Reference                           `bson:"questionnaire,omitempty" json:"questionnaire,omitempty"`
	Status         string                               `bson:"status,omitempty" json:"status,omitempty"`
	Subject        *Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Context        *Reference                           `bson:"context,omitempty" json:"context,omitempty"`
	Authored       *FHIRDateTime                        `bson:"authored,omitempty" json:"authored,omitempty"`
	Author         *Reference                           `bson:"author,omitempty" json:"author,omitempty"`
	Source         *Reference                           `bson:"source,omitempty" json:"source,omitempty"`
	Item           []QuestionnaireResponseItemComponent `bson:"item,omitempty" json:"item,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "QuestionnaireResponse"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to QuestionnaireResponse), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *QuestionnaireResponse) GetBSON() (interface{}, error) {
	x.ResourceType = "QuestionnaireResponse"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "questionnaireResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type questionnaireResponse QuestionnaireResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *QuestionnaireResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := questionnaireResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = QuestionnaireResponse(x2)
		return x.checkResourceType()
	}
	return
}

func (x *QuestionnaireResponse) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "QuestionnaireResponse"
	} else if x.ResourceType != "QuestionnaireResponse" {
		return errors.New(fmt.Sprintf("Expected resourceType to be QuestionnaireResponse, instead received %s", x.ResourceType))
	}
	return nil
}

type QuestionnaireResponseItemComponent struct {
	BackboneElement `bson:",inline"`
	LinkId          string                                     `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Definition      string                                     `bson:"definition,omitempty" json:"definition,omitempty"`
	Text            string                                     `bson:"text,omitempty" json:"text,omitempty"`
	Subject         *Reference                                 `bson:"subject,omitempty" json:"subject,omitempty"`
	Answer          []QuestionnaireResponseItemAnswerComponent `bson:"answer,omitempty" json:"answer,omitempty"`
	Item            []QuestionnaireResponseItemComponent       `bson:"item,omitempty" json:"item,omitempty"`
}

type QuestionnaireResponseItemAnswerComponent struct {
	BackboneElement `bson:",inline"`
	ValueBoolean    *bool                                `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueDecimal    *float64                             `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueInteger    *int32                               `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate       *FHIRDateTime                        `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime   *FHIRDateTime                        `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueTime       *FHIRDateTime                        `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString     string                               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueUri        string                               `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueAttachment *Attachment                          `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCoding     *Coding                              `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueQuantity   *Quantity                            `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueReference  *Reference                           `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Item            []QuestionnaireResponseItemComponent `bson:"item,omitempty" json:"item,omitempty"`
}

type QuestionnaireResponsePlus struct {
	QuestionnaireResponse                     `bson:",inline"`
	QuestionnaireResponsePlusRelatedResources `bson:",inline"`
}

type QuestionnaireResponsePlusRelatedResources struct {
	IncludedObservationResourcesReferencedByParent                  *[]Observation           `bson:"_includedObservationResourcesReferencedByParent,omitempty"`
	IncludedProcedureResourcesReferencedByParent                    *[]Procedure             `bson:"_includedProcedureResourcesReferencedByParent,omitempty"`
	IncludedQuestionnaireResourcesReferencedByQuestionnaire         *[]Questionnaire         `bson:"_includedQuestionnaireResourcesReferencedByQuestionnaire,omitempty"`
	IncludedReferralRequestResourcesReferencedByBasedon             *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByBasedon,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                    *[]CarePlan              `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedProcedureRequestResourcesReferencedByBasedon            *[]ProcedureRequest      `bson:"_includedProcedureRequestResourcesReferencedByBasedon,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthor                 *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByAuthor                       *[]Device                `bson:"_includedDeviceResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByAuthor                      *[]Patient               `bson:"_includedPatientResourcesReferencedByAuthor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAuthor                *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext               *[]EpisodeOfCare         `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext                   *[]Encounter             `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
	IncludedPractitionerResourcesReferencedBySource                 *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySource,omitempty"`
	IncludedPatientResourcesReferencedBySource                      *[]Patient               `bson:"_includedPatientResourcesReferencedBySource,omitempty"`
	IncludedRelatedPersonResourcesReferencedBySource                *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedBySource,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedConsentResourcesReferencingSource                    *[]Consent               `bson:"_revIncludedConsentResourcesReferencingSource,omitempty"`
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
	RevIncludedObservationResourcesReferencingRelatedtarget         *[]Observation           `bson:"_revIncludedObservationResourcesReferencingRelatedtarget,omitempty"`
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
	RevIncludedClinicalImpressionResourcesReferencingInvestigation  *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingInvestigation,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor          *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof         *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedObservationResourcesReferencedByParent() (observations []Observation, err error) {
	if q.IncludedObservationResourcesReferencedByParent == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *q.IncludedObservationResourcesReferencedByParent
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedProcedureResourcesReferencedByParent() (procedures []Procedure, err error) {
	if q.IncludedProcedureResourcesReferencedByParent == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *q.IncludedProcedureResourcesReferencedByParent
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedQuestionnaireResourceReferencedByQuestionnaire() (questionnaire *Questionnaire, err error) {
	if q.IncludedQuestionnaireResourcesReferencedByQuestionnaire == nil {
		err = errors.New("Included questionnaires not requested")
	} else if len(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire) > 1 {
		err = fmt.Errorf("Expected 0 or 1 questionnaire, but found %d", len(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire))
	} else if len(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire) == 1 {
		questionnaire = &(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedReferralRequestResourcesReferencedByBasedon() (referralRequests []ReferralRequest, err error) {
	if q.IncludedReferralRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *q.IncludedReferralRequestResourcesReferencedByBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if q.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *q.IncludedCarePlanResourcesReferencedByBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedProcedureRequestResourcesReferencedByBasedon() (procedureRequests []ProcedureRequest, err error) {
	if q.IncludedProcedureRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included procedureRequests not requested")
	} else {
		procedureRequests = *q.IncludedProcedureRequestResourcesReferencedByBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if q.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*q.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*q.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*q.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*q.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedDeviceResourceReferencedByAuthor() (device *Device, err error) {
	if q.IncludedDeviceResourcesReferencedByAuthor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*q.IncludedDeviceResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*q.IncludedDeviceResourcesReferencedByAuthor))
	} else if len(*q.IncludedDeviceResourcesReferencedByAuthor) == 1 {
		device = &(*q.IncludedDeviceResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPatientResourceReferencedByAuthor() (patient *Patient, err error) {
	if q.IncludedPatientResourcesReferencedByAuthor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedPatientResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedPatientResourcesReferencedByAuthor))
	} else if len(*q.IncludedPatientResourcesReferencedByAuthor) == 1 {
		patient = &(*q.IncludedPatientResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAuthor() (relatedPerson *RelatedPerson, err error) {
	if q.IncludedRelatedPersonResourcesReferencedByAuthor == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*q.IncludedRelatedPersonResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*q.IncludedRelatedPersonResourcesReferencedByAuthor))
	} else if len(*q.IncludedRelatedPersonResourcesReferencedByAuthor) == 1 {
		relatedPerson = &(*q.IncludedRelatedPersonResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if q.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedPatientResourcesReferencedByPatient))
	} else if len(*q.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*q.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByContext() (episodeOfCare *EpisodeOfCare, err error) {
	if q.IncludedEpisodeOfCareResourcesReferencedByContext == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*q.IncludedEpisodeOfCareResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*q.IncludedEpisodeOfCareResourcesReferencedByContext))
	} else if len(*q.IncludedEpisodeOfCareResourcesReferencedByContext) == 1 {
		episodeOfCare = &(*q.IncludedEpisodeOfCareResourcesReferencedByContext)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedEncounterResourceReferencedByContext() (encounter *Encounter, err error) {
	if q.IncludedEncounterResourcesReferencedByContext == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*q.IncludedEncounterResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*q.IncludedEncounterResourcesReferencedByContext))
	} else if len(*q.IncludedEncounterResourcesReferencedByContext) == 1 {
		encounter = &(*q.IncludedEncounterResourcesReferencedByContext)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPractitionerResourceReferencedBySource() (practitioner *Practitioner, err error) {
	if q.IncludedPractitionerResourcesReferencedBySource == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*q.IncludedPractitionerResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*q.IncludedPractitionerResourcesReferencedBySource))
	} else if len(*q.IncludedPractitionerResourcesReferencedBySource) == 1 {
		practitioner = &(*q.IncludedPractitionerResourcesReferencedBySource)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPatientResourceReferencedBySource() (patient *Patient, err error) {
	if q.IncludedPatientResourcesReferencedBySource == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedPatientResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedPatientResourcesReferencedBySource))
	} else if len(*q.IncludedPatientResourcesReferencedBySource) == 1 {
		patient = &(*q.IncludedPatientResourcesReferencedBySource)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedBySource() (relatedPerson *RelatedPerson, err error) {
	if q.IncludedRelatedPersonResourcesReferencedBySource == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*q.IncludedRelatedPersonResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*q.IncludedRelatedPersonResourcesReferencedBySource))
	} else if len(*q.IncludedRelatedPersonResourcesReferencedBySource) == 1 {
		relatedPerson = &(*q.IncludedRelatedPersonResourcesReferencedBySource)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if q.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *q.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if q.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *q.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedConsentResourcesReferencingSource() (consents []Consent, err error) {
	if q.RevIncludedConsentResourcesReferencingSource == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *q.RevIncludedConsentResourcesReferencingSource
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if q.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *q.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if q.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *q.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if q.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *q.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if q.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *q.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if q.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *q.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if q.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *q.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if q.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *q.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if q.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *q.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if q.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *q.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if q.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *q.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if q.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *q.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if q.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *q.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if q.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *q.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if q.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *q.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if q.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *q.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if q.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *q.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if q.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *q.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if q.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *q.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if q.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *q.RevIncludedListResourcesReferencingItem
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if q.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *q.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if q.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *q.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedObservationResourcesReferencingRelatedtarget() (observations []Observation, err error) {
	if q.RevIncludedObservationResourcesReferencingRelatedtarget == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *q.RevIncludedObservationResourcesReferencingRelatedtarget
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if q.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *q.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if q.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *q.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if q.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *q.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if q.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *q.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *q.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if q.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *q.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingInvestigation() (clinicalImpressions []ClinicalImpression, err error) {
	if q.RevIncludedClinicalImpressionResourcesReferencingInvestigation == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *q.RevIncludedClinicalImpressionResourcesReferencingInvestigation
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.IncludedObservationResourcesReferencedByParent != nil {
		for idx := range *q.IncludedObservationResourcesReferencedByParent {
			rsc := (*q.IncludedObservationResourcesReferencedByParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedProcedureResourcesReferencedByParent != nil {
		for idx := range *q.IncludedProcedureResourcesReferencedByParent {
			rsc := (*q.IncludedProcedureResourcesReferencedByParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedQuestionnaireResourcesReferencedByQuestionnaire != nil {
		for idx := range *q.IncludedQuestionnaireResourcesReferencedByQuestionnaire {
			rsc := (*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedReferralRequestResourcesReferencedByBasedon != nil {
		for idx := range *q.IncludedReferralRequestResourcesReferencedByBasedon {
			rsc := (*q.IncludedReferralRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *q.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*q.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedProcedureRequestResourcesReferencedByBasedon != nil {
		for idx := range *q.IncludedProcedureRequestResourcesReferencedByBasedon {
			rsc := (*q.IncludedProcedureRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*q.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*q.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*q.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByPatient {
			rsc := (*q.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *q.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*q.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *q.IncludedEncounterResourcesReferencedByContext {
			rsc := (*q.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*q.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPatientResourcesReferencedBySource {
			rsc := (*q.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	if q.RevIncludedConsentResourcesReferencingSource != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingSource {
			rsc := (*q.RevIncludedConsentResourcesReferencingSource)[idx]
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
	if q.RevIncludedObservationResourcesReferencingRelatedtarget != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingRelatedtarget {
			rsc := (*q.RevIncludedObservationResourcesReferencingRelatedtarget)[idx]
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
	if q.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
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

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.IncludedObservationResourcesReferencedByParent != nil {
		for idx := range *q.IncludedObservationResourcesReferencedByParent {
			rsc := (*q.IncludedObservationResourcesReferencedByParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedProcedureResourcesReferencedByParent != nil {
		for idx := range *q.IncludedProcedureResourcesReferencedByParent {
			rsc := (*q.IncludedProcedureResourcesReferencedByParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedQuestionnaireResourcesReferencedByQuestionnaire != nil {
		for idx := range *q.IncludedQuestionnaireResourcesReferencedByQuestionnaire {
			rsc := (*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedReferralRequestResourcesReferencedByBasedon != nil {
		for idx := range *q.IncludedReferralRequestResourcesReferencedByBasedon {
			rsc := (*q.IncludedReferralRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *q.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*q.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedProcedureRequestResourcesReferencedByBasedon != nil {
		for idx := range *q.IncludedProcedureRequestResourcesReferencedByBasedon {
			rsc := (*q.IncludedProcedureRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*q.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*q.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*q.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByPatient {
			rsc := (*q.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *q.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*q.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *q.IncludedEncounterResourcesReferencedByContext {
			rsc := (*q.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*q.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPatientResourcesReferencedBySource {
			rsc := (*q.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
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
	if q.RevIncludedConsentResourcesReferencingSource != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingSource {
			rsc := (*q.RevIncludedConsentResourcesReferencingSource)[idx]
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
	if q.RevIncludedObservationResourcesReferencingRelatedtarget != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingRelatedtarget {
			rsc := (*q.RevIncludedObservationResourcesReferencingRelatedtarget)[idx]
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
	if q.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
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
