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

type DiagnosticOrder struct {
	DomainResource        `bson:",inline"`
	Subject               *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Orderer               *Reference                      `bson:"orderer,omitempty" json:"orderer,omitempty"`
	Identifier            []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Encounter             *Reference                      `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Reason                []CodeableConcept               `bson:"reason,omitempty" json:"reason,omitempty"`
	SupportingInformation []Reference                     `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Specimen              []Reference                     `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Status                string                          `bson:"status,omitempty" json:"status,omitempty"`
	Priority              string                          `bson:"priority,omitempty" json:"priority,omitempty"`
	Event                 []DiagnosticOrderEventComponent `bson:"event,omitempty" json:"event,omitempty"`
	Item                  []DiagnosticOrderItemComponent  `bson:"item,omitempty" json:"item,omitempty"`
	Note                  []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DiagnosticOrder) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DiagnosticOrder"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DiagnosticOrder), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DiagnosticOrder) GetBSON() (interface{}, error) {
	x.ResourceType = "DiagnosticOrder"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "diagnosticOrder" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type diagnosticOrder DiagnosticOrder

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DiagnosticOrder) UnmarshalJSON(data []byte) (err error) {
	x2 := diagnosticOrder{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DiagnosticOrder(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DiagnosticOrder) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DiagnosticOrder"
	} else if x.ResourceType != "DiagnosticOrder" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DiagnosticOrder, instead received %s", x.ResourceType))
	}
	return nil
}

type DiagnosticOrderEventComponent struct {
	BackboneElement `bson:",inline"`
	Status          string           `bson:"status,omitempty" json:"status,omitempty"`
	Description     *CodeableConcept `bson:"description,omitempty" json:"description,omitempty"`
	DateTime        *FHIRDateTime    `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Actor           *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
}

