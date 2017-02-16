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

type ExplanationOfBenefit struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                                               `bson:"status,omitempty" json:"status,omitempty"`
	Type                 *CodeableConcept                                     `bson:"type,omitempty" json:"type,omitempty"`
	SubType              []CodeableConcept                                    `bson:"subType,omitempty" json:"subType,omitempty"`
	Patient              *Reference                                           `bson:"patient,omitempty" json:"patient,omitempty"`
	BillablePeriod       *Period                                              `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Created              *FHIRDateTime                                        `bson:"created,omitempty" json:"created,omitempty"`
	Enterer              *Reference                                           `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Insurer              *Reference                                           `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider             *Reference                                           `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization         *Reference                                           `bson:"organization,omitempty" json:"organization,omitempty"`
	Referral             *Reference                                           `bson:"referral,omitempty" json:"referral,omitempty"`
	Facility             *Reference                                           `bson:"facility,omitempty" json:"facility,omitempty"`
	Claim                *Reference                                           `bson:"claim,omitempty" json:"claim,omitempty"`
	ClaimResponse        *Reference                                           `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
	Outcome              *CodeableConcept                                     `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition          string                                               `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Related              []ExplanationOfBenefitRelatedClaimComponent          `bson:"related,omitempty" json:"related,omitempty"`
	Prescription         *Reference                                           `bson:"prescription,omitempty" json:"prescription,omitempty"`
	OriginalPrescription *Reference                                           `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                *ExplanationOfBenefitPayeeComponent                  `bson:"payee,omitempty" json:"payee,omitempty"`
	Information          []ExplanationOfBenefitSupportingInformationComponent `bson:"information,omitempty" json:"information,omitempty"`
	CareTeam             []ExplanationOfBenefitCareTeamComponent              `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	Diagnosis            []ExplanationOfBenefitDiagnosisComponent             `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure            []ExplanationOfBenefitProcedureComponent             `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Precedence           *uint32                                              `bson:"precedence,omitempty" json:"precedence,omitempty"`
	Insurance            *ExplanationOfBenefitInsuranceComponent              `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Accident             *ExplanationOfBenefitAccidentComponent               `bson:"accident,omitempty" json:"accident,omitempty"`
	EmploymentImpacted   *Period                                              `bson:"employmentImpacted,omitempty" json:"employmentImpacted,omitempty"`
	Hospitalization      *Period                                              `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Item                 []ExplanationOfBenefitItemComponent                  `bson:"item,omitempty" json:"item,omitempty"`
	AddItem              []ExplanationOfBenefitAddedItemComponent             `bson:"addItem,omitempty" json:"addItem,omitempty"`
	TotalCost            *Quantity                                            `bson:"totalCost,omitempty" json:"totalCost,omitempty"`
	UnallocDeductable    *Quantity                                            `bson:"unallocDeductable,omitempty" json:"unallocDeductable,omitempty"`
	TotalBenefit         *Quantity                                            `bson:"totalBenefit,omitempty" json:"totalBenefit,omitempty"`
	Payment              *ExplanationOfBenefitPaymentComponent                `bson:"payment,omitempty" json:"payment,omitempty"`
	Form                 *CodeableConcept                                     `bson:"form,omitempty" json:"form,omitempty"`
	Note                 []ExplanationOfBenefitNoteComponent                  `bson:"note,omitempty" json:"note,omitempty"`
	BenefitBalance       []ExplanationOfBenefitBenefitBalanceComponent        `bson:"benefitBalance,omitempty" json:"benefitBalance,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ExplanationOfBenefit) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ExplanationOfBenefit"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ExplanationOfBenefit), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ExplanationOfBenefit) GetBSON() (interface{}, error) {
	x.ResourceType = "ExplanationOfBenefit"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "explanationOfBenefit" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type explanationOfBenefit ExplanationOfBenefit

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ExplanationOfBenefit) UnmarshalJSON(data []byte) (err error) {
	x2 := explanationOfBenefit{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ExplanationOfBenefit(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ExplanationOfBenefit) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ExplanationOfBenefit"
	} else if x.ResourceType != "ExplanationOfBenefit" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ExplanationOfBenefit, instead received %s", x.ResourceType))
	}
	return nil
}

