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

type Procedure struct {
	DomainResource     `bson:",inline"`
	Identifier         []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition         []Reference                     `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn            []Reference                     `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf             []Reference                     `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status             string                          `bson:"status,omitempty" json:"status,omitempty"`
	NotDone            *bool                           `bson:"notDone,omitempty" json:"notDone,omitempty"`
	NotDoneReason      *CodeableConcept                `bson:"notDoneReason,omitempty" json:"notDoneReason,omitempty"`
	Category           *CodeableConcept                `bson:"category,omitempty" json:"category,omitempty"`
	Code               *CodeableConcept                `bson:"code,omitempty" json:"code,omitempty"`
	Subject            *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Context            *Reference                      `bson:"context,omitempty" json:"context,omitempty"`
	PerformedDateTime  *FHIRDateTime                   `bson:"performedDateTime,omitempty" json:"performedDateTime,omitempty"`
	PerformedPeriod    *Period                         `bson:"performedPeriod,omitempty" json:"performedPeriod,omitempty"`
	Performer          []ProcedurePerformerComponent   `bson:"performer,omitempty" json:"performer,omitempty"`
	Location           *Reference                      `bson:"location,omitempty" json:"location,omitempty"`
	ReasonCode         []CodeableConcept               `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []Reference                     `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	BodySite           []CodeableConcept               `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Outcome            *CodeableConcept                `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Report             []Reference                     `bson:"report,omitempty" json:"report,omitempty"`
	Complication       []CodeableConcept               `bson:"complication,omitempty" json:"complication,omitempty"`
	ComplicationDetail []Reference                     `bson:"complicationDetail,omitempty" json:"complicationDetail,omitempty"`
	FollowUp           []CodeableConcept               `bson:"followUp,omitempty" json:"followUp,omitempty"`
	Note               []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
	FocalDevice        []ProcedureFocalDeviceComponent `bson:"focalDevice,omitempty" json:"focalDevice,omitempty"`
	UsedReference      []Reference                     `bson:"usedReference,omitempty" json:"usedReference,omitempty"`
	UsedCode           []CodeableConcept               `bson:"usedCode,omitempty" json:"usedCode,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Procedure) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Procedure"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Procedure), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Procedure) GetBSON() (interface{}, error) {
	x.ResourceType = "Procedure"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "procedure" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type procedure Procedure

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Procedure) UnmarshalJSON(data []byte) (err error) {
	x2 := procedure{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = Procedure(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Procedure) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Procedure"
	} else if x.ResourceType != "Procedure" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Procedure, instead received %s", x.ResourceType))
	}
	return nil
}

type ProcedurePerformerComponent struct {
	BackboneElement `bson:",inline"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Actor           *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
	OnBehalfOf      *Reference       `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}

type ProcedureFocalDeviceComponent struct {
	BackboneElement `bson:",inline"`
	Action          *CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Manipulated     *Reference       `bson:"manipulated,omitempty" json:"manipulated,omitempty"`
}

type ProcedurePlus struct {
	Procedure                     `bson:",inline"`
	ProcedurePlusRelatedResources `bson:",inline"`
}

type ProcedurePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByPerformer              *[]Practitioner             `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer              *[]Organization             `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedDeviceResourcesReferencedByPerformer                    *[]Device                   `bson:"_includedDeviceResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                   *[]Patient                  `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer             *[]RelatedPerson            `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedGroupResourcesReferencedBySubject                       *[]Group                    `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                     *[]Patient                  `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedObservationResourcesReferencedByPartof                  *[]Observation              `bson:"_includedObservationResourcesReferencedByPartof,omitempty"`
	IncludedProcedureResourcesReferencedByPartof                    *[]Procedure                `bson:"_includedProcedureResourcesReferencedByPartof,omitempty"`
	IncludedMedicationAdministrationResourcesReferencedByPartof     *[]MedicationAdministration `bson:"_includedMedicationAdministrationResourcesReferencedByPartof,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                 *[]Encounter                `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedReferralRequestResourcesReferencedByBasedon             *[]ReferralRequest          `bson:"_includedReferralRequestResourcesReferencedByBasedon,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                    *[]CarePlan                 `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedProcedureRequestResourcesReferencedByBasedon            *[]ProcedureRequest         `bson:"_includedProcedureRequestResourcesReferencedByBasedon,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient                  `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext               *[]EpisodeOfCare            `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext                   *[]Encounter                `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
	IncludedLocationResourcesReferencedByLocation                   *[]Location                 `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	IncludedPlanDefinitionResourcesReferencedByDefinition           *[]PlanDefinition           `bson:"_includedPlanDefinitionResourcesReferencedByDefinition,omitempty"`
	IncludedHealthcareServiceResourcesReferencedByDefinition        *[]HealthcareService        `bson:"_includedHealthcareServiceResourcesReferencedByDefinition,omitempty"`
	IncludedActivityDefinitionResourcesReferencedByDefinition       *[]ActivityDefinition       `bson:"_includedActivityDefinitionResourcesReferencedByDefinition,omitempty"`
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
	RevIncludedChargeItemResourcesReferencingService                *[]ChargeItem               `bson:"_revIncludedChargeItemResourcesReferencingService,omitempty"`
	RevIncludedEncounterResourcesReferencingDiagnosis               *[]Encounter                `bson:"_revIncludedEncounterResourcesReferencingDiagnosis,omitempty"`
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
	RevIncludedProcedureResourcesReferencingPartof                  *[]Procedure                `bson:"_revIncludedProcedureResourcesReferencingPartof,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                     `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces         *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon          *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                      *[]Flag                     `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                 *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom               *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor               *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                 *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingPartof        *[]MedicationStatement      `bson:"_revIncludedMedicationStatementResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon      *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                     *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                 *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail          *[]Condition                `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject               *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                 *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated          *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingParent      *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingParent,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject     *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest           *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAction         *[]ClinicalImpression       `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor          *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom        *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor        *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof         *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1     *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2     *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (p *ProcedurePlusRelatedResources) GetIncludedPractitionerResourceReferencedByPerformer() (practitioner *Practitioner, err error) {
	if p.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedPractitionerResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedPractitionerResourcesReferencedByPerformer))
	} else if len(*p.IncludedPractitionerResourcesReferencedByPerformer) == 1 {
		practitioner = &(*p.IncludedPractitionerResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedOrganizationResourceReferencedByPerformer() (organization *Organization, err error) {
	if p.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResourcesReferencedByPerformer))
	} else if len(*p.IncludedOrganizationResourcesReferencedByPerformer) == 1 {
		organization = &(*p.IncludedOrganizationResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedDeviceResourceReferencedByPerformer() (device *Device, err error) {
	if p.IncludedDeviceResourcesReferencedByPerformer == nil {
		err = errors.New("Included devices not requested")
	} else if len(*p.IncludedDeviceResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*p.IncludedDeviceResourcesReferencedByPerformer))
	} else if len(*p.IncludedDeviceResourcesReferencedByPerformer) == 1 {
		device = &(*p.IncludedDeviceResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPatientResourceReferencedByPerformer() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedByPerformer))
	} else if len(*p.IncludedPatientResourcesReferencedByPerformer) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByPerformer() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedRelatedPersonResourcesReferencedByPerformer))
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByPerformer) == 1 {
		relatedPerson = &(*p.IncludedRelatedPersonResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if p.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*p.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*p.IncludedGroupResourcesReferencedBySubject))
	} else if len(*p.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*p.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedBySubject))
	} else if len(*p.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedObservationResourcesReferencedByPartof() (observations []Observation, err error) {
	if p.IncludedObservationResourcesReferencedByPartof == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *p.IncludedObservationResourcesReferencedByPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedProcedureResourcesReferencedByPartof() (procedures []Procedure, err error) {
	if p.IncludedProcedureResourcesReferencedByPartof == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *p.IncludedProcedureResourcesReferencedByPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedMedicationAdministrationResourcesReferencedByPartof() (medicationAdministrations []MedicationAdministration, err error) {
	if p.IncludedMedicationAdministrationResourcesReferencedByPartof == nil {
		err = errors.New("Included medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.IncludedMedicationAdministrationResourcesReferencedByPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if p.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*p.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*p.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*p.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*p.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedReferralRequestResourcesReferencedByBasedon() (referralRequests []ReferralRequest, err error) {
	if p.IncludedReferralRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *p.IncludedReferralRequestResourcesReferencedByBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if p.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *p.IncludedCarePlanResourcesReferencedByBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedProcedureRequestResourcesReferencedByBasedon() (procedureRequests []ProcedureRequest, err error) {
	if p.IncludedProcedureRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included procedureRequests not requested")
	} else {
		procedureRequests = *p.IncludedProcedureRequestResourcesReferencedByBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedByPatient))
	} else if len(*p.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByContext() (episodeOfCare *EpisodeOfCare, err error) {
	if p.IncludedEpisodeOfCareResourcesReferencedByContext == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*p.IncludedEpisodeOfCareResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*p.IncludedEpisodeOfCareResourcesReferencedByContext))
	} else if len(*p.IncludedEpisodeOfCareResourcesReferencedByContext) == 1 {
		episodeOfCare = &(*p.IncludedEpisodeOfCareResourcesReferencedByContext)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedEncounterResourceReferencedByContext() (encounter *Encounter, err error) {
	if p.IncludedEncounterResourcesReferencedByContext == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*p.IncludedEncounterResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*p.IncludedEncounterResourcesReferencedByContext))
	} else if len(*p.IncludedEncounterResourcesReferencedByContext) == 1 {
		encounter = &(*p.IncludedEncounterResourcesReferencedByContext)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if p.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*p.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*p.IncludedLocationResourcesReferencedByLocation))
	} else if len(*p.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*p.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPlanDefinitionResourcesReferencedByDefinition() (planDefinitions []PlanDefinition, err error) {
	if p.IncludedPlanDefinitionResourcesReferencedByDefinition == nil {
		err = errors.New("Included planDefinitions not requested")
	} else {
		planDefinitions = *p.IncludedPlanDefinitionResourcesReferencedByDefinition
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedHealthcareServiceResourcesReferencedByDefinition() (healthcareServices []HealthcareService, err error) {
	if p.IncludedHealthcareServiceResourcesReferencedByDefinition == nil {
		err = errors.New("Included healthcareServices not requested")
	} else {
		healthcareServices = *p.IncludedHealthcareServiceResourcesReferencedByDefinition
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedActivityDefinitionResourcesReferencedByDefinition() (activityDefinitions []ActivityDefinition, err error) {
	if p.IncludedActivityDefinitionResourcesReferencedByDefinition == nil {
		err = errors.New("Included activityDefinitions not requested")
	} else {
		activityDefinitions = *p.IncludedActivityDefinitionResourcesReferencedByDefinition
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if p.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *p.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingService() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingService == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingService
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEncounterResourcesReferencingDiagnosis() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingDiagnosis == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingDiagnosis
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if p.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *p.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPartof() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingItem
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if p.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *p.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingPartof() (medicationStatements []MedicationStatement, err error) {
	if p.RevIncludedMedicationStatementResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *p.RevIncludedMedicationStatementResourcesReferencingPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingParent() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingParent == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingParent
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if p.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *p.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAction() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingAction == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingAction
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*p.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*p.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*p.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*p.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedGroupResourcesReferencedBySubject {
			rsc := (*p.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedPatientResourcesReferencedBySubject {
			rsc := (*p.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedObservationResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedObservationResourcesReferencedByPartof {
			rsc := (*p.IncludedObservationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*p.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedMedicationAdministrationResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedMedicationAdministrationResourcesReferencedByPartof {
			rsc := (*p.IncludedMedicationAdministrationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *p.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*p.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedReferralRequestResourcesReferencedByBasedon != nil {
		for idx := range *p.IncludedReferralRequestResourcesReferencedByBasedon {
			rsc := (*p.IncludedReferralRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *p.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*p.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedProcedureRequestResourcesReferencedByBasedon != nil {
		for idx := range *p.IncludedProcedureRequestResourcesReferencedByBasedon {
			rsc := (*p.IncludedProcedureRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPatient {
			rsc := (*p.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *p.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*p.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *p.IncludedEncounterResourcesReferencedByContext {
			rsc := (*p.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *p.IncludedLocationResourcesReferencedByLocation {
			rsc := (*p.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPlanDefinitionResourcesReferencedByDefinition != nil {
		for idx := range *p.IncludedPlanDefinitionResourcesReferencedByDefinition {
			rsc := (*p.IncludedPlanDefinitionResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedHealthcareServiceResourcesReferencedByDefinition != nil {
		for idx := range *p.IncludedHealthcareServiceResourcesReferencedByDefinition {
			rsc := (*p.IncludedHealthcareServiceResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedActivityDefinitionResourcesReferencedByDefinition != nil {
		for idx := range *p.IncludedActivityDefinitionResourcesReferencedByDefinition {
			rsc := (*p.IncludedActivityDefinitionResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	if p.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*p.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*p.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*p.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTermtopic)[idx]
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
	if p.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingDiagnosis != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingDiagnosis {
			rsc := (*p.RevIncludedEncounterResourcesReferencingDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if p.RevIncludedProcedureResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPartof {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*p.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingPartof {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if p.RevIncludedQuestionnaireResponseResourcesReferencingParent != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingParent {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingParent)[idx]
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
	if p.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *ProcedurePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*p.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*p.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*p.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*p.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedGroupResourcesReferencedBySubject {
			rsc := (*p.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedPatientResourcesReferencedBySubject {
			rsc := (*p.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedObservationResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedObservationResourcesReferencedByPartof {
			rsc := (*p.IncludedObservationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*p.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedMedicationAdministrationResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedMedicationAdministrationResourcesReferencedByPartof {
			rsc := (*p.IncludedMedicationAdministrationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *p.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*p.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedReferralRequestResourcesReferencedByBasedon != nil {
		for idx := range *p.IncludedReferralRequestResourcesReferencedByBasedon {
			rsc := (*p.IncludedReferralRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *p.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*p.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedProcedureRequestResourcesReferencedByBasedon != nil {
		for idx := range *p.IncludedProcedureRequestResourcesReferencedByBasedon {
			rsc := (*p.IncludedProcedureRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPatient {
			rsc := (*p.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *p.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*p.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *p.IncludedEncounterResourcesReferencedByContext {
			rsc := (*p.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *p.IncludedLocationResourcesReferencedByLocation {
			rsc := (*p.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPlanDefinitionResourcesReferencedByDefinition != nil {
		for idx := range *p.IncludedPlanDefinitionResourcesReferencedByDefinition {
			rsc := (*p.IncludedPlanDefinitionResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedHealthcareServiceResourcesReferencedByDefinition != nil {
		for idx := range *p.IncludedHealthcareServiceResourcesReferencedByDefinition {
			rsc := (*p.IncludedHealthcareServiceResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedActivityDefinitionResourcesReferencedByDefinition != nil {
		for idx := range *p.IncludedActivityDefinitionResourcesReferencedByDefinition {
			rsc := (*p.IncludedActivityDefinitionResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
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
	if p.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*p.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*p.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*p.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTermtopic)[idx]
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
	if p.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingDiagnosis != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingDiagnosis {
			rsc := (*p.RevIncludedEncounterResourcesReferencingDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*p.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if p.RevIncludedProcedureResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPartof {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*p.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingPartof {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if p.RevIncludedQuestionnaireResponseResourcesReferencingParent != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingParent {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingParent)[idx]
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
	if p.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
