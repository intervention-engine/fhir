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

type Specimen struct {
	DomainResource      `bson:",inline"`
	Identifier          []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              string                       `bson:"status,omitempty" json:"status,omitempty"`
	Type                *CodeableConcept             `bson:"type,omitempty" json:"type,omitempty"`
	Parent              []Reference                  `bson:"parent,omitempty" json:"parent,omitempty"`
	Subject             *Reference                   `bson:"subject,omitempty" json:"subject,omitempty"`
	AccessionIdentifier *Identifier                  `bson:"accessionIdentifier,omitempty" json:"accessionIdentifier,omitempty"`
	ReceivedTime        *FHIRDateTime                `bson:"receivedTime,omitempty" json:"receivedTime,omitempty"`
	Collection          *SpecimenCollectionComponent `bson:"collection,omitempty" json:"collection,omitempty"`
	Treatment           []SpecimenTreatmentComponent `bson:"treatment,omitempty" json:"treatment,omitempty"`
	Container           []SpecimenContainerComponent `bson:"container,omitempty" json:"container,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Specimen) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Specimen"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Specimen), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Specimen) GetBSON() (interface{}, error) {
	x.ResourceType = "Specimen"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "specimen" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type specimen Specimen

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Specimen) UnmarshalJSON(data []byte) (err error) {
	x2 := specimen{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Specimen(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Specimen) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Specimen"
	} else if x.ResourceType != "Specimen" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Specimen, instead received %s", x.ResourceType))
	}
	return nil
}

type SpecimenCollectionComponent struct {
	Collector         *Reference       `bson:"collector,omitempty" json:"collector,omitempty"`
	Comment           []string         `bson:"comment,omitempty" json:"comment,omitempty"`
	CollectedDateTime *FHIRDateTime    `bson:"collectedDateTime,omitempty" json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period          `bson:"collectedPeriod,omitempty" json:"collectedPeriod,omitempty"`
	Quantity          *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Method            *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	BodySite          *CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
}

type SpecimenTreatmentComponent struct {
	Description string           `bson:"description,omitempty" json:"description,omitempty"`
	Procedure   *CodeableConcept `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Additive    []Reference      `bson:"additive,omitempty" json:"additive,omitempty"`
}

type SpecimenContainerComponent struct {
	Identifier              []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Description             string           `bson:"description,omitempty" json:"description,omitempty"`
	Type                    *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Capacity                *Quantity        `bson:"capacity,omitempty" json:"capacity,omitempty"`
	SpecimenQuantity        *Quantity        `bson:"specimenQuantity,omitempty" json:"specimenQuantity,omitempty"`
	AdditiveCodeableConcept *CodeableConcept `bson:"additiveCodeableConcept,omitempty" json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *Reference       `bson:"additiveReference,omitempty" json:"additiveReference,omitempty"`
}

type SpecimenPlus struct {
	Specimen                     `bson:",inline"`
	SpecimenPlusRelatedResources `bson:",inline"`
}

