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

type Provenance struct {
	DomainResource `bson:",inline"`
	Target         []Reference                 `bson:"target,omitempty" json:"target,omitempty"`
	Period         *Period                     `bson:"period,omitempty" json:"period,omitempty"`
	Recorded       *FHIRDateTime               `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Reason         []Coding                    `bson:"reason,omitempty" json:"reason,omitempty"`
	Activity       *Coding                     `bson:"activity,omitempty" json:"activity,omitempty"`
	Location       *Reference                  `bson:"location,omitempty" json:"location,omitempty"`
	Policy         []string                    `bson:"policy,omitempty" json:"policy,omitempty"`
	Agent          []ProvenanceAgentComponent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Entity         []ProvenanceEntityComponent `bson:"entity,omitempty" json:"entity,omitempty"`
	Signature      []Signature                 `bson:"signature,omitempty" json:"signature,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Provenance) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Provenance"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Provenance), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Provenance) GetBSON() (interface{}, error) {
	x.ResourceType = "Provenance"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "provenance" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type provenance Provenance

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Provenance) UnmarshalJSON(data []byte) (err error) {
	x2 := provenance{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Provenance(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Provenance) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Provenance"
	} else if x.ResourceType != "Provenance" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Provenance, instead received %s", x.ResourceType))
	}
	return nil
}

type ProvenanceAgentComponent struct {
	BackboneElement     `bson:",inline"`
	Role                *Coding          `bson:"role,omitempty" json:"role,omitempty"`
	WhoUri              string           `bson:"whoUri,omitempty" json:"whoUri,omitempty"`
	WhoReference        *Reference       `bson:"whoReference,omitempty" json:"whoReference,omitempty"`
	OnBehalfOfUri       string           `bson:"onBehalfOfUri,omitempty" json:"onBehalfOfUri,omitempty"`
	OnBehalfOfReference *Reference       `bson:"onBehalfOfReference,omitempty" json:"onBehalfOfReference,omitempty"`
	RelatedAgentType    *CodeableConcept `bson:"relatedAgentType,omitempty" json:"relatedAgentType,omitempty"`
}

type ProvenanceEntityComponent struct {
	BackboneElement `bson:",inline"`
	Role            string                     `bson:"role,omitempty" json:"role,omitempty"`
	Reference       *Reference                 `bson:"reference,omitempty" json:"reference,omitempty"`
	Agent           []ProvenanceAgentComponent `bson:"agent,omitempty" json:"agent,omitempty"`
}

type ProvenancePlus struct {
	Provenance                     `bson:",inline"`
	ProvenancePlusRelatedResources `bson:",inline"`
}

type ProvenancePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByAgent              *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAgent,omitempty"`
	IncludedOrganizationResourcesReferencedByAgent              *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAgent,omitempty"`
	IncludedDeviceResourcesReferencedByAgent                    *[]Device                `bson:"_includedDeviceResourcesReferencedByAgent,omitempty"`
	IncludedPatientResourcesReferencedByAgent                   *[]Patient               `bson:"_includedPatientResourcesReferencedByAgent,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAgent             *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByAgent,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedLocationResourcesReferencedByLocation               *[]Location              `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
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
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity             *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
}

func (p *ProvenancePlusRelatedResources) GetIncludedPractitionerResourceReferencedByAgent() (practitioner *Practitioner, err error) {
	if p.IncludedPractitionerResourcesReferencedByAgent == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedPractitionerResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedPractitionerResourcesReferencedByAgent))
	} else if len(*p.IncludedPractitionerResourcesReferencedByAgent) == 1 {
		practitioner = &(*p.IncludedPractitionerResourcesReferencedByAgent)[0]
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetIncludedOrganizationResourceReferencedByAgent() (organization *Organization, err error) {
	if p.IncludedOrganizationResourcesReferencedByAgent == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResourcesReferencedByAgent))
	} else if len(*p.IncludedOrganizationResourcesReferencedByAgent) == 1 {
		organization = &(*p.IncludedOrganizationResourcesReferencedByAgent)[0]
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetIncludedDeviceResourceReferencedByAgent() (device *Device, err error) {
	if p.IncludedDeviceResourcesReferencedByAgent == nil {
		err = errors.New("Included devices not requested")
	} else if len(*p.IncludedDeviceResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*p.IncludedDeviceResourcesReferencedByAgent))
	} else if len(*p.IncludedDeviceResourcesReferencedByAgent) == 1 {
		device = &(*p.IncludedDeviceResourcesReferencedByAgent)[0]
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetIncludedPatientResourceReferencedByAgent() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedByAgent == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedByAgent))
	} else if len(*p.IncludedPatientResourcesReferencedByAgent) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedByAgent)[0]
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAgent() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedRelatedPersonResourcesReferencedByAgent == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedRelatedPersonResourcesReferencedByAgent))
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByAgent) == 1 {
		relatedPerson = &(*p.IncludedRelatedPersonResourcesReferencedByAgent)[0]
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetIncludedPatientResourcesReferencedByPatient() (patients []Patient, err error) {
	if p.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *p.IncludedPatientResourcesReferencedByPatient
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if p.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*p.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*p.IncludedLocationResourcesReferencedByLocation))
	} else if len(*p.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*p.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if p.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *p.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingItem
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if p.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *p.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if p.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *p.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if p.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *p.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if p.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *p.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if p.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *p.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if p.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *p.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if p.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *p.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (p *ProvenancePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByAgent != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByAgent {
			rsc := (*p.IncludedPractitionerResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByAgent != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByAgent {
			rsc := (*p.IncludedOrganizationResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedDeviceResourcesReferencedByAgent != nil {
		for idx := range *p.IncludedDeviceResourcesReferencedByAgent {
			rsc := (*p.IncludedDeviceResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByAgent != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByAgent {
			rsc := (*p.IncludedPatientResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByAgent != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByAgent {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByAgent)[idx]
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
	return resourceMap
}

func (p *ProvenancePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	return resourceMap
}

func (p *ProvenancePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByAgent != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByAgent {
			rsc := (*p.IncludedPractitionerResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByAgent != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByAgent {
			rsc := (*p.IncludedOrganizationResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedDeviceResourcesReferencedByAgent != nil {
		for idx := range *p.IncludedDeviceResourcesReferencedByAgent {
			rsc := (*p.IncludedDeviceResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByAgent != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByAgent {
			rsc := (*p.IncludedPatientResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByAgent != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByAgent {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByAgent)[idx]
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
	return resourceMap
}
