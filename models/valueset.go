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

type ValueSet struct {
	DomainResource `bson:",inline"`
	Url            string                       `bson:"url,omitempty" json:"url,omitempty"`
	Identifier     *Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version        string                       `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                       `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                       `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                        `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                       `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ValueSetContactComponent   `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                `bson:"date,omitempty" json:"date,omitempty"`
	LockedDate     *FHIRDateTime                `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	Description    string                       `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []CodeableConcept            `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Immutable      *bool                        `bson:"immutable,omitempty" json:"immutable,omitempty"`
	Requirements   string                       `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright      string                       `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Extensible     *bool                        `bson:"extensible,omitempty" json:"extensible,omitempty"`
	CodeSystem     *ValueSetCodeSystemComponent `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
	Compose        *ValueSetComposeComponent    `bson:"compose,omitempty" json:"compose,omitempty"`
	Expansion      *ValueSetExpansionComponent  `bson:"expansion,omitempty" json:"expansion,omitempty"`
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
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
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

type ValueSetContactComponent struct {
	BackboneElement `bson:",inline"`
	Name            string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type ValueSetCodeSystemComponent struct {
	BackboneElement `bson:",inline"`
	System          string                               `bson:"system,omitempty" json:"system,omitempty"`
	Version         string                               `bson:"version,omitempty" json:"version,omitempty"`
	CaseSensitive   *bool                                `bson:"caseSensitive,omitempty" json:"caseSensitive,omitempty"`
	Concept         []ValueSetConceptDefinitionComponent `bson:"concept,omitempty" json:"concept,omitempty"`
}

type ValueSetConceptDefinitionComponent struct {
	BackboneElement `bson:",inline"`
	Code            string                                          `bson:"code,omitempty" json:"code,omitempty"`
	Abstract        *bool                                           `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Display         string                                          `bson:"display,omitempty" json:"display,omitempty"`
	Definition      string                                          `bson:"definition,omitempty" json:"definition,omitempty"`
	Designation     []ValueSetConceptDefinitionDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
	Concept         []ValueSetConceptDefinitionComponent            `bson:"concept,omitempty" json:"concept,omitempty"`
}

type ValueSetConceptDefinitionDesignationComponent struct {
	BackboneElement `bson:",inline"`
	Language        string  `bson:"language,omitempty" json:"language,omitempty"`
	Use             *Coding `bson:"use,omitempty" json:"use,omitempty"`
	Value           string  `bson:"value,omitempty" json:"value,omitempty"`
}

type ValueSetComposeComponent struct {
	BackboneElement `bson:",inline"`
	Import          []string                      `bson:"import,omitempty" json:"import,omitempty"`
	Include         []ValueSetConceptSetComponent `bson:"include,omitempty" json:"include,omitempty"`
	Exclude         []ValueSetConceptSetComponent `bson:"exclude,omitempty" json:"exclude,omitempty"`
}

type ValueSetConceptSetComponent struct {
	BackboneElement `bson:",inline"`
	System          string                              `bson:"system,omitempty" json:"system,omitempty"`
	Version         string                              `bson:"version,omitempty" json:"version,omitempty"`
	Concept         []ValueSetConceptReferenceComponent `bson:"concept,omitempty" json:"concept,omitempty"`
	Filter          []ValueSetConceptSetFilterComponent `bson:"filter,omitempty" json:"filter,omitempty"`
}

type ValueSetConceptReferenceComponent struct {
	BackboneElement `bson:",inline"`
	Code            string                                          `bson:"code,omitempty" json:"code,omitempty"`
	Display         string                                          `bson:"display,omitempty" json:"display,omitempty"`
	Designation     []ValueSetConceptDefinitionDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
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
	System          string                               `bson:"system,omitempty" json:"system,omitempty"`
	Abstract        *bool                                `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Version         string                               `bson:"version,omitempty" json:"version,omitempty"`
	Code            string                               `bson:"code,omitempty" json:"code,omitempty"`
	Display         string                               `bson:"display,omitempty" json:"display,omitempty"`
	Contains        []ValueSetExpansionContainsComponent `bson:"contains,omitempty" json:"contains,omitempty"`
}

type ValueSetPlus struct {
	ValueSet                     `bson:",inline"`
	ValueSetPlusRelatedResources `bson:",inline"`
}

type ValueSetPlusRelatedResources struct {
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedStructureDefinitionResourcesReferencingValueset  *[]StructureDefinition   `bson:"_revIncludedStructureDefinitionResourcesReferencingValueset,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedConceptMapResourcesReferencingSource             *[]ConceptMap            `bson:"_revIncludedConceptMapResourcesReferencingSource,omitempty"`
	RevIncludedConceptMapResourcesReferencingTarget             *[]ConceptMap            `bson:"_revIncludedConceptMapResourcesReferencingTarget,omitempty"`
	RevIncludedConceptMapResourcesReferencingSourceuri          *[]ConceptMap            `bson:"_revIncludedConceptMapResourcesReferencingSourceuri,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if v.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *v.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
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

func (v *ValueSetPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if v.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *v.RevIncludedOrderResourcesReferencingDetail
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

func (v *ValueSetPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if v.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *v.RevIncludedAuditEventResourcesReferencingReference
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

func (v *ValueSetPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if v.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *v.RevIncludedOrderResponseResourcesReferencingFulfillment
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

func (v *ValueSetPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if v.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *v.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if v.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *v.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (v *ValueSetPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (v *ValueSetPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
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
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *v.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*v.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*v.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *v.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*v.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if v.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *v.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*v.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if v.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *v.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*v.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *v.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*v.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (v *ValueSetPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
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
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *v.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*v.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*v.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *v.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*v.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if v.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *v.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*v.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if v.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *v.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*v.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *v.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*v.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
