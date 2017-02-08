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

type ExpansionProfile struct {
	DomainResource         `bson:",inline"`
	Url                    string                                   `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             *Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                string                                   `bson:"version,omitempty" json:"version,omitempty"`
	Name                   string                                   `bson:"name,omitempty" json:"name,omitempty"`
	Status                 string                                   `bson:"status,omitempty" json:"status,omitempty"`
	Experimental           *bool                                    `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher              string                                   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail                          `bson:"contact,omitempty" json:"contact,omitempty"`
	Date                   *FHIRDateTime                            `bson:"date,omitempty" json:"date,omitempty"`
	Description            string                                   `bson:"description,omitempty" json:"description,omitempty"`
	UseContext             []UsageContext                           `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                        `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	FixedVersion           []ExpansionProfileFixedVersionComponent  `bson:"fixedVersion,omitempty" json:"fixedVersion,omitempty"`
	ExcludedSystem         *ExpansionProfileExcludedSystemComponent `bson:"excludedSystem,omitempty" json:"excludedSystem,omitempty"`
	IncludeDesignations    *bool                                    `bson:"includeDesignations,omitempty" json:"includeDesignations,omitempty"`
	Designation            *ExpansionProfileDesignationComponent    `bson:"designation,omitempty" json:"designation,omitempty"`
	IncludeDefinition      *bool                                    `bson:"includeDefinition,omitempty" json:"includeDefinition,omitempty"`
	ActiveOnly             *bool                                    `bson:"activeOnly,omitempty" json:"activeOnly,omitempty"`
	ExcludeNested          *bool                                    `bson:"excludeNested,omitempty" json:"excludeNested,omitempty"`
	ExcludeNotForUI        *bool                                    `bson:"excludeNotForUI,omitempty" json:"excludeNotForUI,omitempty"`
	ExcludePostCoordinated *bool                                    `bson:"excludePostCoordinated,omitempty" json:"excludePostCoordinated,omitempty"`
	DisplayLanguage        string                                   `bson:"displayLanguage,omitempty" json:"displayLanguage,omitempty"`
	LimitedExpansion       *bool                                    `bson:"limitedExpansion,omitempty" json:"limitedExpansion,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ExpansionProfile) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ExpansionProfile"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ExpansionProfile), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ExpansionProfile) GetBSON() (interface{}, error) {
	x.ResourceType = "ExpansionProfile"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "expansionProfile" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type expansionProfile ExpansionProfile

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ExpansionProfile) UnmarshalJSON(data []byte) (err error) {
	x2 := expansionProfile{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ExpansionProfile(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ExpansionProfile) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ExpansionProfile"
	} else if x.ResourceType != "ExpansionProfile" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ExpansionProfile, instead received %s", x.ResourceType))
	}
	return nil
}

type ExpansionProfileFixedVersionComponent struct {
	BackboneElement `bson:",inline"`
	System          string `bson:"system,omitempty" json:"system,omitempty"`
	Version         string `bson:"version,omitempty" json:"version,omitempty"`
	Mode            string `bson:"mode,omitempty" json:"mode,omitempty"`
}

type ExpansionProfileExcludedSystemComponent struct {
	BackboneElement `bson:",inline"`
	System          string `bson:"system,omitempty" json:"system,omitempty"`
	Version         string `bson:"version,omitempty" json:"version,omitempty"`
}

type ExpansionProfileDesignationComponent struct {
	BackboneElement `bson:",inline"`
	Include         *ExpansionProfileDesignationIncludeComponent `bson:"include,omitempty" json:"include,omitempty"`
	Exclude         *ExpansionProfileDesignationExcludeComponent `bson:"exclude,omitempty" json:"exclude,omitempty"`
}

type ExpansionProfileDesignationIncludeComponent struct {
	BackboneElement `bson:",inline"`
	Designation     []ExpansionProfileDesignationIncludeDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
}

type ExpansionProfileDesignationIncludeDesignationComponent struct {
	BackboneElement `bson:",inline"`
	Language        string  `bson:"language,omitempty" json:"language,omitempty"`
	Use             *Coding `bson:"use,omitempty" json:"use,omitempty"`
}

type ExpansionProfileDesignationExcludeComponent struct {
	BackboneElement `bson:",inline"`
	Designation     []ExpansionProfileDesignationExcludeDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
}

type ExpansionProfileDesignationExcludeDesignationComponent struct {
	BackboneElement `bson:",inline"`
	Language        string  `bson:"language,omitempty" json:"language,omitempty"`
	Use             *Coding `bson:"use,omitempty" json:"use,omitempty"`
}

type ExpansionProfilePlus struct {
	ExpansionProfile                     `bson:",inline"`
	ExpansionProfilePlusRelatedResources `bson:",inline"`
}

type ExpansionProfilePlusRelatedResources struct {
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

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if e.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *e.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if e.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *e.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingItem
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if e.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *e.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if e.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *e.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if e.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *e.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if e.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *e.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if e.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *e.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if e.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *e.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if e.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *e.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *e.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if e.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *e.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingData {
			rsc := (*e.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*e.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*e.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*e.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*e.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*e.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *ExpansionProfilePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingData {
			rsc := (*e.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*e.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*e.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*e.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*e.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*e.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
