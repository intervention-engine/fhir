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

type Immunization struct {
	DomainResource      `bson:",inline"`
	Identifier          []Identifier                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              string                                     `bson:"status,omitempty" json:"status,omitempty"`
	Date                *FHIRDateTime                              `bson:"date,omitempty" json:"date,omitempty"`
	VaccineCode         *CodeableConcept                           `bson:"vaccineCode,omitempty" json:"vaccineCode,omitempty"`
	Patient             *Reference                                 `bson:"patient,omitempty" json:"patient,omitempty"`
	WasNotGiven         *bool                                      `bson:"wasNotGiven,omitempty" json:"wasNotGiven,omitempty"`
	Reported            *bool                                      `bson:"reported,omitempty" json:"reported,omitempty"`
	Performer           *Reference                                 `bson:"performer,omitempty" json:"performer,omitempty"`
	Requester           *Reference                                 `bson:"requester,omitempty" json:"requester,omitempty"`
	Encounter           *Reference                                 `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Manufacturer        *Reference                                 `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Location            *Reference                                 `bson:"location,omitempty" json:"location,omitempty"`
	LotNumber           string                                     `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	ExpirationDate      *FHIRDateTime                              `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	Site                *CodeableConcept                           `bson:"site,omitempty" json:"site,omitempty"`
	Route               *CodeableConcept                           `bson:"route,omitempty" json:"route,omitempty"`
	DoseQuantity        *Quantity                                  `bson:"doseQuantity,omitempty" json:"doseQuantity,omitempty"`
	Note                []Annotation                               `bson:"note,omitempty" json:"note,omitempty"`
	Explanation         *ImmunizationExplanationComponent          `bson:"explanation,omitempty" json:"explanation,omitempty"`
	Reaction            []ImmunizationReactionComponent            `bson:"reaction,omitempty" json:"reaction,omitempty"`
	VaccinationProtocol []ImmunizationVaccinationProtocolComponent `bson:"vaccinationProtocol,omitempty" json:"vaccinationProtocol,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Immunization) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Immunization"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Immunization), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Immunization) GetBSON() (interface{}, error) {
	x.ResourceType = "Immunization"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "immunization" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type immunization Immunization

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Immunization) UnmarshalJSON(data []byte) (err error) {
	x2 := immunization{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Immunization(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Immunization) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Immunization"
	} else if x.ResourceType != "Immunization" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Immunization, instead received %s", x.ResourceType))
	}
	return nil
}

type ImmunizationExplanationComponent struct {
	Reason         []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	ReasonNotGiven []CodeableConcept `bson:"reasonNotGiven,omitempty" json:"reasonNotGiven,omitempty"`
}

type ImmunizationReactionComponent struct {
	Date     *FHIRDateTime `bson:"date,omitempty" json:"date,omitempty"`
	Detail   *Reference    `bson:"detail,omitempty" json:"detail,omitempty"`
	Reported *bool         `bson:"reported,omitempty" json:"reported,omitempty"`
}

type ImmunizationVaccinationProtocolComponent struct {
	DoseSequence     *uint32           `bson:"doseSequence,omitempty" json:"doseSequence,omitempty"`
	Description      string            `bson:"description,omitempty" json:"description,omitempty"`
	Authority        *Reference        `bson:"authority,omitempty" json:"authority,omitempty"`
	Series           string            `bson:"series,omitempty" json:"series,omitempty"`
	SeriesDoses      *uint32           `bson:"seriesDoses,omitempty" json:"seriesDoses,omitempty"`
	TargetDisease    []CodeableConcept `bson:"targetDisease,omitempty" json:"targetDisease,omitempty"`
	DoseStatus       *CodeableConcept  `bson:"doseStatus,omitempty" json:"doseStatus,omitempty"`
	DoseStatusReason *CodeableConcept  `bson:"doseStatusReason,omitempty" json:"doseStatusReason,omitempty"`
}

type ImmunizationPlus struct {
	Immunization                     `bson:",inline"`
	ImmunizationPlusRelatedResources `bson:",inline"`
}

type ImmunizationPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByRequester               *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByRequester,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer               *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedObservationResourcesReferencedByReaction                 *[]Observation                `bson:"_includedObservationResourcesReferencedByReaction,omitempty"`
	IncludedOrganizationResourcesReferencedByManufacturer            *[]Organization               `bson:"_includedOrganizationResourcesReferencedByManufacturer,omitempty"`
	IncludedPatientResourcesReferencedByPatient                      *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedLocationResourcesReferencedByLocation                    *[]Location                   `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                  *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref        *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref        *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                          *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref       *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                       *[]Order                      `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                      *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference               *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                  *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated           *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment          *[]OrderResponse              `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject      *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest            *[]ProcessResponse            `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger         *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                 *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingSupport *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingSupport,omitempty"`
}

