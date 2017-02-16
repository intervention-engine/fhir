// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
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

type RequestGroup struct {
	DomainResource        `bson:",inline"`
	Identifier            *Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject               *Reference                    `bson:"subject,omitempty" json:"subject,omitempty"`
	Context               *Reference                    `bson:"context,omitempty" json:"context,omitempty"`
	OccurrenceDateTime    *FHIRDateTime                 `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	Author                *Reference                    `bson:"author,omitempty" json:"author,omitempty"`
	ReasonCodeableConcept *CodeableConcept              `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                    `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                  []Annotation                  `bson:"note,omitempty" json:"note,omitempty"`
	Action                []RequestGroupActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *RequestGroup) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "RequestGroup"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to RequestGroup), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *RequestGroup) GetBSON() (interface{}, error) {
	x.ResourceType = "RequestGroup"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "requestGroup" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type requestGroup RequestGroup

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *RequestGroup) UnmarshalJSON(data []byte) (err error) {
	x2 := requestGroup{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = RequestGroup(x2)
		return x.checkResourceType()
	}
	return
}

func (x *RequestGroup) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "RequestGroup"
	} else if x.ResourceType != "RequestGroup" {
		return errors.New(fmt.Sprintf("Expected resourceType to be RequestGroup, instead received %s", x.ResourceType))
	}
	return nil
}

type RequestGroupActionComponent struct {
	BackboneElement     `bson:",inline"`
	ActionIdentifier    *Identifier                                `bson:"actionIdentifier,omitempty" json:"actionIdentifier,omitempty"`
	Label               string                                     `bson:"label,omitempty" json:"label,omitempty"`
	Title               string                                     `bson:"title,omitempty" json:"title,omitempty"`
	Description         string                                     `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent      string                                     `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	Code                []CodeableConcept                          `bson:"code,omitempty" json:"code,omitempty"`
	Documentation       []RelatedArtifact                          `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Condition           []RequestGroupActionConditionComponent     `bson:"condition,omitempty" json:"condition,omitempty"`
	RelatedAction       []RequestGroupActionRelatedActionComponent `bson:"relatedAction,omitempty" json:"relatedAction,omitempty"`
	TimingDateTime      *FHIRDateTime                              `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	TimingPeriod        *Period                                    `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDuration      *Quantity                                  `bson:"timingDuration,omitempty" json:"timingDuration,omitempty"`
	TimingRange         *Range                                     `bson:"timingRange,omitempty" json:"timingRange,omitempty"`
	TimingTiming        *Timing                                    `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	Participant         []Reference                                `bson:"participant,omitempty" json:"participant,omitempty"`
	Type                *Coding                                    `bson:"type,omitempty" json:"type,omitempty"`
	GroupingBehavior    string                                     `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty"`
	SelectionBehavior   string                                     `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty"`
	RequiredBehavior    string                                     `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty"`
	PrecheckBehavior    string                                     `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty"`
	CardinalityBehavior string                                     `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty"`
	Resource            *Reference                                 `bson:"resource,omitempty" json:"resource,omitempty"`
	Action              []RequestGroupActionComponent              `bson:"action,omitempty" json:"action,omitempty"`
}

type RequestGroupActionConditionComponent struct {
	BackboneElement `bson:",inline"`
	Kind            string `bson:"kind,omitempty" json:"kind,omitempty"`
	Description     string `bson:"description,omitempty" json:"description,omitempty"`
	Language        string `bson:"language,omitempty" json:"language,omitempty"`
	Expression      string `bson:"expression,omitempty" json:"expression,omitempty"`
}

type RequestGroupActionRelatedActionComponent struct {
	BackboneElement  `bson:",inline"`
	ActionIdentifier *Identifier `bson:"actionIdentifier,omitempty" json:"actionIdentifier,omitempty"`
	Relationship     string      `bson:"relationship,omitempty" json:"relationship,omitempty"`
	OffsetDuration   *Quantity   `bson:"offsetDuration,omitempty" json:"offsetDuration,omitempty"`
	OffsetRange      *Range      `bson:"offsetRange,omitempty" json:"offsetRange,omitempty"`
}

type RequestGroupPlus struct {
	RequestGroup                     `bson:",inline"`
	RequestGroupPlusRelatedResources `bson:",inline"`
}

type RequestGroupPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedGroupResourcesReferencedBySubject                   *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthor             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByAuthor                   *[]Device                `bson:"_includedDeviceResourcesReferencedByAuthor,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext           *[]EpisodeOfCare         `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext               *[]Encounter             `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter             *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPractitionerResourcesReferencedByParticipant        *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByParticipant,omitempty"`
	IncludedPatientResourcesReferencedByParticipant             *[]Patient               `bson:"_includedPatientResourcesReferencedByParticipant,omitempty"`
	IncludedPersonResourcesReferencedByParticipant              *[]Person                `bson:"_includedPersonResourcesReferencedByParticipant,omitempty"`
	IncludedRelatedPersonResourcesReferencedByParticipant       *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByParticipant,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                  *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic               *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject              *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse        *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource  *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon         *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                    *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                  *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces    *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition  *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces     *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition   *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity             *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
}

