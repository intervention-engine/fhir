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

type CareTeam struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *CodeableConcept               `bson:"status,omitempty" json:"status,omitempty"`
	Type                 []CodeableConcept              `bson:"type,omitempty" json:"type,omitempty"`
	Name                 string                         `bson:"name,omitempty" json:"name,omitempty"`
	Subject              *Reference                     `bson:"subject,omitempty" json:"subject,omitempty"`
	Period               *Period                        `bson:"period,omitempty" json:"period,omitempty"`
	Participant          []CareTeamParticipantComponent `bson:"participant,omitempty" json:"participant,omitempty"`
	ManagingOrganization *Reference                     `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *CareTeam) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "CareTeam"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to CareTeam), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *CareTeam) GetBSON() (interface{}, error) {
	x.ResourceType = "CareTeam"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "careTeam" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type careTeam CareTeam

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *CareTeam) UnmarshalJSON(data []byte) (err error) {
	x2 := careTeam{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = CareTeam(x2)
		return x.checkResourceType()
	}
	return
}

func (x *CareTeam) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "CareTeam"
	} else if x.ResourceType != "CareTeam" {
		return errors.New(fmt.Sprintf("Expected resourceType to be CareTeam, instead received %s", x.ResourceType))
	}
	return nil
}

type CareTeamParticipantComponent struct {
	BackboneElement `bson:",inline"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Member          *Reference       `bson:"member,omitempty" json:"member,omitempty"`
	Period          *Period          `bson:"period,omitempty" json:"period,omitempty"`
}

type CareTeamPlus struct {
	CareTeam                     `bson:",inline"`
	CareTeamPlusRelatedResources `bson:",inline"`
}

type CareTeamPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                    *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedGroupResourcesReferencedBySubject                      *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                    *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPractitionerResourcesReferencedByParticipant           *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByParticipant,omitempty"`
	IncludedOrganizationResourcesReferencedByParticipant           *[]Organization          `bson:"_includedOrganizationResourcesReferencedByParticipant,omitempty"`
	IncludedPatientResourcesReferencedByParticipant                *[]Patient               `bson:"_includedPatientResourcesReferencedByParticipant,omitempty"`
	IncludedRelatedPersonResourcesReferencedByParticipant          *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByParticipant,omitempty"`
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

func (c *CareTeamPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPatient))
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if c.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*c.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*c.IncludedGroupResourcesReferencedBySubject))
	} else if len(*c.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*c.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedBySubject))
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetIncludedPractitionerResourceReferencedByParticipant() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByParticipant == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByParticipant))
	} else if len(*c.IncludedPractitionerResourcesReferencedByParticipant) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByParticipant)[0]
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetIncludedOrganizationResourceReferencedByParticipant() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByParticipant == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByParticipant))
	} else if len(*c.IncludedOrganizationResourcesReferencedByParticipant) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByParticipant)[0]
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetIncludedPatientResourceReferencedByParticipant() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByParticipant == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByParticipant))
	} else if len(*c.IncludedPatientResourcesReferencedByParticipant) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByParticipant)[0]
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByParticipant() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByParticipant == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedByParticipant))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByParticipant) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedByParticipant)[0]
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (c *CareTeamPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedGroupResourcesReferencedBySubject {
			rsc := (*c.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubject {
			rsc := (*c.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByParticipant {
			rsc := (*c.IncludedPractitionerResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByParticipant {
			rsc := (*c.IncludedOrganizationResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByParticipant {
			rsc := (*c.IncludedPatientResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByParticipant {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *CareTeamPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*c.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
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
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
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
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *CareTeamPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedGroupResourcesReferencedBySubject {
			rsc := (*c.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubject {
			rsc := (*c.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByParticipant {
			rsc := (*c.IncludedPractitionerResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByParticipant {
			rsc := (*c.IncludedOrganizationResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByParticipant {
			rsc := (*c.IncludedPatientResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByParticipant {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByParticipant)[idx]
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
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*c.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingData)[idx]
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
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
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
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
