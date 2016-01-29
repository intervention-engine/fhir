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

type Device struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Note            []Annotation     `bson:"note,omitempty" json:"note,omitempty"`
	Status          string           `bson:"status,omitempty" json:"status,omitempty"`
	Manufacturer    string           `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Model           string           `bson:"model,omitempty" json:"model,omitempty"`
	Version         string           `bson:"version,omitempty" json:"version,omitempty"`
	ManufactureDate *FHIRDateTime    `bson:"manufactureDate,omitempty" json:"manufactureDate,omitempty"`
	Expiry          *FHIRDateTime    `bson:"expiry,omitempty" json:"expiry,omitempty"`
	Udi             string           `bson:"udi,omitempty" json:"udi,omitempty"`
	LotNumber       string           `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	Owner           *Reference       `bson:"owner,omitempty" json:"owner,omitempty"`
	Location        *Reference       `bson:"location,omitempty" json:"location,omitempty"`
	Patient         *Reference       `bson:"patient,omitempty" json:"patient,omitempty"`
	Contact         []ContactPoint   `bson:"contact,omitempty" json:"contact,omitempty"`
	Url             string           `bson:"url,omitempty" json:"url,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Device) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Device"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Device), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Device) GetBSON() (interface{}, error) {
	x.ResourceType = "Device"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "device" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type device Device

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Device) UnmarshalJSON(data []byte) (err error) {
	x2 := device{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Device(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Device) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Device"
	} else if x.ResourceType != "Device" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Device, instead received %s", x.ResourceType))
	}
	return nil
}

type DevicePlus struct {
	Device                     `bson:",inline"`
	DevicePlusRelatedResources `bson:",inline"`
}

type DevicePlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                   *[]Patient                  `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization         *[]Organization             `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedLocationResourcesReferencedByLocation                 *[]Location                 `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor               *[]Appointment              `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                 *[]Account                  `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget               *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref     *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingSubject        *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor         *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref     *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedSpecimenResourcesReferencingSubject                *[]Specimen                 `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingItem                       *[]List                     `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                    *[]List                     `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingSource                     *[]List                     `bson:"_revIncludedListResourcesReferencingSource,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingSubject       *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor        *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref    *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingSubject                   *[]Order                    `bson:"_revIncludedOrderResourcesReferencingSubject,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                    *[]Order                    `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedOrderResourcesReferencingTarget                    *[]Order                    `bson:"_revIncludedOrderResourcesReferencingTarget,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                   *[]Media                    `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingOrderer        *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingOrderer,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDevice         *[]DeviceUseRequest         `bson:"_revIncludedDeviceUseRequestResourcesReferencingDevice,omitempty"`
	RevIncludedDeviceMetricResourcesReferencingSource             *[]DeviceMetric             `bson:"_revIncludedDeviceMetricResourcesReferencingSource,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                     *[]Flag                     `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor       *[]AppointmentResponse      `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedObservationResourcesReferencingSubject             *[]Observation              `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingDevice              *[]Observation              `bson:"_revIncludedObservationResourcesReferencingDevice,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingDevice *[]MedicationAdministration `bson:"_revIncludedMedicationAdministrationResourcesReferencingDevice,omitempty"`
	RevIncludedContractResourcesReferencingActor                  *[]Contract                 `bson:"_revIncludedContractResourcesReferencingActor,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender     *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient  *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingPerformer        *[]RiskAssessment           `bson:"_revIncludedRiskAssessmentResourcesReferencingPerformer,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                   *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedGroupResourcesReferencingMember                    *[]Group                    `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject        *[]DiagnosticReport         `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedImagingObjectSelectionResourcesReferencingAuthor   *[]ImagingObjectSelection   `bson:"_revIncludedImagingObjectSelectionResourcesReferencingAuthor,omitempty"`
	RevIncludedDeviceComponentResourcesReferencingSource          *[]DeviceComponent          `bson:"_revIncludedDeviceComponentResourcesReferencingSource,omitempty"`
	RevIncludedAuditEventResourcesReferencingParticipant          *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingParticipant,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference            *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender            *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient         *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject             *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor              *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry               *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingAuthor            *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingAuthor,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated        *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingSubject         *[]DiagnosticOrder          `bson:"_revIncludedDiagnosticOrderResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingActorPath1      *[]DiagnosticOrder          `bson:"_revIncludedDiagnosticOrderResourcesReferencingActorPath1,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingActorPath2      *[]DiagnosticOrder          `bson:"_revIncludedDiagnosticOrderResourcesReferencingActorPath2,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment       *[]OrderResponse            `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedOrderResponseResourcesReferencingWho               *[]OrderResponse            `bson:"_revIncludedOrderResponseResourcesReferencingWho,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject   *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor    *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedDeviceUseStatementResourcesReferencingDevice       *[]DeviceUseStatement       `bson:"_revIncludedDeviceUseStatementResourcesReferencingDevice,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest         *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                  *[]Schedule                 `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger      *[]ClinicalImpression       `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData              *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingTarget            *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingTarget,omitempty"`
}

