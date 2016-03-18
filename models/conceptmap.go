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

type ConceptMap struct {
	DomainResource  `bson:",inline"`
	Url             string                             `bson:"url,omitempty" json:"url,omitempty"`
	Identifier      *Identifier                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version         string                             `bson:"version,omitempty" json:"version,omitempty"`
	Name            string                             `bson:"name,omitempty" json:"name,omitempty"`
	Status          string                             `bson:"status,omitempty" json:"status,omitempty"`
	Experimental    *bool                              `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher       string                             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact         []ConceptMapContactComponent       `bson:"contact,omitempty" json:"contact,omitempty"`
	Date            *FHIRDateTime                      `bson:"date,omitempty" json:"date,omitempty"`
	Description     string                             `bson:"description,omitempty" json:"description,omitempty"`
	UseContext      []CodeableConcept                  `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Requirements    string                             `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright       string                             `bson:"copyright,omitempty" json:"copyright,omitempty"`
	SourceUri       string                             `bson:"sourceUri,omitempty" json:"sourceUri,omitempty"`
	SourceReference *Reference                         `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	TargetUri       string                             `bson:"targetUri,omitempty" json:"targetUri,omitempty"`
	TargetReference *Reference                         `bson:"targetReference,omitempty" json:"targetReference,omitempty"`
	Element         []ConceptMapSourceElementComponent `bson:"element,omitempty" json:"element,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ConceptMap) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ConceptMap"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ConceptMap), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ConceptMap) GetBSON() (interface{}, error) {
	x.ResourceType = "ConceptMap"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "conceptMap" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type conceptMap ConceptMap

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ConceptMap) UnmarshalJSON(data []byte) (err error) {
	x2 := conceptMap{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ConceptMap(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ConceptMap) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ConceptMap"
	} else if x.ResourceType != "ConceptMap" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ConceptMap, instead received %s", x.ResourceType))
	}
	return nil
}

type ConceptMapContactComponent struct {
	BackboneElement `bson:",inline"`
	Name            string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type ConceptMapSourceElementComponent struct {
	BackboneElement `bson:",inline"`
	CodeSystem      string                             `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
	Code            string                             `bson:"code,omitempty" json:"code,omitempty"`
	Target          []ConceptMapTargetElementComponent `bson:"target,omitempty" json:"target,omitempty"`
}

type ConceptMapTargetElementComponent struct {
	BackboneElement `bson:",inline"`
	CodeSystem      string                            `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
	Code            string                            `bson:"code,omitempty" json:"code,omitempty"`
	Equivalence     string                            `bson:"equivalence,omitempty" json:"equivalence,omitempty"`
	Comments        string                            `bson:"comments,omitempty" json:"comments,omitempty"`
	DependsOn       []ConceptMapOtherElementComponent `bson:"dependsOn,omitempty" json:"dependsOn,omitempty"`
	Product         []ConceptMapOtherElementComponent `bson:"product,omitempty" json:"product,omitempty"`
}

type ConceptMapOtherElementComponent struct {
	BackboneElement `bson:",inline"`
	Element         string `bson:"element,omitempty" json:"element,omitempty"`
	CodeSystem      string `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
	Code            string `bson:"code,omitempty" json:"code,omitempty"`
}

type ConceptMapPlus struct {
	ConceptMap                     `bson:",inline"`
	ConceptMapPlusRelatedResources `bson:",inline"`
}

