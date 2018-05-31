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
	ProcessNote          []ExplanationOfBenefitNoteComponent                  `bson:"processNote,omitempty" json:"processNote,omitempty"`
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
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
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
	Party           *Reference       `bson:"party,omitempty" json:"party,omitempty"`
}

type ExplanationOfBenefitSupportingInformationComponent struct {
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
	Encounter               []Reference                                 `bson:"encounter,omitempty" json:"encounter,omitempty"`
	NoteNumber              []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication            []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail                  []ExplanationOfBenefitDetailComponent       `bson:"detail,omitempty" json:"detail,omitempty"`
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
	BackboneElement    `bson:",inline"`
	Type               *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	AllowedUnsignedInt *uint32          `bson:"allowedUnsignedInt,omitempty" json:"allowedUnsignedInt,omitempty"`
	AllowedString      string           `bson:"allowedString,omitempty" json:"allowedString,omitempty"`
	AllowedMoney       *Quantity        `bson:"allowedMoney,omitempty" json:"allowedMoney,omitempty"`
	UsedUnsignedInt    *uint32          `bson:"usedUnsignedInt,omitempty" json:"usedUnsignedInt,omitempty"`
	UsedMoney          *Quantity        `bson:"usedMoney,omitempty" json:"usedMoney,omitempty"`
}

type ExplanationOfBenefitPlus struct {
	ExplanationOfBenefit                     `bson:",inline"`
	ExplanationOfBenefitPlusRelatedResources `bson:",inline"`
}

type ExplanationOfBenefitPlusRelatedResources struct {
	IncludedCoverageResourcesReferencedByCoverage                   *[]Coverage              `bson:"_includedCoverageResourcesReferencedByCoverage,omitempty"`
	IncludedPractitionerResourcesReferencedByCareteam               *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByCareteam,omitempty"`
	IncludedOrganizationResourcesReferencedByCareteam               *[]Organization          `bson:"_includedOrganizationResourcesReferencedByCareteam,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                 *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPractitionerResourcesReferencedByPayee                  *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByPayee,omitempty"`
	IncludedOrganizationResourcesReferencedByPayee                  *[]Organization          `bson:"_includedOrganizationResourcesReferencedByPayee,omitempty"`
	IncludedPatientResourcesReferencedByPayee                       *[]Patient               `bson:"_includedPatientResourcesReferencedByPayee,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPayee                 *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByPayee,omitempty"`
	IncludedPractitionerResourcesReferencedByProvider               *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByProvider,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization           *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedClaimResourcesReferencedByClaim                         *[]Claim                 `bson:"_includedClaimResourcesReferencedByClaim,omitempty"`
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerResourceReferencedByCareteam() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByCareteam == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByCareteam) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByCareteam))
	} else if len(*e.IncludedPractitionerResourcesReferencedByCareteam) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByCareteam)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedOrganizationResourceReferencedByCareteam() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByCareteam == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByCareteam) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByCareteam))
	} else if len(*e.IncludedOrganizationResourcesReferencedByCareteam) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByCareteam)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedEncounterResourcesReferencedByEncounter() (encounters []Encounter, err error) {
	if e.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else {
		encounters = *e.IncludedEncounterResourcesReferencedByEncounter
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPayee() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByPayee == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByPayee))
	} else if len(*e.IncludedPractitionerResourcesReferencedByPayee) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByPayee)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedOrganizationResourceReferencedByPayee() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByPayee == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByPayee))
	} else if len(*e.IncludedOrganizationResourcesReferencedByPayee) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByPayee)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPatientResourceReferencedByPayee() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedByPayee == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedByPayee))
	} else if len(*e.IncludedPatientResourcesReferencedByPayee) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedByPayee)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByPayee() (relatedPerson *RelatedPerson, err error) {
	if e.IncludedRelatedPersonResourcesReferencedByPayee == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*e.IncludedRelatedPersonResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*e.IncludedRelatedPersonResourcesReferencedByPayee))
	} else if len(*e.IncludedRelatedPersonResourcesReferencedByPayee) == 1 {
		relatedPerson = &(*e.IncludedRelatedPersonResourcesReferencedByPayee)[0]
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerResourceReferencedByEnterer() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByEnterer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByEnterer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByEnterer))
	} else if len(*e.IncludedPractitionerResourcesReferencedByEnterer) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByEnterer)[0]
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if e.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *e.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if e.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *e.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDependsonPath2
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingTermtopic
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if e.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *e.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if e.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *e.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if e.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *e.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if e.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *e.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingPartof
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if e.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *e.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingEntityref
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if e.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *e.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if e.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *e.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *e.RevIncludedCommunicationRequestResourcesReferencingBasedon
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if e.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *e.RevIncludedConditionResourcesReferencingEvidencedetail
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
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
	if e.IncludedPractitionerResourcesReferencedByCareteam != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByCareteam {
			rsc := (*e.IncludedPractitionerResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByCareteam != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByCareteam {
			rsc := (*e.IncludedOrganizationResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *e.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*e.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByPayee {
			rsc := (*e.IncludedPractitionerResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByPayee {
			rsc := (*e.IncludedOrganizationResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPayee {
			rsc := (*e.IncludedPatientResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedRelatedPersonResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedRelatedPersonResourcesReferencedByPayee {
			rsc := (*e.IncludedRelatedPersonResourcesReferencedByPayee)[idx]
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
	if e.IncludedPractitionerResourcesReferencedByEnterer != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByEnterer {
			rsc := (*e.IncludedPractitionerResourcesReferencedByEnterer)[idx]
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
	if e.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*e.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*e.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*e.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTermtopic)[idx]
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
	if e.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if e.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*e.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if e.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*e.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
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
	if e.IncludedPractitionerResourcesReferencedByCareteam != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByCareteam {
			rsc := (*e.IncludedPractitionerResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByCareteam != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByCareteam {
			rsc := (*e.IncludedOrganizationResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *e.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*e.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByPayee {
			rsc := (*e.IncludedPractitionerResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByPayee {
			rsc := (*e.IncludedOrganizationResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPayee {
			rsc := (*e.IncludedPatientResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedRelatedPersonResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedRelatedPersonResourcesReferencedByPayee {
			rsc := (*e.IncludedRelatedPersonResourcesReferencedByPayee)[idx]
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
	if e.IncludedPractitionerResourcesReferencedByEnterer != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByEnterer {
			rsc := (*e.IncludedPractitionerResourcesReferencedByEnterer)[idx]
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
	if e.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*e.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*e.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*e.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*e.RevIncludedContractResourcesReferencingTermtopic)[idx]
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
	if e.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*e.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
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
	if e.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*e.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
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
	if e.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*e.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
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
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
