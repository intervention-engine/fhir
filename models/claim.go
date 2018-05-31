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

type Claim struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                           `bson:"status,omitempty" json:"status,omitempty"`
	Type                 *CodeableConcept                 `bson:"type,omitempty" json:"type,omitempty"`
	SubType              []CodeableConcept                `bson:"subType,omitempty" json:"subType,omitempty"`
	Use                  string                           `bson:"use,omitempty" json:"use,omitempty"`
	Patient              *Reference                       `bson:"patient,omitempty" json:"patient,omitempty"`
	BillablePeriod       *Period                          `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Created              *FHIRDateTime                    `bson:"created,omitempty" json:"created,omitempty"`
	Enterer              *Reference                       `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Insurer              *Reference                       `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider             *Reference                       `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization         *Reference                       `bson:"organization,omitempty" json:"organization,omitempty"`
	Priority             *CodeableConcept                 `bson:"priority,omitempty" json:"priority,omitempty"`
	FundsReserve         *CodeableConcept                 `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	Related              []ClaimRelatedClaimComponent     `bson:"related,omitempty" json:"related,omitempty"`
	Prescription         *Reference                       `bson:"prescription,omitempty" json:"prescription,omitempty"`
	OriginalPrescription *Reference                       `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                *ClaimPayeeComponent             `bson:"payee,omitempty" json:"payee,omitempty"`
	Referral             *Reference                       `bson:"referral,omitempty" json:"referral,omitempty"`
	Facility             *Reference                       `bson:"facility,omitempty" json:"facility,omitempty"`
	CareTeam             []ClaimCareTeamComponent         `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	Information          []ClaimSpecialConditionComponent `bson:"information,omitempty" json:"information,omitempty"`
	Diagnosis            []ClaimDiagnosisComponent        `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure            []ClaimProcedureComponent        `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Insurance            []ClaimInsuranceComponent        `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Accident             *ClaimAccidentComponent          `bson:"accident,omitempty" json:"accident,omitempty"`
	EmploymentImpacted   *Period                          `bson:"employmentImpacted,omitempty" json:"employmentImpacted,omitempty"`
	Hospitalization      *Period                          `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Item                 []ClaimItemComponent             `bson:"item,omitempty" json:"item,omitempty"`
	Total                *Quantity                        `bson:"total,omitempty" json:"total,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Claim) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Claim"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Claim), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Claim) GetBSON() (interface{}, error) {
	x.ResourceType = "Claim"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "claim" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type claim Claim

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Claim) UnmarshalJSON(data []byte) (err error) {
	x2 := claim{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = Claim(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Claim) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Claim"
	} else if x.ResourceType != "Claim" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Claim, instead received %s", x.ResourceType))
	}
	return nil
}

type ClaimRelatedClaimComponent struct {
	BackboneElement `bson:",inline"`
	Claim           *Reference       `bson:"claim,omitempty" json:"claim,omitempty"`
	Relationship    *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference       *Identifier      `bson:"reference,omitempty" json:"reference,omitempty"`
}

type ClaimPayeeComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ResourceType    *Coding          `bson:"resourceType,omitempty" json:"resourceType,omitempty"`
	Party           *Reference       `bson:"party,omitempty" json:"party,omitempty"`
}

type ClaimCareTeamComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32          `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Provider        *Reference       `bson:"provider,omitempty" json:"provider,omitempty"`
	Responsible     *bool            `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Qualification   *CodeableConcept `bson:"qualification,omitempty" json:"qualification,omitempty"`
}

type ClaimSpecialConditionComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32          `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Category        *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	TimingDate      *FHIRDateTime    `bson:"timingDate,omitempty" json:"timingDate,omitempty"`
	TimingPeriod    *Period          `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	ValueString     string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueQuantity   *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueAttachment *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueReference  *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Reason          *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ClaimDiagnosisComponent struct {
	BackboneElement          `bson:",inline"`
	Sequence                 *uint32           `bson:"sequence,omitempty" json:"sequence,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept  `bson:"diagnosisCodeableConcept,omitempty" json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference        `bson:"diagnosisReference,omitempty" json:"diagnosisReference,omitempty"`
	Type                     []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	PackageCode              *CodeableConcept  `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
}

type ClaimProcedureComponent struct {
	BackboneElement          `bson:",inline"`
	Sequence                 *uint32          `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Date                     *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	ProcedureCodeableConcept *CodeableConcept `bson:"procedureCodeableConcept,omitempty" json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *Reference       `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
}

