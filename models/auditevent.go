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

type AuditEvent struct {
	DomainResource `bson:",inline"`
	Type           *Coding                     `bson:"type,omitempty" json:"type,omitempty"`
	Subtype        []Coding                    `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Action         string                      `bson:"action,omitempty" json:"action,omitempty"`
	Recorded       *FHIRDateTime               `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Outcome        string                      `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OutcomeDesc    string                      `bson:"outcomeDesc,omitempty" json:"outcomeDesc,omitempty"`
	PurposeOfEvent []Coding                    `bson:"purposeOfEvent,omitempty" json:"purposeOfEvent,omitempty"`
	Agent          []AuditEventAgentComponent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Source         *AuditEventSourceComponent  `bson:"source,omitempty" json:"source,omitempty"`
	Entity         []AuditEventEntityComponent `bson:"entity,omitempty" json:"entity,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *AuditEvent) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "AuditEvent"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to AuditEvent), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *AuditEvent) GetBSON() (interface{}, error) {
	x.ResourceType = "AuditEvent"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "auditEvent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type auditEvent AuditEvent

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *AuditEvent) UnmarshalJSON(data []byte) (err error) {
	x2 := auditEvent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = AuditEvent(x2)
		return x.checkResourceType()
	}
	return
}

func (x *AuditEvent) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "AuditEvent"
	} else if x.ResourceType != "AuditEvent" {
		return errors.New(fmt.Sprintf("Expected resourceType to be AuditEvent, instead received %s", x.ResourceType))
	}
	return nil
}

