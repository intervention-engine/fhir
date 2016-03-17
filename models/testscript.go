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

type TestScript struct {
	DomainResource `bson:",inline"`
	Url            string                        `bson:"url,omitempty" json:"url,omitempty"`
	Version        string                        `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                        `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                        `bson:"status,omitempty" json:"status,omitempty"`
	Identifier     *Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Experimental   *bool                         `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                        `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []TestScriptContactComponent  `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                 `bson:"date,omitempty" json:"date,omitempty"`
	Description    string                        `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []CodeableConcept             `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Requirements   string                        `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright      string                        `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Metadata       *TestScriptMetadataComponent  `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Multiserver    *bool                         `bson:"multiserver,omitempty" json:"multiserver,omitempty"`
	Fixture        []TestScriptFixtureComponent  `bson:"fixture,omitempty" json:"fixture,omitempty"`
	Profile        []Reference                   `bson:"profile,omitempty" json:"profile,omitempty"`
	Variable       []TestScriptVariableComponent `bson:"variable,omitempty" json:"variable,omitempty"`
	Setup          *TestScriptSetupComponent     `bson:"setup,omitempty" json:"setup,omitempty"`
	Test           []TestScriptTestComponent     `bson:"test,omitempty" json:"test,omitempty"`
	Teardown       *TestScriptTeardownComponent  `bson:"teardown,omitempty" json:"teardown,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *TestScript) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "TestScript"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to TestScript), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *TestScript) GetBSON() (interface{}, error) {
	x.ResourceType = "TestScript"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "testScript" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type testScript TestScript

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *TestScript) UnmarshalJSON(data []byte) (err error) {
	x2 := testScript{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = TestScript(x2)
		return x.checkResourceType()
	}
	return
}

func (x *TestScript) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "TestScript"
	} else if x.ResourceType != "TestScript" {
		return errors.New(fmt.Sprintf("Expected resourceType to be TestScript, instead received %s", x.ResourceType))
	}
	return nil
}

type TestScriptContactComponent struct {
	BackboneElement `bson:",inline"`
	Name            string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type TestScriptMetadataComponent struct {
	BackboneElement `bson:",inline"`
	Link            []TestScriptMetadataLinkComponent       `bson:"link,omitempty" json:"link,omitempty"`
	Capability      []TestScriptMetadataCapabilityComponent `bson:"capability,omitempty" json:"capability,omitempty"`
}

type TestScriptMetadataLinkComponent struct {
	BackboneElement `bson:",inline"`
	Url             string `bson:"url,omitempty" json:"url,omitempty"`
	Description     string `bson:"description,omitempty" json:"description,omitempty"`
}

type TestScriptMetadataCapabilityComponent struct {
	BackboneElement `bson:",inline"`
	Required        *bool      `bson:"required,omitempty" json:"required,omitempty"`
	Validated       *bool      `bson:"validated,omitempty" json:"validated,omitempty"`
	Description     string     `bson:"description,omitempty" json:"description,omitempty"`
	Destination     *int32     `bson:"destination,omitempty" json:"destination,omitempty"`
	Link            []string   `bson:"link,omitempty" json:"link,omitempty"`
	Conformance     *Reference `bson:"conformance,omitempty" json:"conformance,omitempty"`
}

type TestScriptFixtureComponent struct {
	BackboneElement `bson:",inline"`
	Autocreate      *bool      `bson:"autocreate,omitempty" json:"autocreate,omitempty"`
	Autodelete      *bool      `bson:"autodelete,omitempty" json:"autodelete,omitempty"`
	Resource        *Reference `bson:"resource,omitempty" json:"resource,omitempty"`
}

type TestScriptVariableComponent struct {
	BackboneElement `bson:",inline"`
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	HeaderField     string `bson:"headerField,omitempty" json:"headerField,omitempty"`
	Path            string `bson:"path,omitempty" json:"path,omitempty"`
	SourceId        string `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
}

