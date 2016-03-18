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
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*d.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *d.IncludedLocationResourcesReferencedByLocation {
			rsc := (*d.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DevicePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*d.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*d.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*d.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedListResourcesReferencingSubject {
			rsc := (*d.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedListResourcesReferencingSource {
			rsc := (*d.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedOrderResourcesReferencingSubject {
			rsc := (*d.RevIncludedOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *d.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*d.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedOrderResourcesReferencingTarget {
			rsc := (*d.RevIncludedOrderResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*d.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingOrderer {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingOrderer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceMetricResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedDeviceMetricResourcesReferencingSource {
			rsc := (*d.RevIncludedDeviceMetricResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*d.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*d.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*d.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingDevice {
			rsc := (*d.RevIncludedObservationResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMedicationAdministrationResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedMedicationAdministrationResourcesReferencingDevice {
			rsc := (*d.RevIncludedMedicationAdministrationResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingActor {
			rsc := (*d.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*d.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *d.RevIncludedGroupResourcesReferencingMember {
			rsc := (*d.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*d.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			rsc := (*d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceComponentResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedDeviceComponentResourcesReferencingSource {
			rsc := (*d.RevIncludedDeviceComponentResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingParticipant {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*d.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			rsc := (*d.RevIncludedDiagnosticOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 != nil {
		for idx := range *d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 {
			rsc := (*d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 != nil {
		for idx := range *d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 {
			rsc := (*d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *d.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*d.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResponseResourcesReferencingWho != nil {
		for idx := range *d.RevIncludedOrderResponseResourcesReferencingWho {
			rsc := (*d.RevIncludedOrderResponseResourcesReferencingWho)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseStatementResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceUseStatementResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceUseStatementResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*d.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*d.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingTarget {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DevicePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*d.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *d.IncludedLocationResourcesReferencedByLocation {
			rsc := (*d.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*d.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*d.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*d.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedListResourcesReferencingSubject {
			rsc := (*d.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedListResourcesReferencingSource {
			rsc := (*d.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedOrderResourcesReferencingSubject {
			rsc := (*d.RevIncludedOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *d.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*d.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedOrderResourcesReferencingTarget {
			rsc := (*d.RevIncludedOrderResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*d.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingOrderer != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingOrderer {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingOrderer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseRequestResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceUseRequestResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceUseRequestResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceMetricResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedDeviceMetricResourcesReferencingSource {
			rsc := (*d.RevIncludedDeviceMetricResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*d.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*d.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*d.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingDevice {
			rsc := (*d.RevIncludedObservationResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMedicationAdministrationResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedMedicationAdministrationResourcesReferencingDevice {
			rsc := (*d.RevIncludedMedicationAdministrationResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingActor {
			rsc := (*d.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*d.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *d.RevIncludedGroupResourcesReferencingMember {
			rsc := (*d.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*d.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			rsc := (*d.RevIncludedImagingObjectSelectionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceComponentResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedDeviceComponentResourcesReferencingSource {
			rsc := (*d.RevIncludedDeviceComponentResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingParticipant {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*d.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDiagnosticOrderResourcesReferencingSubject {
			rsc := (*d.RevIncludedDiagnosticOrderResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 != nil {
		for idx := range *d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1 {
			rsc := (*d.RevIncludedDiagnosticOrderResourcesReferencingActorPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 != nil {
		for idx := range *d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2 {
			rsc := (*d.RevIncludedDiagnosticOrderResourcesReferencingActorPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *d.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*d.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResponseResourcesReferencingWho != nil {
		for idx := range *d.RevIncludedOrderResponseResourcesReferencingWho {
			rsc := (*d.RevIncludedOrderResponseResourcesReferencingWho)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseStatementResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceUseStatementResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceUseStatementResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*d.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*d.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingTarget {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
