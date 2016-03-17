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

type HealthcareService struct {
	DomainResource         `bson:",inline"`
	Identifier             []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ProvidedBy             *Reference                                `bson:"providedBy,omitempty" json:"providedBy,omitempty"`
	ServiceCategory        *CodeableConcept                          `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType            []HealthcareServiceServiceTypeComponent   `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Location               *Reference                                `bson:"location,omitempty" json:"location,omitempty"`
	ServiceName            string                                    `bson:"serviceName,omitempty" json:"serviceName,omitempty"`
	Comment                string                                    `bson:"comment,omitempty" json:"comment,omitempty"`
	ExtraDetails           string                                    `bson:"extraDetails,omitempty" json:"extraDetails,omitempty"`
	Photo                  *Attachment                               `bson:"photo,omitempty" json:"photo,omitempty"`
	Telecom                []ContactPoint                            `bson:"telecom,omitempty" json:"telecom,omitempty"`
	CoverageArea           []Reference                               `bson:"coverageArea,omitempty" json:"coverageArea,omitempty"`
	ServiceProvisionCode   []CodeableConcept                         `bson:"serviceProvisionCode,omitempty" json:"serviceProvisionCode,omitempty"`
	Eligibility            *CodeableConcept                          `bson:"eligibility,omitempty" json:"eligibility,omitempty"`
	EligibilityNote        string                                    `bson:"eligibilityNote,omitempty" json:"eligibilityNote,omitempty"`
	ProgramName            []string                                  `bson:"programName,omitempty" json:"programName,omitempty"`
	Characteristic         []CodeableConcept                         `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	ReferralMethod         []CodeableConcept                         `bson:"referralMethod,omitempty" json:"referralMethod,omitempty"`
	PublicKey              string                                    `bson:"publicKey,omitempty" json:"publicKey,omitempty"`
	AppointmentRequired    *bool                                     `bson:"appointmentRequired,omitempty" json:"appointmentRequired,omitempty"`
	AvailableTime          []HealthcareServiceAvailableTimeComponent `bson:"availableTime,omitempty" json:"availableTime,omitempty"`
	NotAvailable           []HealthcareServiceNotAvailableComponent  `bson:"notAvailable,omitempty" json:"notAvailable,omitempty"`
	AvailabilityExceptions string                                    `bson:"availabilityExceptions,omitempty" json:"availabilityExceptions,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *HealthcareService) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "HealthcareService"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to HealthcareService), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *HealthcareService) GetBSON() (interface{}, error) {
	x.ResourceType = "HealthcareService"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "healthcareService" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type healthcareService HealthcareService

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *HealthcareService) UnmarshalJSON(data []byte) (err error) {
	x2 := healthcareService{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = HealthcareService(x2)
		return x.checkResourceType()
	}
	return
}

func (x *HealthcareService) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "HealthcareService"
	} else if x.ResourceType != "HealthcareService" {
		return errors.New(fmt.Sprintf("Expected resourceType to be HealthcareService, instead received %s", x.ResourceType))
	}
	return nil
}

type HealthcareServiceServiceTypeComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Specialty       []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
}

type HealthcareServiceAvailableTimeComponent struct {
	BackboneElement    `bson:",inline"`
	DaysOfWeek         []string      `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty"`
	AllDay             *bool         `bson:"allDay,omitempty" json:"allDay,omitempty"`
	AvailableStartTime *FHIRDateTime `bson:"availableStartTime,omitempty" json:"availableStartTime,omitempty"`
	AvailableEndTime   *FHIRDateTime `bson:"availableEndTime,omitempty" json:"availableEndTime,omitempty"`
}

type HealthcareServiceNotAvailableComponent struct {
	BackboneElement `bson:",inline"`
	Description     string  `bson:"description,omitempty" json:"description,omitempty"`
	During          *Period `bson:"during,omitempty" json:"during,omitempty"`
}

type HealthcareServicePlus struct {
	HealthcareService                     `bson:",inline"`
	HealthcareServicePlusRelatedResources `bson:",inline"`
}

type HealthcareServicePlusRelatedResources struct {
	IncludedOrganizationResourcesReferencedByOrganization       *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedLocationResourcesReferencedByLocation               *[]Location              `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor             *[]Appointment           `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAccountResourcesReferencingSubject               *[]Account               `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor     *[]AppointmentResponse   `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                *[]Schedule              `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if h.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*h.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*h.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*h.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*h.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if h.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*h.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*h.IncludedLocationResourcesReferencedByLocation))
	} else if len(*h.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*h.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if h.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *h.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if h.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *h.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if h.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *h.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if h.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *h.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if h.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *h.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if h.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *h.RevIncludedListResourcesReferencingItem
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if h.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *h.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if h.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *h.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if h.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *h.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if h.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *h.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if h.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *h.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if h.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *h.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if h.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *h.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if h.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *h.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if h.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *h.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if h.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *h.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if h.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *h.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if h.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *h.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if h.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *h.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if h.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *h.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if h.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *h.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*h.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *h.IncludedLocationResourcesReferencedByLocation {
			rsc := (*h.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if h.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*h.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*h.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *h.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*h.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *h.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*h.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *h.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*h.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *h.RevIncludedListResourcesReferencingItem {
			rsc := (*h.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *h.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*h.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *h.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*h.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*h.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*h.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *h.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*h.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*h.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *h.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*h.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *h.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*h.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *h.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*h.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*h.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *h.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*h.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*h.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *h.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*h.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *h.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*h.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if h.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *h.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*h.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *h.IncludedLocationResourcesReferencedByLocation {
			rsc := (*h.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*h.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*h.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *h.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*h.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *h.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*h.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *h.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*h.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *h.RevIncludedListResourcesReferencingItem {
			rsc := (*h.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *h.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*h.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *h.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*h.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*h.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*h.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *h.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*h.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*h.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *h.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*h.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *h.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*h.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *h.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*h.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*h.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *h.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*h.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*h.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *h.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*h.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *h.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*h.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
