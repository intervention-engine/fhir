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

type DiagnosticReport struct {
	DomainResource    `bson:",inline"`
	Identifier        []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string                           `bson:"status,omitempty" json:"status,omitempty"`
	Category          *CodeableConcept                 `bson:"category,omitempty" json:"category,omitempty"`
	Code              *CodeableConcept                 `bson:"code,omitempty" json:"code,omitempty"`
	Subject           *Reference                       `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter         *Reference                       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	EffectiveDateTime *FHIRDateTime                    `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period                          `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued            *FHIRDateTime                    `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer         *Reference                       `bson:"performer,omitempty" json:"performer,omitempty"`
	Request           []Reference                      `bson:"request,omitempty" json:"request,omitempty"`
	Specimen          []Reference                      `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Result            []Reference                      `bson:"result,omitempty" json:"result,omitempty"`
	ImagingStudy      []Reference                      `bson:"imagingStudy,omitempty" json:"imagingStudy,omitempty"`
	Image             []DiagnosticReportImageComponent `bson:"image,omitempty" json:"image,omitempty"`
	Conclusion        string                           `bson:"conclusion,omitempty" json:"conclusion,omitempty"`
	CodedDiagnosis    []CodeableConcept                `bson:"codedDiagnosis,omitempty" json:"codedDiagnosis,omitempty"`
	PresentedForm     []Attachment                     `bson:"presentedForm,omitempty" json:"presentedForm,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DiagnosticReport) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DiagnosticReport"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DiagnosticReport), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DiagnosticReport) GetBSON() (interface{}, error) {
	x.ResourceType = "DiagnosticReport"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "diagnosticReport" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type diagnosticReport DiagnosticReport

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DiagnosticReport) UnmarshalJSON(data []byte) (err error) {
	x2 := diagnosticReport{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DiagnosticReport(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DiagnosticReport) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DiagnosticReport"
	} else if x.ResourceType != "DiagnosticReport" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DiagnosticReport, instead received %s", x.ResourceType))
	}
	return nil
}

type DiagnosticReportImageComponent struct {
	BackboneElement `bson:",inline"`
	Comment         string     `bson:"comment,omitempty" json:"comment,omitempty"`
	Link            *Reference `bson:"link,omitempty" json:"link,omitempty"`
}

type DiagnosticReportPlus struct {
	DiagnosticReport                     `bson:",inline"`
	DiagnosticReportPlusRelatedResources `bson:",inline"`
}

