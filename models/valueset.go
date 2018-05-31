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

type ValueSet struct {
	DomainResource `bson:",inline"`
	Url            string                      `bson:"url,omitempty" json:"url,omitempty"`
	Identifier     []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version        string                      `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                      `bson:"name,omitempty" json:"name,omitempty"`
	Title          string                      `bson:"title,omitempty" json:"title,omitempty"`
	Status         string                      `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                       `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date           *FHIRDateTime               `bson:"date,omitempty" json:"date,omitempty"`
	Publisher      string                      `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ContactDetail             `bson:"contact,omitempty" json:"contact,omitempty"`
	Description    string                      `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []UsageContext              `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction   []CodeableConcept           `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Immutable      *bool                       `bson:"immutable,omitempty" json:"immutable,omitempty"`
	Purpose        string                      `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright      string                      `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Extensible     *bool                       `bson:"extensible,omitempty" json:"extensible,omitempty"`
	Compose        *ValueSetComposeComponent   `bson:"compose,omitempty" json:"compose,omitempty"`
	Expansion      *ValueSetExpansionComponent `bson:"expansion,omitempty" json:"expansion,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ValueSet) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ValueSet"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ValueSet), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ValueSet) GetBSON() (interface{}, error) {
	x.ResourceType = "ValueSet"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "valueSet" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type valueSet ValueSet

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ValueSet) UnmarshalJSON(data []byte) (err error) {
	x2 := valueSet{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = ValueSet(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ValueSet) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ValueSet"
	} else if x.ResourceType != "ValueSet" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ValueSet, instead received %s", x.ResourceType))
	}
	return nil
}

type ValueSetComposeComponent struct {
	BackboneElement `bson:",inline"`
	LockedDate      *FHIRDateTime                 `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	Inactive        *bool                         `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Include         []ValueSetConceptSetComponent `bson:"include,omitempty" json:"include,omitempty"`
	Exclude         []ValueSetConceptSetComponent `bson:"exclude,omitempty" json:"exclude,omitempty"`
}

