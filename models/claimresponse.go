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

type ClaimResponse struct {
	DomainResource                `bson:",inline"`
	Identifier                    []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	RequestIdentifier             *Identifier                       `bson:"requestIdentifier,omitempty" json:"requestIdentifier,omitempty"`
	RequestReference              *Reference                        `bson:"requestReference,omitempty" json:"requestReference,omitempty"`
	Ruleset                       *Coding                           `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset               *Coding                           `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created                       *FHIRDateTime                     `bson:"created,omitempty" json:"created,omitempty"`
	OrganizationIdentifier        *Identifier                       `bson:"organizationIdentifier,omitempty" json:"organizationIdentifier,omitempty"`
	OrganizationReference         *Reference                        `bson:"organizationReference,omitempty" json:"organizationReference,omitempty"`
	RequestProviderIdentifier     *Identifier                       `bson:"requestProviderIdentifier,omitempty" json:"requestProviderIdentifier,omitempty"`
	RequestProviderReference      *Reference                        `bson:"requestProviderReference,omitempty" json:"requestProviderReference,omitempty"`
	RequestOrganizationIdentifier *Identifier                       `bson:"requestOrganizationIdentifier,omitempty" json:"requestOrganizationIdentifier,omitempty"`
	RequestOrganizationReference  *Reference                        `bson:"requestOrganizationReference,omitempty" json:"requestOrganizationReference,omitempty"`
	Outcome                       string                            `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition                   string                            `bson:"disposition,omitempty" json:"disposition,omitempty"`
	PayeeType                     *Coding                           `bson:"payeeType,omitempty" json:"payeeType,omitempty"`
	Item                          []ClaimResponseItemsComponent     `bson:"item,omitempty" json:"item,omitempty"`
	AddItem                       []ClaimResponseAddedItemComponent `bson:"addItem,omitempty" json:"addItem,omitempty"`
	Error                         []ClaimResponseErrorsComponent    `bson:"error,omitempty" json:"error,omitempty"`
	TotalCost                     *Quantity                         `bson:"totalCost,omitempty" json:"totalCost,omitempty"`
	UnallocDeductable             *Quantity                         `bson:"unallocDeductable,omitempty" json:"unallocDeductable,omitempty"`
	TotalBenefit                  *Quantity                         `bson:"totalBenefit,omitempty" json:"totalBenefit,omitempty"`
	PaymentAdjustment             *Quantity                         `bson:"paymentAdjustment,omitempty" json:"paymentAdjustment,omitempty"`
	PaymentAdjustmentReason       *Coding                           `bson:"paymentAdjustmentReason,omitempty" json:"paymentAdjustmentReason,omitempty"`
	PaymentDate                   *FHIRDateTime                     `bson:"paymentDate,omitempty" json:"paymentDate,omitempty"`
	PaymentAmount                 *Quantity                         `bson:"paymentAmount,omitempty" json:"paymentAmount,omitempty"`
	PaymentRef                    *Identifier                       `bson:"paymentRef,omitempty" json:"paymentRef,omitempty"`
	Reserved                      *Coding                           `bson:"reserved,omitempty" json:"reserved,omitempty"`
	Form                          *Coding                           `bson:"form,omitempty" json:"form,omitempty"`
	Note                          []ClaimResponseNotesComponent     `bson:"note,omitempty" json:"note,omitempty"`
	Coverage                      []ClaimResponseCoverageComponent  `bson:"coverage,omitempty" json:"coverage,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ClaimResponse) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ClaimResponse"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ClaimResponse), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ClaimResponse) GetBSON() (interface{}, error) {
	x.ResourceType = "ClaimResponse"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "claimResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type claimResponse ClaimResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ClaimResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := claimResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ClaimResponse(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ClaimResponse) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ClaimResponse"
	} else if x.ResourceType != "ClaimResponse" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ClaimResponse, instead received %s", x.ResourceType))
	}
	return nil
}

