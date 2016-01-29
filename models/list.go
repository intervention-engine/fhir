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

type List struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Title          string               `bson:"title,omitempty" json:"title,omitempty"`
	Code           *CodeableConcept     `bson:"code,omitempty" json:"code,omitempty"`
	Subject        *Reference           `bson:"subject,omitempty" json:"subject,omitempty"`
	Source         *Reference           `bson:"source,omitempty" json:"source,omitempty"`
	Encounter      *Reference           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Status         string               `bson:"status,omitempty" json:"status,omitempty"`
	Date           *FHIRDateTime        `bson:"date,omitempty" json:"date,omitempty"`
	OrderedBy      *CodeableConcept     `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Mode           string               `bson:"mode,omitempty" json:"mode,omitempty"`
	Note           string               `bson:"note,omitempty" json:"note,omitempty"`
	Entry          []ListEntryComponent `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason    *CodeableConcept     `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *List) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "List"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to List), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *List) GetBSON() (interface{}, error) {
	x.ResourceType = "List"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "list" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type list List

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *List) UnmarshalJSON(data []byte) (err error) {
	x2 := list{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = List(x2)
		return x.checkResourceType()
	}
	return
}

func (x *List) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "List"
	} else if x.ResourceType != "List" {
		return errors.New(fmt.Sprintf("Expected resourceType to be List, instead received %s", x.ResourceType))
	}
	return nil
}

type ListEntryComponent struct {
	Flag    *CodeableConcept `bson:"flag,omitempty" json:"flag,omitempty"`
	Deleted *bool            `bson:"deleted,omitempty" json:"deleted,omitempty"`
	Date    *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Item    *Reference       `bson:"item,omitempty" json:"item,omitempty"`
}

type ListPlus struct {
	List                     `bson:",inline"`
	ListPlusRelatedResources `bson:",inline"`
}

type ListPlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                   *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                  *[]Device                `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                *[]Location              `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedBySource             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySource,omitempty"`
	IncludedDeviceResourcesReferencedBySource                   *[]Device                `bson:"_includedDeviceResourcesReferencedBySource,omitempty"`
	IncludedPatientResourcesReferencedBySource                  *[]Patient               `bson:"_includedPatientResourcesReferencedBySource,omitempty"`
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

func (l *ListPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if l.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*l.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*l.IncludedGroupResourcesReferencedBySubject))
	} else if len(*l.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*l.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if l.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*l.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*l.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*l.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*l.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if l.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*l.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*l.IncludedPatientResourcesReferencedBySubject))
	} else if len(*l.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*l.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if l.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*l.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*l.IncludedLocationResourcesReferencedBySubject))
	} else if len(*l.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*l.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if l.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*l.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*l.IncludedPatientResourcesReferencedByPatient))
	} else if len(*l.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*l.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySource() (practitioner *Practitioner, err error) {
	if l.IncludedPractitionerResourcesReferencedBySource == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*l.IncludedPractitionerResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*l.IncludedPractitionerResourcesReferencedBySource))
	} else if len(*l.IncludedPractitionerResourcesReferencedBySource) == 1 {
		practitioner = &(*l.IncludedPractitionerResourcesReferencedBySource)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedDeviceResourceReferencedBySource() (device *Device, err error) {
	if l.IncludedDeviceResourcesReferencedBySource == nil {
		err = errors.New("Included devices not requested")
	} else if len(*l.IncludedDeviceResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*l.IncludedDeviceResourcesReferencedBySource))
	} else if len(*l.IncludedDeviceResourcesReferencedBySource) == 1 {
		device = &(*l.IncludedDeviceResourcesReferencedBySource)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedPatientResourceReferencedBySource() (patient *Patient, err error) {
	if l.IncludedPatientResourcesReferencedBySource == nil {
		err = errors.New("Included patients not requested")
	} else if len(*l.IncludedPatientResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*l.IncludedPatientResourcesReferencedBySource))
	} else if len(*l.IncludedPatientResourcesReferencedBySource) == 1 {
		patient = &(*l.IncludedPatientResourcesReferencedBySource)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if l.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*l.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*l.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*l.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*l.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if l.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *l.RevIncludedListResourcesReferencingItem
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if l.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *l.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if l.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *l.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if l.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *l.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *l.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if l.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *l.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if l.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *l.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if l.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *l.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if l.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *l.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *l.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *l.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *l.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedLocationResourcesReferencedBySubject != nil {
		for _, r := range *l.IncludedLocationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *l.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedPractitionerResourcesReferencedBySource != nil {
		for _, r := range *l.IncludedPractitionerResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedDeviceResourcesReferencedBySource != nil {
		for _, r := range *l.IncludedDeviceResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedPatientResourcesReferencedBySource != nil {
		for _, r := range *l.IncludedPatientResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *l.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (l *ListPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *l.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *l.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *l.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *l.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *l.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *l.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *l.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *l.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *l.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (l *ListPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *l.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *l.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *l.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedLocationResourcesReferencedBySubject != nil {
		for _, r := range *l.IncludedLocationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *l.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedPractitionerResourcesReferencedBySource != nil {
		for _, r := range *l.IncludedPractitionerResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedDeviceResourcesReferencedBySource != nil {
		for _, r := range *l.IncludedDeviceResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedPatientResourcesReferencedBySource != nil {
		for _, r := range *l.IncludedPatientResourcesReferencedBySource {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *l.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *l.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *l.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *l.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *l.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *l.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *l.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *l.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *l.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *l.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
