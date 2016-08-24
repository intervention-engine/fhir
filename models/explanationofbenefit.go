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

type ExplanationOfBenefit struct {
	DomainResource                 `bson:",inline"`
	Identifier                     []Identifier                                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                         string                                          `bson:"status,omitempty" json:"status,omitempty"`
	AuthorIdentifier               *Identifier                                     `bson:"authorIdentifier,omitempty" json:"authorIdentifier,omitempty"`
	AuthorReference                *Reference                                      `bson:"authorReference,omitempty" json:"authorReference,omitempty"`
	ClaimIdentifier                *Identifier                                     `bson:"claimIdentifier,omitempty" json:"claimIdentifier,omitempty"`
	ClaimReference                 *Reference                                      `bson:"claimReference,omitempty" json:"claimReference,omitempty"`
	ClaimResponseIdentifier        *Identifier                                     `bson:"claimResponseIdentifier,omitempty" json:"claimResponseIdentifier,omitempty"`
	ClaimResponseReference         *Reference                                      `bson:"claimResponseReference,omitempty" json:"claimResponseReference,omitempty"`
	Type                           *Coding                                         `bson:"type,omitempty" json:"type,omitempty"`
	SubType                        []Coding                                        `bson:"subType,omitempty" json:"subType,omitempty"`
	Ruleset                        *Coding                                         `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset                *Coding                                         `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created                        *FHIRDateTime                                   `bson:"created,omitempty" json:"created,omitempty"`
	BillablePeriod                 *Period                                         `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Outcome                        *Coding                                         `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition                    string                                          `bson:"disposition,omitempty" json:"disposition,omitempty"`
	ProviderIdentifier             *Identifier                                     `bson:"providerIdentifier,omitempty" json:"providerIdentifier,omitempty"`
	ProviderReference              *Reference                                      `bson:"providerReference,omitempty" json:"providerReference,omitempty"`
	OrganizationIdentifier         *Identifier                                     `bson:"organizationIdentifier,omitempty" json:"organizationIdentifier,omitempty"`
	OrganizationReference          *Reference                                      `bson:"organizationReference,omitempty" json:"organizationReference,omitempty"`
	FacilityIdentifier             *Identifier                                     `bson:"facilityIdentifier,omitempty" json:"facilityIdentifier,omitempty"`
	FacilityReference              *Reference                                      `bson:"facilityReference,omitempty" json:"facilityReference,omitempty"`
	Related                        []ExplanationOfBenefitRelatedClaimComponent     `bson:"related,omitempty" json:"related,omitempty"`
	PrescriptionIdentifier         *Identifier                                     `bson:"prescriptionIdentifier,omitempty" json:"prescriptionIdentifier,omitempty"`
	PrescriptionReference          *Reference                                      `bson:"prescriptionReference,omitempty" json:"prescriptionReference,omitempty"`
	OriginalPrescriptionIdentifier *Identifier                                     `bson:"originalPrescriptionIdentifier,omitempty" json:"originalPrescriptionIdentifier,omitempty"`
	OriginalPrescriptionReference  *Reference                                      `bson:"originalPrescriptionReference,omitempty" json:"originalPrescriptionReference,omitempty"`
	Payee                          *ExplanationOfBenefitPayeeComponent             `bson:"payee,omitempty" json:"payee,omitempty"`
	ReferralIdentifier             *Identifier                                     `bson:"referralIdentifier,omitempty" json:"referralIdentifier,omitempty"`
	ReferralReference              *Reference                                      `bson:"referralReference,omitempty" json:"referralReference,omitempty"`
	Information                    []ExplanationOfBenefitSpecialConditionComponent `bson:"information,omitempty" json:"information,omitempty"`
	Diagnosis                      []ExplanationOfBenefitDiagnosisComponent        `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure                      []ExplanationOfBenefitProcedureComponent        `bson:"procedure,omitempty" json:"procedure,omitempty"`
	PatientIdentifier              *Identifier                                     `bson:"patientIdentifier,omitempty" json:"patientIdentifier,omitempty"`
	PatientReference               *Reference                                      `bson:"patientReference,omitempty" json:"patientReference,omitempty"`
	Precedence                     *uint32                                         `bson:"precedence,omitempty" json:"precedence,omitempty"`
	Coverage                       *ExplanationOfBenefitCoverageComponent          `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Accident                       *ExplanationOfBenefitAccidentComponent          `bson:"accident,omitempty" json:"accident,omitempty"`
	EmploymentImpacted             *Period                                         `bson:"employmentImpacted,omitempty" json:"employmentImpacted,omitempty"`
	Hospitalization                *Period                                         `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Item                           []ExplanationOfBenefitItemComponent             `bson:"item,omitempty" json:"item,omitempty"`
	AddItem                        []ExplanationOfBenefitAddedItemComponent        `bson:"addItem,omitempty" json:"addItem,omitempty"`
	MissingTeeth                   []ExplanationOfBenefitMissingTeethComponent     `bson:"missingTeeth,omitempty" json:"missingTeeth,omitempty"`
	TotalCost                      *Quantity                                       `bson:"totalCost,omitempty" json:"totalCost,omitempty"`
	UnallocDeductable              *Quantity                                       `bson:"unallocDeductable,omitempty" json:"unallocDeductable,omitempty"`
	TotalBenefit                   *Quantity                                       `bson:"totalBenefit,omitempty" json:"totalBenefit,omitempty"`
	Payment                        *ExplanationOfBenefitPaymentComponent           `bson:"payment,omitempty" json:"payment,omitempty"`
	Form                           *Coding                                         `bson:"form,omitempty" json:"form,omitempty"`
	Note                           []ExplanationOfBenefitNoteComponent             `bson:"note,omitempty" json:"note,omitempty"`
	BenefitBalance                 []ExplanationOfBenefitBenefitBalanceComponent   `bson:"benefitBalance,omitempty" json:"benefitBalance,omitempty"`
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
	ClaimIdentifier *Identifier `bson:"claimIdentifier,omitempty" json:"claimIdentifier,omitempty"`
	ClaimReference  *Reference  `bson:"claimReference,omitempty" json:"claimReference,omitempty"`
	Relationship    *Coding     `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference       *Identifier `bson:"reference,omitempty" json:"reference,omitempty"`
}