type DiagnosticOrderItemComponent struct {
	BackboneElement `bson:",inline"`
	Code            *CodeableConcept                `bson:"code,omitempty" json:"code,omitempty"`
	Specimen        []Reference                     `bson:"specimen,omitempty" json:"specimen,omitempty"`
	BodySite        *CodeableConcept                `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Status          string                          `bson:"status,omitempty" json:"status,omitempty"`
	Event           []DiagnosticOrderEventComponent `bson:"event,omitempty" json:"event,omitempty"`
}

type DiagnosticOrderPlus struct {
	DiagnosticOrder                     `bson:",inline"`
	DiagnosticOrderPlusRelatedResources `bson:",inline"`
}

type DiagnosticOrderPlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                   *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                  *[]Device                `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                *[]Location              `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter             *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPractitionerResourcesReferencedByActorPath1         *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByActorPath1,omitempty"`
	IncludedPractitionerResourcesReferencedByActorPath2         *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByActorPath2,omitempty"`
	IncludedDeviceResourcesReferencedByActorPath1               *[]Device                `bson:"_includedDeviceResourcesReferencedByActorPath1,omitempty"`
	IncludedDeviceResourcesReferencedByActorPath2               *[]Device                `bson:"_includedDeviceResourcesReferencedByActorPath2,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByOrderer            *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByOrderer,omitempty"`
	IncludedSpecimenResourcesReferencedBySpecimenPath1          *[]Specimen              `bson:"_includedSpecimenResourcesReferencedBySpecimenPath1,omitempty"`
	IncludedSpecimenResourcesReferencedBySpecimenPath2          *[]Specimen              `bson:"_includedSpecimenResourcesReferencedBySpecimenPath2,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference    *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingRequest      *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingRequest,omitempty"`
	RevIncludedImagingStudyResourcesReferencingOrder            *[]ImagingStudy          `bson:"_revIncludedImagingStudyResourcesReferencingOrder,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAction     *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPlan       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPlan,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if d.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedGroupResourcesReferencedBySubject))
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*d.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedBySubject))
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if d.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*d.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*d.IncludedLocationResourcesReferencedBySubject))
	} else if len(*d.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*d.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if d.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*d.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*d.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*d.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*d.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByActorPath1() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByActorPath1 == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedByActorPath1) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedByActorPath1))
	} else if len(*d.IncludedPractitionerResourcesReferencedByActorPath1) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedByActorPath1)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByActorPath2() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByActorPath2 == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedByActorPath2) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedByActorPath2))
	} else if len(*d.IncludedPractitionerResourcesReferencedByActorPath2) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedByActorPath2)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedDeviceResourceReferencedByActorPath1() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedByActorPath1 == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedByActorPath1) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedByActorPath1))
	} else if len(*d.IncludedDeviceResourcesReferencedByActorPath1) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedByActorPath1)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedDeviceResourceReferencedByActorPath2() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedByActorPath2 == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedByActorPath2) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedByActorPath2))
	} else if len(*d.IncludedDeviceResourcesReferencedByActorPath2) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedByActorPath2)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedByPatient))
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByOrderer() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByOrderer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedByOrderer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedByOrderer))
	} else if len(*d.IncludedPractitionerResourcesReferencedByOrderer) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedByOrderer)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedSpecimenResourcesReferencedBySpecimenPath1() (specimen []Specimen, err error) {
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath1 == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *d.IncludedSpecimenResourcesReferencedBySpecimenPath1
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedSpecimenResourcesReferencedBySpecimenPath2() (specimen []Specimen, err error) {
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath2 == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *d.IncludedSpecimenResourcesReferencedBySpecimenPath2
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if d.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *d.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if d.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *d.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingRequest() (diagnosticReports []DiagnosticReport, err error) {
	if d.RevIncludedDiagnosticReportResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *d.RevIncludedDiagnosticReportResourcesReferencingRequest
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingOrder() (imagingStudies []ImagingStudy, err error) {
	if d.RevIncludedImagingStudyResourcesReferencingOrder == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *d.RevIncludedImagingStudyResourcesReferencingOrder
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *d.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if d.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *d.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAction() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingAction == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingAction
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPlan() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingPlan == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingPlan
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedGroupResourcesReferencedBySubject {
			rsc := (*d.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*d.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPatientResourcesReferencedBySubject {
			rsc := (*d.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedLocationResourcesReferencedBySubject {
			rsc := (*d.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *d.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*d.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByActorPath1 != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByActorPath1 {
			rsc := (*d.IncludedPractitionerResourcesReferencedByActorPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByActorPath2 != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByActorPath2 {
			rsc := (*d.IncludedPractitionerResourcesReferencedByActorPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByActorPath1 != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByActorPath1 {
			rsc := (*d.IncludedDeviceResourcesReferencedByActorPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByActorPath2 != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByActorPath2 {
			rsc := (*d.IncludedDeviceResourcesReferencedByActorPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByOrderer != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByOrderer {
			rsc := (*d.IncludedPractitionerResourcesReferencedByOrderer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath1 != nil {
		for idx := range *d.IncludedSpecimenResourcesReferencedBySpecimenPath1 {
			rsc := (*d.IncludedSpecimenResourcesReferencedBySpecimenPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath2 != nil {
		for idx := range *d.IncludedSpecimenResourcesReferencedBySpecimenPath2 {
			rsc := (*d.IncludedSpecimenResourcesReferencedBySpecimenPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if d.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *d.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*d.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *d.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*d.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedDiagnosticReportResourcesReferencingRequest {
			rsc := (*d.RevIncludedDiagnosticReportResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImagingStudyResourcesReferencingOrder != nil {
		for idx := range *d.RevIncludedImagingStudyResourcesReferencingOrder {
			rsc := (*d.RevIncludedImagingStudyResourcesReferencingOrder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *d.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*d.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedGroupResourcesReferencedBySubject {
			rsc := (*d.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*d.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPatientResourcesReferencedBySubject {
			rsc := (*d.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedLocationResourcesReferencedBySubject {
			rsc := (*d.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *d.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*d.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByActorPath1 != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByActorPath1 {
			rsc := (*d.IncludedPractitionerResourcesReferencedByActorPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByActorPath2 != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByActorPath2 {
			rsc := (*d.IncludedPractitionerResourcesReferencedByActorPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByActorPath1 != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByActorPath1 {
			rsc := (*d.IncludedDeviceResourcesReferencedByActorPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByActorPath2 != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByActorPath2 {
			rsc := (*d.IncludedDeviceResourcesReferencedByActorPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByOrderer != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByOrderer {
			rsc := (*d.IncludedPractitionerResourcesReferencedByOrderer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath1 != nil {
		for idx := range *d.IncludedSpecimenResourcesReferencedBySpecimenPath1 {
			rsc := (*d.IncludedSpecimenResourcesReferencedBySpecimenPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath2 != nil {
		for idx := range *d.IncludedSpecimenResourcesReferencedBySpecimenPath2 {
			rsc := (*d.IncludedSpecimenResourcesReferencedBySpecimenPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
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
	if d.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *d.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*d.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *d.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*d.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedDiagnosticReportResourcesReferencingRequest {
			rsc := (*d.RevIncludedDiagnosticReportResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImagingStudyResourcesReferencingOrder != nil {
		for idx := range *d.RevIncludedImagingStudyResourcesReferencingOrder {
			rsc := (*d.RevIncludedImagingStudyResourcesReferencingOrder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingReference)[idx]
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
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *d.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*d.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
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
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