func (i *ImmunizationPlusRelatedResources) GetIncludedPractitionerResourceReferencedByRequester() (practitioner *Practitioner, err error) {
	if i.IncludedPractitionerResourcesReferencedByRequester == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*i.IncludedPractitionerResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*i.IncludedPractitionerResourcesReferencedByRequester))
	} else if len(*i.IncludedPractitionerResourcesReferencedByRequester) == 1 {
		practitioner = &(*i.IncludedPractitionerResourcesReferencedByRequester)[0]
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPerformer() (practitioner *Practitioner, err error) {
	if i.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*i.IncludedPractitionerResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*i.IncludedPractitionerResourcesReferencedByPerformer))
	} else if len(*i.IncludedPractitionerResourcesReferencedByPerformer) == 1 {
		practitioner = &(*i.IncludedPractitionerResourcesReferencedByPerformer)[0]
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetIncludedObservationResourceReferencedByReaction() (observation *Observation, err error) {
	if i.IncludedObservationResourcesReferencedByReaction == nil {
		err = errors.New("Included observations not requested")
	} else if len(*i.IncludedObservationResourcesReferencedByReaction) > 1 {
		err = fmt.Errorf("Expected 0 or 1 observation, but found %d", len(*i.IncludedObservationResourcesReferencedByReaction))
	} else if len(*i.IncludedObservationResourcesReferencedByReaction) == 1 {
		observation = &(*i.IncludedObservationResourcesReferencedByReaction)[0]
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetIncludedOrganizationResourceReferencedByManufacturer() (organization *Organization, err error) {
	if i.IncludedOrganizationResourcesReferencedByManufacturer == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*i.IncludedOrganizationResourcesReferencedByManufacturer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*i.IncludedOrganizationResourcesReferencedByManufacturer))
	} else if len(*i.IncludedOrganizationResourcesReferencedByManufacturer) == 1 {
		organization = &(*i.IncludedOrganizationResourcesReferencedByManufacturer)[0]
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if i.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResourcesReferencedByPatient))
	} else if len(*i.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*i.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if i.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*i.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*i.IncludedLocationResourcesReferencedByLocation))
	} else if len(*i.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*i.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if i.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *i.RevIncludedListResourcesReferencingItem
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if i.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *i.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if i.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *i.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if i.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *i.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if i.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *i.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *i.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if i.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *i.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if i.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *i.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if i.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *i.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if i.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *i.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingSupport() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if i.RevIncludedImmunizationRecommendationResourcesReferencingSupport == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *i.RevIncludedImmunizationRecommendationResourcesReferencingSupport
	}
	return
}

func (i *ImmunizationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedPractitionerResourcesReferencedByRequester != nil {
		for _, r := range *i.IncludedPractitionerResourcesReferencedByRequester {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for _, r := range *i.IncludedPractitionerResourcesReferencedByPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedObservationResourcesReferencedByReaction != nil {
		for _, r := range *i.IncludedObservationResourcesReferencedByReaction {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedOrganizationResourcesReferencedByManufacturer != nil {
		for _, r := range *i.IncludedOrganizationResourcesReferencedByManufacturer {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *i.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedLocationResourcesReferencedByLocation != nil {
		for _, r := range *i.IncludedLocationResourcesReferencedByLocation {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (i *ImmunizationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	if i.RevIncludedImmunizationRecommendationResourcesReferencingSupport != nil {
		for _, r := range *i.RevIncludedImmunizationRecommendationResourcesReferencingSupport {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (i *ImmunizationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedPractitionerResourcesReferencedByRequester != nil {
		for _, r := range *i.IncludedPractitionerResourcesReferencedByRequester {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for _, r := range *i.IncludedPractitionerResourcesReferencedByPerformer {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedObservationResourcesReferencedByReaction != nil {
		for _, r := range *i.IncludedObservationResourcesReferencedByReaction {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedOrganizationResourcesReferencedByManufacturer != nil {
		for _, r := range *i.IncludedOrganizationResourcesReferencedByManufacturer {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *i.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if i.IncludedLocationResourcesReferencedByLocation != nil {
		for _, r := range *i.IncludedLocationResourcesReferencedByLocation {
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
	if i.RevIncludedImmunizationRecommendationResourcesReferencingSupport != nil {
		for _, r := range *i.RevIncludedImmunizationRecommendationResourcesReferencingSupport {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
