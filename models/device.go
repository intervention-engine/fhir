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

type Device struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Udi             *DeviceUdiComponent `bson:"udi,omitempty" json:"udi,omitempty"`
	Status          string              `bson:"status,omitempty" json:"status,omitempty"`
	Type            *CodeableConcept    `bson:"type,omitempty" json:"type,omitempty"`
	LotNumber       string              `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	Manufacturer    string              `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	ManufactureDate *FHIRDateTime       `bson:"manufactureDate,omitempty" json:"manufactureDate,omitempty"`
	ExpirationDate  *FHIRDateTime       `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	Model           string              `bson:"model,omitempty" json:"model,omitempty"`
	Version         string              `bson:"version,omitempty" json:"version,omitempty"`
	Patient         *Reference          `bson:"patient,omitempty" json:"patient,omitempty"`
	Owner           *Reference          `bson:"owner,omitempty" json:"owner,omitempty"`
	Contact         []ContactPoint      `bson:"contact,omitempty" json:"contact,omitempty"`
	Location        *Reference          `bson:"location,omitempty" json:"location,omitempty"`
	Url             string              `bson:"url,omitempty" json:"url,omitempty"`
	Note            []Annotation        `bson:"note,omitempty" json:"note,omitempty"`
	Safety          []CodeableConcept   `bson:"safety,omitempty" json:"safety,omitempty"`
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
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
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

type DeviceUdiComponent struct {
	BackboneElement  `bson:",inline"`
	DeviceIdentifier string `bson:"deviceIdentifier,omitempty" json:"deviceIdentifier,omitempty"`
	Name             string `bson:"name,omitempty" json:"name,omitempty"`
	Jurisdiction     string `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	CarrierHRF       string `bson:"carrierHRF,omitempty" json:"carrierHRF,omitempty"`
	CarrierAIDC      string `bson:"carrierAIDC,omitempty" json:"carrierAIDC,omitempty"`
	Issuer           string `bson:"issuer,omitempty" json:"issuer,omitempty"`
	EntryType        string `bson:"entryType,omitempty" json:"entryType,omitempty"`
}

type DevicePlus struct {
	Device                     `bson:",inline"`
	DevicePlusRelatedResources `bson:",inline"`
}

type DevicePlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                      *[]Patient                  `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization            *[]Organization             `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedLocationResourcesReferencedByLocation                    *[]Location                 `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                  *[]Appointment              `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRequester          *[]ReferralRequest          `bson:"_revIncludedReferralRequestResourcesReferencingRequester,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                    *[]Account                  `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref        *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingSubject           *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor            *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref        *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                  *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                  *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedConsentResourcesReferencingActorPath1                 *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingActorPath1,omitempty"`
	RevIncludedConsentResourcesReferencingActorPath2                 *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingActorPath2,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                  *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                 *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1             *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2             *[]Measure                  `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingSubject          *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor           *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref       *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedImagingManifestResourcesReferencingAuthor             *[]ImagingManifest          `bson:"_revIncludedImagingManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingRequester            *[]SupplyRequest            `bson:"_revIncludedSupplyRequestResourcesReferencingRequester,omitempty"`
	RevIncludedContractResourcesReferencingAgent                     *[]Contract                 `bson:"_revIncludedContractResourcesReferencingAgent,omitempty"`
	RevIncludedContractResourcesReferencingSubject                   *[]Contract                 `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                 *[]Contract                 `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingPerformer           *[]RiskAssessment           `bson:"_revIncludedRiskAssessmentResourcesReferencingPerformer,omitempty"`
	RevIncludedGroupResourcesReferencingMember                       *[]Group                    `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest              *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse             *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource       *[]ImplementationGuide      `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedDeviceComponentResourcesReferencingSource             *[]DeviceComponent          `bson:"_revIncludedDeviceComponentResourcesReferencingSource,omitempty"`
	RevIncludedChargeItemResourcesReferencingParticipantactor        *[]ChargeItem               `bson:"_revIncludedChargeItemResourcesReferencingParticipantactor,omitempty"`
	RevIncludedChargeItemResourcesReferencingEnterer                 *[]ChargeItem               `bson:"_revIncludedChargeItemResourcesReferencingEnterer,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor        *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom      *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor      *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof       *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson        *[]ServiceDefinition        `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof               *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon              *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender               *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient            *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor       *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom     *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor     *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof      *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1  *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2  *[]ActivityDefinition       `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDeviceUseStatementResourcesReferencingDevice          *[]DeviceUseStatement       `bson:"_revIncludedDeviceUseStatementResourcesReferencingDevice,omitempty"`
	RevIncludedRequestGroupResourcesReferencingAuthor                *[]RequestGroup             `bson:"_revIncludedRequestGroupResourcesReferencingAuthor,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition            *[]RequestGroup             `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingRequester            *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingRequester,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPerformer            *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingSubject              *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon              *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest         *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingDevice               *[]DeviceRequest            `bson:"_revIncludedDeviceRequestResourcesReferencingDevice,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingTarget               *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingTarget,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref               *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                   *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                  *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingOwner                         *[]Task                     `bson:"_revIncludedTaskResourcesReferencingOwner,omitempty"`
	RevIncludedTaskResourcesReferencingRequester                     *[]Task                     `bson:"_revIncludedTaskResourcesReferencingRequester,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                       *[]Task                     `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                         *[]Task                     `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                       *[]Task                     `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedSpecimenResourcesReferencingSubject                   *[]Specimen                 `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                *[]Procedure                `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedListResourcesReferencingItem                          *[]List                     `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                       *[]List                     `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingSource                        *[]List                     `bson:"_revIncludedListResourcesReferencingSource,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingRequester        *[]MedicationRequest        `bson:"_revIncludedMedicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                      *[]Media                    `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedMediaResourcesReferencingDevice                       *[]Media                    `bson:"_revIncludedMediaResourcesReferencingDevice,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingRequester         *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingRequester,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingPerformer         *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces          *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingSubject           *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon           *[]ProcedureRequest         `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceMetricResourcesReferencingSource                *[]DeviceMetric             `bson:"_revIncludedDeviceMetricResourcesReferencingSource,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                        *[]Flag                     `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor          *[]AppointmentResponse      `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAdverseEventResourcesReferencingSubject               *[]AdverseEvent             `bson:"_revIncludedAdverseEventResourcesReferencingSubject,omitempty"`
	RevIncludedAdverseEventResourcesReferencingSubstance             *[]AdverseEvent             `bson:"_revIncludedAdverseEventResourcesReferencingSubstance,omitempty"`
	RevIncludedObservationResourcesReferencingSubject                *[]Observation              `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingDevice                 *[]Observation              `bson:"_revIncludedObservationResourcesReferencingDevice,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPerformer *[]MedicationAdministration `bson:"_revIncludedMedicationAdministrationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingDevice    *[]MedicationAdministration `bson:"_revIncludedMedicationAdministrationResourcesReferencingDevice,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                  *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                 *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                  *[]Library                  `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRequester     *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon       *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender        *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient     *[]CommunicationRequest     `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                      *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPerformer       *[]MedicationDispense       `bson:"_revIncludedMedicationDispenseResourcesReferencingPerformer,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject           *[]DiagnosticReport         `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingAgent                   *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingAgent,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                  *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail           *[]Condition                `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor                 *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                  *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingAuthor               *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingAuthor,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated           *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject      *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor       *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest            *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                     *[]Schedule                 `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor           *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom         *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor         *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof          *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1      *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2      *[]PlanDefinition           `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
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

