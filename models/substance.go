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

type Substance struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status         string                         `bson:"status,omitempty" json:"status,omitempty"`
	Category       []CodeableConcept              `bson:"category,omitempty" json:"category,omitempty"`
	Code           *CodeableConcept               `bson:"code,omitempty" json:"code,omitempty"`
	Description    string                         `bson:"description,omitempty" json:"description,omitempty"`
	Instance       []SubstanceInstanceComponent   `bson:"instance,omitempty" json:"instance,omitempty"`
	Ingredient     []SubstanceIngredientComponent `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Substance) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Substance"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Substance), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Substance) GetBSON() (interface{}, error) {
	x.ResourceType = "Substance"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "substance" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type substance Substance

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Substance) UnmarshalJSON(data []byte) (err error) {
	x2 := substance{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = Substance(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Substance) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Substance"
	} else if x.ResourceType != "Substance" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Substance, instead received %s", x.ResourceType))
	}
	return nil
}

type SubstanceInstanceComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      *Identifier   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Expiry          *FHIRDateTime `bson:"expiry,omitempty" json:"expiry,omitempty"`
	Quantity        *Quantity     `bson:"quantity,omitempty" json:"quantity,omitempty"`
}

type SubstanceIngredientComponent struct {
	BackboneElement          `bson:",inline"`
	Quantity                 *Ratio           `bson:"quantity,omitempty" json:"quantity,omitempty"`
	SubstanceCodeableConcept *CodeableConcept `bson:"substanceCodeableConcept,omitempty" json:"substanceCodeableConcept,omitempty"`
	SubstanceReference       *Reference       `bson:"substanceReference,omitempty" json:"substanceReference,omitempty"`
}

type SubstancePlus struct {
	Substance                     `bson:",inline"`
	SubstancePlusRelatedResources `bson:",inline"`
}

type SubstancePlusRelatedResources struct {
	IncludedSubstanceResourcesReferencedBySubstancereference        *[]Substance             `bson:"_includedSubstanceResourcesReferencedBySubstancereference,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedMedicationResourcesReferencingIngredient             *[]Medication            `bson:"_revIncludedMedicationResourcesReferencingIngredient,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                 *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1            *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2            *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref      *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingAgent                    *[]Contract              `bson:"_revIncludedContractResourcesReferencingAgent,omitempty"`
	RevIncludedContractResourcesReferencingSubject                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedGroupResourcesReferencingMember                      *[]Group                 `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest             *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse            *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource      *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedSubstanceResourcesReferencingSubstancereference      *[]Substance             `bson:"_revIncludedSubstanceResourcesReferencingSubstancereference,omitempty"`
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
	RevIncludedSpecimenResourcesReferencingSubject                  *[]Specimen              `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces         *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon          *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedAdverseEventResourcesReferencingSubstance            *[]AdverseEvent          `bson:"_revIncludedAdverseEventResourcesReferencingSubstance,omitempty"`
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

func (s *SubstancePlusRelatedResources) GetIncludedSubstanceResourceReferencedBySubstancereference() (substance *Substance, err error) {
	if s.IncludedSubstanceResourcesReferencedBySubstancereference == nil {
		err = errors.New("Included substances not requested")
	} else if len(*s.IncludedSubstanceResourcesReferencedBySubstancereference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*s.IncludedSubstanceResourcesReferencedBySubstancereference))
	} else if len(*s.IncludedSubstanceResourcesReferencedBySubstancereference) == 1 {
		substance = &(*s.IncludedSubstanceResourcesReferencedBySubstancereference)[0]
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if s.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *s.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if s.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *s.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedMedicationResourcesReferencingIngredient() (medications []Medication, err error) {
	if s.RevIncludedMedicationResourcesReferencingIngredient == nil {
		err = errors.New("RevIncluded medications not requested")
	} else {
		medications = *s.RevIncludedMedicationResourcesReferencingIngredient
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedContractResourcesReferencingAgent() (contracts []Contract, err error) {
	if s.RevIncludedContractResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *s.RevIncludedContractResourcesReferencingAgent
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if s.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *s.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if s.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *s.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if s.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *s.RevIncludedGroupResourcesReferencingMember
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if s.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *s.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedSubstanceResourcesReferencingSubstancereference() (substances []Substance, err error) {
	if s.RevIncludedSubstanceResourcesReferencingSubstancereference == nil {
		err = errors.New("RevIncluded substances not requested")
	} else {
		substances = *s.RevIncludedSubstanceResourcesReferencingSubstancereference
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if s.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *s.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if s.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *s.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if s.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *s.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if s.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *s.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if s.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *s.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if s.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *s.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if s.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *s.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if s.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *s.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *s.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if s.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *s.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingSubject() (specimen []Specimen, err error) {
	if s.RevIncludedSpecimenResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *s.RevIncludedSpecimenResourcesReferencingSubject
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if s.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *s.RevIncludedListResourcesReferencingItem
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if s.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *s.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if s.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *s.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubstance() (adverseEvents []AdverseEvent, err error) {
	if s.RevIncludedAdverseEventResourcesReferencingSubstance == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *s.RevIncludedAdverseEventResourcesReferencingSubstance
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *s.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if s.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *s.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if s.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *s.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if s.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *s.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *s.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if s.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *s.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *SubstancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *SubstancePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedSubstanceResourcesReferencedBySubstancereference != nil {
		for idx := range *s.IncludedSubstanceResourcesReferencedBySubstancereference {
			rsc := (*s.IncludedSubstanceResourcesReferencedBySubstancereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *SubstancePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*s.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*s.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMedicationResourcesReferencingIngredient != nil {
		for idx := range *s.RevIncludedMedicationResourcesReferencingIngredient {
			rsc := (*s.RevIncludedMedicationResourcesReferencingIngredient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*s.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingAgent {
			rsc := (*s.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingSubject {
			rsc := (*s.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*s.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *s.RevIncludedGroupResourcesReferencingMember {
			rsc := (*s.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *s.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*s.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSubstanceResourcesReferencingSubstancereference != nil {
		for idx := range *s.RevIncludedSubstanceResourcesReferencingSubstancereference {
			rsc := (*s.RevIncludedSubstanceResourcesReferencingSubstancereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*s.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *s.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*s.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*s.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*s.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*s.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*s.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *s.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*s.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *s.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*s.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*s.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *s.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*s.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*s.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *SubstancePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedSubstanceResourcesReferencedBySubstancereference != nil {
		for idx := range *s.IncludedSubstanceResourcesReferencedBySubstancereference {
			rsc := (*s.IncludedSubstanceResourcesReferencedBySubstancereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*s.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*s.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMedicationResourcesReferencingIngredient != nil {
		for idx := range *s.RevIncludedMedicationResourcesReferencingIngredient {
			rsc := (*s.RevIncludedMedicationResourcesReferencingIngredient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*s.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingAgent {
			rsc := (*s.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingSubject {
			rsc := (*s.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*s.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *s.RevIncludedGroupResourcesReferencingMember {
			rsc := (*s.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *s.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*s.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSubstanceResourcesReferencingSubstancereference != nil {
		for idx := range *s.RevIncludedSubstanceResourcesReferencingSubstancereference {
			rsc := (*s.RevIncludedSubstanceResourcesReferencingSubstancereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*s.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *s.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*s.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*s.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*s.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*s.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*s.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *s.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*s.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *s.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*s.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*s.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *s.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*s.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*s.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
