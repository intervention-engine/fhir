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

type Coverage struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status         string                  `bson:"status,omitempty" json:"status,omitempty"`
	Type           *CodeableConcept        `bson:"type,omitempty" json:"type,omitempty"`
	PolicyHolder   *Reference              `bson:"policyHolder,omitempty" json:"policyHolder,omitempty"`
	Subscriber     *Reference              `bson:"subscriber,omitempty" json:"subscriber,omitempty"`
	SubscriberId   string                  `bson:"subscriberId,omitempty" json:"subscriberId,omitempty"`
	Beneficiary    *Reference              `bson:"beneficiary,omitempty" json:"beneficiary,omitempty"`
	Relationship   *CodeableConcept        `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Period         *Period                 `bson:"period,omitempty" json:"period,omitempty"`
	Payor          []Reference             `bson:"payor,omitempty" json:"payor,omitempty"`
	Group          *CoverageGroupComponent `bson:"group,omitempty" json:"group,omitempty"`
	Dependent      string                  `bson:"dependent,omitempty" json:"dependent,omitempty"`
	Sequence       string                  `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Order          *uint32                 `bson:"order,omitempty" json:"order,omitempty"`
	Network        string                  `bson:"network,omitempty" json:"network,omitempty"`
	Contract       []Reference             `bson:"contract,omitempty" json:"contract,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Coverage) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Coverage"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Coverage), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Coverage) GetBSON() (interface{}, error) {
	x.ResourceType = "Coverage"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "coverage" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type coverage Coverage

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Coverage) UnmarshalJSON(data []byte) (err error) {
	x2 := coverage{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Coverage(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Coverage) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Coverage"
	} else if x.ResourceType != "Coverage" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Coverage, instead received %s", x.ResourceType))
	}
	return nil
}

type CoverageGroupComponent struct {
	BackboneElement `bson:",inline"`
	Group           string `bson:"group,omitempty" json:"group,omitempty"`
	GroupDisplay    string `bson:"groupDisplay,omitempty" json:"groupDisplay,omitempty"`
	SubGroup        string `bson:"subGroup,omitempty" json:"subGroup,omitempty"`
	SubGroupDisplay string `bson:"subGroupDisplay,omitempty" json:"subGroupDisplay,omitempty"`
	Plan            string `bson:"plan,omitempty" json:"plan,omitempty"`
	PlanDisplay     string `bson:"planDisplay,omitempty" json:"planDisplay,omitempty"`
	SubPlan         string `bson:"subPlan,omitempty" json:"subPlan,omitempty"`
	SubPlanDisplay  string `bson:"subPlanDisplay,omitempty" json:"subPlanDisplay,omitempty"`
	Class           string `bson:"class,omitempty" json:"class,omitempty"`
	ClassDisplay    string `bson:"classDisplay,omitempty" json:"classDisplay,omitempty"`
	SubClass        string `bson:"subClass,omitempty" json:"subClass,omitempty"`
	SubClassDisplay string `bson:"subClassDisplay,omitempty" json:"subClassDisplay,omitempty"`
}

type CoveragePlus struct {
	Coverage                     `bson:",inline"`
	CoveragePlusRelatedResources `bson:",inline"`
}

type CoveragePlusRelatedResources struct {
	IncludedOrganizationResourcesReferencedByPolicyholder       *[]Organization          `bson:"_includedOrganizationResourcesReferencedByPolicyholder,omitempty"`
	IncludedPatientResourcesReferencedByPolicyholder            *[]Patient               `bson:"_includedPatientResourcesReferencedByPolicyholder,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPolicyholder      *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByPolicyholder,omitempty"`
	IncludedPatientResourcesReferencedBySubscriber              *[]Patient               `bson:"_includedPatientResourcesReferencedBySubscriber,omitempty"`
	IncludedRelatedPersonResourcesReferencedBySubscriber        *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedBySubscriber,omitempty"`
	IncludedOrganizationResourcesReferencedByPayor              *[]Organization          `bson:"_includedOrganizationResourcesReferencedByPayor,omitempty"`
	IncludedPatientResourcesReferencedByPayor                   *[]Patient               `bson:"_includedPatientResourcesReferencedByPayor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPayor             *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByPayor,omitempty"`
	IncludedPatientResourcesReferencedByBeneficiary             *[]Patient               `bson:"_includedPatientResourcesReferencedByBeneficiary,omitempty"`
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
	RevIncludedExplanationOfBenefitResourcesReferencingCoverage *[]ExplanationOfBenefit  `bson:"_revIncludedExplanationOfBenefitResourcesReferencingCoverage,omitempty"`
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

