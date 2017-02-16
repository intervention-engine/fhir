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

type Location struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                     `bson:"status,omitempty" json:"status,omitempty"`
	Name                 string                     `bson:"name,omitempty" json:"name,omitempty"`
	Alias                []string                   `bson:"alias,omitempty" json:"alias,omitempty"`
	Description          string                     `bson:"description,omitempty" json:"description,omitempty"`
	Mode                 string                     `bson:"mode,omitempty" json:"mode,omitempty"`
	Type                 *CodeableConcept           `bson:"type,omitempty" json:"type,omitempty"`
	Telecom              []ContactPoint             `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address              *Address                   `bson:"address,omitempty" json:"address,omitempty"`
	PhysicalType         *CodeableConcept           `bson:"physicalType,omitempty" json:"physicalType,omitempty"`
	Position             *LocationPositionComponent `bson:"position,omitempty" json:"position,omitempty"`
	ManagingOrganization *Reference                 `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	PartOf               *Reference                 `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Endpoint             []Reference                `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Location) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Location"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Location), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Location) GetBSON() (interface{}, error) {
	x.ResourceType = "Location"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "location" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type location Location

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Location) UnmarshalJSON(data []byte) (err error) {
	x2 := location{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Location(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Location) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Location"
	} else if x.ResourceType != "Location" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Location, instead received %s", x.ResourceType))
	}
	return nil
}

type LocationPositionComponent struct {
	BackboneElement `bson:",inline"`
	Longitude       *float64 `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude        *float64 `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Altitude        *float64 `bson:"altitude,omitempty" json:"altitude,omitempty"`
}

type LocationPlus struct {
	Location                     `bson:",inline"`
	LocationPlusRelatedResources `bson:",inline"`
}

