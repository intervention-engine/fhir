// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
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
	Active                 *bool                                     `bson:"active,omitempty" json:"active,omitempty"`
	ProvidedBy             *Reference                                `bson:"providedBy,omitempty" json:"providedBy,omitempty"`
	Category               *CodeableConcept                          `bson:"category,omitempty" json:"category,omitempty"`
	Type                   []CodeableConcept                         `bson:"type,omitempty" json:"type,omitempty"`
	Specialty              []CodeableConcept                         `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Location               []Reference                               `bson:"location,omitempty" json:"location,omitempty"`
	Name                   string                                    `bson:"name,omitempty" json:"name,omitempty"`
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
	AppointmentRequired    *bool                                     `bson:"appointmentRequired,omitempty" json:"appointmentRequired,omitempty"`
	AvailableTime          []HealthcareServiceAvailableTimeComponent `bson:"availableTime,omitempty" json:"availableTime,omitempty"`
	NotAvailable           []HealthcareServiceNotAvailableComponent  `bson:"notAvailable,omitempty" json:"notAvailable,omitempty"`
	AvailabilityExceptions string                                    `bson:"availabilityExceptions,omitempty" json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                               `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
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
	IncludedEndpointResourcesReferencedByEndpoint                   *[]Endpoint              `bson:"_includedEndpointResourcesReferencedByEndpoint,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization           *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedLocationResourcesReferencedByLocation                   *[]Location              `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                 *[]Appointment           `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRecipient         *[]ReferralRequest       `bson:"_revIncludedReferralRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                   *[]Account               `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                 *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1            *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2            *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref      *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedPractitionerRoleResourcesReferencingService          *[]PractitionerRole      `bson:"_revIncludedPractitionerRoleResourcesReferencingService,omitempty"`
	RevIncludedContractResourcesReferencingSubject                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest             *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse            *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource      *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor       *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom     *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor     *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof      *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson       *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof              *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon             *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor      *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom    *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor    *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof     *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition           *[]RequestGroup          `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPerformer           *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon             *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest        *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                 *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                        *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedProcedureResourcesReferencingDefinition              *[]Procedure             `bson:"_revIncludedProcedureResourcesReferencingDefinition,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingPerformer        *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces         *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon          *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor         *[]AppointmentResponse   `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom               *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor               *[]Library               `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                *[]Library               `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon      *[]CommunicationRequest  `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                     *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                 *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail          *[]Condition             `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject               *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                 *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated          *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject     *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest           *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                    *[]Schedule              `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor          *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof         *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedEndpointResourcesReferencedByEndpoint() (endpoints []Endpoint, err error) {
	if h.IncludedEndpointResourcesReferencedByEndpoint == nil {
		err = errors.New("Included endpoints not requested")
	} else {
		endpoints = *h.IncludedEndpointResourcesReferencedByEndpoint
	}
	return
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

func (h *HealthcareServicePlusRelatedResources) GetIncludedLocationResourcesReferencedByLocation() (locations []Location, err error) {
	if h.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else {
		locations = *h.IncludedLocationResourcesReferencedByLocation
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

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingRecipient() (referralRequests []ReferralRequest, err error) {
	if h.RevIncludedReferralRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *h.RevIncludedReferralRequestResourcesReferencingRecipient
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

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if h.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *h.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if h.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *h.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingDependsonPath2
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

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPractitionerRoleResourcesReferencingService() (practitionerRoles []PractitionerRole, err error) {
	if h.RevIncludedPractitionerRoleResourcesReferencingService == nil {
		err = errors.New("RevIncluded practitionerRoles not requested")
	} else {
		practitionerRoles = *h.RevIncludedPractitionerRoleResourcesReferencingService
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if h.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *h.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if h.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *h.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if h.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *h.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if h.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *h.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if h.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *h.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if h.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *h.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if h.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *h.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if h.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *h.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if h.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *h.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if h.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *h.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if h.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *h.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if h.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *h.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if h.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *h.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPerformer() (deviceRequests []DeviceRequest, err error) {
	if h.RevIncludedDeviceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *h.RevIncludedDeviceRequestResourcesReferencingPerformer
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if h.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *h.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if h.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *h.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if h.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *h.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if h.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *h.RevIncludedProvenanceResourcesReferencingEntityref
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

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if h.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *h.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if h.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *h.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if h.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *h.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedProcedureResourcesReferencingDefinition() (procedures []Procedure, err error) {
	if h.RevIncludedProcedureResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *h.RevIncludedProcedureResourcesReferencingDefinition
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

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingPerformer() (procedureRequests []ProcedureRequest, err error) {
	if h.RevIncludedProcedureRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *h.RevIncludedProcedureRequestResourcesReferencingPerformer
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if h.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *h.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if h.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *h.RevIncludedProcedureRequestResourcesReferencingBasedon
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

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if h.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *h.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if h.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *h.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if h.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *h.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if h.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *h.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if h.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *h.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if h.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *h.RevIncludedCommunicationRequestResourcesReferencingBasedon
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

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if h.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *h.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if h.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *h.RevIncludedConditionResourcesReferencingEvidencedetail
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

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if h.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *h.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*h.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
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
	if h.RevIncludedReferralRequestResourcesReferencingRecipient != nil {
		for idx := range *h.RevIncludedReferralRequestResourcesReferencingRecipient {
			rsc := (*h.RevIncludedReferralRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*h.RevIncludedAccountResourcesReferencingSubject)[idx]
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
	if h.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *h.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*h.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *h.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*h.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*h.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *h.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*h.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPractitionerRoleResourcesReferencingService != nil {
		for idx := range *h.RevIncludedPractitionerRoleResourcesReferencingService {
			rsc := (*h.RevIncludedPractitionerRoleResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedContractResourcesReferencingSubject {
			rsc := (*h.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *h.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*h.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *h.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*h.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *h.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*h.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *h.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*h.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*h.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *h.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*h.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *h.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*h.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *h.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*h.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *h.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*h.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*h.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*h.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*h.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProcedureResourcesReferencingDefinition != nil {
		for idx := range *h.RevIncludedProcedureResourcesReferencingDefinition {
			rsc := (*h.RevIncludedProcedureResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *h.RevIncludedListResourcesReferencingItem {
			rsc := (*h.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *h.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*h.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *h.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*h.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*h.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*h.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*h.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*h.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*h.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*h.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *h.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*h.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *h.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*h.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if h.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if h.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *h.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*h.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
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
	if h.RevIncludedReferralRequestResourcesReferencingRecipient != nil {
		for idx := range *h.RevIncludedReferralRequestResourcesReferencingRecipient {
			rsc := (*h.RevIncludedReferralRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*h.RevIncludedAccountResourcesReferencingSubject)[idx]
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
	if h.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *h.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*h.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *h.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*h.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*h.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *h.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*h.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPractitionerRoleResourcesReferencingService != nil {
		for idx := range *h.RevIncludedPractitionerRoleResourcesReferencingService {
			rsc := (*h.RevIncludedPractitionerRoleResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedContractResourcesReferencingSubject {
			rsc := (*h.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *h.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*h.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *h.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*h.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *h.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*h.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *h.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*h.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*h.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *h.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*h.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *h.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*h.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *h.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*h.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *h.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*h.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*h.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*h.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*h.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProcedureResourcesReferencingDefinition != nil {
		for idx := range *h.RevIncludedProcedureResourcesReferencingDefinition {
			rsc := (*h.RevIncludedProcedureResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *h.RevIncludedListResourcesReferencingItem {
			rsc := (*h.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *h.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*h.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *h.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*h.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*h.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*h.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*h.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*h.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*h.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*h.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *h.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*h.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *h.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*h.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if h.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