func (d *DevicePlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingRequester() (referralRequests []ReferralRequest, err error) {
	if d.RevIncludedReferralRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *d.RevIncludedReferralRequestResourcesReferencingRequester
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

func (d *DevicePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedConsentResourcesReferencingActorPath1() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingActorPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingActorPath1
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedConsentResourcesReferencingActorPath2() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingActorPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingActorPath2
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDependsonPath2
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

func (d *DevicePlusRelatedResources) GetRevIncludedImagingManifestResourcesReferencingAuthor() (imagingManifests []ImagingManifest, err error) {
	if d.RevIncludedImagingManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded imagingManifests not requested")
	} else {
		imagingManifests = *d.RevIncludedImagingManifestResourcesReferencingAuthor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingRequester() (supplyRequests []SupplyRequest, err error) {
	if d.RevIncludedSupplyRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *d.RevIncludedSupplyRequestResourcesReferencingRequester
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedContractResourcesReferencingAgent() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingAgent
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingTermtopic
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

func (d *DevicePlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if d.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *d.RevIncludedGroupResourcesReferencingMember
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if d.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *d.RevIncludedImplementationGuideResourcesReferencingResource
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

func (d *DevicePlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingParticipantactor() (chargeItems []ChargeItem, err error) {
	if d.RevIncludedChargeItemResourcesReferencingParticipantactor == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *d.RevIncludedChargeItemResourcesReferencingParticipantactor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingEnterer() (chargeItems []ChargeItem, err error) {
	if d.RevIncludedChargeItemResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *d.RevIncludedChargeItemResourcesReferencingEnterer
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if d.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *d.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if d.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *d.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if d.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *d.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if d.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *d.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if d.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *d.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingBasedon
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

func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
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

func (d *DevicePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingAuthor() (requestGroups []RequestGroup, err error) {
	if d.RevIncludedRequestGroupResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *d.RevIncludedRequestGroupResourcesReferencingAuthor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if d.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *d.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingRequester() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingRequester
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPerformer() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingPerformer
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingSubject() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingDevice() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingDevice == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingDevice
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingFocus
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

func (d *DevicePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingEntityref
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

func (d *DevicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingOwner() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingOwner
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingRequester() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingRequester
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingBasedon
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

func (d *DevicePlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if d.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *d.RevIncludedProcedureResourcesReferencingPerformer
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

func (d *DevicePlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingRequester() (medicationRequests []MedicationRequest, err error) {
	if d.RevIncludedMedicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *d.RevIncludedMedicationRequestResourcesReferencingRequester
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

func (d *DevicePlusRelatedResources) GetRevIncludedMediaResourcesReferencingDevice() (media []Media, err error) {
	if d.RevIncludedMediaResourcesReferencingDevice == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *d.RevIncludedMediaResourcesReferencingDevice
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingRequester() (procedureRequests []ProcedureRequest, err error) {
	if d.RevIncludedProcedureRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *d.RevIncludedProcedureRequestResourcesReferencingRequester
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingPerformer() (procedureRequests []ProcedureRequest, err error) {
	if d.RevIncludedProcedureRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *d.RevIncludedProcedureRequestResourcesReferencingPerformer
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if d.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *d.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingSubject() (procedureRequests []ProcedureRequest, err error) {
	if d.RevIncludedProcedureRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *d.RevIncludedProcedureRequestResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if d.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *d.RevIncludedProcedureRequestResourcesReferencingBasedon
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

func (d *DevicePlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubject() (adverseEvents []AdverseEvent, err error) {
	if d.RevIncludedAdverseEventResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *d.RevIncludedAdverseEventResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubstance() (adverseEvents []AdverseEvent, err error) {
	if d.RevIncludedAdverseEventResourcesReferencingSubstance == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *d.RevIncludedAdverseEventResourcesReferencingSubstance
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

func (d *DevicePlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPerformer() (medicationAdministrations []MedicationAdministration, err error) {
	if d.RevIncludedMedicationAdministrationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *d.RevIncludedMedicationAdministrationResourcesReferencingPerformer
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

func (d *DevicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRequester() (communicationRequests []CommunicationRequest, err error) {
	if d.RevIncludedCommunicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *d.RevIncludedCommunicationRequestResourcesReferencingRequester
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *d.RevIncludedCommunicationRequestResourcesReferencingBasedon
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

func (d *DevicePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPerformer() (medicationDispenses []MedicationDispense, err error) {
	if d.RevIncludedMedicationDispenseResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *d.RevIncludedMedicationDispenseResourcesReferencingPerformer
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

func (d *DevicePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingAgent() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingAgent
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if d.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *d.RevIncludedConditionResourcesReferencingEvidencedetail
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

func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
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
	if d.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedReferralRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedReferralRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*d.RevIncludedAccountResourcesReferencingSubject)[idx]
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
	if d.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*d.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*d.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingActorPath1 != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingActorPath1 {
			rsc := (*d.RevIncludedConsentResourcesReferencingActorPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingActorPath2 != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingActorPath2 {
			rsc := (*d.RevIncludedConsentResourcesReferencingActorPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*d.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
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
	if d.RevIncludedImagingManifestResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedImagingManifestResourcesReferencingAuthor {
			rsc := (*d.RevIncludedImagingManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingAgent {
			rsc := (*d.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*d.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*d.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *d.RevIncludedGroupResourcesReferencingMember {
			rsc := (*d.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceComponentResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedDeviceComponentResourcesReferencingSource {
			rsc := (*d.RevIncludedDeviceComponentResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedChargeItemResourcesReferencingParticipantactor != nil {
		for idx := range *d.RevIncludedChargeItemResourcesReferencingParticipantactor {
			rsc := (*d.RevIncludedChargeItemResourcesReferencingParticipantactor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *d.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*d.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*d.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
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
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseStatementResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceUseStatementResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceUseStatementResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedRequestGroupResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedRequestGroupResourcesReferencingAuthor {
			rsc := (*d.RevIncludedRequestGroupResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*d.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingTarget {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if d.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*d.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*d.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*d.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*d.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*d.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*d.RevIncludedProcedureResourcesReferencingPerformer)[idx]
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
	if d.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*d.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingDevice {
			rsc := (*d.RevIncludedMediaResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingSubject {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
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
	if d.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*d.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *d.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*d.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
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
	if d.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*d.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMedicationAdministrationResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedMedicationAdministrationResourcesReferencingDevice {
			rsc := (*d.RevIncludedMedicationAdministrationResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*d.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*d.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*d.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *d.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*d.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
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
	if d.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedReferralRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedReferralRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*d.RevIncludedAccountResourcesReferencingSubject)[idx]
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
	if d.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*d.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*d.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingActorPath1 != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingActorPath1 {
			rsc := (*d.RevIncludedConsentResourcesReferencingActorPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingActorPath2 != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingActorPath2 {
			rsc := (*d.RevIncludedConsentResourcesReferencingActorPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*d.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
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
	if d.RevIncludedImagingManifestResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedImagingManifestResourcesReferencingAuthor {
			rsc := (*d.RevIncludedImagingManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingAgent {
			rsc := (*d.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*d.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*d.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *d.RevIncludedGroupResourcesReferencingMember {
			rsc := (*d.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceComponentResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedDeviceComponentResourcesReferencingSource {
			rsc := (*d.RevIncludedDeviceComponentResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedChargeItemResourcesReferencingParticipantactor != nil {
		for idx := range *d.RevIncludedChargeItemResourcesReferencingParticipantactor {
			rsc := (*d.RevIncludedChargeItemResourcesReferencingParticipantactor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *d.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*d.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*d.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
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
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceUseStatementResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceUseStatementResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceUseStatementResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedRequestGroupResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedRequestGroupResourcesReferencingAuthor {
			rsc := (*d.RevIncludedRequestGroupResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*d.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingTarget {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if d.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*d.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*d.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*d.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*d.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*d.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*d.RevIncludedProcedureResourcesReferencingPerformer)[idx]
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
	if d.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*d.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingDevice {
			rsc := (*d.RevIncludedMediaResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingSubject {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
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
	if d.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*d.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *d.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*d.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
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
	if d.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*d.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMedicationAdministrationResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedMedicationAdministrationResourcesReferencingDevice {
			rsc := (*d.RevIncludedMedicationAdministrationResourcesReferencingDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*d.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*d.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*d.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *d.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*d.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
