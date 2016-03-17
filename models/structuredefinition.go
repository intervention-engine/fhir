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

type StructureDefinition struct {
	DomainResource  `bson:",inline"`
	Url             string                                    `bson:"url,omitempty" json:"url,omitempty"`
	Identifier      []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version         string                                    `bson:"version,omitempty" json:"version,omitempty"`
	Name            string                                    `bson:"name,omitempty" json:"name,omitempty"`
	Display         string                                    `bson:"display,omitempty" json:"display,omitempty"`
	Status          string                                    `bson:"status,omitempty" json:"status,omitempty"`
	Experimental    *bool                                     `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher       string                                    `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact         []StructureDefinitionContactComponent     `bson:"contact,omitempty" json:"contact,omitempty"`
	Date            *FHIRDateTime                             `bson:"date,omitempty" json:"date,omitempty"`
	Description     string                                    `bson:"description,omitempty" json:"description,omitempty"`
	UseContext      []CodeableConcept                         `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Requirements    string                                    `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright       string                                    `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Code            []Coding                                  `bson:"code,omitempty" json:"code,omitempty"`
	FhirVersion     string                                    `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Mapping         []StructureDefinitionMappingComponent     `bson:"mapping,omitempty" json:"mapping,omitempty"`
	Kind            string                                    `bson:"kind,omitempty" json:"kind,omitempty"`
	ConstrainedType string                                    `bson:"constrainedType,omitempty" json:"constrainedType,omitempty"`
	Abstract        *bool                                     `bson:"abstract,omitempty" json:"abstract,omitempty"`
	ContextType     string                                    `bson:"contextType,omitempty" json:"contextType,omitempty"`
	Context         []string                                  `bson:"context,omitempty" json:"context,omitempty"`
	Base            string                                    `bson:"base,omitempty" json:"base,omitempty"`
	Snapshot        *StructureDefinitionSnapshotComponent     `bson:"snapshot,omitempty" json:"snapshot,omitempty"`
	Differential    *StructureDefinitionDifferentialComponent `bson:"differential,omitempty" json:"differential,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *StructureDefinition) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "StructureDefinition"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to StructureDefinition), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *StructureDefinition) GetBSON() (interface{}, error) {
	x.ResourceType = "StructureDefinition"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "structureDefinition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type structureDefinition StructureDefinition

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *StructureDefinition) UnmarshalJSON(data []byte) (err error) {
	x2 := structureDefinition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = StructureDefinition(x2)
		return x.checkResourceType()
	}
	return
}

func (x *StructureDefinition) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "StructureDefinition"
	} else if x.ResourceType != "StructureDefinition" {
		return errors.New(fmt.Sprintf("Expected resourceType to be StructureDefinition, instead received %s", x.ResourceType))
	}
	return nil
}

type StructureDefinitionContactComponent struct {
	BackboneElement `bson:",inline"`
	Name            string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type StructureDefinitionMappingComponent struct {
	BackboneElement `bson:",inline"`
	Identity        string `bson:"identity,omitempty" json:"identity,omitempty"`
	Uri             string `bson:"uri,omitempty" json:"uri,omitempty"`
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	Comments        string `bson:"comments,omitempty" json:"comments,omitempty"`
}

type StructureDefinitionSnapshotComponent struct {
	BackboneElement `bson:",inline"`
	Element         []ElementDefinition `bson:"element,omitempty" json:"element,omitempty"`
}

type StructureDefinitionDifferentialComponent struct {
	BackboneElement `bson:",inline"`
	Element         []ElementDefinition `bson:"element,omitempty" json:"element,omitempty"`
}

type StructureDefinitionPlus struct {
	StructureDefinition                     `bson:",inline"`
	StructureDefinitionPlusRelatedResources `bson:",inline"`
}

type StructureDefinitionPlusRelatedResources struct {
	IncludedValueSetResourcesReferencedByValueset               *[]ValueSet              `bson:"_includedValueSetResourcesReferencedByValueset,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedConceptMapResourcesReferencingSource             *[]ConceptMap            `bson:"_revIncludedConceptMapResourcesReferencingSource,omitempty"`
	RevIncludedConceptMapResourcesReferencingTarget             *[]ConceptMap            `bson:"_revIncludedConceptMapResourcesReferencingTarget,omitempty"`
	RevIncludedConceptMapResourcesReferencingSourceuri          *[]ConceptMap            `bson:"_revIncludedConceptMapResourcesReferencingSourceuri,omitempty"`
	RevIncludedOperationDefinitionResourcesReferencingProfile   *[]OperationDefinition   `bson:"_revIncludedOperationDefinitionResourcesReferencingProfile,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedConformanceResourcesReferencingProfile           *[]Conformance           `bson:"_revIncludedConformanceResourcesReferencingProfile,omitempty"`
	RevIncludedConformanceResourcesReferencingSupportedprofile  *[]Conformance           `bson:"_revIncludedConformanceResourcesReferencingSupportedprofile,omitempty"`
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

func (s *StructureDefinitionPlusRelatedResources) GetIncludedValueSetResourceReferencedByValueset() (valueSet *ValueSet, err error) {
	if s.IncludedValueSetResourcesReferencedByValueset == nil {
		err = errors.New("Included valuesets not requested")
	} else if len(*s.IncludedValueSetResourcesReferencedByValueset) > 1 {
		err = fmt.Errorf("Expected 0 or 1 valueSet, but found %d", len(*s.IncludedValueSetResourcesReferencedByValueset))
	} else if len(*s.IncludedValueSetResourcesReferencedByValueset) == 1 {
		valueSet = &(*s.IncludedValueSetResourcesReferencedByValueset)[0]
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if s.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *s.RevIncludedListResourcesReferencingItem
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedConceptMapResourcesReferencingSource() (conceptMaps []ConceptMap, err error) {
	if s.RevIncludedConceptMapResourcesReferencingSource == nil {
		err = errors.New("RevIncluded conceptMaps not requested")
	} else {
		conceptMaps = *s.RevIncludedConceptMapResourcesReferencingSource
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedConceptMapResourcesReferencingTarget() (conceptMaps []ConceptMap, err error) {
	if s.RevIncludedConceptMapResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded conceptMaps not requested")
	} else {
		conceptMaps = *s.RevIncludedConceptMapResourcesReferencingTarget
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedConceptMapResourcesReferencingSourceuri() (conceptMaps []ConceptMap, err error) {
	if s.RevIncludedConceptMapResourcesReferencingSourceuri == nil {
		err = errors.New("RevIncluded conceptMaps not requested")
	} else {
		conceptMaps = *s.RevIncludedConceptMapResourcesReferencingSourceuri
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedOperationDefinitionResourcesReferencingProfile() (operationDefinitions []OperationDefinition, err error) {
	if s.RevIncludedOperationDefinitionResourcesReferencingProfile == nil {
		err = errors.New("RevIncluded operationDefinitions not requested")
	} else {
		operationDefinitions = *s.RevIncludedOperationDefinitionResourcesReferencingProfile
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if s.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *s.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedConformanceResourcesReferencingProfile() (conformances []Conformance, err error) {
	if s.RevIncludedConformanceResourcesReferencingProfile == nil {
		err = errors.New("RevIncluded conformances not requested")
	} else {
		conformances = *s.RevIncludedConformanceResourcesReferencingProfile
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedConformanceResourcesReferencingSupportedprofile() (conformances []Conformance, err error) {
	if s.RevIncludedConformanceResourcesReferencingSupportedprofile == nil {
		err = errors.New("RevIncluded conformances not requested")
	} else {
		conformances = *s.RevIncludedConformanceResourcesReferencingSupportedprofile
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if s.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *s.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if s.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *s.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *s.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *s.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if s.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *s.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *s.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if s.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *s.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (s *StructureDefinitionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedValueSetResourcesReferencedByValueset != nil {
		for idx := range *s.IncludedValueSetResourcesReferencedByValueset {
			rsc := (*s.IncludedValueSetResourcesReferencedByValueset)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *StructureDefinitionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConceptMapResourcesReferencingSource != nil {
		for idx := range *s.RevIncludedConceptMapResourcesReferencingSource {
			rsc := (*s.RevIncludedConceptMapResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConceptMapResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedConceptMapResourcesReferencingTarget {
			rsc := (*s.RevIncludedConceptMapResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConceptMapResourcesReferencingSourceuri != nil {
		for idx := range *s.RevIncludedConceptMapResourcesReferencingSourceuri {
			rsc := (*s.RevIncludedConceptMapResourcesReferencingSourceuri)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedOperationDefinitionResourcesReferencingProfile != nil {
		for idx := range *s.RevIncludedOperationDefinitionResourcesReferencingProfile {
			rsc := (*s.RevIncludedOperationDefinitionResourcesReferencingProfile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *s.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*s.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConformanceResourcesReferencingProfile != nil {
		for idx := range *s.RevIncludedConformanceResourcesReferencingProfile {
			rsc := (*s.RevIncludedConformanceResourcesReferencingProfile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConformanceResourcesReferencingSupportedprofile != nil {
		for idx := range *s.RevIncludedConformanceResourcesReferencingSupportedprofile {
			rsc := (*s.RevIncludedConformanceResourcesReferencingSupportedprofile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *s.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*s.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *StructureDefinitionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedValueSetResourcesReferencedByValueset != nil {
		for idx := range *s.IncludedValueSetResourcesReferencedByValueset {
			rsc := (*s.IncludedValueSetResourcesReferencedByValueset)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConceptMapResourcesReferencingSource != nil {
		for idx := range *s.RevIncludedConceptMapResourcesReferencingSource {
			rsc := (*s.RevIncludedConceptMapResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConceptMapResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedConceptMapResourcesReferencingTarget {
			rsc := (*s.RevIncludedConceptMapResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConceptMapResourcesReferencingSourceuri != nil {
		for idx := range *s.RevIncludedConceptMapResourcesReferencingSourceuri {
			rsc := (*s.RevIncludedConceptMapResourcesReferencingSourceuri)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedOperationDefinitionResourcesReferencingProfile != nil {
		for idx := range *s.RevIncludedOperationDefinitionResourcesReferencingProfile {
			rsc := (*s.RevIncludedOperationDefinitionResourcesReferencingProfile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *s.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*s.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConformanceResourcesReferencingProfile != nil {
		for idx := range *s.RevIncludedConformanceResourcesReferencingProfile {
			rsc := (*s.RevIncludedConformanceResourcesReferencingProfile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConformanceResourcesReferencingSupportedprofile != nil {
		for idx := range *s.RevIncludedConformanceResourcesReferencingSupportedprofile {
			rsc := (*s.RevIncludedConformanceResourcesReferencingSupportedprofile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *s.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*s.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
