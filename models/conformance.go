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

type Conformance struct {
	DomainResource `bson:",inline"`
	Url            string                              `bson:"url,omitempty" json:"url,omitempty"`
	Version        string                              `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                              `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                              `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                               `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                              `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ConformanceContactComponent       `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                       `bson:"date,omitempty" json:"date,omitempty"`
	Description    string                              `bson:"description,omitempty" json:"description,omitempty"`
	Requirements   string                              `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright      string                              `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Kind           string                              `bson:"kind,omitempty" json:"kind,omitempty"`
	Software       *ConformanceSoftwareComponent       `bson:"software,omitempty" json:"software,omitempty"`
	Implementation *ConformanceImplementationComponent `bson:"implementation,omitempty" json:"implementation,omitempty"`
	FhirVersion    string                              `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	AcceptUnknown  string                              `bson:"acceptUnknown,omitempty" json:"acceptUnknown,omitempty"`
	Format         []string                            `bson:"format,omitempty" json:"format,omitempty"`
	Profile        []Reference                         `bson:"profile,omitempty" json:"profile,omitempty"`
	Rest           []ConformanceRestComponent          `bson:"rest,omitempty" json:"rest,omitempty"`
	Messaging      []ConformanceMessagingComponent     `bson:"messaging,omitempty" json:"messaging,omitempty"`
	Document       []ConformanceDocumentComponent      `bson:"document,omitempty" json:"document,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Conformance) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Conformance"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Conformance), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Conformance) GetBSON() (interface{}, error) {
	x.ResourceType = "Conformance"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "conformance" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type conformance Conformance

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Conformance) UnmarshalJSON(data []byte) (err error) {
	x2 := conformance{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Conformance(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Conformance) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Conformance"
	} else if x.ResourceType != "Conformance" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Conformance, instead received %s", x.ResourceType))
	}
	return nil
}

