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

type QuestionnaireResponse struct {
	DomainResource `bson:",inline"`
	Identifier     *Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Questionnaire  *Reference                           `bson:"questionnaire,omitempty" json:"questionnaire,omitempty"`
	Status         string                               `bson:"status,omitempty" json:"status,omitempty"`
	Subject        *Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Author         *Reference                           `bson:"author,omitempty" json:"author,omitempty"`
	Authored       *FHIRDateTime                        `bson:"authored,omitempty" json:"authored,omitempty"`
	Source         *Reference                           `bson:"source,omitempty" json:"source,omitempty"`
	Encounter      *Reference                           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Group          *QuestionnaireResponseGroupComponent `bson:"group,omitempty" json:"group,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "QuestionnaireResponse"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to QuestionnaireResponse), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *QuestionnaireResponse) GetBSON() (interface{}, error) {
	x.ResourceType = "QuestionnaireResponse"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "questionnaireResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type questionnaireResponse QuestionnaireResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *QuestionnaireResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := questionnaireResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = QuestionnaireResponse(x2)
		return x.checkResourceType()
	}
	return
}

func (x *QuestionnaireResponse) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "QuestionnaireResponse"
	} else if x.ResourceType != "QuestionnaireResponse" {
		return errors.New(fmt.Sprintf("Expected resourceType to be QuestionnaireResponse, instead received %s", x.ResourceType))
	}
	return nil
}

type QuestionnaireResponseGroupComponent struct {
	BackboneElement `bson:",inline"`
	LinkId          string                                   `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Title           string                                   `bson:"title,omitempty" json:"title,omitempty"`
	Text            string                                   `bson:"text,omitempty" json:"text,omitempty"`
	Subject         *Reference                               `bson:"subject,omitempty" json:"subject,omitempty"`
	Group           []QuestionnaireResponseGroupComponent    `bson:"group,omitempty" json:"group,omitempty"`
	Question        []QuestionnaireResponseQuestionComponent `bson:"question,omitempty" json:"question,omitempty"`
}

type QuestionnaireResponseQuestionComponent struct {
	BackboneElement `bson:",inline"`
	LinkId          string                                         `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Text            string                                         `bson:"text,omitempty" json:"text,omitempty"`
	Answer          []QuestionnaireResponseQuestionAnswerComponent `bson:"answer,omitempty" json:"answer,omitempty"`
}

type QuestionnaireResponseQuestionAnswerComponent struct {
	BackboneElement `bson:",inline"`
	ValueBoolean    *bool                                 `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueDecimal    *float64                              `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueInteger    *int32                                `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate       *FHIRDateTime                         `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime   *FHIRDateTime                         `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueInstant    *FHIRDateTime                         `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueTime       *FHIRDateTime                         `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString     string                                `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueUri        string                                `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueAttachment *Attachment                           `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCoding     *Coding                               `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueQuantity   *Quantity                             `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueReference  *Reference                            `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Group           []QuestionnaireResponseGroupComponent `bson:"group,omitempty" json:"group,omitempty"`
}

type QuestionnaireResponsePlus struct {
	QuestionnaireResponse                     `bson:",inline"`
	QuestionnaireResponsePlusRelatedResources `bson:",inline"`
}