type ExplanationOfBenefitRelatedClaimComponent struct {
	BackboneElement `bson:",inline"`
	Claim           *Reference       `bson:"claim,omitempty" json:"claim,omitempty"`
	Relationship    *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference       *Identifier      `bson:"reference,omitempty" json:"reference,omitempty"`
}

type ExplanationOfBenefitPayeeComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ResourceType    *CodeableConcept `bson:"resourceType,omitempty" json:"resourceType,omitempty"`
	PartyIdentifier *Identifier      `bson:"partyIdentifier,omitempty" json:"partyIdentifier,omitempty"`
	PartyReference  *Reference       `bson:"partyReference,omitempty" json:"partyReference,omitempty"`
}

type ExplanationOfBenefitSupportingInformationComponent struct {
	BackboneElement `bson:",inline"`
	Category        *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	TimingDate      *FHIRDateTime    `bson:"timingDate,omitempty" json:"timingDate,omitempty"`
	TimingPeriod    *Period          `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	ValueString     string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueQuantity   *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueAttachment *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueReference  *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Reason          *Coding          `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ExplanationOfBenefitCareTeamComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32          `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Provider        *Reference       `bson:"provider,omitempty" json:"provider,omitempty"`
	Responsible     *bool            `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Qualification   *CodeableConcept `bson:"qualification,omitempty" json:"qualification,omitempty"`
}

type ExplanationOfBenefitDiagnosisComponent struct {
	BackboneElement          `bson:",inline"`
	Sequence                 *uint32           `bson:"sequence,omitempty" json:"sequence,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept  `bson:"diagnosisCodeableConcept,omitempty" json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference        `bson:"diagnosisReference,omitempty" json:"diagnosisReference,omitempty"`
	Type                     []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	PackageCode              *CodeableConcept  `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
}

type ExplanationOfBenefitProcedureComponent struct {
	BackboneElement          `bson:",inline"`
	Sequence                 *uint32          `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Date                     *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	ProcedureCodeableConcept *CodeableConcept `bson:"procedureCodeableConcept,omitempty" json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *Reference       `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
}

type ExplanationOfBenefitInsuranceComponent struct {
	BackboneElement `bson:",inline"`
	Coverage        *Reference `bson:"coverage,omitempty" json:"coverage,omitempty"`
	PreAuthRef      []string   `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
}

type ExplanationOfBenefitAccidentComponent struct {
	BackboneElement   `bson:",inline"`
	Date              *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	LocationAddress   *Address         `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference *Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
}

type ExplanationOfBenefitItemComponent struct {
	BackboneElement         `bson:",inline"`
	Sequence                *uint32                                     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	CareTeamLinkId          []uint32                                    `bson:"careTeamLinkId,omitempty" json:"careTeamLinkId,omitempty"`
	DiagnosisLinkId         []uint32                                    `bson:"diagnosisLinkId,omitempty" json:"diagnosisLinkId,omitempty"`
	ProcedureLinkId         []uint32                                    `bson:"procedureLinkId,omitempty" json:"procedureLinkId,omitempty"`
	InformationLinkId       []uint32                                    `bson:"informationLinkId,omitempty" json:"informationLinkId,omitempty"`
	Revenue                 *CodeableConcept                            `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category                *CodeableConcept                            `bson:"category,omitempty" json:"category,omitempty"`
	Service                 *CodeableConcept                            `bson:"service,omitempty" json:"service,omitempty"`
	Modifier                []CodeableConcept                           `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept                           `bson:"programCode,omitempty" json:"programCode,omitempty"`
	ServicedDate            *FHIRDateTime                               `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                                     `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                            `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                                    `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference       *Reference                                  `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity                *Quantity                                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice               *Quantity                                   `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                  *float64                                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net                     *Quantity                                   `bson:"net,omitempty" json:"net,omitempty"`
	Udi                     []Reference                                 `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite                *CodeableConcept                            `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept                           `bson:"subSite,omitempty" json:"subSite,omitempty"`
	NoteNumber              []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication            []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail                  []ExplanationOfBenefitDetailComponent       `bson:"detail,omitempty" json:"detail,omitempty"`
	Prosthesis              *ExplanationOfBenefitProsthesisComponent    `bson:"prosthesis,omitempty" json:"prosthesis,omitempty"`
}