type AuditEventAgentComponent struct {
	BackboneElement `bson:",inline"`
	Role            []CodeableConcept                `bson:"role,omitempty" json:"role,omitempty"`
	Reference       *Reference                       `bson:"reference,omitempty" json:"reference,omitempty"`
	UserId          *Identifier                      `bson:"userId,omitempty" json:"userId,omitempty"`
	AltId           string                           `bson:"altId,omitempty" json:"altId,omitempty"`
	Name            string                           `bson:"name,omitempty" json:"name,omitempty"`
	Requestor       *bool                            `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Location        *Reference                       `bson:"location,omitempty" json:"location,omitempty"`
	Policy          []string                         `bson:"policy,omitempty" json:"policy,omitempty"`
	Media           *Coding                          `bson:"media,omitempty" json:"media,omitempty"`
	Network         *AuditEventAgentNetworkComponent `bson:"network,omitempty" json:"network,omitempty"`
	PurposeOfUse    []Coding                         `bson:"purposeOfUse,omitempty" json:"purposeOfUse,omitempty"`
}

type AuditEventAgentNetworkComponent struct {
	BackboneElement `bson:",inline"`
	Address         string `bson:"address,omitempty" json:"address,omitempty"`
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventSourceComponent struct {
	BackboneElement `bson:",inline"`
	Site            string      `bson:"site,omitempty" json:"site,omitempty"`
	Identifier      *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type            []Coding    `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventEntityComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      *Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Reference       *Reference                        `bson:"reference,omitempty" json:"reference,omitempty"`
	Type            *Coding                           `bson:"type,omitempty" json:"type,omitempty"`
	Role            *Coding                           `bson:"role,omitempty" json:"role,omitempty"`
	Lifecycle       *Coding                           `bson:"lifecycle,omitempty" json:"lifecycle,omitempty"`
	SecurityLabel   []Coding                          `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Name            string                            `bson:"name,omitempty" json:"name,omitempty"`
	Description     string                            `bson:"description,omitempty" json:"description,omitempty"`
	Query           string                            `bson:"query,omitempty" json:"query,omitempty"`
	Detail          []AuditEventEntityDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type AuditEventEntityDetailComponent struct {
	BackboneElement `bson:",inline"`
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
	Value           string `bson:"value,omitempty" json:"value,omitempty"`
}

type AuditEventPlus struct {
	AuditEvent                     `bson:",inline"`
	AuditEventPlusRelatedResources `bson:",inline"`
}

type AuditEventPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByAgent                 *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAgent,omitempty"`
	IncludedOrganizationResourcesReferencedByAgent                 *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAgent,omitempty"`
	IncludedDeviceResourcesReferencedByAgent                       *[]Device                `bson:"_includedDeviceResourcesReferencedByAgent,omitempty"`
	IncludedPatientResourcesReferencedByAgent                      *[]Patient               `bson:"_includedPatientResourcesReferencedByAgent,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAgent                *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByAgent,omitempty"`
	IncludedPatientResourcesReferencedByPatientPath1               *[]Patient               `bson:"_includedPatientResourcesReferencedByPatientPath1,omitempty"`
	IncludedPatientResourcesReferencedByPatientPath2               *[]Patient               `bson:"_includedPatientResourcesReferencedByPatientPath2,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref     *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                   *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponsereference  *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponsereference,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequestreference   *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequestreference,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource     *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment        *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                     *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingItem                        *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                     *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                    *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated         *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject    *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestreference *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequestreference,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
}

func (a *AuditEventPlusRelatedResources) GetIncludedPractitionerResourceReferencedByAgent() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByAgent == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByAgent))
	} else if len(*a.IncludedPractitionerResourcesReferencedByAgent) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByAgent)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedOrganizationResourceReferencedByAgent() (organization *Organization, err error) {
	if a.IncludedOrganizationResourcesReferencedByAgent == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*a.IncludedOrganizationResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*a.IncludedOrganizationResourcesReferencedByAgent))
	} else if len(*a.IncludedOrganizationResourcesReferencedByAgent) == 1 {
		organization = &(*a.IncludedOrganizationResourcesReferencedByAgent)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedDeviceResourceReferencedByAgent() (device *Device, err error) {
	if a.IncludedDeviceResourcesReferencedByAgent == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedDeviceResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedDeviceResourcesReferencedByAgent))
	} else if len(*a.IncludedDeviceResourcesReferencedByAgent) == 1 {
		device = &(*a.IncludedDeviceResourcesReferencedByAgent)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPatientResourceReferencedByAgent() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByAgent == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByAgent))
	} else if len(*a.IncludedPatientResourcesReferencedByAgent) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByAgent)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAgent() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedRelatedPersonResourcesReferencedByAgent == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedRelatedPersonResourcesReferencedByAgent))
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByAgent) == 1 {
		relatedPerson = &(*a.IncludedRelatedPersonResourcesReferencedByAgent)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPatientResourceReferencedByPatientPath1() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatientPath1 == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath1) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatientPath1))
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath1) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatientPath1)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPatientResourceReferencedByPatientPath2() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatientPath2 == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath2) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatientPath2))
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath2) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatientPath2)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingRequestreference
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if a.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *a.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *a.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if a.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *a.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if a.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *a.RevIncludedListResourcesReferencingItem
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if a.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *a.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if a.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *a.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if a.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *a.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *a.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if a.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *a.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *a.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByAgent {
			rsc := (*a.IncludedPractitionerResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedOrganizationResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedOrganizationResourcesReferencedByAgent {
			rsc := (*a.IncludedOrganizationResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedByAgent {
			rsc := (*a.IncludedDeviceResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByAgent {
			rsc := (*a.IncludedPatientResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByAgent {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath1 != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatientPath1 {
			rsc := (*a.IncludedPatientResourcesReferencedByPatientPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath2 != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatientPath2 {
			rsc := (*a.IncludedPatientResourcesReferencedByPatientPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *a.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*a.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *a.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*a.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *a.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*a.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AuditEventPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByAgent {
			rsc := (*a.IncludedPractitionerResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedOrganizationResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedOrganizationResourcesReferencedByAgent {
			rsc := (*a.IncludedOrganizationResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedByAgent {
			rsc := (*a.IncludedDeviceResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByAgent {
			rsc := (*a.IncludedPatientResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByAgent {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath1 != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatientPath1 {
			rsc := (*a.IncludedPatientResourcesReferencedByPatientPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath2 != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatientPath2 {
			rsc := (*a.IncludedPatientResourcesReferencedByPatientPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *a.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*a.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *a.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*a.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *a.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*a.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
