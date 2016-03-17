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

type OrderResponse struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Request        *Reference    `bson:"request,omitempty" json:"request,omitempty"`
	Date           *FHIRDateTime `bson:"date,omitempty" json:"date,omitempty"`
	Who            *Reference    `bson:"who,omitempty" json:"who,omitempty"`
	OrderStatus    string        `bson:"orderStatus,omitempty" json:"orderStatus,omitempty"`
	Description    string        `bson:"description,omitempty" json:"description,omitempty"`
	Fulfillment    []Reference   `bson:"fulfillment,omitempty" json:"fulfillment,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *OrderResponse) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "OrderResponse"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to OrderResponse), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *OrderResponse) GetBSON() (interface{}, error) {
	x.ResourceType = "OrderResponse"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "orderResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type orderResponse OrderResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *OrderResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := orderResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = OrderResponse(x2)
		return x.checkResourceType()
	}
	return
}

func (x *OrderResponse) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "OrderResponse"
	} else if x.ResourceType != "OrderResponse" {
		return errors.New(fmt.Sprintf("Expected resourceType to be OrderResponse, instead received %s", x.ResourceType))
	}
	return nil
}

type OrderResponsePlus struct {
	OrderResponse                     `bson:",inline"`
	OrderResponsePlusRelatedResources `bson:",inline"`
}

type OrderResponsePlusRelatedResources struct {
	IncludedOrderResourcesReferencedByRequest                   *[]Order                 `bson:"_includedOrderResourcesReferencedByRequest,omitempty"`
	IncludedPractitionerResourcesReferencedByWho                *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByWho,omitempty"`
	IncludedOrganizationResourcesReferencedByWho                *[]Organization          `bson:"_includedOrganizationResourcesReferencedByWho,omitempty"`
	IncludedDeviceResourcesReferencedByWho                      *[]Device                `bson:"_includedDeviceResourcesReferencedByWho,omitempty"`
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

func (o *OrderResponsePlusRelatedResources) GetIncludedOrderResourceReferencedByRequest() (order *Order, err error) {
	if o.IncludedOrderResourcesReferencedByRequest == nil {
		err = errors.New("Included orders not requested")
	} else if len(*o.IncludedOrderResourcesReferencedByRequest) > 1 {
		err = fmt.Errorf("Expected 0 or 1 order, but found %d", len(*o.IncludedOrderResourcesReferencedByRequest))
	} else if len(*o.IncludedOrderResourcesReferencedByRequest) == 1 {
		order = &(*o.IncludedOrderResourcesReferencedByRequest)[0]
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetIncludedPractitionerResourceReferencedByWho() (practitioner *Practitioner, err error) {
	if o.IncludedPractitionerResourcesReferencedByWho == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*o.IncludedPractitionerResourcesReferencedByWho) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*o.IncludedPractitionerResourcesReferencedByWho))
	} else if len(*o.IncludedPractitionerResourcesReferencedByWho) == 1 {
		practitioner = &(*o.IncludedPractitionerResourcesReferencedByWho)[0]
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetIncludedOrganizationResourceReferencedByWho() (organization *Organization, err error) {
	if o.IncludedOrganizationResourcesReferencedByWho == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*o.IncludedOrganizationResourcesReferencedByWho) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*o.IncludedOrganizationResourcesReferencedByWho))
	} else if len(*o.IncludedOrganizationResourcesReferencedByWho) == 1 {
		organization = &(*o.IncludedOrganizationResourcesReferencedByWho)[0]
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetIncludedDeviceResourceReferencedByWho() (device *Device, err error) {
	if o.IncludedDeviceResourcesReferencedByWho == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedDeviceResourcesReferencedByWho) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedDeviceResourcesReferencedByWho))
	} else if len(*o.IncludedDeviceResourcesReferencedByWho) == 1 {
		device = &(*o.IncludedDeviceResourcesReferencedByWho)[0]
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if o.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *o.RevIncludedListResourcesReferencingItem
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if o.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *o.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *o.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *o.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (o *OrderResponsePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedOrderResourcesReferencedByRequest != nil {
		for idx := range *o.IncludedOrderResourcesReferencedByRequest {
			rsc := (*o.IncludedOrderResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPractitionerResourcesReferencedByWho != nil {
		for idx := range *o.IncludedPractitionerResourcesReferencedByWho {
			rsc := (*o.IncludedPractitionerResourcesReferencedByWho)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedOrganizationResourcesReferencedByWho != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByWho {
			rsc := (*o.IncludedOrganizationResourcesReferencedByWho)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedByWho != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedByWho {
			rsc := (*o.IncludedDeviceResourcesReferencedByWho)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *OrderResponsePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *o.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*o.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *o.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*o.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *OrderResponsePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedOrderResourcesReferencedByRequest != nil {
		for idx := range *o.IncludedOrderResourcesReferencedByRequest {
			rsc := (*o.IncludedOrderResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPractitionerResourcesReferencedByWho != nil {
		for idx := range *o.IncludedPractitionerResourcesReferencedByWho {
			rsc := (*o.IncludedPractitionerResourcesReferencedByWho)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedOrganizationResourcesReferencedByWho != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByWho {
			rsc := (*o.IncludedOrganizationResourcesReferencedByWho)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedByWho != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedByWho {
			rsc := (*o.IncludedDeviceResourcesReferencedByWho)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *o.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*o.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *o.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*o.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