type ClaimInsuranceComponent struct {
	BackboneElement     `bson:",inline"`
	Sequence            *uint32    `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Focal               *bool      `bson:"focal,omitempty" json:"focal,omitempty"`
	Coverage            *Reference `bson:"coverage,omitempty" json:"coverage,omitempty"`
	BusinessArrangement string     `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	PreAuthRef          []string   `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
}

type ClaimAccidentComponent struct {
	BackboneElement   `bson:",inline"`
	Date              *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	LocationAddress   *Address         `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference *Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
}

type ClaimItemComponent struct {
	BackboneElement         `bson:",inline"`
	Sequence                *uint32                `bson:"sequence,omitempty" json:"sequence,omitempty"`
	CareTeamLinkId          []uint32               `bson:"careTeamLinkId,omitempty" json:"careTeamLinkId,omitempty"`
	DiagnosisLinkId         []uint32               `bson:"diagnosisLinkId,omitempty" json:"diagnosisLinkId,omitempty"`
	ProcedureLinkId         []uint32               `bson:"procedureLinkId,omitempty" json:"procedureLinkId,omitempty"`
	InformationLinkId       []uint32               `bson:"informationLinkId,omitempty" json:"informationLinkId,omitempty"`
	Revenue                 *CodeableConcept       `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category                *CodeableConcept       `bson:"category,omitempty" json:"category,omitempty"`
	Service                 *CodeableConcept       `bson:"service,omitempty" json:"service,omitempty"`
	Modifier                []CodeableConcept      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept      `bson:"programCode,omitempty" json:"programCode,omitempty"`
	ServicedDate            *FHIRDateTime          `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept       `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address               `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference       *Reference             `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity                *Quantity              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice               *Quantity              `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                  *float64               `bson:"factor,omitempty" json:"factor,omitempty"`
	Net                     *Quantity              `bson:"net,omitempty" json:"net,omitempty"`
	Udi                     []Reference            `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite                *CodeableConcept       `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept      `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Encounter               []Reference            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Detail                  []ClaimDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ClaimDetailComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32                   `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Revenue         *CodeableConcept          `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *CodeableConcept          `bson:"category,omitempty" json:"category,omitempty"`
	Service         *CodeableConcept          `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []CodeableConcept         `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode     []CodeableConcept         `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity        *Quantity                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64                  `bson:"factor,omitempty" json:"factor,omitempty"`
	Net             *Quantity                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi             []Reference               `bson:"udi,omitempty" json:"udi,omitempty"`
	SubDetail       []ClaimSubDetailComponent `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ClaimSubDetailComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32           `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Revenue         *CodeableConcept  `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Service         *CodeableConcept  `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode     []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity        *Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity         `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64          `bson:"factor,omitempty" json:"factor,omitempty"`
	Net             *Quantity         `bson:"net,omitempty" json:"net,omitempty"`
	Udi             []Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
}

type ClaimPlus struct {
	Claim                     `bson:",inline"`
	ClaimPlusRelatedResources `bson:",inline"`
}

type ClaimPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByCareteam               *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByCareteam,omitempty"`
	IncludedOrganizationResourcesReferencedByCareteam               *[]Organization          `bson:"_includedOrganizationResourcesReferencedByCareteam,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                 *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPractitionerResourcesReferencedByPayee                  *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByPayee,omitempty"`
	IncludedOrganizationResourcesReferencedByPayee                  *[]Organization          `bson:"_includedOrganizationResourcesReferencedByPayee,omitempty"`
	IncludedPatientResourcesReferencedByPayee                       *[]Patient               `bson:"_includedPatientResourcesReferencedByPayee,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPayee                 *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByPayee,omitempty"`
	IncludedPractitionerResourcesReferencedByProvider               *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByProvider,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByInsurer                *[]Organization          `bson:"_includedOrganizationResourcesReferencedByInsurer,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization           *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedPractitionerResourcesReferencedByEnterer                *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByEnterer,omitempty"`
	IncludedLocationResourcesReferencedByFacility                   *[]Location              `bson:"_includedLocationResourcesReferencedByFacility,omitempty"`
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
	RevIncludedExplanationOfBenefitResourcesReferencingClaim        *[]ExplanationOfBenefit  `bson:"_revIncludedExplanationOfBenefitResourcesReferencingClaim,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces         *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon          *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom               *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor               *[]Library               `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                *[]Library               `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon      *[]CommunicationRequest  `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                     *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedClaimResponseResourcesReferencingRequest             *[]ClaimResponse         `bson:"_revIncludedClaimResponseResourcesReferencingRequest,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                 *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail          *[]Condition             `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject               *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                 *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated          *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject     *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest           *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor          *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof         *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerResourceReferencedByCareteam() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByCareteam == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByCareteam) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByCareteam))
	} else if len(*c.IncludedPractitionerResourcesReferencedByCareteam) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByCareteam)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedOrganizationResourceReferencedByCareteam() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByCareteam == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByCareteam) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByCareteam))
	} else if len(*c.IncludedOrganizationResourcesReferencedByCareteam) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByCareteam)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedEncounterResourcesReferencedByEncounter() (encounters []Encounter, err error) {
	if c.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else {
		encounters = *c.IncludedEncounterResourcesReferencedByEncounter
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPayee() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByPayee == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByPayee))
	} else if len(*c.IncludedPractitionerResourcesReferencedByPayee) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByPayee)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedOrganizationResourceReferencedByPayee() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByPayee == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByPayee))
	} else if len(*c.IncludedOrganizationResourcesReferencedByPayee) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByPayee)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPatientResourceReferencedByPayee() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPayee == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPayee))
	} else if len(*c.IncludedPatientResourcesReferencedByPayee) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPayee)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByPayee() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByPayee == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedByPayee))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByPayee) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedByPayee)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerResourceReferencedByProvider() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByProvider == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByProvider))
	} else if len(*c.IncludedPractitionerResourcesReferencedByProvider) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByProvider)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPatient))
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedOrganizationResourceReferencedByInsurer() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByInsurer == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByInsurer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByInsurer))
	} else if len(*c.IncludedOrganizationResourcesReferencedByInsurer) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByInsurer)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*c.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerResourceReferencedByEnterer() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByEnterer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByEnterer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByEnterer))
	} else if len(*c.IncludedPractitionerResourcesReferencedByEnterer) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByEnterer)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedLocationResourceReferencedByFacility() (location *Location, err error) {
	if c.IncludedLocationResourcesReferencedByFacility == nil {
		err = errors.New("Included locations not requested")
	} else if len(*c.IncludedLocationResourcesReferencedByFacility) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*c.IncludedLocationResourcesReferencedByFacility))
	} else if len(*c.IncludedLocationResourcesReferencedByFacility) == 1 {
		location = &(*c.IncludedLocationResourcesReferencedByFacility)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if c.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *c.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingClaim() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if c.RevIncludedExplanationOfBenefitResourcesReferencingClaim == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *c.RevIncludedExplanationOfBenefitResourcesReferencingClaim
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if c.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *c.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if c.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *c.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *c.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedClaimResponseResourcesReferencingRequest() (claimResponses []ClaimResponse, err error) {
	if c.RevIncludedClaimResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded claimResponses not requested")
	} else {
		claimResponses = *c.RevIncludedClaimResponseResourcesReferencingRequest
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if c.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *c.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPractitionerResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByCareteam {
			rsc := (*c.IncludedPractitionerResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByCareteam {
			rsc := (*c.IncludedOrganizationResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*c.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByPayee {
			rsc := (*c.IncludedPractitionerResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPayee {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPayee {
			rsc := (*c.IncludedPatientResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPayee {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*c.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByInsurer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByInsurer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*c.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByEnterer != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByEnterer {
			rsc := (*c.IncludedPractitionerResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByFacility != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByFacility {
			rsc := (*c.IncludedLocationResourcesReferencedByFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*c.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedExplanationOfBenefitResourcesReferencingClaim != nil {
		for idx := range *c.RevIncludedExplanationOfBenefitResourcesReferencingClaim {
			rsc := (*c.RevIncludedExplanationOfBenefitResourcesReferencingClaim)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClaimResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedClaimResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedClaimResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ClaimPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPractitionerResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByCareteam {
			rsc := (*c.IncludedPractitionerResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByCareteam {
			rsc := (*c.IncludedOrganizationResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*c.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByPayee {
			rsc := (*c.IncludedPractitionerResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPayee {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPayee {
			rsc := (*c.IncludedPatientResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPayee {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*c.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByInsurer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByInsurer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*c.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByEnterer != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByEnterer {
			rsc := (*c.IncludedPractitionerResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByFacility != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByFacility {
			rsc := (*c.IncludedLocationResourcesReferencedByFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*c.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedExplanationOfBenefitResourcesReferencingClaim != nil {
		for idx := range *c.RevIncludedExplanationOfBenefitResourcesReferencingClaim {
			rsc := (*c.RevIncludedExplanationOfBenefitResourcesReferencingClaim)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClaimResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedClaimResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedClaimResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
