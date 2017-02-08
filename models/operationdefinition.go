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
	Date           *FHIRDateTime                           `bson:"date,omitempty" json:"date,omitempty"`
	Publisher      string                                  `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ContactDetail                         `bson:"contact,omitempty" json:"contact,omitempty"`
	Description    string                                  `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []UsageContext                          `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction   []CodeableConcept                       `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose        string                                  `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Idempotent     *bool                                   `bson:"idempotent,omitempty" json:"idempotent,omitempty"`
	Code           string                                  `bson:"code,omitempty" json:"code,omitempty"`
	Comment        string                                  `bson:"comment,omitempty" json:"comment,omitempty"`
	Base           *Reference                              `bson:"base,omitempty" json:"base,omitempty"`
	Resource       []string                                `bson:"resource,omitempty" json:"resource,omitempty"`
	System         *bool                                   `bson:"system,omitempty" json:"system,omitempty"`
	Type           *bool                                   `bson:"type,omitempty" json:"type,omitempty"`
	Instance       *bool                                   `bson:"instance,omitempty" json:"instance,omitempty"`
	Parameter      []OperationDefinitionParameterComponent `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Overload       []OperationDefinitionOverloadComponent  `bson:"overload,omitempty" json:"overload,omitempty"`
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

type OperationDefinitionParameterComponent struct {
	BackboneElement `bson:",inline"`
	Name            string                                        `bson:"name,omitempty" json:"name,omitempty"`
	Use             string                                        `bson:"use,omitempty" json:"use,omitempty"`
	Min             *int32                                        `bson:"min,omitempty" json:"min,omitempty"`
	Max             string                                        `bson:"max,omitempty" json:"max,omitempty"`
	Documentation   string                                        `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Type            string                                        `bson:"type,omitempty" json:"type,omitempty"`
	SearchType      string                                        `bson:"searchType,omitempty" json:"searchType,omitempty"`
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

type OperationDefinitionOverloadComponent struct {
	BackboneElement `bson:",inline"`
	ParameterName   []string `bson:"parameterName,omitempty" json:"parameterName,omitempty"`
	Comment         string   `bson:"comment,omitempty" json:"comment,omitempty"`
}

type OperationDefinitionPlus struct {
	OperationDefinition                     `bson:",inline"`
	OperationDefinitionPlusRelatedResources `bson:",inline"`
}

type OperationDefinitionPlusRelatedResources struct {
	IncludedStructureDefinitionResourcesReferencedByParamprofile *[]StructureDefinition   `bson:"_includedStructureDefinitionResourcesReferencedByParamprofile,omitempty"`
	IncludedOperationDefinitionResourcesReferencedByBase         *[]OperationDefinition   `bson:"_includedOperationDefinitionResourcesReferencedByBase,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                   *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref   *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject               *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest          *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource   *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon          *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData             *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                   *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                     *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                   *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                      *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedOperationDefinitionResourcesReferencingBase       *[]OperationDefinition   `bson:"_revIncludedOperationDefinitionResourcesReferencingBase,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon      *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition   *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon       *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition    *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                  *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity              *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject            *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated       *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject  *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest        *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
}

func (o *OperationDefinitionPlusRelatedResources) GetIncludedStructureDefinitionResourceReferencedByParamprofile() (structureDefinition *StructureDefinition, err error) {
	if o.IncludedStructureDefinitionResourcesReferencedByParamprofile == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*o.IncludedStructureDefinitionResourcesReferencedByParamprofile) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*o.IncludedStructureDefinitionResourcesReferencedByParamprofile))
	} else if len(*o.IncludedStructureDefinitionResourcesReferencedByParamprofile) == 1 {
		structureDefinition = &(*o.IncludedStructureDefinitionResourcesReferencedByParamprofile)[0]
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

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingData
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

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if o.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *o.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingBasedon
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

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingEntity
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

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingBasedon
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

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if o.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *o.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if o.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *o.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if o.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *o.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if o.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *o.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if o.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *o.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if o.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *o.RevIncludedDeviceUseRequestResourcesReferencingDefinition
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

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingEntity
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

func (o *OperationDefinitionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedStructureDefinitionResourcesReferencedByParamprofile != nil {
		for idx := range *o.IncludedStructureDefinitionResourcesReferencedByParamprofile {
			rsc := (*o.IncludedStructureDefinitionResourcesReferencedByParamprofile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedOperationDefinitionResourcesReferencedByBase != nil {
		for idx := range *o.IncludedOperationDefinitionResourcesReferencedByBase {
			rsc := (*o.IncludedOperationDefinitionResourcesReferencedByBase)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *OperationDefinitionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingData {
			rsc := (*o.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*o.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingTopic {
			rsc := (*o.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*o.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*o.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*o.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOperationDefinitionResourcesReferencingBase != nil {
		for idx := range *o.RevIncludedOperationDefinitionResourcesReferencingBase {
			rsc := (*o.RevIncludedOperationDefinitionResourcesReferencingBase)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *OperationDefinitionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedStructureDefinitionResourcesReferencedByParamprofile != nil {
		for idx := range *o.IncludedStructureDefinitionResourcesReferencedByParamprofile {
			rsc := (*o.IncludedStructureDefinitionResourcesReferencedByParamprofile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedOperationDefinitionResourcesReferencedByBase != nil {
		for idx := range *o.IncludedOperationDefinitionResourcesReferencedByBase {
			rsc := (*o.IncludedOperationDefinitionResourcesReferencedByBase)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingData {
			rsc := (*o.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*o.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingTopic {
			rsc := (*o.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*o.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*o.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*o.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOperationDefinitionResourcesReferencingBase != nil {
		for idx := range *o.RevIncludedOperationDefinitionResourcesReferencingBase {
			rsc := (*o.RevIncludedOperationDefinitionResourcesReferencingBase)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *o.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*o.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *o.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*o.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
