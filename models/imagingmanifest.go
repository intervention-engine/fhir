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

type ImagingManifest struct {
	DomainResource `bson:",inline"`
	Uid            string                          `bson:"uid,omitempty" json:"uid,omitempty"`
	Patient        *Reference                      `bson:"patient,omitempty" json:"patient,omitempty"`
	AuthoringTime  *FHIRDateTime                   `bson:"authoringTime,omitempty" json:"authoringTime,omitempty"`
	Author         *Reference                      `bson:"author,omitempty" json:"author,omitempty"`
	Title          *CodeableConcept                `bson:"title,omitempty" json:"title,omitempty"`
	Description    string                          `bson:"description,omitempty" json:"description,omitempty"`
	Study          []ImagingManifestStudyComponent `bson:"study,omitempty" json:"study,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ImagingManifest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ImagingManifest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ImagingManifest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ImagingManifest) GetBSON() (interface{}, error) {
	x.ResourceType = "ImagingManifest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "imagingManifest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type imagingManifest ImagingManifest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ImagingManifest) UnmarshalJSON(data []byte) (err error) {
	x2 := imagingManifest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ImagingManifest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ImagingManifest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ImagingManifest"
	} else if x.ResourceType != "ImagingManifest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ImagingManifest, instead received %s", x.ResourceType))
	}
	return nil
}

type ImagingManifestStudyComponent struct {
	BackboneElement `bson:",inline"`
	Uid             string                                      `bson:"uid,omitempty" json:"uid,omitempty"`
	ImagingStudy    *Reference                                  `bson:"imagingStudy,omitempty" json:"imagingStudy,omitempty"`
	BaseLocation    []ImagingManifestStudyBaseLocationComponent `bson:"baseLocation,omitempty" json:"baseLocation,omitempty"`
	Series          []ImagingManifestSeriesComponent            `bson:"series,omitempty" json:"series,omitempty"`
}

type ImagingManifestStudyBaseLocationComponent struct {
	BackboneElement `bson:",inline"`
	Type            *Coding `bson:"type,omitempty" json:"type,omitempty"`
	Url             string  `bson:"url,omitempty" json:"url,omitempty"`
}

type ImagingManifestSeriesComponent struct {
	BackboneElement `bson:",inline"`
	Uid             string                                       `bson:"uid,omitempty" json:"uid,omitempty"`
	BaseLocation    []ImagingManifestSeriesBaseLocationComponent `bson:"baseLocation,omitempty" json:"baseLocation,omitempty"`
	Instance        []ImagingManifestInstanceComponent           `bson:"instance,omitempty" json:"instance,omitempty"`
}

type ImagingManifestSeriesBaseLocationComponent struct {
	BackboneElement `bson:",inline"`
	Type            *Coding `bson:"type,omitempty" json:"type,omitempty"`
	Url             string  `bson:"url,omitempty" json:"url,omitempty"`
}

type ImagingManifestInstanceComponent struct {
	BackboneElement `bson:",inline"`
	SopClass        string `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	Uid             string `bson:"uid,omitempty" json:"uid,omitempty"`
}

type ImagingManifestPlus struct {
	ImagingManifest                     `bson:",inline"`
	ImagingManifestPlusRelatedResources `bson:",inline"`
}

type ImagingManifestPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByAuthor             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthor             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByAuthor                   *[]Device                `bson:"_includedDeviceResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByAuthor                  *[]Patient               `bson:"_includedPatientResourcesReferencedByAuthor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAuthor            *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
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

func (i *ImagingManifestPlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if i.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*i.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*i.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*i.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*i.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetIncludedOrganizationResourceReferencedByAuthor() (organization *Organization, err error) {
	if i.IncludedOrganizationResourcesReferencedByAuthor == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*i.IncludedOrganizationResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*i.IncludedOrganizationResourcesReferencedByAuthor))
	} else if len(*i.IncludedOrganizationResourcesReferencedByAuthor) == 1 {
		organization = &(*i.IncludedOrganizationResourcesReferencedByAuthor)[0]
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetIncludedDeviceResourceReferencedByAuthor() (device *Device, err error) {
	if i.IncludedDeviceResourcesReferencedByAuthor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*i.IncludedDeviceResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*i.IncludedDeviceResourcesReferencedByAuthor))
	} else if len(*i.IncludedDeviceResourcesReferencedByAuthor) == 1 {
		device = &(*i.IncludedDeviceResourcesReferencedByAuthor)[0]
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetIncludedPatientResourceReferencedByAuthor() (patient *Patient, err error) {
	if i.IncludedPatientResourcesReferencedByAuthor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResourcesReferencedByAuthor))
	} else if len(*i.IncludedPatientResourcesReferencedByAuthor) == 1 {
		patient = &(*i.IncludedPatientResourcesReferencedByAuthor)[0]
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAuthor() (relatedPerson *RelatedPerson, err error) {
	if i.IncludedRelatedPersonResourcesReferencedByAuthor == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*i.IncludedRelatedPersonResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*i.IncludedRelatedPersonResourcesReferencedByAuthor))
	} else if len(*i.IncludedRelatedPersonResourcesReferencedByAuthor) == 1 {
		relatedPerson = &(*i.IncludedRelatedPersonResourcesReferencedByAuthor)[0]
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if i.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResourcesReferencedByPatient))
	} else if len(*i.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*i.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if i.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *i.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if i.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *i.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if i.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *i.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if i.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *i.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if i.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *i.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if i.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *i.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if i.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *i.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if i.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *i.RevIncludedListResourcesReferencingItem
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if i.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *i.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if i.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *i.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if i.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *i.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if i.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *i.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if i.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *i.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if i.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *i.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if i.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *i.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if i.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *i.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *i.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if i.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *i.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (i *ImagingManifestPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *i.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*i.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *i.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*i.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *i.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*i.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *i.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*i.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *i.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*i.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *i.IncludedPatientResourcesReferencedByPatient {
			rsc := (*i.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (i *ImagingManifestPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *i.RevIncludedConsentResourcesReferencingData {
			rsc := (*i.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*i.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingSubject {
			rsc := (*i.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingTopic {
			rsc := (*i.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*i.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*i.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*i.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*i.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*i.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*i.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*i.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*i.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*i.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*i.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (i *ImagingManifestPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *i.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*i.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *i.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*i.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *i.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*i.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *i.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*i.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *i.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*i.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *i.IncludedPatientResourcesReferencedByPatient {
			rsc := (*i.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *i.RevIncludedConsentResourcesReferencingData {
			rsc := (*i.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*i.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingSubject {
			rsc := (*i.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingTopic {
			rsc := (*i.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*i.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*i.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*i.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *i.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*i.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *i.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*i.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*i.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*i.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*i.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*i.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*i.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*i.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*i.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
