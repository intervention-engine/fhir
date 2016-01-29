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

type Flag struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Category       *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Status         string           `bson:"status,omitempty" json:"status,omitempty"`
	Period         *Period          `bson:"period,omitempty" json:"period,omitempty"`
	Subject        *Reference       `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter      *Reference       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Author         *Reference       `bson:"author,omitempty" json:"author,omitempty"`
	Code           *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
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
	IncludedPractitionerResourcesReferencedBySubject            *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySubject,omitempty"`
	IncludedGroupResourcesReferencedBySubject                   *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedOrganizationResourcesReferencedBySubject            *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                *[]Location              `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthor             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthor             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByAuthor                   *[]Device                `bson:"_includedDeviceResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByAuthor                  *[]Patient               `bson:"_includedPatientResourcesReferencedByAuthor,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter             *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
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

func (f *FlagPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if f.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *f.RevIncludedProvenanceResourcesReferencingTarget
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

func (f *FlagPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if f.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *f.RevIncludedListResourcesReferencingItem
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

func (f *FlagPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if f.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *f.RevIncludedOrderResourcesReferencingDetail
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

func (f *FlagPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if f.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *f.RevIncludedAuditEventResourcesReferencingReference
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

func (f *FlagPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if f.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *f.RevIncludedOrderResponseResourcesReferencingFulfillment
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

func (f *FlagPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if f.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *f.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if f.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *f.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedPractitionerResourcesReferencedBySubject != nil {
		for _, r := range *f.IncludedPractitionerResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *f.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedOrganizationResourcesReferencedBySubject != nil {
		for _, r := range *f.IncludedOrganizationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *f.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedLocationResourcesReferencedBySubject != nil {
		for _, r := range *f.IncludedLocationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *f.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for _, r := range *f.IncludedPractitionerResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for _, r := range *f.IncludedOrganizationResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedDeviceResourcesReferencedByAuthor != nil {
		for _, r := range *f.IncludedDeviceResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedPatientResourcesReferencedByAuthor != nil {
		for _, r := range *f.IncludedPatientResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *f.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (f *FlagPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *f.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *f.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *f.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *f.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *f.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *f.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *f.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *f.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *f.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *f.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *f.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *f.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *f.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *f.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *f.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (f *FlagPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedPractitionerResourcesReferencedBySubject != nil {
		for _, r := range *f.IncludedPractitionerResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *f.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedOrganizationResourcesReferencedBySubject != nil {
		for _, r := range *f.IncludedOrganizationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *f.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedLocationResourcesReferencedBySubject != nil {
		for _, r := range *f.IncludedLocationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *f.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for _, r := range *f.IncludedPractitionerResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for _, r := range *f.IncludedOrganizationResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedDeviceResourcesReferencedByAuthor != nil {
		for _, r := range *f.IncludedDeviceResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedPatientResourcesReferencedByAuthor != nil {
		for _, r := range *f.IncludedPatientResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if f.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *f.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *f.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *f.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *f.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *f.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *f.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *f.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *f.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *f.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *f.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *f.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *f.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *f.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *f.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *f.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if f.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *f.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
