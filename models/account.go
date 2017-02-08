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

type Account struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name           string                      `bson:"name,omitempty" json:"name,omitempty"`
	Type           *CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	Status         string                      `bson:"status,omitempty" json:"status,omitempty"`
	Active         *Period                     `bson:"active,omitempty" json:"active,omitempty"`
	Currency       *Coding                     `bson:"currency,omitempty" json:"currency,omitempty"`
	Balance        *Quantity                   `bson:"balance,omitempty" json:"balance,omitempty"`
	Coverage       []Reference                 `bson:"coverage,omitempty" json:"coverage,omitempty"`
	CoveragePeriod *Period                     `bson:"coveragePeriod,omitempty" json:"coveragePeriod,omitempty"`
	Subject        *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Owner          *Reference                  `bson:"owner,omitempty" json:"owner,omitempty"`
	Description    string                      `bson:"description,omitempty" json:"description,omitempty"`
	Guarantor      []AccountGuarantorComponent `bson:"guarantor,omitempty" json:"guarantor,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Account) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Account"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Account), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Account) GetBSON() (interface{}, error) {
	x.ResourceType = "Account"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "account" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type account Account

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Account) UnmarshalJSON(data []byte) (err error) {
	x2 := account{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Account(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Account) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Account"
	} else if x.ResourceType != "Account" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Account, instead received %s", x.ResourceType))
	}
	return nil
}

type AccountGuarantorComponent struct {
	BackboneElement `bson:",inline"`
	Party           *Reference `bson:"party,omitempty" json:"party,omitempty"`
	OnHold          *bool      `bson:"onHold,omitempty" json:"onHold,omitempty"`
	Period          *Period    `bson:"period,omitempty" json:"period,omitempty"`
}

type AccountPlus struct {
	Account                     `bson:",inline"`
	AccountPlusRelatedResources `bson:",inline"`
}

type AccountPlusRelatedResources struct {
	IncludedOrganizationResourcesReferencedByOwner              *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOwner,omitempty"`
	IncludedPractitionerResourcesReferencedBySubject            *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySubject,omitempty"`
	IncludedOrganizationResourcesReferencedBySubject            *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                  *[]Device                `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedHealthcareServiceResourcesReferencedBySubject       *[]HealthcareService     `bson:"_includedHealthcareServiceResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                *[]Location              `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
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

func (a *AccountPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOwner() (organization *Organization, err error) {
	if a.IncludedOrganizationResourcesReferencedByOwner == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*a.IncludedOrganizationResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*a.IncludedOrganizationResourcesReferencedByOwner))
	} else if len(*a.IncludedOrganizationResourcesReferencedByOwner) == 1 {
		organization = &(*a.IncludedOrganizationResourcesReferencedByOwner)[0]
	}
	return
}

func (a *AccountPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySubject() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedBySubject == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedBySubject))
	} else if len(*a.IncludedPractitionerResourcesReferencedBySubject) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedBySubject)[0]
	}
	return
}

func (a *AccountPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySubject() (organization *Organization, err error) {
	if a.IncludedOrganizationResourcesReferencedBySubject == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*a.IncludedOrganizationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*a.IncludedOrganizationResourcesReferencedBySubject))
	} else if len(*a.IncludedOrganizationResourcesReferencedBySubject) == 1 {
		organization = &(*a.IncludedOrganizationResourcesReferencedBySubject)[0]
	}
	return
}

func (a *AccountPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if a.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*a.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*a.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (a *AccountPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedBySubject))
	} else if len(*a.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (a *AccountPlusRelatedResources) GetIncludedHealthcareServiceResourceReferencedBySubject() (healthcareService *HealthcareService, err error) {
	if a.IncludedHealthcareServiceResourcesReferencedBySubject == nil {
		err = errors.New("Included healthcareservices not requested")
	} else if len(*a.IncludedHealthcareServiceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 healthcareService, but found %d", len(*a.IncludedHealthcareServiceResourcesReferencedBySubject))
	} else if len(*a.IncludedHealthcareServiceResourcesReferencedBySubject) == 1 {
		healthcareService = &(*a.IncludedHealthcareServiceResourcesReferencedBySubject)[0]
	}
	return
}

func (a *AccountPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if a.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*a.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*a.IncludedLocationResourcesReferencedBySubject))
	} else if len(*a.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*a.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (a *AccountPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatient))
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if a.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *a.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if a.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *a.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if a.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *a.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if a.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *a.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if a.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *a.RevIncludedListResourcesReferencingItem
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if a.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *a.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if a.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *a.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if a.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *a.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if a.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *a.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if a.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *a.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if a.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *a.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if a.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *a.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if a.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *a.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *a.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (a *AccountPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if a.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *a.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (a *AccountPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedOrganizationResourcesReferencedByOwner != nil {
		for idx := range *a.IncludedOrganizationResourcesReferencedByOwner {
			rsc := (*a.IncludedOrganizationResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedBySubject {
			rsc := (*a.IncludedPractitionerResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedOrganizationResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedOrganizationResourcesReferencedBySubject {
			rsc := (*a.IncludedOrganizationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*a.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedPatientResourcesReferencedBySubject {
			rsc := (*a.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedHealthcareServiceResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedHealthcareServiceResourcesReferencedBySubject {
			rsc := (*a.IncludedHealthcareServiceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedLocationResourcesReferencedBySubject {
			rsc := (*a.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AccountPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	return resourceMap
}

func (a *AccountPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedOrganizationResourcesReferencedByOwner != nil {
		for idx := range *a.IncludedOrganizationResourcesReferencedByOwner {
			rsc := (*a.IncludedOrganizationResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedBySubject {
			rsc := (*a.IncludedPractitionerResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedOrganizationResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedOrganizationResourcesReferencedBySubject {
			rsc := (*a.IncludedOrganizationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*a.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedPatientResourcesReferencedBySubject {
			rsc := (*a.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedHealthcareServiceResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedHealthcareServiceResourcesReferencedBySubject {
			rsc := (*a.IncludedHealthcareServiceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *a.IncludedLocationResourcesReferencedBySubject {
			rsc := (*a.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
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
	return resourceMap
}