type ExplanationOfBenefitPayeeComponent struct {
	BackboneElement `bson:",inline"`
	Type            *Coding     `bson:"type,omitempty" json:"type,omitempty"`
	ResourceType    *Coding     `bson:"resourceType,omitempty" json:"resourceType,omitempty"`
	PartyIdentifier *Identifier `bson:"partyIdentifier,omitempty" json:"partyIdentifier,omitempty"`
	PartyReference  *Reference  `bson:"partyReference,omitempty" json:"partyReference,omitempty"`
}

type ExplanationOfBenefitSpecialConditionComponent struct {
	BackboneElement `bson:",inline"`
	Category        *Coding       `bson:"category,omitempty" json:"category,omitempty"`
	Code            *Coding       `bson:"code,omitempty" json:"code,omitempty"`
	TimingDate      *FHIRDateTime `bson:"timingDate,omitempty" json:"timingDate,omitempty"`
	TimingPeriod    *Period       `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	ValueString     string        `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueQuantity   *Quantity     `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
}

type ExplanationOfBenefitDiagnosisComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32  `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Diagnosis       *Coding  `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Type            []Coding `bson:"type,omitempty" json:"type,omitempty"`
	Drg             *Coding  `bson:"drg,omitempty" json:"drg,omitempty"`
}

type ExplanationOfBenefitProcedureComponent struct {
	BackboneElement    `bson:",inline"`
	Sequence           *uint32       `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Date               *FHIRDateTime `bson:"date,omitempty" json:"date,omitempty"`
	ProcedureCoding    *Coding       `bson:"procedureCoding,omitempty" json:"procedureCoding,omitempty"`
	ProcedureReference *Reference    `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
}

type ExplanationOfBenefitCoverageComponent struct {
	BackboneElement    `bson:",inline"`
	CoverageIdentifier *Identifier `bson:"coverageIdentifier,omitempty" json:"coverageIdentifier,omitempty"`
	CoverageReference  *Reference  `bson:"coverageReference,omitempty" json:"coverageReference,omitempty"`
	PreAuthRef         []string    `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
}

type ExplanationOfBenefitAccidentComponent struct {
	BackboneElement   `bson:",inline"`
	Date              *FHIRDateTime `bson:"date,omitempty" json:"date,omitempty"`
	Type              *Coding       `bson:"type,omitempty" json:"type,omitempty"`
	LocationAddress   *Address      `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference *Reference    `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
}