type ConformanceContactComponent struct {
	BackboneElement `bson:",inline"`
	Name            string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type ConformanceSoftwareComponent struct {
	BackboneElement `bson:",inline"`
	Name            string        `bson:"name,omitempty" json:"name,omitempty"`
	Version         string        `bson:"version,omitempty" json:"version,omitempty"`
	ReleaseDate     *FHIRDateTime `bson:"releaseDate,omitempty" json:"releaseDate,omitempty"`
}

type ConformanceImplementationComponent struct {
	BackboneElement `bson:",inline"`
	Description     string `bson:"description,omitempty" json:"description,omitempty"`
	Url             string `bson:"url,omitempty" json:"url,omitempty"`
}

type ConformanceRestComponent struct {
	BackboneElement `bson:",inline"`
	Mode            string                                        `bson:"mode,omitempty" json:"mode,omitempty"`
	Documentation   string                                        `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Security        *ConformanceRestSecurityComponent             `bson:"security,omitempty" json:"security,omitempty"`
	Resource        []ConformanceRestResourceComponent            `bson:"resource,omitempty" json:"resource,omitempty"`
	Interaction     []ConformanceSystemInteractionComponent       `bson:"interaction,omitempty" json:"interaction,omitempty"`
	TransactionMode string                                        `bson:"transactionMode,omitempty" json:"transactionMode,omitempty"`
	SearchParam     []ConformanceRestResourceSearchParamComponent `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	Operation       []ConformanceRestOperationComponent           `bson:"operation,omitempty" json:"operation,omitempty"`
	Compartment     []string                                      `bson:"compartment,omitempty" json:"compartment,omitempty"`
}

type ConformanceRestSecurityComponent struct {
	BackboneElement `bson:",inline"`
	Cors            *bool                                         `bson:"cors,omitempty" json:"cors,omitempty"`
	Service         []CodeableConcept                             `bson:"service,omitempty" json:"service,omitempty"`
	Description     string                                        `bson:"description,omitempty" json:"description,omitempty"`
	Certificate     []ConformanceRestSecurityCertificateComponent `bson:"certificate,omitempty" json:"certificate,omitempty"`
}

type ConformanceRestSecurityCertificateComponent struct {
	BackboneElement `bson:",inline"`
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
	Blob            string `bson:"blob,omitempty" json:"blob,omitempty"`
}

type ConformanceRestResourceComponent struct {
	BackboneElement   `bson:",inline"`
	Type              string                                        `bson:"type,omitempty" json:"type,omitempty"`
	Profile           *Reference                                    `bson:"profile,omitempty" json:"profile,omitempty"`
	Interaction       []ConformanceResourceInteractionComponent     `bson:"interaction,omitempty" json:"interaction,omitempty"`
	Versioning        string                                        `bson:"versioning,omitempty" json:"versioning,omitempty"`
	ReadHistory       *bool                                         `bson:"readHistory,omitempty" json:"readHistory,omitempty"`
	UpdateCreate      *bool                                         `bson:"updateCreate,omitempty" json:"updateCreate,omitempty"`
	ConditionalCreate *bool                                         `bson:"conditionalCreate,omitempty" json:"conditionalCreate,omitempty"`
	ConditionalUpdate *bool                                         `bson:"conditionalUpdate,omitempty" json:"conditionalUpdate,omitempty"`
	ConditionalDelete string                                        `bson:"conditionalDelete,omitempty" json:"conditionalDelete,omitempty"`
	SearchInclude     []string                                      `bson:"searchInclude,omitempty" json:"searchInclude,omitempty"`
	SearchRevInclude  []string                                      `bson:"searchRevInclude,omitempty" json:"searchRevInclude,omitempty"`
	SearchParam       []ConformanceRestResourceSearchParamComponent `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
}

type ConformanceResourceInteractionComponent struct {
	BackboneElement `bson:",inline"`
	Code            string `bson:"code,omitempty" json:"code,omitempty"`
	Documentation   string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type ConformanceRestResourceSearchParamComponent struct {
	BackboneElement `bson:",inline"`
	Name            string   `bson:"name,omitempty" json:"name,omitempty"`
	Definition      string   `bson:"definition,omitempty" json:"definition,omitempty"`
	Type            string   `bson:"type,omitempty" json:"type,omitempty"`
	Documentation   string   `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Target          []string `bson:"target,omitempty" json:"target,omitempty"`
	Modifier        []string `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Chain           []string `bson:"chain,omitempty" json:"chain,omitempty"`
}

type ConformanceSystemInteractionComponent struct {
	BackboneElement `bson:",inline"`
	Code            string `bson:"code,omitempty" json:"code,omitempty"`
	Documentation   string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type ConformanceRestOperationComponent struct {
	BackboneElement `bson:",inline"`
	Name            string     `bson:"name,omitempty" json:"name,omitempty"`
	Definition      *Reference `bson:"definition,omitempty" json:"definition,omitempty"`
}

type ConformanceMessagingComponent struct {
	BackboneElement `bson:",inline"`
	Endpoint        []ConformanceMessagingEndpointComponent `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	ReliableCache   *uint32                                 `bson:"reliableCache,omitempty" json:"reliableCache,omitempty"`
	Documentation   string                                  `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Event           []ConformanceMessagingEventComponent    `bson:"event,omitempty" json:"event,omitempty"`
}

type ConformanceMessagingEndpointComponent struct {
	BackboneElement `bson:",inline"`
	Protocol        *Coding `bson:"protocol,omitempty" json:"protocol,omitempty"`
	Address         string  `bson:"address,omitempty" json:"address,omitempty"`
}

type ConformanceMessagingEventComponent struct {
	BackboneElement `bson:",inline"`
	Code            *Coding    `bson:"code,omitempty" json:"code,omitempty"`
	Category        string     `bson:"category,omitempty" json:"category,omitempty"`
	Mode            string     `bson:"mode,omitempty" json:"mode,omitempty"`
	Focus           string     `bson:"focus,omitempty" json:"focus,omitempty"`
	Request         *Reference `bson:"request,omitempty" json:"request,omitempty"`
	Response        *Reference `bson:"response,omitempty" json:"response,omitempty"`
	Documentation   string     `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type ConformanceDocumentComponent struct {
	BackboneElement `bson:",inline"`
	Mode            string     `bson:"mode,omitempty" json:"mode,omitempty"`
	Documentation   string     `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Profile         *Reference `bson:"profile,omitempty" json:"profile,omitempty"`
}

type ConformancePlus struct {
	Conformance                     `bson:",inline"`
	ConformancePlusRelatedResources `bson:",inline"`
}

type ConformancePlusRelatedResources struct {
	IncludedStructureDefinitionResourcesReferencedByProfile          *[]StructureDefinition   `bson:"_includedStructureDefinitionResourcesReferencedByProfile,omitempty"`
	IncludedStructureDefinitionResourcesReferencedBySupportedprofile *[]StructureDefinition   `bson:"_includedStructureDefinitionResourcesReferencedBySupportedprofile,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                  *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref        *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref        *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                          *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref       *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                       *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                      *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference               *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                  *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated           *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment          *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject      *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest            *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger         *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                 *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (c *ConformancePlusRelatedResources) GetIncludedStructureDefinitionResourceReferencedByProfile() (structureDefinition *StructureDefinition, err error) {
	if c.IncludedStructureDefinitionResourcesReferencedByProfile == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*c.IncludedStructureDefinitionResourcesReferencedByProfile) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*c.IncludedStructureDefinitionResourcesReferencedByProfile))
	} else if len(*c.IncludedStructureDefinitionResourcesReferencedByProfile) == 1 {
		structureDefinition = &(*c.IncludedStructureDefinitionResourcesReferencedByProfile)[0]
	}
	return
}

func (c *ConformancePlusRelatedResources) GetIncludedStructureDefinitionResourcesReferencedBySupportedprofile() (structureDefinitions []StructureDefinition, err error) {
	if c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile == nil {
		err = errors.New("Included structureDefinitions not requested")
	} else {
		structureDefinitions = *c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (c *ConformancePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *ConformancePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedStructureDefinitionResourcesReferencedByProfile != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedByProfile {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedByProfile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ConformancePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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

func (c *ConformancePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedStructureDefinitionResourcesReferencedByProfile != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedByProfile {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedByProfile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile)[idx]
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
