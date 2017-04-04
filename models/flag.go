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

type Flag struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status         string           `bson:"status,omitempty" json:"status,omitempty"`
	Category       *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code           *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Subject        *Reference       `bson:"subject,omitempty" json:"subject,omitempty"`
	Period         *Period          `bson:"period,omitempty" json:"period,omitempty"`
	Encounter      *Reference       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Author         *Reference       `bson:"author,omitempty" json:"author,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Flag) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Flag"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Flag), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Flag) GetBSON() (interface{}, error) {
	x.ResourceType = "Flag"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "flag" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type flag Flag

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Flag) UnmarshalJSON(data []byte) (err error) {
	x2 := flag{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Flag(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Flag) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Flag"
	} else if x.ResourceType != "Flag" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Flag, instead received %s", x.ResourceType))
	}
	return nil
}

type FlagPlus struct {
	Flag                     `bson:",inline"`
	FlagPlusRelatedResources `bson:",inline"`
}

type FlagPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedBySubject                *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySubject,omitempty"`
	IncludedGroupResourcesReferencedBySubject                       *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedOrganizationResourcesReferencedBySubject                *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySubject,omitempty"`
	IncludedMedicationResourcesReferencedBySubject                  *[]Medication            `bson:"_includedMedicationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                     *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPlanDefinitionResourcesReferencedBySubject              *[]PlanDefinition        `bson:"_includedPlanDefinitionResourcesReferencedBySubject,omitempty"`
	IncludedProcedureResourcesReferencedBySubject                   *[]Procedure             `bson:"_includedProcedureResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                    *[]Location              `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthor                 *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthor                 *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByAuthor                       *[]Device                `bson:"_includedDeviceResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByAuthor                      *[]Patient               `bson:"_includedPatientResourcesReferencedByAuthor,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                 *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
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
	RevIncludedDeviceRequestResourcesReferencingBasedon             *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest        *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                 *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                        *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces         *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon          *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
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
	RevIncludedPlanDefinitionResourcesReferencingSuccessor          *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof         *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (f *FlagPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySubject() (practitioner *Practitioner, err error) {
	if f.IncludedPractitionerResourcesReferencedBySubject == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*f.IncludedPractitionerResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*f.IncludedPractitionerResourcesReferencedBySubject))
	} else if len(*f.IncludedPractitionerResourcesReferencedBySubject) == 1 {
		practitioner = &(*f.IncludedPractitionerResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if f.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*f.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*f.IncludedGroupResourcesReferencedBySubject))
	} else if len(*f.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*f.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySubject() (organization *Organization, err error) {
	if f.IncludedOrganizationResourcesReferencedBySubject == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*f.IncludedOrganizationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*f.IncludedOrganizationResourcesReferencedBySubject))
	} else if len(*f.IncludedOrganizationResourcesReferencedBySubject) == 1 {
		organization = &(*f.IncludedOrganizationResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedMedicationResourceReferencedBySubject() (medication *Medication, err error) {
	if f.IncludedMedicationResourcesReferencedBySubject == nil {
		err = errors.New("Included medications not requested")
	} else if len(*f.IncludedMedicationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*f.IncludedMedicationResourcesReferencedBySubject))
	} else if len(*f.IncludedMedicationResourcesReferencedBySubject) == 1 {
		medication = &(*f.IncludedMedicationResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if f.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedPatientResourcesReferencedBySubject))
	} else if len(*f.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*f.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedPlanDefinitionResourceReferencedBySubject() (planDefinition *PlanDefinition, err error) {
	if f.IncludedPlanDefinitionResourcesReferencedBySubject == nil {
		err = errors.New("Included plandefinitions not requested")
	} else if len(*f.IncludedPlanDefinitionResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 planDefinition, but found %d", len(*f.IncludedPlanDefinitionResourcesReferencedBySubject))
	} else if len(*f.IncludedPlanDefinitionResourcesReferencedBySubject) == 1 {
		planDefinition = &(*f.IncludedPlanDefinitionResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedProcedureResourceReferencedBySubject() (procedure *Procedure, err error) {
	if f.IncludedProcedureResourcesReferencedBySubject == nil {
		err = errors.New("Included procedures not requested")
	} else if len(*f.IncludedProcedureResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 procedure, but found %d", len(*f.IncludedProcedureResourcesReferencedBySubject))
	} else if len(*f.IncludedProcedureResourcesReferencedBySubject) == 1 {
		procedure = &(*f.IncludedProcedureResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if f.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*f.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*f.IncludedLocationResourcesReferencedBySubject))
	} else if len(*f.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*f.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if f.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedPatientResourcesReferencedByPatient))
	} else if len(*f.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*f.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if f.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*f.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*f.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*f.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*f.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedOrganizationResourceReferencedByAuthor() (organization *Organization, err error) {
	if f.IncludedOrganizationResourcesReferencedByAuthor == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*f.IncludedOrganizationResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*f.IncludedOrganizationResourcesReferencedByAuthor))
	} else if len(*f.IncludedOrganizationResourcesReferencedByAuthor) == 1 {
		organization = &(*f.IncludedOrganizationResourcesReferencedByAuthor)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedDeviceResourceReferencedByAuthor() (device *Device, err error) {
	if f.IncludedDeviceResourcesReferencedByAuthor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*f.IncludedDeviceResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*f.IncludedDeviceResourcesReferencedByAuthor))
	} else if len(*f.IncludedDeviceResourcesReferencedByAuthor) == 1 {
		device = &(*f.IncludedDeviceResourcesReferencedByAuthor)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedPatientResourceReferencedByAuthor() (patient *Patient, err error) {
	if f.IncludedPatientResourcesReferencedByAuthor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedPatientResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedPatientResourcesReferencedByAuthor))
	} else if len(*f.IncludedPatientResourcesReferencedByAuthor) == 1 {
		patient = &(*f.IncludedPatientResourcesReferencedByAuthor)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if f.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*f.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*f.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*f.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*f.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if f.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *f.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *f.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if f.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *f.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if f.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *f.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if f.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *f.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if f.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *f.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if f.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *f.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if f.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *f.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if f.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *f.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if f.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *f.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if f.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *f.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if f.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *f.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if f.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *f.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if f.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *f.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if f.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *f.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if f.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *f.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if f.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *f.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if f.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *f.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if f.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *f.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if f.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *f.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if f.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *f.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if f.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *f.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if f.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *f.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if f.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *f.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if f.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *f.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if f.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *f.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if f.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *f.RevIncludedListResourcesReferencingItem
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if f.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *f.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if f.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *f.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if f.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *f.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if f.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *f.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if f.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *f.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if f.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *f.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if f.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *f.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if f.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *f.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *f.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if f.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *f.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedPractitionerResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPractitionerResourcesReferencedBySubject {
			rsc := (*f.IncludedPractitionerResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedGroupResourcesReferencedBySubject {
			rsc := (*f.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedOrganizationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedOrganizationResourcesReferencedBySubject {
			rsc := (*f.IncludedOrganizationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedMedicationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedMedicationResourcesReferencedBySubject {
			rsc := (*f.IncludedMedicationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPatientResourcesReferencedBySubject {
			rsc := (*f.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPlanDefinitionResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPlanDefinitionResourcesReferencedBySubject {
			rsc := (*f.IncludedPlanDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedProcedureResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedProcedureResourcesReferencedBySubject {
			rsc := (*f.IncludedProcedureResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedLocationResourcesReferencedBySubject {
			rsc := (*f.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByPatient {
			rsc := (*f.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*f.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*f.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*f.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*f.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *f.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*f.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (f *FlagPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *f.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*f.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *f.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*f.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*f.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedContractResourcesReferencingSubject {
			rsc := (*f.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *f.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*f.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *f.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*f.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*f.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *f.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*f.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*f.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*f.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*f.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*f.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedListResourcesReferencingItem {
			rsc := (*f.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *f.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*f.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*f.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*f.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *f.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*f.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *f.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*f.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*f.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*f.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *f.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*f.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*f.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *f.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*f.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (f *FlagPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedPractitionerResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPractitionerResourcesReferencedBySubject {
			rsc := (*f.IncludedPractitionerResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedGroupResourcesReferencedBySubject {
			rsc := (*f.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedOrganizationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedOrganizationResourcesReferencedBySubject {
			rsc := (*f.IncludedOrganizationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedMedicationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedMedicationResourcesReferencedBySubject {
			rsc := (*f.IncludedMedicationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPatientResourcesReferencedBySubject {
			rsc := (*f.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPlanDefinitionResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPlanDefinitionResourcesReferencedBySubject {
			rsc := (*f.IncludedPlanDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedProcedureResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedProcedureResourcesReferencedBySubject {
			rsc := (*f.IncludedProcedureResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedLocationResourcesReferencedBySubject {
			rsc := (*f.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByPatient {
			rsc := (*f.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*f.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*f.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*f.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*f.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *f.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*f.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *f.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*f.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *f.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*f.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*f.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedContractResourcesReferencingSubject {
			rsc := (*f.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *f.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*f.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *f.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*f.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*f.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *f.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*f.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*f.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*f.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*f.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*f.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedListResourcesReferencingItem {
			rsc := (*f.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *f.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*f.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*f.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*f.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *f.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*f.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *f.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*f.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*f.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*f.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *f.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*f.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*f.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *f.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*f.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