type ValueSetConceptSetComponent struct {
	BackboneElement `bson:",inline"`
	System          string                              `bson:"system,omitempty" json:"system,omitempty"`
	Version         string                              `bson:"version,omitempty" json:"version,omitempty"`
	Concept         []ValueSetConceptReferenceComponent `bson:"concept,omitempty" json:"concept,omitempty"`
	Filter          []ValueSetConceptSetFilterComponent `bson:"filter,omitempty" json:"filter,omitempty"`
	ValueSet        []string                            `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
}

type ValueSetConceptReferenceComponent struct {
	BackboneElement `bson:",inline"`
	Code            string                                         `bson:"code,omitempty" json:"code,omitempty"`
	Display         string                                         `bson:"display,omitempty" json:"display,omitempty"`
	Designation     []ValueSetConceptReferenceDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
}

type ValueSetConceptReferenceDesignationComponent struct {
	BackboneElement `bson:",inline"`
	Language        string  `bson:"language,omitempty" json:"language,omitempty"`
	Use             *Coding `bson:"use,omitempty" json:"use,omitempty"`
	Value           string  `bson:"value,omitempty" json:"value,omitempty"`
}

type ValueSetConceptSetFilterComponent struct {
	BackboneElement `bson:",inline"`
	Property        string `bson:"property,omitempty" json:"property,omitempty"`
	Op              string `bson:"op,omitempty" json:"op,omitempty"`
	Value           string `bson:"value,omitempty" json:"value,omitempty"`
}

type ValueSetExpansionComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      string                                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Timestamp       *FHIRDateTime                         `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Total           *int32                                `bson:"total,omitempty" json:"total,omitempty"`
	Offset          *int32                                `bson:"offset,omitempty" json:"offset,omitempty"`
	Parameter       []ValueSetExpansionParameterComponent `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Contains        []ValueSetExpansionContainsComponent  `bson:"contains,omitempty" json:"contains,omitempty"`
}

type ValueSetExpansionParameterComponent struct {
	BackboneElement `bson:",inline"`
	Name            string   `bson:"name,omitempty" json:"name,omitempty"`
	ValueString     string   `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean    *bool    `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger    *int32   `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDecimal    *float64 `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueUri        string   `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueCode       string   `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
}

type ValueSetExpansionContainsComponent struct {
	BackboneElement `bson:",inline"`
	System          string                                         `bson:"system,omitempty" json:"system,omitempty"`
	Abstract        *bool                                          `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Inactive        *bool                                          `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Version         string                                         `bson:"version,omitempty" json:"version,omitempty"`
	Code            string                                         `bson:"code,omitempty" json:"code,omitempty"`
	Display         string                                         `bson:"display,omitempty" json:"display,omitempty"`
	Designation     []ValueSetConceptReferenceDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
	Contains        []ValueSetExpansionContainsComponent           `bson:"contains,omitempty" json:"contains,omitempty"`
}

type ValueSetPlus struct {
	ValueSet                     `bson:",inline"`
	ValueSetPlusRelatedResources `bson:",inline"`
}

type ValueSetPlusRelatedResources struct {
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
	RevIncludedStructureDefinitionResourcesReferencingValueset      *[]StructureDefinition   `bson:"_revIncludedStructureDefinitionResourcesReferencingValueset,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedConceptMapResourcesReferencingSource                 *[]ConceptMap            `bson:"_revIncludedConceptMapResourcesReferencingSource,omitempty"`
	RevIncludedConceptMapResourcesReferencingTarget                 *[]ConceptMap            `bson:"_revIncludedConceptMapResourcesReferencingTarget,omitempty"`
	RevIncludedConceptMapResourcesReferencingSourceuri              *[]ConceptMap            `bson:"_revIncludedConceptMapResourcesReferencingSourceuri,omitempty"`
	RevIncludedConceptMapResourcesReferencingTargeturi              *[]ConceptMap            `bson:"_revIncludedConceptMapResourcesReferencingTargeturi,omitempty"`
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

func (v *ValueSetPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if v.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *v.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *v.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if v.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *v.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if v.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *v.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if v.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *v.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if v.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *v.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *v.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *v.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if v.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *v.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if v.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *v.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if v.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *v.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if v.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *v.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if v.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *v.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if v.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *v.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if v.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *v.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if v.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *v.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if v.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *v.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if v.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *v.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if v.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *v.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if v.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *v.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if v.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *v.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedStructureDefinitionResourcesReferencingValueset() (structureDefinitions []StructureDefinition, err error) {
	if v.RevIncludedStructureDefinitionResourcesReferencingValueset == nil {
		err = errors.New("RevIncluded structureDefinitions not requested")
	} else {
		structureDefinitions = *v.RevIncludedStructureDefinitionResourcesReferencingValueset
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if v.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *v.RevIncludedListResourcesReferencingItem
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedConceptMapResourcesReferencingSource() (conceptMaps []ConceptMap, err error) {
	if v.RevIncludedConceptMapResourcesReferencingSource == nil {
		err = errors.New("RevIncluded conceptMaps not requested")
	} else {
		conceptMaps = *v.RevIncludedConceptMapResourcesReferencingSource
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedConceptMapResourcesReferencingTarget() (conceptMaps []ConceptMap, err error) {
	if v.RevIncludedConceptMapResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded conceptMaps not requested")
	} else {
		conceptMaps = *v.RevIncludedConceptMapResourcesReferencingTarget
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedConceptMapResourcesReferencingSourceuri() (conceptMaps []ConceptMap, err error) {
	if v.RevIncludedConceptMapResourcesReferencingSourceuri == nil {
		err = errors.New("RevIncluded conceptMaps not requested")
	} else {
		conceptMaps = *v.RevIncludedConceptMapResourcesReferencingSourceuri
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedConceptMapResourcesReferencingTargeturi() (conceptMaps []ConceptMap, err error) {
	if v.RevIncludedConceptMapResourcesReferencingTargeturi == nil {
		err = errors.New("RevIncluded conceptMaps not requested")
	} else {
		conceptMaps = *v.RevIncludedConceptMapResourcesReferencingTargeturi
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if v.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *v.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if v.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *v.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if v.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *v.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if v.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *v.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if v.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *v.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if v.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *v.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if v.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *v.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if v.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *v.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *v.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if v.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *v.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *v.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*v.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *v.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*v.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*v.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedContractResourcesReferencingSubject {
			rsc := (*v.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *v.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*v.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *v.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*v.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *v.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*v.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*v.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*v.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*v.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*v.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedStructureDefinitionResourcesReferencingValueset != nil {
		for idx := range *v.RevIncludedStructureDefinitionResourcesReferencingValueset {
			rsc := (*v.RevIncludedStructureDefinitionResourcesReferencingValueset)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedListResourcesReferencingItem {
			rsc := (*v.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingSource != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingSource {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingTarget {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingSourceuri != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingSourceuri {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingSourceuri)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingTargeturi != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingTargeturi {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingTargeturi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *v.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*v.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*v.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*v.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *v.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*v.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *v.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*v.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*v.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*v.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *v.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*v.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*v.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *v.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*v.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (v *ValueSetPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *v.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*v.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *v.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*v.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*v.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedContractResourcesReferencingSubject {
			rsc := (*v.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *v.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*v.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *v.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*v.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *v.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*v.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*v.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*v.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*v.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*v.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedStructureDefinitionResourcesReferencingValueset != nil {
		for idx := range *v.RevIncludedStructureDefinitionResourcesReferencingValueset {
			rsc := (*v.RevIncludedStructureDefinitionResourcesReferencingValueset)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedListResourcesReferencingItem {
			rsc := (*v.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingSource != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingSource {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingTarget {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingSourceuri != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingSourceuri {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingSourceuri)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingTargeturi != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingTargeturi {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingTargeturi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *v.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*v.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*v.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*v.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *v.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*v.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *v.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*v.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*v.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*v.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *v.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*v.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*v.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *v.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*v.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
