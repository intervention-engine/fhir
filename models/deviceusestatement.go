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

type DeviceUseStatement struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status         string            `bson:"status,omitempty" json:"status,omitempty"`
	Subject        *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	WhenUsed       *Period           `bson:"whenUsed,omitempty" json:"whenUsed,omitempty"`
	TimingTiming   *Timing           `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingPeriod   *Period           `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDateTime *FHIRDateTime     `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	RecordedOn     *FHIRDateTime     `bson:"recordedOn,omitempty" json:"recordedOn,omitempty"`
	Source         *Reference        `bson:"source,omitempty" json:"source,omitempty"`
	Device         *Reference        `bson:"device,omitempty" json:"device,omitempty"`
	Indication     []CodeableConcept `bson:"indication,omitempty" json:"indication,omitempty"`
	BodySite       *CodeableConcept  `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Note           []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DeviceUseStatement) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DeviceUseStatement"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DeviceUseStatement), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DeviceUseStatement) GetBSON() (interface{}, error) {
	x.ResourceType = "DeviceUseStatement"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "deviceUseStatement" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type deviceUseStatement DeviceUseStatement

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DeviceUseStatement) UnmarshalJSON(data []byte) (err error) {
	x2 := deviceUseStatement{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = DeviceUseStatement(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DeviceUseStatement) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DeviceUseStatement"
	} else if x.ResourceType != "DeviceUseStatement" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DeviceUseStatement, instead received %s", x.ResourceType))
	}
	return nil
}

type DeviceUseStatementPlus struct {
	DeviceUseStatement                     `bson:",inline"`
	DeviceUseStatementPlusRelatedResources `bson:",inline"`
}

type DeviceUseStatementPlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                       *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                     *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedGroupResourcesReferencedByPatient                       *[]Group                 `bson:"_includedGroupResourcesReferencedByPatient,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedDeviceResourcesReferencedByDevice                       *[]Device                `bson:"_includedDeviceResourcesReferencedByDevice,omitempty"`
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

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if d.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedGroupResourcesReferencedBySubject))
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*d.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedBySubject))
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedGroupResourceReferencedByPatient() (group *Group, err error) {
	if d.IncludedGroupResourcesReferencedByPatient == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedGroupResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedGroupResourcesReferencedByPatient))
	} else if len(*d.IncludedGroupResourcesReferencedByPatient) == 1 {
		group = &(*d.IncludedGroupResourcesReferencedByPatient)[0]
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedByPatient))
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedDeviceResourceReferencedByDevice() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedByDevice == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedByDevice) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedByDevice))
	} else if len(*d.IncludedDeviceResourcesReferencedByDevice) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedByDevice)[0]
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if d.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *d.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if d.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *d.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if d.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *d.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if d.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *d.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if d.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *d.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if d.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *d.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if d.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *d.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if d.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *d.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if d.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *d.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *d.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if d.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *d.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if d.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *d.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedGroupResourcesReferencedBySubject {
			rsc := (*d.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPatientResourcesReferencedBySubject {
			rsc := (*d.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedGroupResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedGroupResourcesReferencedByPatient {
			rsc := (*d.IncludedGroupResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByDevice != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByDevice {
			rsc := (*d.IncludedDeviceResourcesReferencedByDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
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
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
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
	if d.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*d.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
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
	if d.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
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
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
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
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
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
	if d.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*d.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
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

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedGroupResourcesReferencedBySubject {
			rsc := (*d.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPatientResourcesReferencedBySubject {
			rsc := (*d.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedGroupResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedGroupResourcesReferencedByPatient {
			rsc := (*d.IncludedGroupResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByDevice != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByDevice {
			rsc := (*d.IncludedDeviceResourcesReferencedByDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
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
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
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
	if d.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *d.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*d.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
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
	if d.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
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
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
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
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
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
	if d.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*d.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
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