type ConceptMapPlusRelatedResources struct {
	IncludedStructureDefinitionResourcesReferencedBySource      *[]StructureDefinition   `bson:"_includedStructureDefinitionResourcesReferencedBySource,omitempty"`
	IncludedValueSetResourcesReferencedBySource                 *[]ValueSet              `bson:"_includedValueSetResourcesReferencedBySource,omitempty"`
	IncludedStructureDefinitionResourcesReferencedByTarget      *[]StructureDefinition   `bson:"_includedStructureDefinitionResourcesReferencedByTarget,omitempty"`
	IncludedValueSetResourcesReferencedByTarget                 *[]ValueSet              `bson:"_includedValueSetResourcesReferencedByTarget,omitempty"`
	IncludedStructureDefinitionResourcesReferencedBySourceuri   *[]StructureDefinition   `bson:"_includedStructureDefinitionResourcesReferencedBySourceuri,omitempty"`
	IncludedValueSetResourcesReferencedBySourceuri              *[]ValueSet              `bson:"_includedValueSetResourcesReferencedBySourceuri,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
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

func (c *ConceptMapPlusRelatedResources) GetIncludedStructureDefinitionResourceReferencedBySource() (structureDefinition *StructureDefinition, err error) {
	if c.IncludedStructureDefinitionResourcesReferencedBySource == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*c.IncludedStructureDefinitionResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*c.IncludedStructureDefinitionResourcesReferencedBySource))
	} else if len(*c.IncludedStructureDefinitionResourcesReferencedBySource) == 1 {
		structureDefinition = &(*c.IncludedStructureDefinitionResourcesReferencedBySource)[0]
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetIncludedValueSetResourceReferencedBySource() (valueSet *ValueSet, err error) {
	if c.IncludedValueSetResourcesReferencedBySource == nil {
		err = errors.New("Included valuesets not requested")
	} else if len(*c.IncludedValueSetResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 valueSet, but found %d", len(*c.IncludedValueSetResourcesReferencedBySource))
	} else if len(*c.IncludedValueSetResourcesReferencedBySource) == 1 {
		valueSet = &(*c.IncludedValueSetResourcesReferencedBySource)[0]
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetIncludedStructureDefinitionResourceReferencedByTarget() (structureDefinition *StructureDefinition, err error) {
	if c.IncludedStructureDefinitionResourcesReferencedByTarget == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*c.IncludedStructureDefinitionResourcesReferencedByTarget) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*c.IncludedStructureDefinitionResourcesReferencedByTarget))
	} else if len(*c.IncludedStructureDefinitionResourcesReferencedByTarget) == 1 {
		structureDefinition = &(*c.IncludedStructureDefinitionResourcesReferencedByTarget)[0]
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetIncludedValueSetResourceReferencedByTarget() (valueSet *ValueSet, err error) {
	if c.IncludedValueSetResourcesReferencedByTarget == nil {
		err = errors.New("Included valuesets not requested")
	} else if len(*c.IncludedValueSetResourcesReferencedByTarget) > 1 {
		err = fmt.Errorf("Expected 0 or 1 valueSet, but found %d", len(*c.IncludedValueSetResourcesReferencedByTarget))
	} else if len(*c.IncludedValueSetResourcesReferencedByTarget) == 1 {
		valueSet = &(*c.IncludedValueSetResourcesReferencedByTarget)[0]
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetIncludedStructureDefinitionResourceReferencedBySourceuri() (structureDefinition *StructureDefinition, err error) {
	if c.IncludedStructureDefinitionResourcesReferencedBySourceuri == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*c.IncludedStructureDefinitionResourcesReferencedBySourceuri) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*c.IncludedStructureDefinitionResourcesReferencedBySourceuri))
	} else if len(*c.IncludedStructureDefinitionResourcesReferencedBySourceuri) == 1 {
		structureDefinition = &(*c.IncludedStructureDefinitionResourcesReferencedBySourceuri)[0]
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetIncludedValueSetResourceReferencedBySourceuri() (valueSet *ValueSet, err error) {
	if c.IncludedValueSetResourcesReferencedBySourceuri == nil {
		err = errors.New("Included valuesets not requested")
	} else if len(*c.IncludedValueSetResourcesReferencedBySourceuri) > 1 {
		err = fmt.Errorf("Expected 0 or 1 valueSet, but found %d", len(*c.IncludedValueSetResourcesReferencedBySourceuri))
	} else if len(*c.IncludedValueSetResourcesReferencedBySourceuri) == 1 {
		valueSet = &(*c.IncludedValueSetResourcesReferencedBySourceuri)[0]
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *ConceptMapPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedStructureDefinitionResourcesReferencedBySource != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedBySource {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedValueSetResourcesReferencedBySource != nil {
		for idx := range *c.IncludedValueSetResourcesReferencedBySource {
			rsc := (*c.IncludedValueSetResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedStructureDefinitionResourcesReferencedByTarget != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedByTarget {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedByTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedValueSetResourcesReferencedByTarget != nil {
		for idx := range *c.IncludedValueSetResourcesReferencedByTarget {
			rsc := (*c.IncludedValueSetResourcesReferencedByTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedStructureDefinitionResourcesReferencedBySourceuri != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedBySourceuri {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedBySourceuri)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedValueSetResourcesReferencedBySourceuri != nil {
		for idx := range *c.IncludedValueSetResourcesReferencedBySourceuri {
			rsc := (*c.IncludedValueSetResourcesReferencedBySourceuri)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ConceptMapPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *c.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*c.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*c.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ConceptMapPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedStructureDefinitionResourcesReferencedBySource != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedBySource {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedValueSetResourcesReferencedBySource != nil {
		for idx := range *c.IncludedValueSetResourcesReferencedBySource {
			rsc := (*c.IncludedValueSetResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedStructureDefinitionResourcesReferencedByTarget != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedByTarget {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedByTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedValueSetResourcesReferencedByTarget != nil {
		for idx := range *c.IncludedValueSetResourcesReferencedByTarget {
			rsc := (*c.IncludedValueSetResourcesReferencedByTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedStructureDefinitionResourcesReferencedBySourceuri != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedBySourceuri {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedBySourceuri)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedValueSetResourcesReferencedBySourceuri != nil {
		for idx := range *c.IncludedValueSetResourcesReferencedBySourceuri {
			rsc := (*c.IncludedValueSetResourcesReferencedBySourceuri)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *c.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*c.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*c.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
