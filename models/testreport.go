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

type TestReport struct {
	DomainResource `bson:",inline"`
	Identifier     *Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name           string                           `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                           `bson:"status,omitempty" json:"status,omitempty"`
	Score          *float64                         `bson:"score,omitempty" json:"score,omitempty"`
	Tester         string                           `bson:"tester,omitempty" json:"tester,omitempty"`
	TestScript     *Reference                       `bson:"testScript,omitempty" json:"testScript,omitempty"`
	Issued         *FHIRDateTime                    `bson:"issued,omitempty" json:"issued,omitempty"`
	Participant    []TestReportParticipantComponent `bson:"participant,omitempty" json:"participant,omitempty"`
	Setup          *TestReportSetupComponent        `bson:"setup,omitempty" json:"setup,omitempty"`
	Test           []TestReportTestComponent        `bson:"test,omitempty" json:"test,omitempty"`
	Teardown       *TestReportTeardownComponent     `bson:"teardown,omitempty" json:"teardown,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *TestReport) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "TestReport"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to TestReport), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *TestReport) GetBSON() (interface{}, error) {
	x.ResourceType = "TestReport"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "testReport" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type testReport TestReport

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *TestReport) UnmarshalJSON(data []byte) (err error) {
	x2 := testReport{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = TestReport(x2)
		return x.checkResourceType()
	}
	return
}

func (x *TestReport) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "TestReport"
	} else if x.ResourceType != "TestReport" {
		return errors.New(fmt.Sprintf("Expected resourceType to be TestReport, instead received %s", x.ResourceType))
	}
	return nil
}

type TestReportParticipantComponent struct {
	BackboneElement `bson:",inline"`
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
	Uri             string `bson:"uri,omitempty" json:"uri,omitempty"`
	Display         string `bson:"display,omitempty" json:"display,omitempty"`
}

type TestReportSetupComponent struct {
	BackboneElement `bson:",inline"`
	Action          []TestReportSetupActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestReportSetupActionComponent struct {
	BackboneElement `bson:",inline"`
	Operation       *TestReportSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert          *TestReportSetupActionAssertComponent    `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestReportSetupActionOperationComponent struct {
	BackboneElement `bson:",inline"`
	Result          string `bson:"result,omitempty" json:"result,omitempty"`
	Message         string `bson:"message,omitempty" json:"message,omitempty"`
	Detail          string `bson:"detail,omitempty" json:"detail,omitempty"`
}

type TestReportSetupActionAssertComponent struct {
	BackboneElement `bson:",inline"`
	Result          string `bson:"result,omitempty" json:"result,omitempty"`
	Message         string `bson:"message,omitempty" json:"message,omitempty"`
	Detail          string `bson:"detail,omitempty" json:"detail,omitempty"`
}

type TestReportTestComponent struct {
	BackboneElement `bson:",inline"`
	Name            string                          `bson:"name,omitempty" json:"name,omitempty"`
	Description     string                          `bson:"description,omitempty" json:"description,omitempty"`
	Action          []TestReportTestActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestReportTestActionComponent struct {
	BackboneElement `bson:",inline"`
	Operation       *TestReportSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert          *TestReportSetupActionAssertComponent    `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestReportTeardownComponent struct {
	BackboneElement `bson:",inline"`
	Action          []TestReportTeardownActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestReportTeardownActionComponent struct {
	BackboneElement `bson:",inline"`
	Operation       *TestReportSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
}

type TestReportPlus struct {
	TestReport                     `bson:",inline"`
	TestReportPlusRelatedResources `bson:",inline"`
}

type TestReportPlusRelatedResources struct {
	IncludedTestScriptResourcesReferencedByTestscript           *[]TestScript            `bson:"_includedTestScriptResourcesReferencedByTestscript,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                  *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic               *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject              *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse        *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource  *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon         *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                    *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces    *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition  *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces     *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition   *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity             *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
}

func (t *TestReportPlusRelatedResources) GetIncludedTestScriptResourceReferencedByTestscript() (testScript *TestScript, err error) {
	if t.IncludedTestScriptResourcesReferencedByTestscript == nil {
		err = errors.New("Included testscripts not requested")
	} else if len(*t.IncludedTestScriptResourcesReferencedByTestscript) > 1 {
		err = fmt.Errorf("Expected 0 or 1 testScript, but found %d", len(*t.IncludedTestScriptResourcesReferencedByTestscript))
	} else if len(*t.IncludedTestScriptResourcesReferencedByTestscript) == 1 {
		testScript = &(*t.IncludedTestScriptResourcesReferencedByTestscript)[0]
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if t.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *t.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if t.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *t.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if t.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *t.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if t.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *t.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if t.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *t.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if t.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *t.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if t.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *t.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if t.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *t.RevIncludedListResourcesReferencingItem
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if t.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *t.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if t.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *t.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if t.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *t.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if t.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *t.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if t.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *t.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if t.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *t.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if t.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *t.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if t.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *t.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *t.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (t *TestReportPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if t.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *t.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (t *TestReportPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.IncludedTestScriptResourcesReferencedByTestscript != nil {
		for idx := range *t.IncludedTestScriptResourcesReferencedByTestscript {
			rsc := (*t.IncludedTestScriptResourcesReferencedByTestscript)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (t *TestReportPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
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
	if t.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingData {
			rsc := (*t.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*t.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingTopic {
			rsc := (*t.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*t.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*t.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *t.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*t.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *t.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*t.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *t.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*t.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *t.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*t.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
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
	return resourceMap
}

func (t *TestReportPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.IncludedTestScriptResourcesReferencedByTestscript != nil {
		for idx := range *t.IncludedTestScriptResourcesReferencedByTestscript {
			rsc := (*t.IncludedTestScriptResourcesReferencedByTestscript)[idx]
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
	if t.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingData {
			rsc := (*t.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*t.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingTopic {
			rsc := (*t.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*t.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*t.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *t.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*t.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *t.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*t.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *t.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*t.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *t.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*t.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
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
	return resourceMap
}