type QuestionnaireResponsePlusRelatedResources struct {
	IncludedQuestionnaireResourcesReferencedByQuestionnaire        *[]Questionnaire         `bson:"_includedQuestionnaireResourcesReferencedByQuestionnaire,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthor                *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByAuthor                      *[]Device                `bson:"_includedDeviceResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByAuthor                     *[]Patient               `bson:"_includedPatientResourcesReferencedByAuthor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAuthor               *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByPatient                    *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPractitionerResourcesReferencedBySource                *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySource,omitempty"`
	IncludedPatientResourcesReferencedBySource                     *[]Patient               `bson:"_includedPatientResourcesReferencedBySource,omitempty"`
	IncludedRelatedPersonResourcesReferencedBySource               *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedBySource,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedListResourcesReferencingItem                        *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref     *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                     *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedObservationResourcesReferencingRelatedtarget        *[]Observation           `bson:"_revIncludedObservationResourcesReferencingRelatedtarget,omitempty"`
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

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedQuestionnaireResourceReferencedByQuestionnaire() (questionnaire *Questionnaire, err error) {
	if q.IncludedQuestionnaireResourcesReferencedByQuestionnaire == nil {
		err = errors.New("Included questionnaires not requested")
	} else if len(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire) > 1 {
		err = fmt.Errorf("Expected 0 or 1 questionnaire, but found %d", len(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire))
	} else if len(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire) == 1 {
		questionnaire = &(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if q.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*q.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*q.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*q.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*q.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedDeviceResourceReferencedByAuthor() (device *Device, err error) {
	if q.IncludedDeviceResourcesReferencedByAuthor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*q.IncludedDeviceResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*q.IncludedDeviceResourcesReferencedByAuthor))
	} else if len(*q.IncludedDeviceResourcesReferencedByAuthor) == 1 {
		device = &(*q.IncludedDeviceResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPatientResourceReferencedByAuthor() (patient *Patient, err error) {
	if q.IncludedPatientResourcesReferencedByAuthor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedPatientResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedPatientResourcesReferencedByAuthor))
	} else if len(*q.IncludedPatientResourcesReferencedByAuthor) == 1 {
		patient = &(*q.IncludedPatientResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAuthor() (relatedPerson *RelatedPerson, err error) {
	if q.IncludedRelatedPersonResourcesReferencedByAuthor == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*q.IncludedRelatedPersonResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*q.IncludedRelatedPersonResourcesReferencedByAuthor))
	} else if len(*q.IncludedRelatedPersonResourcesReferencedByAuthor) == 1 {
		relatedPerson = &(*q.IncludedRelatedPersonResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if q.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedPatientResourcesReferencedByPatient))
	} else if len(*q.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*q.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if q.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*q.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*q.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*q.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*q.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPractitionerResourceReferencedBySource() (practitioner *Practitioner, err error) {
	if q.IncludedPractitionerResourcesReferencedBySource == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*q.IncludedPractitionerResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*q.IncludedPractitionerResourcesReferencedBySource))
	} else if len(*q.IncludedPractitionerResourcesReferencedBySource) == 1 {
		practitioner = &(*q.IncludedPractitionerResourcesReferencedBySource)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPatientResourceReferencedBySource() (patient *Patient, err error) {
	if q.IncludedPatientResourcesReferencedBySource == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedPatientResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedPatientResourcesReferencedBySource))
	} else if len(*q.IncludedPatientResourcesReferencedBySource) == 1 {
		patient = &(*q.IncludedPatientResourcesReferencedBySource)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedBySource() (relatedPerson *RelatedPerson, err error) {
	if q.IncludedRelatedPersonResourcesReferencedBySource == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*q.IncludedRelatedPersonResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*q.IncludedRelatedPersonResourcesReferencedBySource))
	} else if len(*q.IncludedRelatedPersonResourcesReferencedBySource) == 1 {
		relatedPerson = &(*q.IncludedRelatedPersonResourcesReferencedBySource)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if q.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *q.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if q.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *q.RevIncludedListResourcesReferencingItem
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if q.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *q.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedObservationResourcesReferencingRelatedtarget() (observations []Observation, err error) {
	if q.RevIncludedObservationResourcesReferencingRelatedtarget == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *q.RevIncludedObservationResourcesReferencingRelatedtarget
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if q.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *q.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if q.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *q.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *q.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if q.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *q.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if q.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *q.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if q.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *q.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingInvestigation() (clinicalImpressions []ClinicalImpression, err error) {
	if q.RevIncludedClinicalImpressionResourcesReferencingInvestigation == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *q.RevIncludedClinicalImpressionResourcesReferencingInvestigation
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if q.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *q.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.IncludedQuestionnaireResourcesReferencedByQuestionnaire != nil {
		for idx := range *q.IncludedQuestionnaireResourcesReferencedByQuestionnaire {
			rsc := (*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*q.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*q.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*q.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByPatient {
			rsc := (*q.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *q.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*q.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*q.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPatientResourcesReferencedBySource {
			rsc := (*q.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *q.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*q.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedObservationResourcesReferencingRelatedtarget != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingRelatedtarget {
			rsc := (*q.RevIncludedObservationResourcesReferencingRelatedtarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*q.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *q.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*q.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*q.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*q.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *q.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*q.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *q.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*q.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*q.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.IncludedQuestionnaireResourcesReferencedByQuestionnaire != nil {
		for idx := range *q.IncludedQuestionnaireResourcesReferencedByQuestionnaire {
			rsc := (*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*q.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*q.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*q.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByPatient {
			rsc := (*q.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *q.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*q.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*q.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPatientResourcesReferencedBySource {
			rsc := (*q.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *q.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*q.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedObservationResourcesReferencingRelatedtarget != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingRelatedtarget {
			rsc := (*q.RevIncludedObservationResourcesReferencingRelatedtarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*q.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *q.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*q.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*q.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*q.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *q.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*q.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *q.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*q.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*q.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
