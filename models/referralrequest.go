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

type ReferralRequest struct {
	DomainResource     `bson:",inline"`
	Identifier         []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition         []Reference                        `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn            []Reference                        `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces           []Reference                        `bson:"replaces,omitempty" json:"replaces,omitempty"`
	GroupIdentifier    *Identifier                        `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status             string                             `bson:"status,omitempty" json:"status,omitempty"`
	Intent             string                             `bson:"intent,omitempty" json:"intent,omitempty"`
	Type               *CodeableConcept                   `bson:"type,omitempty" json:"type,omitempty"`
	Priority           string                             `bson:"priority,omitempty" json:"priority,omitempty"`
	ServiceRequested   []CodeableConcept                  `bson:"serviceRequested,omitempty" json:"serviceRequested,omitempty"`
	Subject            *Reference                         `bson:"subject,omitempty" json:"subject,omitempty"`
	Context            *Reference                         `bson:"context,omitempty" json:"context,omitempty"`
	OccurrenceDateTime *FHIRDateTime                      `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                            `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	AuthoredOn         *FHIRDateTime                      `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester          *ReferralRequestRequesterComponent `bson:"requester,omitempty" json:"requester,omitempty"`
	Specialty          *CodeableConcept                   `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Recipient          []Reference                        `bson:"recipient,omitempty" json:"recipient,omitempty"`
	ReasonCode         []CodeableConcept                  `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []Reference                        `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Description        string                             `bson:"description,omitempty" json:"description,omitempty"`
	SupportingInfo     []Reference                        `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Note               []Annotation                       `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory    []Reference                        `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ReferralRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ReferralRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ReferralRequest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ReferralRequest) GetBSON() (interface{}, error) {
	x.ResourceType = "ReferralRequest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "referralRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type referralRequest ReferralRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ReferralRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := referralRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = ReferralRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ReferralRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ReferralRequest"
	} else if x.ResourceType != "ReferralRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ReferralRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type ReferralRequestRequesterComponent struct {
	BackboneElement `bson:",inline"`
	Agent           *Reference `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf      *Reference `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}

type ReferralRequestPlus struct {
	ReferralRequest                     `bson:",inline"`
	ReferralRequestPlusRelatedResources `bson:",inline"`
}

type ReferralRequestPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByRequester              *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRequester,omitempty"`
	IncludedOrganizationResourcesReferencedByRequester              *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRequester,omitempty"`
	IncludedDeviceResourcesReferencedByRequester                    *[]Device                `bson:"_includedDeviceResourcesReferencedByRequester,omitempty"`
	IncludedPatientResourcesReferencedByRequester                   *[]Patient               `bson:"_includedPatientResourcesReferencedByRequester,omitempty"`
	IncludedRelatedPersonResourcesReferencedByRequester             *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByRequester,omitempty"`
	IncludedReferralRequestResourcesReferencedByReplaces            *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByReplaces,omitempty"`
	IncludedGroupResourcesReferencedBySubject                       *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                     *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                 *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedReferralRequestResourcesReferencedByBasedon             *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByBasedon,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                    *[]CarePlan              `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedProcedureRequestResourcesReferencedByBasedon            *[]ProcedureRequest      `bson:"_includedProcedureRequestResourcesReferencedByBasedon,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByRecipient              *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRecipient,omitempty"`
	IncludedOrganizationResourcesReferencedByRecipient              *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRecipient,omitempty"`
	IncludedHealthcareServiceResourcesReferencedByRecipient         *[]HealthcareService     `bson:"_includedHealthcareServiceResourcesReferencedByRecipient,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext               *[]EpisodeOfCare         `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext                   *[]Encounter             `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
	IncludedPlanDefinitionResourcesReferencedByDefinition           *[]PlanDefinition        `bson:"_includedPlanDefinitionResourcesReferencedByDefinition,omitempty"`
	IncludedActivityDefinitionResourcesReferencedByDefinition       *[]ActivityDefinition    `bson:"_includedActivityDefinitionResourcesReferencedByDefinition,omitempty"`
	RevIncludedAppointmentResourcesReferencingIncomingreferral      *[]Appointment           `bson:"_revIncludedAppointmentResourcesReferencingIncomingreferral,omitempty"`
	RevIncludedReferralRequestResourcesReferencingReplaces          *[]ReferralRequest       `bson:"_revIncludedReferralRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedReferralRequestResourcesReferencingBasedon           *[]ReferralRequest       `bson:"_revIncludedReferralRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                 *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1            *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2            *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref      *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingSubject                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest             *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse            *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource      *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingBasedon              *[]ImagingStudy          `bson:"_revIncludedImagingStudyResourcesReferencingBasedon,omitempty"`
	RevIncludedEncounterResourcesReferencingIncomingreferral        *[]Encounter             `bson:"_revIncludedEncounterResourcesReferencingIncomingreferral,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor       *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom     *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor     *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof      *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson       *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof              *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon             *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor      *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom    *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor    *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof     *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition           *[]RequestGroup          `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon             *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest        *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                 *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                        *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference        *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral    *[]EpisodeOfCare         `bson:"_revIncludedEpisodeOfCareResourcesReferencingIncomingreferral,omitempty"`
	RevIncludedProcedureResourcesReferencingBasedon                 *[]Procedure             `bson:"_revIncludedProcedureResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces         *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon          *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedObservationResourcesReferencingBasedon               *[]Observation           `bson:"_revIncludedObservationResourcesReferencingBasedon,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom               *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor               *[]Library               `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                *[]Library               `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon      *[]CommunicationRequest  `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                     *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingBasedon          *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingBasedon,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                 *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail          *[]Condition             `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject               *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                 *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated          *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingBasedon     *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingBasedon,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject     *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest           *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAction         *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor          *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof         *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedPractitionerResourceReferencedByRequester() (practitioner *Practitioner, err error) {
	if r.IncludedPractitionerResourcesReferencedByRequester == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*r.IncludedPractitionerResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*r.IncludedPractitionerResourcesReferencedByRequester))
	} else if len(*r.IncludedPractitionerResourcesReferencedByRequester) == 1 {
		practitioner = &(*r.IncludedPractitionerResourcesReferencedByRequester)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedOrganizationResourceReferencedByRequester() (organization *Organization, err error) {
	if r.IncludedOrganizationResourcesReferencedByRequester == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*r.IncludedOrganizationResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*r.IncludedOrganizationResourcesReferencedByRequester))
	} else if len(*r.IncludedOrganizationResourcesReferencedByRequester) == 1 {
		organization = &(*r.IncludedOrganizationResourcesReferencedByRequester)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedDeviceResourceReferencedByRequester() (device *Device, err error) {
	if r.IncludedDeviceResourcesReferencedByRequester == nil {
		err = errors.New("Included devices not requested")
	} else if len(*r.IncludedDeviceResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*r.IncludedDeviceResourcesReferencedByRequester))
	} else if len(*r.IncludedDeviceResourcesReferencedByRequester) == 1 {
		device = &(*r.IncludedDeviceResourcesReferencedByRequester)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByRequester() (patient *Patient, err error) {
	if r.IncludedPatientResourcesReferencedByRequester == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResourcesReferencedByRequester))
	} else if len(*r.IncludedPatientResourcesReferencedByRequester) == 1 {
		patient = &(*r.IncludedPatientResourcesReferencedByRequester)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByRequester() (relatedPerson *RelatedPerson, err error) {
	if r.IncludedRelatedPersonResourcesReferencedByRequester == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*r.IncludedRelatedPersonResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*r.IncludedRelatedPersonResourcesReferencedByRequester))
	} else if len(*r.IncludedRelatedPersonResourcesReferencedByRequester) == 1 {
		relatedPerson = &(*r.IncludedRelatedPersonResourcesReferencedByRequester)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedReferralRequestResourcesReferencedByReplaces() (referralRequests []ReferralRequest, err error) {
	if r.IncludedReferralRequestResourcesReferencedByReplaces == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *r.IncludedReferralRequestResourcesReferencedByReplaces
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if r.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*r.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*r.IncludedGroupResourcesReferencedBySubject))
	} else if len(*r.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*r.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if r.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResourcesReferencedBySubject))
	} else if len(*r.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*r.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if r.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*r.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*r.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*r.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*r.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedReferralRequestResourcesReferencedByBasedon() (referralRequests []ReferralRequest, err error) {
	if r.IncludedReferralRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included referralRequests not requested")
	} else {
		referralRequests = *r.IncludedReferralRequestResourcesReferencedByBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if r.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *r.IncludedCarePlanResourcesReferencedByBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedProcedureRequestResourcesReferencedByBasedon() (procedureRequests []ProcedureRequest, err error) {
	if r.IncludedProcedureRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included procedureRequests not requested")
	} else {
		procedureRequests = *r.IncludedProcedureRequestResourcesReferencedByBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if r.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResourcesReferencedByPatient))
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*r.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByRecipient() (practitioners []Practitioner, err error) {
	if r.IncludedPractitionerResourcesReferencedByRecipient == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *r.IncludedPractitionerResourcesReferencedByRecipient
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByRecipient() (organizations []Organization, err error) {
	if r.IncludedOrganizationResourcesReferencedByRecipient == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *r.IncludedOrganizationResourcesReferencedByRecipient
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedHealthcareServiceResourcesReferencedByRecipient() (healthcareServices []HealthcareService, err error) {
	if r.IncludedHealthcareServiceResourcesReferencedByRecipient == nil {
		err = errors.New("Included healthcareServices not requested")
	} else {
		healthcareServices = *r.IncludedHealthcareServiceResourcesReferencedByRecipient
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByContext() (episodeOfCare *EpisodeOfCare, err error) {
	if r.IncludedEpisodeOfCareResourcesReferencedByContext == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*r.IncludedEpisodeOfCareResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*r.IncludedEpisodeOfCareResourcesReferencedByContext))
	} else if len(*r.IncludedEpisodeOfCareResourcesReferencedByContext) == 1 {
		episodeOfCare = &(*r.IncludedEpisodeOfCareResourcesReferencedByContext)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedEncounterResourceReferencedByContext() (encounter *Encounter, err error) {
	if r.IncludedEncounterResourcesReferencedByContext == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*r.IncludedEncounterResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*r.IncludedEncounterResourcesReferencedByContext))
	} else if len(*r.IncludedEncounterResourcesReferencedByContext) == 1 {
		encounter = &(*r.IncludedEncounterResourcesReferencedByContext)[0]
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedPlanDefinitionResourcesReferencedByDefinition() (planDefinitions []PlanDefinition, err error) {
	if r.IncludedPlanDefinitionResourcesReferencedByDefinition == nil {
		err = errors.New("Included planDefinitions not requested")
	} else {
		planDefinitions = *r.IncludedPlanDefinitionResourcesReferencedByDefinition
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedActivityDefinitionResourcesReferencedByDefinition() (activityDefinitions []ActivityDefinition, err error) {
	if r.IncludedActivityDefinitionResourcesReferencedByDefinition == nil {
		err = errors.New("Included activityDefinitions not requested")
	} else {
		activityDefinitions = *r.IncludedActivityDefinitionResourcesReferencedByDefinition
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingIncomingreferral() (appointments []Appointment, err error) {
	if r.RevIncludedAppointmentResourcesReferencingIncomingreferral == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *r.RevIncludedAppointmentResourcesReferencingIncomingreferral
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingReplaces() (referralRequests []ReferralRequest, err error) {
	if r.RevIncludedReferralRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *r.RevIncludedReferralRequestResourcesReferencingReplaces
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingBasedon() (referralRequests []ReferralRequest, err error) {
	if r.RevIncludedReferralRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *r.RevIncludedReferralRequestResourcesReferencingBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if r.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *r.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if r.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *r.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if r.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *r.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingBasedon() (imagingStudies []ImagingStudy, err error) {
	if r.RevIncludedImagingStudyResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *r.RevIncludedImagingStudyResourcesReferencingBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingIncomingreferral() (encounters []Encounter, err error) {
	if r.RevIncludedEncounterResourcesReferencingIncomingreferral == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *r.RevIncludedEncounterResourcesReferencingIncomingreferral
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if r.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *r.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if r.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *r.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if r.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *r.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if r.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *r.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if r.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *r.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if r.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *r.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *r.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if r.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *r.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if r.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *r.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingIncomingreferral() (episodeOfCares []EpisodeOfCare, err error) {
	if r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingBasedon() (procedures []Procedure, err error) {
	if r.RevIncludedProcedureResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *r.RevIncludedProcedureResourcesReferencingBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if r.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *r.RevIncludedListResourcesReferencingItem
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if r.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *r.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if r.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *r.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedObservationResourcesReferencingBasedon() (observations []Observation, err error) {
	if r.RevIncludedObservationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *r.RevIncludedObservationResourcesReferencingBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *r.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if r.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *r.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingBasedon() (diagnosticReports []DiagnosticReport, err error) {
	if r.RevIncludedDiagnosticReportResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *r.RevIncludedDiagnosticReportResourcesReferencingBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if r.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *r.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if r.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *r.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *r.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingBasedon() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingBasedon
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if r.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *r.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAction() (clinicalImpressions []ClinicalImpression, err error) {
	if r.RevIncludedClinicalImpressionResourcesReferencingAction == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *r.RevIncludedClinicalImpressionResourcesReferencingAction
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*r.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*r.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*r.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByRequester {
			rsc := (*r.IncludedPatientResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedRelatedPersonResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedRelatedPersonResourcesReferencedByRequester {
			rsc := (*r.IncludedRelatedPersonResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedReferralRequestResourcesReferencedByReplaces != nil {
		for idx := range *r.IncludedReferralRequestResourcesReferencedByReplaces {
			rsc := (*r.IncludedReferralRequestResourcesReferencedByReplaces)[idx]
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
	if r.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *r.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*r.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedReferralRequestResourcesReferencedByBasedon != nil {
		for idx := range *r.IncludedReferralRequestResourcesReferencedByBasedon {
			rsc := (*r.IncludedReferralRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *r.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*r.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedProcedureRequestResourcesReferencedByBasedon != nil {
		for idx := range *r.IncludedProcedureRequestResourcesReferencedByBasedon {
			rsc := (*r.IncludedProcedureRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByPatient {
			rsc := (*r.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPractitionerResourcesReferencedByRecipient != nil {
		for idx := range *r.IncludedPractitionerResourcesReferencedByRecipient {
			rsc := (*r.IncludedPractitionerResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedOrganizationResourcesReferencedByRecipient != nil {
		for idx := range *r.IncludedOrganizationResourcesReferencedByRecipient {
			rsc := (*r.IncludedOrganizationResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedHealthcareServiceResourcesReferencedByRecipient != nil {
		for idx := range *r.IncludedHealthcareServiceResourcesReferencedByRecipient {
			rsc := (*r.IncludedHealthcareServiceResourcesReferencedByRecipient)[idx]
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
	if r.IncludedPlanDefinitionResourcesReferencedByDefinition != nil {
		for idx := range *r.IncludedPlanDefinitionResourcesReferencedByDefinition {
			rsc := (*r.IncludedPlanDefinitionResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedActivityDefinitionResourcesReferencedByDefinition != nil {
		for idx := range *r.IncludedActivityDefinitionResourcesReferencedByDefinition {
			rsc := (*r.IncludedActivityDefinitionResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.RevIncludedAppointmentResourcesReferencingIncomingreferral != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingIncomingreferral {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedReferralRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedReferralRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedReferralRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedReferralRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedReferralRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedReferralRequestResourcesReferencingBasedon)[idx]
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
	if r.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*r.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*r.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*r.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*r.RevIncludedContractResourcesReferencingTermtopic)[idx]
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
	if r.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*r.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEncounterResourcesReferencingIncomingreferral != nil {
		for idx := range *r.RevIncludedEncounterResourcesReferencingIncomingreferral {
			rsc := (*r.RevIncludedEncounterResourcesReferencingIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*r.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *r.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*r.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if r.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *r.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*r.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral != nil {
		for idx := range *r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral {
			rsc := (*r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedProcedureResourcesReferencingBasedon {
			rsc := (*r.RevIncludedProcedureResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*r.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *r.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*r.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if r.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
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
	if r.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *ReferralRequestPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*r.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*r.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*r.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByRequester {
			rsc := (*r.IncludedPatientResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedRelatedPersonResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedRelatedPersonResourcesReferencedByRequester {
			rsc := (*r.IncludedRelatedPersonResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedReferralRequestResourcesReferencedByReplaces != nil {
		for idx := range *r.IncludedReferralRequestResourcesReferencedByReplaces {
			rsc := (*r.IncludedReferralRequestResourcesReferencedByReplaces)[idx]
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
	if r.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *r.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*r.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedReferralRequestResourcesReferencedByBasedon != nil {
		for idx := range *r.IncludedReferralRequestResourcesReferencedByBasedon {
			rsc := (*r.IncludedReferralRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *r.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*r.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedProcedureRequestResourcesReferencedByBasedon != nil {
		for idx := range *r.IncludedProcedureRequestResourcesReferencedByBasedon {
			rsc := (*r.IncludedProcedureRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByPatient {
			rsc := (*r.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPractitionerResourcesReferencedByRecipient != nil {
		for idx := range *r.IncludedPractitionerResourcesReferencedByRecipient {
			rsc := (*r.IncludedPractitionerResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedOrganizationResourcesReferencedByRecipient != nil {
		for idx := range *r.IncludedOrganizationResourcesReferencedByRecipient {
			rsc := (*r.IncludedOrganizationResourcesReferencedByRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedHealthcareServiceResourcesReferencedByRecipient != nil {
		for idx := range *r.IncludedHealthcareServiceResourcesReferencedByRecipient {
			rsc := (*r.IncludedHealthcareServiceResourcesReferencedByRecipient)[idx]
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
	if r.IncludedPlanDefinitionResourcesReferencedByDefinition != nil {
		for idx := range *r.IncludedPlanDefinitionResourcesReferencedByDefinition {
			rsc := (*r.IncludedPlanDefinitionResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedActivityDefinitionResourcesReferencedByDefinition != nil {
		for idx := range *r.IncludedActivityDefinitionResourcesReferencedByDefinition {
			rsc := (*r.IncludedActivityDefinitionResourcesReferencedByDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAppointmentResourcesReferencingIncomingreferral != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingIncomingreferral {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedReferralRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedReferralRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedReferralRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedReferralRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedReferralRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedReferralRequestResourcesReferencingBasedon)[idx]
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
	if r.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*r.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*r.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*r.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*r.RevIncludedContractResourcesReferencingTermtopic)[idx]
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
	if r.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*r.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEncounterResourcesReferencingIncomingreferral != nil {
		for idx := range *r.RevIncludedEncounterResourcesReferencingIncomingreferral {
			rsc := (*r.RevIncludedEncounterResourcesReferencingIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*r.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *r.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*r.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if r.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *r.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*r.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral != nil {
		for idx := range *r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral {
			rsc := (*r.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedProcedureResourcesReferencingBasedon {
			rsc := (*r.RevIncludedProcedureResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *r.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*r.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*r.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *r.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*r.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if r.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
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
	if r.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