func (d *DevicePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedByPatient))
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (d *DevicePlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if d.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*d.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*d.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*d.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*d.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (d *DevicePlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if d.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*d.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*d.IncludedLocationResourcesReferencedByLocation))
	} else if len(*d.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*d.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if d.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *d.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if d.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *d.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingSubject() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingSubject() (specimen []Specimen, err error) {
	if d.RevIncludedSpecimenResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *d.RevIncludedSpecimenResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedListResourcesReferencingSubject() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedListResourcesReferencingSource() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingSource == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingSource
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingSubject() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedOrderResourcesReferencingSubject() (orders []Order, err error) {
	if d.RevIncludedOrderResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *d.RevIncludedOrderResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if d.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *d.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedOrderResourcesReferencingTarget() (orders []Order, err error) {
	if d.RevIncludedOrderResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *d.RevIncludedOrderResourcesReferencingTarget
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if d.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *d.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingOrderer() (procedureRequests []ProcedureRequest, err error) {
	if d.RevIncludedProcedureRequestResourcesReferencingOrderer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *d.RevIncludedProcedureRequestResourcesReferencingOrderer
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDevice() (deviceUseRequests []DeviceUseRequest, err error) {
	if d.RevIncludedDeviceUseRequestResourcesReferencingDevice == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *d.RevIncludedDeviceUseRequestResourcesReferencingDevice
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDeviceMetricResourcesReferencingSource() (deviceMetrics []DeviceMetric, err error) {
	if d.RevIncludedDeviceMetricResourcesReferencingSource == nil {
		err = errors.New("RevIncluded deviceMetrics not requested")
	} else {
		deviceMetrics = *d.RevIncludedDeviceMetricResourcesReferencingSource
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedFlagResourcesReferencingAuthor() (flags []Flag, err error) {
	if d.RevIncludedFlagResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *d.RevIncludedFlagResourcesReferencingAuthor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if d.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *d.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedObservationResourcesReferencingSubject() (observations []Observation, err error) {
	if d.RevIncludedObservationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *d.RevIncludedObservationResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedObservationResourcesReferencingDevice() (observations []Observation, err error) {
	if d.RevIncludedObservationResourcesReferencingDevice == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *d.RevIncludedObservationResourcesReferencingDevice
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingDevice() (medicationAdministrations []MedicationAdministration, err error) {
	if d.RevIncludedMedicationAdministrationResourcesReferencingDevice == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *d.RevIncludedMedicationAdministrationResourcesReferencingDevice
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedContractResourcesReferencingActor() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingActor == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingActor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if d.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *d.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if d.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *d.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingPerformer() (riskAssessments []RiskAssessment, err error) {
	if d.RevIncludedRiskAssessmentResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *d.RevIncludedRiskAssessmentResourcesReferencingPerformer
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if d.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *d.RevIncludedGroupResourcesReferencingMember
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingSubject() (diagnosticReports []DiagnosticReport, err error) {
	if d.RevIncludedDiagnosticReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *d.RevIncludedDiagnosticReportResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedImagingObjectSelectionResourcesReferencingAuthor() (imagingObjectSelections []ImagingObjectSelection, err error) {
	if d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded imagingObjectSelections not requested")
	} else {
		imagingObjectSelections = *d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDeviceComponentResourcesReferencingSource() (deviceComponents []DeviceComponent, err error) {
	if d.RevIncludedDeviceComponentResourcesReferencingSource == nil {
		err = errors.New("RevIncluded deviceComponents not requested")
	} else {
		deviceComponents = *d.RevIncludedDeviceComponentResourcesReferencingSource
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingParticipant() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingParticipant
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAuthor() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingAuthor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingAuthor() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingAuthor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingSubject() (diagnosticOrders []DiagnosticOrder, err error) {
	if d.RevIncludedDiagnosticOrderResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *d.RevIncludedDiagnosticOrderResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingActorPath1() (diagnosticOrders []DiagnosticOrder, err error) {
	if d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingActorPath2() (diagnosticOrders []DiagnosticOrder, err error) {
	if d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *d.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingWho() (orderResponses []OrderResponse, err error) {
	if d.RevIncludedOrderResponseResourcesReferencingWho == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *d.RevIncludedOrderResponseResourcesReferencingWho
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDeviceUseStatementResourcesReferencingDevice() (deviceUseStatements []DeviceUseStatement, err error) {
	if d.RevIncludedDeviceUseStatementResourcesReferencingDevice == nil {
		err = errors.New("RevIncluded deviceUseStatements not requested")
	} else {
		deviceUseStatements = *d.RevIncludedDeviceUseStatementResourcesReferencingDevice
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if d.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *d.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if d.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *d.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingTarget() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingTarget
	}
	return
}

func (d *DevicePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *d.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedLocationResourcesReferencedByLocation != nil {
		for _, r := range *d.IncludedLocationResourcesReferencedByLocation {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (d *DevicePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.RevIncludedAppointmentResourcesReferencingActor != nil {
		for _, r := range *d.RevIncludedAppointmentResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedAccountResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedAccountResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for _, r := range *d.RevIncludedProvenanceResourcesReferencingAgent {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedSpecimenResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *d.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedListResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedListResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedListResourcesReferencingSource != nil {
		for _, r := range *d.RevIncludedListResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedDocumentReferenceResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *d.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResourcesReferencingTarget != nil {
		for _, r := range *d.RevIncludedOrderResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedMediaResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedMediaResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for _, r := range *d.RevIncludedProcedureRequestResourcesReferencingOrderer {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingDevice != nil {
		for _, r := range *d.RevIncludedDeviceUseRequestResourcesReferencingDevice {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDeviceMetricResourcesReferencingSource != nil {
		for _, r := range *d.RevIncludedDeviceMetricResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedFlagResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedFlagResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for _, r := range *d.RevIncludedAppointmentResponseResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedObservationResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedObservationResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedObservationResourcesReferencingDevice != nil {
		for _, r := range *d.RevIncludedObservationResourcesReferencingDevice {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedMedicationAdministrationResourcesReferencingDevice != nil {
		for _, r := range *d.RevIncludedMedicationAdministrationResourcesReferencingDevice {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedContractResourcesReferencingActor != nil {
		for _, r := range *d.RevIncludedContractResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for _, r := range *d.RevIncludedCommunicationRequestResourcesReferencingSender {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for _, r := range *d.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for _, r := range *d.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedGroupResourcesReferencingMember != nil {
		for _, r := range *d.RevIncludedGroupResourcesReferencingMember {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedDiagnosticReportResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDeviceComponentResourcesReferencingSource != nil {
		for _, r := range *d.RevIncludedDeviceComponentResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for _, r := range *d.RevIncludedAuditEventResourcesReferencingParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *d.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingSender != nil {
		for _, r := range *d.RevIncludedCommunicationResourcesReferencingSender {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for _, r := range *d.RevIncludedCommunicationResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedCompositionResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *d.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedDetectedIssueResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 != nil {
		for _, r := range *d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 != nil {
		for _, r := range *d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *d.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResponseResourcesReferencingWho != nil {
		for _, r := range *d.RevIncludedOrderResponseResourcesReferencingWho {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDeviceUseStatementResourcesReferencingDevice != nil {
		for _, r := range *d.RevIncludedDeviceUseStatementResourcesReferencingDevice {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *d.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedScheduleResourcesReferencingActor != nil {
		for _, r := range *d.RevIncludedScheduleResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *d.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingTarget != nil {
		for _, r := range *d.RevIncludedMessageHeaderResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (d *DevicePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *d.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedLocationResourcesReferencedByLocation != nil {
		for _, r := range *d.IncludedLocationResourcesReferencedByLocation {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedAppointmentResourcesReferencingActor != nil {
		for _, r := range *d.RevIncludedAppointmentResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedAccountResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedAccountResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for _, r := range *d.RevIncludedProvenanceResourcesReferencingAgent {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedSpecimenResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *d.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedListResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedListResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedListResourcesReferencingSource != nil {
		for _, r := range *d.RevIncludedListResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedDocumentReferenceResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *d.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResourcesReferencingTarget != nil {
		for _, r := range *d.RevIncludedOrderResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedMediaResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedMediaResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for _, r := range *d.RevIncludedProcedureRequestResourcesReferencingOrderer {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingDevice != nil {
		for _, r := range *d.RevIncludedDeviceUseRequestResourcesReferencingDevice {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDeviceMetricResourcesReferencingSource != nil {
		for _, r := range *d.RevIncludedDeviceMetricResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedFlagResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedFlagResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for _, r := range *d.RevIncludedAppointmentResponseResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedObservationResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedObservationResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedObservationResourcesReferencingDevice != nil {
		for _, r := range *d.RevIncludedObservationResourcesReferencingDevice {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedMedicationAdministrationResourcesReferencingDevice != nil {
		for _, r := range *d.RevIncludedMedicationAdministrationResourcesReferencingDevice {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedContractResourcesReferencingActor != nil {
		for _, r := range *d.RevIncludedContractResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for _, r := range *d.RevIncludedCommunicationRequestResourcesReferencingSender {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for _, r := range *d.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for _, r := range *d.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedGroupResourcesReferencingMember != nil {
		for _, r := range *d.RevIncludedGroupResourcesReferencingMember {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedDiagnosticReportResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDeviceComponentResourcesReferencingSource != nil {
		for _, r := range *d.RevIncludedDeviceComponentResourcesReferencingSource {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for _, r := range *d.RevIncludedAuditEventResourcesReferencingParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *d.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingSender != nil {
		for _, r := range *d.RevIncludedCommunicationResourcesReferencingSender {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for _, r := range *d.RevIncludedCommunicationResourcesReferencingRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedCompositionResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *d.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedDetectedIssueResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 != nil {
		for _, r := range *d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 != nil {
		for _, r := range *d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *d.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResponseResourcesReferencingWho != nil {
		for _, r := range *d.RevIncludedOrderResponseResourcesReferencingWho {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for _, r := range *d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDeviceUseStatementResourcesReferencingDevice != nil {
		for _, r := range *d.RevIncludedDeviceUseStatementResourcesReferencingDevice {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *d.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedScheduleResourcesReferencingActor != nil {
		for _, r := range *d.RevIncludedScheduleResourcesReferencingActor {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *d.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingTarget != nil {
		for _, r := range *d.RevIncludedMessageHeaderResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