type SpecimenPlusRelatedResources struct {
	IncludedSpecimenResourcesReferencedByParent                 *[]Specimen              `bson:"_includedSpecimenResourcesReferencedByParent,omitempty"`
	IncludedGroupResourcesReferencedBySubject                   *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                  *[]Device                `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedSubstanceResourcesReferencedBySubject               *[]Substance             `bson:"_includedSubstanceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByCollector          *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByCollector,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedSpecimenResourcesReferencingParent               *[]Specimen              `bson:"_revIncludedSpecimenResourcesReferencingParent,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                 *[]Media                 `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingSpecimen          *[]Observation           `bson:"_revIncludedObservationResourcesReferencingSpecimen,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSpecimen     *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingSpecimen,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath1 *[]DiagnosticOrder       `bson:"_revIncludedDiagnosticOrderResourcesReferencingSpecimenPath1,omitempty"`
	RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath2 *[]DiagnosticOrder       `bson:"_revIncludedDiagnosticOrderResourcesReferencingSpecimenPath2,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (s *SpecimenPlusRelatedResources) GetIncludedSpecimenResourcesReferencedByParent() (specimen []Specimen, err error) {
	if s.IncludedSpecimenResourcesReferencedByParent == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *s.IncludedSpecimenResourcesReferencedByParent
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if s.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*s.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*s.IncludedGroupResourcesReferencedBySubject))
	} else if len(*s.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*s.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if s.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*s.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*s.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*s.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*s.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if s.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResourcesReferencedBySubject))
	} else if len(*s.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*s.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedSubstanceResourceReferencedBySubject() (substance *Substance, err error) {
	if s.IncludedSubstanceResourcesReferencedBySubject == nil {
		err = errors.New("Included substances not requested")
	} else if len(*s.IncludedSubstanceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*s.IncludedSubstanceResourcesReferencedBySubject))
	} else if len(*s.IncludedSubstanceResourcesReferencedBySubject) == 1 {
		substance = &(*s.IncludedSubstanceResourcesReferencedBySubject)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if s.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResourcesReferencedByPatient))
	} else if len(*s.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*s.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedPractitionerResourceReferencedByCollector() (practitioner *Practitioner, err error) {
	if s.IncludedPractitionerResourcesReferencedByCollector == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*s.IncludedPractitionerResourcesReferencedByCollector) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*s.IncludedPractitionerResourcesReferencedByCollector))
	} else if len(*s.IncludedPractitionerResourcesReferencedByCollector) == 1 {
		practitioner = &(*s.IncludedPractitionerResourcesReferencedByCollector)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingParent() (specimen []Specimen, err error) {
	if s.RevIncludedSpecimenResourcesReferencingParent == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *s.RevIncludedSpecimenResourcesReferencingParent
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if s.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *s.RevIncludedListResourcesReferencingItem
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if s.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *s.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if s.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *s.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedObservationResourcesReferencingSpecimen() (observations []Observation, err error) {
	if s.RevIncludedObservationResourcesReferencingSpecimen == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *s.RevIncludedObservationResourcesReferencingSpecimen
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if s.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *s.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingSpecimen() (diagnosticReports []DiagnosticReport, err error) {
	if s.RevIncludedDiagnosticReportResourcesReferencingSpecimen == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *s.RevIncludedDiagnosticReportResourcesReferencingSpecimen
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if s.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *s.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *s.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingSpecimenPath1() (diagnosticOrders []DiagnosticOrder, err error) {
	if s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath1 == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath1
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDiagnosticOrderResourcesReferencingSpecimenPath2() (diagnosticOrders []DiagnosticOrder, err error) {
	if s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath2 == nil {
		err = errors.New("RevIncluded diagnosticOrders not requested")
	} else {
		diagnosticOrders = *s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath2
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *s.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if s.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *s.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *s.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if s.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *s.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedSpecimenResourcesReferencedByParent != nil {
		for _, r := range *s.IncludedSpecimenResourcesReferencedByParent {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *s.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *s.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *s.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedSubstanceResourcesReferencedBySubject != nil {
		for _, r := range *s.IncludedSubstanceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *s.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedPractitionerResourcesReferencedByCollector != nil {
		for _, r := range *s.IncludedPractitionerResourcesReferencedByCollector {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedSpecimenResourcesReferencingParent != nil {
		for _, r := range *s.RevIncludedSpecimenResourcesReferencingParent {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *s.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *s.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedMediaResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedMediaResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedObservationResourcesReferencingSpecimen != nil {
		for _, r := range *s.RevIncludedObservationResourcesReferencingSpecimen {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDiagnosticReportResourcesReferencingSpecimen != nil {
		for _, r := range *s.RevIncludedDiagnosticReportResourcesReferencingSpecimen {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *s.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath1 != nil {
		for _, r := range *s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath2 != nil {
		for _, r := range *s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *s.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *s.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *s.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *s.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (s *SpecimenPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedSpecimenResourcesReferencedByParent != nil {
		for _, r := range *s.IncludedSpecimenResourcesReferencedByParent {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *s.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *s.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *s.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedSubstanceResourcesReferencedBySubject != nil {
		for _, r := range *s.IncludedSubstanceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *s.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if s.IncludedPractitionerResourcesReferencedByCollector != nil {
		for _, r := range *s.IncludedPractitionerResourcesReferencedByCollector {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedSpecimenResourcesReferencingParent != nil {
		for _, r := range *s.RevIncludedSpecimenResourcesReferencingParent {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *s.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *s.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedMediaResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedMediaResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedObservationResourcesReferencingSpecimen != nil {
		for _, r := range *s.RevIncludedObservationResourcesReferencingSpecimen {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDiagnosticReportResourcesReferencingSpecimen != nil {
		for _, r := range *s.RevIncludedDiagnosticReportResourcesReferencingSpecimen {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *s.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath1 != nil {
		for _, r := range *s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath2 != nil {
		for _, r := range *s.RevIncludedDiagnosticOrderResourcesReferencingSpecimenPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *s.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *s.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *s.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *s.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
