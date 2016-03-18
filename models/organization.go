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

type Organization struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active         *bool                          `bson:"active,omitempty" json:"active,omitempty"`
	Type           *CodeableConcept               `bson:"type,omitempty" json:"type,omitempty"`
	Name           string                         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom        []ContactPoint                 `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address        []Address                      `bson:"address,omitempty" json:"address,omitempty"`
	PartOf         *Reference                     `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Contact        []OrganizationContactComponent `bson:"contact,omitempty" json:"contact,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Organization) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Organization"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Organization), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Organization) GetBSON() (interface{}, error) {
	x.ResourceType = "Organization"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "organization" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type organization Organization

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Organization) UnmarshalJSON(data []byte) (err error) {
	x2 := organization{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Organization(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Organization) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Organization"
	} else if x.ResourceType != "Organization" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Organization, instead received %s", x.ResourceType))
	}
	return nil
}

type OrganizationContactComponent struct {
	BackboneElement `bson:",inline"`
	Purpose         *CodeableConcept `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Name            *HumanName       `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint   `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address         *Address         `bson:"address,omitempty" json:"address,omitempty"`
}

type OrganizationPlus struct {
	Organization                     `bson:",inline"`
	OrganizationPlusRelatedResources `bson:",inline"`
}

type OrganizationPlusRelatedResources struct {
	IncludedOrganizationResourcesReferencedByPartof                   *[]Organization           `bson:"_includedOrganizationResourcesReferencedByPartof,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRequester           *[]ReferralRequest        `bson:"_revIncludedReferralRequestResourcesReferencingRequester,omitempty"`
	RevIncludedReferralRequestResourcesReferencingRecipient           *[]ReferralRequest        `bson:"_revIncludedReferralRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedAccountResourcesReferencingOwner                       *[]Account                `bson:"_revIncludedAccountResourcesReferencingOwner,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                     *[]Account                `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                    *[]Provenance             `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                   *[]Provenance             `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref         *[]DocumentManifest       `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor             *[]DocumentManifest       `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref         *[]DocumentManifest       `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRecipient          *[]DocumentManifest       `bson:"_revIncludedDocumentManifestResourcesReferencingRecipient,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                  *[]CarePlan               `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedCarePlanResourcesReferencingParticipant                *[]CarePlan               `bson:"_revIncludedCarePlanResourcesReferencingParticipant,omitempty"`
	RevIncludedGoalResourcesReferencingSubject                        *[]Goal                   `bson:"_revIncludedGoalResourcesReferencingSubject,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingOrganization          *[]EpisodeOfCare          `bson:"_revIncludedEpisodeOfCareResourcesReferencingOrganization,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingTeammember            *[]EpisodeOfCare          `bson:"_revIncludedEpisodeOfCareResourcesReferencingTeammember,omitempty"`
	RevIncludedMedicationResourcesReferencingManufacturer             *[]Medication             `bson:"_revIncludedMedicationResourcesReferencingManufacturer,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                 *[]Procedure              `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedListResourcesReferencingItem                           *[]List                   `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthenticator     *[]DocumentReference      `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthenticator,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingCustodian         *[]DocumentReference      `bson:"_revIncludedDocumentReferenceResourcesReferencingCustodian,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor            *[]DocumentReference      `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref        *[]DocumentReference      `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedOrderResourcesReferencingSource                        *[]Order                  `bson:"_revIncludedOrderResourcesReferencingSource,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                        *[]Order                  `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedOrderResourcesReferencingTarget                        *[]Order                  `bson:"_revIncludedOrderResourcesReferencingTarget,omitempty"`
	RevIncludedImmunizationResourcesReferencingManufacturer           *[]Immunization           `bson:"_revIncludedImmunizationResourcesReferencingManufacturer,omitempty"`
	RevIncludedDeviceResourcesReferencingOrganization                 *[]Device                 `bson:"_revIncludedDeviceResourcesReferencingOrganization,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingPerformer          *[]ProcedureRequest       `bson:"_revIncludedProcedureRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                        *[]Flag                   `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                         *[]Flag                   `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingSupplier              *[]SupplyRequest          `bson:"_revIncludedSupplyRequestResourcesReferencingSupplier,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingSource                *[]SupplyRequest          `bson:"_revIncludedSupplyRequestResourcesReferencingSource,omitempty"`
	RevIncludedPractitionerResourcesReferencingOrganization           *[]Practitioner           `bson:"_revIncludedPractitionerResourcesReferencingOrganization,omitempty"`
	RevIncludedObservationResourcesReferencingPerformer               *[]Observation            `bson:"_revIncludedObservationResourcesReferencingPerformer,omitempty"`
	RevIncludedPersonResourcesReferencingOrganization                 *[]Person                 `bson:"_revIncludedPersonResourcesReferencingOrganization,omitempty"`
	RevIncludedContractResourcesReferencingActor                      *[]Contract               `bson:"_revIncludedContractResourcesReferencingActor,omitempty"`
	RevIncludedContractResourcesReferencingSigner                     *[]Contract               `bson:"_revIncludedContractResourcesReferencingSigner,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender         *[]CommunicationRequest   `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient      *[]CommunicationRequest   `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                       *[]Basic                  `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedOrganizationResourcesReferencingPartof                 *[]Organization           `bson:"_revIncludedOrganizationResourcesReferencingPartof,omitempty"`
	RevIncludedProcessRequestResourcesReferencingOrganization         *[]ProcessRequest         `bson:"_revIncludedProcessRequestResourcesReferencingOrganization,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingPerformer          *[]DiagnosticReport       `bson:"_revIncludedDiagnosticReportResourcesReferencingPerformer,omitempty"`
	RevIncludedImagingObjectSelectionResourcesReferencingAuthor       *[]ImagingObjectSelection `bson:"_revIncludedImagingObjectSelectionResourcesReferencingAuthor,omitempty"`
	RevIncludedHealthcareServiceResourcesReferencingOrganization      *[]HealthcareService      `bson:"_revIncludedHealthcareServiceResourcesReferencingOrganization,omitempty"`
	RevIncludedAuditEventResourcesReferencingParticipant              *[]AuditEvent             `bson:"_revIncludedAuditEventResourcesReferencingParticipant,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference                *[]AuditEvent             `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender                *[]Communication          `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient             *[]Communication          `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                 *[]Composition            `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAttester                *[]Composition            `bson:"_revIncludedCompositionResourcesReferencingAttester,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                   *[]Composition            `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated            *[]DetectedIssue          `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedPatientResourcesReferencingCareprovider                *[]Patient                `bson:"_revIncludedPatientResourcesReferencingCareprovider,omitempty"`
	RevIncludedPatientResourcesReferencingOrganization                *[]Patient                `bson:"_revIncludedPatientResourcesReferencingOrganization,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment           *[]OrderResponse          `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedOrderResponseResourcesReferencingWho                   *[]OrderResponse          `bson:"_revIncludedOrderResponseResourcesReferencingWho,omitempty"`
	RevIncludedCoverageResourcesReferencingIssuer                     *[]Coverage               `bson:"_revIncludedCoverageResourcesReferencingIssuer,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject       *[]QuestionnaireResponse  `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest             *[]ProcessResponse        `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedProcessResponseResourcesReferencingOrganization        *[]ProcessResponse        `bson:"_revIncludedProcessResponseResourcesReferencingOrganization,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestorganization *[]ProcessResponse        `bson:"_revIncludedProcessResponseResourcesReferencingRequestorganization,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger          *[]ClinicalImpression     `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData                  *[]MessageHeader          `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingReceiver              *[]MessageHeader          `bson:"_revIncludedMessageHeaderResourcesReferencingReceiver,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingResponsible           *[]MessageHeader          `bson:"_revIncludedMessageHeaderResourcesReferencingResponsible,omitempty"`
	RevIncludedLocationResourcesReferencingOrganization               *[]Location               `bson:"_revIncludedLocationResourcesReferencingOrganization,omitempty"`
}

func (o *OrganizationPlusRelatedResources) GetIncludedOrganizationResourceReferencedByPartof() (organization *Organization, err error) {
	if o.IncludedOrganizationResourcesReferencedByPartof == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*o.IncludedOrganizationResourcesReferencedByPartof) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*o.IncludedOrganizationResourcesReferencedByPartof))
	} else if len(*o.IncludedOrganizationResourcesReferencedByPartof) == 1 {
		organization = &(*o.IncludedOrganizationResourcesReferencedByPartof)[0]
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingRequester() (referralRequests []ReferralRequest, err error) {
	if o.RevIncludedReferralRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *o.RevIncludedReferralRequestResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedReferralRequestResourcesReferencingRecipient() (referralRequests []ReferralRequest, err error) {
	if o.RevIncludedReferralRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded referralRequests not requested")
	} else {
		referralRequests = *o.RevIncludedReferralRequestResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAccountResourcesReferencingOwner() (accounts []Account, err error) {
	if o.RevIncludedAccountResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *o.RevIncludedAccountResourcesReferencingOwner
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if o.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *o.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRecipient() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if o.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *o.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingParticipant() (carePlans []CarePlan, err error) {
	if o.RevIncludedCarePlanResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *o.RevIncludedCarePlanResourcesReferencingParticipant
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedGoalResourcesReferencingSubject() (goals []Goal, err error) {
	if o.RevIncludedGoalResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded goals not requested")
	} else {
		goals = *o.RevIncludedGoalResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingOrganization() (episodeOfCares []EpisodeOfCare, err error) {
	if o.RevIncludedEpisodeOfCareResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *o.RevIncludedEpisodeOfCareResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingTeammember() (episodeOfCares []EpisodeOfCare, err error) {
	if o.RevIncludedEpisodeOfCareResourcesReferencingTeammember == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *o.RevIncludedEpisodeOfCareResourcesReferencingTeammember
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationResourcesReferencingManufacturer() (medications []Medication, err error) {
	if o.RevIncludedMedicationResourcesReferencingManufacturer == nil {
		err = errors.New("RevIncluded medications not requested")
	} else {
		medications = *o.RevIncludedMedicationResourcesReferencingManufacturer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if o.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *o.RevIncludedProcedureResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if o.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *o.RevIncludedListResourcesReferencingItem
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthenticator() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingCustodian() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingCustodian == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingCustodian
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedOrderResourcesReferencingSource() (orders []Order, err error) {
	if o.RevIncludedOrderResourcesReferencingSource == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *o.RevIncludedOrderResourcesReferencingSource
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if o.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *o.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedOrderResourcesReferencingTarget() (orders []Order, err error) {
	if o.RevIncludedOrderResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *o.RevIncludedOrderResourcesReferencingTarget
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingManufacturer() (immunizations []Immunization, err error) {
	if o.RevIncludedImmunizationResourcesReferencingManufacturer == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *o.RevIncludedImmunizationResourcesReferencingManufacturer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceResourcesReferencingOrganization() (devices []Device, err error) {
	if o.RevIncludedDeviceResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded devices not requested")
	} else {
		devices = *o.RevIncludedDeviceResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingPerformer() (procedureRequests []ProcedureRequest, err error) {
	if o.RevIncludedProcedureRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *o.RevIncludedProcedureRequestResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if o.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *o.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedFlagResourcesReferencingAuthor() (flags []Flag, err error) {
	if o.RevIncludedFlagResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *o.RevIncludedFlagResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingSupplier() (supplyRequests []SupplyRequest, err error) {
	if o.RevIncludedSupplyRequestResourcesReferencingSupplier == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *o.RevIncludedSupplyRequestResourcesReferencingSupplier
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingSource() (supplyRequests []SupplyRequest, err error) {
	if o.RevIncludedSupplyRequestResourcesReferencingSource == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *o.RevIncludedSupplyRequestResourcesReferencingSource
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPractitionerResourcesReferencingOrganization() (practitioners []Practitioner, err error) {
	if o.RevIncludedPractitionerResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded practitioners not requested")
	} else {
		practitioners = *o.RevIncludedPractitionerResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPerformer() (observations []Observation, err error) {
	if o.RevIncludedObservationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *o.RevIncludedObservationResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPersonResourcesReferencingOrganization() (people []Person, err error) {
	if o.RevIncludedPersonResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *o.RevIncludedPersonResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingActor() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingActor == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingActor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSigner() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingSigner == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingSigner
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if o.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *o.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if o.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *o.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedOrganizationResourcesReferencingPartof() (organizations []Organization, err error) {
	if o.RevIncludedOrganizationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded organizations not requested")
	} else {
		organizations = *o.RevIncludedOrganizationResourcesReferencingPartof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcessRequestResourcesReferencingOrganization() (processRequests []ProcessRequest, err error) {
	if o.RevIncludedProcessRequestResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded processRequests not requested")
	} else {
		processRequests = *o.RevIncludedProcessRequestResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingPerformer() (diagnosticReports []DiagnosticReport, err error) {
	if o.RevIncludedDiagnosticReportResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *o.RevIncludedDiagnosticReportResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedImagingObjectSelectionResourcesReferencingAuthor() (imagingObjectSelections []ImagingObjectSelection, err error) {
	if o.RevIncludedImagingObjectSelectionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded imagingObjectSelections not requested")
	} else {
		imagingObjectSelections = *o.RevIncludedImagingObjectSelectionResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedHealthcareServiceResourcesReferencingOrganization() (healthcareServices []HealthcareService, err error) {
	if o.RevIncludedHealthcareServiceResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded healthcareServices not requested")
	} else {
		healthcareServices = *o.RevIncludedHealthcareServiceResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingParticipant() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingParticipant
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAttester() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingAttester == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingAttester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *o.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPatientResourcesReferencingCareprovider() (patients []Patient, err error) {
	if o.RevIncludedPatientResourcesReferencingCareprovider == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *o.RevIncludedPatientResourcesReferencingCareprovider
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPatientResourcesReferencingOrganization() (patients []Patient, err error) {
	if o.RevIncludedPatientResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *o.RevIncludedPatientResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *o.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingWho() (orderResponses []OrderResponse, err error) {
	if o.RevIncludedOrderResponseResourcesReferencingWho == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *o.RevIncludedOrderResponseResourcesReferencingWho
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingIssuer() (coverages []Coverage, err error) {
	if o.RevIncludedCoverageResourcesReferencingIssuer == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *o.RevIncludedCoverageResourcesReferencingIssuer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingOrganization() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestorganization() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingRequestorganization == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingRequestorganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingReceiver() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingReceiver
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingResponsible() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingResponsible == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingResponsible
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLocationResourcesReferencingOrganization() (locations []Location, err error) {
	if o.RevIncludedLocationResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded locations not requested")
	} else {
		locations = *o.RevIncludedLocationResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedOrganizationResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByPartof {
			rsc := (*o.IncludedOrganizationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedReferralRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedReferralRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedReferralRequestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedReferralRequestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedReferralRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingOwner != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingOwner {
			rsc := (*o.RevIncludedAccountResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*o.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*o.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCarePlanResourcesReferencingParticipant != nil {
		for idx := range *o.RevIncludedCarePlanResourcesReferencingParticipant {
			rsc := (*o.RevIncludedCarePlanResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*o.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEpisodeOfCareResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEpisodeOfCareResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEpisodeOfCareResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEpisodeOfCareResourcesReferencingTeammember != nil {
		for idx := range *o.RevIncludedEpisodeOfCareResourcesReferencingTeammember {
			rsc := (*o.RevIncludedEpisodeOfCareResourcesReferencingTeammember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedMedicationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedMedicationResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*o.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingCustodian != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingCustodian {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingCustodian)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedOrderResourcesReferencingSource {
			rsc := (*o.RevIncludedOrderResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *o.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*o.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedOrderResourcesReferencingTarget {
			rsc := (*o.RevIncludedOrderResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedDeviceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedDeviceResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*o.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*o.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*o.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSupplier != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSupplier {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSource {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPractitionerResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPractitionerResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPractitionerResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*o.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPersonResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPersonResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPersonResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingActor {
			rsc := (*o.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSigner {
			rsc := (*o.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrganizationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedOrganizationResourcesReferencingPartof {
			rsc := (*o.RevIncludedOrganizationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessRequestResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedProcessRequestResourcesReferencingOrganization {
			rsc := (*o.RevIncludedProcessRequestResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			rsc := (*o.RevIncludedImagingObjectSelectionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedHealthcareServiceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedHealthcareServiceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedHealthcareServiceResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingParticipant {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*o.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingCareprovider != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingCareprovider {
			rsc := (*o.RevIncludedPatientResourcesReferencingCareprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPatientResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *o.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*o.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResponseResourcesReferencingWho != nil {
		for idx := range *o.RevIncludedOrderResponseResourcesReferencingWho {
			rsc := (*o.RevIncludedOrderResponseResourcesReferencingWho)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageResourcesReferencingIssuer != nil {
		for idx := range *o.RevIncludedCoverageResourcesReferencingIssuer {
			rsc := (*o.RevIncludedCoverageResourcesReferencingIssuer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingOrganization {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequestorganization != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequestorganization {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequestorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLocationResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedLocationResourcesReferencingOrganization {
			rsc := (*o.RevIncludedLocationResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *OrganizationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedOrganizationResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByPartof {
			rsc := (*o.IncludedOrganizationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedReferralRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedReferralRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedReferralRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedReferralRequestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedReferralRequestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedReferralRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingOwner != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingOwner {
			rsc := (*o.RevIncludedAccountResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*o.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*o.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCarePlanResourcesReferencingParticipant != nil {
		for idx := range *o.RevIncludedCarePlanResourcesReferencingParticipant {
			rsc := (*o.RevIncludedCarePlanResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*o.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEpisodeOfCareResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEpisodeOfCareResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEpisodeOfCareResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEpisodeOfCareResourcesReferencingTeammember != nil {
		for idx := range *o.RevIncludedEpisodeOfCareResourcesReferencingTeammember {
			rsc := (*o.RevIncludedEpisodeOfCareResourcesReferencingTeammember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedMedicationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedMedicationResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*o.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingCustodian != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingCustodian {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingCustodian)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedOrderResourcesReferencingSource {
			rsc := (*o.RevIncludedOrderResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResourcesReferencingDetail != nil {
		for idx := range *o.RevIncludedOrderResourcesReferencingDetail {
			rsc := (*o.RevIncludedOrderResourcesReferencingDetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedOrderResourcesReferencingTarget {
			rsc := (*o.RevIncludedOrderResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedDeviceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedDeviceResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureRequestResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedProcedureRequestResourcesReferencingPerformer {
			rsc := (*o.RevIncludedProcedureRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*o.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*o.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSupplier != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSupplier {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSource {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPractitionerResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPractitionerResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPractitionerResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*o.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPersonResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPersonResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPersonResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingActor != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingActor {
			rsc := (*o.RevIncludedContractResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSigner {
			rsc := (*o.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrganizationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedOrganizationResourcesReferencingPartof {
			rsc := (*o.RevIncludedOrganizationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessRequestResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedProcessRequestResourcesReferencingOrganization {
			rsc := (*o.RevIncludedProcessRequestResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImagingObjectSelectionResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedImagingObjectSelectionResourcesReferencingAuthor {
			rsc := (*o.RevIncludedImagingObjectSelectionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedHealthcareServiceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedHealthcareServiceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedHealthcareServiceResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingParticipant != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingParticipant {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingReference != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingReference {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingReference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*o.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingCareprovider != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingCareprovider {
			rsc := (*o.RevIncludedPatientResourcesReferencingCareprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPatientResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for idx := range *o.RevIncludedOrderResponseResourcesReferencingFulfillment {
			rsc := (*o.RevIncludedOrderResponseResourcesReferencingFulfillment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrderResponseResourcesReferencingWho != nil {
		for idx := range *o.RevIncludedOrderResponseResourcesReferencingWho {
			rsc := (*o.RevIncludedOrderResponseResourcesReferencingWho)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageResourcesReferencingIssuer != nil {
		for idx := range *o.RevIncludedCoverageResourcesReferencingIssuer {
			rsc := (*o.RevIncludedCoverageResourcesReferencingIssuer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingOrganization {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequestorganization != nil {
		for idx := range *o.RevIncludedProcessResponseResourcesReferencingRequestorganization {
			rsc := (*o.RevIncludedProcessResponseResourcesReferencingRequestorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingTrigger)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLocationResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedLocationResourcesReferencingOrganization {
			rsc := (*o.RevIncludedLocationResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
