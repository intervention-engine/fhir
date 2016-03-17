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

type Order struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Date                  *FHIRDateTime       `bson:"date,omitempty" json:"date,omitempty"`
	Subject               *Reference          `bson:"subject,omitempty" json:"subject,omitempty"`
	Source                *Reference          `bson:"source,omitempty" json:"source,omitempty"`
	Target                *Reference          `bson:"target,omitempty" json:"target,omitempty"`
	ReasonCodeableConcept *CodeableConcept    `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference          `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	When                  *OrderWhenComponent `bson:"when,omitempty" json:"when,omitempty"`
	Detail                []Reference         `bson:"detail,omitempty" json:"detail,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Order) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Order"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Order), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Order) GetBSON() (interface{}, error) {
	x.ResourceType = "Order"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "order" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type order Order

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Order) UnmarshalJSON(data []byte) (err error) {
	x2 := order{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Order(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Order) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Order"
	} else if x.ResourceType != "Order" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Order, instead received %s", x.ResourceType))
	}
	return nil
}

type OrderWhenComponent struct {
	BackboneElement `bson:",inline"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Schedule        *Timing          `bson:"schedule,omitempty" json:"schedule,omitempty"`
}

type OrderPlus struct {
	Order                     `bson:",inline"`
	OrderPlusRelatedResources `bson:",inline"`
}

type OrderPlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                   *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                  *[]Device                `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedSubstanceResourcesReferencedBySubject               *[]Substance             `bson:"_includedSubstanceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedBySource             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySource,omitempty"`
	IncludedOrganizationResourcesReferencedBySource             *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySource,omitempty"`
	IncludedPractitionerResourcesReferencedByTarget             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByTarget,omitempty"`
	IncludedOrganizationResourcesReferencedByTarget             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByTarget,omitempty"`
	IncludedDeviceResourcesReferencedByTarget                   *[]Device                `bson:"_includedDeviceResourcesReferencedByTarget,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference    *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingRequest         *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingRequest,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPlan       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPlan,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (o *OrderPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if o.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*o.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*o.IncludedGroupResourcesReferencedBySubject))
	} else if len(*o.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*o.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (o *OrderPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if o.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*o.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*o.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (o *OrderPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if o.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedPatientResourcesReferencedBySubject))
	} else if len(*o.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*o.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (o *OrderPlusRelatedResources) GetIncludedSubstanceResourceReferencedBySubject() (substance *Substance, err error) {
	if o.IncludedSubstanceResourcesReferencedBySubject == nil {
		err = errors.New("Included substances not requested")
	} else if len(*o.IncludedSubstanceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*o.IncludedSubstanceResourcesReferencedBySubject))
	} else if len(*o.IncludedSubstanceResourcesReferencedBySubject) == 1 {
		substance = &(*o.IncludedSubstanceResourcesReferencedBySubject)[0]
	}
	return
}

func (o *OrderPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if o.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedPatientResourcesReferencedByPatient))
	} else if len(*o.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*o.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (o *OrderPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySource() (practitioner *Practitioner, err error) {
	if o.IncludedPractitionerResourcesReferencedBySource == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*o.IncludedPractitionerResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*o.IncludedPractitionerResourcesReferencedBySource))
	} else if len(*o.IncludedPractitionerResourcesReferencedBySource) == 1 {
		practitioner = &(*o.IncludedPractitionerResourcesReferencedBySource)[0]
	}
	return
}

func (o *OrderPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySource() (organization *Organization, err error) {
	if o.IncludedOrganizationResourcesReferencedBySource == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*o.IncludedOrganizationResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*o.IncludedOrganizationResourcesReferencedBySource))
	} else if len(*o.IncludedOrganizationResourcesReferencedBySource) == 1 {
		organization = &(*o.IncludedOrganizationResourcesReferencedBySource)[0]
	}
	return
}

func (o *OrderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByTarget() (practitioner *Practitioner, err error) {
	if o.IncludedPractitionerResourcesReferencedByTarget == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*o.IncludedPractitionerResourcesReferencedByTarget) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*o.IncludedPractitionerResourcesReferencedByTarget))
	} else if len(*o.IncludedPractitionerResourcesReferencedByTarget) == 1 {
		practitioner = &(*o.IncludedPractitionerResourcesReferencedByTarget)[0]
	}
	return
}

func (o *OrderPlusRelatedResources) GetIncludedOrganizationResourceReferencedByTarget() (organization *Organization, err error) {
	if o.IncludedOrganizationResourcesReferencedByTarget == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*o.IncludedOrganizationResourcesReferencedByTarget) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*o.IncludedOrganizationResourcesReferencedByTarget))
	} else if len(*o.IncludedOrganizationResourcesReferencedByTarget) == 1 {
		organization = &(*o.IncludedOrganizationResourcesReferencedByTarget)[0]
	}
	return
}

func (o *OrderPlusRelatedResources) GetIncludedDeviceResourceReferencedByTarget() (device *Device, err error) {
	if o.IncludedDeviceResourcesReferencedByTarget == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedDeviceResourcesReferencedByTarget) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedDeviceResourcesReferencedByTarget))
	} else if len(*o.IncludedDeviceResourcesReferencedByTarget) == 1 {
		device = &(*o.IncludedDeviceResourcesReferencedByTarget)[0]
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if o.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *o.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if o.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *o.RevIncludedListResourcesReferencingItem
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if o.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *o.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *o.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingRequest() (orderResponses []OrderResponse, err error) {
	if o.RevIncludedOrderResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *o.RevIncludedOrderResponseResourcesReferencingRequest
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *o.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPlan() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingPlan == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingPlan
	}
	return
}

func (o *OrderPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (o *OrderPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedGroupResourcesReferencedBySubject {
			rsc := (*o.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*o.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedPatientResourcesReferencedBySubject {
			rsc := (*o.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedSubstanceResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedSubstanceResourcesReferencedBySubject {
			rsc := (*o.IncludedSubstanceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *o.IncludedPatientResourcesReferencedByPatient {
			rsc := (*o.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *o.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*o.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedOrganizationResourcesReferencedBySource != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedBySource {
			rsc := (*o.IncludedOrganizationResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPractitionerResourcesReferencedByTarget != nil {
		for idx := range *o.IncludedPractitionerResourcesReferencedByTarget {
			rsc := (*o.IncludedPractitionerResourcesReferencedByTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedOrganizationResourcesReferencedByTarget != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByTarget {
			rsc := (*o.IncludedOrganizationResourcesReferencedByTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedByTarget != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedByTarget {
			rsc := (*o.IncludedDeviceResourcesReferencedByTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *OrderPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	if o.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *o.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*o.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
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
	if o.RevIncludedOrderResponseResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedOrderResponseResourcesReferencingRequest {
			rsc := (*o.RevIncludedOrderResponseResourcesReferencingRequest)[idx]
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
	if o.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
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

func (o *OrderPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedGroupResourcesReferencedBySubject {
			rsc := (*o.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*o.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedPatientResourcesReferencedBySubject {
			rsc := (*o.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedSubstanceResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedSubstanceResourcesReferencedBySubject {
			rsc := (*o.IncludedSubstanceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *o.IncludedPatientResourcesReferencedByPatient {
			rsc := (*o.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *o.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*o.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedOrganizationResourcesReferencedBySource != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedBySource {
			rsc := (*o.IncludedOrganizationResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedPractitionerResourcesReferencedByTarget != nil {
		for idx := range *o.IncludedPractitionerResourcesReferencedByTarget {
			rsc := (*o.IncludedPractitionerResourcesReferencedByTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedOrganizationResourcesReferencedByTarget != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByTarget {
			rsc := (*o.IncludedOrganizationResourcesReferencedByTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedByTarget != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedByTarget {
			rsc := (*o.IncludedDeviceResourcesReferencedByTarget)[idx]
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
	if o.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *o.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*o.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
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
	if o.RevIncludedOrderResponseResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedOrderResponseResourcesReferencingRequest {
			rsc := (*o.RevIncludedOrderResponseResourcesReferencingRequest)[idx]
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
	if o.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
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
