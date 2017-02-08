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

type MedicationRequest struct {
	DomainResource            `bson:",inline"`
	Identifier                []Identifier                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition                []Reference                                `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn                   []Reference                                `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Requisition               *Identifier                                `bson:"requisition,omitempty" json:"requisition,omitempty"`
	Status                    string                                     `bson:"status,omitempty" json:"status,omitempty"`
	Stage                     *CodeableConcept                           `bson:"stage,omitempty" json:"stage,omitempty"`
	MedicationCodeableConcept *CodeableConcept                           `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                                 `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Patient                   *Reference                                 `bson:"patient,omitempty" json:"patient,omitempty"`
	Context                   *Reference                                 `bson:"context,omitempty" json:"context,omitempty"`
	SupportingInformation     []Reference                                `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	DateWritten               *FHIRDateTime                              `bson:"dateWritten,omitempty" json:"dateWritten,omitempty"`
	Requester                 *Reference                                 `bson:"requester,omitempty" json:"requester,omitempty"`
	ReasonCode                []CodeableConcept                          `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference           []Reference                                `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                      []Annotation                               `bson:"note,omitempty" json:"note,omitempty"`
	Category                  *CodeableConcept                           `bson:"category,omitempty" json:"category,omitempty"`
	DosageInstruction         []DosageInstruction                        `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	DispenseRequest           *MedicationRequestDispenseRequestComponent `bson:"dispenseRequest,omitempty" json:"dispenseRequest,omitempty"`
	Substitution              *MedicationRequestSubstitutionComponent    `bson:"substitution,omitempty" json:"substitution,omitempty"`
	PriorPrescription         *Reference                                 `bson:"priorPrescription,omitempty" json:"priorPrescription,omitempty"`
	EventHistory              []Reference                                `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *MedicationRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "MedicationRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to MedicationRequest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *MedicationRequest) GetBSON() (interface{}, error) {
	x.ResourceType = "MedicationRequest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "medicationRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medicationRequest MedicationRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *MedicationRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := medicationRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MedicationRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *MedicationRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "MedicationRequest"
	} else if x.ResourceType != "MedicationRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be MedicationRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type MedicationRequestDispenseRequestComponent struct {
	BackboneElement        `bson:",inline"`
	ValidityPeriod         *Period    `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	NumberOfRepeatsAllowed *uint32    `bson:"numberOfRepeatsAllowed,omitempty" json:"numberOfRepeatsAllowed,omitempty"`
	Quantity               *Quantity  `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ExpectedSupplyDuration *Quantity  `bson:"expectedSupplyDuration,omitempty" json:"expectedSupplyDuration,omitempty"`
	Performer              *Reference `bson:"performer,omitempty" json:"performer,omitempty"`
}

type MedicationRequestSubstitutionComponent struct {
	BackboneElement `bson:",inline"`
	Allowed         *bool            `bson:"allowed,omitempty" json:"allowed,omitempty"`
	Reason          *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}

type MedicationRequestPlus struct {
	MedicationRequest                     `bson:",inline"`
	MedicationRequestPlusRelatedResources `bson:",inline"`
}

type MedicationRequestPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByRequester                  *[]Practitioner             `bson:"_includedPractitionerResourcesReferencedByRequester,omitempty"`
	IncludedOrganizationResourcesReferencedByRequester                  *[]Organization             `bson:"_includedOrganizationResourcesReferencedByRequester,omitempty"`
	IncludedDeviceResourcesReferencedByRequester                        *[]Device                   `bson:"_includedDeviceResourcesReferencedByRequester,omitempty"`
	IncludedPatientResourcesReferencedByRequester                       *[]Patient                  `bson:"_includedPatientResourcesReferencedByRequester,omitempty"`
	IncludedRelatedPersonResourcesReferencedByRequester                 *[]RelatedPerson            `bson:"_includedRelatedPersonResourcesReferencedByRequester,omitempty"`
	IncludedOrganizationResourcesReferencedByIntendeddispenser          *[]Organization             `bson:"_includedOrganizationResourcesReferencedByIntendeddispenser,omitempty"`
	IncludedPatientResourcesReferencedByPatient                         *[]Patient                  `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext                   *[]EpisodeOfCare            `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext                       *[]Encounter                `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
	IncludedMedicationResourcesReferencedByMedication                   *[]Medication               `bson:"_includedMedicationResourcesReferencedByMedication,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref           *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref           *[]DocumentManifest         `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                          *[]Consent                  `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref          *[]DocumentReference        `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                       *[]Contract                 `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                      *[]Contract                 `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                        *[]Contract                 `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                 *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                *[]PaymentNotice            `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource          *[]ImplementationGuide      `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                 *[]Communication            `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                    *[]MessageHeader            `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                     *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                     *[]Provenance               `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                          *[]Task                     `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                            *[]Task                     `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                          *[]Task                     `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference            *[]CarePlan                 `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                             *[]List                     `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces            *[]DiagnosticRequest        `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon             *[]DiagnosticRequest        `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition          *[]DiagnosticRequest        `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces             *[]DeviceUseRequest         `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon              *[]DeviceUseRequest         `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition           *[]DeviceUseRequest         `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPrescription *[]MedicationAdministration `bson:"_revIncludedMedicationAdministrationResourcesReferencingPrescription,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                         *[]Basic                    `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPrescription       *[]MedicationDispense       `bson:"_revIncludedMedicationDispenseResourcesReferencingPrescription,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                     *[]AuditEvent               `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                   *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                     *[]Composition              `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated              *[]DetectedIssue            `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject         *[]QuestionnaireResponse    `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest               *[]ProcessResponse          `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAction             *[]ClinicalImpression       `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedPractitionerResourceReferencedByRequester() (practitioner *Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByRequester == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedPractitionerResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedPractitionerResourcesReferencedByRequester))
	} else if len(*m.IncludedPractitionerResourcesReferencedByRequester) == 1 {
		practitioner = &(*m.IncludedPractitionerResourcesReferencedByRequester)[0]
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedOrganizationResourceReferencedByRequester() (organization *Organization, err error) {
	if m.IncludedOrganizationResourcesReferencedByRequester == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*m.IncludedOrganizationResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*m.IncludedOrganizationResourcesReferencedByRequester))
	} else if len(*m.IncludedOrganizationResourcesReferencedByRequester) == 1 {
		organization = &(*m.IncludedOrganizationResourcesReferencedByRequester)[0]
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedDeviceResourceReferencedByRequester() (device *Device, err error) {
	if m.IncludedDeviceResourcesReferencedByRequester == nil {
		err = errors.New("Included devices not requested")
	} else if len(*m.IncludedDeviceResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*m.IncludedDeviceResourcesReferencedByRequester))
	} else if len(*m.IncludedDeviceResourcesReferencedByRequester) == 1 {
		device = &(*m.IncludedDeviceResourcesReferencedByRequester)[0]
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByRequester() (patient *Patient, err error) {
	if m.IncludedPatientResourcesReferencedByRequester == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResourcesReferencedByRequester))
	} else if len(*m.IncludedPatientResourcesReferencedByRequester) == 1 {
		patient = &(*m.IncludedPatientResourcesReferencedByRequester)[0]
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByRequester() (relatedPerson *RelatedPerson, err error) {
	if m.IncludedRelatedPersonResourcesReferencedByRequester == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*m.IncludedRelatedPersonResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*m.IncludedRelatedPersonResourcesReferencedByRequester))
	} else if len(*m.IncludedRelatedPersonResourcesReferencedByRequester) == 1 {
		relatedPerson = &(*m.IncludedRelatedPersonResourcesReferencedByRequester)[0]
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedOrganizationResourceReferencedByIntendeddispenser() (organization *Organization, err error) {
	if m.IncludedOrganizationResourcesReferencedByIntendeddispenser == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*m.IncludedOrganizationResourcesReferencedByIntendeddispenser) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*m.IncludedOrganizationResourcesReferencedByIntendeddispenser))
	} else if len(*m.IncludedOrganizationResourcesReferencedByIntendeddispenser) == 1 {
		organization = &(*m.IncludedOrganizationResourcesReferencedByIntendeddispenser)[0]
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if m.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResourcesReferencedByPatient))
	} else if len(*m.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*m.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByContext() (episodeOfCare *EpisodeOfCare, err error) {
	if m.IncludedEpisodeOfCareResourcesReferencedByContext == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*m.IncludedEpisodeOfCareResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*m.IncludedEpisodeOfCareResourcesReferencedByContext))
	} else if len(*m.IncludedEpisodeOfCareResourcesReferencedByContext) == 1 {
		episodeOfCare = &(*m.IncludedEpisodeOfCareResourcesReferencedByContext)[0]
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedEncounterResourceReferencedByContext() (encounter *Encounter, err error) {
	if m.IncludedEncounterResourcesReferencedByContext == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*m.IncludedEncounterResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*m.IncludedEncounterResourcesReferencedByContext))
	} else if len(*m.IncludedEncounterResourcesReferencedByContext) == 1 {
		encounter = &(*m.IncludedEncounterResourcesReferencedByContext)[0]
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedMedicationResourceReferencedByMedication() (medication *Medication, err error) {
	if m.IncludedMedicationResourcesReferencedByMedication == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedMedicationResourcesReferencedByMedication) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedMedicationResourcesReferencedByMedication))
	} else if len(*m.IncludedMedicationResourcesReferencedByMedication) == 1 {
		medication = &(*m.IncludedMedicationResourcesReferencedByMedication)[0]
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if m.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *m.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if m.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *m.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if m.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *m.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if m.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *m.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if m.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *m.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if m.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *m.RevIncludedListResourcesReferencingItem
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if m.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *m.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if m.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *m.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if m.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *m.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if m.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *m.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if m.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *m.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if m.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *m.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPrescription() (medicationAdministrations []MedicationAdministration, err error) {
	if m.RevIncludedMedicationAdministrationResourcesReferencingPrescription == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *m.RevIncludedMedicationAdministrationResourcesReferencingPrescription
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if m.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *m.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPrescription() (medicationDispenses []MedicationDispense, err error) {
	if m.RevIncludedMedicationDispenseResourcesReferencingPrescription == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *m.RevIncludedMedicationDispenseResourcesReferencingPrescription
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if m.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *m.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *m.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if m.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *m.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAction() (clinicalImpressions []ClinicalImpression, err error) {
	if m.RevIncludedClinicalImpressionResourcesReferencingAction == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *m.RevIncludedClinicalImpressionResourcesReferencingAction
	}
	return
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*m.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*m.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *m.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*m.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByRequester {
			rsc := (*m.IncludedPatientResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedRelatedPersonResourcesReferencedByRequester != nil {
		for idx := range *m.IncludedRelatedPersonResourcesReferencedByRequester {
			rsc := (*m.IncludedRelatedPersonResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedByIntendeddispenser != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedByIntendeddispenser {
			rsc := (*m.IncludedOrganizationResourcesReferencedByIntendeddispenser)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPatient {
			rsc := (*m.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*m.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEncounterResourcesReferencedByContext {
			rsc := (*m.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByMedication != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByMedication {
			rsc := (*m.IncludedMedicationResourcesReferencedByMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MedicationRequestPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingData {
			rsc := (*m.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *m.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*m.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationAdministrationResourcesReferencingPrescription != nil {
		for idx := range *m.RevIncludedMedicationAdministrationResourcesReferencingPrescription {
			rsc := (*m.RevIncludedMedicationAdministrationResourcesReferencingPrescription)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationDispenseResourcesReferencingPrescription != nil {
		for idx := range *m.RevIncludedMedicationDispenseResourcesReferencingPrescription {
			rsc := (*m.RevIncludedMedicationDispenseResourcesReferencingPrescription)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MedicationRequestPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*m.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*m.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *m.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*m.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByRequester {
			rsc := (*m.IncludedPatientResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedRelatedPersonResourcesReferencedByRequester != nil {
		for idx := range *m.IncludedRelatedPersonResourcesReferencedByRequester {
			rsc := (*m.IncludedRelatedPersonResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedByIntendeddispenser != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedByIntendeddispenser {
			rsc := (*m.IncludedOrganizationResourcesReferencedByIntendeddispenser)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPatient {
			rsc := (*m.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*m.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEncounterResourcesReferencedByContext {
			rsc := (*m.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByMedication != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByMedication {
			rsc := (*m.IncludedMedicationResourcesReferencedByMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingData {
			rsc := (*m.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingTopic {
			rsc := (*m.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *m.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*m.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *m.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*m.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *m.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*m.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationAdministrationResourcesReferencingPrescription != nil {
		for idx := range *m.RevIncludedMedicationAdministrationResourcesReferencingPrescription {
			rsc := (*m.RevIncludedMedicationAdministrationResourcesReferencingPrescription)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationDispenseResourcesReferencingPrescription != nil {
		for idx := range *m.RevIncludedMedicationDispenseResourcesReferencingPrescription {
			rsc := (*m.RevIncludedMedicationDispenseResourcesReferencingPrescription)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*m.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