func (c *CoveragePlusRelatedResources) GetIncludedOrganizationResourceReferencedByPolicyholder() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByPolicyholder == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByPolicyholder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByPolicyholder))
	} else if len(*c.IncludedOrganizationResourcesReferencedByPolicyholder) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByPolicyholder)[0]
	}
	return
}

func (c *CoveragePlusRelatedResources) GetIncludedPatientResourceReferencedByPolicyholder() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPolicyholder == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPolicyholder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPolicyholder))
	} else if len(*c.IncludedPatientResourcesReferencedByPolicyholder) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPolicyholder)[0]
	}
	return
}

func (c *CoveragePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByPolicyholder() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByPolicyholder == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByPolicyholder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedByPolicyholder))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByPolicyholder) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedByPolicyholder)[0]
	}
	return
}

func (c *CoveragePlusRelatedResources) GetIncludedPatientResourceReferencedBySubscriber() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySubscriber == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedBySubscriber) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedBySubscriber))
	} else if len(*c.IncludedPatientResourcesReferencedBySubscriber) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedBySubscriber)[0]
	}
	return
}

func (c *CoveragePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedBySubscriber() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedBySubscriber == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedBySubscriber) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedBySubscriber))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedBySubscriber) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedBySubscriber)[0]
	}
	return
}

func (c *CoveragePlusRelatedResources) GetIncludedOrganizationResourcesReferencedByPayor() (organizations []Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByPayor == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *c.IncludedOrganizationResourcesReferencedByPayor
	}
	return
}

func (c *CoveragePlusRelatedResources) GetIncludedPatientResourcesReferencedByPayor() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPayor == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedByPayor
	}
	return
}

func (c *CoveragePlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByPayor() (relatedPeople []RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByPayor == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *c.IncludedRelatedPersonResourcesReferencedByPayor
	}
	return
}

func (c *CoveragePlusRelatedResources) GetIncludedPatientResourceReferencedByBeneficiary() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByBeneficiary == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByBeneficiary) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByBeneficiary))
	} else if len(*c.IncludedPatientResourcesReferencedByBeneficiary) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByBeneficiary)[0]
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingCoverage() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if c.RevIncludedExplanationOfBenefitResourcesReferencingCoverage == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *c.RevIncludedExplanationOfBenefitResourcesReferencingCoverage
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *CoveragePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *CoveragePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedOrganizationResourcesReferencedByPolicyholder != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPolicyholder {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPolicyholder != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPolicyholder {
			rsc := (*c.IncludedPatientResourcesReferencedByPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPolicyholder != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPolicyholder {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubscriber != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubscriber {
			rsc := (*c.IncludedPatientResourcesReferencedBySubscriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedBySubscriber != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedBySubscriber {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedBySubscriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPayor != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPayor {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPayor != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPayor {
			rsc := (*c.IncludedPatientResourcesReferencedByPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPayor != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPayor {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByBeneficiary != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByBeneficiary {
			rsc := (*c.IncludedPatientResourcesReferencedByBeneficiary)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *CoveragePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
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
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedExplanationOfBenefitResourcesReferencingCoverage != nil {
		for idx := range *c.RevIncludedExplanationOfBenefitResourcesReferencingCoverage {
			rsc := (*c.RevIncludedExplanationOfBenefitResourcesReferencingCoverage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
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
	return resourceMap
}

func (c *CoveragePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedOrganizationResourcesReferencedByPolicyholder != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPolicyholder {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPolicyholder != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPolicyholder {
			rsc := (*c.IncludedPatientResourcesReferencedByPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPolicyholder != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPolicyholder {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubscriber != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubscriber {
			rsc := (*c.IncludedPatientResourcesReferencedBySubscriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedBySubscriber != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedBySubscriber {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedBySubscriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPayor != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPayor {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPayor != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPayor {
			rsc := (*c.IncludedPatientResourcesReferencedByPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPayor != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPayor {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByBeneficiary != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByBeneficiary {
			rsc := (*c.IncludedPatientResourcesReferencedByBeneficiary)[idx]
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
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedExplanationOfBenefitResourcesReferencingCoverage != nil {
		for idx := range *c.RevIncludedExplanationOfBenefitResourcesReferencingCoverage {
			rsc := (*c.RevIncludedExplanationOfBenefitResourcesReferencingCoverage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*c.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
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
	return resourceMap
}
