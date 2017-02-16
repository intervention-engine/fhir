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

type Appointment struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status          string                            `bson:"status,omitempty" json:"status,omitempty"`
	ServiceCategory *CodeableConcept                  `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType     []CodeableConcept                 `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty       []CodeableConcept                 `bson:"specialty,omitempty" json:"specialty,omitempty"`
	AppointmentType *CodeableConcept                  `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	Reason          *CodeableConcept                  `bson:"reason,omitempty" json:"reason,omitempty"`
	Priority        *uint32                           `bson:"priority,omitempty" json:"priority,omitempty"`
	Description     string                            `bson:"description,omitempty" json:"description,omitempty"`
	Start           *FHIRDateTime                     `bson:"start,omitempty" json:"start,omitempty"`
	End             *FHIRDateTime                     `bson:"end,omitempty" json:"end,omitempty"`
	MinutesDuration *uint32                           `bson:"minutesDuration,omitempty" json:"minutesDuration,omitempty"`
	Slot            []Reference                       `bson:"slot,omitempty" json:"slot,omitempty"`
	Created         *FHIRDateTime                     `bson:"created,omitempty" json:"created,omitempty"`
	Comment         string                            `bson:"comment,omitempty" json:"comment,omitempty"`
	Participant     []AppointmentParticipantComponent `bson:"participant,omitempty" json:"participant,omitempty"`
	RequestedPeriod []Period                          `bson:"requestedPeriod,omitempty" json:"requestedPeriod,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Appointment) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Appointment"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Appointment), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Appointment) GetBSON() (interface{}, error) {
	x.ResourceType = "Appointment"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "appointment" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type appointment Appointment

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Appointment) UnmarshalJSON(data []byte) (err error) {
	x2 := appointment{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Appointment(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Appointment) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Appointment"
	} else if x.ResourceType != "Appointment" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Appointment, instead received %s", x.ResourceType))
	}
	return nil
}

type AppointmentParticipantComponent struct {
	BackboneElement `bson:",inline"`
	Type            []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Actor           *Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
	Required        string            `bson:"required,omitempty" json:"required,omitempty"`
	Status          string            `bson:"status,omitempty" json:"status,omitempty"`
}

type AppointmentPlus struct {
	Appointment                     `bson:",inline"`
	AppointmentPlusRelatedResources `bson:",inline"`
}

type AppointmentPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByActor                *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByActor,omitempty"`
	IncludedDeviceResourcesReferencedByActor                      *[]Device                `bson:"_includedDeviceResourcesReferencedByActor,omitempty"`
	IncludedPatientResourcesReferencedByActor                     *[]Patient               `bson:"_includedPatientResourcesReferencedByActor,omitempty"`
	IncludedHealthcareServiceResourcesReferencedByActor           *[]HealthcareService     `bson:"_includedHealthcareServiceResourcesReferencedByActor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByActor               *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByActor,omitempty"`
	IncludedLocationResourcesReferencedByActor                    *[]Location              `bson:"_includedLocationResourcesReferencedByActor,omitempty"`
	IncludedPractitionerResourcesReferencedByPractitioner         *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByPractitioner,omitempty"`
	IncludedPatientResourcesReferencedByPatient                   *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedLocationResourcesReferencedByLocation                 *[]Location              `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref     *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref     *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                    *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref    *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest           *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse          *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource    *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedEncounterResourcesReferencingAppointment           *[]Encounter             `bson:"_revIncludedEncounterResourcesReferencingAppointment,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon           *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData              *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity               *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget               *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                    *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                    *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference      *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                       *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces      *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon       *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition    *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces       *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon        *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition     *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingAppointment *[]AppointmentResponse   `bson:"_revIncludedAppointmentResponseResourcesReferencingAppointment,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                   *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity               *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry               *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated        *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject   *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest         *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAction       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
}

