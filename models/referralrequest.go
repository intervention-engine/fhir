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

type ReferralRequest struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn               []Reference       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Parent                *Identifier       `bson:"parent,omitempty" json:"parent,omitempty"`
	Status                string            `bson:"status,omitempty" json:"status,omitempty"`
	Category              string            `bson:"category,omitempty" json:"category,omitempty"`
	Type                  *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Priority              *CodeableConcept  `bson:"priority,omitempty" json:"priority,omitempty"`
	Patient               *Reference        `bson:"patient,omitempty" json:"patient,omitempty"`
	Context               *Reference        `bson:"context,omitempty" json:"context,omitempty"`
	FulfillmentTime       *Period           `bson:"fulfillmentTime,omitempty" json:"fulfillmentTime,omitempty"`
	Authored              *FHIRDateTime     `bson:"authored,omitempty" json:"authored,omitempty"`
	Requester             *Reference        `bson:"requester,omitempty" json:"requester,omitempty"`
	Specialty             *CodeableConcept  `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Recipient             []Reference       `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Reason                *CodeableConcept  `bson:"reason,omitempty" json:"reason,omitempty"`
	Description           string            `bson:"description,omitempty" json:"description,omitempty"`
	ServiceRequested      []CodeableConcept `bson:"serviceRequested,omitempty" json:"serviceRequested,omitempty"`
	SupportingInformation []Reference       `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
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
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
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

type ReferralRequestPlus struct {
	ReferralRequest                     `bson:",inline"`
	ReferralRequestPlusRelatedResources `bson:",inline"`
}

type ReferralRequestPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByRequester             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRequester,omitempty"`
	IncludedOrganizationResourcesReferencedByRequester             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRequester,omitempty"`
	IncludedPatientResourcesReferencedByRequester                  *[]Patient               `bson:"_includedPatientResourcesReferencedByRequester,omitempty"`
	IncludedPatientResourcesReferencedByPatient                    *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByRecipient             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRecipient,omitempty"`
	IncludedOrganizationResourcesReferencedByRecipient             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRecipient,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext              *[]EpisodeOfCare         `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext                  *[]Encounter             `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
	IncludedReferralRequestResourcesReferencedByBasedon            *[]ReferralRequest       `bson:"_includedReferralRequestResourcesReferencedByBasedon,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                   *[]CarePlan              `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedDiagnosticOrderResourcesReferencedByBasedon            *[]DiagnosticOrder       `bson:"_includedDiagnosticOrderResourcesReferencedByBasedon,omitempty"`
	IncludedProcedureRequestResourcesReferencedByBasedon           *[]ProcedureRequest      `bson:"_includedProcedureRequestResourcesReferencedByBasedon,omitempty"`
	RevIncludedReferralRequestResourcesReferencingBasedon          *[]ReferralRequest       `bson:"_revIncludedReferralRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref     *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                   *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponsereference  *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponsereference,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequestreference   *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequestreference,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource     *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedEncounterResourcesReferencingIncomingreferral       *[]Encounter             `bson:"_revIncludedEncounterResourcesReferencingIncomingreferral,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment        *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                     *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference       *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral   *[]EpisodeOfCare         `bson:"_revIncludedEpisodeOfCareResourcesReferencingIncomingreferral,omitempty"`
	RevIncludedListResourcesReferencingItem                        *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                     *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                    *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingRequest         *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingRequest,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated         *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject    *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestreference *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequestreference,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAction        *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPlan          *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPlan,omitempty"`
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

func (r *ReferralRequestPlusRelatedResources) GetIncludedDiagnosticOrderResourcesReferencedByBasedon() (diagnosticOrders []DiagnosticOrder, err error) {
	if r.IncludedDiagnosticOrderResourcesReferencedByBasedon == nil {
		err = errors.New("Included diagnosticOrders not requested")
	} else {
		diagnosticOrders = *r.IncludedDiagnosticOrderResourcesReferencedByBasedon
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

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if r.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingTtopic
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

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingResponsereference
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingRequestreference
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

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingIncomingreferral() (encounters []Encounter, err error) {
	if r.RevIncludedEncounterResourcesReferencingIncomingreferral == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *r.RevIncludedEncounterResourcesReferencingIncomingreferral
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if r.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *r.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if r.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *r.RevIncludedMessageHeaderResourcesReferencingData
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

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if r.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *r.RevIncludedListResourcesReferencingItem
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if r.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *r.RevIncludedOrderResourcesReferencingDetail
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

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingRequest() (diagnosticReports []DiagnosticReport, err error) {
	if r.RevIncludedDiagnosticReportResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *r.RevIncludedDiagnosticReportResourcesReferencingRequest
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

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if r.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *r.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if r.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *r.RevIncludedClinicalImpressionResourcesReferencingTrigger
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

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPlan() (clinicalImpressions []ClinicalImpression, err error) {
	if r.RevIncludedClinicalImpressionResourcesReferencingPlan == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *r.RevIncludedClinicalImpressionResourcesReferencingPlan
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
	if r.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByRequester {
			rsc := (*r.IncludedPatientResourcesReferencedByRequester)[idx]
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
	if r.IncludedDiagnosticOrderResourcesReferencedByBasedon != nil {
		for idx := range *r.IncludedDiagnosticOrderResourcesReferencedByBasedon {
			rsc := (*r.IncludedDiagnosticOrderResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedProcedureRequestResourcesReferencedByBasedon != nil {
		for idx := range *r.IncludedProcedureRequestResourcesReferencedByBasedon {
			rsc := (*r.IncludedProcedureRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *ReferralRequestPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
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
	if r.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEncounterResourcesReferencingIncomingreferral != nil {
		for idx := range *r.RevIncludedEncounterResourcesReferencingIncomingreferral {
			rsc := (*r.RevIncludedEncounterResourcesReferencingIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *r.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*r.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingData)[idx]
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
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *r.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*r.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticReportResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedDiagnosticReportResourcesReferencingRequest {
			rsc := (*r.RevIncludedDiagnosticReportResourcesReferencingRequest)[idx]
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
	if r.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *r.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*r.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
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
	if r.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByRequester {
			rsc := (*r.IncludedPatientResourcesReferencedByRequester)[idx]
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
	if r.IncludedDiagnosticOrderResourcesReferencedByBasedon != nil {
		for idx := range *r.IncludedDiagnosticOrderResourcesReferencedByBasedon {
			rsc := (*r.IncludedDiagnosticOrderResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedProcedureRequestResourcesReferencedByBasedon != nil {
		for idx := range *r.IncludedProcedureRequestResourcesReferencedByBasedon {
			rsc := (*r.IncludedProcedureRequestResourcesReferencedByBasedon)[idx]
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
	if r.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEncounterResourcesReferencingIncomingreferral != nil {
		for idx := range *r.RevIncludedEncounterResourcesReferencingIncomingreferral {
			rsc := (*r.RevIncludedEncounterResourcesReferencingIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *r.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*r.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingData)[idx]
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
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *r.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*r.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDiagnosticReportResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedDiagnosticReportResourcesReferencingRequest {
			rsc := (*r.RevIncludedDiagnosticReportResourcesReferencingRequest)[idx]
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
	if r.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *r.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*r.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingAction {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingAction)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingPlan {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingPlan)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