type ExplanationOfBenefitItemComponent struct {
	BackboneElement   `bson:",inline"`
	Sequence          *uint32                                     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	CareTeam          []ExplanationOfBenefitCareTeamComponent     `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	DiagnosisLinkId   []uint32                                    `bson:"diagnosisLinkId,omitempty" json:"diagnosisLinkId,omitempty"`
	Revenue           *Coding                                     `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *Coding                                     `bson:"category,omitempty" json:"category,omitempty"`
	Service           *Coding                                     `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []Coding                                    `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []Coding                                    `bson:"programCode,omitempty" json:"programCode,omitempty"`
	ServicedDate      *FHIRDateTime                               `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod    *Period                                     `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCoding    *Coding                                     `bson:"locationCoding,omitempty" json:"locationCoding,omitempty"`
	LocationAddress   *Address                                    `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference *Reference                                  `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity          *Quantity                                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Quantity                                   `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *float64                                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Points            *float64                                    `bson:"points,omitempty" json:"points,omitempty"`
	Net               *Quantity                                   `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []Reference                                 `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite          *Coding                                     `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite           []Coding                                    `bson:"subSite,omitempty" json:"subSite,omitempty"`
	NoteNumber        []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail            []ExplanationOfBenefitDetailComponent       `bson:"detail,omitempty" json:"detail,omitempty"`
	Prosthesis        *ExplanationOfBenefitProsthesisComponent    `bson:"prosthesis,omitempty" json:"prosthesis,omitempty"`
}