func (a *AppointmentPlusRelatedResources) GetIncludedPractitionerResourceReferencedByActor() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByActor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByActor))
	} else if len(*a.IncludedPractitionerResourcesReferencedByActor) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByActor)[0]
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetIncludedDeviceResourceReferencedByActor() (device *Device, err error) {
	if a.IncludedDeviceResourcesReferencedByActor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedDeviceResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedDeviceResourcesReferencedByActor))
	} else if len(*a.IncludedDeviceResourcesReferencedByActor) == 1 {
		device = &(*a.IncludedDeviceResourcesReferencedByActor)[0]
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetIncludedPatientResourceReferencedByActor() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByActor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByActor))
	} else if len(*a.IncludedPatientResourcesReferencedByActor) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByActor)[0]
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetIncludedHealthcareServiceResourceReferencedByActor() (healthcareService *HealthcareService, err error) {
	if a.IncludedHealthcareServiceResourcesReferencedByActor == nil {
		err = errors.New("Included healthcareservices not requested")
	} else if len(*a.IncludedHealthcareServiceResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 healthcareService, but found %d", len(*a.IncludedHealthcareServiceResourcesReferencedByActor))
	} else if len(*a.IncludedHealthcareServiceResourcesReferencedByActor) == 1 {
		healthcareService = &(*a.IncludedHealthcareServiceResourcesReferencedByActor)[0]
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByActor() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedRelatedPersonResourcesReferencedByActor == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedRelatedPersonResourcesReferencedByActor))
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByActor) == 1 {
		relatedPerson = &(*a.IncludedRelatedPersonResourcesReferencedByActor)[0]
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetIncludedLocationResourceReferencedByActor() (location *Location, err error) {
	if a.IncludedLocationResourcesReferencedByActor == nil {
		err = errors.New("Included locations not requested")
	} else if len(*a.IncludedLocationResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*a.IncludedLocationResourcesReferencedByActor))
	} else if len(*a.IncludedLocationResourcesReferencedByActor) == 1 {
		location = &(*a.IncludedLocationResourcesReferencedByActor)[0]
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPractitioner() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByPractitioner == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByPractitioner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByPractitioner))
	} else if len(*a.IncludedPractitionerResourcesReferencedByPractitioner) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByPractitioner)[0]
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatient))
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if a.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*a.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*a.IncludedLocationResourcesReferencedByLocation))
	} else if len(*a.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*a.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if a.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *a.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if a.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *a.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingAppointment() (encounters []Encounter, err error) {
	if a.RevIncludedEncounterResourcesReferencingAppointment == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *a.RevIncludedEncounterResourcesReferencingAppointment
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if a.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *a.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if a.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *a.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if a.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *a.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if a.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *a.RevIncludedListResourcesReferencingItem
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if a.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *a.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if a.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *a.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if a.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *a.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if a.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *a.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if a.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *a.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if a.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *a.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingAppointment() (appointmentResponses []AppointmentResponse, err error) {
	if a.RevIncludedAppointmentResponseResourcesReferencingAppointment == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *a.RevIncludedAppointmentResponseResourcesReferencingAppointment
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if a.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *a.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if a.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *a.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *a.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if a.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *a.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAction() (clinicalImpressions []ClinicalImpression, err error) {
	if a.RevIncludedClinicalImpressionResourcesReferencingAction == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *a.RevIncludedClinicalImpressionResourcesReferencingAction
	}
	return
}

func (a *AppointmentPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByActor != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByActor {
			rsc := (*a.IncludedPractitionerResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedByActor != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedByActor {
			rsc := (*a.IncludedDeviceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByActor != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByActor {
			rsc := (*a.IncludedPatientResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedHealthcareServiceResourcesReferencedByActor != nil {
		for idx := range *a.IncludedHealthcareServiceResourcesReferencedByActor {
			rsc := (*a.IncludedHealthcareServiceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByActor != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByActor {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedByActor != nil {
		for idx := range *a.IncludedLocationResourcesReferencedByActor {
			rsc := (*a.IncludedLocationResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByPractitioner {
			rsc := (*a.IncludedPractitionerResourcesReferencedByPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *a.IncludedLocationResourcesReferencedByLocation {
			rsc := (*a.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AppointmentPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingData {
			rsc := (*a.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEncounterResourcesReferencingAppointment != nil {
		for idx := range *a.RevIncludedEncounterResourcesReferencingAppointment {
			rsc := (*a.RevIncludedEncounterResourcesReferencingAppointment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *a.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*a.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAppointmentResponseResourcesReferencingAppointment != nil {
		for idx := range *a.RevIncludedAppointmentResponseResourcesReferencingAppointment {
			rsc := (*a.RevIncludedAppointmentResponseResourcesReferencingAppointment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*a.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AppointmentPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByActor != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByActor {
			rsc := (*a.IncludedPractitionerResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedByActor != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedByActor {
			rsc := (*a.IncludedDeviceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByActor != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByActor {
			rsc := (*a.IncludedPatientResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedHealthcareServiceResourcesReferencedByActor != nil {
		for idx := range *a.IncludedHealthcareServiceResourcesReferencedByActor {
			rsc := (*a.IncludedHealthcareServiceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByActor != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByActor {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedByActor != nil {
		for idx := range *a.IncludedLocationResourcesReferencedByActor {
			rsc := (*a.IncludedLocationResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByPractitioner {
			rsc := (*a.IncludedPractitionerResourcesReferencedByPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *a.IncludedLocationResourcesReferencedByLocation {
			rsc := (*a.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingData {
			rsc := (*a.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEncounterResourcesReferencingAppointment != nil {
		for idx := range *a.RevIncludedEncounterResourcesReferencingAppointment {
			rsc := (*a.RevIncludedEncounterResourcesReferencingAppointment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *a.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*a.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAppointmentResponseResourcesReferencingAppointment != nil {
		for idx := range *a.RevIncludedAppointmentResponseResourcesReferencingAppointment {
			rsc := (*a.RevIncludedAppointmentResponseResourcesReferencingAppointment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*a.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