type ClaimResponseItemsComponent struct {
	BackboneElement `bson:",inline"`
	SequenceLinkId  *uint32                                  `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	NoteNumber      []uint32                                 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ClaimResponseItemAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail          []ClaimResponseItemDetailComponent       `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ClaimResponseItemAdjudicationComponent struct {
	BackboneElement `bson:",inline"`
	Category        *Coding   `bson:"category,omitempty" json:"category,omitempty"`
	Reason          *Coding   `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount          *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value           *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ClaimResponseItemDetailComponent struct {
	BackboneElement `bson:",inline"`
	SequenceLinkId  *uint32                                    `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Adjudication    []ClaimResponseDetailAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail       []ClaimResponseSubDetailComponent          `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ClaimResponseDetailAdjudicationComponent struct {
	BackboneElement `bson:",inline"`
	Category        *Coding   `bson:"category,omitempty" json:"category,omitempty"`
	Reason          *Coding   `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount          *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value           *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ClaimResponseSubDetailComponent struct {
	BackboneElement `bson:",inline"`
	SequenceLinkId  *uint32                                       `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Adjudication    []ClaimResponseSubdetailAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ClaimResponseSubdetailAdjudicationComponent struct {
	BackboneElement `bson:",inline"`
	Category        *Coding   `bson:"category,omitempty" json:"category,omitempty"`
	Reason          *Coding   `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount          *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value           *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ClaimResponseAddedItemComponent struct {
	BackboneElement  `bson:",inline"`
	SequenceLinkId   []uint32                                      `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Service          *Coding                                       `bson:"service,omitempty" json:"service,omitempty"`
	Fee              *Quantity                                     `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumberLinkId []uint32                                      `bson:"noteNumberLinkId,omitempty" json:"noteNumberLinkId,omitempty"`
	Adjudication     []ClaimResponseAddedItemAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail           []ClaimResponseAddedItemsDetailComponent      `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ClaimResponseAddedItemAdjudicationComponent struct {
	BackboneElement `bson:",inline"`
	Category        *Coding   `bson:"category,omitempty" json:"category,omitempty"`
	Reason          *Coding   `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount          *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value           *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ClaimResponseAddedItemsDetailComponent struct {
	BackboneElement `bson:",inline"`
	Service         *Coding                                             `bson:"service,omitempty" json:"service,omitempty"`
	Fee             *Quantity                                           `bson:"fee,omitempty" json:"fee,omitempty"`
	Adjudication    []ClaimResponseAddedItemDetailAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ClaimResponseAddedItemDetailAdjudicationComponent struct {
	BackboneElement `bson:",inline"`
	Category        *Coding   `bson:"category,omitempty" json:"category,omitempty"`
	Reason          *Coding   `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount          *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value           *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ClaimResponseErrorsComponent struct {
	BackboneElement         `bson:",inline"`
	SequenceLinkId          *uint32 `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	DetailSequenceLinkId    *uint32 `bson:"detailSequenceLinkId,omitempty" json:"detailSequenceLinkId,omitempty"`
	SubdetailSequenceLinkId *uint32 `bson:"subdetailSequenceLinkId,omitempty" json:"subdetailSequenceLinkId,omitempty"`
	Code                    *Coding `bson:"code,omitempty" json:"code,omitempty"`
}

type ClaimResponseNotesComponent struct {
	BackboneElement `bson:",inline"`
	Number          *uint32 `bson:"number,omitempty" json:"number,omitempty"`
	Type            *Coding `bson:"type,omitempty" json:"type,omitempty"`
	Text            string  `bson:"text,omitempty" json:"text,omitempty"`
}

type ClaimResponseCoverageComponent struct {
	BackboneElement     `bson:",inline"`
	Sequence            *uint32     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Focal               *bool       `bson:"focal,omitempty" json:"focal,omitempty"`
	CoverageIdentifier  *Identifier `bson:"coverageIdentifier,omitempty" json:"coverageIdentifier,omitempty"`
	CoverageReference   *Reference  `bson:"coverageReference,omitempty" json:"coverageReference,omitempty"`
	BusinessArrangement string      `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	PreAuthRef          []string    `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference  `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
}

type ClaimResponsePlus struct {
	ClaimResponse                     `bson:",inline"`
	ClaimResponsePlusRelatedResources `bson:",inline"`
}

type ClaimResponsePlusRelatedResources struct {
	IncludedOrganizationResourcesReferencedByOrganizationreference *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganizationreference,omitempty"`
	IncludedClaimResourcesReferencedByRequestreference             *[]Claim                 `bson:"_includedClaimResourcesReferencedByRequestreference,omitempty"`
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

func (c *ClaimResponsePlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganizationreference() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByOrganizationreference == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByOrganizationreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByOrganizationreference))
	} else if len(*c.IncludedOrganizationResourcesReferencedByOrganizationreference) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByOrganizationreference)[0]
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetIncludedClaimResourceReferencedByRequestreference() (claim *Claim, err error) {
	if c.IncludedClaimResourcesReferencedByRequestreference == nil {
		err = errors.New("Included claims not requested")
	} else if len(*c.IncludedClaimResourcesReferencedByRequestreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 claim, but found %d", len(*c.IncludedClaimResourcesReferencedByRequestreference))
	} else if len(*c.IncludedClaimResourcesReferencedByRequestreference) == 1 {
		claim = &(*c.IncludedClaimResourcesReferencedByRequestreference)[0]
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequestreference
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedOrganizationResourcesReferencedByOrganizationreference != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByOrganizationreference {
			rsc := (*c.IncludedOrganizationResourcesReferencedByOrganizationreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedClaimResourcesReferencedByRequestreference != nil {
		for idx := range *c.IncludedClaimResourcesReferencedByRequestreference {
			rsc := (*c.IncludedClaimResourcesReferencedByRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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

func (c *ClaimResponsePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedOrganizationResourcesReferencedByOrganizationreference != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByOrganizationreference {
			rsc := (*c.IncludedOrganizationResourcesReferencedByOrganizationreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedClaimResourcesReferencedByRequestreference != nil {
		for idx := range *c.IncludedClaimResourcesReferencedByRequestreference {
			rsc := (*c.IncludedClaimResourcesReferencedByRequestreference)[idx]
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
