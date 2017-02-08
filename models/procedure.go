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

type Procedure struct {
	DomainResource     `bson:",inline"`
	Identifier         []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status             string                          `bson:"status,omitempty" json:"status,omitempty"`
	Category           *CodeableConcept                `bson:"category,omitempty" json:"category,omitempty"`
	Code               *CodeableConcept                `bson:"code,omitempty" json:"code,omitempty"`
	Subject            *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter          *Reference                      `bson:"encounter,omitempty" json:"encounter,omitempty"`
	PerformedDateTime  *FHIRDateTime                   `bson:"performedDateTime,omitempty" json:"performedDateTime,omitempty"`
	PerformedPeriod    *Period                         `bson:"performedPeriod,omitempty" json:"performedPeriod,omitempty"`
	Performer          []ProcedurePerformerComponent   `bson:"performer,omitempty" json:"performer,omitempty"`
	Location           *Reference                      `bson:"location,omitempty" json:"location,omitempty"`
	ReasonReference    []Reference                     `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	ReasonCode         []CodeableConcept               `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	NotPerformed       *bool                           `bson:"notPerformed,omitempty" json:"notPerformed,omitempty"`
	ReasonNotPerformed []CodeableConcept               `bson:"reasonNotPerformed,omitempty" json:"reasonNotPerformed,omitempty"`
	BodySite           []CodeableConcept               `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Outcome            *CodeableConcept                `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Report             []Reference                     `bson:"report,omitempty" json:"report,omitempty"`
	Complication       []CodeableConcept               `bson:"complication,omitempty" json:"complication,omitempty"`
	FollowUp           []CodeableConcept               `bson:"followUp,omitempty" json:"followUp,omitempty"`
	Request            *Reference                      `bson:"request,omitempty" json:"request,omitempty"`
	Notes              []Annotation                    `bson:"notes,omitempty" json:"notes,omitempty"`
	FocalDevice        []ProcedureFocalDeviceComponent `bson:"focalDevice,omitempty" json:"focalDevice,omitempty"`
	UsedReference      []Reference                     `bson:"usedReference,omitempty" json:"usedReference,omitempty"`
	UsedCode           []CodeableConcept               `bson:"usedCode,omitempty" json:"usedCode,omitempty"`
	Component          []Reference                     `bson:"component,omitempty" json:"component,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Procedure) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Procedure"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Procedure), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Procedure) GetBSON() (interface{}, error) {
	x.ResourceType = "Procedure"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "procedure" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type procedure Procedure

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Procedure) UnmarshalJSON(data []byte) (err error) {
	x2 := procedure{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Procedure(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Procedure) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Procedure"
	} else if x.ResourceType != "Procedure" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Procedure, instead received %s", x.ResourceType))
	}
	return nil
}

type ProcedurePerformerComponent struct {
	BackboneElement `bson:",inline"`
	Actor           *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ProcedureFocalDeviceComponent struct {
	BackboneElement `bson:",inline"`
	Action          *CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Manipulated     *Reference       `bson:"manipulated,omitempty" json:"manipulated,omitempty"`
}

type ProcedurePlus struct {
	Procedure                     `bson:",inline"`
	ProcedurePlusRelatedResources `bson:",inline"`
}

type ProcedurePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByPerformer          *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer          *[]Organization          `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer               *[]Patient               `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer         *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedGroupResourcesReferencedBySubject                   *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedLocationResourcesReferencedByLocation               *[]Location              `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter             *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                  *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic               *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject              *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse        *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource  *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedEncounterResourcesReferencingProcedure           *[]Encounter             `bson:"_revIncludedEncounterResourcesReferencingProcedure,omitempty"`
	RevIncludedEncounterResourcesReferencingIndication          *[]Encounter             `bson:"_revIncludedEncounterResourcesReferencingIndication,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon         *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                    *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces    *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition  *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces     *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition   *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                  *[]Flag                  `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity             *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingParent  *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingParent,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAction     *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
}

func (p *ProcedurePlusRelatedResources) GetIncludedPractitionerResourceReferencedByPerformer() (practitioner *Practitioner, err error) {
	if p.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedPractitionerResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedPractitionerResourcesReferencedByPerformer))
	} else if len(*p.IncludedPractitionerResourcesReferencedByPerformer) == 1 {
		practitioner = &(*p.IncludedPractitionerResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedOrganizationResourceReferencedByPerformer() (organization *Organization, err error) {
	if p.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResourcesReferencedByPerformer))
	} else if len(*p.IncludedOrganizationResourcesReferencedByPerformer) == 1 {
		organization = &(*p.IncludedOrganizationResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPatientResourceReferencedByPerformer() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedByPerformer))
	} else if len(*p.IncludedPatientResourcesReferencedByPerformer) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByPerformer() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedRelatedPersonResourcesReferencedByPerformer))
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByPerformer) == 1 {
		relatedPerson = &(*p.IncludedRelatedPersonResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if p.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*p.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*p.IncludedGroupResourcesReferencedBySubject))
	} else if len(*p.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*p.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedBySubject))
	} else if len(*p.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedByPatient))
	} else if len(*p.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if p.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*p.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*p.IncludedLocationResourcesReferencedByLocation))
	} else if len(*p.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*p.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if p.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*p.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*p.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*p.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*p.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if p.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *p.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEncounterResourcesReferencingProcedure() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingProcedure == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingProcedure
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEncounterResourcesReferencingIndication() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingIndication == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingIndication
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingItem
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if p.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *p.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if p.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *p.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if p.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *p.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if p.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *p.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if p.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *p.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if p.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *p.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingParent() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingParent == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingParent
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if p.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *p.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAction() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingAction == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingAction
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*p.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*p.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*p.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedGroupResourcesReferencedBySubject {
			rsc := (*p.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedPatientResourcesReferencedBySubject {
			rsc := (*p.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPatient {
			rsc := (*p.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *p.IncludedLocationResourcesReferencedByLocation {
			rsc := (*p.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *p.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*p.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingData {
			rsc := (*p.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingProcedure != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingProcedure {
			rsc := (*p.RevIncludedEncounterResourcesReferencingProcedure)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingIndication != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingIndication {
			rsc := (*p.RevIncludedEncounterResourcesReferencingIndication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingParent != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingParent {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*p.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *ProcedurePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*p.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*p.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*p.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedGroupResourcesReferencedBySubject {
			rsc := (*p.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedPatientResourcesReferencedBySubject {
			rsc := (*p.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPatient {
			rsc := (*p.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *p.IncludedLocationResourcesReferencedByLocation {
			rsc := (*p.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *p.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*p.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingData {
			rsc := (*p.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingTopic {
			rsc := (*p.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingProcedure != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingProcedure {
			rsc := (*p.RevIncludedEncounterResourcesReferencingProcedure)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingIndication != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingIndication {
			rsc := (*p.RevIncludedEncounterResourcesReferencingIndication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*p.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *p.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*p.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingParent != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingParent {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*p.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
