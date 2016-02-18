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

type OperationDefinition struct {
	DomainResource `bson:",inline"`
	Url            string                                  `bson:"url,omitempty" json:"url,omitempty"`
	Version        string                                  `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                                  `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                                  `bson:"status,omitempty" json:"status,omitempty"`
	Kind           string                                  `bson:"kind,omitempty" json:"kind,omitempty"`
	Experimental   *bool                                   `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                                  `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []OperationDefinitionContactComponent   `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                           `bson:"date,omitempty" json:"date,omitempty"`
	Description    string                                  `bson:"description,omitempty" json:"description,omitempty"`
	Requirements   string                                  `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Idempotent     *bool                                   `bson:"idempotent,omitempty" json:"idempotent,omitempty"`
	Code           string                                  `bson:"code,omitempty" json:"code,omitempty"`
	Notes          string                                  `bson:"notes,omitempty" json:"notes,omitempty"`
	Base           *Reference                              `bson:"base,omitempty" json:"base,omitempty"`
	System         *bool                                   `bson:"system,omitempty" json:"system,omitempty"`
	Type           []string                                `bson:"type,omitempty" json:"type,omitempty"`
	Instance       *bool                                   `bson:"instance,omitempty" json:"instance,omitempty"`
	Parameter      []OperationDefinitionParameterComponent `bson:"parameter,omitempty" json:"parameter,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *OperationDefinition) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "OperationDefinition"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to OperationDefinition), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *OperationDefinition) GetBSON() (interface{}, error) {
	x.ResourceType = "OperationDefinition"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "operationDefinition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type operationDefinition OperationDefinition

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *OperationDefinition) UnmarshalJSON(data []byte) (err error) {
	x2 := operationDefinition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = OperationDefinition(x2)
		return x.checkResourceType()
	}
	return
}

func (x *OperationDefinition) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "OperationDefinition"
	} else if x.ResourceType != "OperationDefinition" {
		return errors.New(fmt.Sprintf("Expected resourceType to be OperationDefinition, instead received %s", x.ResourceType))
	}
	return nil
}

type OperationDefinitionContactComponent struct {
	BackboneElement `bson:",inline"`
	Name            string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type OperationDefinitionParameterComponent struct {
	BackboneElement `bson:",inline"`
	Name            string                                        `bson:"name,omitempty" json:"name,omitempty"`
	Use             string                                        `bson:"use,omitempty" json:"use,omitempty"`
	Min             *int32                                        `bson:"min,omitempty" json:"min,omitempty"`
	Max             string                                        `bson:"max,omitempty" json:"max,omitempty"`
	Documentation   string                                        `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Type            string                                        `bson:"type,omitempty" json:"type,omitempty"`
	Profile         *Reference                                    `bson:"profile,omitempty" json:"profile,omitempty"`
	Binding         *OperationDefinitionParameterBindingComponent `bson:"binding,omitempty" json:"binding,omitempty"`
	Part            []OperationDefinitionParameterComponent       `bson:"part,omitempty" json:"part,omitempty"`
}

type OperationDefinitionParameterBindingComponent struct {
	BackboneElement   `bson:",inline"`
	Strength          string     `bson:"strength,omitempty" json:"strength,omitempty"`
	ValueSetUri       string     `bson:"valueSetUri,omitempty" json:"valueSetUri,omitempty"`
	ValueSetReference *Reference `bson:"valueSetReference,omitempty" json:"valueSetReference,omitempty"`
}

type OperationDefinitionPlus struct {
	OperationDefinition                     `bson:",inline"`
	OperationDefinitionPlusRelatedResources `bson:",inline"`
}

type OperationDefinitionPlusRelatedResources struct {
	IncludedStructureDefinitionResourcesReferencedByProfile     *[]StructureDefinition   `bson:"_includedStructureDefinitionResourcesReferencedByProfile,omitempty"`
	IncludedOperationDefinitionResourcesReferencedByBase        *[]OperationDefinition   `bson:"_includedOperationDefinitionResourcesReferencedByBase,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedOperationDefinitionResourcesReferencingBase      *[]OperationDefinition   `bson:"_revIncludedOperationDefinitionResourcesReferencingBase,omitempty"`
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

func (o *OperationDefinitionPlusRelatedResources) GetIncludedStructureDefinitionResourceReferencedByProfile() (structureDefinition *StructureDefinition, err error) {
	if o.IncludedStructureDefinitionResourcesReferencedByProfile == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*o.IncludedStructureDefinitionResourcesReferencedByProfile) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*o.IncludedStructureDefinitionResourcesReferencedByProfile))
	} else if len(*o.IncludedStructureDefinitionResourcesReferencedByProfile) == 1 {
		structureDefinition = &(*o.IncludedStructureDefinitionResourcesReferencedByProfile)[0]
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetIncludedOperationDefinitionResourceReferencedByBase() (operationDefinition *OperationDefinition, err error) {
	if o.IncludedOperationDefinitionResourcesReferencedByBase == nil {
		err = errors.New("Included operationdefinitions not requested")
	} else if len(*o.IncludedOperationDefinitionResourcesReferencedByBase) > 1 {
		err = fmt.Errorf("Expected 0 or 1 operationDefinition, but found %d", len(*o.IncludedOperationDefinitionResourcesReferencedByBase))
	} else if len(*o.IncludedOperationDefinitionResourcesReferencedByBase) == 1 {
		operationDefinition = &(*o.IncludedOperationDefinitionResourcesReferencedByBase)[0]
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if o.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *o.RevIncludedListResourcesReferencingItem
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedOperationDefinitionResourcesReferencingBase() (operationDefinitions []OperationDefinition, err error) {
	if o.RevIncludedOperationDefinitionResourcesReferencingBase == nil {
		err = errors.New("RevIncluded operationDefinitions not requested")
	} else {
		operationDefinitions = *o.RevIncludedOperationDefinitionResourcesReferencingBase
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if o.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *o.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *o.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *o.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedStructureDefinitionResourcesReferencedByProfile != nil {
		for _, r := range *o.IncludedStructureDefinitionResourcesReferencedByProfile {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedOperationDefinitionResourcesReferencedByBase != nil {
		for _, r := range *o.IncludedOperationDefinitionResourcesReferencedByBase {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *o.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOperationDefinitionResourcesReferencingBase != nil {
		for _, r := range *o.RevIncludedOperationDefinitionResourcesReferencingBase {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *o.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *o.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *o.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *o.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (o *OperationDefinitionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedStructureDefinitionResourcesReferencedByProfile != nil {
		for _, r := range *o.IncludedStructureDefinitionResourcesReferencedByProfile {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedOperationDefinitionResourcesReferencedByBase != nil {
		for _, r := range *o.IncludedOperationDefinitionResourcesReferencedByBase {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *o.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOperationDefinitionResourcesReferencingBase != nil {
		for _, r := range *o.RevIncludedOperationDefinitionResourcesReferencingBase {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *o.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *o.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *o.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *o.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
