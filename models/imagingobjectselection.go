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

type ImagingObjectSelection struct {
	DomainResource `bson:",inline"`
	Uid            string                                 `bson:"uid,omitempty" json:"uid,omitempty"`
	Patient        *Reference                             `bson:"patient,omitempty" json:"patient,omitempty"`
	Title          *CodeableConcept                       `bson:"title,omitempty" json:"title,omitempty"`
	Description    string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Author         *Reference                             `bson:"author,omitempty" json:"author,omitempty"`
	AuthoringTime  *FHIRDateTime                          `bson:"authoringTime,omitempty" json:"authoringTime,omitempty"`
	Study          []ImagingObjectSelectionStudyComponent `bson:"study,omitempty" json:"study,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ImagingObjectSelection) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ImagingObjectSelection"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ImagingObjectSelection), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ImagingObjectSelection) GetBSON() (interface{}, error) {
	x.ResourceType = "ImagingObjectSelection"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "imagingObjectSelection" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type imagingObjectSelection ImagingObjectSelection

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ImagingObjectSelection) UnmarshalJSON(data []byte) (err error) {
	x2 := imagingObjectSelection{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ImagingObjectSelection(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ImagingObjectSelection) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ImagingObjectSelection"
	} else if x.ResourceType != "ImagingObjectSelection" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ImagingObjectSelection, instead received %s", x.ResourceType))
	}
	return nil
}

type ImagingObjectSelectionStudyComponent struct {
	Uid          string                                  `bson:"uid,omitempty" json:"uid,omitempty"`
	Url          string                                  `bson:"url,omitempty" json:"url,omitempty"`
	ImagingStudy *Reference                              `bson:"imagingStudy,omitempty" json:"imagingStudy,omitempty"`
	Series       []ImagingObjectSelectionSeriesComponent `bson:"series,omitempty" json:"series,omitempty"`
}

type ImagingObjectSelectionSeriesComponent struct {
	Uid      string                                    `bson:"uid,omitempty" json:"uid,omitempty"`
	Url      string                                    `bson:"url,omitempty" json:"url,omitempty"`
	Instance []ImagingObjectSelectionInstanceComponent `bson:"instance,omitempty" json:"instance,omitempty"`
}

type ImagingObjectSelectionInstanceComponent struct {
	SopClass string                                  `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	Uid      string                                  `bson:"uid,omitempty" json:"uid,omitempty"`
	Url      string                                  `bson:"url,omitempty" json:"url,omitempty"`
	Frames   []ImagingObjectSelectionFramesComponent `bson:"frames,omitempty" json:"frames,omitempty"`
}

type ImagingObjectSelectionFramesComponent struct {
	FrameNumbers []uint32 `bson:"frameNumbers,omitempty" json:"frameNumbers,omitempty"`
	Url          string   `bson:"url,omitempty" json:"url,omitempty"`
}

type ImagingObjectSelectionPlus struct {
	ImagingObjectSelection                     `bson:",inline"`
	ImagingObjectSelectionPlusRelatedResources `bson:",inline"`
}

type ImagingObjectSelectionPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByAuthor             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthor             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByAuthor                   *[]Device                `bson:"_includedDeviceResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByAuthor                  *[]Patient               `bson:"_includedPatientResourcesReferencedByAuthor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAuthor            *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByAuthor,omitempty"`
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

func (i *ImagingObjectSelectionPlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if i.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*i.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*i.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*i.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*i.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetIncludedOrganizationResourceReferencedByAuthor() (organization *Organization, err error) {
	if i.IncludedOrganizationResourcesReferencedByAuthor == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*i.IncludedOrganizationResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*i.IncludedOrganizationResourcesReferencedByAuthor))
	} else if len(*i.IncludedOrganizationResourcesReferencedByAuthor) == 1 {
		organization = &(*i.IncludedOrganizationResourcesReferencedByAuthor)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetIncludedDeviceResourceReferencedByAuthor() (device *Device, err error) {
	if i.IncludedDeviceResourcesReferencedByAuthor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*i.IncludedDeviceResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*i.IncludedDeviceResourcesReferencedByAuthor))
	} else if len(*i.IncludedDeviceResourcesReferencedByAuthor) == 1 {
		device = &(*i.IncludedDeviceResourcesReferencedByAuthor)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetIncludedPatientResourceReferencedByAuthor() (patient *Patient, err error) {
	if i.IncludedPatientResourcesReferencedByAuthor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResourcesReferencedByAuthor))
	} else if len(*i.IncludedPatientResourcesReferencedByAuthor) == 1 {
		patient = &(*i.IncludedPatientResourcesReferencedByAuthor)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAuthor() (relatedPerson *RelatedPerson, err error) {
	if i.IncludedRelatedPersonResourcesReferencedByAuthor == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*i.IncludedRelatedPersonResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*i.IncludedRelatedPersonResourcesReferencedByAuthor))
	} else if len(*i.IncludedRelatedPersonResourcesReferencedByAuthor) == 1 {
		relatedPerson = &(*i.IncludedRelatedPersonResourcesReferencedByAuthor)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if i.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResourcesReferencedByPatient))
	} else if len(*i.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*i.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if i.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *i.RevIncludedListResourcesReferencingItem
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if i.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *i.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if i.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *i.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if i.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *i.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *i.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if i.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *i.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if i.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *i.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if i.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *i.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if i.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *i.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for _, r := range *i.IncludedPractitionerResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for _, r := range *i.IncludedOrganizationResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedDeviceResourcesReferencedByAuthor != nil {
		for _, r := range *i.IncludedDeviceResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedPatientResourcesReferencedByAuthor != nil {
		for _, r := range *i.IncludedPatientResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for _, r := range *i.IncludedRelatedPersonResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *i.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *i.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *i.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *i.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *i.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *i.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *i.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *i.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *i.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *i.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (i *ImagingObjectSelectionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for _, r := range *i.IncludedPractitionerResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for _, r := range *i.IncludedOrganizationResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedDeviceResourcesReferencedByAuthor != nil {
		for _, r := range *i.IncludedDeviceResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedPatientResourcesReferencedByAuthor != nil {
		for _, r := range *i.IncludedPatientResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for _, r := range *i.IncludedRelatedPersonResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *i.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *i.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *i.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *i.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *i.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *i.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *i.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *i.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *i.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *i.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
