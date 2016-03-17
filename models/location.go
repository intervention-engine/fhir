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

type Location struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                     `bson:"status,omitempty" json:"status,omitempty"`
	Name                 string                     `bson:"name,omitempty" json:"name,omitempty"`
	Description          string                     `bson:"description,omitempty" json:"description,omitempty"`
	Mode                 string                     `bson:"mode,omitempty" json:"mode,omitempty"`
	Type                 *CodeableConcept           `bson:"type,omitempty" json:"type,omitempty"`
	Telecom              []ContactPoint             `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address              *Address                   `bson:"address,omitempty" json:"address,omitempty"`
	PhysicalType         *CodeableConcept           `bson:"physicalType,omitempty" json:"physicalType,omitempty"`
	Position             *LocationPositionComponent `bson:"position,omitempty" json:"position,omitempty"`
	ManagingOrganization *Reference                 `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	PartOf               *Reference                 `bson:"partOf,omitempty" json:"partOf,omitempty"`
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
	IncludedOrganizationResourcesReferencedByOrganization        *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor              *[]Appointment           `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingLocation           *[]Appointment           `bson:"_revIncludedAppointmentResourcesReferencingLocation,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                *[]Account               `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedProvenanceResourcesReferencingLocation            *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingLocation,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedProcedureResourcesReferencingLocation             *[]Procedure             `bson:"_revIncludedProcedureResourcesReferencingLocation,omitempty"`
	RevIncludedListResourcesReferencingItem                      *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                   *[]List                  `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref   *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                   *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedImmunizationResourcesReferencingLocation          *[]Immunization          `bson:"_revIncludedImmunizationResourcesReferencingLocation,omitempty"`
	RevIncludedDeviceResourcesReferencingLocation                *[]Device                `bson:"_revIncludedDeviceResourcesReferencingLocation,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                   *[]Flag                  `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedPractitionerResourcesReferencingLocation          *[]Practitioner          `bson:"_revIncludedPractitionerResourcesReferencingLocation,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor      *[]AppointmentResponse   `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingLocation   *[]AppointmentResponse   `bson:"_revIncludedAppointmentResponseResourcesReferencingLocation,omitempty"`
	RevIncludedObservationResourcesReferencingSubject            *[]Observation           `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingActor                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingActor,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                  *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingDestination *[]MedicationDispense    `bson:"_revIncludedMedicationDispenseResourcesReferencingDestination,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject       *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedHealthcareServiceResourcesReferencingLocation     *[]HealthcareService     `bson:"_revIncludedHealthcareServiceResourcesReferencingLocation,omitempty"`
	RevIncludedEncounterResourcesReferencingLocation             *[]Encounter             `bson:"_revIncludedEncounterResourcesReferencingLocation,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference           *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject            *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated       *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingSubject        *[]DiagnosticOrder       `bson:"_revIncludedDiagnosticOrderResourcesReferencingSubject,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment      *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject  *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest        *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                 *[]Schedule              `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger     *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData             *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
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

func (l *LocationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingLocation() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingLocation
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

func (l *LocationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if l.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *l.RevIncludedOrderResourcesReferencingDetail
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

func (l *LocationPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if l.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *l.RevIncludedFlagResourcesReferencingSubject
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

func (l *LocationPlusRelatedResources) GetRevIncludedContractResourcesReferencingActor() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingActor == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingActor
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

func (l *LocationPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingLocation() (encounters []Encounter, err error) {
	if l.RevIncludedEncounterResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *l.RevIncludedEncounterResourcesReferencingLocation
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if l.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *l.RevIncludedAuditEventResourcesReferencingReference
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

func (l *LocationPlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingSubject() (diagnosticOrders []DiagnosticOrder, err error) {
	if l.RevIncludedDiagnosticOrderResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *l.RevIncludedDiagnosticOrderResourcesReferencingSubject
	}
	return
}

func (l *LocationPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if l.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *l.RevIncludedOrderResponseResourcesReferencingFulfillment
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

func (l *LocationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if l.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *l.RevIncludedClinicalImpressionResourcesReferencingTrigger
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
	if l.RevIncludedProvenanceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingLocation {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *l.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*l.RevIncludedOrderResourcesReferencingDetail)[idx]
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
	if l.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*l.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPractitionerResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedPractitionerResourcesReferencingLocation {
			rsc := (*l.RevIncludedPractitionerResourcesReferencingLocation)[idx]
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
	if l.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingActor {
			rsc := (*l.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*l.RevIncludedBasicResourcesReferencingSubject)[idx]
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
	if l.RevIncludedEncounterResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedEncounterResourcesReferencingLocation {
			rsc := (*l.RevIncludedEncounterResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *l.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*l.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if l.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			rsc := (*l.RevIncludedDiagnosticOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *l.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*l.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if l.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *l.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*l.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *l.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*l.RevIncludedMessageHeaderResourcesReferencingData)[idx]
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
	if l.RevIncludedProvenanceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingLocation {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *l.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*l.RevIncludedOrderResourcesReferencingDetail)[idx]
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
	if l.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*l.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPractitionerResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedPractitionerResourcesReferencingLocation {
			rsc := (*l.RevIncludedPractitionerResourcesReferencingLocation)[idx]
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
	if l.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingActor {
			rsc := (*l.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*l.RevIncludedBasicResourcesReferencingSubject)[idx]
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
	if l.RevIncludedEncounterResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedEncounterResourcesReferencingLocation {
			rsc := (*l.RevIncludedEncounterResourcesReferencingLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *l.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*l.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if l.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			rsc := (*l.RevIncludedDiagnosticOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *l.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*l.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if l.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *l.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*l.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *l.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*l.RevIncludedMessageHeaderResourcesReferencingData)[idx]
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
