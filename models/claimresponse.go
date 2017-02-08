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
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                            `bson:"status,omitempty" json:"status,omitempty"`
	Created              *FHIRDateTime                     `bson:"created,omitempty" json:"created,omitempty"`
	Insurer              *Reference                        `bson:"insurer,omitempty" json:"insurer,omitempty"`
	RequestProvider      *Reference                        `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization  *Reference                        `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Request              *Reference                        `bson:"request,omitempty" json:"request,omitempty"`
	Outcome              *CodeableConcept                  `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition          string                            `bson:"disposition,omitempty" json:"disposition,omitempty"`
	PayeeType            *CodeableConcept                  `bson:"payeeType,omitempty" json:"payeeType,omitempty"`
	Item                 []ClaimResponseItemComponent      `bson:"item,omitempty" json:"item,omitempty"`
	AddItem              []ClaimResponseAddedItemComponent `bson:"addItem,omitempty" json:"addItem,omitempty"`
	Error                []ClaimResponseErrorComponent     `bson:"error,omitempty" json:"error,omitempty"`
	TotalCost            *Quantity                         `bson:"totalCost,omitempty" json:"totalCost,omitempty"`
	UnallocDeductable    *Quantity                         `bson:"unallocDeductable,omitempty" json:"unallocDeductable,omitempty"`
	TotalBenefit         *Quantity                         `bson:"totalBenefit,omitempty" json:"totalBenefit,omitempty"`
	Payment              *ClaimResponsePaymentComponent    `bson:"payment,omitempty" json:"payment,omitempty"`
	Reserved             *Coding                           `bson:"reserved,omitempty" json:"reserved,omitempty"`
	Form                 *CodeableConcept                  `bson:"form,omitempty" json:"form,omitempty"`
	Note                 []ClaimResponseNoteComponent      `bson:"note,omitempty" json:"note,omitempty"`
	CommunicationRequest []Reference                       `bson:"communicationRequest,omitempty" json:"communicationRequest,omitempty"`
	Insurance            []ClaimResponseInsuranceComponent `bson:"insurance,omitempty" json:"insurance,omitempty"`
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

type ClaimResponseItemComponent struct {
	BackboneElement `bson:",inline"`
	SequenceLinkId  *uint32                              `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	NoteNumber      []uint32                             `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ClaimResponseAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail          []ClaimResponseItemDetailComponent   `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ClaimResponseAdjudicationComponent struct {
	BackboneElement `bson:",inline"`
	Category        *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Reason          *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount          *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
	Value           *float64         `bson:"value,omitempty" json:"value,omitempty"`
}

type ClaimResponseItemDetailComponent struct {
	BackboneElement `bson:",inline"`
	SequenceLinkId  *uint32                              `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	NoteNumber      []uint32                             `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ClaimResponseAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail       []ClaimResponseSubDetailComponent    `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ClaimResponseSubDetailComponent struct {
	BackboneElement `bson:",inline"`
	SequenceLinkId  *uint32                              `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	NoteNumber      []uint32                             `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ClaimResponseAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ClaimResponseAddedItemComponent struct {
	BackboneElement `bson:",inline"`
	SequenceLinkId  []uint32                                 `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Revenue         *CodeableConcept                         `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *CodeableConcept                         `bson:"category,omitempty" json:"category,omitempty"`
	Service         *CodeableConcept                         `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []CodeableConcept                        `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee             *Quantity                                `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber      []uint32                                 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ClaimResponseAdjudicationComponent     `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail          []ClaimResponseAddedItemsDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ClaimResponseAddedItemsDetailComponent struct {
	BackboneElement `bson:",inline"`
	Revenue         *CodeableConcept                     `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *CodeableConcept                     `bson:"category,omitempty" json:"category,omitempty"`
	Service         *CodeableConcept                     `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []CodeableConcept                    `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee             *Quantity                            `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber      []uint32                             `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ClaimResponseAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ClaimResponseErrorComponent struct {
	BackboneElement         `bson:",inline"`
	SequenceLinkId          *uint32          `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	DetailSequenceLinkId    *uint32          `bson:"detailSequenceLinkId,omitempty" json:"detailSequenceLinkId,omitempty"`
	SubdetailSequenceLinkId *uint32          `bson:"subdetailSequenceLinkId,omitempty" json:"subdetailSequenceLinkId,omitempty"`
	Code                    *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}

type ClaimResponsePaymentComponent struct {
	BackboneElement  `bson:",inline"`
	Type             *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Adjustment       *Quantity        `bson:"adjustment,omitempty" json:"adjustment,omitempty"`
	AdjustmentReason *CodeableConcept `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty"`
	Date             *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Amount           *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
	Identifier       *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
}

type ClaimResponseNoteComponent struct {
	BackboneElement `bson:",inline"`
	Number          *uint32          `bson:"number,omitempty" json:"number,omitempty"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Text            string           `bson:"text,omitempty" json:"text,omitempty"`
	Language        *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}

type ClaimResponseInsuranceComponent struct {
	BackboneElement     `bson:",inline"`
	Sequence            *uint32    `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Focal               *bool      `bson:"focal,omitempty" json:"focal,omitempty"`
	Coverage            *Reference `bson:"coverage,omitempty" json:"coverage,omitempty"`
	BusinessArrangement string     `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	PreAuthRef          []string   `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
}

type ClaimResponsePlus struct {
	ClaimResponse                     `bson:",inline"`
	ClaimResponsePlusRelatedResources `bson:",inline"`
}

type ClaimResponsePlusRelatedResources struct {
	IncludedClaimResourcesReferencedByRequest                   *[]Claim                 `bson:"_includedClaimResourcesReferencedByRequest,omitempty"`
	IncludedOrganizationResourcesReferencedByInsurer            *[]Organization          `bson:"_includedOrganizationResourcesReferencedByInsurer,omitempty"`
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

func (c *ClaimResponsePlusRelatedResources) GetIncludedClaimResourceReferencedByRequest() (claim *Claim, err error) {
	if c.IncludedClaimResourcesReferencedByRequest == nil {
		err = errors.New("Included claims not requested")
	} else if len(*c.IncludedClaimResourcesReferencedByRequest) > 1 {
		err = fmt.Errorf("Expected 0 or 1 claim, but found %d", len(*c.IncludedClaimResourcesReferencedByRequest))
	} else if len(*c.IncludedClaimResourcesReferencedByRequest) == 1 {
		claim = &(*c.IncludedClaimResourcesReferencedByRequest)[0]
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetIncludedOrganizationResourceReferencedByInsurer() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByInsurer == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByInsurer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByInsurer))
	} else if len(*c.IncludedOrganizationResourcesReferencedByInsurer) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByInsurer)[0]
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

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingData
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

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponse
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

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingBasedon
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

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingEntity
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

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingBasedon
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

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if c.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *c.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if c.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *c.RevIncludedDeviceUseRequestResourcesReferencingDefinition
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

func (c *ClaimResponsePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *ClaimResponsePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedClaimResourcesReferencedByRequest != nil {
		for idx := range *c.IncludedClaimResourcesReferencedByRequest {
			rsc := (*c.IncludedClaimResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByInsurer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByInsurer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByInsurer)[idx]
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

func (c *ClaimResponsePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedClaimResourcesReferencedByRequest != nil {
		for idx := range *c.IncludedClaimResourcesReferencedByRequest {
			rsc := (*c.IncludedClaimResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByInsurer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByInsurer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByInsurer)[idx]
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