type ExplanationOfBenefitCareTeamComponent struct {
	BackboneElement    `bson:",inline"`
	ProviderIdentifier *Identifier `bson:"providerIdentifier,omitempty" json:"providerIdentifier,omitempty"`
	ProviderReference  *Reference  `bson:"providerReference,omitempty" json:"providerReference,omitempty"`
	Responsible        *bool       `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Role               *Coding     `bson:"role,omitempty" json:"role,omitempty"`
	Qualification      *Coding     `bson:"qualification,omitempty" json:"qualification,omitempty"`
}

type ExplanationOfBenefitAdjudicationComponent struct {
	BackboneElement `bson:",inline"`
	Category        *Coding   `bson:"category,omitempty" json:"category,omitempty"`
	Reason          *Coding   `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount          *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value           *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ExplanationOfBenefitDetailComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32                                     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type            *Coding                                     `bson:"type,omitempty" json:"type,omitempty"`
	Revenue         *Coding                                     `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *Coding                                     `bson:"category,omitempty" json:"category,omitempty"`
	Service         *Coding                                     `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []Coding                                    `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode     []Coding                                    `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity        *Quantity                                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity                                   `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64                                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Points          *float64                                    `bson:"points,omitempty" json:"points,omitempty"`
	Net             *Quantity                                   `bson:"net,omitempty" json:"net,omitempty"`
	Udi             []Reference                                 `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber      []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail       []ExplanationOfBenefitSubDetailComponent    `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ExplanationOfBenefitSubDetailComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32                                     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type            *Coding                                     `bson:"type,omitempty" json:"type,omitempty"`
	Revenue         *Coding                                     `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *Coding                                     `bson:"category,omitempty" json:"category,omitempty"`
	Service         *Coding                                     `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []Coding                                    `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode     []Coding                                    `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity        *Quantity                                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity                                   `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64                                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Points          *float64                                    `bson:"points,omitempty" json:"points,omitempty"`
	Net             *Quantity                                   `bson:"net,omitempty" json:"net,omitempty"`
	Udi             []Reference                                 `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber      []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ExplanationOfBenefitProsthesisComponent struct {
	BackboneElement `bson:",inline"`
	Initial         *bool         `bson:"initial,omitempty" json:"initial,omitempty"`
	PriorDate       *FHIRDateTime `bson:"priorDate,omitempty" json:"priorDate,omitempty"`
	PriorMaterial   *Coding       `bson:"priorMaterial,omitempty" json:"priorMaterial,omitempty"`
}

type ExplanationOfBenefitAddedItemComponent struct {
	BackboneElement `bson:",inline"`
	SequenceLinkId  []uint32                                        `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Revenue         *Coding                                         `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *Coding                                         `bson:"category,omitempty" json:"category,omitempty"`
	Service         *Coding                                         `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []Coding                                        `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee             *Quantity                                       `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber      []uint32                                        `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ExplanationOfBenefitAdjudicationComponent     `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail          []ExplanationOfBenefitAddedItemsDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ExplanationOfBenefitAddedItemsDetailComponent struct {
	BackboneElement `bson:",inline"`
	Revenue         *Coding                                     `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category        *Coding                                     `bson:"category,omitempty" json:"category,omitempty"`
	Service         *Coding                                     `bson:"service,omitempty" json:"service,omitempty"`
	Modifier        []Coding                                    `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee             *Quantity                                   `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber      []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ExplanationOfBenefitMissingTeethComponent struct {
	BackboneElement `bson:",inline"`
	Tooth           *Coding       `bson:"tooth,omitempty" json:"tooth,omitempty"`
	Reason          *Coding       `bson:"reason,omitempty" json:"reason,omitempty"`
	ExtractionDate  *FHIRDateTime `bson:"extractionDate,omitempty" json:"extractionDate,omitempty"`
}

type ExplanationOfBenefitPaymentComponent struct {
	BackboneElement  `bson:",inline"`
	Type             *Coding       `bson:"type,omitempty" json:"type,omitempty"`
	Adjustment       *Quantity     `bson:"adjustment,omitempty" json:"adjustment,omitempty"`
	AdjustmentReason *Coding       `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty"`
	Date             *FHIRDateTime `bson:"date,omitempty" json:"date,omitempty"`
	Amount           *Quantity     `bson:"amount,omitempty" json:"amount,omitempty"`
	Identifier       *Identifier   `bson:"identifier,omitempty" json:"identifier,omitempty"`
}

type ExplanationOfBenefitNoteComponent struct {
	BackboneElement `bson:",inline"`
	Number          *uint32 `bson:"number,omitempty" json:"number,omitempty"`
	Type            *Coding `bson:"type,omitempty" json:"type,omitempty"`
	Text            string  `bson:"text,omitempty" json:"text,omitempty"`
	Language        *Coding `bson:"language,omitempty" json:"language,omitempty"`
}

type ExplanationOfBenefitBenefitBalanceComponent struct {
	BackboneElement `bson:",inline"`
	Category        *Coding                                `bson:"category,omitempty" json:"category,omitempty"`
	SubCategory     *Coding                                `bson:"subCategory,omitempty" json:"subCategory,omitempty"`
	Name            string                                 `bson:"name,omitempty" json:"name,omitempty"`
	Description     string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Network         *Coding                                `bson:"network,omitempty" json:"network,omitempty"`
	Unit            *Coding                                `bson:"unit,omitempty" json:"unit,omitempty"`
	Term            *Coding                                `bson:"term,omitempty" json:"term,omitempty"`
	Financial       []ExplanationOfBenefitBenefitComponent `bson:"financial,omitempty" json:"financial,omitempty"`
}

type ExplanationOfBenefitBenefitComponent struct {
	BackboneElement        `bson:",inline"`
	Type                   *Coding   `bson:"type,omitempty" json:"type,omitempty"`
	BenefitUnsignedInt     *uint32   `bson:"benefitUnsignedInt,omitempty" json:"benefitUnsignedInt,omitempty"`
	BenefitString          string    `bson:"benefitString,omitempty" json:"benefitString,omitempty"`
	BenefitMoney           *Quantity `bson:"benefitMoney,omitempty" json:"benefitMoney,omitempty"`
	BenefitUsedUnsignedInt *uint32   `bson:"benefitUsedUnsignedInt,omitempty" json:"benefitUsedUnsignedInt,omitempty"`
	BenefitUsedMoney       *Quantity `bson:"benefitUsedMoney,omitempty" json:"benefitUsedMoney,omitempty"`
}

type ExplanationOfBenefitPlus struct {
	ExplanationOfBenefit                     `bson:",inline"`
	ExplanationOfBenefitPlusRelatedResources `bson:",inline"`
}

type ExplanationOfBenefitPlusRelatedResources struct {
	IncludedClaimResourcesReferencedByClaimreference               *[]Claim                 `bson:"_includedClaimResourcesReferencedByClaimreference,omitempty"`
	IncludedPatientResourcesReferencedByPatientreference           *[]Patient               `bson:"_includedPatientResourcesReferencedByPatientreference,omitempty"`
	IncludedPractitionerResourcesReferencedByProviderreference     *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByProviderreference,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganizationreference *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganizationreference,omitempty"`
	IncludedLocationResourcesReferencedByFacilityreference         *[]Location              `bson:"_includedLocationResourcesReferencedByFacilityreference,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref      *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                     *[]Consent               `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref     *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingTtopic                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingTtopic,omitempty"`
	RevIncludedContractResourcesReferencingSubject                 *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTopic                   *[]Contract              `bson:"_revIncludedContractResourcesReferencingTopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequestreference   *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequestreference,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponsereference  *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponsereference,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource     *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon            *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                       *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedListResourcesReferencingItem                        *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingReplaces       *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingBasedon        *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDiagnosticRequestResourcesReferencingDefinition     *[]DiagnosticRequest     `bson:"_revIncludedDiagnosticRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingReplaces        *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingBasedon         *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceUseRequestResourcesReferencingDefinition      *[]DeviceUseRequest      `bson:"_revIncludedDeviceUseRequestResourcesReferencingDefinition,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                    *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated         *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject    *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequestreference *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequestreference,omitempty"`
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedClaimResourceReferencedByClaimreference() (claim *Claim, err error) {
	if e.IncludedClaimResourcesReferencedByClaimreference == nil {
		err = errors.New("Included claims not requested")
	} else if len(*e.IncludedClaimResourcesReferencedByClaimreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 claim, but found %d", len(*e.IncludedClaimResourcesReferencedByClaimreference))
	} else if len(*e.IncludedClaimResourcesReferencedByClaimreference) == 1 {
		claim = &(*e.IncludedClaimResourcesReferencedByClaimreference)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPatientResourceReferencedByPatientreference() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedByPatientreference == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedByPatientreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedByPatientreference))
	} else if len(*e.IncludedPatientResourcesReferencedByPatientreference) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedByPatientreference)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerResourceReferencedByProviderreference() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByProviderreference == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByProviderreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByProviderreference))
	} else if len(*e.IncludedPractitionerResourcesReferencedByProviderreference) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByProviderreference)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganizationreference() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByOrganizationreference == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganizationreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByOrganizationreference))
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganizationreference) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByOrganizationreference)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedLocationResourceReferencedByFacilityreference() (location *Location, err error) {
	if e.IncludedLocationResourcesReferencedByFacilityreference == nil {
		err = errors.New("Included locations not requested")
	} else if len(*e.IncludedLocationResourcesReferencedByFacilityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*e.IncludedLocationResourcesReferencedByFacilityreference))
	} else if len(*e.IncludedLocationResourcesReferencedByFacilityreference) == 1 {
		location = &(*e.IncludedLocationResourcesReferencedByFacilityreference)[0]
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequestreference() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequestreference
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponsereference() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponsereference == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponsereference
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequestreference() (processResponses []ProcessResponse, err error) {
	if e.RevIncludedProcessResponseResourcesReferencingRequestreference == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *e.RevIncludedProcessResponseResourcesReferencingRequestreference
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedClaimResourcesReferencedByClaimreference != nil {
		for idx := range *e.IncludedClaimResourcesReferencedByClaimreference {
			rsc := (*e.IncludedClaimResourcesReferencedByClaimreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatientreference != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatientreference {
			rsc := (*e.IncludedPatientResourcesReferencedByPatientreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByProviderreference != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByProviderreference {
			rsc := (*e.IncludedPractitionerResourcesReferencedByProviderreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganizationreference != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByOrganizationreference {
			rsc := (*e.IncludedOrganizationResourcesReferencedByOrganizationreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedLocationResourcesReferencedByFacilityreference != nil {
		for idx := range *e.IncludedLocationResourcesReferencedByFacilityreference {
			rsc := (*e.IncludedLocationResourcesReferencedByFacilityreference)[idx]
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
	if e.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
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
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
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
	if e.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *e.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*e.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedClaimResourcesReferencedByClaimreference != nil {
		for idx := range *e.IncludedClaimResourcesReferencedByClaimreference {
			rsc := (*e.IncludedClaimResourcesReferencedByClaimreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatientreference != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatientreference {
			rsc := (*e.IncludedPatientResourcesReferencedByPatientreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByProviderreference != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByProviderreference {
			rsc := (*e.IncludedPractitionerResourcesReferencedByProviderreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganizationreference != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByOrganizationreference {
			rsc := (*e.IncludedOrganizationResourcesReferencedByOrganizationreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedLocationResourcesReferencedByFacilityreference != nil {
		for idx := range *e.IncludedLocationResourcesReferencedByFacilityreference {
			rsc := (*e.IncludedLocationResourcesReferencedByFacilityreference)[idx]
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
	if e.RevIncludedPaymentNoticeResourcesReferencingRequestreference != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequestreference {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponsereference != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponsereference {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponsereference)[idx]
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
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
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
	if e.RevIncludedProcessResponseResourcesReferencingRequestreference != nil {
		for idx := range *e.RevIncludedProcessResponseResourcesReferencingRequestreference {
			rsc := (*e.RevIncludedProcessResponseResourcesReferencingRequestreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
