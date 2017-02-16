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

type AllergyIntolerance struct {
	DomainResource     `bson:",inline"`
	Identifier         []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ClinicalStatus     string                                `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus string                                `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Type               string                                `bson:"type,omitempty" json:"type,omitempty"`
	Category           []string                              `bson:"category,omitempty" json:"category,omitempty"`
	Criticality        string                                `bson:"criticality,omitempty" json:"criticality,omitempty"`
	Code               *CodeableConcept                      `bson:"code,omitempty" json:"code,omitempty"`
	Patient            *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	OnsetDateTime      *FHIRDateTime                         `bson:"onsetDateTime,omitempty" json:"onsetDateTime,omitempty"`
	OnsetAge           *Quantity                             `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                               `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                                `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetString        string                                `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	AssertedDate       *FHIRDateTime                         `bson:"assertedDate,omitempty" json:"assertedDate,omitempty"`
	Recorder           *Reference                            `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Asserter           *Reference                            `bson:"asserter,omitempty" json:"asserter,omitempty"`
	LastOccurrence     *FHIRDateTime                         `bson:"lastOccurrence,omitempty" json:"lastOccurrence,omitempty"`
	Note               []Annotation                          `bson:"note,omitempty" json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReactionComponent `bson:"reaction,omitempty" json:"reaction,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *AllergyIntolerance) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "AllergyIntolerance"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to AllergyIntolerance), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *AllergyIntolerance) GetBSON() (interface{}, error) {
	x.ResourceType = "AllergyIntolerance"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "allergyIntolerance" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type allergyIntolerance AllergyIntolerance

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *AllergyIntolerance) UnmarshalJSON(data []byte) (err error) {
	x2 := allergyIntolerance{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = AllergyIntolerance(x2)
		return x.checkResourceType()
	}
	return
}

func (x *AllergyIntolerance) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "AllergyIntolerance"
	} else if x.ResourceType != "AllergyIntolerance" {
		return errors.New(fmt.Sprintf("Expected resourceType to be AllergyIntolerance, instead received %s", x.ResourceType))
	}
	return nil
}

type AllergyIntoleranceReactionComponent struct {
	BackboneElement `bson:",inline"`
	Substance       *CodeableConcept  `bson:"substance,omitempty" json:"substance,omitempty"`
	Certainty       string            `bson:"certainty,omitempty" json:"certainty,omitempty"`
	Manifestation   []CodeableConcept `bson:"manifestation,omitempty" json:"manifestation,omitempty"`
	Description     string            `bson:"description,omitempty" json:"description,omitempty"`
	Onset           *FHIRDateTime     `bson:"onset,omitempty" json:"onset,omitempty"`
	Severity        string            `bson:"severity,omitempty" json:"severity,omitempty"`
	ExposureRoute   *CodeableConcept  `bson:"exposureRoute,omitempty" json:"exposureRoute,omitempty"`
	Note            []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}

type AllergyIntolerancePlus struct {
	AllergyIntolerance                     `bson:",inline"`
	AllergyIntolerancePlusRelatedResources `bson:",inline"`
}

type AllergyIntolerancePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByRecorder                    *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByRecorder,omitempty"`
	IncludedPatientResourcesReferencedByRecorder                         *[]Patient                    `bson:"_includedPatientResourcesReferencedByRecorder,omitempty"`
	IncludedPractitionerResourcesReferencedByAsserter                    *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByAsserter,omitempty"`
	IncludedPatientResourcesReferencedByAsserter                         *[]Patient                    `bson:"_includedPatientResourcesReferencedByAsserter,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAsserter                   *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByAsserter,omitempty"`
	IncludedPatientResourcesReferencedByPatient                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref            *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref            *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                           *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref           *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                        *[]Contract                   `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                       *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                         *[]Contract                   `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                  *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                 *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource           *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                  *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                     *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                      *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                      *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                           *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                           *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                              *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces             *[]DiagnosticRequest          `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon              *[]DiagnosticRequest          `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition           *[]DiagnosticRequest          `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces              *[]DeviceUseRequest           `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon               *[]DeviceUseRequest           `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition            *[]DeviceUseRequest           `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                          *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                      *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                    *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                      *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated               *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject          *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest                *[]ProcessResponse            `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingProblem             *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingProblem,omitempty"`
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPractitionerResourceReferencedByRecorder() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByRecorder == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByRecorder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByRecorder))
	} else if len(*a.IncludedPractitionerResourcesReferencedByRecorder) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByRecorder)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPatientResourceReferencedByRecorder() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByRecorder == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByRecorder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByRecorder))
	} else if len(*a.IncludedPatientResourcesReferencedByRecorder) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByRecorder)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPractitionerResourceReferencedByAsserter() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByAsserter == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByAsserter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByAsserter))
	} else if len(*a.IncludedPractitionerResourcesReferencedByAsserter) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByAsserter)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPatientResourceReferencedByAsserter() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByAsserter == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByAsserter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByAsserter))
	} else if len(*a.IncludedPatientResourcesReferencedByAsserter) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByAsserter)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAsserter() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedRelatedPersonResourcesReferencedByAsserter == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByAsserter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedRelatedPersonResourcesReferencedByAsserter))
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByAsserter) == 1 {
		relatedPerson = &(*a.IncludedRelatedPersonResourcesReferencedByAsserter)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatient))
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if a.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *a.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if a.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *a.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if a.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *a.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if a.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *a.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if a.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *a.RevIncludedListResourcesReferencingItem
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if a.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *a.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if a.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *a.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if a.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *a.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if a.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *a.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if a.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *a.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if a.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *a.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if a.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *a.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if a.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *a.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *a.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if a.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *a.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingProblem() (clinicalImpressions []ClinicalImpression, err error) {
	if a.RevIncludedClinicalImpressionResourcesReferencingProblem == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *a.RevIncludedClinicalImpressionResourcesReferencingProblem
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByRecorder {
			rsc := (*a.IncludedPractitionerResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByRecorder {
			rsc := (*a.IncludedPatientResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByAsserter {
			rsc := (*a.IncludedPractitionerResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByAsserter {
			rsc := (*a.IncludedPatientResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByAsserter {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingData {
			rsc := (*a.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*a.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*a.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingProblem != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingProblem {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingProblem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByRecorder {
			rsc := (*a.IncludedPractitionerResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByRecorder {
			rsc := (*a.IncludedPatientResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByAsserter {
			rsc := (*a.IncludedPractitionerResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByAsserter {
			rsc := (*a.IncludedPatientResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByAsserter {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingData {
			rsc := (*a.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingTopic {
			rsc := (*a.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*a.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *a.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*a.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *a.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*a.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*a.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingProblem != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingProblem {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingProblem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