type ExplanationOfBenefitAdjudicationComponent struct {
	BackboneElement `bson:",inline"`
	Category        *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Reason          *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount          *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
	Value           *float64         `bson:"value,omitempty" json:"value,omitempty"`
}

type ExplanationOfBenefitDetailComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32                                     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type            *CodeableConcept                            `bson:"type,omitempty" json:"type,omitempty"`
	Revenue         *CodeableConcept                            `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *CodeableConcept                            `bson:"category,omitempty" json:"category,omitempty"`
	Service         *CodeableConcept                            `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []CodeableConcept                           `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode     []CodeableConcept                           `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity        *Quantity                                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity                                   `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64                                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net             *Quantity                                   `bson:"net,omitempty" json:"net,omitempty"`
	Udi             []Reference                                 `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber      []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail       []ExplanationOfBenefitSubDetailComponent    `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ExplanationOfBenefitSubDetailComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32                                     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type            *CodeableConcept                            `bson:"type,omitempty" json:"type,omitempty"`
	Revenue         *CodeableConcept                            `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *CodeableConcept                            `bson:"category,omitempty" json:"category,omitempty"`
	Service         *CodeableConcept                            `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []CodeableConcept                           `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode     []CodeableConcept                           `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity        *Quantity                                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity                                   `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64                                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net             *Quantity                                   `bson:"net,omitempty" json:"net,omitempty"`
	Udi             []Reference                                 `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber      []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ExplanationOfBenefitProsthesisComponent struct {
	BackboneElement `bson:",inline"`
	Initial         *bool            `bson:"initial,omitempty" json:"initial,omitempty"`
	PriorDate       *FHIRDateTime    `bson:"priorDate,omitempty" json:"priorDate,omitempty"`
	PriorMaterial   *CodeableConcept `bson:"priorMaterial,omitempty" json:"priorMaterial,omitempty"`
}

type ExplanationOfBenefitAddedItemComponent struct {
	BackboneElement `bson:",inline"`
	SequenceLinkId  []uint32                                        `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Revenue         *CodeableConcept                                `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *CodeableConcept                                `bson:"category,omitempty" json:"category,omitempty"`
	Service         *CodeableConcept                                `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []CodeableConcept                               `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee             *Quantity                                       `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber      []uint32                                        `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ExplanationOfBenefitAdjudicationComponent     `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail          []ExplanationOfBenefitAddedItemsDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ExplanationOfBenefitAddedItemsDetailComponent struct {
	BackboneElement `bson:",inline"`
	Revenue         *CodeableConcept                            `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *CodeableConcept                            `bson:"category,omitempty" json:"category,omitempty"`
	Service         *CodeableConcept                            `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []CodeableConcept                           `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee             *Quantity                                   `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber      []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ExplanationOfBenefitPaymentComponent struct {
	BackboneElement  `bson:",inline"`
	Type             *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Adjustment       *Quantity        `bson:"adjustment,omitempty" json:"adjustment,omitempty"`
	AdjustmentReason *CodeableConcept `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty"`
	Date             *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Amount           *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
	Identifier       *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
}

type ExplanationOfBenefitNoteComponent struct {
	BackboneElement `bson:",inline"`
	Number          *uint32          `bson:"number,omitempty" json:"number,omitempty"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Text            string           `bson:"text,omitempty" json:"text,omitempty"`
	Language        *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}

type ExplanationOfBenefitBenefitBalanceComponent struct {
	BackboneElement `bson:",inline"`
	Category        *CodeableConcept                       `bson:"category,omitempty" json:"category,omitempty"`
	SubCategory     *CodeableConcept                       `bson:"subCategory,omitempty" json:"subCategory,omitempty"`
	Excluded        *bool                                  `bson:"excluded,omitempty" json:"excluded,omitempty"`
	Name            string                                 `bson:"name,omitempty" json:"name,omitempty"`
	Description     string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Network         *CodeableConcept                       `bson:"network,omitempty" json:"network,omitempty"`
	Unit            *CodeableConcept                       `bson:"unit,omitempty" json:"unit,omitempty"`
	Term            *CodeableConcept                       `bson:"term,omitempty" json:"term,omitempty"`
	Financial       []ExplanationOfBenefitBenefitComponent `bson:"financial,omitempty" json:"financial,omitempty"`
}

type ExplanationOfBenefitBenefitComponent struct {
	BackboneElement        `bson:",inline"`
	Type                   *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	BenefitUnsignedInt     *uint32          `bson:"benefitUnsignedInt,omitempty" json:"benefitUnsignedInt,omitempty"`
	BenefitString          string           `bson:"benefitString,omitempty" json:"benefitString,omitempty"`
	BenefitMoney           *Quantity        `bson:"benefitMoney,omitempty" json:"benefitMoney,omitempty"`
	BenefitUsedUnsignedInt *uint32          `bson:"benefitUsedUnsignedInt,omitempty" json:"benefitUsedUnsignedInt,omitempty"`
	BenefitUsedMoney       *Quantity        `bson:"benefitUsedMoney,omitempty" json:"benefitUsedMoney,omitempty"`
}

type ExplanationOfBenefitPlus struct {
	ExplanationOfBenefit                     `bson:",inline"`
	ExplanationOfBenefitPlusRelatedResources `bson:",inline"`
}

type ExplanationOfBenefitPlusRelatedResources struct {
	IncludedCoverageResourcesReferencedByCoverage               *[]Coverage              `bson:"_includedCoverageResourcesReferencedByCoverage,omitempty"`
	IncludedPractitionerResourcesReferencedByProvider           *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByProvider,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization       *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedClaimResourcesReferencedByClaim                     *[]Claim                 `bson:"_includedClaimResourcesReferencedByClaim,omitempty"`
	IncludedLocationResourcesReferencedByFacility               *[]Location              `bson:"_includedLocationResourcesReferencedByFacility,omitempty"`
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedCoverageResourceReferencedByCoverage() (coverage *Coverage, err error) {
	if e.IncludedCoverageResourcesReferencedByCoverage == nil {
		err = errors.New("Included coverages not requested")
	} else if len(*e.IncludedCoverageResourcesReferencedByCoverage) > 1 {
		err = fmt.Errorf("Expected 0 or 1 coverage, but found %d", len(*e.IncludedCoverageResourcesReferencedByCoverage))
	} else if len(*e.IncludedCoverageResourcesReferencedByCoverage) == 1 {
		coverage = &(*e.IncludedCoverageResourcesReferencedByCoverage)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerResourceReferencedByProvider() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByProvider == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByProvider))
	} else if len(*e.IncludedPractitionerResourcesReferencedByProvider) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByProvider)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedByPatient))
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedClaimResourceReferencedByClaim() (claim *Claim, err error) {
	if e.IncludedClaimResourcesReferencedByClaim == nil {
		err = errors.New("Included claims not requested")
	} else if len(*e.IncludedClaimResourcesReferencedByClaim) > 1 {
		err = fmt.Errorf("Expected 0 or 1 claim, but found %d", len(*e.IncludedClaimResourcesReferencedByClaim))
	} else if len(*e.IncludedClaimResourcesReferencedByClaim) == 1 {
		claim = &(*e.IncludedClaimResourcesReferencedByClaim)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedLocationResourceReferencedByFacility() (location *Location, err error) {
	if e.IncludedLocationResourcesReferencedByFacility == nil {
		err = errors.New("Included locations not requested")
	} else if len(*e.IncludedLocationResourcesReferencedByFacility) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*e.IncludedLocationResourcesReferencedByFacility))
	} else if len(*e.IncludedLocationResourcesReferencedByFacility) == 1 {
		location = &(*e.IncludedLocationResourcesReferencedByFacility)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if e.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *e.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedContractResourcesReferencingTtopic() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingTtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingTtopic
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedContractResourcesReferencingTopic() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingTopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingTopic
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if e.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *e.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingItem
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingReplaces() (diagnosticRequests []DiagnosticRequest, err error) {
	if e.RevIncludedDiagnosticRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *e.RevIncludedDiagnosticRequestResourcesReferencingReplaces
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingBasedon() (diagnosticRequests []DiagnosticRequest, err error) {
	if e.RevIncludedDiagnosticRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *e.RevIncludedDiagnosticRequestResourcesReferencingBasedon
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDiagnosticRequestResourcesReferencingDefinition() (diagnosticRequests []DiagnosticRequest, err error) {
	if e.RevIncludedDiagnosticRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded diagnosticRequests not requested")
	} else {
		diagnosticRequests = *e.RevIncludedDiagnosticRequestResourcesReferencingDefinition
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingReplaces() (deviceUseRequests []DeviceUseRequest, err error) {
	if e.RevIncludedDeviceUseRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *e.RevIncludedDeviceUseRequestResourcesReferencingReplaces
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingBasedon() (deviceUseRequests []DeviceUseRequest, err error) {
	if e.RevIncludedDeviceUseRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *e.RevIncludedDeviceUseRequestResourcesReferencingBasedon
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDeviceUseRequestResourcesReferencingDefinition() (deviceUseRequests []DeviceUseRequest, err error) {
	if e.RevIncludedDeviceUseRequestResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded deviceUseRequests not requested")
	} else {
		deviceUseRequests = *e.RevIncludedDeviceUseRequestResourcesReferencingDefinition
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if e.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *e.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *e.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if e.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *e.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedCoverageResourcesReferencedByCoverage != nil {
		for idx := range *e.IncludedCoverageResourcesReferencedByCoverage {
			rsc := (*e.IncludedCoverageResourcesReferencedByCoverage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*e.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatient {
			rsc := (*e.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*e.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedClaimResourcesReferencedByClaim != nil {
		for idx := range *e.IncludedClaimResourcesReferencedByClaim {
			rsc := (*e.IncludedClaimResourcesReferencedByClaim)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedLocationResourcesReferencedByFacility != nil {
		for idx := range *e.IncludedLocationResourcesReferencedByFacility {
			rsc := (*e.IncludedLocationResourcesReferencedByFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingData {
			rsc := (*e.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*e.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*e.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*e.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*e.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*e.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedCoverageResourcesReferencedByCoverage != nil {
		for idx := range *e.IncludedCoverageResourcesReferencedByCoverage {
			rsc := (*e.IncludedCoverageResourcesReferencedByCoverage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*e.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatient {
			rsc := (*e.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*e.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedClaimResourcesReferencedByClaim != nil {
		for idx := range *e.IncludedClaimResourcesReferencedByClaim {
			rsc := (*e.IncludedClaimResourcesReferencedByClaim)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedLocationResourcesReferencedByFacility != nil {
		for idx := range *e.IncludedLocationResourcesReferencedByFacility {
			rsc := (*e.IncludedLocationResourcesReferencedByFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingData {
			rsc := (*e.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTtopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTtopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*e.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*e.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDiagnosticRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDiagnosticRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceUseRequestResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedDeviceUseRequestResourcesReferencingDefinition {
			rsc := (*e.RevIncludedDeviceUseRequestResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*e.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*e.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*e.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
