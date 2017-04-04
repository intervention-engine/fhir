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

type AdverseEvent struct {
	DomainResource        `bson:",inline"`
	Identifier            *Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Category              string                               `bson:"category,omitempty" json:"category,omitempty"`
	Type                  *CodeableConcept                     `bson:"type,omitempty" json:"type,omitempty"`
	Subject               *Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Date                  *FHIRDateTime                        `bson:"date,omitempty" json:"date,omitempty"`
	Reaction              []Reference                          `bson:"reaction,omitempty" json:"reaction,omitempty"`
	Location              *Reference                           `bson:"location,omitempty" json:"location,omitempty"`
	Seriousness           *CodeableConcept                     `bson:"seriousness,omitempty" json:"seriousness,omitempty"`
	Outcome               *CodeableConcept                     `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Recorder              *Reference                           `bson:"recorder,omitempty" json:"recorder,omitempty"`
	EventParticipant      *Reference                           `bson:"eventParticipant,omitempty" json:"eventParticipant,omitempty"`
	Description           string                               `bson:"description,omitempty" json:"description,omitempty"`
	SuspectEntity         []AdverseEventSuspectEntityComponent `bson:"suspectEntity,omitempty" json:"suspectEntity,omitempty"`
	SubjectMedicalHistory []Reference                          `bson:"subjectMedicalHistory,omitempty" json:"subjectMedicalHistory,omitempty"`
	ReferenceDocument     []Reference                          `bson:"referenceDocument,omitempty" json:"referenceDocument,omitempty"`
	Study                 []Reference                          `bson:"study,omitempty" json:"study,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *AdverseEvent) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "AdverseEvent"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to AdverseEvent), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *AdverseEvent) GetBSON() (interface{}, error) {
	x.ResourceType = "AdverseEvent"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "adverseEvent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type adverseEvent AdverseEvent

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *AdverseEvent) UnmarshalJSON(data []byte) (err error) {
	x2 := adverseEvent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = AdverseEvent(x2)
		return x.checkResourceType()
	}
	return
}

func (x *AdverseEvent) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "AdverseEvent"
	} else if x.ResourceType != "AdverseEvent" {
		return errors.New(fmt.Sprintf("Expected resourceType to be AdverseEvent, instead received %s", x.ResourceType))
	}
	return nil
}

type AdverseEventSuspectEntityComponent struct {
	BackboneElement             `bson:",inline"`
	Instance                    *Reference       `bson:"instance,omitempty" json:"instance,omitempty"`
	Causality                   string           `bson:"causality,omitempty" json:"causality,omitempty"`
	CausalityAssessment         *CodeableConcept `bson:"causalityAssessment,omitempty" json:"causalityAssessment,omitempty"`
	CausalityProductRelatedness string           `bson:"causalityProductRelatedness,omitempty" json:"causalityProductRelatedness,omitempty"`
	CausalityMethod             *CodeableConcept `bson:"causalityMethod,omitempty" json:"causalityMethod,omitempty"`
	CausalityAuthor             *Reference       `bson:"causalityAuthor,omitempty" json:"causalityAuthor,omitempty"`
	CausalityResult             *CodeableConcept `bson:"causalityResult,omitempty" json:"causalityResult,omitempty"`
}

type AdverseEventPlus struct {
	AdverseEvent                     `bson:",inline"`
	AdverseEventPlusRelatedResources `bson:",inline"`
}

type AdverseEventPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByRecorder               *[]Practitioner             `bson:"_includedPractitionerResourcesReferencedByRecorder,omitempty"`
	IncludedPatientResourcesReferencedByRecorder                    *[]Patient                  `bson:"_includedPatientResourcesReferencedByRecorder,omitempty"`
	IncludedRelatedPersonResourcesReferencedByRecorder              *[]RelatedPerson            `bson:"_includedRelatedPersonResourcesReferencedByRecorder,omitempty"`
	IncludedResearchStudyResourcesReferencedByStudy                 *[]ResearchStudy            `bson:"_includedResearchStudyResourcesReferencedByStudy,omitempty"`
	IncludedConditionResourcesReferencedByReaction                  *[]Condition                `bson:"_includedConditionResourcesReferencedByReaction,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                      *[]Device                   `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedMedicationResourcesReferencedBySubject                  *[]Medication               `bson:"_includedMedicationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                     *[]Patient                  `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedResearchSubjectResourcesReferencedBySubject             *[]ResearchSubject          `bson:"_includedResearchSubjectResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubstance                    *[]Device                   `bson:"_includedDeviceResourcesReferencedBySubstance,omitempty"`
	IncludedMedicationResourcesReferencedBySubstance                *[]Medication               `bson:"_includedMedicationResourcesReferencedBySubstance,omitempty"`
	IncludedSubstanceResourcesReferencedBySubstance                 *[]Substance                `bson:"_includedSubstanceResourcesReferencedBySubstance,omitempty"`
	IncludedMedicationAdministrationResourcesReferencedBySubstance  *[]MedicationAdministration `bson:"_includedMedicationAdministrationResourcesReferencedBySubstance,omitempty"`
	IncludedMedicationStatementResourcesReferencedBySubstance       *[]MedicationStatement      `bson:"_includedMedicationStatementResourcesReferencedBySubstance,omitempty"`
	IncludedLocationResourcesReferencedByLocation                   *[]Location                 `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref       *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref       *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                 *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                 *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                 *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom               *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor               *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1            *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2            *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref      *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingSubject                  *[]Contract                 `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                *[]Contract                 `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest             *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse            *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource      *[]ImplementationGuide      `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor       *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom     *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor     *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof      *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson       *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof              *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon             *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor      *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom    *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor    *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof     *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition           *[]RequestGroup             `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon             *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest        *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus               *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref              *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                 *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                      *[]Task                     `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                        *[]Task                     `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                      *[]Task                     `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                     `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces         *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon          *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                 *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom               *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor               *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                 *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon      *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                     *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                 *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail          *[]Condition                `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject               *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                 *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated          *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject     *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest           *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor          *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom        *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor        *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof         *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1     *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2     *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (a *AdverseEventPlusRelatedResources) GetIncludedPractitionerResourceReferencedByRecorder() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByRecorder == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByRecorder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByRecorder))
	} else if len(*a.IncludedPractitionerResourcesReferencedByRecorder) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByRecorder)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedPatientResourceReferencedByRecorder() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByRecorder == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByRecorder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByRecorder))
	} else if len(*a.IncludedPatientResourcesReferencedByRecorder) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByRecorder)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByRecorder() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedRelatedPersonResourcesReferencedByRecorder == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByRecorder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedRelatedPersonResourcesReferencedByRecorder))
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByRecorder) == 1 {
		relatedPerson = &(*a.IncludedRelatedPersonResourcesReferencedByRecorder)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedResearchStudyResourcesReferencedByStudy() (researchStudies []ResearchStudy, err error) {
	if a.IncludedResearchStudyResourcesReferencedByStudy == nil {
		err = errors.New("Included researchStudies not requested")
	} else {
		researchStudies = *a.IncludedResearchStudyResourcesReferencedByStudy
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedConditionResourcesReferencedByReaction() (conditions []Condition, err error) {
	if a.IncludedConditionResourcesReferencedByReaction == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *a.IncludedConditionResourcesReferencedByReaction
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if a.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*a.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*a.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedMedicationResourceReferencedBySubject() (medication *Medication, err error) {
	if a.IncludedMedicationResourcesReferencedBySubject == nil {
		err = errors.New("Included medications not requested")
	} else if len(*a.IncludedMedicationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*a.IncludedMedicationResourcesReferencedBySubject))
	} else if len(*a.IncludedMedicationResourcesReferencedBySubject) == 1 {
		medication = &(*a.IncludedMedicationResourcesReferencedBySubject)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedBySubject))
	} else if len(*a.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedResearchSubjectResourceReferencedBySubject() (researchSubject *ResearchSubject, err error) {
	if a.IncludedResearchSubjectResourcesReferencedBySubject == nil {
		err = errors.New("Included researchsubjects not requested")
	} else if len(*a.IncludedResearchSubjectResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 researchSubject, but found %d", len(*a.IncludedResearchSubjectResourcesReferencedBySubject))
	} else if len(*a.IncludedResearchSubjectResourcesReferencedBySubject) == 1 {
		researchSubject = &(*a.IncludedResearchSubjectResourcesReferencedBySubject)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubstance() (device *Device, err error) {
	if a.IncludedDeviceResourcesReferencedBySubstance == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedDeviceResourcesReferencedBySubstance) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedDeviceResourcesReferencedBySubstance))
	} else if len(*a.IncludedDeviceResourcesReferencedBySubstance) == 1 {
		device = &(*a.IncludedDeviceResourcesReferencedBySubstance)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedMedicationResourceReferencedBySubstance() (medication *Medication, err error) {
	if a.IncludedMedicationResourcesReferencedBySubstance == nil {
		err = errors.New("Included medications not requested")
	} else if len(*a.IncludedMedicationResourcesReferencedBySubstance) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*a.IncludedMedicationResourcesReferencedBySubstance))
	} else if len(*a.IncludedMedicationResourcesReferencedBySubstance) == 1 {
		medication = &(*a.IncludedMedicationResourcesReferencedBySubstance)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedSubstanceResourceReferencedBySubstance() (substance *Substance, err error) {
	if a.IncludedSubstanceResourcesReferencedBySubstance == nil {
		err = errors.New("Included substances not requested")
	} else if len(*a.IncludedSubstanceResourcesReferencedBySubstance) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*a.IncludedSubstanceResourcesReferencedBySubstance))
	} else if len(*a.IncludedSubstanceResourcesReferencedBySubstance) == 1 {
		substance = &(*a.IncludedSubstanceResourcesReferencedBySubstance)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedMedicationAdministrationResourceReferencedBySubstance() (medicationAdministration *MedicationAdministration, err error) {
	if a.IncludedMedicationAdministrationResourcesReferencedBySubstance == nil {
		err = errors.New("Included medicationadministrations not requested")
	} else if len(*a.IncludedMedicationAdministrationResourcesReferencedBySubstance) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medicationAdministration, but found %d", len(*a.IncludedMedicationAdministrationResourcesReferencedBySubstance))
	} else if len(*a.IncludedMedicationAdministrationResourcesReferencedBySubstance) == 1 {
		medicationAdministration = &(*a.IncludedMedicationAdministrationResourcesReferencedBySubstance)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedMedicationStatementResourceReferencedBySubstance() (medicationStatement *MedicationStatement, err error) {
	if a.IncludedMedicationStatementResourcesReferencedBySubstance == nil {
		err = errors.New("Included medicationstatements not requested")
	} else if len(*a.IncludedMedicationStatementResourcesReferencedBySubstance) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medicationStatement, but found %d", len(*a.IncludedMedicationStatementResourcesReferencedBySubstance))
	} else if len(*a.IncludedMedicationStatementResourcesReferencedBySubstance) == 1 {
		medicationStatement = &(*a.IncludedMedicationStatementResourcesReferencedBySubstance)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if a.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*a.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*a.IncludedLocationResourcesReferencedByLocation))
	} else if len(*a.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*a.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if a.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *a.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if a.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *a.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if a.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *a.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if a.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *a.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if a.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *a.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if a.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *a.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if a.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *a.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if a.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *a.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if a.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *a.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if a.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *a.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if a.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *a.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *a.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if a.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *a.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if a.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *a.RevIncludedListResourcesReferencingItem
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if a.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *a.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if a.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *a.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *a.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if a.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *a.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if a.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *a.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if a.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *a.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *a.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if a.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *a.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (a *AdverseEventPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByRecorder {
			rsc := (*a.IncludedPractitionerResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByRecorder {
			rsc := (*a.IncludedPatientResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByRecorder {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedResearchStudyResourcesReferencedByStudy != nil {
		for idx := range *a.IncludedResearchStudyResourcesReferencedByStudy {
			rsc := (*a.IncludedResearchStudyResourcesReferencedByStudy)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedConditionResourcesReferencedByReaction != nil {
		for idx := range *a.IncludedConditionResourcesReferencedByReaction {
			rsc := (*a.IncludedConditionResourcesReferencedByReaction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*a.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedMedicationResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedMedicationResourcesReferencedBySubject {
			rsc := (*a.IncludedMedicationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedPatientResourcesReferencedBySubject {
			rsc := (*a.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedResearchSubjectResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedResearchSubjectResourcesReferencedBySubject {
			rsc := (*a.IncludedResearchSubjectResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedBySubstance != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedBySubstance {
			rsc := (*a.IncludedDeviceResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedMedicationResourcesReferencedBySubstance != nil {
		for idx := range *a.IncludedMedicationResourcesReferencedBySubstance {
			rsc := (*a.IncludedMedicationResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedSubstanceResourcesReferencedBySubstance != nil {
		for idx := range *a.IncludedSubstanceResourcesReferencedBySubstance {
			rsc := (*a.IncludedSubstanceResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedMedicationAdministrationResourcesReferencedBySubstance != nil {
		for idx := range *a.IncludedMedicationAdministrationResourcesReferencedBySubstance {
			rsc := (*a.IncludedMedicationAdministrationResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedMedicationStatementResourcesReferencedBySubstance != nil {
		for idx := range *a.IncludedMedicationStatementResourcesReferencedBySubstance {
			rsc := (*a.IncludedMedicationStatementResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *a.IncludedLocationResourcesReferencedByLocation {
			rsc := (*a.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AdverseEventPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*a.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*a.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*a.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*a.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *a.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*a.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *a.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*a.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*a.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *a.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*a.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*a.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AdverseEventPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByRecorder {
			rsc := (*a.IncludedPractitionerResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByRecorder {
			rsc := (*a.IncludedPatientResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByRecorder {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedResearchStudyResourcesReferencedByStudy != nil {
		for idx := range *a.IncludedResearchStudyResourcesReferencedByStudy {
			rsc := (*a.IncludedResearchStudyResourcesReferencedByStudy)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedConditionResourcesReferencedByReaction != nil {
		for idx := range *a.IncludedConditionResourcesReferencedByReaction {
			rsc := (*a.IncludedConditionResourcesReferencedByReaction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*a.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedMedicationResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedMedicationResourcesReferencedBySubject {
			rsc := (*a.IncludedMedicationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedPatientResourcesReferencedBySubject {
			rsc := (*a.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedResearchSubjectResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedResearchSubjectResourcesReferencedBySubject {
			rsc := (*a.IncludedResearchSubjectResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedBySubstance != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedBySubstance {
			rsc := (*a.IncludedDeviceResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedMedicationResourcesReferencedBySubstance != nil {
		for idx := range *a.IncludedMedicationResourcesReferencedBySubstance {
			rsc := (*a.IncludedMedicationResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedSubstanceResourcesReferencedBySubstance != nil {
		for idx := range *a.IncludedSubstanceResourcesReferencedBySubstance {
			rsc := (*a.IncludedSubstanceResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedMedicationAdministrationResourcesReferencedBySubstance != nil {
		for idx := range *a.IncludedMedicationAdministrationResourcesReferencedBySubstance {
			rsc := (*a.IncludedMedicationAdministrationResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedMedicationStatementResourcesReferencedBySubstance != nil {
		for idx := range *a.IncludedMedicationStatementResourcesReferencedBySubstance {
			rsc := (*a.IncludedMedicationStatementResourcesReferencedBySubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *a.IncludedLocationResourcesReferencedByLocation {
			rsc := (*a.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*a.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*a.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*a.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*a.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *a.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*a.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *a.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*a.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*a.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *a.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*a.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*a.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