func (r *RequestGroupPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if r.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResourcesReferencedByPatient))
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*r.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if r.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*r.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*r.IncludedGroupResourcesReferencedBySubject))
	} else if len(*r.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*r.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if r.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResourcesReferencedBySubject))
	} else if len(*r.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*r.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if r.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*r.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*r.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*r.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*r.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedDeviceResourceReferencedByAuthor() (device *Device, err error) {
	if r.IncludedDeviceResourcesReferencedByAuthor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*r.IncludedDeviceResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*r.IncludedDeviceResourcesReferencedByAuthor))
	} else if len(*r.IncludedDeviceResourcesReferencedByAuthor) == 1 {
		device = &(*r.IncludedDeviceResourcesReferencedByAuthor)[0]
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByContext() (episodeOfCare *EpisodeOfCare, err error) {
	if r.IncludedEpisodeOfCareResourcesReferencedByContext == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*r.IncludedEpisodeOfCareResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*r.IncludedEpisodeOfCareResourcesReferencedByContext))
	} else if len(*r.IncludedEpisodeOfCareResourcesReferencedByContext) == 1 {
		episodeOfCare = &(*r.IncludedEpisodeOfCareResourcesReferencedByContext)[0]
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedEncounterResourceReferencedByContext() (encounter *Encounter, err error) {
	if r.IncludedEncounterResourcesReferencedByContext == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*r.IncludedEncounterResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*r.IncludedEncounterResourcesReferencedByContext))
	} else if len(*r.IncludedEncounterResourcesReferencedByContext) == 1 {
		encounter = &(*r.IncludedEncounterResourcesReferencedByContext)[0]
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if r.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*r.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*r.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*r.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*r.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByParticipant() (practitioners []Practitioner, err error) {
	if r.IncludedPractitionerResourcesReferencedByParticipant == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *r.IncludedPractitionerResourcesReferencedByParticipant
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedPatientResourcesReferencedByParticipant() (patients []Patient, err error) {
	if r.IncludedPatientResourcesReferencedByParticipant == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *r.IncludedPatientResourcesReferencedByParticipant
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedPersonResourcesReferencedByParticipant() (people []Person, err error) {
	if r.IncludedPersonResourcesReferencedByParticipant == nil {
		err = errors.New("Included people not requested")
	} else {
		people = *r.IncludedPersonResourcesReferencedByParticipant
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByParticipant() (relatedPeople []RelatedPerson, err error) {
	if r.IncludedRelatedPersonResourcesReferencedByParticipant == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *r.IncludedRelatedPersonResourcesReferencedByParticipant
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if r.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *r.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if r.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *r.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if r.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *r.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if r.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *r.RevIncludedListResourcesReferencingItem
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if r.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *r.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if r.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *r.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if r.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *r.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if r.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *r.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if r.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *r.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if r.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *r.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if r.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *r.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if r.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *r.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *r.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if r.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *r.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (r *RequestGroupPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByPatient {
			rsc := (*r.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedGroupResourcesReferencedBySubject {
			rsc := (*r.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedPatientResourcesReferencedBySubject {
			rsc := (*r.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *r.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*r.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *r.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*r.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *r.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*r.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *r.IncludedEncounterResourcesReferencedByContext {
			rsc := (*r.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *r.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*r.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for idx := range *r.IncludedPractitionerResourcesReferencedByParticipant {
			rsc := (*r.IncludedPractitionerResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPatientResourcesReferencedByParticipant != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByParticipant {
			rsc := (*r.IncludedPatientResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPersonResourcesReferencedByParticipant != nil {
		for idx := range *r.IncludedPersonResourcesReferencedByParticipant {
			rsc := (*r.IncludedPersonResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for idx := range *r.IncludedRelatedPersonResourcesReferencedByParticipant {
			rsc := (*r.IncludedRelatedPersonResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *RequestGroupPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingData {
			rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*r.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingTopic {
			rsc := (*r.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*r.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *RequestGroupPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByPatient {
			rsc := (*r.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedGroupResourcesReferencedBySubject {
			rsc := (*r.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedPatientResourcesReferencedBySubject {
			rsc := (*r.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *r.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*r.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *r.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*r.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *r.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*r.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *r.IncludedEncounterResourcesReferencedByContext {
			rsc := (*r.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *r.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*r.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for idx := range *r.IncludedPractitionerResourcesReferencedByParticipant {
			rsc := (*r.IncludedPractitionerResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPatientResourcesReferencedByParticipant != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByParticipant {
			rsc := (*r.IncludedPatientResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPersonResourcesReferencedByParticipant != nil {
		for idx := range *r.IncludedPersonResourcesReferencedByParticipant {
			rsc := (*r.IncludedPersonResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for idx := range *r.IncludedRelatedPersonResourcesReferencedByParticipant {
			rsc := (*r.IncludedRelatedPersonResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingData {
			rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*r.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingTopic {
			rsc := (*r.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *r.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*r.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *r.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*r.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*r.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
