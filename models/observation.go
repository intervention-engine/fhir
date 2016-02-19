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

type Observation struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                               `bson:"status,omitempty" json:"status,omitempty"`
	Category             *CodeableConcept                     `bson:"category,omitempty" json:"category,omitempty"`
	Code                 *CodeableConcept                     `bson:"code,omitempty" json:"code,omitempty"`
	Subject              *Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter            *Reference                           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	EffectiveDateTime    *FHIRDateTime                        `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod      *Period                              `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued               *FHIRDateTime                        `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer            []Reference                          `bson:"performer,omitempty" json:"performer,omitempty"`
	ValueQuantity        *Quantity                            `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept                     `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          string                               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueRange           *Range                               `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                               `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                         `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueAttachment      *Attachment                          `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueTime            *FHIRDateTime                        `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime        *FHIRDateTime                        `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                              `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept                     `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation       *CodeableConcept                     `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	Comments             string                               `bson:"comments,omitempty" json:"comments,omitempty"`
	BodySite             *CodeableConcept                     `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Method               *CodeableConcept                     `bson:"method,omitempty" json:"method,omitempty"`
	Specimen             *Reference                           `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Device               *Reference                           `bson:"device,omitempty" json:"device,omitempty"`
	ReferenceRange       []ObservationReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
	Related              []ObservationRelatedComponent        `bson:"related,omitempty" json:"related,omitempty"`
	Component            []ObservationComponentComponent      `bson:"component,omitempty" json:"component,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Observation) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Observation"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Observation), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Observation) GetBSON() (interface{}, error) {
	x.ResourceType = "Observation"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "observation" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type observation Observation

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Observation) UnmarshalJSON(data []byte) (err error) {
	x2 := observation{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Observation(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Observation) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Observation"
	} else if x.ResourceType != "Observation" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Observation, instead received %s", x.ResourceType))
	}
	return nil
}

type ObservationReferenceRangeComponent struct {
	BackboneElement `bson:",inline"`
	Low             *Quantity        `bson:"low,omitempty" json:"low,omitempty"`
	High            *Quantity        `bson:"high,omitempty" json:"high,omitempty"`
	Meaning         *CodeableConcept `bson:"meaning,omitempty" json:"meaning,omitempty"`
	Age             *Range           `bson:"age,omitempty" json:"age,omitempty"`
	Text            string           `bson:"text,omitempty" json:"text,omitempty"`
}

type ObservationRelatedComponent struct {
	BackboneElement `bson:",inline"`
	Type            string     `bson:"type,omitempty" json:"type,omitempty"`
	Target          *Reference `bson:"target,omitempty" json:"target,omitempty"`
}

type ObservationComponentComponent struct {
	BackboneElement      `bson:",inline"`
	Code                 *CodeableConcept                     `bson:"code,omitempty" json:"code,omitempty"`
	ValueQuantity        *Quantity                            `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept                     `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          string                               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueRange           *Range                               `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                               `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                         `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueAttachment      *Attachment                          `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueTime            *FHIRDateTime                        `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime        *FHIRDateTime                        `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                              `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept                     `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	ReferenceRange       []ObservationReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}

type ObservationPlus struct {
	Observation                     `bson:",inline"`
	ObservationPlusRelatedResources `bson:",inline"`
}

type ObservationPlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                            *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                           *[]Device                     `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                          *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                         *[]Location                   `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedSpecimenResourcesReferencedBySpecimen                        *[]Specimen                   `bson:"_includedSpecimenResourcesReferencedBySpecimen,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer                   *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer                   *[]Organization               `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                        *[]Patient                    `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                  *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                      *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedObservationResourcesReferencedByRelatedtarget                *[]Observation                `bson:"_includedObservationResourcesReferencedByRelatedtarget,omitempty"`
	IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget      *[]QuestionnaireResponse      `bson:"_includedQuestionnaireResponseResourcesReferencedByRelatedtarget,omitempty"`
	IncludedDeviceResourcesReferencedByDevice                            *[]Device                     `bson:"_includedDeviceResourcesReferencedByDevice,omitempty"`
	IncludedDeviceMetricResourcesReferencedByDevice                      *[]DeviceMetric               `bson:"_includedDeviceMetricResourcesReferencedByDevice,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                      *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref            *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref            *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                              *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref           *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                           *[]Order                      `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedImmunizationResourcesReferencingReaction                  *[]Immunization               `bson:"_revIncludedImmunizationResourcesReferencingReaction,omitempty"`
	RevIncludedObservationResourcesReferencingRelatedtarget              *[]Observation                `bson:"_revIncludedObservationResourcesReferencingRelatedtarget,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                          *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingResult                *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingResult,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference                   *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                    *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                      *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated               *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment              *[]OrderResponse              `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject          *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest                *[]ProcessResponse            `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger             *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingInvestigation       *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingInvestigation,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                     *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
}

func (o *ObservationPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if o.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*o.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*o.IncludedGroupResourcesReferencedBySubject))
	} else if len(*o.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*o.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if o.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*o.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*o.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if o.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedPatientResourcesReferencedBySubject))
	} else if len(*o.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*o.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if o.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*o.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*o.IncludedLocationResourcesReferencedBySubject))
	} else if len(*o.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*o.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if o.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedPatientResourcesReferencedByPatient))
	} else if len(*o.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*o.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedSpecimenResourceReferencedBySpecimen() (specimen *Specimen, err error) {
	if o.IncludedSpecimenResourcesReferencedBySpecimen == nil {
		err = errors.New("Included specimen not requested")
	} else if len(*o.IncludedSpecimenResourcesReferencedBySpecimen) > 1 {
		err = fmt.Errorf("Expected 0 or 1 specimen, but found %d", len(*o.IncludedSpecimenResourcesReferencedBySpecimen))
	} else if len(*o.IncludedSpecimenResourcesReferencedBySpecimen) == 1 {
		specimen = &(*o.IncludedSpecimenResourcesReferencedBySpecimen)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByPerformer() (practitioners []Practitioner, err error) {
	if o.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *o.IncludedPractitionerResourcesReferencedByPerformer
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByPerformer() (organizations []Organization, err error) {
	if o.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *o.IncludedOrganizationResourcesReferencedByPerformer
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedPatientResourcesReferencedByPerformer() (patients []Patient, err error) {
	if o.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *o.IncludedPatientResourcesReferencedByPerformer
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByPerformer() (relatedPeople []RelatedPerson, err error) {
	if o.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *o.IncludedRelatedPersonResourcesReferencedByPerformer
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if o.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*o.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*o.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*o.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*o.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedObservationResourceReferencedByRelatedtarget() (observation *Observation, err error) {
	if o.IncludedObservationResourcesReferencedByRelatedtarget == nil {
		err = errors.New("Included observations not requested")
	} else if len(*o.IncludedObservationResourcesReferencedByRelatedtarget) > 1 {
		err = fmt.Errorf("Expected 0 or 1 observation, but found %d", len(*o.IncludedObservationResourcesReferencedByRelatedtarget))
	} else if len(*o.IncludedObservationResourcesReferencedByRelatedtarget) == 1 {
		observation = &(*o.IncludedObservationResourcesReferencedByRelatedtarget)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedQuestionnaireResponseResourceReferencedByRelatedtarget() (questionnaireResponse *QuestionnaireResponse, err error) {
	if o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget == nil {
		err = errors.New("Included questionnaireresponses not requested")
	} else if len(*o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget) > 1 {
		err = fmt.Errorf("Expected 0 or 1 questionnaireResponse, but found %d", len(*o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget))
	} else if len(*o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget) == 1 {
		questionnaireResponse = &(*o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedDeviceResourceReferencedByDevice() (device *Device, err error) {
	if o.IncludedDeviceResourcesReferencedByDevice == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedDeviceResourcesReferencedByDevice) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedDeviceResourcesReferencedByDevice))
	} else if len(*o.IncludedDeviceResourcesReferencedByDevice) == 1 {
		device = &(*o.IncludedDeviceResourcesReferencedByDevice)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedDeviceMetricResourceReferencedByDevice() (deviceMetric *DeviceMetric, err error) {
	if o.IncludedDeviceMetricResourcesReferencedByDevice == nil {
		err = errors.New("Included devicemetrics not requested")
	} else if len(*o.IncludedDeviceMetricResourcesReferencedByDevice) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceMetric, but found %d", len(*o.IncludedDeviceMetricResourcesReferencedByDevice))
	} else if len(*o.IncludedDeviceMetricResourcesReferencedByDevice) == 1 {
		deviceMetric = &(*o.IncludedDeviceMetricResourcesReferencedByDevice)[0]
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if o.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *o.RevIncludedListResourcesReferencingItem
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if o.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *o.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingReaction() (immunizations []Immunization, err error) {
	if o.RevIncludedImmunizationResourcesReferencingReaction == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *o.RevIncludedImmunizationResourcesReferencingReaction
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingRelatedtarget() (observations []Observation, err error) {
	if o.RevIncludedObservationResourcesReferencingRelatedtarget == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *o.RevIncludedObservationResourcesReferencingRelatedtarget
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingResult() (diagnosticReports []DiagnosticReport, err error) {
	if o.RevIncludedDiagnosticReportResourcesReferencingResult == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *o.RevIncludedDiagnosticReportResourcesReferencingResult
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *o.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *o.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingInvestigation() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingInvestigation == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingInvestigation
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (o *ObservationPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (o *ObservationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *o.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *o.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *o.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedLocationResourcesReferencedBySubject != nil {
		for _, r := range *o.IncludedLocationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *o.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedSpecimenResourcesReferencedBySpecimen != nil {
		for _, r := range *o.IncludedSpecimenResourcesReferencedBySpecimen {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for _, r := range *o.IncludedPractitionerResourcesReferencedByPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for _, r := range *o.IncludedOrganizationResourcesReferencedByPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPatientResourcesReferencedByPerformer != nil {
		for _, r := range *o.IncludedPatientResourcesReferencedByPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for _, r := range *o.IncludedRelatedPersonResourcesReferencedByPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *o.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedObservationResourcesReferencedByRelatedtarget != nil {
		for _, r := range *o.IncludedObservationResourcesReferencedByRelatedtarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget != nil {
		for _, r := range *o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedDeviceResourcesReferencedByDevice != nil {
		for _, r := range *o.IncludedDeviceResourcesReferencedByDevice {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedDeviceMetricResourcesReferencedByDevice != nil {
		for _, r := range *o.IncludedDeviceMetricResourcesReferencedByDevice {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (o *ObservationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *o.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *o.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingReaction != nil {
		for _, r := range *o.RevIncludedImmunizationResourcesReferencingReaction {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedObservationResourcesReferencingRelatedtarget != nil {
		for _, r := range *o.RevIncludedObservationResourcesReferencingRelatedtarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingResult != nil {
		for _, r := range *o.RevIncludedDiagnosticReportResourcesReferencingResult {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *o.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *o.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *o.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for _, r := range *o.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for _, r := range *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (o *ObservationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *o.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *o.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *o.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedLocationResourcesReferencedBySubject != nil {
		for _, r := range *o.IncludedLocationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *o.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedSpecimenResourcesReferencedBySpecimen != nil {
		for _, r := range *o.IncludedSpecimenResourcesReferencedBySpecimen {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for _, r := range *o.IncludedPractitionerResourcesReferencedByPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for _, r := range *o.IncludedOrganizationResourcesReferencedByPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedPatientResourcesReferencedByPerformer != nil {
		for _, r := range *o.IncludedPatientResourcesReferencedByPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for _, r := range *o.IncludedRelatedPersonResourcesReferencedByPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *o.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedObservationResourcesReferencedByRelatedtarget != nil {
		for _, r := range *o.IncludedObservationResourcesReferencedByRelatedtarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget != nil {
		for _, r := range *o.IncludedQuestionnaireResponseResourcesReferencedByRelatedtarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedDeviceResourcesReferencedByDevice != nil {
		for _, r := range *o.IncludedDeviceResourcesReferencedByDevice {
			resourceMap[r.Id] = &r
		}
	}
	if o.IncludedDeviceMetricResourcesReferencedByDevice != nil {
		for _, r := range *o.IncludedDeviceMetricResourcesReferencedByDevice {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *o.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *o.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingReaction != nil {
		for _, r := range *o.RevIncludedImmunizationResourcesReferencingReaction {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedObservationResourcesReferencingRelatedtarget != nil {
		for _, r := range *o.RevIncludedObservationResourcesReferencingRelatedtarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingResult != nil {
		for _, r := range *o.RevIncludedDiagnosticReportResourcesReferencingResult {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *o.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *o.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *o.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for _, r := range *o.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for _, r := range *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
