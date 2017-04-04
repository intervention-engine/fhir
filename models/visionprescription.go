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

type VisionPrescription struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                string                                `bson:"status,omitempty" json:"status,omitempty"`
	Patient               *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	Encounter             *Reference                            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateWritten           *FHIRDateTime                         `bson:"dateWritten,omitempty" json:"dateWritten,omitempty"`
	Prescriber            *Reference                            `bson:"prescriber,omitempty" json:"prescriber,omitempty"`
	ReasonCodeableConcept *CodeableConcept                      `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                            `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Dispense              []VisionPrescriptionDispenseComponent `bson:"dispense,omitempty" json:"dispense,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *VisionPrescription) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "VisionPrescription"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to VisionPrescription), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *VisionPrescription) GetBSON() (interface{}, error) {
	x.ResourceType = "VisionPrescription"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "visionPrescription" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type visionPrescription VisionPrescription

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *VisionPrescription) UnmarshalJSON(data []byte) (err error) {
	x2 := visionPrescription{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = VisionPrescription(x2)
		return x.checkResourceType()
	}
	return
}

func (x *VisionPrescription) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "VisionPrescription"
	} else if x.ResourceType != "VisionPrescription" {
		return errors.New(fmt.Sprintf("Expected resourceType to be VisionPrescription, instead received %s", x.ResourceType))
	}
	return nil
}

type VisionPrescriptionDispenseComponent struct {
	BackboneElement `bson:",inline"`
	Product         *CodeableConcept `bson:"product,omitempty" json:"product,omitempty"`
	Eye             string           `bson:"eye,omitempty" json:"eye,omitempty"`
	Sphere          *float64         `bson:"sphere,omitempty" json:"sphere,omitempty"`
	Cylinder        *float64         `bson:"cylinder,omitempty" json:"cylinder,omitempty"`
	Axis            *int32           `bson:"axis,omitempty" json:"axis,omitempty"`
	Prism           *float64         `bson:"prism,omitempty" json:"prism,omitempty"`
	Base            string           `bson:"base,omitempty" json:"base,omitempty"`
	Add             *float64         `bson:"add,omitempty" json:"add,omitempty"`
	Power           *float64         `bson:"power,omitempty" json:"power,omitempty"`
	BackCurve       *float64         `bson:"backCurve,omitempty" json:"backCurve,omitempty"`
	Diameter        *float64         `bson:"diameter,omitempty" json:"diameter,omitempty"`
	Duration        *Quantity        `bson:"duration,omitempty" json:"duration,omitempty"`
	Color           string           `bson:"color,omitempty" json:"color,omitempty"`
	Brand           string           `bson:"brand,omitempty" json:"brand,omitempty"`
	Note            []Annotation     `bson:"note,omitempty" json:"note,omitempty"`
}

type VisionPrescriptionPlus struct {
	VisionPrescription                     `bson:",inline"`
	VisionPrescriptionPlusRelatedResources `bson:",inline"`
}

type VisionPrescriptionPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByPrescriber             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByPrescriber,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
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
	RevIncludedCarePlanResourcesReferencingActivityreference        *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
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

func (v *VisionPrescriptionPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPrescriber() (practitioner *Practitioner, err error) {
	if v.IncludedPractitionerResourcesReferencedByPrescriber == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*v.IncludedPractitionerResourcesReferencedByPrescriber) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*v.IncludedPractitionerResourcesReferencedByPrescriber))
	} else if len(*v.IncludedPractitionerResourcesReferencedByPrescriber) == 1 {
		practitioner = &(*v.IncludedPractitionerResourcesReferencedByPrescriber)[0]
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if v.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*v.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*v.IncludedPatientResourcesReferencedByPatient))
	} else if len(*v.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*v.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if v.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*v.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*v.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*v.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*v.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if v.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *v.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *v.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if v.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *v.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if v.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *v.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if v.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *v.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if v.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *v.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *v.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *v.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if v.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *v.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if v.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *v.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if v.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *v.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if v.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *v.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if v.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *v.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if v.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *v.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if v.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *v.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if v.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *v.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if v.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *v.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if v.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *v.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if v.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *v.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if v.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *v.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if v.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *v.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if v.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *v.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if v.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *v.RevIncludedListResourcesReferencingItem
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if v.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *v.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if v.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *v.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if v.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *v.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if v.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *v.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if v.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *v.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if v.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *v.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if v.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *v.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if v.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *v.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *v.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if v.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *v.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.IncludedPractitionerResourcesReferencedByPrescriber != nil {
		for idx := range *v.IncludedPractitionerResourcesReferencedByPrescriber {
			rsc := (*v.IncludedPractitionerResourcesReferencedByPrescriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *v.IncludedPatientResourcesReferencedByPatient {
			rsc := (*v.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *v.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*v.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *v.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*v.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *v.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*v.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*v.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedContractResourcesReferencingSubject {
			rsc := (*v.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *v.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*v.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *v.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*v.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *v.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*v.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*v.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*v.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*v.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*v.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *v.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*v.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedListResourcesReferencingItem {
			rsc := (*v.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *v.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*v.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*v.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*v.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *v.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*v.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *v.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*v.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*v.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*v.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *v.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*v.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*v.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *v.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*v.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (v *VisionPrescriptionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.IncludedPractitionerResourcesReferencedByPrescriber != nil {
		for idx := range *v.IncludedPractitionerResourcesReferencedByPrescriber {
			rsc := (*v.IncludedPractitionerResourcesReferencedByPrescriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *v.IncludedPatientResourcesReferencedByPatient {
			rsc := (*v.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *v.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*v.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *v.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*v.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *v.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*v.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*v.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedContractResourcesReferencingSubject {
			rsc := (*v.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *v.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*v.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *v.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*v.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*v.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *v.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*v.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*v.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*v.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*v.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*v.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *v.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*v.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedListResourcesReferencingItem {
			rsc := (*v.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *v.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*v.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*v.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*v.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *v.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*v.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *v.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*v.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*v.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*v.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *v.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*v.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*v.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *v.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*v.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