type DiagnosticReportPlusRelatedResources struct {
	IncludedMediaResourcesReferencedByImage                        *[]Media                 `bson:"_includedMediaResourcesReferencedByImage,omitempty"`
	IncludedReferralRequestResourcesReferencedByRequest            *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByRequest,omitempty"`
	IncludedDiagnosticOrderResourcesReferencedByRequest            *[]DiagnosticOrder       `bson:"_includedDiagnosticOrderResourcesReferencedByRequest,omitempty"`
	IncludedProcedureRequestResourcesReferencedByRequest           *[]ProcedureRequest      `bson:"_includedProcedureRequestResourcesReferencedByRequest,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedGroupResourcesReferencedBySubject                      *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                     *[]Device                `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                    *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                   *[]Location              `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedObservationResourcesReferencedByResult                 *[]Observation           `bson:"_includedObservationResourcesReferencedByResult,omitempty"`
	IncludedPatientResourcesReferencedByPatient                    *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedSpecimenResourcesReferencedBySpecimen                  *[]Specimen              `bson:"_includedSpecimenResourcesReferencedBySpecimen,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                        *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref     *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                     *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                    *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference             *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated         *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment        *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject    *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest          *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingInvestigation *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingInvestigation,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedMediaResourceReferencedByImage() (media *Media, err error) {
	if d.IncludedMediaResourcesReferencedByImage == nil {
		err = errors.New("Included media not requested")
	} else if len(*d.IncludedMediaResourcesReferencedByImage) > 1 {
		err = fmt.Errorf("Expected 0 or 1 media, but found %d", len(*d.IncludedMediaResourcesReferencedByImage))
	} else if len(*d.IncludedMediaResourcesReferencedByImage) == 1 {
		media = &(*d.IncludedMediaResourcesReferencedByImage)[0]
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedReferralRequestResourcesReferencedByRequest() (referralRequests []ReferralRequest, err error) {
	if d.IncludedReferralRequestResourcesReferencedByRequest == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *d.IncludedReferralRequestResourcesReferencedByRequest
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedDiagnosticOrderResourcesReferencedByRequest() (diagnosticOrders []DiagnosticOrder, err error) {
	if d.IncludedDiagnosticOrderResourcesReferencedByRequest == nil {
		err = errors.New("Included diagnosticOrders not requested")
	} else {
		diagnosticOrders = *d.IncludedDiagnosticOrderResourcesReferencedByRequest
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedProcedureRequestResourcesReferencedByRequest() (procedureRequests []ProcedureRequest, err error) {
	if d.IncludedProcedureRequestResourcesReferencedByRequest == nil {
		err = errors.New("Included procedureRequests not requested")
	} else {
		procedureRequests = *d.IncludedProcedureRequestResourcesReferencedByRequest
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPerformer() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedByPerformer))
	} else if len(*d.IncludedPractitionerResourcesReferencedByPerformer) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedByPerformer)[0]
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedOrganizationResourceReferencedByPerformer() (organization *Organization, err error) {
	if d.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*d.IncludedOrganizationResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*d.IncludedOrganizationResourcesReferencedByPerformer))
	} else if len(*d.IncludedOrganizationResourcesReferencedByPerformer) == 1 {
		organization = &(*d.IncludedOrganizationResourcesReferencedByPerformer)[0]
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if d.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedGroupResourcesReferencedBySubject))
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*d.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedBySubject))
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if d.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*d.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*d.IncludedLocationResourcesReferencedBySubject))
	} else if len(*d.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*d.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if d.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*d.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*d.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*d.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*d.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedObservationResourcesReferencedByResult() (observations []Observation, err error) {
	if d.IncludedObservationResourcesReferencedByResult == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *d.IncludedObservationResourcesReferencedByResult
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedByPatient))
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedSpecimenResourcesReferencedBySpecimen() (specimen []Specimen, err error) {
	if d.IncludedSpecimenResourcesReferencedBySpecimen == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *d.IncludedSpecimenResourcesReferencedBySpecimen
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if d.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *d.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *d.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if d.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *d.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingInvestigation() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingInvestigation == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingInvestigation
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (d *DiagnosticReportPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedMediaResourcesReferencedByImage != nil {
		for idx := range *d.IncludedMediaResourcesReferencedByImage {
			rsc := (*d.IncludedMediaResourcesReferencedByImage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedReferralRequestResourcesReferencedByRequest != nil {
		for idx := range *d.IncludedReferralRequestResourcesReferencedByRequest {
			rsc := (*d.IncludedReferralRequestResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDiagnosticOrderResourcesReferencedByRequest != nil {
		for idx := range *d.IncludedDiagnosticOrderResourcesReferencedByRequest {
			rsc := (*d.IncludedDiagnosticOrderResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedProcedureRequestResourcesReferencedByRequest != nil {
		for idx := range *d.IncludedProcedureRequestResourcesReferencedByRequest {
			rsc := (*d.IncludedProcedureRequestResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*d.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*d.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
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
	if d.IncludedObservationResourcesReferencedByResult != nil {
		for idx := range *d.IncludedObservationResourcesReferencedByResult {
			rsc := (*d.IncludedObservationResourcesReferencedByResult)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedSpecimenResourcesReferencedBySpecimen != nil {
		for idx := range *d.IncludedSpecimenResourcesReferencedBySpecimen {
			rsc := (*d.IncludedSpecimenResourcesReferencedBySpecimen)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DiagnosticReportPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	if d.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
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

func (d *DiagnosticReportPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedMediaResourcesReferencedByImage != nil {
		for idx := range *d.IncludedMediaResourcesReferencedByImage {
			rsc := (*d.IncludedMediaResourcesReferencedByImage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedReferralRequestResourcesReferencedByRequest != nil {
		for idx := range *d.IncludedReferralRequestResourcesReferencedByRequest {
			rsc := (*d.IncludedReferralRequestResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDiagnosticOrderResourcesReferencedByRequest != nil {
		for idx := range *d.IncludedDiagnosticOrderResourcesReferencedByRequest {
			rsc := (*d.IncludedDiagnosticOrderResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedProcedureRequestResourcesReferencedByRequest != nil {
		for idx := range *d.IncludedProcedureRequestResourcesReferencedByRequest {
			rsc := (*d.IncludedProcedureRequestResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *d.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*d.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*d.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
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
	if d.IncludedObservationResourcesReferencedByResult != nil {
		for idx := range *d.IncludedObservationResourcesReferencedByResult {
			rsc := (*d.IncludedObservationResourcesReferencedByResult)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedSpecimenResourcesReferencedBySpecimen != nil {
		for idx := range *d.IncludedSpecimenResourcesReferencedBySpecimen {
			rsc := (*d.IncludedSpecimenResourcesReferencedBySpecimen)[idx]
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
	if d.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
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
