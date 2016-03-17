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
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name           string           `bson:"name,omitempty" json:"name,omitempty"`
	Type           *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Status         string           `bson:"status,omitempty" json:"status,omitempty"`
	ActivePeriod   *Period          `bson:"activePeriod,omitempty" json:"activePeriod,omitempty"`
	Currency       *Coding          `bson:"currency,omitempty" json:"currency,omitempty"`
	Balance        *Quantity        `bson:"balance,omitempty" json:"balance,omitempty"`
	CoveragePeriod *Period          `bson:"coveragePeriod,omitempty" json:"coveragePeriod,omitempty"`
	Subject        *Reference       `bson:"subject,omitempty" json:"subject,omitempty"`
	Owner          *Reference       `bson:"owner,omitempty" json:"owner,omitempty"`
	Description    string           `bson:"description,omitempty" json:"description,omitempty"`
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
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
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

func (a *AccountPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingTarget
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

func (a *AccountPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if a.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *a.RevIncludedListResourcesReferencingItem
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

func (a *AccountPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if a.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *a.RevIncludedOrderResourcesReferencingDetail
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

func (a *AccountPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if a.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *a.RevIncludedAuditEventResourcesReferencingReference
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

func (a *AccountPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *a.RevIncludedOrderResponseResourcesReferencingFulfillment
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

func (a *AccountPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *a.RevIncludedClinicalImpressionResourcesReferencingTrigger
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
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *a.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*a.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *a.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*a.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingData)[idx]
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
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *a.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*a.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *a.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*a.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