type LocationPlusRelatedResources struct {
	IncludedLocationResourcesReferencedByPartof                  *[]Location              `bson:"_includedLocationResourcesReferencedByPartof,omitempty"`
	IncludedEndpointResourcesReferencedByEndpoint                *[]Endpoint              `bson:"_includedEndpointResourcesReferencedByEndpoint,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization        *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor              *[]Appointment           `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingLocation           *[]Appointment           `bson:"_revIncludedAppointmentResourcesReferencingLocation,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                *[]Account               `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                   *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref   *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedPractitionerRoleResourcesReferencingLocation      *[]PractitionerRole      `bson:"_revIncludedPractitionerRoleResourcesReferencingLocation,omitempty"`
	RevIncludedPractitionerResourcesReferencingLocation          *[]Practitioner          `bson:"_revIncludedPractitionerResourcesReferencingLocation,omitempty"`
	RevIncludedContractResourcesReferencingAgent                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingAgent,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject               *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingDomain                *[]Contract              `bson:"_revIncludedContractResourcesReferencingDomain,omitempty"`
	RevIncludedContractResourcesReferencingTopic                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest          *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource   *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedEncounterResourcesReferencingLocation             *[]Encounter             `bson:"_revIncludedEncounterResourcesReferencingLocation,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon          *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData             *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingLocation            *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingLocation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                   *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                     *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                   *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingFacility  *[]ExplanationOfBenefit  `bson:"_revIncludedExplanationOfBenefitResourcesReferencingFacility,omitempty"`
	RevIncludedResearchStudyResourcesReferencingSite             *[]ResearchStudy         `bson:"_revIncludedResearchStudyResourcesReferencingSite,omitempty"`
	RevIncludedProcedureResourcesReferencingLocation             *[]Procedure             `bson:"_revIncludedProcedureResourcesReferencingLocation,omitempty"`
	RevIncludedListResourcesReferencingItem                      *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                   *[]List                  `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingSubject      *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon      *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition   *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedImmunizationResourcesReferencingLocation          *[]Immunization          `bson:"_revIncludedImmunizationResourcesReferencingLocation,omitempty"`
	RevIncludedDeviceResourcesReferencingLocation                *[]Device                `bson:"_revIncludedDeviceResourcesReferencingLocation,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingSubject       *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon       *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition    *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                   *[]Flag                  `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor      *[]AppointmentResponse   `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingLocation   *[]AppointmentResponse   `bson:"_revIncludedAppointmentResponseResourcesReferencingLocation,omitempty"`
	RevIncludedObservationResourcesReferencingSubject            *[]Observation           `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                  *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedEligibilityRequestResourcesReferencingFacility    *[]EligibilityRequest    `bson:"_revIncludedEligibilityRequestResourcesReferencingFacility,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingDestination *[]MedicationDispense    `bson:"_revIncludedMedicationDispenseResourcesReferencingDestination,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject       *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedHealthcareServiceResourcesReferencingLocation     *[]HealthcareService     `bson:"_revIncludedHealthcareServiceResourcesReferencingLocation,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity              *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject            *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated       *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject  *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest        *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                 *[]Schedule              `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedClaimResourcesReferencingFacility                 *[]Claim                 `bson:"_revIncludedClaimResourcesReferencingFacility,omitempty"`
	RevIncludedLocationResourcesReferencingPartof                *[]Location              `bson:"_revIncludedLocationResourcesReferencingPartof,omitempty"`
}

func (l *LocationPlusRelatedResources) GetIncludedLocationResourceReferencedByPartof() (location *Location, err error) {
	if l.IncludedLocationResourcesReferencedByPartof == nil {
		err = errors.New("Included locations not requested")
	} else if len(*l.IncludedLocationResourcesReferencedByPartof) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*l.IncludedLocationResourcesReferencedByPartof))
	} else if len(*l.IncludedLocationResourcesReferencedByPartof) == 1 {
		location = &(*l.IncludedLocationResourcesReferencedByPartof)[0]
	}
	return
}

func (l *LocationPlusRelatedResources) GetIncludedEndpointResourcesReferencedByEndpoint() (endpoints []Endpoint, err error) {
	if l.IncludedEndpointResourcesReferencedByEndpoint == nil {
		err = errors.New("Included endpoints not requested")
	} else {
		endpoints = *l.IncludedEndpointResourcesReferencedByEndpoint
	}
	return
}

func (l *LocationPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if l.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*l.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*l.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*l.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*l.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if l.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *l.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingLocation() (appointments []Appointment, err error) {
	if l.RevIncludedAppointmentResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *l.RevIncludedAppointmentResourcesReferencingLocation
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if l.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *l.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if l.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *l.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedPractitionerRoleResourcesReferencingLocation() (practitionerRoles []PractitionerRole, err error) {
	if l.RevIncludedPractitionerRoleResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded practitionerRoles not requested")
	} else {
		practitionerRoles = *l.RevIncludedPractitionerRoleResourcesReferencingLocation
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedPractitionerResourcesReferencingLocation() (practitioners []Practitioner, err error) {
	if l.RevIncludedPractitionerResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded practitioners not requested")
	} else {
		practitioners = *l.RevIncludedPractitionerResourcesReferencingLocation
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedContractResourcesReferencingAgent() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingAgent
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedContractResourcesReferencingDomain() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingDomain == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingDomain
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *l.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *l.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if l.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *l.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingLocation() (encounters []Encounter, err error) {
	if l.RevIncludedEncounterResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *l.RevIncludedEncounterResourcesReferencingLocation
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if l.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *l.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if l.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *l.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingLocation() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingLocation
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingFacility() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if l.RevIncludedExplanationOfBenefitResourcesReferencingFacility == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *l.RevIncludedExplanationOfBenefitResourcesReferencingFacility
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedResearchStudyResourcesReferencingSite() (researchStudies []ResearchStudy, err error) {
	if l.RevIncludedResearchStudyResourcesReferencingSite == nil {
		err = errors.New("RevIncluded researchStudies not requested")
	} else {
		researchStudies = *l.RevIncludedResearchStudyResourcesReferencingSite
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingLocation() (procedures []Procedure, err error) {
	if l.RevIncludedProcedureResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *l.RevIncludedProcedureResourcesReferencingLocation
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if l.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *l.RevIncludedListResourcesReferencingItem
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedListResourcesReferencingSubject() (lists []List, err error) {
	if l.RevIncludedListResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *l.RevIncludedListResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if l.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *l.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingSubject() (diagnosticRequests []DiagnosticRequest, err error) {
	if l.RevIncludedDiagnosticRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *l.RevIncludedDiagnosticRequestResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if l.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *l.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if l.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *l.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingLocation() (immunizations []Immunization, err error) {
	if l.RevIncludedImmunizationResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *l.RevIncludedImmunizationResourcesReferencingLocation
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDeviceResourcesReferencingLocation() (devices []Device, err error) {
	if l.RevIncludedDeviceResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded devices not requested")
	} else {
		devices = *l.RevIncludedDeviceResourcesReferencingLocation
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if l.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *l.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingSubject() (deviceUseRequests []DeviceUseRequest, err error) {
	if l.RevIncludedDeviceUseRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *l.RevIncludedDeviceUseRequestResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if l.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *l.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if l.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *l.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if l.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *l.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if l.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *l.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingLocation() (appointmentResponses []AppointmentResponse, err error) {
	if l.RevIncludedAppointmentResponseResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *l.RevIncludedAppointmentResponseResourcesReferencingLocation
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingSubject() (observations []Observation, err error) {
	if l.RevIncludedObservationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *l.RevIncludedObservationResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if l.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *l.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedEligibilityRequestResourcesReferencingFacility() (eligibilityRequests []EligibilityRequest, err error) {
	if l.RevIncludedEligibilityRequestResourcesReferencingFacility == nil {
		err = errors.New("RevIncluded eligibilityRequests not requested")
	} else {
		eligibilityRequests = *l.RevIncludedEligibilityRequestResourcesReferencingFacility
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingDestination() (medicationDispenses []MedicationDispense, err error) {
	if l.RevIncludedMedicationDispenseResourcesReferencingDestination == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *l.RevIncludedMedicationDispenseResourcesReferencingDestination
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingSubject() (diagnosticReports []DiagnosticReport, err error) {
	if l.RevIncludedDiagnosticReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *l.RevIncludedDiagnosticReportResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedHealthcareServiceResourcesReferencingLocation() (healthcareServices []HealthcareService, err error) {
	if l.RevIncludedHealthcareServiceResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded healthcareServices not requested")
	} else {
		healthcareServices = *l.RevIncludedHealthcareServiceResourcesReferencingLocation
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if l.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *l.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *l.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if l.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *l.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if l.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *l.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedClaimResourcesReferencingFacility() (claims []Claim, err error) {
	if l.RevIncludedClaimResourcesReferencingFacility == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *l.RevIncludedClaimResourcesReferencingFacility
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedLocationResourcesReferencingPartof() (locations []Location, err error) {
	if l.RevIncludedLocationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded locations not requested")
	} else {
		locations = *l.RevIncludedLocationResourcesReferencingPartof
	}
	return
}

func (l *LocationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedLocationResourcesReferencedByPartof != nil {
		for idx := range *l.IncludedLocationResourcesReferencedByPartof {
			rsc := (*l.IncludedLocationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *l.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*l.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *l.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*l.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (l *LocationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAppointmentResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingLocation {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*l.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *l.RevIncludedConsentResourcesReferencingData {
			rsc := (*l.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPractitionerRoleResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedPractitionerRoleResourcesReferencingLocation {
			rsc := (*l.RevIncludedPractitionerRoleResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPractitionerResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedPractitionerResourcesReferencingLocation {
			rsc := (*l.RevIncludedPractitionerResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingAgent {
			rsc := (*l.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*l.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingSubject {
			rsc := (*l.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingDomain != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingDomain {
			rsc := (*l.RevIncludedContractResourcesReferencingDomain)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingTopic {
			rsc := (*l.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *l.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*l.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEncounterResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedEncounterResourcesReferencingLocation {
			rsc := (*l.RevIncludedEncounterResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *l.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*l.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingLocation {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*l.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*l.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*l.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedExplanationOfBenefitResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedExplanationOfBenefitResourcesReferencingFacility {
			rsc := (*l.RevIncludedExplanationOfBenefitResourcesReferencingFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchStudyResourcesReferencingSite != nil {
		for idx := range *l.RevIncludedResearchStudyResourcesReferencingSite {
			rsc := (*l.RevIncludedResearchStudyResourcesReferencingSite)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProcedureResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedProcedureResourcesReferencingLocation {
			rsc := (*l.RevIncludedProcedureResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedListResourcesReferencingItem {
			rsc := (*l.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedListResourcesReferencingSubject {
			rsc := (*l.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingSubject {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedImmunizationResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedImmunizationResourcesReferencingLocation {
			rsc := (*l.RevIncludedImmunizationResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedDeviceResourcesReferencingLocation {
			rsc := (*l.RevIncludedDeviceResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingSubject {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*l.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*l.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAppointmentResponseResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedAppointmentResponseResourcesReferencingLocation {
			rsc := (*l.RevIncludedAppointmentResponseResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*l.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*l.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEligibilityRequestResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedEligibilityRequestResourcesReferencingFacility {
			rsc := (*l.RevIncludedEligibilityRequestResourcesReferencingFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMedicationDispenseResourcesReferencingDestination != nil {
		for idx := range *l.RevIncludedMedicationDispenseResourcesReferencingDestination {
			rsc := (*l.RevIncludedMedicationDispenseResourcesReferencingDestination)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*l.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedHealthcareServiceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedHealthcareServiceResourcesReferencingLocation {
			rsc := (*l.RevIncludedHealthcareServiceResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*l.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*l.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*l.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*l.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*l.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*l.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*l.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedClaimResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedClaimResourcesReferencingFacility {
			rsc := (*l.RevIncludedClaimResourcesReferencingFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLocationResourcesReferencingPartof != nil {
		for idx := range *l.RevIncludedLocationResourcesReferencingPartof {
			rsc := (*l.RevIncludedLocationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (l *LocationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedLocationResourcesReferencedByPartof != nil {
		for idx := range *l.IncludedLocationResourcesReferencedByPartof {
			rsc := (*l.IncludedLocationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *l.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*l.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *l.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*l.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAppointmentResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingLocation {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*l.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *l.RevIncludedConsentResourcesReferencingData {
			rsc := (*l.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPractitionerRoleResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedPractitionerRoleResourcesReferencingLocation {
			rsc := (*l.RevIncludedPractitionerRoleResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPractitionerResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedPractitionerResourcesReferencingLocation {
			rsc := (*l.RevIncludedPractitionerResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingAgent {
			rsc := (*l.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*l.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingSubject {
			rsc := (*l.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingDomain != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingDomain {
			rsc := (*l.RevIncludedContractResourcesReferencingDomain)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingTopic {
			rsc := (*l.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *l.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*l.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEncounterResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedEncounterResourcesReferencingLocation {
			rsc := (*l.RevIncludedEncounterResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *l.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*l.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingLocation {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*l.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*l.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*l.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedExplanationOfBenefitResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedExplanationOfBenefitResourcesReferencingFacility {
			rsc := (*l.RevIncludedExplanationOfBenefitResourcesReferencingFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchStudyResourcesReferencingSite != nil {
		for idx := range *l.RevIncludedResearchStudyResourcesReferencingSite {
			rsc := (*l.RevIncludedResearchStudyResourcesReferencingSite)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProcedureResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedProcedureResourcesReferencingLocation {
			rsc := (*l.RevIncludedProcedureResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedListResourcesReferencingItem {
			rsc := (*l.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedListResourcesReferencingSubject {
			rsc := (*l.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingSubject {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *l.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*l.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedImmunizationResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedImmunizationResourcesReferencingLocation {
			rsc := (*l.RevIncludedImmunizationResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedDeviceResourcesReferencingLocation {
			rsc := (*l.RevIncludedDeviceResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingSubject {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *l.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*l.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*l.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*l.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAppointmentResponseResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedAppointmentResponseResourcesReferencingLocation {
			rsc := (*l.RevIncludedAppointmentResponseResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*l.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*l.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEligibilityRequestResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedEligibilityRequestResourcesReferencingFacility {
			rsc := (*l.RevIncludedEligibilityRequestResourcesReferencingFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMedicationDispenseResourcesReferencingDestination != nil {
		for idx := range *l.RevIncludedMedicationDispenseResourcesReferencingDestination {
			rsc := (*l.RevIncludedMedicationDispenseResourcesReferencingDestination)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*l.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedHealthcareServiceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedHealthcareServiceResourcesReferencingLocation {
			rsc := (*l.RevIncludedHealthcareServiceResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*l.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*l.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*l.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*l.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*l.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*l.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*l.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedClaimResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedClaimResourcesReferencingFacility {
			rsc := (*l.RevIncludedClaimResourcesReferencingFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLocationResourcesReferencingPartof != nil {
		for idx := range *l.RevIncludedLocationResourcesReferencingPartof {
			rsc := (*l.RevIncludedLocationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