type TestScriptSetupComponent struct {
	BackboneElement `bson:",inline"`
	Metadata        *TestScriptMetadataComponent     `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Action          []TestScriptSetupActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestScriptSetupActionComponent struct {
	BackboneElement `bson:",inline"`
	Operation       *TestScriptSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert          *TestScriptSetupActionAssertComponent    `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestScriptSetupActionOperationComponent struct {
	BackboneElement  `bson:",inline"`
	Type             *Coding                                                `bson:"type,omitempty" json:"type,omitempty"`
	Resource         string                                                 `bson:"resource,omitempty" json:"resource,omitempty"`
	Label            string                                                 `bson:"label,omitempty" json:"label,omitempty"`
	Description      string                                                 `bson:"description,omitempty" json:"description,omitempty"`
	Accept           string                                                 `bson:"accept,omitempty" json:"accept,omitempty"`
	ContentType      string                                                 `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Destination      *int32                                                 `bson:"destination,omitempty" json:"destination,omitempty"`
	EncodeRequestUrl *bool                                                  `bson:"encodeRequestUrl,omitempty" json:"encodeRequestUrl,omitempty"`
	Params           string                                                 `bson:"params,omitempty" json:"params,omitempty"`
	RequestHeader    []TestScriptSetupActionOperationRequestHeaderComponent `bson:"requestHeader,omitempty" json:"requestHeader,omitempty"`
	ResponseId       string                                                 `bson:"responseId,omitempty" json:"responseId,omitempty"`
	SourceId         string                                                 `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	TargetId         string                                                 `bson:"targetId,omitempty" json:"targetId,omitempty"`
	Url              string                                                 `bson:"url,omitempty" json:"url,omitempty"`
}

type TestScriptSetupActionOperationRequestHeaderComponent struct {
	BackboneElement `bson:",inline"`
	Field           string `bson:"field,omitempty" json:"field,omitempty"`
	Value           string `bson:"value,omitempty" json:"value,omitempty"`
}

type TestScriptSetupActionAssertComponent struct {
	BackboneElement     `bson:",inline"`
	Label               string `bson:"label,omitempty" json:"label,omitempty"`
	Description         string `bson:"description,omitempty" json:"description,omitempty"`
	Direction           string `bson:"direction,omitempty" json:"direction,omitempty"`
	CompareToSourceId   string `bson:"compareToSourceId,omitempty" json:"compareToSourceId,omitempty"`
	CompareToSourcePath string `bson:"compareToSourcePath,omitempty" json:"compareToSourcePath,omitempty"`
	ContentType         string `bson:"contentType,omitempty" json:"contentType,omitempty"`
	HeaderField         string `bson:"headerField,omitempty" json:"headerField,omitempty"`
	MinimumId           string `bson:"minimumId,omitempty" json:"minimumId,omitempty"`
	NavigationLinks     *bool  `bson:"navigationLinks,omitempty" json:"navigationLinks,omitempty"`
	Operator            string `bson:"operator,omitempty" json:"operator,omitempty"`
	Path                string `bson:"path,omitempty" json:"path,omitempty"`
	Resource            string `bson:"resource,omitempty" json:"resource,omitempty"`
	Response            string `bson:"response,omitempty" json:"response,omitempty"`
	ResponseCode        string `bson:"responseCode,omitempty" json:"responseCode,omitempty"`
	SourceId            string `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	ValidateProfileId   string `bson:"validateProfileId,omitempty" json:"validateProfileId,omitempty"`
	Value               string `bson:"value,omitempty" json:"value,omitempty"`
	WarningOnly         *bool  `bson:"warningOnly,omitempty" json:"warningOnly,omitempty"`
}

type TestScriptTestComponent struct {
	BackboneElement `bson:",inline"`
	Name            string                          `bson:"name,omitempty" json:"name,omitempty"`
	Description     string                          `bson:"description,omitempty" json:"description,omitempty"`
	Metadata        *TestScriptMetadataComponent    `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Action          []TestScriptTestActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestScriptTestActionComponent struct {
	BackboneElement `bson:",inline"`
	Operation       *TestScriptSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert          *TestScriptSetupActionAssertComponent    `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestScriptTeardownComponent struct {
	BackboneElement `bson:",inline"`
	Action          []TestScriptTeardownActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestScriptTeardownActionComponent struct {
	BackboneElement `bson:",inline"`
	Operation       *TestScriptSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
}

type TestScriptPlus struct {
	TestScript                     `bson:",inline"`
	TestScriptPlusRelatedResources `bson:",inline"`
}

type TestScriptPlusRelatedResources struct {
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

func (t *TestScriptPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if t.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *t.RevIncludedListResourcesReferencingItem
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if t.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *t.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if t.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *t.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if t.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *t.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *t.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if t.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *t.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if t.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *t.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if t.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *t.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if t.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *t.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *t.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*t.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *t.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*t.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*t.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *t.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*t.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (t *TestScriptPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *t.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*t.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *t.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*t.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*t.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *t.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*t.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
